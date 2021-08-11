
package main
/*
import (
	"fmt"
	"net"
	"crypto/sha1"
	"encoding/base64"
	"strings"
	"io"
)
func Bytesf(s string, a ...interface{}) []byte {
	return []byte(fmt.Sprintf(s, a...))
}
const DefaultPort uint16 = 19134
const KeyAccept string = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
func serve(c net.Conn) {
	data := make([]byte, 0)
	c.Read(data)
	Sdata := string(data)
	fmt.Println(Sdata)
	Hdata := strings.Split(Sdata, "\r\n")[0]
	Key := strings.Split(Hdata ,"Sec-WebSocket-Key: ")[1]
	RecvKey := strings.Split(Key, "\r\n")[0]
	RecvKey = RecvKey + KeyAccept
	sha := sha1.New()
	io.Copy(sha, strings.NewReader(RecvKey))
	SendKey := base64.URLEncoding.EncodeToString([]byte(sha.Sum(nil)))
	c.Write(Bytesf("Connection: Upgrade\r\nSec-WebSocket-Accept: %s\r\nUpgrade: websocket\r\n\r\n", SendKey))
}
func main() {
lnr, err := net.Listen("tcp", fmt.Sprintf(":%d", DefaultPort))
if err != nil {
	fmt.Println(err)
	return
}
for {
	conn ,err := lnr.Accept()
	if err != nil {
		continue
	}
	go serve(conn)
}
}
*/