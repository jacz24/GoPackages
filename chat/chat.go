package chat

import (
	"fmt"
"github.com/hewiefreeman/GopherGameServer/core"

)
func CreateChatRoomType() *core.Room{
	// Make a Room type and set broadcasts and callbacks
	chatRoomType := core.NewRoomType("chat", true)
	chatRoomType.EnableBroadcastUserEnter().EnableBroadcastUserLeave().
		SetUserEnterCallback(onEnterChat).SetUserLeaveCallback(onLeaveChat)

	// Open a Room
	chatRoom := openNewChatRoom()
	return chatRoom
}

func openNewChatRoom() *core.Room{
	chatRoom, roomErr := core.NewRoom("chatExample", "chat", false, 0, "")
	if roomErr != nil {
		fmt.Println("Error while opening Room:", roomErr)
		return nil
	} else {
		return chatRoom
	}
}

//////////////// Chat Call Backs ////////////////
func onEnterChat(room *core.Room, user *core.RoomUser) {
	// Example of using parameters to send a welcome message to the entering User
	//log.Println("Entering Chat")
	message := "Welcome! Please read the chat room rules, and have fun!"
	messageErr := room.ServerMessage(message, core.ServerMessageNotice, []string{user.User().Name()})
	if messageErr != nil {
		fmt.Println("Error while messaging User:", messageErr)
	}
}

func onLeaveChat(room *core.Room, user *core.RoomUser) {
	// ...

	// To convert RoomUser to User:
	// u := user.User()
}
