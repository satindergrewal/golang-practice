package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

func main() {
	x := getIP()
	fmt.Printf("detected public IP: %s\n", x)
	ips, err := net.LookupIP("dev.khoji.io")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		os.Exit(1)
	}
	for _, ip := range ips {
		fmt.Printf("dev.khoji.io: %s\n", ip.String())
		fmt.Printf("Public IP (%v) == dev.khoji.io IP (%v): %v\n", x, ip, isIPEqual(net.IP(x), ip))
	}
}

type IP struct {
	Query string
}

func getIP() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}

	// fmt.Println(string(body))

	var ip IP
	json.Unmarshal(body, &ip)

	return ip.Query
}

func isIPEqual(ip1, ip2 net.IP) bool {

	isEqual := ip1.Equal(ip2)

	return isEqual

}
