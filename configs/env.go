package configs

import (
	"log"
	"os"
)

func EnvCloudName() string {
	return getEnv("CLOUDINARY_CLOUD_NAME")
}

func EnvCloudAPIKey() string {
	return getEnv("CLOUDINARY_API_KEY")
}

func EnvCloudAPISecret() string {
	return getEnv("CLOUDINARY_API_SECRET")
}

func EnvCloudUploadFolder() string {
	return getEnv("CLOUDINARY_UPLOAD_FOLDER")
}

func EnvPort() string {
	return getEnv("IMAGESERVICE_PORT")
}

func EnvIP() string {
	return getEnv("IMAGESERVICE_IP")
}

func EnvGWPort() string {
	return getEnv("GATEWAY_PORT")
}

func EnvGWIP() string {
	return getEnv("GATEWAY_IP")
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	log.Default().Printf("failed to get env %s\n", key)
	return ""
}
