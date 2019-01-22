package main

import (
    "fmt"
    "log"
    "errors"
    "encoding/json"
    "strconv"
    "net/http"
    "io/ioutil"
    "bytes"
    "github.com/satindergrewal/kmdgo/kmdutil"
)


/* Common code START */

type AppType string

type APIResult struct {
    Result interface{}  `json:"result"`
    Error interface{}   `json:"error"`
    ID string           `json:"id"`
}

type APIQuery struct {
    Method string `json:"method"`
    Params string `json:"params"`
}

type Error struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

func (appName AppType) APICall(q APIQuery) string {
    rpcuser, rpcpass, rpcport := kmdutil.AppRPCInfo(string(appName))

    client := &http.Client{}
    url := `http://127.0.0.1:`+rpcport

    var query_str string
    query_str = `{"jsonrpc": "1.0", "id":"kmdgo", "method": "`+q.Method+`", "params": `+q.Params+` }`
    //fmt.Println(query_str)

    query_byte := []byte(query_str)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(query_byte))
    req.Header.Set("Content-Type", "application/json")

    //req, err := http.NewRequest("POST", , nil)
    req.SetBasicAuth(rpcuser, rpcpass)
    resp, err := client.Do(req)
    if err != nil{
        log.Fatal(err)
    }
    bodyText, err := ioutil.ReadAll(resp.Body)

    var query_result map[string]interface{}
    if err := json.Unmarshal(bodyText, &query_result); err != nil {
        panic(err)
    }

    //fmt.Println(string(bodyText))

    s := string(bodyText)
    return s
}

/* Common code END */

type ListTransactions struct {
    Result  []struct {
        InvolvesWatchonly bool          `json:"involvesWatchonly"`
        Account           string        `json:"account"`
        Address           string        `json:"address"`
        Category          string        `json:"category"`
        Amount            float64       `json:"amount"`
        Vout              int           `json:"vout"`
        Fee               float64       `json:"fee,omitempty"`
        Rawconfirmations  int           `json:"rawconfirmations"`
        Confirmations     int           `json:"confirmations"`
        Blockhash         string        `json:"blockhash"`
        Blockindex        int           `json:"blockindex"`
        Blocktime         int           `json:"blocktime"`
        Expiryheight      int           `json:"expiryheight"`
        Txid              string        `json:"txid"`
        Walletconflicts   []interface{} `json:"walletconflicts"`
        Time              int           `json:"time"`
        Timereceived      int           `json:"timereceived"`
        Vjoinsplit        []interface{} `json:"vjoinsplit"`
        Size              int           `json:"size"`
        }   `json:"result"`
    Error   Error   `json:"error"`
    ID      string  `json:"id"`
    
}

type ListTransactionsParams struct {
    actname string
    count int
    frm int
    incwch bool    
}

func (appName AppType) ListTransactions(actname string, count int, frm int, incwch bool) (ListTransactions, error) {
    query := APIQuery {
        Method: `listtransactions`,
        Params: `["`+actname+`", `+strconv.Itoa(count)+`, `+strconv.Itoa(frm)+`, `+strconv.FormatBool(incwch)+`]`,
    }
    //fmt.Println(query)

    var listtransactions ListTransactions

    listtransactionsJson := appName.APICall(query)

    var result APIResult

    json.Unmarshal([]byte(listtransactionsJson), &result)

    if result.Error != nil {
        answer_error, err := json.Marshal(result.Error)
        if err != nil {
        }
        json.Unmarshal([]byte(listtransactionsJson), &listtransactions)
        return listtransactions, errors.New(string(answer_error))
    }

    json.Unmarshal([]byte(listtransactionsJson), &listtransactions)
    return listtransactions, nil
}

func main() {  
    var appName AppType
    appName = `komodo`

    var lstx ListTransactions

    account_name := `*` //Either use * or the "account name".
    numoftx := 10
    skiptx := 100
    watchonly := true

    lstx, err := appName.ListTransactions(account_name, numoftx, skiptx, watchonly)
    if err != nil {
        fmt.Printf("Code: %v\n", lstx.Error.Code)
        fmt.Printf("Message: %v\n\n", lstx.Error.Message)
        log.Fatalln("Err happened", err)
    }

    fmt.Println("lstx value", lstx)
    //fmt.Println("-------")
    //fmt.Println(lstx.Result)
    fmt.Println("-------")
    
    //fmt.Printf("\n\n\n-------")

    for i, v := range lstx.Result {
        fmt.Printf("\n-------\n")
        fmt.Println(i)
        //fmt.Println(v)
        fmt.Println("InvolvesWatchonly: ", v.InvolvesWatchonly)
        fmt.Println("Account: ", v.Account)
        fmt.Println("Address: ", v.Address)
        fmt.Println("Category: ", v.Category)
        fmt.Println("Amount: ", v.Amount)
        fmt.Println("Vout: ", v.Vout)
        fmt.Println("Fee: ", v.Fee)
        fmt.Println("Rawconfirmations: ", v.Rawconfirmations)
        fmt.Println("Confirmations: ", v.Confirmations)
        fmt.Println("Blockhash: ", v.Blockhash)
        fmt.Println("Blockindex: ", v.Blockindex)
        fmt.Println("Blocktime: ", v.Blocktime)
        fmt.Println("Expiryheight: ", v.Expiryheight)
        fmt.Println("Txid: ", v.Txid)
        fmt.Println("Walletconflicts: ", v.Walletconflicts)
        fmt.Println("Time: ", v.Time)
        fmt.Println("Timereceived: ", v.Timereceived)
        fmt.Println("Vjoinsplit: ", v.Vjoinsplit)
        fmt.Println("Size: ", v.Size)
    }
}
