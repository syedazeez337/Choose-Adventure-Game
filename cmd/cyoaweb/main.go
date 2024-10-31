package main

import (
	"cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "The powrt to start the CYOA web application on")
	filename := flag.String("file", "gopher.json", "the JSON file with cyoa game")
	flag.Parse()
	fmt.Println(*filename)

	// Opening the json file
	f, err := os.Open(*filename)
	if err != nil {
		log.Println(err)
	}

	// decode the json file
	story, err := cyoa.JsonStory(f)
	if err != nil {
		log.Println(err)
	}

	h := cyoa.NewHandler(story, cyoa.WithTemplate(nil))
	fmt.Printf("Starting the server at: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
