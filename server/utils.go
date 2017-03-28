package main

import (
        "fmt"
        "net"
)

// Return whether port is open or not
// TODO: Return service name as well 
// GOAL 
// PORT     STATE  SERVICE
// 8000/tcp closed http-alt
// I think openbsd-libc uses /etc/protocols and /etc/services
func port_scan(port int) bool {
	var closed bool
        conn, err := net.Dial("tcp", fmt.Sprintf("23.79.234.237:%d", port))
        if err != nil {
                fmt.Println("Port closed")
                fmt.Printf("%s\n", err)
		closed = true
        } else {
                fmt.Println("Port open")
                fmt.Printf("%s\n", conn)
		closed = false
        }
	return closed
}
