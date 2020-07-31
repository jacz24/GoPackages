package lobby

import (
	"fmt"
	"github.com/hewiefreeman/GopherGameServer/actions"
	"github.com/hewiefreeman/GopherGameServer/core"
	"github.com/jacz24/GoPackages/pokkkker"
	"log"
	"strconv"
	"strings"
)


func actionValidateRoomCode(actionData interface{}, client *actions.Client) {
	log.Println(actionData) // it does the job but nothing much else, it can do a lot more like checks and validation
	str := fmt.Sprintf("%v", actionData)
	validateRoomCode(str)
	client.Respond(actionData, actions.NoError()) // TODO Improve on this to make it smarter,
}

func actionCreateTable(actionData interface{}, client *actions.Client){
	s := make([]string, 4)

	str := fmt.Sprintf("%v", actionData)
	s = strings.Split(str, ",")
	private, PrivErr := strconv.ParseBool(s[2])
	max, MaxErr := strconv.Atoi(s[3])
	if PrivErr != nil {
		log.Println(PrivErr)
		return
	}
	if MaxErr != nil {
		log.Println(MaxErr)
		return
	}
	log.Println(private)
	log.Println(max)
	userCreatedRoom := pokkkker.UserCreatePokerRoom(s[0], private, max, client.User().Name())

	client.Respond(userCreatedRoom.Name(), actions.NoError())
}

func UserCreatePokerRoom(roomName string, isPrivate bool, maxUsers int, userOwner string) *core.Room{ // Unpacks the createrpokerRoom action and returns a created room
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
