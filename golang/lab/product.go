package main

type Product interface {
	GetPrice() float64
	SetPrice(price float64)
	ApplyDiscount(percent float64)
	GetDetails() string
}
