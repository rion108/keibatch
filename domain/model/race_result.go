package model

import "metis/domain/types"

type RaceResult struct {
	ID            int
	RaceMappingID types.RaceMappingID
	Jockey        string
	HorseNumber   int
	Time          string
	Odds          float32
	Position      string
	Result        string
	HorseWeight   string
	WeightGain    int
	Gender        string
	Age           int
	Weight        float32
	ThreeFurlong  float32
	Favorite      int
	RaceCourse    string
}

type RaceResults Slice[*RaceResult]
