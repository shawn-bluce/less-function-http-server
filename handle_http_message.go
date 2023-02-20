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
	httpMessageList := strings.Split(httpMessage, "\r\n")

	contentStartLine := 0
	for index, line := range httpMessageList {
		if index == 0 { // star line
			headLine := strings.Split(line, " ")
			method := headLine[0]
			path := headLine[1]
			version := headLine[2]
			logrus.Info("========== HTTP STAR LINE ==========")
			logrus.Info("Method: ", method)
			logrus.Info("PATH: ", path)
			logrus.Info("Version: ", version)
			logrus.Info("========== HTTP STAR LINE END ==========")
			logrus.Info("========== HTTP HEADER ==========")
		} else if line == "" { // http message blank line
			logrus.Info("========== HTTP HEADER END ==========")
			contentStartLine = index
			break
		} else { // http header
			keyValue := strings.Split(line, ":")
			headerKey := keyValue[0]
			headerValue := strings.Join(keyValue[1:len(keyValue)], ":")
			logrus.Info(headerKey, ": ", headerValue)
		}
	}
	logrus.Info("========== CONTENT START ==========")
	content := strings.Join(httpMessageList[contentStartLine:len(httpMessageList)], "\n")
	logrus.Info(content)
	logrus.Info("========== CONTENT END ==========")
}
