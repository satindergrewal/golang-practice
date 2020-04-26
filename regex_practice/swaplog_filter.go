package main

import (
	"fmt"
	"regexp"
	"strings"
)

// var logString string = ``

// var logString string = `3013947264 openrequest.4195714048 -> (0133f63a3d4ae8db7e9efe8b8702e10ecc9ef44901dd2321c92132889bc6656b4e)`
// var logString string = `3013947264 iambob.0 (PIRATE/KMD) channelapproved origid.3013947264 status.1`

var logString string = `
rel.KMD/KMD  openrequest 3013947264 status.0 (RBthCSgNLE3rwvAKce8JNSo7xDpxQEiRTX/zs1zqks0tergf6nk69evm6awte4xmhf2fd9epnv946vzlhfxztkls6a9lyfmuafda00krvkvj0xagp)
3013947264 openrequest.4195714048 -> (0133f63a3d4ae8db7e9efe8b8702e10ecc9ef44901dd2321c92132889bc6656b4e)
3013947264 iambob.0 (PIRATE/KMD) channelapproved origid.3013947264 status.1
3013947264 approvalid.1437135040 (0133f63a3d4ae8db7e9efe8b8702e10ecc9ef44901dd2321c92132889bc6656b4e)
3013947264 iambob.0 (PIRATE/KMD) incomingchannel status.2
3013947264 got txid.4a1e6502c5434f0c9abbce3c73f5bfb88c71dc1d2ceb2a0669168e31c6d7ef8b
3013947264: 0.10000000 KMD -> RQPZrM4yQaTZpEuoGmcGwzE4SaG2Tn9QiB, paymentid[0] 16445472
3013947264 iambob.0 (PIRATE/KMD) incomingpayment status.4
3013947264 alice waits for PIRATE.5a2166bcded6cf1dc8f3f7fd3dd0ead99b57a39bf491ca68951242bc6f080dcb to be in mempool (1.00000000 -> zs1zqks0tergf6nk69evm6awte4xmhf2fd9epnv946vzlhfxztkls6a9lyfmuafda00krvkvj0xagp)
zs1zqks0tergf6nk69evm6awte4xmhf2fd9epnv946vzlhfxztkls6a9lyfmuafda00krvkvj0xagp received 1.00000000 vs 1.00000000
3013947264 SWAP COMPLETE <<<<<<<<<<<<<<<<
3013947264 paidid.1345205632
3013947264 iambob.0 (PIRATE/KMD) incomingfullypaid status.5
3013947264 closedid.851672672
3013947264 iambob.0 (PIRATE/KMD) incomingclose status.6
alice 3013947264 10000000 4195714048 finished
subatomic_channel_alice (KMD/KMD) 3013947264 3013947264 with 0.10000000 10000000
initialized 43 messages, updated 156 out of total.156
start subatomic_loop iambob.0 PIRATE -> KMD, 3013947264 10000000 4195714048
sendtoaddress RQPZrM4yQaTZpEuoGmcGwzE4SaG2Tn9QiB 0.10000000 txid.(4a1e6502c5434f0c9abbce3c73f5bfb88c71dc1d2ceb2a0669168e31c6d7ef8b)
dpow_broadcast.(completed/4a1e6502c5434f0c9abbce3c73f5bfb88c71dc1d2ceb2a0669168e31c6d7ef8b) [ ] 4a1e6502c5434f0c9abbce3c73f5bfb88c71dc1d2ceb2a0669168e31c6d7ef8b error.(-1)
`

// SwapStatus defines the data type to store filtered data and push to application in JSON format for UI side rendering.
type SwapStatus struct {
	State      string `json:"state,omitempty"`
	StateID    string `json:"state_id,omitempty"`
	Status     string `json:"status,omitempty"`
	StateHash  string `json:"state_hash,omitempty"`
	BaseTxID   string `json:"base_txid,omitempty"`
	RelTxID    string `json:"rel_txid,omitempty"`
	SwapID     string `json:"swap_id,omitempty"`
	SendToAddr string `json:"sendtoaddr,omitempty"`
	RecvAddr   string `json:"recvaddr,omitempty"`
	Base       string `json:"base,omitempty"`
	Rel        string `json:"rel,omitempty"`
	BaseAmount string `json:"base_amount,omitempty"`
	RelAmount  string `json:"rel_amount,omitempty"`
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

		fmt.Println("State 0:", state0)
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

		fmt.Println("state 1:", state1)
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

		fmt.Println("state 1:", state1)
	} else {
		fmt.Printf("length of aprovIDSf is lower: %d\n", len(aprovIDSf))
	}
}
