package main

import (
	_ "fmt"
	"storyapi/cmd"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.WithFields(
		logrus.Fields{
			"Function": "Main()",
		}).Info("App : Begin")
	cmd.Begin()
}
