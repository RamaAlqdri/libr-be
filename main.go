package main

import (
	"fmt"
	"libr-be/config"
	"os"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	port := os.Getenv("APP_PORT")
	fmt.Println("Server running on port", port)
	// gqlgen akan disiapkan nanti
}
