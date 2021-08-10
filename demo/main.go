package main

import (
	core "github.com/OdorajBotoj/WetSponge"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
)

var PORT uint16 = 19134
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     checkOrigin,
}

type RecvEventBodyProperties struct {
	Message string `json:"Message,omitempty"`
}
type RecvEventBody struct {
	Properties    *RecvEventBodyProperties `json:"properties,omitempty"`
	StatusMessage string                   `json:"statusMessage,omitempty"`
}
type RecvEventHeader struct {
	MessagePurpose string `json:"messagePurpose,omitempty"`
}
type RecvEvent struct {
	Body   *RecvEventBody   `json:"body,omitempty"`
	Header *RecvEventHeader `json:"header,omitempty"`
}

func checkOrigin(r *http.Request) bool {
	return true
}
func fakeMsg(c *websocket.Conn, u string, s string) {
	msg, _ := core.MakeCmdReq(core.COMMANDREQUEST, core.UUID{}, fmt.Sprintf(`tellraw @a {"rawtext":[{"text":"<%s>%s"}]}`, u, s))
	c.WriteMessage(websocket.TextMessage, msg)
}
func fakeJoin(c *websocket.Conn, u string) {
	msg, _ := core.MakeCmdReq(core.COMMANDREQUEST, core.UUID{}, fmt.Sprintf(`tellraw @a {"rawtext":[{"text":"§e%s Join the game"}]}`, u))
	c.WriteMessage(websocket.TextMessage, msg)
}
func Aserf(conn *websocket.Conn) {
	UName := "WetSponge"
	// fmt.Printf("Test:%T\n", conn)
	msg, _ := core.MakeCmdReq(core.SUBSCRIBE, core.UUID{}, "PlayerMessage")
	conn.WriteMessage(websocket.TextMessage, msg)

	msg, _ = core.MakeCmdReq(core.COMMANDREQUEST, core.UUID{}, "say §e[Debug]§r Test1")
	conn.WriteMessage(websocket.TextMessage, msg)

	fakeJoin(conn, UName)

	msg, _ = core.MakeCmdReq(core.COMMANDREQUEST, core.UUID{}, "say §e[Debug]§r Test2")
	conn.WriteMessage(websocket.TextMessage, msg)

	msg, _ = core.MakeCmdReq(core.COMMANDREQUEST, core.UUID{}, "say Hello！You can type in §e$<Command>§r to let the server execute the command.")
	conn.WriteMessage(websocket.TextMessage, msg)

	CHAN := make(chan string)
	go func(conn *websocket.Conn) {
		for {
			_, readMsg, Cerr := conn.ReadMessage()
			if Cerr != nil {
				return
			}
			var recvJson RecvEvent
			if len(readMsg) == 0 {
				continue
			}
			err := json.Unmarshal(readMsg, &recvJson)
			if err != nil {
				fmt.Println(err)
				continue
			}
			var bodyMsg string = ""
			switch recvJson.Header.MessagePurpose {
			case "commandResponse":
				bodyMsg = recvJson.Body.StatusMessage
			case "event":
				bodyMsg = recvJson.Body.Properties.Message
			}
			if strings.HasPrefix(bodyMsg, "$") {
				CHAN <- bodyMsg
				fmt.Println(bodyMsg)
			}
		}
	}(conn)
	go func(conn *websocket.Conn) {
		for {
			msgRecv := <-CHAN
			s := strings.TrimPrefix(msgRecv, "$")
			msgSend, _ := core.MakeCmdReq(core.COMMANDREQUEST, core.UUID{}, s)
			ok := conn.WriteMessage(websocket.TextMessage, msgSend)
			if ok != nil {
				return
			}
		}
	}(conn)
}
func main() {
	fmt.Print("Version Description:\n")
	fmt.Println(core.VERSION.GetInfo())
	ws := new(core.WswsS)
	ws.SetPort(PORT)
	ws.SetUpg(&upgrader)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		header := r.Header
		for i, j := range header {
			w.Write([]byte(fmt.Sprintf("%s:%s\n", i, j)))
		}
	})
	ws.SetMux(mux)
	ws.SetHdl("mcws")
	ws.SetFunc(Aserf)
	wsser := ws.GetSer()
	wsser.ListenAndServe()
}
