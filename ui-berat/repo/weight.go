package repo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"sort"

	model "github.com/salimkun/sirclo-test/ui-berat/models"
	"github.com/schwarmco/go-pagination"
)

type WeighRepo interface {
	AddWeight(*model.Weight) error
	UpdateWeight(*model.Weight) error
	DeleteWeight(string) error
	ViewAllWeight(page, limit int) *model.ListWeight
}

type WeightData struct {
	Data []*model.Weight `json:"data"`
}

func (p *WeightData) AddWeight(wg *model.Weight) error {
	dt := p.ViewAllWeight(0, 0)
	for _, i := range dt.Data {
		if wg.Date == i.Date {
			return errors.New("error date already input")
		}
	}

	wg.Range = wg.Max - wg.Min
	dt.Data = append(dt.Data, wg)
	p.Data = dt.Data
	// write file
	file, _ := json.MarshalIndent(p, "", " ")
	_ = ioutil.WriteFile("weight-db.txt", file, 0644)
	return nil
}

func (p *WeightData) UpdateWeight(wg *model.Weight) error {
	dt := p.ViewAllWeight(0, 0)
	match := false
	for c, i := range dt.Data {
		if wg.Date == i.Date {
			dt.Data[c] = wg
			match = true
		}
	}

	if !match {
		return errors.New("error data not found")
	}

	wg.Range = wg.Max - wg.Min
	p.Data = dt.Data
	// write file
	file, _ := json.MarshalIndent(p, "", " ")
	_ = ioutil.WriteFile("weight-db.txt", file, 0644)
	return nil
}

func (p *WeightData) DeleteWeight(date string) error {
	dt := p.ViewAllWeight(0, 0)
	match := false
	newData := []*model.Weight{}
	for _, i := range dt.Data {
		if date == i.Date {
			match = true
		} else {
			newData = append(newData, i)
		}
	}

	if !match {
		return errors.New("error data not found")
	}

	p.Data = newData
	// write file
	file, _ := json.MarshalIndent(p, "", " ")
	_ = ioutil.WriteFile("weight-db.txt", file, 0644)
	return nil
}

func (p *WeightData) ViewAllWeight(page, limit int) *model.ListWeight {
	// get all data
	data, err := ioutil.ReadFile("weight-db.txt")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	dataWeight := WeightData{}
	_ = json.Unmarshal([]byte(data), &dataWeight)

	sort.SliceStable(dataWeight.Data, func(i, j int) bool {
		return dataWeight.Data[i].Date > dataWeight.Data[j].Date
	})
	if page == 0 {
		page = 0
	}

	if limit == 0 {
		limit = len(dataWeight.Data)
	}

	totalData := int32(len(dataWeight.Data))
	totalPage := pagination.Calculate(int(page), int(limit), len(dataWeight.Data))
	resp := &model.ListWeight{
		Data: paginate(dataWeight.Data, page, limit),
		Meta: model.Meta{
			Page:      int32(page),
			Limit:     int32(limit),
			TotalData: totalData,
			TotalPage: int32(totalPage.NumPages),
		},
	}

	return resp
}

func paginate(x []*model.Weight, skip int, size int) []*model.Weight {
	if skip > len(x) {
		skip = len(x)
	}

	end := skip + size
	if end > len(x) {
		end = len(x)
	}

	return x[skip:end]
}
