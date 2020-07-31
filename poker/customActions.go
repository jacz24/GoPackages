package poker

import (
	"github.com/hewiefreeman/GopherGameServer/actions"
	"github.com/hewiefreeman/GopherGameServer/core"
	Utils "github.com/jacz24/GoPackages/gameUtils"
	"github.com/jacz24/GoPackages/lobby"
	"github.com/jacz24/GoPackages/table"
	"log"
	"strconv"
)

//var UIDList map[int]bool
var UIDStruct = Utils.UniqueRand{}

var UIDList = make(map[int]bool)

func createCustomActions() {
	// CREATES UID GENERATION ACTION
	guestUIDCustomAction()
	lobby.RoomCodeCustomAction()
	lobby.CreateTableCustomAction()

	// Creates User Input Actions While In Game!
	table.TakeSeatCustomAction() // Registers when the users sits down.
}

func guestUIDCustomAction() {
	UIDStruct.Generated = UIDList
	UIDStruct.MaxNumber = 9999999 // 7 digit long possible unique ids TODO ADD CHANGE MAYBE
	log.Println("Creating Custom Actions")
	err := actions.New("generateGuest", actions.DataTypeString, actionGenerateGuestUsername)
	if err != nil {
		log.Println(err)
		return
	}
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

func actionGenerateGuestUsername(actionData interface{}, client *actions.Client) {
	UID := strconv.Itoa(UIDStruct.Int())
	guestUsername := "guest" + UID
	log.Println("Generating Guest Username: ", guestUsername)
	log.Println("total possible unique guest names left!", Utils.CalculateTotalPossibleUIDS(UIDStruct))
	//client.Respond(UID + actionData.(string) + "!", actions.NoError())
	client.Respond(guestUsername, actions.NoError())
}
