package main

import (
	//"encoding/hex"
	"errors"
	"fmt"

	"github.com/satindergrewal/kmdgo/btcec"
	"github.com/satindergrewal/kmdgo/chaincfg"
	"github.com/satindergrewal/kmdgo/kmdutil"
)

type Network struct {
	name        string
	symbol      string
	xpubkey     byte
	xprivatekey byte
}

var network = map[string]Network{
	"rdd": {name: "reddcoin", symbol: "rdd", xpubkey: 0x3d, xprivatekey: 0xbd},
	"dgb": {name: "digibyte", symbol: "dgb", xpubkey: 0x1e, xprivatekey: 0x80},
	"btc": {name: "bitcoin", symbol: "btc", xpubkey: 0x00, xprivatekey: 0x80},
	"ltc": {name: "litecoin", symbol: "ltc", xpubkey: 0x30, xprivatekey: 0xb0},
	"kmd": {name: "Komodo", symbol: "kmd", xpubkey: 0x3c, xprivatekey: 0xb4},
}

func (network Network) GetNetworkParams(display bool) *chaincfg.Params {
	networkParams := &chaincfg.MainNetParams
	networkParams.PubKeyHashAddrID = network.xpubkey
	networkParams.PrivateKeyID = network.xprivatekey

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
	//fmt.Println(wif)
	address, _ := network["kmd"].GetAddress(wif)
	fmt.Printf("Wif Key: %s\nAddress: %s\n\n", wif.String(), address.EncodeAddress())
}
