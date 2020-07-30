package lobby

import (
	"fmt"
	"github.com/hewiefreeman/GopherGameServer/actions"
	"github.com/hewiefreeman/GopherGameServer/core"
	"log"
	"strings"
)


func actionValidateRoomCode(actionData interface{}, client *actions.Client) {
	log.Println(actionData) // it does the job but nothing much else, it can do a lot more like checks and validation
	str := fmt.Sprintf("%v", actionData)
	validateRoomCode(str)
	client.Respond(actionData, actions.NoError()) // TODO Improve on this to make it smarter,
}

func actionCreateTable(actionData interface{}, client *actions.Client){
	str := fmt.Sprintf("%v", actionData)
	formatted := strings.Split(str, ",")
	log.Println("its wokring", formatted)
	//userCreatePokerRoom()
}
func userCreatePokerRoom(roomName string, isPrivate bool, maxUsers int, userOwner string) *core.Room{ // Unpacks the createrpokerRoom action and returns a created room
	log.Println("Creating table ", roomName)
	room ,roomErr := core.NewRoom(roomName, "PokerTable", isPrivate, maxUsers, userOwner)
	if roomErr != nil {
		log.Println("Error while opening Room:", roomErr)
		return nil
	} else {
		return room
	}
}

func CreateTableCustomAction(){
	err := actions.New("sendTableSetup", actions.DataTypeString, actionCreateTable)
	if err != nil {
		log.Println(err)
		return
	}
}

func RoomCodeCustomAction() {
	err := actions.New("sendRoomCode", actions.DataTypeString, actionValidateRoomCode)
	if err != nil {
		log.Println(err)
		return
	}
}
