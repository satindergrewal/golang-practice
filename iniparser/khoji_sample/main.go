package main

import (
	"fmt"
	"golang-practice/iniparser/khoji_sample/sub"
	"os"

	"gopkg.in/ini.v1"
)

var rDB string

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	rDB = cfg.Section("DATABASE").Key("RDB_DB").String()
	fmt.Println("main - init - rDB:", rDB)
}

func main() {

	sub.Sub()
	fmt.Println("main - main - rDB:", rDB)
}

// $ go run main.go
// sub - init - rDB: vrsctest
// main - init - rDB: vrsctest
// sub - Sub func - rDB: vrsctest
// main - main - rDB: vrsctest
