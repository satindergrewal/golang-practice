package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	//"crypto/elliptic"
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
	/*"rdd": {name: "reddcoin", symbol: "rdd", xpubkey: 0x3d, xprivatekey: 0xbd},
	"dgb": {name: "digibyte", symbol: "dgb", xpubkey: 0x1e, xprivatekey: 0x80},
	"btc": {name: "bitcoin", symbol: "btc", xpubkey: 0x00, xprivatekey: 0x80},
	"ltc": {name: "litecoin", symbol: "ltc", xpubkey: 0x30, xprivatekey: 0xb0},*/
	"kmd": {name: "Komodo", symbol: "kmd", xpubkey: 0x3c, xprivatekey: 0xbc, scripthashaddr: 0x55},
}

func (network Network) GetNetworkParams(display bool) *chaincfg.Params {
	networkParams := &chaincfg.MainNetParams
	networkParams.PubKeyHashAddrID = network.xpubkey
	networkParams.PrivateKeyID = network.xprivatekey
	networkParams.ScriptHashAddrID = network.scripthashaddr

	if display == true {
		fmt.Println("\n~~~~~~~~")
		fmt.Println("COIN Name: ", network.name)
		//fmt.Println(networkParams)
		fmt.Println("Name: ", networkParams.Name)
		fmt.Println("Net: ", networkParams.Net)
		//fmt.Println("DefaultPort: ", networkParams.DefaultPort)
		//fmt.Println("DNSSeeds: ", networkParams.DNSSeeds)
		//fmt.Println("GenesisBlock: ", networkParams.GenesisBlock)
		//fmt.Println("GenesisHash: ", networkParams.GenesisHash)
		//fmt.Println("PowLimit: ", networkParams.PowLimit)
		//fmt.Println("PowLimitBits: ", networkParams.PowLimitBits)
		//fmt.Println("BIP0034Height: ", networkParams.BIP0034Height)
		//fmt.Println("BIP0065Height: ", networkParams.BIP0065Height)
		//fmt.Println("BIP0066Height: ", networkParams.BIP0066Height)
		//fmt.Println("CoinbaseMaturity: ", networkParams.CoinbaseMaturity)
		//fmt.Println("SubsidyReductionInterval: ", networkParams.SubsidyReductionInterval)
		//fmt.Println("TargetTimespan: ", networkParams.TargetTimespan)
		//fmt.Println("TargetTimePerBlock: ", networkParams.TargetTimePerBlock)
		//fmt.Println("RetargetAdjustmentFactor: ", networkParams.RetargetAdjustmentFactor)
		//fmt.Println("ReduceMinDifficulty: ", networkParams.ReduceMinDifficulty)
		//fmt.Println("MinDiffReductionTime: ", networkParams.MinDiffReductionTime)
		//fmt.Println("GenerateSupported: ", networkParams.GenerateSupported)
		//fmt.Println("Checkpoints: ", networkParams.Checkpoints)
		//fmt.Println("RuleChangeActivationThreshold: ", networkParams.RuleChangeActivationThreshold)
		//fmt.Println("MinerConfirmationWindow: ", networkParams.MinerConfirmationWindow)
		//fmt.Println("Deployments: ", networkParams.Deployments)
		//fmt.Println("RelayNonStdTxs: ", networkParams.RelayNonStdTxs)
		//fmt.Println("Bech32HRPSegwit: ", networkParams.Bech32HRPSegwit)
		fmt.Println("PubKeyHashAddrID: ", networkParams.PubKeyHashAddrID)
		fmt.Println("ScriptHashAddrID: ", networkParams.ScriptHashAddrID)
		fmt.Println("PrivateKeyID: ", networkParams.PrivateKeyID)
		//fmt.Println("WitnessPubKeyHashAddrID: ", networkParams.WitnessPubKeyHashAddrID)
		//fmt.Println("WitnessScriptHashAddrID: ", networkParams.WitnessScriptHashAddrID)
		//fmt.Println("HDPrivateKeyID: ", networkParams.HDPrivateKeyID)
		//fmt.Println("HDPublicKeyID: ", networkParams.HDPublicKeyID)
		//fmt.Println("HDCoinType: ", networkParams.HDCoinType)
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
	wif, _ := network["kmd"].CreatePrivateKey()

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
	fmt.Println("~~~~~~~")

	fmt.Printf("\n")

	fmt.Println("~~~~~~~")
	PrivKeyBytes := wif.PrivKey.Serialize()
	PrivKeyHex := make([]byte, hex.EncodedLen(len(PrivKeyBytes)))
	hex.Encode(PrivKeyHex, PrivKeyBytes)

	PubKeyBytes := wif.PrivKey.PubKey().SerializeCompressed()
	PubKeyHex := make([]byte, hex.EncodedLen(len(PubKeyBytes)))
	hex.Encode(PubKeyHex, PubKeyBytes)

	PubKeyBytesUn := wif.PrivKey.PubKey().SerializeUncompressed()
	PubKeyHexUn := make([]byte, hex.EncodedLen(len(PubKeyBytesUn)))
	hex.Encode(PubKeyHexUn, PubKeyBytesUn)

	fmt.Printf("PrivKey Hex: %s\n", PrivKeyHex)
	fmt.Printf("PubKey Compressed Hex: %s\n", PubKeyHex)
	fmt.Printf("PubKey Uncompressed Hex: %s\n", PubKeyHexUn)
	fmt.Println("~~~~~~~")

	//fmt.Printf("\n")

	address, _ := network["kmd"].GetAddress(wif)
	fmt.Printf("Wif Key: %s\nAddress: %s\n\n", wif.String(), address.EncodeAddress())

	curve := btcec.S256()
	randread := rand.Reader
	fmt.Println("\nSHA256: ", curve)
	fmt.Println("\nRandom: ", randread)

	key, _ := ecdsa.GenerateKey(curve, randread)
	/*if err != nil {
		return nil, err
	}*/
	fmt.Println("Private Key: ", *key)
	fmt.Printf("Private Key Type: %T\n\n", *key)
	fmt.Println("Public Key: ", key.PublicKey)
	fmt.Printf("Public Key Type: %T\n\n", key.PublicKey)
	fmt.Printf("D: %s\nD Type: %T\n\n", key.D, key.D)

	fmt.Printf("0x%x\n", 188)
	fmt.Printf("0x%x\n", 60)
	fmt.Printf("0x%x\n", 85)
}
