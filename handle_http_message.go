package main

import (
	"github.com/sirupsen/logrus"
	"net"
	"regexp"
	strings "strings"
)

func analyzeHttpMessage(conn net.Conn, routers string) (bool, string) {
	httpMessageWithByte := make([]byte, 4096)
	conn.Read(httpMessageWithByte)
	httpMessage := string(httpMessageWithByte)
	httpMessageList := strings.Split(httpMessage, "\r\n")

	contentStartLine := 0
	for index, line := range httpMessageList {
		if index == 0 { // star line: method, path, version
			headLine := strings.Split(line, " ")
			method := headLine[0]
			path := headLine[1]
			version := headLine[2]

			if version != "HTTP/1.0" && version != "HTTP/1.1" {
				return false, "505"
			}

			// range all the routers config
			for _, line := range strings.Split(routers, "\n") {
				pathMatched := false
				if line != "" { // exclude blank line
					res := strings.Split(line, " ")
					confMethod, confPath := res[0], res[1]
					pathCompile, err := regexp.Compile(confPath)
					if err != nil {
						return false, "500"
					}

					if pathCompile.Match([]byte(path)) {
						pathMatched = true
						if confMethod == method {
							logrus.Info("Math path: ", confPath)
							return true, ""
						}
					}
				}
				if pathMatched {
					return false, "405"
				}
			}
			return false, "404"

		} else if line == "" { // http message blank line
			contentStartLine = index
			break
		} else { // http header
			keyValue := strings.Split(line, ":")
			headerKey := keyValue[0]
			headerValue := strings.Join(keyValue[1:len(keyValue)], ":")
			logrus.Debug("HTTP Request Header -> ", headerKey, ": ", headerValue)
		}
	}
	logrus.Debug("========== CONTENT START ==========")
	content := strings.Join(httpMessageList[contentStartLine:len(httpMessageList)], "\n")
	logrus.Debug(content)
	logrus.Debug("========== CONTENT END ==========")
	return true, ""
}
