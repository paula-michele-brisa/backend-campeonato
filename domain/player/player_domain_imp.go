package player

func NewPlayerDomain(name, photo, position, teamID string, number, weight, age, height int32) PlayerDomainInterface {
	return &playerDomain{
		name:     name,
		position: position,
		teamID:   teamID,
		number:   number,
		weight:   weight,
		height:   height,
		age:      age,
		photo:    photo,
	}
}
