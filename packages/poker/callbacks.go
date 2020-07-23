package poker

import (
	"fmt"
	"github.com/hewiefreeman/GopherGameServer"
	"net/http"
)
//////////////// Main Server Call Backs ////////////////
func createServerCallbacks(){
	// Creates clientConnected Call back
	setErr := gopher.SetClientConnectCallback(clientConnected)
	if setErr != nil {
		fmt.Println(setErr)
	}
}
func clientConnected(*http.ResponseWriter, *http.Request) bool {
	// Runs BEFORE a client finishes handshaking server, but after
	// the initial HTTP connection.
	//
	// Returns a boolean, which when false will prevent the client
	// from finishing connecting, and send back an http.StatusForbidden
	// message.
	//log.Println("Client Connected!")
	return true
}


//////////////// Game Call Backs ////////////////

