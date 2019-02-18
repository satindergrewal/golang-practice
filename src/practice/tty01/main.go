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
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/tty", tty)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8081", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	f := req.FormValue("first")
	l := req.FormValue("last")
	s := req.FormValue("subscribe") == "on"

	err := tpl.ExecuteTemplate(w, "index.gohtml", person{f, l, s})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}


func tty(w http.ResponseWriter, req *http.Request) {

	v := req.FormValue("q")
	fmt.Fprintln(w, "Do my search: "+v)

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

	var cmd string
	if v != `exit` || v == "" {
		cmd = v
	} else {
		exit(nil, 1)
	}
	cmd_args := []string{}

	factory, err := localcommand.NewFactory(cmd, cmd_args, backendOptions)
	if err != nil {
		exit(err, 3)
	}

	hostname, _ := os.Hostname()
	appOptions.Port = `8082`
	appOptions.PermitWrite = true
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

	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
	ctx, _ := context.WithCancel(context.Background())
	gCtx, _ := context.WithCancel(context.Background())
	srv.Run(ctx, server.WithGracefullContext(gCtx))
}

func exit(err error, code int) {
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}