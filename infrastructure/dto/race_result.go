package dto

import (
	"time"
)

type RaceResultDto struct {
	ID            int     `gorm:"autoIncrement column:id"`
	RaceMappingID int     `gorm:"column:race_mapping_id"`
	Jockey        string  `gorm:"column:jockey"`
	HorseNumber   int     `gorm:"column:horse_number"`
	Time          string  `gorm:"column:time"`
	Odds          float32 `gorm:"column:odds"`
	Position      string  `gorm:"column:position"`
	Result        string  `gorm:"column:result"`
	HorseWeight   string  `gorm:"column:horse_weight"`
	WeightGain    int     `gorm:"column:weight_gain"`
	Gender        string  `gorm:"column:gender"`
	Age           int     `gorm:"column:age"`
	Weight        float32 `gorm:"column:weight"`
	ThreeFurlong  float32 `gorm:"column:three_furlong"`
	Favorite      int     `gorm:"column:favorite"`
	RaceCourse    string  `gorm:"column:race_course"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type RaceResultDTOs []*RaceResultDto

func NewRaceResultsDTOs(results ResultDTOs, raceHorseMappings RaceHorseMappingDTOs) RaceResultDTOs {
	var ds RaceResultDTOs
	for _, result := range results {
		for _, raceHorseMapping := range raceHorseMappings {
			if raceHorseMapping.RaceID == result.ID && raceHorseMapping.HorseName == result.HorseName {
				raceResult := &RaceResultDto{
					RaceMappingID: raceHorseMapping.ID,
					Jockey:        result.Jockey,
					HorseNumber:   result.HorseNumber,
					Time:          result.Time,
					Odds:          result.Odds,
					Position:      result.Position,
					Result:        result.Result,
					HorseWeight:   result.HorseWeight,
					WeightGain:    result.WeightGain,
					Gender:        result.Gender,
					Age:           result.Age,
					Weight:        result.Weight,
					ThreeFurlong:  result.ThreeFurlong,
					Favorite:      result.Favorite,
					RaceCourse:    result.RaceCourse,
					UpdatedAt:     time.Now(),
				}
				ds = append(ds, raceResult)
			}
		}
	}
	return ds
}

func (*RaceResultDto) TableName() string {
	return "race_results"
}
