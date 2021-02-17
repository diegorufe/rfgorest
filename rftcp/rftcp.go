package rftcp

import (
	"bufio"
	"fmt"
	"net"
	"rfgocore/utils/utilsbytes"
	"rfgocore/utils/utilsstring"
	"rfgorest/rftcp/constants"
)

// RFTcp : struct for store data tcp
type RFTcp struct {
}

// NewRFTcp : method for instance new RFTcp
func NewRFTcp() *RFTcp {
	var rftcp *RFTcp = new(RFTcp)

	return rftcp
}

// handleConnection method for handle conecction and realice busines logics
func handleConnection(connection net.Conn, rftcp *RFTcp) {

	fmt.Printf("Serving %s\n", connection.RemoteAddr().String())

	// TODO check security connection for RemoteAddr. If not accept connexion send error

loopHandleConnection:
	for {
		// Read data for connection
		netData, err := bufio.NewReader(connection).ReadString(utilsbytes.BreakLine)

		// If have error break connection
		if err != nil {
			fmt.Println(err)
			break loopHandleConnection
		}

		// Proccess command
		processCommandResult, err := proccessCommand(netData, connection, rftcp)

		// If have error or stop command break connection
		if err != nil || processCommandResult == constants.StopCommand {
			// Print error
			if err != nil {
				fmt.Println(err)
			}
			// Break loop connection
			break loopHandleConnection
		}

	}

	// Close conection
	connection.Close()
}

// proccessCommand : method for porcess command
func proccessCommand(netData string, connection net.Conn, rftcp *RFTcp) (string, error) {

	parsedNetData := utilsstring.TrimAllSpace(string(netData))

	switch parsedNetData {
	default:
		// TODO check send keep alive
		sendKeepAlive(connection, rftcp)
		break
	}

	return parsedNetData, nil
}

// sendKeepAlive method for send keep alive command
func sendKeepAlive(connection net.Conn, rftcp *RFTcp) {
	connection.Write([]byte(string(constants.KeepAliveCommand)))
}
