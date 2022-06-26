package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetIp() string {
	ipify, ipapi := http.Get("https://api.ipify.org")
	if ipapi != nil {
		panic(ipapi)
	}
	defer ipify.Body.Close()
	ipaddr, addrerr := ioutil.ReadAll(ipify.Body)
	if addrerr != nil {
		panic(addrerr)
	}
	return string(ipaddr)

}
func GetIp64() string {
	ipify, ipapi := http.Get("https://api64.ipify.org")
	if ipapi != nil {
		panic(ipapi)
	}
	defer ipify.Body.Close()
	ipaddr, addrerr := ioutil.ReadAll(ipify.Body)
	if addrerr != nil {
		panic(addrerr)
	}
	return string(ipaddr)
}
func main() {
	fmt.Printf("About Public Internet Protocol\n")
	fmt.Printf("\t * IPv4 : %s\n", string(GetIp()))
	if GetIp() != GetIp64() {
		fmt.Printf("\t * IPv6 : %s\n", string(GetIp64()))
	} else {
		fmt.Printf("\t * IPv6 : Not activated\n")
	}
}
