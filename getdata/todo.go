package anillkas

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Todos struct {
	TodoData []Todo `json:"todos"`
}

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func GetTodo() Todos {
	dosya, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer dosya.Close()

	arr, err := ioutil.ReadAll(dosya)
	if err != nil {
		log.Fatal(err)
	}

	var listeler Todos

	json.Unmarshal(arr, &listeler)

	return listeler

}
