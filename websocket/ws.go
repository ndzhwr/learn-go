package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// listenNotifier := func(w http.ResponseWriter, r *http.Response) {
// 	fmt.Println("🚀 Server running on port" + os.Getenv("PORT"))
// }

func main() {
	godotenv.Load()
	http.HandleFunc("/echo", func(res http.ResponseWriter, req *http.Request) {
		conn, _ := upgrader.Upgrade(res, req, nil)
		
		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			fmt.Print(conn)
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			conn.WriteMessage(msgType, msg)
		}

	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Host)
		http.ServeFile(w, r, "./websocket.html")
	})

	port := os.Getenv("PORT")
	exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://localhost:"+port+"/echo")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
