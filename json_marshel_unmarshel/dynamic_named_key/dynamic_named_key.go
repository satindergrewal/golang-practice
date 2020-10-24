package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var data = `[
	{
	  "VRSC-BTC-ETH-USD": {
		"name": "VRSC-BTC-ETH-USD",
		"version": 1,
		"options": 97,
		"parent": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
		"systemid": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
		"currencyid": "iFiiKcz41hzxaGXvBmiB44ZHdZ1eT1BiA9",
		"notarizationprotocol": 1,
		"proofprotocol": 1,
		"idregistrationprice": 0,
		"idreferrallevels": 0,
		"minnotariesconfirm": 0,
		"billingperiod": 0,
		"notarizationreward": 0,
		"startblock": 147,
		"endblock": 0,
		"currencies": [
		  "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
		  "iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3",
		  "iHk5GXF4XU2AodPk4KwqwHxbfkaLpdJ56K",
		  "iFawzbS99RqGs7J2TNxME1TmmayBGuRkA2"
		],
		"weights": [
		  0.25000000,
		  0.25000000,
		  0.25000000,
		  0.25000000
		],
		"conversions": [
		  0.00000000,
		  0.00000000,
		  0.00000000,
		  0.00000000
		],
		"initialsupply": 4000000.00000000,
		"prelaunchcarveout": 0.00000000,
		"initialcontributions": [
		  1000000.00000000,
		  94.13377439,
		  2893.76965247,
		  1000000.00000000
		],
		"preconversions": [
		  1000000.00000000,
		  94.13377439,
		  2893.76965247,
		  1000000.00000000
		],
		"eras": [
		]
	  },
	  "lastnotarization": {
		"version": 1,
		"currencyid": "iFiiKcz41hzxaGXvBmiB44ZHdZ1eT1BiA9",
		"notaryaddress": "iFiiKcz41hzxaGXvBmiB44ZHdZ1eT1BiA9",
		"notarizationheight": 25985,
		"mmrroot": "9d6e2844e0ec2ed7d47443ef9faeee1e625a3e85fe30a4df461b3cb0f355e22a",
		"notarizationprehash": "a519ecb213f5d0616febce696ebad8987786fa9abd52b11a6da717350eb681c2",
		"work": "00000000000000000000000000000000000000000000000000000000172b3a06",
		"stake": "00000000000000000000000000000000000000000000000000002a884679ae47",
		"currencystate": {
		  "flags": 3,
		  "currencyid": "iFiiKcz41hzxaGXvBmiB44ZHdZ1eT1BiA9",
		  "reservecurrencies": [
			{
			  "currencyid": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
			  "weight": 0.25000000,
			  "reserves": 983676.98969257,
			  "priceinreserve": 0.98367698
			},
			{
			  "currencyid": "iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3",
			  "weight": 0.25000000,
			  "reserves": 95.60656616,
			  "priceinreserve": 0.00009560
			},
			{
			  "currencyid": "iHk5GXF4XU2AodPk4KwqwHxbfkaLpdJ56K",
			  "weight": 0.25000000,
			  "reserves": 2894.49327578,
			  "priceinreserve": 0.00289449
			},
			{
			  "currencyid": "iFawzbS99RqGs7J2TNxME1TmmayBGuRkA2",
			  "weight": 0.25000000,
			  "reserves": 1000682.99455323,
			  "priceinreserve": 1.00068299
			}
		  ],
		  "initialsupply": 4000000.00000000,
		  "emitted": 0.00000000,
		  "supply": 4000000.00000000,
		  "currencies": {
			"iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq": {
			  "reservein": 0.00000000,
			  "nativein": 0.00000000,
			  "reserveout": 4.72812000,
			  "lastconversionprice": 0.98368171,
			  "viaconversionprice": 0.97895361,
			  "fees": 0.00000000,
			  "conversionfees": 0.00000000
			},
			"iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3": {
			  "reservein": 1.82444378,
			  "nativein": 0.00000000,
			  "reserveout": 0.00000000,
			  "lastconversionprice": 0.00009446,
			  "viaconversionprice": 0.00009514,
			  "fees": 0.00045622,
			  "conversionfees": 0.00091244
			},
			"iHk5GXF4XU2AodPk4KwqwHxbfkaLpdJ56K": {
			  "reservein": 0.00000000,
			  "nativein": 0.00000000,
			  "reserveout": 0.00000000,
			  "lastconversionprice": 0.00289449,
			  "viaconversionprice": 0.00288058,
			  "fees": 0.00000000,
			  "conversionfees": 0.00000000
			},
			"iFawzbS99RqGs7J2TNxME1TmmayBGuRkA2": {
			  "reservein": 0.00000000,
			  "nativein": 0.00000000,
			  "reserveout": 19463.09271239,
			  "lastconversionprice": 1.02014608,
			  "viaconversionprice": 1.00794752,
			  "fees": 0.00000000,
			  "conversionfees": 0.00000000
			}
		  },
		  "nativefees": 472812000,
		  "nativeconversionfees": 0
		},
		"prevnotarization": "354598e9fedbc0b320cfd49a0d718560cfd0094629771ab1c411b8eb710821f6",
		"prevheight": 24778,
		"crossnotarization": "0000000000000000000000000000000000000000000000000000000000000000",
		"crossheight": 0,
		"nodes": [
		]
	  }
	},
	{
	  "VRSC-BTC-ETH-KMD": {
		"name": "VRSC-BTC-ETH-KMD",
		"version": 1,
		"options": 97,
		"parent": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
		"systemid": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
		"currencyid": "i8JRDSbog7jBEpbeVd1ib3huTYuNVYeXSy",
		"notarizationprotocol": 1,
		"proofprotocol": 1,
		"idregistrationprice": 0,
		"idreferrallevels": 0,
		"minnotariesconfirm": 0,
		"billingperiod": 0,
		"notarizationreward": 0,
		"startblock": 147,
		"endblock": 0,
		"currencies": [
		  "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
		  "iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3",
		  "iHk5GXF4XU2AodPk4KwqwHxbfkaLpdJ56K",
		  "i8xcvrfKTw8eEiTMjt1dAQxpWcbZeaVMFV"
		],
		"weights": [
		  0.25000000,
		  0.25000000,
		  0.25000000,
		  0.25000000
		],
		"conversions": [
		  0.00000000,
		  0.00000000,
		  0.00000000,
		  0.00000000
		],
		"initialsupply": 4000000.00000000,
		"prelaunchcarveout": 0.00000000,
		"initialcontributions": [
		  1000000.00000000,
		  94.13377439,
		  2893.76965247,
		  250000.00000000
		],
		"preconversions": [
		  1000000.00000000,
		  94.13377439,
		  2893.76965247,
		  250000.00000000
		],
		"eras": [
		]
	  },
	  "lastnotarization": {
		"version": 1,
		"currencyid": "i8JRDSbog7jBEpbeVd1ib3huTYuNVYeXSy",
		"notaryaddress": "i8JRDSbog7jBEpbeVd1ib3huTYuNVYeXSy",
		"notarizationheight": 30439,
		"mmrroot": "e3d89fbe4857f54c63e03dc63c0be8c135c2421a38d174a6173ad6f8a456bfc4",
		"notarizationprehash": "767636b20f78c6bdf1c828d1b6081bf099bcf5b3ae11bb458afe3c515173d1ab",
		"work": "000000000000000000000000000000000000000000000000000000001f41ebe0",
		"stake": "0000000000000000000000000000000000000000000000000000000000000000",
		"currencystate": {
		  "flags": 3,
		  "currencyid": "i8JRDSbog7jBEpbeVd1ib3huTYuNVYeXSy",
		  "reservecurrencies": [
			{
			  "currencyid": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
			  "weight": 0.25000000,
			  "reserves": 1000246.95597757,
			  "priceinreserve": 0.99999692
			},
			{
			  "currencyid": "iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3",
			  "weight": 0.25000000,
			  "reserves": 93.04175960,
			  "priceinreserve": 0.00009301
			},
			{
			  "currencyid": "iHk5GXF4XU2AodPk4KwqwHxbfkaLpdJ56K",
			  "weight": 0.25000000,
			  "reserves": 2894.49327578,
			  "priceinreserve": 0.00289376
			},
			{
			  "currencyid": "i8xcvrfKTw8eEiTMjt1dAQxpWcbZeaVMFV",
			  "weight": 0.25000000,
			  "reserves": 253061.76562890,
			  "priceinreserve": 0.25299850
			}
		  ],
		  "initialsupply": 4000000.00000000,
		  "emitted": 0.00000000,
		  "supply": 4001000.12802839,
		  "currencies": {
			"iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq": {
			  "reservein": 999.75000000,
			  "nativein": 0.00000000,
			  "reserveout": 0.00000000,
			  "lastconversionprice": 0.99962202,
			  "viaconversionprice": 0.99627543,
			  "fees": 0.12500000,
			  "conversionfees": 0.25000000
			},
			"iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3": {
			  "reservein": 0.00000000,
			  "nativein": 0.00000000,
			  "reserveout": 0.00000000,
			  "lastconversionprice": 0.00009304,
			  "viaconversionprice": 0.00009345,
			  "fees": 0.00000000,
			  "conversionfees": 0.00000000
			},
			"iHk5GXF4XU2AodPk4KwqwHxbfkaLpdJ56K": {
			  "reservein": 0.00000000,
			  "nativein": 0.00000000,
			  "reserveout": 0.00000000,
			  "lastconversionprice": 0.00289449,
			  "viaconversionprice": 0.00288587,
			  "fees": 0.00000000,
			  "conversionfees": 0.00000000
			},
			"i8xcvrfKTw8eEiTMjt1dAQxpWcbZeaVMFV": {
			  "reservein": 0.00000000,
			  "nativein": 0.00000000,
			  "reserveout": 0.00000000,
			  "lastconversionprice": 0.25306176,
			  "viaconversionprice": 0.25230859,
			  "fees": 0.00000000,
			  "conversionfees": 0.00000000
			}
		  },
		  "nativefees": 12500000,
		  "nativeconversionfees": 25000000
		},
		"prevnotarization": "69a4ca2d7962a44e894b71220440f55bc4110b3e32fd781ffa1864b078576691",
		"prevheight": 1479,
		"crossnotarization": "0000000000000000000000000000000000000000000000000000000000000000",
		"crossheight": 0,
		"nodes": [
		]
	  }
	},
	{
	  "VRSC-BTC-ETH-GOLD": {
		"name": "VRSC-BTC-ETH-GOLD",
		"version": 1,
		"options": 97,
		"parent": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
		"systemid": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
		"currencyid": "iGaHsN6TyXXwfKQRQfWV1DSKa2Y6e7frPx",
		"notarizationprotocol": 1,
		"proofprotocol": 1,
		"idregistrationprice": 0,
		"idreferrallevels": 0,
		"minnotariesconfirm": 0,
		"billingperiod": 0,
		"notarizationreward": 0,
		"startblock": 147,
		"endblock": 0,
		"currencies": [
		  "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
		  "iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3",
		  "iHk5GXF4XU2AodPk4KwqwHxbfkaLpdJ56K",
		  "iQP7TeWNDNsF7aaaCkQzNyS4jDjdKncNWf"
		],
		"weights": [
		  0.25000000,
		  0.25000000,
		  0.25000000,
		  0.25000000
		],
		"conversions": [
		  0.00000000,
		  0.00000000,
		  0.00000000,
		  0.00000000
		],
		"initialsupply": 4000000.00000000,
		"prelaunchcarveout": 0.00000000,
		"initialcontributions": [
		  1000000.00000000,
		  94.13377439,
		  2893.76965247,
		  533.26792469
		],
		"preconversions": [
		  1000000.00000000,
		  94.13377439,
		  2893.76965247,
		  533.26792469
		],
		"eras": [
		]
	  },
	  "lastnotarization": {
		"version": 1,
		"currencyid": "iGaHsN6TyXXwfKQRQfWV1DSKa2Y6e7frPx",
		"notaryaddress": "iGaHsN6TyXXwfKQRQfWV1DSKa2Y6e7frPx",
		"notarizationheight": 148,
		"mmrroot": "65cc5163afcc30067a25ebe2427a3e211ad769722d1a87769326a927cf889e88",
		"notarizationprehash": "bd850c19d5e2f61ca567b096f3682d1f349e4ae29f8f1b088cb2d0cf728d11fc",
		"work": "000000000000000000000000000000000000000000000000000000000282dffd",
		"stake": "0000000000000000000000000000000000000000000000000000000000000000",
		"currencystate": {
		  "flags": 3,
		  "currencyid": "iGaHsN6TyXXwfKQRQfWV1DSKa2Y6e7frPx",
		  "reservecurrencies": [
			{
			  "currencyid": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
			  "weight": 0.25000000,
			  "reserves": 999250.18169948,
			  "priceinreserve": 0.99925018
			},
			{
			  "currencyid": "iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3",
			  "weight": 0.25000000,
			  "reserves": 94.15731371,
			  "priceinreserve": 0.00009415
			},
			{
			  "currencyid": "iHk5GXF4XU2AodPk4KwqwHxbfkaLpdJ56K",
			  "weight": 0.25000000,
			  "reserves": 2894.49327578,
			  "priceinreserve": 0.00289449
			},
			{
			  "currencyid": "iQP7TeWNDNsF7aaaCkQzNyS4jDjdKncNWf",
			  "weight": 0.25000000,
			  "reserves": 533.40127500,
			  "priceinreserve": 0.00053340
			}
		  ],
		  "initialsupply": 4000000.00000000,
		  "emitted": 0.00000000,
		  "supply": 4000000.00000000,
		  "currencies": {
			"iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq": {
			  "reservein": 1000000.00000000,
			  "nativein": 0.00000000,
			  "reserveout": 749.81830052,
			  "lastconversionprice": 1.00000000,
			  "viaconversionprice": 0.99953131,
			  "fees": 250.06331562,
			  "conversionfees": 250.06251562
			},
			"iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3": {
			  "reservein": 94.15731371,
			  "nativein": 0.00000000,
			  "reserveout": 0.00000000,
			  "lastconversionprice": 0.00009414,
			  "viaconversionprice": 0.00009413,
			  "fees": 0.02353932,
			  "conversionfees": 0.02353932
			},
			"iHk5GXF4XU2AodPk4KwqwHxbfkaLpdJ56K": {
			  "reservein": 2894.49327578,
			  "nativein": 0.00000000,
			  "reserveout": 0.00000000,
			  "lastconversionprice": 0.00289377,
			  "viaconversionprice": 0.00289395,
			  "fees": 0.72362331,
			  "conversionfees": 0.72362331
			},
			"iQP7TeWNDNsF7aaaCkQzNyS4jDjdKncNWf": {
			  "reservein": 533.40127500,
			  "nativein": 0.00000000,
			  "reserveout": 0.00000000,
			  "lastconversionprice": 0.00053327,
			  "viaconversionprice": 0.00053330,
			  "fees": 0.13335031,
			  "conversionfees": 0.13335031
			}
		  },
		  "nativefees": 99988161614,
		  "nativeconversionfees": 99988081614
		},
		"prevnotarization": "c2523226de37a1a0ee9c400f4285ce9feaa337244a4fdde5096f097381b41ecc",
		"prevheight": 0,
		"crossnotarization": "0000000000000000000000000000000000000000000000000000000000000000",
		"crossheight": 0,
		"nodes": [
		]
	  }
	},
	{
	  "multifractional": {
		"name": "multifractional",
		"version": 1,
		"options": 97,
		"parent": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
		"systemid": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
		"currencyid": "iJJUSbSbsc8FAuBGSYhZ9kcmJ7NVQKqd5D",
		"notarizationprotocol": 1,
		"proofprotocol": 1,
		"idregistrationprice": 0,
		"idreferrallevels": 0,
		"minnotariesconfirm": 0,
		"billingperiod": 0,
		"notarizationreward": 0,
		"startblock": 475,
		"endblock": 0,
		"currencies": [
		  "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
		  "iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3",
		  "iHk5GXF4XU2AodPk4KwqwHxbfkaLpdJ56K"
		],
		"weights": [
		  0.33333334,
		  0.33333333,
		  0.33333333
		],
		"conversions": [
		  0.00000000,
		  0.00000000,
		  0.00000000
		],
		"initialsupply": 2000.00000000,
		"prelaunchcarveout": 0.50000000,
		"initialcontributions": [
		  1000.00000000,
		  0.10000000,
		  2.50000000
		],
		"preconversions": [
		  1000.00000000,
		  0.10000000,
		  2.50000000
		],
		"eras": [
		]
	  },
	  "lastnotarization": {
		"version": 1,
		"currencyid": "iJJUSbSbsc8FAuBGSYhZ9kcmJ7NVQKqd5D",
		"notaryaddress": "iJJUSbSbsc8FAuBGSYhZ9kcmJ7NVQKqd5D",
		"notarizationheight": 478,
		"mmrroot": "87859317a8db63b8f4384f3d70c453a5fd17b62a98f2f6bf6fb19abcefe216a8",
		"notarizationprehash": "1c947fa17b7f0fbcb26add9e571386a483b5fa61e236ce373f685c2dac920db9",
		"work": "000000000000000000000000000000000000000000000000000000001ff91d7b",
		"stake": "0000000000000000000000000000000000000000000000000000000000000000",
		"currencystate": {
		  "flags": 3,
		  "currencyid": "iJJUSbSbsc8FAuBGSYhZ9kcmJ7NVQKqd5D",
		  "reservecurrencies": [
			{
			  "currencyid": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
			  "weight": 0.16666667,
			  "reserves": 498.87672231,
			  "priceinreserve": 1.49663013
			},
			{
			  "currencyid": "iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3",
			  "weight": 0.16666666,
			  "reserves": 0.05010000,
			  "priceinreserve": 0.00015030
			},
			{
			  "currencyid": "iHk5GXF4XU2AodPk4KwqwHxbfkaLpdJ56K",
			  "weight": 0.16666666,
			  "reserves": 1.25031257,
			  "priceinreserve": 0.00375093
			}
		  ],
		  "initialsupply": 2000.00000000,
		  "emitted": 0.00000000,
		  "supply": 2000.00000000,
		  "currencies": {
			"iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq": {
			  "reservein": 1000.00000000,
			  "nativein": 0.00000000,
			  "reserveout": 2.24655537,
			  "lastconversionprice": 1.49999998,
			  "viaconversionprice": 1.49775284,
			  "fees": 0.25066251,
			  "conversionfees": 0.25006251
			},
			"iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3": {
			  "reservein": 0.10020000,
			  "nativein": 0.00000000,
			  "reserveout": 0.00000000,
			  "lastconversionprice": 0.00015001,
			  "viaconversionprice": 0.00015018,
			  "fees": 0.00020000,
			  "conversionfees": 0.00020000
			},
			"iHk5GXF4XU2AodPk4KwqwHxbfkaLpdJ56K": {
			  "reservein": 2.50062515,
			  "nativein": 0.00000000,
			  "reserveout": 0.00000000,
			  "lastconversionprice": 0.00375001,
			  "viaconversionprice": 0.00374812,
			  "fees": 0.00062515,
			  "conversionfees": 0.00062515
			}
		  },
		  "nativefees": 249721788,
		  "nativeconversionfees": 249661788
		},
		"prevnotarization": "b6cc5359c07ffc2ffdc13352753f4ce0b37faabe57d0d9d487abe925ec4abdda",
		"prevheight": 0,
		"crossnotarization": "0000000000000000000000000000000000000000000000000000000000000000",
		"crossheight": 0,
		"nodes": [
		]
	  }
	},
	{
	  "VRSC-BTC-ETH": {
		"name": "VRSC-BTC-ETH",
		"version": 1,
		"options": 97,
		"parent": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
		"systemid": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
		"currencyid": "iL8hNCMca68MUrfzy8hD8ZFAxR3PkqESat",
		"notarizationprotocol": 1,
		"proofprotocol": 1,
		"idregistrationprice": 0,
		"idreferrallevels": 0,
		"minnotariesconfirm": 0,
		"billingperiod": 0,
		"notarizationreward": 0,
		"startblock": 147,
		"endblock": 0,
		"currencies": [
		  "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
		  "iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3",
		  "iHk5GXF4XU2AodPk4KwqwHxbfkaLpdJ56K"
		],
		"weights": [
		  0.33333334,
		  0.33333333,
		  0.33333333
		],
		"conversions": [
		  0.00000000,
		  0.00000000,
		  0.00000000
		],
		"initialsupply": 3000000.00000000,
		"prelaunchcarveout": 0.00000000,
		"initialcontributions": [
		  1000000.00000000,
		  94.13377439,
		  2893.76965247
		],
		"preconversions": [
		  1000000.00000000,
		  94.13377439,
		  2893.76965247
		],
		"eras": [
		]
	  },
	  "lastnotarization": {
		"version": 1,
		"currencyid": "iL8hNCMca68MUrfzy8hD8ZFAxR3PkqESat",
		"notaryaddress": "iL8hNCMca68MUrfzy8hD8ZFAxR3PkqESat",
		"notarizationheight": 1541,
		"mmrroot": "3af123cb189e66beeb789fedf5670e42acfa18c1efff4f881df82e46d3c4bb52",
		"notarizationprehash": "e197aafb7198cd8fa5a0fa8595816bd512b0d561ddea6bc53688452a2ac342ac",
		"work": "00000000000000000000000000000000000000000000000000000000bdb42d72",
		"stake": "00000000000000000000000000000000000000000000000000000207ecc87395",
		"currencystate": {
		  "flags": 3,
		  "currencyid": "iL8hNCMca68MUrfzy8hD8ZFAxR3PkqESat",
		  "reservecurrencies": [
			{
			  "currencyid": "iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq",
			  "weight": 0.33333334,
			  "reserves": 997254.27969994,
			  "priceinreserve": 0.99725397
			},
			{
			  "currencyid": "iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3",
			  "weight": 0.33333333,
			  "reserves": 94.15736371,
			  "priceinreserve": 0.00009415
			},
			{
			  "currencyid": "iHk5GXF4XU2AodPk4KwqwHxbfkaLpdJ56K",
			  "weight": 0.33333333,
			  "reserves": 2901.01246078,
			  "priceinreserve": 0.00290101
			}
		  ],
		  "initialsupply": 3000000.00000000,
		  "emitted": 0.00000000,
		  "supply": 3000000.84409421,
		  "currencies": {
			"iJhCezBExJHvtyH3fGhNnt2NhU4Ztkf2yq": {
			  "reservein": 0.00000000,
			  "nativein": 0.00000000,
			  "reserveout": 0.52960919,
			  "lastconversionprice": 0.99725450,
			  "viaconversionprice": 0.99725413,
			  "fees": 0.00000000,
			  "conversionfees": 0.00000000
			},
			"iBBRjDbPf3wdFpghLotJQ3ESjtPBxn6NS3": {
			  "reservein": 0.00005000,
			  "nativein": 0.00000000,
			  "reserveout": 0.00000000,
			  "lastconversionprice": 0.00009415,
			  "viaconversionprice": 0.00009415,
			  "fees": 0.00005000,
			  "conversionfees": 0.00010000
			},
			"iHk5GXF4XU2AodPk4KwqwHxbfkaLpdJ56K": {
			  "reservein": 0.00000000,
			  "nativein": 0.00000000,
			  "reserveout": 0.00000000,
			  "lastconversionprice": 0.00290101,
			  "viaconversionprice": 0.00290101,
			  "fees": 0.00000000,
			  "conversionfees": 0.00000000
			}
		  },
		  "nativefees": 52960919,
		  "nativeconversionfees": 0
		},
		"prevnotarization": "494e2036b701d6942cf8d419bb9d64a0dc211c8004ab2d6471d99bb1d595db1a",
		"prevheight": 810,
		"crossnotarization": "0000000000000000000000000000000000000000000000000000000000000000",
		"crossheight": 0,
		"nodes": [
		]
	  }
	}
  ]`

