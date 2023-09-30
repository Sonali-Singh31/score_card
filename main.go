package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Score struct {
	Filters struct {
		DateFrom string `json:"dateFrom"`
		DateTo   string `json:"dateTo"`
	} `json:"filters"`
	ResultSet struct {
		Count int `json:"count"`
	} `json:"resultSet"`
	Matches []struct {
	} `json:"matches"`
}

func main() {
	// fmt.Println("Ready to go")

	res, err := http.Get("https://api.football-data.org/v4/matches")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Score Api not available to real time")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(body))
	var score Score
	err = json.Unmarshal(body, &score)
	if err != nil {
		panic(err)
	}
	fmt.Println(score)
	time, result := score.Filters, score.ResultSet
	fmt.Printf("%s, %s, %v\n", time.DateFrom, time.DateTo, result.Count)
}
