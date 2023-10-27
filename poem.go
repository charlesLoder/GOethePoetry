package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type Poem struct {
	Title string `json:"title"`
	Author string `json:"author"`
	Lines []string `json:"lines"`
	LineCount string `json:"linecount"`
}

func GetRandomPoem() (Poem, error)  {
	resp, err := http.Get("http://poetrydb.org/random")

	if err != nil {
		return Poem{}, err
	}
	defer resp.Body.Close()

	// Read the body into a byte slice
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Poem{}, err
	}

	var data []Poem

	err = json.Unmarshal(body, &data)
	if err != nil {
		return Poem{}, err
	}

	if len(data) == 0 {
		return Poem{}, errors.New("no poems found in the response")
	}

	log.Println("• Found a poem: ", data[0].Title)

	return data[0], nil
}