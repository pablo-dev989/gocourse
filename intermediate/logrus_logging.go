package intermediate

import "github.com/sirupsen/logrus"

func main() {

	log := logrus.New()

	//Set log level
	log.SetLevel(logrus.InfoLevel)

	// Set log format
	log.SetFormatter(&logrus.JSONFormatter{})

	//Logging examples
	log.Info("This is a info message")
	log.Warn("This is a warning message")
	log.Error("This is a error message")

	log.WithFields(logrus.Fields{
		"username": "Simon Golden",
		"method":   "GET",
	}).Info("User logged in.")

}
