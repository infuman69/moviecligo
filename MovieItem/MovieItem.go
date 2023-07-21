package movieitem

type MovieItem struct {
	Adult bool `json:"adult"`

	Backdrop_path string `json:"backdrop_path"`

	Genre_ids []int `json:"genre_ids"`

	Id int `json:"id"`

	Original_language string `json:"original_language"`

	Original_title string `json:"original_title"`

	Overview string `json:"overview"`

	Popularity float32 `json:"popularity"`

	Poster_path string `json:"poster_path"`

	Release_date string `json:"release_date"`

	Title string `json:"title"`

	Video bool `json:"video"`

	Vote_average float32 `json:"vote_average"`

	Vote_count int `json:"vote_count"`

}