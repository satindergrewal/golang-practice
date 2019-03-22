package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	//"crypto/rand"
	//"crypto/ecdsa"
	//"crypto/elliptic"
	//"crypto/sha256"
	//"golang.org/x/crypto/ripemd160"
	"github.com/satindergrewal/kmdgo/btcec"
	"github.com/satindergrewal/kmdgo/chaincfg"
	"github.com/satindergrewal/kmdgo/kmdutil"
)

type Network struct {
	name           string
	symbol         string
	xpubkey        byte
	xprivatekey    byte
	scripthashaddr byte
}

var network = map[string]Network{
	"rdd": {name: "reddcoin", symbol: "rdd", xpubkey: 0x3d, xprivatekey: 0xbd},
	"dgb": {name: "digibyte", symbol: "dgb", xpubkey: 0x1e, xprivatekey: 0x80},
	"btc": {name: "bitcoin", symbol: "btc", xpubkey: 0x00, xprivatekey: 0x80},
	"ltc": {name: "litecoin", symbol: "ltc", xpubkey: 0x30, xprivatekey: 0xb0},
	"kmd": {name: "Komodo", symbol: "kmd", xpubkey: 0x3c, xprivatekey: 0xbc, scripthashaddr: 0x55},
}

func (network Network) GetNetworkParams(display bool) *chaincfg.Params {
	networkParams := &chaincfg.MainNetParams
	networkParams.PubKeyHashAddrID = network.xpubkey
	networkParams.PrivateKeyID = network.xprivatekey
	if network.symbol == `kmd` {
		networkParams.ScriptHashAddrID = network.scripthashaddr
	}

	if display == true {
		fmt.Println("\n~~~~~~~~")
		fmt.Println("COIN Name: ", network.name)
		//fmt.Println(networkParams)
		fmt.Println("Name: ", networkParams.Name)
		fmt.Println("Net: ", networkParams.Net)
		fmt.Println("PubKeyHashAddrID: ", networkParams.PubKeyHashAddrID)
		fmt.Println("ScriptHashAddrID: ", networkParams.ScriptHashAddrID)
		fmt.Println("PrivateKeyID: ", networkParams.PrivateKeyID)
		fmt.Println("~~~~~~~~\n")
	}

	return networkParams
}

func (network Network) CreatePrivateKey() (*kmdutil.WIF, error) {
	secret, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return nil, err
	}
	//fmt.Println("secret: ", secret)
	return kmdutil.NewWIF(secret, network.GetNetworkParams(false), true)
}

//func (network Network) ImportPrivateKey(secretHex string) (*kmdutil.WIF, error) { }

func (network Network) ImportWIF(wifStr string) (*kmdutil.WIF, error) {
	wif, err := kmdutil.DecodeWIF(wifStr)
	if err != nil {
		return nil, err
	}
	//fmt.Println(*wif)
	if !wif.IsForNet(network.GetNetworkParams(false)) {
		return nil, errors.New("The WIF string is not valid for the `" + network.name + "` network")
	}
	return wif, nil
}

func (network Network) GetAddress(wif *kmdutil.WIF) (*kmdutil.AddressPubKey, error) {
	//fmt.Println(wif.PrivKey.PubKey().SerializeCompressed())
	return kmdutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeCompressed(), network.GetNetworkParams(true))
}

