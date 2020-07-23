package poker

import (
	"fmt"
	"github.com/hewiefreeman/GopherGameServer/actions"
	"log"
	"strconv"
)

func createCustomActions() {
	// CREATES UID GENERATION ACTION
	guestUIDCustomAction()
	roomCodeCustomAction()

	// Creates User Input Actions While in Lobby!
	createPokerRoomCustomAction()

	// Creates User Input Actions While In Game!
	takeSeatCustomAction() // Registers when the users sits down.
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

func guestUIDCustomAction() {
	UIDList = make(map[int]bool)
	UIDStruct.generated = UIDList
	UIDStruct.maxNumber = 9999999 // 7 digit long possible unique ids TODO ADD CHANGE MAYBE
	log.Println("Creating Custom Actions")
	err := actions.New("generateGuest", actions.DataTypeString, actionGenerateGuestUsername)
	if err != nil {
		log.Println(err)
		return
	}
}

func roomCodeCustomAction() {
	err := actions.New("sendRoomCode", actions.DataTypeString, actionValidateRoomCode)
	if err != nil {
		log.Println(err)
		return
	}
}
func actionCreatePokerRoom(actionData interface{}, client *actions.Client) {
	log.Println("Seating a user")
	roomUIDList = make(map[int]bool)
	roomUID.generated = roomUIDList
	roomUID.maxNumber = 9999999999

	UserCreatePokerRoom(actionData)
	//client.Respond()
}

func actionUserTookSeat(actionData interface{}, client *actions.Client) {
	log.Println("Seating a user")
}

func actionValidateRoomCode(actionData interface{}, client *actions.Client) {
	log.Println(actionData) // it does the job but nothing much else, it can do a lot more like checks and validation
	str := fmt.Sprintf("%v", actionData)
	validateRoomCode(str)
	client.Respond(actionData, actions.NoError()) // TODO Improve on this to make it smarter,
}

func actionGenerateGuestUsername(actionData interface{}, client *actions.Client) {
	UID := strconv.Itoa(UIDStruct.Int())
	guestUsername := "guest" + UID
	log.Println("Generating Guest Username: ", guestUsername)
	log.Println("total possible unique guest names left!", calculateTotalPossibleUIDS(UIDStruct))
	//client.Respond(UID + actionData.(string) + "!", actions.NoError())
	client.Respond(guestUsername, actions.NoError())
}
