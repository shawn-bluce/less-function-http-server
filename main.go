package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"net"
	"os"
)

var host = flag.String("host", "0.0.0.0", "host")
var port = flag.String("port", "2333", "port")

func welcomeMessage() {
	logrus.Info("Listening on " + *host + ":" + *port)
}

func main() {
	flag.Parse()
	var l net.Listener
	var err error

	l, err = net.Listen("tcp", *host+":"+*port)
	if err != nil { // 拦截监听异常情况
		logrus.Errorln("Error listening:", err)
		os.Exit(1)
	}
	defer l.Close() //向 defer 关键字传入的函数会在函数返回之前运行
	welcomeMessage()

	for {
		conn, err := l.Accept() // 阻塞，开始等待连接

		// 连接异常处理
		if err != nil {
			logrus.Errorln("Error accepting: ", err)
			os.Exit(1)
		}

		// 输出内容，以表示接到了来自remoteAddr->localAddr的连接
		logrus.Info("connection: ", conn.RemoteAddr(), " -> ", conn.LocalAddr())

		// 将连接转交给handle函数
		go handleRequest(conn)
	}
}
func handleRequest(conn net.Conn) {
	defer conn.Close()

	analyzeHttpMessage(conn)

	responseMessage := buildResponse()
	conn.Write([]byte(responseMessage))
}
