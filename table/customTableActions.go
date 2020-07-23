package table

import (
	"github.com/hewiefreeman/GopherGameServer/actions"
	"github.com/jacz24/GoPackages/poker"
	"log"
)

var roomUIDList map[int]bool

var roomUID = poker.UniqueRand{}


func actionCreatePokerRoom(actionData interface{}, client *actions.Client) {
	log.Println("Seating a user")
	roomUIDList = make(map[int]bool)
	roomUID.generated = roomUIDList
	roomUID.maxNumber = 9999999999

	userCreatePokerRoom(actionData)
	//client.Respond()
}

func actionUserTookSeat(actionData interface{}, client *actions.Client) {
	log.Println("Seating a user")
}

func takeSeatCustomAction() {
	err := actions.New("takeSeat", actions.DataTypeString, actionUserTookSeat)
	if err != nil {
		log.Println(err)
		return
	}
}

func createPokerRoomCustomAction(){
	err := actions.New("userCreatePokerRoom", actions.DataTypeString, actionCreatePokerRoom)
	if err != nil {
		log.Println(err)
		return
	}
}
