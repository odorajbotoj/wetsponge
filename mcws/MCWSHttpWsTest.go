package main
import (
	"fmt"
	"net/http"
	"strings"
	"crypto/sha1"
	"encoding/base64"
	"io"
)
func Bytesf(s string, a ...interface{}) []byte {
	return []byte(fmt.Sprintf(s, a...))
}
func KeySum(s string) string {
	RecvKey := s + "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
	sha := sha1.New()
	io.Copy(sha, strings.NewReader(RecvKey))
	SendKey := base64.URLEncoding.EncodeToString([]byte(sha.Sum(nil)))
	return SendKey
}
func WSserveFunc(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	w.Header().Set("Connection", "Upgrade")
	w.Header().Set("Sec-WebSocket-Accept", KeySum(header["Sec-WebSocket-Key"][0]))
	w.Header().Set("Upgrade", "websocket")
}
func main() {
	fmt.Println("test.")
	mux := http.NewServeMux()
	mux.HandleFunc("/", WSserveFunc)
	srv := new(http.Server)
	srv.Addr = ":60000"
	srv.Handler = mux
	srv.ListenAndServe()
}