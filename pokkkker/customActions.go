package pokkkker

import (
	//"github.com/jacz24/GoPackages/login"
	"github.com/jacz24/GoPackages/lobby"
	"github.com/jacz24/GoPackages/table"
)

func createCustomActions() {
	// LOGIN CUSTOM ACTIONS
	//login.



	// LOBBY CUSTOM ACTIONS
	lobby.RoomCodeCustomAction()
	lobby.CreateTableCustomAction()

	// Creates User Input Actions While In Game!
	table.TakeSeatCustomAction() // Registers when the users sits down.
}
