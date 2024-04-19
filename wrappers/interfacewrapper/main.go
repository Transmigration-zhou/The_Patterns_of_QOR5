package main

import "fmt"

type Ecommerce interface {
	MakeOrder(userId string, productCode string, quantity int, price float64) (id string)
	ShipOrder(userId string, orderId string) bool
	Refund(userId string, orderId string) bool
}

type EcommerceDefault struct {
	// implementation of Ecommerce interface
}

func (e *EcommerceDefault) MakeOrder(userId string, productCode string, quantity int, price float64) (id string) {
	return "1"
}

func (e *EcommerceDefault) ShipOrder(userId string, orderId string) bool {
	return true
}

func (e *EcommerceDefault) Refund(userId string, orderId string) bool {
	return true
}

func main() {
	// create an instance of EcommerceLogger with an instance of EcommerceDefault
	e := EcommerceLogger{Ecommerce: &EcommerceDefault{}}
	// call methods on EcommerceLogger instance
	e.MakeOrder("123", "ABC123", 1, 10.0)
	e.ShipOrder("123", "456")
	e.Refund("123", "789")
}

type EcommerceLogger struct {
	Ecommerce
}

func (e *EcommerceLogger) MakeOrder(userId string, productCode string, quantity int, price float64) (id string) {
	fmt.Println("MakeOrder called", userId, productCode, quantity, price)
	id = e.Ecommerce.MakeOrder(userId, productCode, quantity, price)
	fmt.Println("MakeOrder returned", id)
	return id
}

func (e *EcommerceLogger) ShipOrder(userId string, orderId string) bool {
	fmt.Println("ShipOrder called", userId, orderId)
	r := e.Ecommerce.ShipOrder(userId, orderId)
	fmt.Println("ShipOrder returned", r)
	return r
}

func (e *EcommerceLogger) Refund(userId string, orderId string) bool {
	fmt.Println("Refund called", userId, orderId)
	r := e.Ecommerce.Refund(userId, orderId)
	fmt.Println("Refund returned", r)
	return r
}
