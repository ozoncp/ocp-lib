package tasks

// EventType - тип события
type EventType uint

// DifficultyType - сложность задачи
type DifficultyType uint

const (
	DifficultyBegin DifficultyType = iota
	DifficultyLow
	DifficultyMedium
	DifficultyHigh
)

const (
	Created EventType = iota
	Removed
	Updated
)

// Event - событие о задачах
type Event struct {
	Type EventType
	Body map[string]interface{}
}
