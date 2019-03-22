package main

import (
	"fmt"
	"encoding/hex"
)

func main() {
	rand32Bytes, err := hex.DecodeString("3b6778bcf53dec4a39c8ebdc39f1d0e1e9852ce7bfcc39cb625b418ed5c1da90")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\nrand32Bytes: ", rand32Bytes)
	fmt.Printf("\nrand32Bytes Type: %T\n", rand32Bytes)
	fmt.Println("\nrand32Bytes Length: ", len(rand32Bytes))


	//s := "hello world"
	s := "aff51dad774a1c612dc82e63f85f07b992b665836b0f0efbcb26ee679f4f4848"
	bdata := []byte(s)
	fmt.Println(bdata)

	fmt.Println([]byte("hello"))

	b2s := BytesToString(bdata)
	fmt.Println(b2s)
	fmt.Println(len(b2s))
}


func BytesToString(data []byte) string {
	return string(data[:])
}