package wetsponge

import (
	"strings"
)

/*
func SlowBuilder(wsconn *WsConn, s string) {
	var cmdPipe = CmdReqPipe{wsconn}
	if strings.HasPrefix(s, "void") {
		ss := strings.Split(s, " ")
		if len(ss) != 3 {
			wsconn.WriteError("参数不完整")
			return
		}
	}
}
*/
