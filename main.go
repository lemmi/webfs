// Package main provides a small http server with dir listing
package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"path/filepath"
)

type webfsConfig struct {
	net   string
	laddr string
	dir   string
}

func parseArgs() webfsConfig {
	var config webfsConfig
	config.dir = "./"
	flag.StringVar(&config.net, "net", "tcp", "specify the network (tcp, tcp4, tcp6, unix, unixpacket)")
	flag.StringVar(&config.laddr, "laddr", ":8008", "specify host and port")
	flag.Parse()
	if flag.NArg() > 0 {
		config.dir = filepath.Clean(flag.Args()[0])
	}
	return config
}

func main() {
	config := parseArgs()
	log.Printf("Starting listing of %s on %s...\n", config.dir, config.net)
	l, err := net.Listen(config.net, config.laddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.Serve(l, http.FileServer(http.Dir(config.dir))))
}
