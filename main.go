package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
	"github.com/gorilla/websocket"
)

var dbName string = "test"

func input() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press ESC to quit")
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyArrowUp {
			fmt.Println(" Up ")
			continue
		}
		if key == keyboard.KeyArrowDown {
			fmt.Println(" Down ")
			continue
		}
		if key == keyboard.KeyEsc {
			fmt.Println()
			break
		}
	}
}

func main() {

	input()

	socketUrl := "ws://localhost:1111/ws"

	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		fmt.Println("Error! :", err)
		return
	}
	defer conn.Close()

	// handle std input
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(dbName + "> ")
		inputQuery, _ := reader.ReadString('\n')
		if inputQuery == "" {
			println("zeroval")
			continue
		}

		// Send an echo packet every second
		err = conn.WriteMessage(websocket.TextMessage, []byte(inputQuery))
		if err != nil {
			fmt.Println("Error! :", err)
			return
		}

		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error in receive:", err)
			return
		}
		fmt.Println(string(msg))

	}
}
