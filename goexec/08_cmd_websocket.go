package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"text/template"

	"github.com/gorilla/websocket"
)

var html = template.Must(template.New("").Parse(`<html>
	<body>
		<h1>ls -lah</h1>
		<code></code>
		<script>
			// var ws = new WebSocket("ws://127.0.0.1:8000/ws")
			var ws = new WebSocket("{{.}}");
			ws.onmessage = function(e) {
				document.querySelector("code").innerHTML += e.data + "<br>"
			}
		</script>
	</body>
</html>
`))

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/ws", cmdToResponse)
	http.HandleFunc("/", serveHtml)

	log.Println("Listening on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func cmdToResponse(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("", err)))
		return
	}
	defer ws.Close()

	// discard received messages
	go func(c *websocket.Conn) {
		for {
			if _, _, err := c.NextReader(); err != nil {
				c.Close()
				break
			}
		}
	}(ws)

	ws.WriteMessage(1, []byte("Starting...\n"))

	// execute and get a pipe
	// cmd := exec.Command("tail", "-f", "debug.log")
	// cmd.Dir = kmdutil.AppDataDir("komodo", false)
	// cmd := exec.Command("ls", "-lah")
	cmd := exec.Command("./subatomic", "PIRATE", "", "3784961472", "1")
	cmd.Dir = "/Users/satinder/repositories/jl777/komodo/src"
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(err)
		fmt.Println("StdOut Nil")
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Println(err)
		fmt.Println("Err Nil")
		return
	}

	if err := cmd.Start(); err != nil {
		log.Println(err)
		fmt.Println("Start")
		return
	}

	s := bufio.NewScanner(io.MultiReader(stdout, stderr))
	for s.Scan() {
		ws.WriteMessage(1, s.Bytes())
	}

	if err := cmd.Wait(); err != nil {
		log.Println(err)
		fmt.Println("Wait")
		return
	}

	ws.WriteMessage(1, []byte("Finished\n"))
}

func serveHtml(w http.ResponseWriter, r *http.Request) {
	// w.Write(html)
	html.Execute(w, "ws://"+r.Host+"/ws")
}
