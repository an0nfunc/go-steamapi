package dota

type GameMode int

const (
	AnyMode = -1

	AllPick GameMode = iota
	SingleDraft
	AllRandom
	RandomDraft
	CaptainsDraft
	CaptainsMode
	DeathMode
	Diretide
	ReverseCaptainsMode
	TheGreeviling
	TutorialGame
	MidOnly
	LeastPlayed
	NewPlayerPool
	CompendiumMatchmaking
)

type Skill uint

const (
	AnySkill Skill = iota
	Normal
	High
	VeryHigh
)

type LeaverStatus uint

const (
	None LeaverStatus = iota
	Disconnected
	DisconnectedTooLong
	Abandoned
	AFK
	NeverConnected
	NeverConnectedTooLong
)

type LobbyType int

const (
	Invalid LobbyType = -1

	PublicMatchMaking LobbyType = iota
	Practice
	Tournament
	Tutorial
	Coop
	TeamMatch
	SoloQueue
)

type PlayerSlot uint8

func (d PlayerSlot) IsDire() bool {
	if d&(1<<7) > 0 {
		return true
	}
	return false
}

func (d PlayerSlot) GetPosition() (p uint) {
	p = uint(d & ((1 << 7) - 1))
	return
}

type Team uint

const (
	Radiant Team = iota
	Dire
)

type TowerStatus uint16

//TODO: add methods that read information from bits

type BarracksStatus uint16

// TODO: add methods that read information from bits
