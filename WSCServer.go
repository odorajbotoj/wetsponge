package wetsponge

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// 两个默认项
var DefaultPORT uint16 = 19134
var DefaultUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// WswsS结构体 存储开启服务的参数
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

// 用这个函数获得一个http.Server
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

type WsConn struct {
	Conn        *websocket.Conn
	WriteLock   bool
	UuidPool    *UUIDPool
	PMSGFuncMap map[string]func(wsconn *WsConn, s string)
}

func (wc *WsConn) WriteMsg(bs []byte) error {
	if wc.WriteLock {
		return fmt.Errorf("WRITER has been LOCKED")
	}
	ok := wc.Conn.WriteMessage(websocket.TextMessage, bs)
	if ok != nil {
		return ok
	}
	return nil
}
func (wc *WsConn) ReadMsg(ch chan []byte) {
	for {
		_, readMsg, Cerr := wc.Conn.ReadMessage()
		if Cerr != nil {
			fmt.Println("Error happened or Connection has been Closed: ", Cerr)
			return
		}
		if len(readMsg) == 0 {
			continue
		}
		ch <- readMsg
	}
}

// Listen PlayerMessage Event.
func (wc *WsConn) ListenPMSG(ch chan []byte) {
	for {
		MSGrecv := <-ch
		var pmsg RECVGameEvent
		err := json.Unmarshal(MSGrecv, &pmsg)
		if err != nil {
			continue
		}
		if pmsg.Body.EventName != "PlayerMessage" {
			continue
		}
		msgText := pmsg.Body.Properties.Message
		foundFlag := false
		for i, f := range wc.PMSGFuncMap {
			if strings.HasPrefix(msgText, i) {
				f(wc, strings.TrimPrefix(msgText, i))
				foundFlag = true
			}
		}
		if !foundFlag {
			wc.WriteError("Can't find command: " + msgText)
		}
	}

}

// 封装Subscribe方法
func (wc *WsConn) Subscribe(name string) error {
	id := wc.UuidPool.NewEvent(WsSerEvent{0, name, time.Now().Unix()})
	msg, err := MakeCmdReq(0, id, name)
	if err != nil {
		return err
	}
	if wc.WriteLock {
		return fmt.Errorf("WRITER has been LOCKED")
	}
	ok := wc.Conn.WriteMessage(websocket.TextMessage, msg)
	if ok != nil {
		return ok
	}
	return nil
}

// 封装Unsubscribe方法
func (wc *WsConn) Unsubscribe(name string) error {
	id := wc.UuidPool.NewEvent(WsSerEvent{1, name, time.Now().Unix()})
	msg, err := MakeCmdReq(1, id, name)
	if err != nil {
		return err
	}
	if wc.WriteLock {
		return fmt.Errorf("WRITER has been LOCKED")
	}
	ok := wc.Conn.WriteMessage(websocket.TextMessage, msg)
	if ok != nil {
		return ok
	}
	return nil
}

// 封装CommandRequest方法
func (wc *WsConn) CommandRequest(name string) error {
	id := wc.UuidPool.NewEvent(WsSerEvent{2, name, time.Now().Unix()})
	msg, err := MakeCmdReq(2, id, name)
	if err != nil {
		return err
	}
	if wc.WriteLock {
		return fmt.Errorf("WRITER has been LOCKED")
	}
	ok := wc.Conn.WriteMessage(websocket.TextMessage, msg)
	if ok != nil {
		return ok
	}
	return nil
}

// 封装WriteWarning方法, 向客户端发送警告
func (wc *WsConn) WriteWarning(msg string) error {
	id := wc.UuidPool.NewEvent(WsSerEvent{2, fmt.Sprintf("WARNING: %s", msg), time.Now().Unix()})
	msg = fmt.Sprintf("tellraw @a {\"rawtext\":[{\"text\":\"%s[WS Waring]%s\"}]}", YELLOW, msg)
	sendMsg, err := MakeCmdReq(2, id, msg)
	if err != nil {
		return err
	}
	ok := wc.Conn.WriteMessage(websocket.TextMessage, sendMsg)
	if ok != nil {
		return ok
	}
	return nil
}

// 封装WriteError方法, 向客户端发送警告
func (wc *WsConn) WriteError(msg string) error {
	id := wc.UuidPool.NewEvent(WsSerEvent{2, fmt.Sprintf("ERROR: %s", msg), time.Now().Unix()})
	msg = fmt.Sprintf("tellraw @a {\"rawtext\":[{\"text\":\"%s[WS Error]%s\"}]}", RED, msg)
	sendMsg, err := MakeCmdReq(2, id, msg)
	if err != nil {
		return err
	}
	ok := wc.Conn.WriteMessage(websocket.TextMessage, sendMsg)
	if ok != nil {
		return ok
	}
	return nil
}

// System -> Game, 将输入的信息传给游戏(tellraw), 实现io.Writer接口
func (wc *WsConn) Write(p []byte) (n int, err error) {
	msg := fmt.Sprintf("tellraw @a {\"rawtext\":[{\"text\":\"%s[WS HostSystem]%s\"}]}", AQUA, string(p))
	sendMsg, err := MakeCmdReq(2, NewUUID(), msg)
	if err != nil {
		return 0, err
	}
	ok := wc.Conn.WriteMessage(websocket.TextMessage, sendMsg)
	if ok != nil {
		return 0, ok
	}
	return len(p), nil
}
