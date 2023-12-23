package main

import (
	"gin_realword/server"
	_ "gin_realword/storage"
)

func main() {
	server.RunHTTPServer()
}
