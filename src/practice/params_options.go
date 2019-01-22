package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Actname	string	`json:",actname"`
	Count	int		`json:"count"`
	Frm		int		`json:"frm"`
	Incwch	bool	`json:"incwch"`
}

type params []interface{}

func main() {
	fmt.Println("#########################")
	p1 := person{"savings", 33, 74, true}
	bs, _ := json.Marshal(p1)
	fmt.Println(p1)
	fmt.Println(string(bs))
	fmt.Println("-----")
	
	p2 := person{
		Actname: "eftpos",
		Count: 11,
	}
	if p2.Count == 0 {
		p2.Count = 10
	}
	fmt.Println(p2)
	fmt.Println(p2.Actname)
	bs2, _ := json.Marshal(p2)
	fmt.Println(string(bs2))

	fmt.Println("-----")

	p3 := params{
		"eftpos",
		11,
	}
	fmt.Println(p3)
	bs3, _ := json.Marshal(p3)
	fmt.Println(string(bs3))

	fmt.Println("-----")

	p4 := make(params, 4)
	fmt.Println(p4)
	
	if p4[0] == nil {
		p4[0] = "*"
	}
	if p4[1] == nil {
		p4[1] = 0
	}

	if p4[2] == nil {
		p4[2] = 10
	}
	if p4[3] == nil {
		p4[3] = false
	}
	bs4, _ := json.Marshal(p4)
	fmt.Println(string(bs4))

}


/*

Arguments:
1. "account"    (string, optional) DEPRECATED. The account name. Should be "*".
2. count          (numeric, optional, default=10) The number of transactions to return
3. from           (numeric, optional, default=0) The number of transactions to skip
4. includeWatchonly (bool, optional, default=false) Include transactions to watchonly addresses (see 'importaddress')


Target outputs:

["*", 10, 0, false]

["*", 20, 100, true]

*/