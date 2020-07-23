package poker

import (
	Utils "github.com/jacz24/GoPackages/gameUtils"
	"github.com/hewiefreeman/GopherGameServer/actions"
	"github.com/jacz24/GoPackages/table"
	"github.com/jacz24/GoPackages/lobby"
	"log"
	"strconv"
)

func createCustomActions() {
	// CREATES UID GENERATION ACTION
	guestUIDCustomAction()
	lobby.RoomCodeCustomAction()

	// Creates User Input Actions While in Lobby!
	table.CreatePokerRoomCustomAction()

	// Creates User Input Actions While In Game!
	table.TakeSeatCustomAction() // Registers when the users sits down.
}

func guestUIDCustomAction() {
	Utils.UIDList = make(map[int]bool)
	Utils.UIDStruct.Generated = Utils.UIDList
	Utils.UIDStruct.MaxNumber = 9999999 // 7 digit long possible unique ids TODO ADD CHANGE MAYBE
	log.Println("Creating Custom Actions")
	err := actions.New("generateGuest", actions.DataTypeString, actionGenerateGuestUsername)
	if err != nil {
		log.Println(err)
		return
	}
}


func actionGenerateGuestUsername(actionData interface{}, client *actions.Client) {
	UID := strconv.Itoa(Utils.UIDStruct.Int())
	guestUsername := "guest" + UID
	log.Println("Generating Guest Username: ", guestUsername)
	log.Println("total possible unique guest names left!", Utils.CalculateTotalPossibleUIDS(Utils.UIDStruct))
	//client.Respond(UID + actionData.(string) + "!", actions.NoError())
	client.Respond(guestUsername, actions.NoError())
}
