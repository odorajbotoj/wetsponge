package wetsponge

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type CmdReqPipe struct {
	*WsConn
}

// Command -> Game, 将输入的信息传给游戏并执行, 实现io.Writer接口
func (pipe *CmdReqPipe) Write(p []byte) (n int, err error) {
	sendMsg, err := MakeCmdReq(2, NewUUID(), string(p))
	if err != nil {
		return 0, err
	}
	ok := pipe.Conn.WriteMessage(websocket.TextMessage, sendMsg)
	if ok != nil {
		return 0, ok
	}
	return len(p), nil
}

type SysInfoPipe struct {
	*WsConn
}

// System -> Game, 将输入的信息传给游戏(tellraw), 实现io.Writer接口
func (pipe *SysInfoPipe) Write(p []byte) (n int, err error) {
	msg := fmt.Sprintf("tellraw @a {\"rawtext\":[{\"text\":\"%s[WS HostSystem]%s\"}]}", AQUA, string(p))
	sendMsg, err := MakeCmdReq(2, NewUUID(), msg)
	if err != nil {
		return 0, err
	}
	ok := pipe.Conn.WriteMessage(websocket.TextMessage, sendMsg)
	if ok != nil {
		return 0, ok
	}
	return len(p), nil
}

type ErrInfoPipe struct {
	*WsConn
}

// ErrInfo -> Game, 错误信息传给游戏(tellraw), 实现io.Writer接口
func (pipe *ErrInfoPipe) Write(p []byte) (n int, err error) {
	msg := fmt.Sprintf("tellraw @a {\"rawtext\":[{\"text\":\"%s[WS Error]%s\"}]}", RED, string(p))
	sendMsg, err := MakeCmdReq(2, NewUUID(), msg)
	if err != nil {
		return 0, err
	}
	ok := pipe.Conn.WriteMessage(websocket.TextMessage, sendMsg)
	if ok != nil {
		return 0, ok
	}
	return len(p), nil
}
