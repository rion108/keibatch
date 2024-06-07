package dto

import (
	"metis/domain/model"
	"metis/domain/types"
	"time"
)

type HorseDto struct {
	ID        types.HorseID   `gorm:"primaryKey autoIncrement column:id"`
	Name      types.HorseName `gorm:"column:name"`
	Gender    string          `gorm:"column:gender"`
	Age       int             `gorm:"column:age"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewHorses(horses model.Horses) []*HorseDto {
	var horsesDto []*HorseDto
	for _, horse := range horses {
		horsesDto = append(horsesDto, newHorse(horse))
	}

	return horsesDto
}

func newHorse(horse *model.Horse) *HorseDto {
	return &HorseDto{
		Name:   horse.Name,
		Gender: horse.Gender,
		Age:    horse.Age,
	}
}

func (d *HorseDto) TableName() string {
	return "horses"
}
