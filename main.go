package main

import (
	"github.com/cdgProcessor/dbWriter/db"
	"github.com/cdgProcessor/dbWriter/logger"
	"github.com/cdgProcessor/dbWriter/messageQ"
	"go.uber.org/zap"
)

func main() {
	logger.InitLogger("./logs/dbWriter.log")
	zap.L().Info("Processor sms to DB starting...")

	in2dbChan := make(chan []byte)

	go messageQ.MQRead(in2dbChan, "inboundSMS", "smsToDB", "consumeToDB")

	db.Writer(in2dbChan)
}
