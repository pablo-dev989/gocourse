package intermediate

import (
	"log"

	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		log.Println("Error in initializing Zap logger")
	}

	defer logger.Sync()

	logger.Info("This is a info message")

	logger.Info("User logged in", zap.String("username", "Simon GOlden"), zap.String("method", "GET"))

}
