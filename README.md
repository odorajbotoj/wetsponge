# WetSponge
A simple websocket server for Minecraft Bedrock Edition (and Education Edition)
* (Written with Go)
* It's really a simple server :smile:
* I'm just a chinese middle school student, and I'm really not good at English...
## Update infomation
* 我不想写英文了
* 原先的包名是WetSpongeCore，提交的时候没注意，现在改成WetSponge了。
* 由于我可能会做一些不兼容的更改，1.0.0中的TODO等计划将在2.0.0实现（暂定）。
## It depends on ...
+ fmt
+ encoding/json
+ github.com/gorilla/websocket
+ time
+ net/http
+ strings
+ github.com/satori/go.uuid
+ and your love (not LOVE :D)
## How to use it
+ First, install it.
``` shell
$ go get github.com/OdorajBotoj/WetSponge
```
+ you may need:
``` shell
$ go get github.com/gorilla/websocket
$ go get github.com/satori/go.uuid
```
+ Then, import it.
``` go
import (
	"github.com/OdorajBotoj/WetSponge"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)
```
+ Er, that's all...?No.
+ Use it:
``` go
// you need a upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     checkOrigin,
}
func checkOrigin(r *http.Request) bool {
	return true
}

// this is a DEMO
	ws := new(WetSponge.WswsS)
	ws.SetPort(19134)
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
```
+ Check "demo/main.go" for more infomation...
###### Have a nice day ;P
