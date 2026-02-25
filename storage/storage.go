package storage

import (
	"RestGo/structs"
	"encoding/json"
	"log"
	"os"
)

func ReadStorage() []structs.TodoStorage {
	file, err := os.ReadFile("storage/storage.json")
	if err != nil {
		log.Fatal(err)
		return []structs.TodoStorage{}
	}
	var structToDo = []structs.TodoStorage{}
	err = json.Unmarshal(file, &structToDo)
	if err != nil {
		log.Fatal(err)
		return []structs.TodoStorage{}
	}
	return structToDo

}
