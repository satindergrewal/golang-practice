package main

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
)

// var logString string = ``

// var logString string = `3013947264 openrequest.4195714048 -> (0133f63a3d4ae8db7e9efe8b8702e10ecc9ef44901dd2321c92132889bc6656b4e)`
// var logString string = `3013947264 iambob.0 (PIRATE/KMD) channelapproved origid.3013947264 status.1`

// var logString string = `
// rel.KMD/KMD  openrequest 3013947264 status.0 (RBthCSgNLE3rwvAKce8JNSo7xDpxQEiRTX/zs1zqks0tergf6nk69evm6awte4xmhf2fd9epnv946vzlhfxztkls6a9lyfmuafda00krvkvj0xagp)
// 3013947264 openrequest.4195714048 -> (0133f63a3d4ae8db7e9efe8b8702e10ecc9ef44901dd2321c92132889bc6656b4e)
// 3013947264 iambob.0 (PIRATE/KMD) channelapproved origid.3013947264 status.1
// 3013947264 approvalid.1437135040 (0133f63a3d4ae8db7e9efe8b8702e10ecc9ef44901dd2321c92132889bc6656b4e)
// 3013947264 iambob.0 (PIRATE/KMD) incomingchannel status.2
// 3013947264 got txid.4a1e6502c5434f0c9abbce3c73f5bfb88c71dc1d2ceb2a0669168e31c6d7ef8b
// 3013947264: 0.10000000 KMD -> RQPZrM4yQaTZpEuoGmcGwzE4SaG2Tn9QiB, paymentid[0] 16445472
// 3013947264 iambob.0 (PIRATE/KMD) incomingpayment status.4
// 3013947264 alice waits for PIRATE.5a2166bcded6cf1dc8f3f7fd3dd0ead99b57a39bf491ca68951242bc6f080dcb to be in mempool (1.00000000 -> zs1zqks0tergf6nk69evm6awte4xmhf2fd9epnv946vzlhfxztkls6a9lyfmuafda00krvkvj0xagp)
// zs1zqks0tergf6nk69evm6awte4xmhf2fd9epnv946vzlhfxztkls6a9lyfmuafda00krvkvj0xagp received 1.00000000 vs 1.00000000
// 3013947264 SWAP COMPLETE <<<<<<<<<<<<<<<<
// 3013947264 paidid.1345205632
// 3013947264 iambob.0 (PIRATE/KMD) incomingfullypaid status.5
// 3013947264 closedid.851672672
// 3013947264 iambob.0 (PIRATE/KMD) incomingclose status.6
// alice 3013947264 10000000 4195714048 finished
// subatomic_channel_alice (KMD/KMD) 3013947264 3013947264 with 0.10000000 10000000
// initialized 43 messages, updated 156 out of total.156
// start subatomic_loop iambob.0 PIRATE -> KMD, 3013947264 10000000 4195714048
// sendtoaddress RQPZrM4yQaTZpEuoGmcGwzE4SaG2Tn9QiB 0.10000000 txid.(4a1e6502c5434f0c9abbce3c73f5bfb88c71dc1d2ceb2a0669168e31c6d7ef8b)
// dpow_broadcast.(completed/4a1e6502c5434f0c9abbce3c73f5bfb88c71dc1d2ceb2a0669168e31c6d7ef8b) [ ] 4a1e6502c5434f0c9abbce3c73f5bfb88c71dc1d2ceb2a0669168e31c6d7ef8b error.(-1)
// `

