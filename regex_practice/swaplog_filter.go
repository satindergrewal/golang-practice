package main

import (
	"fmt"
	"regexp"
	"strings"
)

var log_string string = `
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

func BytesToString(data []byte) string {
	return string(data[:])
}

func main() {

	// fmt.Println(log_string)

	var expOpenReq = regexp.MustCompile(`(?m)openrequest\..+$`)
	openReq := expOpenReq.FindString(log_string)
	fmt.Println(openReq)
	openReqID := strings.TrimLeft(openReq, `openrequest`)[1:10]
	fmt.Println(openReqID)

	fmt.Println(`----`)
	var expChAprov = regexp.MustCompile(`(?m)channelapproved.+$`)
	chAprov := expChAprov.FindString(log_string)
	fmt.Println(chAprov)
}