func main() {
	fmt.Println("Starting the application...")

	//fmt.Println(sha256.New())
	//fmt.Println(ripemd160.New())

	wif, _ := network["kmd"].CreatePrivateKey()
	fmt.Println("wif: ", *wif)

	fmt.Printf("\n")

	fmt.Println("~~~~~~~")
	fmt.Println("wif PrivKey: ", wif.PrivKey)
	fmt.Println("wif PrivKey Serialize: ", wif.PrivKey.Serialize())
	fmt.Println("wif PrivKey Serialize Length: ", len(wif.PrivKey.Serialize()))
	fmt.Println("wif PubKey: ", wif.PrivKey.PubKey())
	fmt.Println("wif PubKey SerializeCompressed: ", wif.PrivKey.PubKey().SerializeCompressed())
	fmt.Println("wif PubKey SerializeCompressed Length: ", len(wif.PrivKey.PubKey().SerializeCompressed()))
	fmt.Println("wif PubKey SerializeUncompressed: ", wif.PrivKey.PubKey().SerializeUncompressed())
	fmt.Println("wif PubKey SerializeUncompressed Length: ", len(wif.PrivKey.PubKey().SerializeUncompressed()))
	fmt.Println("wif CompressPubKey: ", wif.CompressPubKey)
	fmt.Println("~~~~~~~")

	fmt.Printf("\n")

	fmt.Println("~~~~~~~")
	PrivKeyBytes := wif.PrivKey.Serialize()
	PrivKeyHex := make([]byte, hex.EncodedLen(len(PrivKeyBytes)))
	hex.Encode(PrivKeyHex, PrivKeyBytes)
	fmt.Printf("PrivKey Hex: %s\n", PrivKeyHex)

	PubKeyBytes := wif.PrivKey.PubKey().SerializeCompressed()
	PubKeyHex := make([]byte, hex.EncodedLen(len(PubKeyBytes)))
	hex.Encode(PubKeyHex, PubKeyBytes)
	fmt.Printf("PubKey Compressed Hex: %s\n", PubKeyHex)

	PubKeyBytesUn := wif.PrivKey.PubKey().SerializeUncompressed()
	PubKeyHexUn := make([]byte, hex.EncodedLen(len(PubKeyBytesUn)))
	hex.Encode(PubKeyHexUn, PubKeyBytesUn)
	fmt.Printf("PubKey Uncompressed Hex: %s\n", PubKeyHexUn)
	fmt.Println("~~~~~~~")

	fmt.Printf("\n")
	address, _ := network["kmd"].GetAddress(wif)
	fmt.Printf("Wif Key: %s\nAddress: %s\n\n", wif.String(), address.EncodeAddress())

	/*
		// Get secp256k1 hash
		curve := btcec.S256()

		// Get Random number
		randread := rand.Reader

		fmt.Println("\nsecp256k1: ", curve)
		fmt.Println("\nRandom: ", randread)

		key, _ := ecdsa.GenerateKey(curve, randread)
		fmt.Println("Private Key: ", *key)
		fmt.Printf("Private Key Type: %T\n\n", *key)
		fmt.Println("Public Key: ", key.PublicKey)
		fmt.Printf("Public Key Type: %T\n\n", key.PublicKey)
		fmt.Printf("D: %s\nD Type: %T\n\n", key.D, key.D)


		type WIF struct {
		// PrivKey is the private key being imported or exported.
			PrivKey *ecdsa.PrivateKey

			// CompressPubKey specifies whether the address controlled by the
			// imported or exported private key was created by hashing a
			// compressed (33-byte) serialized public key, rather than an
			// uncompressed (65-byte) one.
			CompressPubKey bool

			// netID is the bitcoin network identifier byte used when
			// WIF encoding the private key.
			netID byte
		}

		_wif := &WIF{key, true, 0xbc}
		fmt.Println(*_wif)
	*/

	//fmt.Printf("\n")

	//fmt.Println("~~~~~~~")
	//fmt.Println("_wif PrivKey: ", _wif.PrivKey)
	//fmt.Println("_wif PrivKey Serialize: ", _wif.PrivKey.Serialize())
	//fmt.Println("_wif PrivKey Serialize Length: ", len(_wif.PrivKey.Serialize()))
	//fmt.Println("_wif PubKey: ", _wif.PrivKey.PubKey())
	//fmt.Println("_wif PubKey SerializeCompressed: ", _wif.PrivKey.PubKey().SerializeCompressed())
	//fmt.Println("_wif PubKey SerializeCompressed Length: ", len(_wif.PrivKey.PubKey().SerializeCompressed()))
	//fmt.Println("_wif PubKey SerializeUncompressed: ", _wif.PrivKey.PubKey().SerializeUncompressed())
	//fmt.Println("_wif PubKey SerializeUncompressed Length: ", len(_wif.PrivKey.PubKey().SerializeUncompressed()))
	//fmt.Println("_wif CompressPubKey: ", _wif.CompressPubKey)
	//fmt.Println("_wif netID: ", _wif)
	//fmt.Println("~~~~~~~")

	//fmt.Printf("0x%x\n", 188)
	//fmt.Printf("0x%x\n", 60)
	//fmt.Printf("0x%x\n", 85)

	//fmt.Printf("\n")

	//fmt.Println("~~~~~~~")

	//hsh := kmdutil.Hash160(wif.PrivKey.Serialize())
	//fmt.Println(hsh)

	//h := sha256.New()
	//h.Write([]byte("satinder"))
	//fmt.Printf("SHA256: %x\n", h.Sum(nil))

	//RPdW5oL5icgDEA9gfDMecwvcwKzreLDHrH
}
