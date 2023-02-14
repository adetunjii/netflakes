package model

type Movie struct {
	EpisodeID       int32        `json:"episode_id"`
	Title           string       `json:"title"`
	OpeningCrawl    string       `json:"opening_crawl"`
	Characters      []string     `json:"characters"`
	MovieCharacters []*Character `json:"movie_characters"`
	Created         string       `json:"created"`
	Url             string       `json:"url"`
	ReleaseDate     string       `json:"release_date"`
	CommentCount    int64        `json:"comment_count"`
}
