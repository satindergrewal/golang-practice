// https://api.github.com/repos/veruscoin/veruscoin/releases
// https://api.github.com/repos/veruscoin/veruscoin/releases/latest

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"runtime"

	"github.com/valyala/fasthttp"
	// "net/http"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	// gi, _ := goInfo.GetInfo()
	out, err := GetDlURL("linux", runtime.GOARCH)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("out:", out)
}

func GetDlURL(str ...string) (string, error) {
	var dlLinuxArm64, dlLinuxAmd64, dlmacOS, dlWin64 string
	fmt.Println(str)

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
	if len(bodyBytes) != 0 {
		fmt.Println("bodyBytes len:", len(bodyBytes))
		// fmt.Println("bodyBytes:", string(bodyBytes))

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
		if len(str) != 0 {
			switch str[0] {
			case "darwin":
				return dlmacOS, nil
			case "linux":
				switch str[1] {
				case "x86_64":
					return dlLinuxAmd64, nil
				case "amd64":
					return dlLinuxAmd64, nil
				case "arm64":
					return dlLinuxArm64, nil
				}
			case "windows":
				return dlWin64, nil
			}
		} else {
			switch runtime.GOOS {
			case "darwin":
				return dlmacOS, nil
			case "linux":
				switch runtime.GOARCH {
				case "x86_64":
					return dlLinuxAmd64, nil
				case "amd64":
					return dlLinuxAmd64, nil
				case "arm64":
					return dlLinuxArm64, nil
				}
			case "windows":
				return dlWin64, nil
			}
		}
	} else {
		return "", errors.New("downloads are unreachable")
	}
	return "", errors.New("something went wrong processing this request")
}
