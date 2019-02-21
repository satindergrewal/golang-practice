package main

import (
	"html/template"
	"log"
	"net/http"
	"fmt"

	"context"
	"os"
	//"os/signal"
	"strings"
	//"syscall"

	"github.com/yudai/gotty/backend/localcommand"
	"github.com/yudai/gotty/server"
	"github.com/yudai/gotty/utils"

	"github.com/satindergrewal/kmdgo/kmdutil"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type ttycommand struct {
	Command  string
	Arguments   string
	//Subscribed bool
}

func main() {
	
	appName := "ROGUE"
	dir := kmdutil.AppDataDir(appName, false)
	fmt.Println(dir, "\n")

	err := os.Chdir(dir)
	if err != nil {
		fmt.Println(err)
	}
	d, _ := os.Getwd()
	fmt.Println(d)

	http.HandleFunc("/", foo)
	http.HandleFunc("/tty", tty)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8081", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	c := req.FormValue("cmd")
	a := req.FormValue("args")

	err := tpl.ExecuteTemplate(w, "index.gohtml", ttycommand{c, a})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}


func tty(w http.ResponseWriter, req *http.Request) {

	//v := req.FormValue("q")
	//fmt.Fprintln(w, "Do my search: "+v)
	c := req.FormValue("cmd")
	a := req.FormValue("args")
	fmt.Println(c)
	//fmt.Println(a)

	// visit this page:
	// http://localhost:8080/tty?q=dog

	appOptions := &server.Options{}
	if err := utils.ApplyDefaultValues(appOptions); err != nil {
		exit(err, 1)
	}
	backendOptions := &localcommand.Options{}
	if err := utils.ApplyDefaultValues(backendOptions); err != nil {
		exit(err, 1)
	}

	//cmd := `echo`
	//cmd_args := []string{"hello"}

	//var cmd string
	//if c == `exit` || c != "" {
		//http.Redirect(w, req, "http://localhost:8081", http.StatusSeeOther)
	//}
	//cmd_args := []string{}
	cmd := c
	cmd_args := strings.Fields(a)
	fmt.Printf("%T\n", cmd_args)

	factory, err := localcommand.NewFactory(cmd, cmd_args, backendOptions)
	if err != nil {
		exit(err, 3)
	}

	d, _ := os.Getwd()
	fmt.Println(d)

	hostname, _ := os.Hostname()
	appOptions.Port = `8082`
	appOptions.PermitWrite = true
	appOptions.Once = true
	appOptions.TitleVariables = map[string]interface{}{
		"command":  cmd,
		"argv":     cmd_args,
		"hostname": hostname,
	}

	srv, err := server.New(factory, appOptions)
	if err != nil {
		exit(err, 3)
	}

	log.Printf("GoTTY is starting with command: %s %s", cmd, strings.Join(cmd_args, " "))

	ctx, _ := context.WithCancel(context.Background())
	gCtx, _ := context.WithCancel(context.Background())
	go srv.Run(ctx, server.WithGracefullContext(gCtx))
	//http.Redirect(w, req, "http://localhost:8082", http.StatusSeeOther)

	err = tpl.ExecuteTemplate(w, "play.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

func exit(err error, code int) {
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}