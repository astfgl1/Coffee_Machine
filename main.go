package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	water       int
	milk        int
	coffeeBeans int
	cups        int
	money       int
)

func printIngredient() {
	fmt.Println("The coffee machine has:")
	fmt.Println(strconv.Itoa(water) + " ml of water")
	fmt.Println(strconv.Itoa(milk) + " ml of milk")
	fmt.Println(strconv.Itoa(coffeeBeans) + " g of coffee beans")
	fmt.Println(strconv.Itoa(cups) + " disposable cups")
	fmt.Println("$" + strconv.Itoa(money) + " of money")
}

func initIngredient() {
	water = 400
	milk = 540
	coffeeBeans = 120
	cups = 9
	money = 550
}

func fill() {
	var w, m, cb, c int
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&w)
	water += w
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&m)
	milk += m
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&cb)
	coffeeBeans += cb
	fmt.Println("Write how many disposable cups of coffee you want to add:")
	fmt.Scan(&c)
	cups += c
}

func take() {
	fmt.Println("I gave you $" + strconv.Itoa(money))
	money = 0
}

func buyCoffee(waterRequest int, coffeeBeansRequest int, milkRequest int, cost int) {
	if waterRequest > water {
		fmt.Println("Sorry, not enough water!")
		return
	}
	if coffeeBeansRequest > coffeeBeans {
		fmt.Println("Sorry, not enough coffee beans!")
		return
	}
	if milkRequest > milk {
		fmt.Println("Sorry, not enough milk!")
		return
	}
	if cups == 0 {
		fmt.Println("Sorry, not enough cups!")
		return
	}
	water -= waterRequest
	coffeeBeans -= coffeeBeansRequest
	milk -= milkRequest
	cups -= 1
	money += cost
	fmt.Println("I have enough resources, making you a coffee!")

	return
}

func buy() {
	var coffee string
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	fmt.Scan(&coffee)

	switch coffee {
	case "1":
		buyCoffee(250, 16, 0, 4)
	case "2":
		buyCoffee(350, 20, 75, 7)
	case "3":
		buyCoffee(200, 12, 100, 6)
	case "back":
		return
	}
}

func action() {
	var action string
	fmt.Println("Write action (buy, fill, take, remaining, exit):")
	fmt.Scan(&action)

	switch action {
	case "buy":
		buy()
	case "fill":
		fill()
	case "take":
		take()
	case "remaining":
		printIngredient()
	case "exit":
		os.Exit(0)
	}
}

func main() {
	initIngredient()
	for {
		action()
	}
}
