package model

import "metis/domain/types"

type Horse struct {
	Name   types.HorseName
	Gender string
	Age    int
}

type Horses Slice[*Horse]

func (ms Horses) GetHorseNames() types.HorseNames {
	var horseNames types.HorseNames
	m := make(map[types.HorseName]bool)
	for _, horse := range ms {
		if !m[horse.Name] {
			m[horse.Name] = true
			horseNames = append(horseNames, horse.Name)
		}
	}

	return horseNames
}
