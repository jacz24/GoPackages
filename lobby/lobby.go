package lobby

import (
	"github.com/hewiefreeman/GopherGameServer/core"
	"log"
)
func CreateLobbyRoomType() *core.Room{
	// Make a Room type and set broadcasts and callbacks
	lobbyRoomType := core.NewRoomType("Lobby", true)// TODO FIX NEEDING TO LOGIN FIRST!
	lobbyRoomType.EnableBroadcastUserEnter().EnableBroadcastUserLeave().
		SetUserEnterCallback(onEnterLobby).SetUserLeaveCallback(onLeaveLobby)

	// Create lobby room on startup so people auto join in
	lobbyroom, roomErr := core.NewRoom("mainLobby", "Lobby", false, 0, "")
	if roomErr != nil {
		log.Println("Error while opening Room:", roomErr)
		return nil
	} else {
		return lobbyroom
	}
}

