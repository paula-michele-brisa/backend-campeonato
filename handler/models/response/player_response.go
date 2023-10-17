package response

type PlayerResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Photo    string `json:"photo"`
	Height   int32  `json:"height"`
	Weight   int32  `json:"weight"`
	Age      int32  `json:"age"`
	Position string `json:"position"`
	Number   int32  `json:"number"`
	TeamID   string `json:"teamID"`
}
