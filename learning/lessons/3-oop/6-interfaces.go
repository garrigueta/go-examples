package main

import (
	"fmt"
)

type Food interface {
	Nutrition() string
	FoodType() string
}

type Apple struct {
}

func (a Apple) Nutrition() string {

	return "Apples are crab heavy!"
}

func (a Apple) FoodType() string {

	return "Apples are fruit"
}

type Celery struct {
}

func (a Celery) Nutrition() string {

	return "Celery has zero everything!"
}

func (a Celery) FoodType() string {

	return "Celery is a vegetable"
}

func main() {

	fmt.Println("###managing interfaces###")

	fmt.Println("#operate#")
	foods := []Food{Apple{}, Celery{}}
	for _, f := range foods {
		fmt.Println(f.Nutrition(), f.FoodType())
	}
}
