package main

import (
	"github.com/rgdaddio/gobotics/server"
)

func main() {
	server.Serve() // TODO have the package return a new instance of the server
			// s = server.new server
			// s.start()
			// select {} // this is a forever
}
