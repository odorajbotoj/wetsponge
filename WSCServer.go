package WetSponge

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
)

type WswsS struct {
	Port      uint16
	HttpMux   *http.ServeMux
	WsUpg     *websocket.Upgrader
	WsHandler string
	ServeFunc func(conn *websocket.Conn)
}

func (wswss *WswsS) SetPort(p uint16) {
	wswss.Port = p
}
func (wswss *WswsS) SetUpg(g *websocket.Upgrader) {
	wswss.WsUpg = g
}
func (wswss *WswsS) SetMux(m *http.ServeMux) {
	wswss.HttpMux = m
}
func (wswss *WswsS) SetHdl(h string) {
	if !strings.HasPrefix(h, "/") {
		h = "/" + h
	}
	wswss.WsHandler = h
}
func (wswss *WswsS) SetFunc(f func(conn *websocket.Conn)) {
	wswss.ServeFunc = f
}
func (wswss *WswsS) GetSer() *http.Server {
	var httpSer = new(http.Server)
	httpSer.Addr = fmt.Sprintf(":%d", wswss.Port)
	wswss.HttpMux.HandleFunc(wswss.WsHandler, func(w http.ResponseWriter, r *http.Request) {
		if websocket.IsWebSocketUpgrade(r) {
			conn, _ := wswss.WsUpg.Upgrade(w, r, w.Header())
			wswss.ServeFunc(conn)
		} else {
			w.Write([]byte(fmt.Sprintf("Please connect <Server address>:%d%s with Minecraft:Bedrock Edition", wswss.Port, wswss.WsHandler)))
		}
	})
	httpSer.Handler = wswss.HttpMux
	return httpSer
}
