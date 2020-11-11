package session

import (
	"fmt"
	"time"
)

const restDurationBetweenSides = time.Duration(2 * time.Second)

type exercise struct {
	Name, Description          string
	Occurrences                int
	WorkDuration, RestDuration time.Duration
	RepeatPerSide              bool
}

type routine []exercise

// Session of phy
type Session struct {
	Routine   routine
	StartTime time.Time
	EndTime   time.Time
}

// NewSession creates a new Session
func NewSession() Session {
	S := Session{}
	S.Routine = exampleRoutine()
	return S
}

// Start starts the session
func (S *Session) Start() {
	S.StartTime = time.Now()
	fmt.Printf("session started at: %v", S.StartTime)

	for _, ex := range S.Routine {
		ex.GoThrough()
	}

	S.sayGoodbye()
}

func (ex exercise) GoThrough() {
	ex.Declare()
	ex.ExecuteSides()
}

func (S Session) sayGoodbye() {
	fmt.Println("")
	fmt.Println("********")
	fmt.Println("Well Done! Goodbye")
}

func (ex exercise) Declare() {
	fmt.Println("")
	fmt.Println("******")
	fmt.Println("- " + ex.Name + " -")
	fmt.Println(ex.Description)
	fmt.Println("")
}

func (ex exercise) ExecuteSides() {
	if ex.RepeatPerSide {
		for i := 0; i < 2; i++ {
			ex.DeclareSide(i)
			ex.ExecuteOccurrences()
			time.Sleep(restDurationBetweenSides)
		}
	} else {
		ex.ExecuteOccurrences()
	}

}
func (ex exercise) DeclareSide(i int) {
	fmt.Println("")
	fmt.Println(">")
	fmt.Printf([]string{"Left", "Right"}[i])
}

func (ex exercise) ExecuteOccurrences() {
	for i := 0; i < ex.Occurrences; i++ {
		ex.Countdown()
		time.Sleep(ex.RestDuration)
	}
}

func (ex exercise) Countdown() {
	fmt.Println("")
	fmt.Printf("Start Exercise")
	time.Sleep(ex.WorkDuration)
	fmt.Println("")
	fmt.Printf("You can rest now")
}

func exampleRoutine() routine {
	return routine{
		exercise{
			Name:          "Standing",
			Description:   "Feet together, eyes open",
			Occurrences:   1,
			WorkDuration:  time.Duration(2 * time.Second),
			RestDuration:  time.Duration(2 * time.Second),
			RepeatPerSide: false,
		},
		exercise{
			Name:          "Standing",
			Description:   "Feet together, eyes closed",
			Occurrences:   1,
			WorkDuration:  time.Duration(3 * time.Second),
			RestDuration:  time.Duration(1 * time.Second),
			RepeatPerSide: false,
		},
		exercise{
			Name:          "Stand on one leg",
			Description:   "Stand on one leg, eyes open",
			Occurrences:   3,
			WorkDuration:  time.Duration(1 * time.Second),
			RestDuration:  time.Duration(1 * time.Second),
			RepeatPerSide: true,
		},
	}
}
