package response

type SongResponse struct {
	Id   int    `json: "id"`
	Name string `json: "name"`
}
