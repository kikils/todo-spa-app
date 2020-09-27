package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/kikils/golang-todo/infrastructure"
	"github.com/rs/cors"
)

func main() {
	port, err := strconv.Atoi(os.Args[1])
	addr := strconv.Itoa(port)
	if err != nil {
		log.Fatal("Please input port.")
	}
	mux := infrastructure.SetUpRouting()
	handler := cors.Default().Handler(mux)
	fmt.Printf("Start Server! Listen to %s port\n", addr)
	log.Fatal(http.ListenAndServe(":"+addr, infrastructure.LogHandler(handler)))
}
