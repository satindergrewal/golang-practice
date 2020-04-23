package main

// https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func main() {
	cmd := exec.Command("ls", "-lah")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("tasklist")
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}

// func main() {
// 	cmd := exec.Command("programToExecute")

// 	additionalEnv := "FOO=bar"
// 	newEnv := append(os.Environ(), additionalEnv)
// 	cmd.Env = newEnv

// 	out, err := cmd.CombinedOutput()
// 	if err != nil {
// 		log.Fatalf("cmd.Run() failed with %s\n", err)
// 	}
// 	fmt.Printf("%s", out)
// }
