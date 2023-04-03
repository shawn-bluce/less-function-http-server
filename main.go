package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"net"
	"os"
)

// define CLI arguments
var bindHost = flag.String("bind", "0.0.0.0", "Binding address(only ipv4)")
var port = flag.String("port", "8080", "Listening port")
var logLevel = flag.String("loglevel", "Info", "Debug/Info/Error")

func setLogConfig() {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logrus.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true

	targetLoglevel := logrus.InfoLevel
	if *logLevel == "Debug" {
		targetLoglevel = logrus.DebugLevel
	} else if *logLevel == "Info" {
		targetLoglevel = logrus.InfoLevel
	} else if *logLevel == "Error" {
		targetLoglevel = logrus.ErrorLevel
	}
	logrus.SetLevel(targetLoglevel)
	logrus.Debug("Set loglevel ", targetLoglevel)
}

func main() {
	flag.Parse()
	var listener net.Listener
	var err error

	COUNTER := 0

	setLogConfig()
	routers := readRouters()

	listener, err = net.Listen("tcp", *bindHost+":"+*port)
	if err != nil { // listen bindHost:port error
		logrus.Error("Error listening:", err)
		os.Exit(1)
	}
	defer listener.Close()

	logrus.Info("Listening on " + *bindHost + ":" + *port)

	for {
		conn, err := listener.Accept() // waiting for connection

		// 监听步骤出现异常
		if err != nil {
			logrus.Errorln("Error accepting: ", err)
			os.Exit(1)
		}

		// 输出内容，以表示接到了来自remoteAddr->localAddr的连接
		COUNTER++
		logrus.Info("connection ", COUNTER, ": ", conn.RemoteAddr(), " -> ", conn.LocalAddr())

		// 将连接转交给handle函数
		go handleRequest(conn, routers)
	}
}

func handleRequest(conn net.Conn, routers string) {
	defer conn.Close()

	success, statusCode := analyzeHttpMessage(conn, routers)

	responseMessage := buildResponse(success, statusCode)
	conn.Write([]byte(responseMessage))
}