var logString string = `
rel.PIRATE/PIRATE  openrequest 4105997824 status.0 (RBthCSgNLE3rwvAKce8JNSo7xDpxQEiRTX/zs1zqks0tergf6nk69evm6awte4xmhf2fd9epnv946vzlhfxztkls6a9lyfmuafda00krvkvj0xagp)
4105997824 openrequest.3128973312 -> (0133f63a3d4ae8db7e9efe8b8702e10ecc9ef44901dd2321c92132889bc6656b4e)
4105997824 iambob.0 (KMD/PIRATE) channelapproved origid.4105997824 status.1
4105997824 approvalid.682507392 (0133f63a3d4ae8db7e9efe8b8702e10ecc9ef44901dd2321c92132889bc6656b4e)
4105997824 iambob.0 (KMD/PIRATE) incomingchannel status.2
z_sendmany.( PIRATE) from.(zs1zqks0tergf6nk69evm6awte4xmhf2fd9epnv946vzlhfxztkls6a9lyfmuafda00krvkvj0xagp) -> '[{"address":"zs1wq40g4wvrzc2eq9xw7wtstshgar68ash659eq20ellm5jeqsyfwe5qs3tex9l3mjnrj2yf34hw0","amount":1.50000000,"memo":"3431303539393738323420"}]'
4105997824: 1.50000000 PIRATE -> zs1wq40g4wvrzc2eq9xw7wtstshgar68ash659eq20ellm5jeqsyfwe5qs3tex9l3mjnrj2yf34hw0, paymentid[0] 576533696
4105997824 iambob.0 (KMD/PIRATE) incomingpayment status.4
4105997824 alice waits for KMD.a2e484460c6c0e003ead566072f23cda5b746dc443e6ec0268d89d773e25bc1a to be in mempool (0.10000000 -> RBthCSgNLE3rwvAKce8JNSo7xDpxQEiRTX)
RBthCSgNLE3rwvAKce8JNSo7xDpxQEiRTX received 0.10000000 vs 0.10000000
4105997824 SWAP COMPLETE <<<<<<<<<<<<<<<<
4105997824 paidid.2056384576
4105997824 iambob.0 (KMD/PIRATE) incomingfullypaid status.5
4105997824 closedid.1230418208
4105997824 iambob.0 (KMD/PIRATE) incomingclose status.6
alice 4105997824 150000000 3128973312 finished
subatomic_channel_alice (PIRATE/PIRATE) 4105997824 4105997824 with 1.50000000 150000000
initialized 42 messages, updated 151 out of total.151
start subatomic_loop iambob.0 KMD -> PIRATE, 4105997824 150000000 3128973312
z_sendmany.() -> opid.(opid-a8124036-d646-45d2-9639-731e363d480f)
dpow_broadcast.(completed/e90dfe302ecda41454d33b1a1fdf1d5d07a00894888cecfe1870df8a541e1f72) [ ] e90dfe302ecda41454d33b1a1fdf1d5d07a00894888cecfe1870df8a541e1f72 error.(-1)
`

// SwapStatus defines the data type to store filtered data and push to application in JSON format for UI side rendering.
type SwapStatus struct {
	State      string  `json:"state,omitempty"`
	StateID    string  `json:"state_id,omitempty"`
	Status     string  `json:"status,omitempty"`
	StateHash  string  `json:"state_hash,omitempty"`
	BaseTxID   string  `json:"base_txid,omitempty"`
	RelTxID    string  `json:"rel_txid,omitempty"`
	SwapID     string  `json:"swap_id,omitempty"`
	SendToAddr string  `json:"sendtoaddr,omitempty"`
	RecvAddr   string  `json:"recvaddr,omitempty"`
	Base       string  `json:"base,omitempty"`
	Rel        string  `json:"rel,omitempty"`
	BaseAmount float64 `json:"base_amount,omitempty"`
	RelAmount  float64 `json:"rel_amount,omitempty"`
}

//ZFrom to render the JSON data from "from." log entery coming from subatomic stdout
type ZFrom []struct {
	Address string  `json:"address,omitempty"`
	Amount  float64 `json:"amount,omitempty"`
	Memo    string  `json:"memo,omitempty"`
}

