package storage

import (
	"RestGo/structs"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func ReadStorage() {

	dec := json.NewDecoder(strings.NewReader("storage/storage.json"))
	for {
		var m structs.TodoStorage
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Description, m.Done)
	}
}
