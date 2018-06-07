package main

import (
	"github.com/saheienko/kube-api-proxy/pkg/proxy"
	"log"
)

func main() {
	prx := proxy.New("", "")
	log.Fatal(prx.Run())
}
