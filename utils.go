package main

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"os"
)

func readRouters() string {
	file, err := os.Open("routers.txt")
	if err != nil {
		logrus.Error("open routers.txt error", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fileContent := ""

	for scanner.Scan() {
		fileContent = fileContent + scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		logrus.Error("read routers.txt error：", err)
		os.Exit(1)
	}

	return fileContent
}
