package main

import (
	"cyoa"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := flag.String("file", "gopher.json", "the JSON file with cyoa game")
	flag.Parse()
	fmt.Println(*filename)

	// Opening the json file
	f, err := os.Open(*filename)
	if err != nil {
		log.Println(err)
	}

	// decode the json file
	d := json.NewDecoder(f)
	var story cyoa.Story
	if err := d.Decode(&story); err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v\n", story)
}
