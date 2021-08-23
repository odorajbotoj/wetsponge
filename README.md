# WetSponge
A simple websocket server for Minecraft Bedrock Edition (and Education Edition)
* (Written with Go)
* It's really a simple server :smile:
* I'm just a chinese middle school student, and I'm really not good at English...
## Update infomation
* 现在是2.0.0-dev3
* 但这是一个稳定版
* 我要上初三了,所以会很久不动这个仓库
* 今年是2021年,我们四年后见
## It depends on ...
+ github.com/gorilla/websocket
+ STL, and your love (not LOVE :D)
## 我已经对go的包管理工具失去信心了...所以wetsponge不再作为一个库发布,而是一个软件
+ 好吧,如果你一定要把它作为一个库来使用,那也不要紧
+ go get github.com/odorajbotoj/wetsponge
+ cd /path/to/wetsponge/demo (不会有人直接输入/path/to吧?)
+ go build main.go
+ ./main 19134 (19134可以换成你想开启的端口号)
## 用法
+ 首先开一个http Server Mux
+ 然后你可以进行一些常规配置(针对http Server)
+ 获取一个WswsS实例
```go
uport := uint16(19134)
ws := &wetsponge.WswsS{uport, mux, &wetsponge.DefaultUpgrader, "/mcws", Serve}
// 其中,uport为服务端口号,mux为先前创建的路由,DefaultUpgrader是websocket upgrader(详见gorilla/websocket中的upgrader)
// "/mcws"是开启服务的地址(如上代码开启的服务需要在<ip address>:19134/mcws连接)
// Serve是一个函数
// func Serve(conn *websocket.Conn)
// 关于这个函数还请自行看demo    ;P
```
+ 接下来调用ws.GetSer(),该函数返回一个http.Server
+ 接下来就可以ListenAndServe了
###### Have a nice day ;P
