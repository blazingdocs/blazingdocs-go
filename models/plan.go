package models

type PlanModel struct {
	Id           string
	Name         string
	Price        float64
	PricePerUnit float64
	Quota        int
}
