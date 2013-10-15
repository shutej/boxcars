package main

import (
	"flag"
	"fmt"
	"github.com/shutej/boxcars"
	"os"
)

var (
	filename string
	port     int
	userId   int
	groupId  int
	certFile string
	keyFile  string
	secure   bool
)

func main() {
	flag.IntVar(&port, "port", 8080, "Port to listen")
	flag.BoolVar(&secure, "secure", false, "Enables secure mode to avoid running as sudo.")
	flag.IntVar(&userId, "uid", 1000, "User id that'll own the system process.")
	flag.IntVar(&groupId, "gid", 1000, "Group id that'll own the system process.")
	flag.StringVar(&certFile, "cert_file", "", "Path to the certificate file.")
	flag.StringVar(&keyFile, "key_file", "", "Path to the key file.")
	flag.Parse()

	filename = flag.Arg(0)

	if filename == "" {
		fmt.Printf("Usage: boxcars config.json\n")
		os.Exit(1)
	}

	useTLS := false
	if certFile != "" || keyFile != "" {
		if certFile == "" {
			fmt.Printf("Must specify --cert_file when --key_file is specified.")
		}
		if keyFile == "" {
			fmt.Printf("Must specify --key_file when --cert_file is specified.")
		}
		useTLS = true
	}

	boxcars.SetFilename(filename)
	go boxcars.ReadConfig()
	go boxcars.AutoReload()

	if secure {
		go boxcars.Secure(userId, groupId)
	}

	if useTLS {
		boxcars.ListenTLS(port, certFile, keyFile)
	} else {
		boxcars.Listen(port)
	}
}
