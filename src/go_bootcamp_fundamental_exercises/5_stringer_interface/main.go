package main

import "fmt"
type IPAddr [4]byte

// Implementing  fmt.Stringer in the function
func (ip IPAddr) String() string {
	fmt.Println("String method", )
	res := fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
	return res
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

