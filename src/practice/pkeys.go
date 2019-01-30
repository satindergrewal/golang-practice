package main

import (
	"encoding/hex"
	//"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
)

func main() {
	x := btcec.S256()
	fmt.Println(x)

	tmp_str := []byte(`hello world komodo platform light house parse code temp king quen`)

	dst_str := make([]byte, hex.EncodedLen(len(tmp_str)))
	hex.Encode(dst_str, tmp_str)

	fmt.Printf("%s\n", dst_str)

	tmp_hex := "04115c42e757b2efb7671c578530ec191a1" + "359381e6a71127a9d37c486fd30dae57e76dc58f693bd7e7010358ce6b165e483a29" + "21010db67ac11b1b51b651953d2"

	//pkBytes, err := hex.DecodeString(string(dst_str))
	pkBytes, err := hex.DecodeString(tmp_hex)
	if err != nil {
		fmt.Println(err)
	return
	}
	fmt.Println("pkBytes: ", pkBytes)
	fmt.Println("pkBytes Length: ", len(pkBytes))

	privKey, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), pkBytes)

	fmt.Println("privKey: ", privKey)
	fmt.Printf("privKey Type: %T\n", privKey)
	fmt.Println("pubKey: ", pubKey)
	fmt.Printf("pubKey Type: %T\n", pubKey)


	publicKey, err := btcec.ParsePubKey(pkBytes, btcec.S256())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("publicKey: ", *publicKey)
	fmt.Printf("publicKey: %T\n", publicKey)
}