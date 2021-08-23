package wetsponge

import (
	"crypto/rand"
)

func NewUUID() string {
	// 生成随机数
	b := make([]byte, 30)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	ret := "" // 8-4-4-4-12
	for i, j := range b {
		// 在特定位置插入特定字符
		if i == 8 || i == 20 {
			ret += "-"
		}
		if i == 12 {
			ret += "-4"
			continue
		}
		if i == 16 {
			ret += "-8"
			continue
		}
		// 用取余来获得随机字符
		switch j % 16 {
		case 0:
			ret += "0"
		case 1:
			ret += "1"
		case 2:
			ret += "2"
		case 3:
			ret += "3"
		case 4:
			ret += "4"
		case 5:
			ret += "5"
		case 6:
			ret += "6"
		case 7:
			ret += "7"
		case 8:
			ret += "8"
		case 9:
			ret += "9"
		case 10:
			ret += "a"
		case 11:
			ret += "b"
		case 12:
			ret += "c"
		case 13:
			ret += "d"
		case 14:
			ret += "e"
		case 15:
			ret += "f"
		}
	}

	return ret
}

// WsSerEvent结构体 服务器事件信息
type WsSerEvent struct {
	Type      uint8
	Name      string
	TimeStamp int64
}

// UUIDPool结构体 存储生成的uuid及请求信息,以防uuid重复.
type UUIDPool struct {
	Pool map[string]WsSerEvent
}

// 由这个方法生成的UUID会自动检测重复,并将信息存入Pool.
func (up *UUIDPool) NewEvent(wse WsSerEvent) string {
	var id string
	for {
		id = NewUUID()
		_, ok := up.Pool[id]
		if !ok {
			up.Pool[id] = wse
			break
		} else {
			continue
		}
	}
	return id
}

// 新建UUIDPool, 初始化并返回
func NewUUIDPool() *UUIDPool {
	var pool UUIDPool
	pool.Pool = make(map[string]WsSerEvent)
	return &pool
}
