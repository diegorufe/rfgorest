package rftcp

import (
	"bufio"
	"fmt"
	"net"
	"rfgocore/logger"
	"rfgocore/utils/utilsbytes"
	"rfgocore/utils/utilsstring"
	"rfgorest/rftcp/beans"
	"rfgorest/rftcp/constants"
)

// RFTcp : struct for store data tcp
type RFTcp struct {
	Properties beans.RFTcpProperties

	// Function for chec secure access
	FunctionCheckSecureAccess func(connection net.Conn, rftcp *RFTcp) (bool, error)

	// Function for process data recived
	FunctionProcessDataReceived func(connection net.Conn, rftcp *RFTcp, parsedNetData string) error
}

// NewRFTcp : method for instance new RFTcp
func NewRFTcp() *RFTcp {
	var rftcp *RFTcp = new(RFTcp)

	// Init default port
	rftcp.Properties.Port = 5000

	return rftcp
}

// Listen : method for start server on port
func (rftcp *RFTcp) Listen() {
	var hostAndPort string = rftcp.Properties.Host +
		":" +
		utilsstring.IntToString(rftcp.Properties.Port)

	connection, err := net.Listen("tcp4", hostAndPort)

	if err != nil {
		if logger.IsErrorEnabled() {
			logger.Error(err)
		}
		return
	}

	// Close connection on error
	defer connection.Close()

	for {
		clientConnection, err := connection.Accept()
		if err != nil {
			if logger.IsErrorEnabled() {
				logger.Error(err)
			}
			return
		}

		// Handle connection client to go rutine (thread)
		go handleClientConnection(clientConnection, rftcp)
	}
}

// handleClientConnection method for handle conecction and realice busines logics
func handleClientConnection(connection net.Conn, rftcp *RFTcp) {

	if logger.IsDebugEnabled() {
		logger.Debug(fmt.Sprintf("Serving %s\n", connection.RemoteAddr().String()))
	}

	// Check security connection
	if rftcp.FunctionCheckSecureAccess != nil {
		secured, err := rftcp.FunctionCheckSecureAccess(connection, rftcp)

		if !secured || err != nil {

			if logger.IsErrorEnabled() {
				logger.Error(fmt.Sprintf("Connection not secured"))
			}

			if err != nil {
				if logger.IsErrorEnabled() {
					logger.Error(err)
				}
			}

			connection.Close()

			return
		}
	}

loopHandleConnection:
	for {
		// Read data for connection
		netData, err := bufio.NewReader(connection).ReadString(utilsbytes.BreakLine)

		// If have error break connection
		if err != nil {
			if logger.IsErrorEnabled() {
				logger.Error(err)
			}
			break loopHandleConnection
		}

		// Proccess command
		processCommandResult, err := proccessCommand(netData, connection, rftcp)

		// If have error or stop command break connection
		if err != nil || processCommandResult == constants.StopCommand {
			// Print error
			if err != nil {
				if logger.IsErrorEnabled() {
					logger.Error(err)
				}
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

	var err error = nil

	switch parsedNetData {
	case constants.KeepAliveCommand:
		sendKeepAlive(connection)
		break

	default:
		if rftcp.FunctionProcessDataReceived != nil {
			err = rftcp.FunctionProcessDataReceived(connection, rftcp, parsedNetData)
		}
		break
	}

	return parsedNetData, err
}

// SendCommandToClient method for send command to client
func SendCommandToClient(connection net.Conn, command string) {
	connection.Write([]byte(command))
}

// sendKeepAlive method for send keep alive command
func sendKeepAlive(connection net.Conn) {
	SendCommandToClient(connection, string(constants.KeepAliveCommand))
}
