package main

import "github.com/DistributedClocks/GoVector/govec"

func main() {
	//Initialize GoVector logger
	Logger := govec.InitGoVector("MyProcess", "LogFile", govec.GetDefaultConfig())
	
	//Encode message, and update vector clock
	messagepayload := []byte("samplepayload")
	vectorclockmessage := Logger.PrepareSend("Sending Message", messagepayload, govec.GetDefaultLogOptions())
	
	//send message
	connection.Write(vectorclockmessage)

	//In Receiving Process
	connection.Read(vectorclockmessage)
	//Decode message, and update local vector clock with received clock
	Logger.UnpackReceive("Receiving Message", vectorclockmessage, &messagepayload, govec.GetDefaultLogOptions())

	//Log a local event
	Logger.LogLocalEvent("Example Complete", govec.GetDefaultLogOptions())
}