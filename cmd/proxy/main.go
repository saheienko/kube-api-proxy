package main

import (
	"log"

	"github.com/saheienko/kube-api-proxy/pkg/proxy"
)

func main() {
	prx := proxy.New("", "")
	log.Fatal(prx.Run())
}
