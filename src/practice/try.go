package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	//"github.com/codegangsta/cli"

	"github.com/yudai/gotty/backend/localcommand"
	//"github.com/yudai/gotty/pkg/homedir"
	"github.com/yudai/gotty/server"
	"github.com/yudai/gotty/utils"
)

func main() {
	appOptions := &server.Options{}
	if err := utils.ApplyDefaultValues(appOptions); err != nil {
		exit(err, 1)
	}
	backendOptions := &localcommand.Options{}
	if err := utils.ApplyDefaultValues(backendOptions); err != nil {
		exit(err, 1)
	}

	cmd := `echo`
	cmd_args := []string{"hello"}

	factory, err := localcommand.NewFactory(cmd, cmd_args, backendOptions)
	if err != nil {
		exit(err, 3)
	}

	hostname, _ := os.Hostname()
	appOptions.TitleVariables = map[string]interface{}{
		"command":  cmd,
		"argv":     cmd_args,
		"hostname": hostname,
	}

	srv, err := server.New(factory, appOptions)
	if err != nil {
		exit(err, 3)
	}

	ctx, cancel := context.WithCancel(context.Background())
	gCtx, gCancel := context.WithCancel(context.Background())

	log.Printf("GoTTY is starting with command: %s %s", cmd, strings.Join(cmd_args, " "))

	errs := make(chan error, 1)
	go func() {
		errs <- srv.Run(ctx, server.WithGracefullContext(gCtx))
	}()
	err = waitSignals(errs, cancel, gCancel)

	if err != nil && err != context.Canceled {
		fmt.Printf("Error: %s\n", err)
		exit(err, 8)
	}
}

func exit(err error, code int) {
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

func waitSignals(errs chan error, cancel context.CancelFunc, gracefullCancel context.CancelFunc) error {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(
		sigChan,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	select {
	case err := <-errs:
		return err

	case s := <-sigChan:
		switch s {
		case syscall.SIGINT:
			gracefullCancel()
			fmt.Println("C-C to force close")
			select {
			case err := <-errs:
				return err
			case <-sigChan:
				fmt.Println("Force closing...")
				cancel()
				return <-errs
			}
		default:
			cancel()
			return <-errs
		}
	}
}
