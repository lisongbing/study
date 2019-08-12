package main

import (
	"os"
	"logrus"
)

// Create a new instance of the logger. You can have any number of instances.
var log2 = logrus.New()

func main() {
	// The API for setting attributes is a little different than the package level
	// exported logger. See Godoc.
	log2.Out = os.Stdout

	// You could set this to any `io.Writer` such as a file
	// file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err == nil {
	//  log.Out = file
	// } else {
	//  log.Info("Failed to log to file, using default stderr")
	// }

	log2.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
		"level": logrus.GetLevel(),
		"enable":logrus.IsLevelEnabled(logrus.FatalLevel),
	}).Info("A group of walrus emerges from the ocean")
	print_req("mylog","127.0.0.1")

}

func print_req(user_id,user_ip string){
	requestLogger := log2.WithFields(logrus.Fields{"request_id": user_id, "user_ip": user_ip})
	requestLogger.Info("something happened on that request")
	requestLogger.Warn("something not great happened")
}