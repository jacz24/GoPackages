package lobby

import (
	"github.com/hewiefreeman/GopherGameServer/core"
	"log"
)

//////////////// Lobby Call Backs ////////////////


func OnEnterLobby(room *core.Room, user *core.RoomUser) {
	// Example of using parameters to send a welcome message to the entering User
	message := "Welcome To Pokkker! " + user.User().Name() + " You are currently at the Lobby"
	messageErr := room.ServerMessage(message, core.ServerMessageNotice, []string{user.User().Name()})
	if messageErr != nil {
		log.Println("Error while messaging User:", messageErr)
	}
}

func OnLeaveLobby(room *core.Room, user *core.RoomUser) {
	message := "Good luck!"
	messageErr := room.ServerMessage(message, core.ServerMessageNotice, []string{user.User().Name()})
	if messageErr != nil {
		log.Println("Error while messaging User:", messageErr)
	}
	// To convert RoomUser to User:
	// u := user.User()
}
