package main

import (
	"bytes"
	"fmt"
	"github.com/rainu/go-command-chain"
	"io/ioutil"
	"net/http"
)

func GetPrivateIp() {
	output := &bytes.Buffer{}

	err := cmdchain.Builder().
		Join("ipconfig").
		Join("findstr", "IPv4").
		Finalize().
		WithOutput(output).
		Run()
	if err != nil {
		panic(err)
	}
	var IPoutput = output.String()
	fmt.Println("\t * IPv4 :", IPoutput[33:50])
}

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
	fmt.Printf("About Private IP\n")
	GetPrivateIp()
}
