package poker

import (
	"github.com/hewiefreeman/GopherGameServer"
	"github.com/hewiefreeman/GopherGameServer/core"
	"github.com/jacz24/GoPackages/chat"
	"github.com/jacz24/GoPackages/lobby"
	"github.com/jacz24/GoPackages/table"
	"log"
)
func run() {

	settings := gopher.ServerSettings{
		ServerName:     "!s!",
		MaxConnections: 0,

		HostName:  "https://dijidocs.com/",
		HostAlias: "https://www.dijidocs.com/",
		IP:        "0.0.0.0",
		Port:      443,

		TLS: false,
		CertFile: "C:/Users/irwin/OneDrive/Documents/GitHub/pokkker.github.io/Cert.pem",
		PrivKeyFile: "C:/Users/irwin/Sync/MaxResearchDevelopment/pokkker/GopherGameServerEx/Priv.pem",

		AdminLogin: "admin",
		AdminPassword: "admin",
		//OriginOnly: true,
	}
	// Make a Room type and set broadcasts and callbacks
	createCustomActions() // Creates all custom actions required for client
	createServerCallbacks()
	mainchat := chat.CreateChatRoomType()
	lobbyRoomType := lobby.CreateLobbyRoomType()
	//lobbyRoom := lobby.

	//chat := createChat()
	log.Println(mainchat)
	log.Println(lobbyRoomType)
	//createPokerTable() // if button click
	table.CreatePokerTableType()

	openSecondTestTable() // FOR DEBUGGING!
	openTestTable() // FOR DEBUGGING!
	gopher.Start(&settings)
}

func Init(){
	run()
}

func openRoom(rName string, rtype string) *core.Room{
	// Open a Room
	room, roomErr := core.NewRoom(rName, rtype, false, 0, "")
	if roomErr != nil {
		log.Println("Error while opening Room:", roomErr)
		return nil
	} else {
		return room
	}
}

func openPrivateRoom(rName string, rtype string) *core.Room{ // TODO IDK IF THIS IS NEEDED
	// Open a Room
	room, roomErr := core.NewRoom(rName, rtype, true, 0, "")
	if roomErr != nil {
		log.Println("Error while opening Room:", roomErr)
		return nil
	} else {
		return room
	}
}




