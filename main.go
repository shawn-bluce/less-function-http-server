package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"net"
	"os"
)

var bindHost = flag.String("bind", "0.0.0.0", "Binding address(only ipv4)")
var port = flag.String("port", "8080", "Listening port")
var logLevel = flag.String("loglevel", "Info", "Debug/Info/Error")

func welcomeMessage() {
	logrus.Info("Listening on " + *bindHost + ":" + *port)
}

func setLogConfig() {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logrus.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true

	if *logLevel == "Debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else if *logLevel == "Info" {
		logrus.SetLevel(logrus.InfoLevel)
	} else if *logLevel == "Error" {
		logrus.SetLevel(logrus.ErrorLevel)
	} else { // default
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func main() {
	flag.Parse()
	var listener net.Listener
	var err error

	setLogConfig()

	listener, err = net.Listen("tcp", *bindHost+":"+*port)
	if err != nil { // 拦截监听异常情况
		logrus.Errorln("Error listening:", err)
		os.Exit(1)
	}
	defer listener.Close() //向 defer 关键字传入的函数会在函数返回之前运行
	welcomeMessage()

	for {
		conn, err := listener.Accept() // 阻塞，开始等待连接

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
