package main

import "github.com/sirupsen/logrus"

const HTTP_200 = "OK"
const HTTP_201 = "OK"
const HTTP_404 = "NOT FOUND"
const HTTP_405 = "Method Not Allowed"
const HTTP_500 = "INTERNAL ERROR"
const HTTP_505 = "HTTP Version Not Supported"
const HTTP_ERROR = "SUPPORT CODE NOT MATCH"

func buildResponse(success bool, statusCode string) string {
	if !success {
		httpMessage := HTTP_ERROR
		if statusCode == "404" {
			httpMessage = HTTP_404
		} else if statusCode == "405" {
			httpMessage = HTTP_405
		} else if statusCode == "500" {
			httpMessage = HTTP_500
		} else if statusCode == "505" {
			httpMessage = HTTP_505
		}
		if statusCode[0] == '3' {
			logrus.Info("HTTP/1.1 ", statusCode)
		} else if statusCode[0] == '4' {
			logrus.Warning("HTTP/1.1 ", statusCode)
		} else if statusCode[0] == '5' {
			logrus.Error("HTTP/1.1 ", statusCode)
		}
		return "HTTP/1.1 " + statusCode + " " + httpMessage
	}

	return "HTTP/1.1 200 OK\nAAA: 111\nBBB: 222\nCCC: 333\n\nhello, world"
}
