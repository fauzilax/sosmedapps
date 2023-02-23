package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvCloudName() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("CLOUDINARY_CLOUD_NAME")
}

func EnvCloudAPIKey() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("CLOUDINARY_API_KEY")
}

func EnvCloudAPISecret() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("CLOUDINARY_API_SECRET")
}

func EnvCloudUploadFolder() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
}
