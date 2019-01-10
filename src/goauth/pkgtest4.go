package main

import (
    "fmt"
    "log"
    "goauth/mypkg3"
)

func main() {
    var appName mypkg.AppType
    appName = `komodo`
    var err error
    fmt.Printf("\n")

    var bh mypkg.GetBestBlockhash
    bh, err = appName.GetBestBlockhash()
    if err != nil {
        log.Println("err happened", err)
    }
    fmt.Println("bh value", bh)
    fmt.Println(bh.Result)
    fmt.Println(bh.Error.Code)
    fmt.Println(bh.Error.Message)
    fmt.Println(bh.ID)

    /*
    var info mypkg.GetInfo

    info, err = appName.GetInfo()
    if err != nil {
        log.Println("err happened", err)
    }
    fmt.Println("getinfo value", info)
    fmt.Println(info.Result)
    fmt.Println(info.Error.Code)
    fmt.Println(info.Error.Message)

    fmt.Printf("\n\n\n")

    fmt.Println(info.Result.Version)
    fmt.Println(info.Result.Balance)
    fmt.Println(info.Result.Blocks)
    fmt.Println(info.Result.Name)
    */
}