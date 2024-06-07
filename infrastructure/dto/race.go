package dto

import (
	"metis/domain/model"
	"metis/domain/types"
	"time"
)

type RaceDto struct {
	ID             types.RaceID `gorm:"column:id"`
	RaceCourse     string       `gorm:"column:race_course"`
	Name           string       `gorm:"column:name"`
	Date           string       `gorm:"column:date"`
	Class          string       `gorm:"column:class"`
	RaceType       string       `gorm:"column:type"`
	Distance       int          `gorm:"column:distance"`
	Turn           string       `gorm:"column:turn"`
	TrackCondition string       `gorm:"column:track_condition"`
	Weather        string       `gorm:"column:weather"`
	CreatedAt      *time.Time
	UpdatedAt      time.Time
}

func NewRaces(races model.Races) []*RaceDto {
	var racesDto []*RaceDto
	m := make(map[types.RaceID]bool)
	for _, race := range races {
		if !m[race.ID] {
			m[race.ID] = true
			racesDto = append(racesDto, newRace(race))
		}
	}
	return racesDto
}

func newRace(race *model.Race) *RaceDto {
	return &RaceDto{
		ID:             race.ID,
		Name:           race.Name,
		RaceCourse:     race.RaceCourse,
		Date:           race.Date,
		Class:          race.Class,
		RaceType:       race.RaceType,
		Distance:       race.Distance,
		Turn:           race.Turn,
		TrackCondition: race.TrackCondition,
		Weather:        race.Weather,
	}
}

func (d *RaceDto) TableName() string {
	return "races"
}
