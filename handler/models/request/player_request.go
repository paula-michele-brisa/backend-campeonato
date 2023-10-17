package request

type PlayerRequest struct {
	Id       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required,min=4,max=100"`
	Photo    string `json:"photo"`
	Height   int32  `json:"height"`
	Weight   int32  `json:"weight"`
	Age      int32  `json:"age"`
	Position string `json:"position"`
	Number   int32  `json:"number"`
	TeamID   string `json:"teamID" binding:"required"`
}

type PlayerUpdateRequest struct {
	Id       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required,min=4,max=100"`
	Photo    string `json:"photo"`
	Height   int32  `json:"height"`
	Weight   int32  `json:"weight"`
	Age      int32  `json:"age"`
	Position string `json:"position"`
	Number   int32  `json:"number"`
	TeamID   string `json:"teamID" binding:"required"`
}
