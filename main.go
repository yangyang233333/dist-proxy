package main

import (
	"github.com/yangyang233333/dist-proxy/src"
	_ "github.com/yangyang233333/dist-proxy/src"
)

func main() {
	logger := src.GetLogInstance()
	logger.Info("aaaaaaaaaaaaaa")

}
