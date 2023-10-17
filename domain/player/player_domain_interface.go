package player

type PlayerDomainInterface interface {
	GetID() string
	GetName() string
	GetPhoto() string
	GetHeight() int32
	GetWeight() int32
	GetAge() int32
	GetPosition() string
	GetNumber() int32
	GetTeamID() string
}