func main() {

	// fmt.Println(logString)

	var expOpenReq = regexp.MustCompile(`(?m)openrequest\..+$`)
	openReq := expOpenReq.FindString(logString)
	// fmt.Println(openReq)
	openReqSf := strings.Fields(openReq)
	// fmt.Println(openReqSf[2])

	if len(openReqSf) > 0 {
		// fmt.Printf("length of openReqSf is greater: %d\n", len(openReqSf))

		openReqHash := strings.ReplaceAll(openReqSf[2], "(", "")
		openReqHash = strings.ReplaceAll(openReqHash, ")", "")
		// fmt.Println(openReqHash)
		openReqID := strings.Split(openReqSf[0], ".")
		// fmt.Println(openReqID)
		// fmt.Println(openReqID[0])
		// fmt.Println(openReqID[1])
		// fmt.Println("Channel Open Request Sent:", openReqID[1])
		state0 := SwapStatus{
			Status:    "0",
			State:     openReqID[0],
			StateID:   openReqID[1],
			StateHash: openReqHash,
		}

		// fmt.Println("State 0:", state0)
		state0JSON, _ := json.Marshal(state0)
		fmt.Println("Channel Opened")
		fmt.Println("state0 JSON:", string(state0JSON))

	} else {
		fmt.Printf("length of openReqSf is lower: %d\n", len(openReqSf))
	}

	fmt.Println(`----`)
	var expChAprov = regexp.MustCompile(`(?m)channelapproved.+$`)
	chAprov := expChAprov.FindString(logString)
	// fmt.Println(chAprov)
	chAprovSf := strings.Fields(chAprov)
	if len(chAprovSf) > 0 {
		// fmt.Printf("length of chAprovSf is greater: %d\n", len(chAprovSf))

		// fmt.Println(chAprovSf[0])
		// fmt.Println(chAprovSf[1])
		// fmt.Println(chAprovSf[2])
		aprStatus := strings.Split(chAprovSf[2], ".")
		// fmt.Println(aprStatus[1])
		// fmt.Printf("Channel with ID approved with status %s\n", aprStatus[1])

		state1 := SwapStatus{
			Status: aprStatus[1],
			State:  chAprovSf[0],
		}

		// fmt.Println("state 1:", state1)
		state1JSON, _ := json.Marshal(state1)
		fmt.Println("Channel Approved")
		fmt.Println("state1 JSON:", string(state1JSON))
	} else {
		fmt.Printf("length of chAprovSf is lower: %d\n", len(chAprovSf))
	}

	// if len(aprovIDSf) > 0 {
	// 	// fmt.Printf("length of aprovIDSf is greater: %d\n", len(aprovIDSf))
	// } else {
	// 	fmt.Printf("length of aprovIDSf is lower: %d\n", len(aprovIDSf))
	// }

	fmt.Println(`----`)
	var expAprovID = regexp.MustCompile(`(?m)approvalid.+$`)
	aprovID := expAprovID.FindString(logString)
	// fmt.Println(aprovID)
	aprovIDSf := strings.Fields(aprovID)
	if len(aprovIDSf) > 0 {
		// fmt.Printf("length of aprovIDSf is greater: %d\n", len(aprovIDSf))

		// fmt.Println(aprovIDSf)
		// fmt.Println(aprovIDSf[0])
		aprovStatus := strings.Split(aprovIDSf[0], ".")
		// fmt.Println(aprovStatus[0])
		// fmt.Println(aprovStatus[1])
		// fmt.Println(aprovIDSf[1])

		state1 := SwapStatus{
			State:     aprovStatus[0],
			Status:    "1",
			StateHash: aprovIDSf[1],
		}

		// fmt.Println("state 1:", state1)
		state1JSON, _ := json.Marshal(state1)
		fmt.Println("Channel Approval ID")
		fmt.Println("state1 JSON:", string(state1JSON))
	} else {
		fmt.Printf("length of aprovIDSf is lower: %d\n", len(aprovIDSf))
	}

	fmt.Println(`----`)
	var expIncCh = regexp.MustCompile(`(?m)incomingchannel.+$`)
	incCh := expIncCh.FindString(logString)
	// fmt.Println(incCh)
	incChSf := strings.Fields(incCh)
	// fmt.Println(incChSf)
	if len(incChSf) > 0 {
		// fmt.Printf("length of incChSf is greater: %d\n", len(incChSf))

		// fmt.Println(incChSf[0])
		incChStatus := strings.Split(incChSf[1], ".")
		// fmt.Println(incChStatus[0])
		// fmt.Println(incChStatus[1])

		state2 := SwapStatus{
			State:  incChSf[0],
			Status: incChStatus[1],
		}

		// fmt.Println("state 1:", state2)
		state2JSON, _ := json.Marshal(state2)
		fmt.Println("Incoming Channel")
		fmt.Println("state2 JSON:", string(state2JSON))
	} else {
		fmt.Printf("length of incChSf is lower: %d\n", len(incChSf))
	}

	fmt.Println(`----`)
	var expGotTxID = regexp.MustCompile(`(?m)got txid.+$`)
	TxID := expGotTxID.FindString(logString)
	// fmt.Println(TxID)
	TxIDSf := strings.Fields(TxID)
	// fmt.Println(TxIDSf)
	if len(TxIDSf) > 0 {
		// fmt.Printf("length of TxIDSf is greater: %d\n", len(TxIDSf))

		// fmt.Println(TxIDSf[0])
		TxIDStatus := strings.Split(TxIDSf[1], ".")
		// fmt.Println(TxIDStatus[0])
		// fmt.Println(TxIDStatus[1])

		state3 := SwapStatus{
			State:    TxIDSf[0],
			BaseTxID: TxIDStatus[1],
			Status:   "3",
		}

		// fmt.Println("state 1:", state3)
		state3JSON, _ := json.Marshal(state3)
		fmt.Println("Recieved TxID")
		fmt.Println("state3 JSON:", string(state3JSON))
	} else {
		fmt.Printf("length of TxIDSf is lower: %d\n", len(TxIDSf))
	}

	fmt.Println(`----`)
	var expZFrom = regexp.MustCompile(`(?m)from..+$`)
	zFrom := expZFrom.FindString(logString)
	// fmt.Println(zFrom)
	zFromSf := strings.Fields(zFrom)
	// fmt.Println(zFromSf)

	if len(zFromSf) > 0 {
		// fmt.Printf("length of zFromSf is greater: %d\n", len(zFromSf))

		// fmt.Println(zFromSf[0])
		zFromSl := strings.Split(zFromSf[0], ".")
		zFromAddr := strings.ReplaceAll(zFromSl[1], "(", "")
		zFromAddr = strings.ReplaceAll(zFromAddr, ")", "")
		// fmt.Println(zFromAddr)
		// fmt.Println(zFromSf[2])
		zFromJSON := strings.ReplaceAll(zFromSf[2], "'", "")
		zFromJSON = strings.ReplaceAll(zFromJSON, "'", "")
		// fmt.Printf("%s\n", zFromJSON)
		var zj ZFrom
		err := json.Unmarshal([]byte(zFromJSON), &zj)
		if err != nil {
			log.Println(err)
		}
		// fmt.Println(zj[0].Address)
		// fmt.Println(zj[0].Amount)
		// fmt.Println(zj[0].Memo)

		state3 := SwapStatus{
			State:      "Sending Z Transaction",
			Status:     "3",
			SendToAddr: zj[0].Address,
			BaseAmount: zj[0].Amount,
		}

		// fmt.Println("state 1:", state3)
		state3JSON, _ := json.Marshal(state3)
		fmt.Println("Sending Z tx")
		fmt.Println("state3 JSON:", string(state3JSON))
	} else {
		fmt.Printf("length of zFromSf is lower: %d\n", len(zFromSf))
	}
}
