package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//AllRooms is the global hashmap for the server
var AllRooms RoomMap

//CreteRoomRequestHandler creates a room and return roomID
func CreteRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	roomID := AllRooms.CreateRoom()

	type resp struct {
		RoomID string `json:"room_id"`
	}
	log.Println(AllRooms.Map)
	json.NewEncoder(w).Encode(resp{RoomID: roomID})
}

//JoinRoomRequestHandler will join the client in particular room
func JoinRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write("test")
	fmt.Fprintf(w, "Test")
}
