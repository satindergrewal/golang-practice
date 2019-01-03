package main

import (
	"fmt"
	"goauth/mypkg2"
)

func main() {
	var rval mypkg.GetInfo
    rval = mypkg.ResultGetInfo()
    //fmt.Println(rval)
    //fmt.Println(rval.Result)
    fmt.Println("Version:", rval.Result.Version)
    fmt.Println("Balance:", rval.Result.Balance)
    fmt.Println("Blocks:", rval.Result.Blocks)
    fmt.Println("Name:", rval.Result.Name)
    fmt.Println("Connections:", rval.Result.Connections)
    fmt.Println("Difficulty:", rval.Result.Difficulty)
    fmt.Println("Magic:", rval.Result.Magic)
    //fmt.Println(rval.Error)
    //fmt.Println(rval.ID)
}