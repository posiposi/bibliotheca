package main

import (
	"fmt"
	"log"

	"github.com/posiposi/project/backend/api"
)

func main() {
	fmt.Println("Hello World")
	r := api.NewRouter()
	log.Print("Server is running on port 8080")
	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
