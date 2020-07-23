package poker

import (
	"fmt"
	"github.com/jacz24/GoPackages/gameUtils"
	"github.com/hewiefreeman/GopherGameServer/actions"
	"github.com/jacz24/GoPackages/table"
	"log"
	"strconv"
)

func createCustomActions() {
	// CREATES UID GENERATION ACTION
	guestUIDCustomAction()
	roomCodeCustomAction()

	// Creates User Input Actions While in Lobby!
	table.CreatePokerRoomCustomAction()

	// Creates User Input Actions While In Game!
	table.TakeSeatCustomAction() // Registers when the users sits down.
}

func guestUIDCustomAction() {
	gameUtils.UIDList = make(map[int]bool)
	gameUtils.UIDStruct.Generated = UIDList
	gameUtils.UIDStruct.MaxNumber = 9999999 // 7 digit long possible unique ids TODO ADD CHANGE MAYBE
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
	log.Println("total possible unique guest names left!", CalculateTotalPossibleUIDS(UIDStruct))
	//client.Respond(UID + actionData.(string) + "!", actions.NoError())
	client.Respond(guestUsername, actions.NoError())
}
