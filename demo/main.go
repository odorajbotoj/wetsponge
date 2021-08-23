package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/odorajbotoj/wetsponge"
)

func LO(s string) string {
	return "§" + s + "§l§o"
}
func Serve(conn *websocket.Conn) {
	fmt.Println("New Connection: ", time.Now().Unix())
	funcMap := make(map[string]func(wsconn *wetsponge.WsConn, s string))

	funcMap[":"] = func(wsconn *wetsponge.WsConn, s string) {
		err := wsconn.CommandRequest(s)
		if err != nil {
			fmt.Println(err)
		}
	}
	/*
		按需求使用
		***注意:使用该命令可以执行sudo, 这可能危害您的服务器!***
		funcMap["~$ "] = func(wsconn *wetsponge.WsConn, s string) {
			// Linux:
			cmd := exec.Command("/bin/bash", "-c", s)
			var pipe = wetsponge.SysInfoPipe{wsconn}
			cmd.Stdout = &pipe
			cmd.Stderr = &pipe
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
		}

		有严重bug
		funcMap["P> "] = func(wsconn *wetsponge.WsConn, s string) {
			// Linux
			ss := strings.Split(s, " ")
			cmd := exec.Command("./plugins/bin/painter-linux64", ss...)
			var pipe = wetsponge.CmdReqPipe{wsconn}
			var errpipe = wetsponge.SysInfoPipe{wsconn}
			cmd.Stdout = &pipe
			cmd.Stderr = &errpipe
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
		}
	*/
	funcMap["+"] = func(wsconn *wetsponge.WsConn, s string) {
		ss := strings.Split(s, " ")
		var cmd *exec.Cmd
		if len(ss) == 0 {
			wsconn.WriteError("没有传入数据!")
		} else if len(ss) == 1 {
			cmd = exec.Command("./plugins/bin" + ss[0])
		} else {
			cmd = exec.Command("./plugins/bin"+ss[0], ss[1:]...)
		}
		var pipe = wetsponge.CmdReqPipe{wsconn}
		var errpipe = wetsponge.SysInfoPipe{wsconn}
		cmd.Stdout = &pipe
		cmd.Stderr = &errpipe
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
	funcMap["&"] = func(wsconn *wetsponge.WsConn, s string) {
		ss := strings.Split(" ")
		if len(ss) != 2 {
			wsconn.WriteError("传入参数错误.")
		}
		mcfunc, err := os.Open("./plugins/functions/" + ss[0] + ".mcfunction")
		defer mcfunc.Close()
		if err != nil {
			wsconn.WriteError("找不到文件: " + ss[0] + ".mcfunction")
		}
		var pipe = wetsponge.CmdReqPipe{wsconn}
		bf := bufio.NewReader(mcfunc)
		for {
			l, _, err1 := bf.ReadLine()
			if err1 != nil {
				if err == io.EOF {
					break
				} else {
					wsconn.WriteError(fmt.Sprint(err1))
					break
				}
			}
			wl := strings.ReplaceAll(string(l), "@initiator", ss[1])
			wl = strings.ReplaceAll(wl, "@s", ss[1])
			_, err2 := pipe.Write([]byte(wl))
			if err2 != nil {
				wsconn.WriteError(fmt.Sprint(err2))
			}
		}
	}
	funcMap["builder"] = wetsponge.SlowBuilder
	/*
		由于系统适配问题, 不予执行
		funcMap["F& "] = func(wsconn *wetsponge.WsConn, s string) {
			s1 := strings.Replace(s, "%RF%", "./plugins/bin/runFunc", 1)
			s2 := strings.Replace(s1, "%FD%", "./plugins/functions", 1)
			cmd := exec.Command("/bin/bash", "-c", s2)
			var pipe = wetsponge.CmdReqPipe{wsconn}
			var errpipe = wetsponge.ErrInfoPipe{wsconn}
			out, Oerr := cmd.StdoutPipe()
			if Oerr != nil {
				wsconn.WriteError(fmt.Sprint(Oerr))
			}
			bf := bufio.NewReader(out)
			cmd.Stderr = &errpipe
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
			for {
				l, _, err1 := bf.ReadLine()
				if err1 != nil {
					if err1 == io.EOF {
						break
					} else {
						wsconn.WriteError(fmt.Sprint(err1))
						break
					}
				}
				_, err2 := pipe.Write(l)
				if err2 != nil {
					wsconn.WriteError(fmt.Sprint(err2))
				}
			}
		}

		funcMap["Stdin<"] = func(wsconn *wetsponge.WsConn, s string) {
			_, err := os.Stdin.Write([]byte(s))
			if err != nil {
				wsconn.WriteError(fmt.Sprint(err))
			}
		}
	*/
	funcMap["?"] = func(wsconn *wetsponge.WsConn, s string) {
		err := wsconn.CommandRequest(`tellraw @a {"rawtext":[{"text":"` + wetsponge.LIGHT_PURPLE + `|:| -> 让Websocket服务器执行命令\n|+| -> 执行在 ./plugins/bin 下的可执行文件\n|&| -> 执行 ./plugins/functions/ 下后缀为 .mcfunction 的文件\n|?| -> 无参数,打印帮助信息.` + `"}]}`)
		if err != nil {
			wsconn.WriteError(fmt.Sprint(err))
		}
	}

	wsconn := wetsponge.WsConn{conn, false, wetsponge.NewUUIDPool(), funcMap}
	wsconn.Subscribe("PlayerMessage")
	wsconn.CommandRequest(fmt.Sprintf(`tellraw @a { "rawtext" : [ { "text" : "%sw%se%st%ss%sp%so%sn%sg%se" } ] }`, LO("6"), LO("7"), LO("8"), LO("9"), LO("a"), LO("b"), LO("c"), LO("d"), LO("e")))
	wsconn.CommandRequest(fmt.Sprintf(`tellraw @a { "rawtext" : [ { "text" : "%s%s" } ] }`, wetsponge.GREEN, wetsponge.VERSION.GetInfo()))
	readCH := make(chan []byte)
	go wsconn.ReadMsg(readCH)
	go wsconn.ListenPMSG(readCH)
}
func main() {
	if len(os.Args) != 2 {
		panic(fmt.Errorf("Hmm~ Where is the PORT?"))
		return
	}
	fmt.Println("HELLO MINECRAFT_BE")
	fmt.Print("Version Description:\n")
	fmt.Println(wetsponge.VERSION.GetInfo())
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte(wetsponge.VERSION.GetInfo())) })
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
		return
	}
	uport := uint16(port)
	ws := &wetsponge.WswsS{uport, mux, &wetsponge.DefaultUpgrader, "/mcws", Serve}
	wsser := ws.GetSer()
	wsser.ListenAndServe()
}
