package anillkas

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Posts struct {
	PostData []Post `json:"posts"`
}

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func GetPost() Posts {
	dosya, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer dosya.Close()
	var PostList Posts
	PostArray, err := ioutil.ReadAll(dosya)
	json.Unmarshal(PostArray, &PostList)
	if err != nil {
		log.Fatal(err)
	}

	return PostList
}
