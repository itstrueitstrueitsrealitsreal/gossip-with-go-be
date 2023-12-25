package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/itstrueitstrueitsrealitsreal/gossip-with-go-be/internal/router"
)

func main() {
	r := router.Setup()
	fmt.Print("Listening on port 8000 at https://gossip-with-go.onrender.com!\n")

	log.Fatalln(http.ListenAndServe(":8000", r))
}
