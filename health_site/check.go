package main
import (
	"fmt"
	"net"
	"time"

)

func Check(destination string, port string) string {

	address := destination + ":" + port
	timeout := time.Duration(5*time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)

	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable \n err: %v", destination, err)

	} else {
		status = fmt.Sprintf("[UP] %v is reachable from %v to %v", destination,
	    conn.LocalAddr(), conn.RemoteAddr())
	}
	return status

} 