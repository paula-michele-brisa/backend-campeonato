package response

type PlayerResponse struct {
	Name     string `json:"name"`
	Photo    string `json:"photo"`
	Height   int32  `json:"height"`
	Weight   int32  `json:"weight"`
	Age      int32  `json:"age"`
	Position string `json:"position"`
	Number   int32  `json:"number"`
	TeamId   int32  `json:"teamId"`
}
