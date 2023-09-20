package model

type Stat struct {
	Event  TEvent
	Kind   int
	Pos    Pos
	UserID int
	Score  int
	Jotai  [][]int
}

type Pos struct {
	X int
	Y int
}

type Input struct {
	Input  TInput
	UserID int
}

type TInput int

const (
	Up TInput = iota
	Left
	Right
	Bottom
)

func (i TInput) String() string {
	switch i {
	case Up:
		return "Up"
	case Left:
		return "Left"
	case Right:
		return "Right"
	case Bottom:
		return "Bottom"
	default:
		return "Unknown"
	}
}

type TEvent int

const (
	Clean TEvent = iota
	New
	Over
	Move
	R_move
	Stop
	Create
	Spin
)

func (e TEvent) String() string {
	switch e {
	case Clean:
		return "Clean"
	case New:
		return "New"
	case Over:
		return "Over"
	case Move:
		return "Move"
	case R_move:
		return "R"
	case Stop:
		return "Stop"
	case Create:
		return "Create"
	case Spin:
		return "Spin"
	default:
		return "Unknown"
	}
}
