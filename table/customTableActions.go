package table

import (
	"github.com/hewiefreeman/GopherGameServer/actions"
	Utils "github.com/jacz24/GoPackages/gameUtils"
	"log"
)

var roomUIDList map[int]bool

var roomUID = Utils.UniqueRand{}
func actionCreatePokerRoom(actionData interface{}, client *actions.Client) {
	log.Println("Seating a user")
	roomUIDList = make(map[int]bool)
	roomUID.Generated = roomUIDList
	roomUID.MaxNumber = 9999999999

	userCreatePokerRoom(actionData)
	//client.Respond()
}

func actionUserTookSeat(actionData interface{}, client *actions.Client) {
	log.Println("Seating a user")
}

func TakeSeatCustomAction() {
	err := actions.New("takeSeat", actions.DataTypeString, actionUserTookSeat)
	if err != nil {
		log.Println(err)
		return
	}
}

func CreatePokerRoomCustomAction(){
	err := actions.New("userCreatePokerRoom", actions.DataTypeString, actionCreatePokerRoom)
	if err != nil {
		log.Println(err)
		return
	}
}