var data2 = `[
	{
	  "iwuyraasdf": {
		"name": "john"
	  },
	  "otherinfo": {
		"city": "ABC"
	  }
	},
	{
	  "iugwsdf": {
		"name": "bob"
	  },
	  "otherinfo": {
		"city": "DEF"
	  }
	},
	{
	  "y8uwegkhfj": {
		"name": "jane"
	  },
	  "otherinfo": {
		"city": "GHI"
	  }
	}
  ]`

var data3 = `{
    "iwuyraasdf": {
      "name": "john"
    },
    "otherinfo": {
      "city": "ABC"
    }
  }`

func main() {

	// fmt.Println(data)

	var curConv []GetCurrencyConverter

	json.Unmarshal([]byte(data), &curConv)

	// fmt.Printf("%+v\n", curConv)
	// fmt.Println(curConv[0])
	// fmt.Printf("CurrencyInfo --- %+v\n", curConv[0].CurrencyInfo)

	var f interface{}
	// m := f.(map[string]interface{})
	err := json.Unmarshal([]byte(data), &f)
	if err != nil {
		log.Printf("%v", err)
	}
	// fmt.Printf("%+v\n", f)

	// fmt.Printf("%T\n", f)

	m := f.([]interface{})
	// m := f.(map[string]interface{})

	// fmt.Println(m)

	for k, v := range m {
		// fmt.Printf("k -- %+v\n", k)
		// fmt.Println("v --- ", v)
		n := v.(map[string]interface{})
		if _, ok := v.(map[string]interface{})["lastnotarization"]; ok {
			// fmt.Println("lastnotarization ---", val)
			delete(v.(map[string]interface{}), "lastnotarization")
		}
		if _, ok := v.(map[string]interface{})["multifractional"]; ok {
			// fmt.Println("multifractional ---", val)
			delete(v.(map[string]interface{}), "multifractional")
		}
		for nk, nv := range n {
			// fmt.Printf("nk - %+v\n", nk)
			// fmt.Printf("nv - %T\n", nv)
			switch vv := nv.(type) {
			case string:
				fmt.Println(nk, "is string", vv)
			case float64:
				fmt.Println(nk, "is float64", vv)
			case []interface{}:
				fmt.Println(nk, "is an array:")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				// fmt.Println(nk, "is of a type I don't know how to handle")
				// fmt.Printf("%T\n", vv)
				// fmt.Printf("vv -- %+v\n", vv)
				if val, ok := vv.(map[string]interface{})["name"]; ok {
					// fmt.Println("name ---", val)
					curConv[k].CurrencyInfo.Name = val.(string)
				}
				if val, ok := vv.(map[string]interface{})["version"]; ok {
					// fmt.Println("version ---", val)
					curConv[k].CurrencyInfo.Version = int(val.(float64))
				}
				if val, ok := vv.(map[string]interface{})["options"]; ok {
					// fmt.Println("options ---", val)
					curConv[k].CurrencyInfo.Options = int(val.(float64))
				}
				if val, ok := vv.(map[string]interface{})["parent"]; ok {
					// fmt.Println("parent ---", val)
					curConv[k].CurrencyInfo.Parent = val.(string)
				}
				if val, ok := vv.(map[string]interface{})["systemid"]; ok {
					// fmt.Println("systemid ---", val)
					curConv[k].CurrencyInfo.Systemid = val.(string)
				}
				if val, ok := vv.(map[string]interface{})["currencyid"]; ok {
					// fmt.Println("currencyid ---", val)
					curConv[k].CurrencyInfo.Currencyid = val.(string)
				}
				if val, ok := vv.(map[string]interface{})["notarizationprotocol"]; ok {
					// fmt.Println("notarizationprotocol ---", val)
					curConv[k].CurrencyInfo.Notarizationprotocol = int(val.(float64))
				}
				if val, ok := vv.(map[string]interface{})["proofprotocol"]; ok {
					// fmt.Println("proofprotocol ---", val)
					curConv[k].CurrencyInfo.Proofprotocol = int(val.(float64))
				}
				if val, ok := vv.(map[string]interface{})["idregistrationprice"]; ok {
					// fmt.Println("idregistrationprice ---", val)
					curConv[k].CurrencyInfo.Idregistrationprice = int(val.(float64))
				}
				if val, ok := vv.(map[string]interface{})["idreferrallevels"]; ok {
					// fmt.Println("idreferrallevels ---", val)
					curConv[k].CurrencyInfo.Idreferrallevels = int(val.(float64))
				}
				if val, ok := vv.(map[string]interface{})["minnotariesconfirm"]; ok {
					// fmt.Println("minnotariesconfirm ---", val)
					curConv[k].CurrencyInfo.Minnotariesconfirm = int(val.(float64))
				}
				if val, ok := vv.(map[string]interface{})["billingperiod"]; ok {
					// fmt.Println("billingperiod ---", val)
					curConv[k].CurrencyInfo.Billingperiod = int(val.(float64))
				}
				if val, ok := vv.(map[string]interface{})["notarizationreward"]; ok {
					// fmt.Println("notarizationreward ---", val)
					curConv[k].CurrencyInfo.Notarizationreward = int(val.(float64))
				}
				if val, ok := vv.(map[string]interface{})["startblock"]; ok {
					// fmt.Println("startblock ---", val)
					curConv[k].CurrencyInfo.Startblock = int(val.(float64))
				}
				if val, ok := vv.(map[string]interface{})["endblock"]; ok {
					// fmt.Println("endblock ---", val)
					curConv[k].CurrencyInfo.Endblock = int(val.(float64))
				}
				if val, ok := vv.(map[string]interface{})["currencies"]; ok {
					// fmt.Printf("currencies type --- %T\n", val)
					// fmt.Println("currencies ---", val)
					var _tmpCurrencies []string
					for _, curv := range val.([]interface{}) {
						// fmt.Println(curv)
						// fmt.Printf("curv -- %T\n", curv)
						_tmpCurrencies = append(_tmpCurrencies, fmt.Sprintf("%v", curv))
					}
					// fmt.Printf("%T\n", _tmpCurrencies)
					// fmt.Printf("%+v\n", _tmpCurrencies)
					curConv[k].CurrencyInfo.Currencies = _tmpCurrencies
				}
				if val, ok := vv.(map[string]interface{})["weights"]; ok {
					// fmt.Println("weights ---", val)
					var _tmpWeights []float64
					for _, wghtv := range val.([]interface{}) {
						// fmt.Println(wghtv)
						// fmt.Printf("wghtv -- %T\n", wghtv)
						_tmpWeights = append(_tmpWeights, wghtv.(float64))
					}
					// fmt.Printf("%T\n", _tmpWeights)
					// fmt.Printf("%+v\n", _tmpWeights)
					curConv[k].CurrencyInfo.Weights = _tmpWeights
				}
				if val, ok := vv.(map[string]interface{})["conversions"]; ok {
					// fmt.Println("conversions ---", val)
					var _tmpConversions []float64
					for _, cnvrsv := range val.([]interface{}) {
						// fmt.Println(cnvrsv)
						// fmt.Printf("cnvrsv -- %T\n", cnvrsv)
						_tmpConversions = append(_tmpConversions, cnvrsv.(float64))
					}
					// fmt.Printf("%T\n", _tmpConversions)
					// fmt.Printf("%+v\n", _tmpConversions)
					curConv[k].CurrencyInfo.Conversions = _tmpConversions
				}
				if val, ok := vv.(map[string]interface{})["initialsupply"]; ok {
					// fmt.Println("initialsupply ---", val)
					curConv[k].CurrencyInfo.Initialsupply = val.(float64)
				}
				if val, ok := vv.(map[string]interface{})["prelaunchcarveout"]; ok {
					// fmt.Println("prelaunchcarveout ---", val)
					curConv[k].CurrencyInfo.Prelaunchcarveout = val.(float64)
				}
				if val, ok := vv.(map[string]interface{})["initialcontributions"]; ok {
					// fmt.Println("initialcontributions ---", val)
					var _tmpInitContri []float64
					for _, initcontv := range val.([]interface{}) {
						// fmt.Println(initcontv)
						// fmt.Printf("initcontv -- %T\n", initcontv)
						_tmpInitContri = append(_tmpInitContri, initcontv.(float64))
					}
					// fmt.Printf("%T\n", _tmpInitContri)
					// fmt.Printf("%+v\n", _tmpInitContri)
					curConv[k].CurrencyInfo.Initialcontributions = _tmpInitContri
				}
				if val, ok := vv.(map[string]interface{})["preconversions"]; ok {
					// fmt.Println("preconversions ---", val)
					var _tmpPreConv []float64
					for _, preconv := range val.([]interface{}) {
						// fmt.Println(preconv)
						// fmt.Printf("preconv -- %T\n", preconv)
						_tmpPreConv = append(_tmpPreConv, preconv.(float64))
					}
					// fmt.Printf("%T\n", _tmpPreConv)
					// fmt.Printf("%+v\n", _tmpPreConv)
					curConv[k].CurrencyInfo.Preconversions = _tmpPreConv
				}
				if val, ok := vv.(map[string]interface{})["eras"]; ok {
					// fmt.Println("eras ---", val)
					curConv[k].CurrencyInfo.Eras = val.([]interface{})
				}
			}
		}
		// break
	}

	for i, v := range curConv {
		fmt.Println(i)
		fmt.Printf("CurrencyInfo --- %+v\n", v.CurrencyInfo)
		fmt.Printf("Lastnotarization --- %+v\n", v.Lastnotarization)
		if v.Multifractional.Name != "" {
			fmt.Printf("\nMultifractional --- %+v\n\n", v.Multifractional)
		}
		// break
	}
}
