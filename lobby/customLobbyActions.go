package lobby

import (
	"fmt"
	"github.com/hewiefreeman/GopherGameServer/actions"
	//jaczpoker "github.com/jacz24/GoPackages/pokkkker"
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
	//pokkkker.
	//userCreatedRoom := jaczpoker.UserCreatePokerRoom(s[0], private, max, client.User().Name())

	//client.Respond(userCreatedRoom.Name(), actions.NoError())
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
