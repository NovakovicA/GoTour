package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var result string

	for i := 0; i < len(ip); i++ {
		result += strconv.Itoa(int(ip[i]))

		if i != len(ip) - 1 {
			result += "."
		}
	}

	return result
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
