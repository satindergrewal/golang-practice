package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

var MARKETS_AVAILABLE = map[string][]string{}

var fixer_data = `{
	"success":true,
	"timestamp":1635570135,
	"base":"EUR",
	"date":"2021-10-30",
	"rates":{
	  "AED":4.246413,
	  "AFN":104.801184,
	  "ALL":121.680284,
	  "AMD":552.390515,
	  "ANG":2.084152,
	  "AOA":690.193783,
	  "ARS":115.25928,
	  "AUD":1.536962,
	  "AWG":2.081563,
	  "AZN":1.969964,
	  "BAM":1.940818,
	  "BBD":2.33484,
	  "BDT":99.025359,
	  "BGN":1.956213,
	  "BHD":0.435882,
	  "BIF":2309.89335,
	  "BMD":1.156103,
	  "BND":1.556289,
	  "BOB":7.973129,
	  "BRL":6.515569,
	  "BSD":1.156356,
	  "BTC":1.8794063e-5,
	  "BTN":86.600562,
	  "BWP":13.177971,
	  "BYN":2.804012,
	  "BYR":22659.61444,
	  "BZD":2.330871,
	  "CAD":1.432238,
	  "CDF":2326.079207,
	  "CHF":1.058645,
	  "CLF":0.034084,
	  "CLP":940.494132,
	  "CNY":7.405652,
	  "COP":4346.946444,
	  "CRC":737.446003,
	  "CUC":1.156103,
	  "CUP":30.636724,
	  "CVE":109.487438,
	  "CZK":25.652422,
	  "DJF":205.463041,
	  "DKK":7.439757,
	  "DOP":65.267829,
	  "DZD":159.145006,
	  "EGP":18.164887,
	  "ERN":17.342974,
	  "ETB":54.510699,
	  "EUR":1,
	  "FJD":2.405146,
	  "FKP":0.84768,
	  "GBP":0.844796,
	  "GEL":3.642176,
	  "GGP":0.84768,
	  "GHS":7.041117,
	  "GIP":0.84768,
	  "GMD":60.117756,
	  "GNF":11139.050673,
	  "GTQ":8.950271,
	  "GYD":242.057732,
	  "HKD":8.992139,
	  "HNL":27.955014,
	  "HRK":7.503556,
	  "HTG":116.2151,
	  "HUF":359.978617,
	  "IDR":16449.666175,
	  "ILS":3.656927,
	  "IMP":0.84768,
	  "INR":86.625105,
	  "IQD":1687.910055,
	  "IRR":48845.342758,
	  "ISK":150.004782,
	  "JEP":0.84768,
	  "JMD":177.768,
	  "JOD":0.819723,
	  "JPY":131.820041,
	  "KES":128.563026,
	  "KGS":98.04191,
	  "KHR":4699.558196,
	  "KMF":487.186099,
	  "KPW":1040.492153,
	  "KRW":1357.970327,
	  "KWD":0.348762,
	  "KYD":0.963646,
	  "KZT":494.858092,
	  "LAK":11861.614901,
	  "LBP":1771.149859,
	  "LKR":233.59388,
	  "LRD":175.153935,
	  "LSL":17.480715,
	  "LTL":3.413671,
	  "LVL":0.699315,
	  "LYD":5.260708,
	  "MAD":10.437879,
	  "MDL":20.166417,
	  "MGA":4581.0616,
	  "MKD":61.144178,
	  "MMK":2075.749315,
	  "MNT":3295.961672,
	  "MOP":9.262913,
	  "MRO":412.728493,
	  "MUR":49.779689,
	  "MVR":17.862225,
	  "MWK":942.224163,
	  "MXN":23.766872,
	  "MYR":4.787466,
	  "MZN":73.794477,
	  "NAD":17.48071,
	  "NGN":474.360967,
	  "NIO":40.699124,
	  "NOK":9.760259,
	  "NPR":138.566014,
	  "NZD":1.612405,
	  "OMR":0.445167,
	  "PAB":1.156356,
	  "PEN":4.598403,
	  "PGK":4.086868,
	  "PHP":58.428635,
	  "PKR":199.023526,
	  "PLN":4.609618,
	  "PYG":7994.242733,
	  "QAR":4.209415,
	  "RON":4.947546,
	  "RSD":116.676567,
	  "RUB":82.01382,
	  "RWF":1156.102778,
	  "SAR":4.336429,
	  "SBD":9.28602,
	  "SCR":17.036894,
	  "SDG":509.267526,
	  "SEK":9.932385,
	  "SGD":1.559398,
	  "SHP":1.59242,
	  "SLL":12457.00782,
	  "SOS":675.164413,
	  "SRD":24.864346,
	  "STD":23928.993333,
	  "SVC":10.117741,
	  "SYP":1453.188848,
	  "SZL":17.480701,
	  "THB":38.479769,
	  "TJS":13.00321,
	  "TMT":4.05214,
	  "TND":3.251544,
	  "TOP":2.582098,
	  "TRY":11.098229,
	  "TTD":7.838471,
	  "TWD":32.17631,
	  "TZS":2664.81729,
	  "UAH":30.360666,
	  "UGX":4109.712203,
	  "USD":1.156103,
	  "UYU":50.528917,
	  "UZS":12358.739079,
	  "VEF":247209713966.38382,
	  "VND":26303.072343,
	  "VUV":129.826218,
	  "WST":2.979269,
	  "XAF":650.939009,
	  "XAG":0.04838,
	  "XAU":0.000648,
	  "XCD":3.124426,
	  "XDR":0.818093,
	  "XOF":653.780285,
	  "XPF":118.504696,
	  "YER":289.315141,
	  "ZAR":17.624614,
	  "ZMK":10406.316475,
	  "ZMW":19.964677,
	  "ZWL":372.264623
	}
  }`

