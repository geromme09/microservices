package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ride-sharing/shared/contracts"
	"ride-sharing/shared/util"

	"github.com/gorilla/websocket"
)

// upgrading an HTTP connection to a WebSocket connection
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleRiderWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Websocket upgrade failed %v", err)
		return
	}
	defer conn.Close()

	userID := r.URL.Query().Get("userID")
	if userID == "" {
		fmt.Println("userID is required")
		return
	}

	for {
		//read message from rider
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Read message failed: %v", err)
			break
		}
		fmt.Printf("Received message from rider %s: %s\n", userID, string(message))

	}

}

func handleDriverWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Websocket upgrade failed %v", err)
		return
	}
	defer conn.Close()

	userID := r.URL.Query().Get("userID")
	if userID == "" {
		fmt.Println("userID is required")
		return
	}

	packageSlug := r.URL.Query().Get("packageSlug")
	if packageSlug == "" {
		fmt.Println("packageSlug is required")
		return
	}

	driver := Driver{
		Id:             userID,
		Name:           "Geromme Beligon",
		ProfilePicture: util.GetRandomAvatar(1),
		CarPlate:       "ABC123",
		PackageSlug:    packageSlug,
	}

	data, err := json.Marshal(driver)
	if err != nil {
		fmt.Printf("Marshal failed: %v", err)
		return
	}

	msg := contracts.WSDriverMessage{
		Type: "driver.cmd.register",
		Data: data,
	}

	if err := conn.WriteJSON(msg); err != nil {
		fmt.Printf("Write message failed: %v", err)
		return
	}

	//we can check the package slug if its valid

	for {
		//read message from driver
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Read message failed: %v", err)
			break
		}
		fmt.Printf("Received message from rider %s: %s\n", userID, string(message))

	}

}

type Driver struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	ProfilePicture string `json:"profilePicture"`
	CarPlate       string `json:"carPlate"`
	PackageSlug    string `json:"packageSlug"`
}
