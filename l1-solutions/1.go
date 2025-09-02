package main

import (
	"fmt"
	"time"
)

// Определяем структуру Human с полями и методами
type Human struct {
	firstName string
	lastName  string
	age       int
	gender    string
	height    float64
	weight    float64
	birthDate time.Time
}

func (h Human) Eat() {
	fmt.Println("I'm eating")
}

func (h Human) Sleep() {
	fmt.Println("I'm sleeping")
}

func (h Human) Code() {
	fmt.Println("I'm coding")
}

type Action struct {
	Human // реализуем встраивание структуры Human в структуру Action
}

func main() {
	action := Action{Human: Human{firstName: "Mike"}} // Задаем имя человеку
	fmt.Println(action.firstName) // Теперь action обладаем теми же полями
	action.Code() // и методами
}
