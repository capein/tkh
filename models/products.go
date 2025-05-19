package models

import (
	"context"
	_ "embed"
	"encoding/json"
)

type Products struct {
	IDs      []string  `json:"ids,omitempty"`
	Products []Product `json:"products,omitempty"`
}

func (p *Products) GetKey(ctx context.Context) string {
	return ""
}

//go:embed Products.json
var list string

func (p *Products) Save(ctx context.Context) error {
	return nil
}

func (p *Products) GetData(ctx context.Context) error {
	if len(p.IDs) != 0 {
		pMap := make(map[string]bool)
		p.Products = make([]Product, 0, len(p.IDs))
		pr := []Product{}
		err := json.Unmarshal([]byte(list), &pr)
		if err != nil {
			return err
		}
		for _, id := range p.IDs {
			pMap[id] = true
		}
		for _, val := range pr {
			if pMap[val.Id] {
				p.Products = append(p.Products, val)
			}
		}
		return nil
	}
	return json.Unmarshal([]byte(list), &(p.Products))
}

func (p *Products) Delete(ctx context.Context) error {
	return nil
}

func (p *Products) String() string {
	return ""
}
