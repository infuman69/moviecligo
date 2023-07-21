package main

import (
	"encoding/json"
	"fmt"
	movieitem "moviecligo/MovieItem"
	"net/http"
	"os"
)

func main() {
	var api_key string = os.Getenv("API_KEY")

	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/popular?api_key=%s", api_key)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Add("Application", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJlMGM0ODdkMWJjN2I5MzU1MzlkYjBiOGEwYmFiMjc0YSIsInN1YiI6IjYwNjIzMTVkMTEzMGJkMDAyOWFhYTM2OSIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.ErOS9GUjyH3rhu7KpbvqfgnCOO9m5-swcvJXUmWaJmk")

	res, err := http.DefaultClient.Do(req)

	// fmt.Println(res)

	if err != nil {
		panic(err)
	}

	var movieResponse movieitem.MovieResponse

	err = json.NewDecoder(res.Body).Decode(&movieResponse)

	if err != nil {

		panic(err)

	}
	
	defer res.Body.Close()

	var movieItems []movieitem.MovieItem

	movieItems = append(movieItems, movieResponse.Results...)


}
