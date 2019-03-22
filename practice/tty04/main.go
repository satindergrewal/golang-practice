package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	"math/rand"

	"context"
	"os"
	"runtime"
	"strings"
	"io"

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


// GoTTY command and argument struct
type ttycommand struct {
	Command   string
	Arguments string
	//Subscribed bool
}

// Gorilla Websocket upgrader settings
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


// Some example user struct which is used with websocket url v3/ws
type myStruct struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {


	appName := "ROGUE" //Name of my app
	dir := kmdutil.AppDataDir(appName, false) //Based on the name of my app, it gets the location of this app on the machine.
	fmt.Println(dir, "\n") //Displays the Directory location of the app

	err := os.Chdir(dir) //Just changing the current working directory location to the directory location got from the AppDataDir output
	if err != nil {
		fmt.Println(err)
	}
	d, _ := os.Getwd()
	fmt.Println(d) //Printing out the current working directory to console

	var cat hotcat
	cat = 1
	mux := http.NewServeMux()
	mux.Handle("/cat", cat)

	http.HandleFunc("/", foo) //Displays the template form with input form to input command and argument
	http.HandleFunc("/tty", tty) //Output of foo forwards to this, an iFrame page showing the URL http://localhost:8082, which is running a TTY cosonle app, example a console Rogue game.
	//http.Handle("/cat", cat)
	http.HandleFunc("/sockets", socketshandle) //Just the usual example of gorilla websockets, Read and Write the output coming from browser websocket
	http.HandleFunc("/v3/ws", v3ws) // Just another test example of gorilla websocket for testing.
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8081", mux)
}

func foo(w http.ResponseWriter, req *http.Request) {

	c := req.FormValue("cmd") //Getting the command to pass to GoTTY to start the web SSH with.
	a := req.FormValue("args") //Arguments for that command

	err := tpl.ExecuteTemplate(w, "index.gohtml", ttycommand{c, a})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

func tty(w http.ResponseWriter, req *http.Request) {

	//v := req.FormValue("q")
	//fmt.Fprintln(w, "Do my search: "+v)
	c := req.FormValue("cmd") //Getting the command from foo
	a := req.FormValue("args") //Getting the command arguments from foo
	fmt.Println(c)
	//fmt.Println(a)

	// visit this page:
	// http://localhost:8080/tty?q=dog

	appOptions := &server.Options{} //Getting default GoTTY server options
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

	factory, err := localcommand.NewFactory(cmd, cmd_args, backendOptions) //Passing the command and it's args to GoTTY
	if err != nil {
		exit(err, 3)
	}

	d, _ := os.Getwd()
	fmt.Println(d)

	hostname, _ := os.Hostname()
	appOptions.Port = `8082`
	appOptions.PermitWrite = true //Setting GoTTY's permissions to accept input on web shell, so that Game can be played by user keystrokes on Web
	appOptions.Once = true //Quit the GoTTY after the ROGUE game ends. The http://localhost:8082 will dead/exited as per this option set to true.
	appOptions.TitleVariables = map[string]interface{}{
		"command":  cmd,
		"argv":     cmd_args,
		"hostname": hostname,
	}

	srv, err := server.New(factory, appOptions) //Set GoTTY New ServeHTTP
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
		errs <- srv.Run(ctx, server.WithGracefullContext(gCtx)) //Execute GoTTY Server command and get it's context channel updates in errs
	}()

	err = tpl.ExecuteTemplate(w, "play.gohtml", nil) //the go html template page which has iframe pointing to http://localhost:8082
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}

	// Once the srv.Run is done, execute this part of the code to close any gorotunes of GoTTY
	go func() {
		for {
			err := <-errs
			cancel()
			fmt.Println("<<--errs: ", err)
			fmt.Println("----------")
			fmt.Println("context err after Game end:\t", ctx.Err())
			fmt.Printf("context type after Game end:\t%T\n", ctx)
			fmt.Println("Goroutines after Game end\t", runtime.NumGoroutine())

			//Was just testing Channels between /tty and websockets 
			c := fanIn(boring("Joe"), boring("Ann"))
			for i := 0; i < 10; i++ {
				fmt.Println(<-c)
			}
			fmt.Println("You're both boring; I'm leaving.")
			
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

//Not using this function, but kept it from the early tests of GoTTY.
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



//Gorilla websockets which if establishes just keeps sending user JSON data to browser every 5 seconds untile ws.close().
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


//Testing with example of HandlerFunc from Golang Web Dev course
type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println(c)
	io.WriteString(res, "cat cat cat")
}



/*
code source:
Rob Pike
https://talks.golang.org/2012/concurrency.slide#25

source:
https://blog.golang.org/pipelines
*/
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// FAN IN
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}