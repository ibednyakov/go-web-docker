package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"log"
)

type Item struct {
	product string `db:"product"`
	qty     int    `db:"qty"`
}

type Info struct {
	customer string `db:"customer"`
	item     Item   `db:"items"`
}
type OrderWithId struct {
	id    int  `db:"id"`
	order Info `db:"info"`
}

type OrdersMap map[string]interface{}

func (p OrdersMap) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}
func (p *OrdersMap) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}
	log.Println("Source: " + string(source))

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}
	log.Println(i)

	*p, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("Type assertion .(map[string]interface{}) failed.")
	}

	return nil
}

type Person struct {
	Fn string
	Ln string
}
type ColorGroup struct {
	ID1    int
	ID2    int
	Name   string
	Colors []string
	P      Person `json:"Person"`
}

type ColorGroupWithId struct {
	id    int        `db:"id"`
	group ColorGroup `db:"info"`
}

type ColorGroupMap map[string]interface{}

func (p ColorGroupMap) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}
func (p *ColorGroupMap) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}
	log.Println("Source: " + string(source))

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}
	log.Println(i)

	*p, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("Type assertion .(map[string]interface{}) failed.")
	}

	return nil
}
