package dto

import (
	"encoding/csv"
	"github.com/jszwec/csvutil"
	"io"
	"metis/domain/model"
	"metis/domain/types"
	"strings"
)

type ResultDto struct {
	ID             types.RaceID    `csv:"race_id"`
	HorseName      types.HorseName `csv:"馬"`
	Jockey         string          `csv:"騎手"`
	HorseNumber    int             `csv:"馬番"`
	Time           string          `csv:"走破時間"`
	Odds           float32         `csv:"オッズ"`
	Position       string          `csv:"通過順"`
	Result         string          `csv:"着順"`
	HorseWeight    string          `csv:"体重"`
	WeightGain     int             `csv:"体重変化"`
	Gender         string          `csv:"性"`
	Age            int             `csv:"齢"`
	Weight         float32         `csv:"斤量"`
	ThreeFurlong   float32         `csv:"上がり"`
	Favorite       int             `csv:"人気"`
	Name           string          `csv:"レース名"`
	Date           string          `csv:"日付"`
	Class          string          `csv:"クラス"`
	RaceType       string          `csv:"芝・ダート"`
	Distance       int             `csv:"距離"`
	Turn           string          `csv:"回り"`
	TrackCondition string          `csv:"馬場"`
	Weather        string          `csv:"天気"`
	RaceCourseID   int             `csv:"場id"`
	RaceCourse     string          `csv:"場名"`
}

type ResultDTOs []*ResultDto

func NewResultDTOs(rows string) (ResultDTOs, error) {

	csvReader := csv.NewReader(strings.NewReader(rows))

	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		return nil, err
	}
	var resultDtos ResultDTOs
	for {
		var d *ResultDto
		err := dec.Decode(&d)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		resultDtos = append(resultDtos, d)
	}
	return resultDtos, nil
}

func (ds ResultDTOs) ToHorses() model.Horses {
	var horses model.Horses
	for _, d := range ds {
		horses = append(horses, d.ToHorse())
	}

	return horses
}

func (d *ResultDto) ToHorse() *model.Horse {
	return &model.Horse{
		Name:   d.HorseName,
		Gender: d.Gender,
		Age:    d.Age,
	}
}

func (ds ResultDTOs) ToRaces() model.Races {
	var races model.Races
	for _, d := range ds {
		races = append(races, d.ToRace())
	}

	return races
}

func (d *ResultDto) ToRace() *model.Race {
	return &model.Race{
		ID:             d.ID,
		Name:           d.Name,
		RaceCourse:     d.RaceCourse,
		Date:           d.Date,
		Class:          d.Class,
		RaceType:       d.RaceType,
		Distance:       d.Distance,
		Turn:           d.Turn,
		TrackCondition: d.TrackCondition,
		Weather:        d.Weather,
	}
}
