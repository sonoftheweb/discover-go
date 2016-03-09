package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	i := "tt0372784"
	plot := "short"
	r := "json"

	res, err := http.Get("http://www.omdbapi.com/?i=" + i + "&plot=" + plot + "&r=" + r)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var m movieData
	err2 := json.Unmarshal(body, &m)
	if err2 != nil {
		fmt.Println(err2)
	}
	j, err3 := strconv.ParseFloat(m.imdbRating, 32)
	if err3 != nil {
		fmt.Println(err3)
	}
	j = (j / 10) * 100

	fmt.Printf("The movie : %s was released in %s - the IMDB rating is %d%% with %s votes\n", m.Title, m.Year, int(j), m.imdbVotes)
}
