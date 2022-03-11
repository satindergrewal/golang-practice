package sub

import (
	"fmt"
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
	fmt.Println("sub - init - rDB:", rDB)
}

func Sub() {
	fmt.Println("sub - Sub func - rDB:", rDB)
}
