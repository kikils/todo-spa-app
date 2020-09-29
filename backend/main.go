package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/kikils/golang-todo/infrastructure"
	"github.com/rs/cors"
)

const location = "Asia/Tokyo"

func initTimeZone() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

func main() {
	initTimeZone()
	port, err := strconv.Atoi(os.Args[1])
	addr := strconv.Itoa(port)
	if err != nil {
		log.Fatal("Please input port.")
	}
	mux := infrastructure.SetUpRouting()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
	})
	handler := c.Handler(mux)
	fmt.Printf("Start Server! Listen to %s port\n", addr)
	log.Fatal(http.ListenAndServe(":"+addr, infrastructure.LogHandler(handler)))
}
