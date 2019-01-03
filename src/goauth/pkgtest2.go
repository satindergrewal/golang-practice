package main

import (
	"fmt"
	"goauth/mypkg"
)

func main() {
	var rval mypkg.GetInfo
    rval = mypkg.ResultGetInfo()
    fmt.Println(rval)
    fmt.Println(rval.Result)
    fmt.Println(rval.Result.Version)
    fmt.Println(rval.Result.Balance)
    fmt.Println(rval.Result.Blocks)
    fmt.Println(rval.Result.Name)
    fmt.Println(rval.Error)
    fmt.Println(rval.ID)
}