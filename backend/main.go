package main

import (
	"fmt"
	"log"

	"github.com/posiposi/project/backend/presentation/router"
)

func main() {
	fmt.Println("Hello World")
	r := router.NewRouter()
	log.Print("Server is running on port 8080")
	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
