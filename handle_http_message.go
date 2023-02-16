package main

import (
	"github.com/sirupsen/logrus"
	"net"
	"strings"
)

func analyzeHttpMessage(conn net.Conn) {
	httpMessageWithByte := make([]byte, 4096)
	conn.Read(httpMessageWithByte)
	httpMessage := string(httpMessageWithByte)
	v := strings.Split(httpMessage, "\n")
	for index, line := range v {
		if index == 0 { // head line
			headLine := strings.Split(line, " ")
			method := headLine[0]
			path := headLine[1]
			version := headLine[2]
			logrus.Info("========== HTTP HEAD LINE ==========")
			logrus.Info("Method: ", method)
			logrus.Info("PATH: ", path)
			logrus.Info("Version: ", version)
		} else {
			logrus.Info("xxxxxxxxxxx")
		}
	}
}
