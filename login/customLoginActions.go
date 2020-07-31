package login

import (
	"github.com/hewiefreeman/GopherGameServer/actions"
	Utils "github.com/jacz24/GoPackages/gameUtils"
	"log"
	"strconv"
)
//var UIDList map[int]bool
var UIDStruct = Utils.UniqueRand{}

var UIDList = make(map[int]bool)
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
func actionGenerateGuestUsername(actionData interface{}, client *actions.Client) {
	UID := strconv.Itoa(UIDStruct.Int())
	guestUsername := "guest" + UID
	log.Println("Generating Guest Username: ", guestUsername)
	log.Println("total possible unique guest names left!", Utils.CalculateTotalPossibleUIDS(UIDStruct))
	//client.Respond(UID + actionData.(string) + "!", actions.NoError())
	client.Respond(guestUsername, actions.NoError())
}