type fixerRates struct {
	Success   bool        `json:"success"`
	Timestamp float64     `json:"timestamp"`
	Base      string      `json:"base"`
	Date      string      `json:"date"`
	Rates     interface{} `json:"rates"`
}

func main() {

	symbols := `USD,EUR,INR`
	s := strings.Split(symbols, ",")
	// fmt.Println(s[0])
	// fmt.Println(s[1])
	// fmt.Println(s[2])
	// fmt.Println(len(s))

	var result interface{}
	json.Unmarshal([]byte(fixer_data), &result)
	// fmt.Printf("%T\n", result)
	// fmt.Println(result.(map[string]interface{})["rates"].(map[string]interface{})["AED"].(float64))

	var rates []map[string]float64

	var fx fixerRates
	for _, symbol := range s {
		// fmt.Println(symbol)
		for i, v := range result.(map[string]interface{})["rates"].(map[string]interface{}) {
			if strings.Compare(i, symbol) == 0 {
				// fmt.Println("i -", i)
				// fmt.Println("v -", v)
				rates = append(rates, map[string]float64{i: v.(float64)})
			}
		}
	}
	fmt.Println(rates)
	fx.Rates = rates
	fmt.Println(fx)

	// b, _ := json.MarshalIndent(fx, "", "  ")
	// fmt.Println(string(b))

	// MARKETS_AVAILABLE["binance"] = append(MARKETS_AVAILABLE["binance"], "BTC-USD")

	if _, ok := MARKETS_AVAILABLE["binance"]; ok {
		fmt.Println("element found")
	} else {
		fmt.Println("element not found")
	}

	// for _, v := range MARKETS_AVAILABLE["binance"] {
	// 	// fmt.Println(i)
	// 	// fmt.Println(v)
	// 	if v != "BTC-USDT" {
	// 		MARKETS_AVAILABLE["binance"] = append(MARKETS_AVAILABLE["binance"], "BTC-USDT")
	// 	}
	// }

	// if _, ok := set[1]; ok {
	// 	fmt.Println("element found")
	// } else {
	// 	fmt.Println("element not found")
	// }

	// markets, _ := json.Marshal(MARKETS_AVAILABLE)
	// fmt.Println(string(markets))

	var x map[string]interface{}
	xType := fmt.Sprintf("%T", x)
	if xType != "[]interface {}" {
		fmt.Println("not type []interface{}")
		fmt.Println("xType value:", xType)
		fmt.Printf("xTyle type: %T\n", xType)
	} else {
		fmt.Println("type is []interface{}")
		fmt.Println("xType value:", xType)
		fmt.Printf("xTyle type: %T\n", xType)
	}
}