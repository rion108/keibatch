package model

import (
	"metis/domain/types"
)

type Race struct {
	ID             types.RaceID
	Name           string
	RaceCourse     string
	Date           string
	Class          string
	RaceType       string
	Distance       int
	Turn           string
	TrackCondition string
	Weather        string
}

type Races Slice[*Race]
