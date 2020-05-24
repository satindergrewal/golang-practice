// source: https://socketloop.com/tutorials/golang-daemonizing-a-simple-web-server-process-example

package main

import (
	"fmt"
	"golang-practice/kmdutil"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/satindergrewal/kmdgo"
)

var PIDFile = "./daemonize.pid"
var DexPid = "./kmd_dex.pid"

func savePID(p string, pid int) {

	var file *os.File
	var err error
	if p == "self" {
		file, err = os.Create(PIDFile)
	} else {
		p = "./kmd_" + p + ".pid"
		file, err = os.Create(p)
	}

	if err != nil {
		log.Printf("Unable to create pid file : %v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(pid))

	if err != nil {
		log.Printf("Unable to create pid file : %v\n", err)
		os.Exit(1)
	}

	file.Sync() // flush to disk

}

func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	html := "Hello World"

	w.Write([]byte(html))
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage : %s [start|stop] \n ", os.Args[0]) // return the program name back to %s
		os.Exit(0)                                            // graceful exit
	}

	if strings.ToLower(os.Args[1]) == "main" {

		// Make arrangement to remove PID file upon receiving the SIGTERM from kill command
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)

		go func() {
			signalType := <-ch
			signal.Stop(ch)
			fmt.Println("Exit command received. Exiting...")

			// this is a good place to flush everything to disk
			// before terminating.
			fmt.Println("Received signal type : ", signalType)

			// remove PID file
			os.Remove(PIDFile)

			os.Exit(0)

		}()

		mux := http.NewServeMux()
		mux.HandleFunc("/", SayHelloWorld)
		log.Fatalln(http.ListenAndServe(":8080", mux))
	}

	if strings.ToLower(os.Args[1]) == "start" {

		// check if daemon already running.
		if _, err := os.Stat(PIDFile); err == nil {
			fmt.Println("Already running or ./daemonize.pid file exist.")
			os.Exit(1)
		}

		cmd := exec.Command(os.Args[0], "main")
		cmd.Start()
		fmt.Println("Daemon process ID is : ", cmd.Process.Pid)
		savePID("self", cmd.Process.Pid)
		// os.Exit(0)

		appName := "DEX"
		dir := kmdutil.AppDataDir(appName, false)
		fmt.Println(dir)
		if _, err := os.Stat(dir + "/komodod.pid"); err == nil {
			fmt.Println("Already running or DEX pid file exist.")
			os.Exit(1)
		}

		dexcmd := exec.Command("komodod", "-ac_name=DEX", "-daemon", "-nSPV=1", "-server", "-ac_supply=999999", "-pubkey=03732f8ef851ff234c74d0df575c2c5b159e2bab3faca4ec52b3f217d5cda5361d", "-dexp2p=2", "-addnode=136.243.58.134", "-handle=satinder", "-recvZaddr=zs1zqks0tergf6nk69evm6awte4xmhf2fd9epnv946vzlhfxztkls6a9lyfmuafda00krvkvj0xagp", "-recvTaddr=RBthCSgNLE3rwvAKce8JNSo7xDpxQEiRTX")
		dexcmd.Start()
		fmt.Println("DEX komodod process ID is : ", dexcmd.Process.Pid)
		// savePID("dex", dexcmd.Process.Pid)
		os.Exit(0)

	}

	// upon receiving the stop command
	// read the Process ID stored in PIDfile
	// kill the process using the Process ID
	// and exit. If Process ID does not exist, prompt error and quit

	if strings.ToLower(os.Args[1]) == "stop" {

		appName := kmdgo.NewAppType(`DEX`)
		var info kmdgo.Stop
		info, err := appName.Stop()
		if err != nil {
			fmt.Printf("Code: %v\n", info.Error.Code)
			fmt.Printf("Message: %v\n\n", info.Error.Message)
			log.Fatalln("Err happened", err)
		}
		// fmt.Println(info)
		fmt.Println(info.Result)

		if _, err := os.Stat(PIDFile); err == nil {
			data, err := ioutil.ReadFile(PIDFile)
			if err != nil {
				fmt.Println("Not running")
				os.Exit(1)
			}
			ProcessID, err := strconv.Atoi(string(data))

			if err != nil {
				fmt.Println("Unable to read and parse process id found in ", PIDFile)
				os.Exit(1)
			}

			process, err := os.FindProcess(ProcessID)

			if err != nil {
				fmt.Printf("Unable to find process ID [%v] with error %v \n", ProcessID, err)
				os.Exit(1)
			}
			// remove PID file
			os.Remove(PIDFile)

			fmt.Printf("Killing process ID [%v] now.\n", ProcessID)
			// kill process and exit immediately
			err = process.Kill()

			if err != nil {
				fmt.Printf("Unable to kill process ID [%v] with error %v \n", ProcessID, err)
				os.Exit(1)
			} else {
				fmt.Printf("Killed process ID [%v]\n", ProcessID)
				os.Exit(0)
			}

		} else {

			fmt.Println("Not running.")
			os.Exit(1)
		}
	} else {
		fmt.Printf("Unknown command : %v\n", os.Args[1])
		fmt.Printf("Usage : %s [start|stop]\n", os.Args[0]) // return the program name back to %s
		os.Exit(1)
	}

}
