// https://api.github.com/repos/veruscoin/veruscoin/releases
// https://api.github.com/repos/veruscoin/veruscoin/releases/latest

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"

	// "net/http"
	"github.com/valyala/fasthttp"
)

var dlLinuxArm64, dlLinuxAmd64, dlmacOS, dlWin64 string

func main() {

	client := &fasthttp.Client{}
	url := `https://api.github.com/repos/veruscoin/veruscoin/releases/latest`

	queryByte, err := json.Marshal(map[string]interface{}{
		"id":      "0",
		"jsonrpc": "1.0",
	})
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
	}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("GET")
	req.SetBody(queryByte)

	resp := fasthttp.AcquireResponse()
	client.Do(req, resp)

	bodyBytes := resp.Body()
	// println(string(bodyBytes))

	var res interface{}
	json.Unmarshal(bodyBytes, &res)

	assets := res.(map[string]interface{})["assets"].([]interface{})
	tag_name := res.(map[string]interface{})["tag_name"].(string)
	fmt.Println("tag_name:", tag_name)

	// fmt.Println(`assets -- `, len(assets))
	for i, _ := range assets {
		var arm64 = regexp.MustCompile("(?m)arm64.+$")
		arm64Line := arm64.FindString(assets[i].(map[string]interface{})["name"].(string))
		if arm64Line != "" {
			// fmt.Println(assets[i].(map[string]interface{})["name"])
			// fmt.Println("armLine -", armLine)
			// v1 := strings.TrimLeft(assets[i].(map[string]interface{})["name"].(string), `Verus-CLI-Linux-`)[1:]
			// v2 := strings.TrimRight(v1, `-arm64.tgz`)
			// fmt.Println("version:", v2)
			dlLinuxArm64 = assets[i].(map[string]interface{})["browser_download_url"].(string)
			// fmt.Println("browser_download_url:", assets[i].(map[string]interface{})["browser_download_url"])
		}
		var amd64 = regexp.MustCompile("(?m)x86_64.+$")
		amd64Line := amd64.FindString(assets[i].(map[string]interface{})["name"].(string))
		if amd64Line != "" {
			dlLinuxAmd64 = assets[i].(map[string]interface{})["browser_download_url"].(string)
		}
		var macOS = regexp.MustCompile("(?m)MacOS.+$")
		macOSLine := macOS.FindString(assets[i].(map[string]interface{})["name"].(string))
		if macOSLine != "" {
			dlmacOS = assets[i].(map[string]interface{})["browser_download_url"].(string)
		}
		var win64 = regexp.MustCompile("(?m)Windows.+$")
		win64Line := win64.FindString(assets[i].(map[string]interface{})["name"].(string))
		if win64Line != "" {
			dlWin64 = assets[i].(map[string]interface{})["browser_download_url"].(string)
		}
	}
	fmt.Println("dlLinuxArm64:", dlLinuxArm64)
	fmt.Println("dlLinuxAmd64:", dlLinuxAmd64)
	fmt.Println("dlmacOS:", dlmacOS)
	fmt.Println("dlWin64:", dlWin64)
}
