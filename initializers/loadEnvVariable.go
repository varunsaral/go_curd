package initializers

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func LoadEnvVariable() {
	// loads values from .env into the system
	if err := godotenv.Load(); 
	err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("PORT"))
}