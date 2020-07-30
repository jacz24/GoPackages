package table

import (
	"github.com/hewiefreeman/GopherGameServer/core"
	"log"
)

/*
WARNING: When you use a *Room object in your code, DO NOT dereference it. Instead,
there are many methods for *Room for maniupulating and retrieving any information about them you could possibly need.
Dereferencing them could cause data races in the Room fields that get locked by mutexes.
*/


// Creates A poker player by the user, TODO ADD SETUP CONFIGS LIKE MAX PLAYERS AND STUFF

func CreatePokerTableType() *core.RoomType{ // Serialize these for being able to create multiple
	// Make a Room type and set broadcasts and callbacks
	pokerRoomType := core.NewRoomType("PokerTable", false)// allows the players to create room!
	pokerRoomType.EnableBroadcastUserEnter().EnableBroadcastUserLeave().
		SetUserEnterCallback(onEnterPokerTable).SetUserLeaveCallback(onLeavePokerTable)

	// Create lobby room on startup so people auto join in
	//pokerRoom := openRoom("chat", "PokerTable")

	return pokerRoomType
}
func createComputerPokerTableType() *core.RoomType{ // Serialize these for being able to create multiple
	// Make a Room type and set broadcasts and callbacks
	computerPokerRoomType := core.NewRoomType("ComputerPokerTable", false)// allows the players to create room!
	computerPokerRoomType.EnableBroadcastUserEnter().EnableBroadcastUserLeave().
		SetUserEnterCallback(onEnterComputerPokerTable).SetUserLeaveCallback(onLeaveComputerPokerTable)
	//core.NewRoom("computerGame","ComputerPokerTable", true,8, []string{user.User().Name()})
	// Create lobby room on startup so people auto join in
	//pokerRoom := openRoom("chat", "PokerTable")

	return computerPokerRoomType
}
func onEnterComputerPokerTable(room *core.Room, user *core.RoomUser) {
	// Example of using parameters to send a welcome message to the entering User
	message := "Welcome To The Table Please Play Nicely!"
	messageErr := room.ServerMessage(message, core.ServerMessageNotice, []string{user.User().Name()})
	if messageErr != nil {
		log.Println("Error while messaging User:", messageErr)
	}
}

func onLeaveComputerPokerTable(room *core.Room, user *core.RoomUser) {
	// ...

	// To convert RoomUser to User:
	// u := user.User()
}
func onEnterPokerTable(room *core.Room, user *core.RoomUser) {
	// Example of using parameters to send a welcome message to the entering User
	message := "Welcome To The Table Please Play Nicely!"
	messageErr := room.ServerMessage(message, core.ServerMessageNotice, []string{user.User().Name()})
	if messageErr != nil {
		log.Println("Error while messaging User:", messageErr)
	}
}

func onLeavePokerTable(room *core.Room, user *core.RoomUser) {
	// ...

	// To convert RoomUser to User:
	// u := user.User()
}
