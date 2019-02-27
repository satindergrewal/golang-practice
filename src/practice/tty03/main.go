package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"context"
	"os"
	"runtime"
	"strings"

	"github.com/yudai/gotty/backend/localcommand"
	"github.com/yudai/gotty/server"
	"github.com/yudai/gotty/utils"

	"github.com/satindergrewal/kmdgo/kmdutil"

	"github.com/gorilla/websocket"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type ttycommand struct {
	Command   string
	Arguments string
	//Subscribed bool
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


type myStruct struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
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
	http.HandleFunc("/sockets", socketshandle)
	http.HandleFunc("/v3/ws", v3ws)
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

	//ctx, _ := context.WithCancel(context.Background())
	//gCtx, _ := context.WithCancel(context.Background())
	//go srv.Run(ctx, server.WithGracefullContext(gCtx))
	//http.Redirect(w, req, "http://localhost:8082", http.StatusSeeOther)

	ctx, cancel := context.WithCancel(context.Background())
	gCtx, _ := context.WithCancel(context.Background())

	errs := make(chan error, 1)
	go func() {
		errs <- srv.Run(ctx, server.WithGracefullContext(gCtx))
	}()

	err = tpl.ExecuteTemplate(w, "play.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}

	go func() {
		for {
			err := <-errs
			cancel()
			fmt.Println("<<--errs: ", err)
			fmt.Println("----------")
			fmt.Println("context err after Game end:\t", ctx.Err())
			fmt.Printf("context type after Game end:\t%T\n", ctx)
			fmt.Println("Goroutines after Game end\t", runtime.NumGoroutine())
			return
		}
	}()

	fmt.Println("----------")
	fmt.Println("context before Game end:\t", ctx)
	fmt.Println("context err before Game end:\t", ctx.Err())
	fmt.Printf("context type before Game end:\t%T\n", ctx)
	fmt.Println("Goroutines before Game end\t", runtime.NumGoroutine())

	/*err = waitSignals(errs, cancel, gCancel)

	if err != nil && err != context.Canceled {
		fmt.Printf("Error: %s\n", err)
		//exit(err, 8)
	}*/
}

func exit(err error, code int) {
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}


/*
var ws = new WebSocket("ws://localhost:8081/sockets")
ws.addEventListener("message", function(e) {console.log(e);});
ws.onmessage = function (event) {
	console.log(event.data);
}
ws.send("foo")
ws.send(JSON.stringify({username: "Sat"}))
*/
func socketshandle(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello from go")
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		msgType, msg, err := socket.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(msg))
		if err = socket.WriteMessage(msgType, msg); err != nil {
			fmt.Println(err)
		}
	}
}


func v3ws(w http.ResponseWriter, r *http.Request) {
	var conn, _ = upgrader.Upgrade(w, r, nil)
	go func(conn *websocket.Conn) {
		ch := time.Tick(5 * time.Second)

		for range ch {
			conn.WriteJSON(myStruct{
				Username:  "mvansickle",
				FirstName: "Michael",
				LastName:  "Van Sickle",
			})
		}
	}(conn)
}

