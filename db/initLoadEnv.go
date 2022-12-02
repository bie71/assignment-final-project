package db

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
)

func init() {
	const dirName = "assignment-final-project"
	projectName := regexp.MustCompile(`^(.*` + dirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Println(err)
		log.Fatalf("Error loading .env file")
	}
}
