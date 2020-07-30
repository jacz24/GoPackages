package lobby

import (
	"fmt"
	"github.com/hewiefreeman/GopherGameServer/actions"
	"log"
)


func actionValidateRoomCode(actionData interface{}, client *actions.Client) {
	log.Println(actionData) // it does the job but nothing much else, it can do a lot more like checks and validation
	str := fmt.Sprintf("%v", actionData)
	validateRoomCode(str)
	client.Respond(actionData, actions.NoError()) // TODO Improve on this to make it smarter,
}

func actionCreateTable(actionData interface{}, client *actions.Client){
	log.Println("creating table for user ", actionData)

}

func CreateTableCustomAction(){
	err := actions.New("sendTableSetup", actions.DataTypeString, actionValidateRoomCode)
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
