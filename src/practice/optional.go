package main

import (
	"encoding/json"
	"fmt"
)

type AppType string

type params []interface{}

func (appName AppType) ListTransactions(args params) {
	fmt.Println(appName)
	fmt.Println(args)

	if args[0] == nil {
		args[0] = "*"
	}
	if args[1] == nil {
		args[1] = 10
	}

	if args[2] == nil {
		args[2] = 0
	}
	if args[3] == nil {
		args[3] = false
	}
	bs4, _ := json.Marshal(args)
	fmt.Println(string(bs4))
}

func main() {
	App := AppType(`komodo`)
	//fmt.Println(App)

	Args := make(params, 4)
	fmt.Println(Args)

	App.ListTransactions(Args)

	fmt.Println("~~~~~~~~~~~")

	Args2 := make(params, 4)
	Args2[1] = 20
	Args2[2] = 100
	Args2[3] = true
	fmt.Println(Args2)

	App.ListTransactions(Args2)

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
