package main

import (
	"encoding/hex"
	"errors"
	"fmt"
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

func (network Network) WifFromHex(pkey *btcec.PrivateKey, b bool) (*kmdutil.WIF, error) {
	secret := pkey
	fmt.Println("secret: ", secret)
	return kmdutil.NewWIF(secret, network.GetNetworkParams(b), true)
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

func (network Network) GetAddress(wif *kmdutil.WIF, b bool) (*kmdutil.AddressPubKey, error) {
	//fmt.Println(wif.PrivKey.PubKey().SerializeCompressed())
	return kmdutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeCompressed(), network.GetNetworkParams(b))
}

func main() {
	//tmp_hex := "04115c42e757b2efb7671c578530ec191a1359381e6a71127a9d37c486fd30dae57e76dc58f693bd7e7010358ce6b165e483a2921010db67ac11b1b51b651953d2"
	//tmp_hex := "0481ac2e2cb36247af4668e1f5f03e87f7fc58507d21e863f63a176e643579ed8b638d39a07be407b04f0b1fddd7c0f291d9d55d37549e6f67b5a79f85d94a148d"
	tmp_hex := "0381ac2e2cb36247af4668e1f5f03e87f7fc58507d21e863f63a176e643579ed8b"
	fmt.Println("tmp_hex: ", tmp_hex)

	pkBytes, err := hex.DecodeString(tmp_hex)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("pkBytes: ", pkBytes)
	fmt.Println("pkBytes Length: ", len(pkBytes))
	fmt.Println("~~~~~~~")

	fmt.Printf("\n")

	fmt.Println("~~~~~~~")

	privKey, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), pkBytes)

	fmt.Println("privKey: ", privKey)
	fmt.Printf("privKey Type: %T\n", privKey)
	fmt.Println("~~~~~~~")

	fmt.Printf("\n")

	fmt.Println("~~~~~~~")
	fmt.Println("pubKey: ", pubKey)
	fmt.Printf("pubKey Type: %T\n", pubKey)

	publicKey, err := btcec.ParsePubKey(pkBytes, btcec.S256())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("~~~~~~~")

	fmt.Printf("\n")

	fmt.Println("~~~~~~~")
	fmt.Println("publicKey: ", *publicKey)
	fmt.Printf("publicKey: %T\n", publicKey)

	fmt.Printf("\n")

	type WIF struct {
		PrivKey        *btcec.PrivateKey
		CompressPubKey bool
		netID          byte
	}

	wif := &WIF{privKey, false, 0xbc}
	fmt.Println("wif: ", wif)

	fmt.Println("~~~~~~~")
	fmt.Println("wif PrivKey: ", privKey)
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

	fmt.Println("~~~~~~~")
	wif2, _ := network["kmd"].WifFromHex(privKey, false)
	//wif2 := kmdutil.NewWIF(secret, network.GetNetworkParams(false)
	//wif2 := &kmdutil.WIF{privKey, false, 0xbc}
	fmt.Println("wif2: ", wif2)

	fmt.Printf("\n")
	address, _ := network["kmd"].GetAddress(wif2, false)
	fmt.Printf("Wif Key: %s\nAddress: %s\n\n", wif2.String(), address.EncodeAddress())

}
