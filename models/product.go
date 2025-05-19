package models

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"tkh/logger"
)

var ErrNoData = errors.New("no data found")

type Product struct {
	Id       string  `json:"id"`
	Image    Image   `json:"image"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
}

func (p *Product) Save(ctx context.Context) error {

	return nil
}

func (p *Product) GetKey(ctx context.Context) string {
	return ""
}

func (p *Product) GetData(ctx context.Context) error {
	pL := []Product{}
	err := json.Unmarshal([]byte(list), &pL)
	if err != nil {
		logger.Println("error while unmarshalling json", err)
		return err
	}
	index, err := strconv.Atoi(p.Id)
	if err != nil {
		logger.Println("error while unmarshalling json", err)
		return err
	}
	if index < 1 || index > len(pL) {
		return ErrNoData
	}
	*p = pL[index-1]
	return nil
}

func (p *Product) Delete(ctx context.Context) error {
	return nil
}
func (p *Product) String() string {
	return ""
}

type Image struct {
	Thumbnail string `json:"thumbnail"`
	Mobile    string `json:"mobile"`
	Tablet    string `json:"tablet"`
	Desktop   string `json:"desktop"`
}
