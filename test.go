// import (
// 	"fmt"
// 	"net"
// )

// func main() {
// 	listener, err := net.Listen("tcp", "localhost:1200")
// 	if err != nil{
// 		fmt.Println("Error :", err)
// 	}
// 	defer listener.Close()

// 	fmt.Println("Server listening on 1200...")
// }

package main

import (
	"fmt"
	"net"
)

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			fmt.Println("Wi-Fi IP Address:", ipNet.IP.String())
		}
	}
}
