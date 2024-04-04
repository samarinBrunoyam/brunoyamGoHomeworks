package main

import (
	"brunoyamGoHomeworks/goHomeWorks/lessionOne/div"
	"brunoyamGoHomeworks/goHomeWorks/lessionOne/divRem"
	"brunoyamGoHomeworks/goHomeWorks/lessionOne/multiply"
	"brunoyamGoHomeworks/goHomeWorks/lessionOne/sub"
	"brunoyamGoHomeworks/goHomeWorks/lessionOne/sum"
	"fmt"
)

func main() {

	// EASY
	a := 10
	b := 2
	a = divRem.DivRem(a, b)
	a = div.Div(10, 2)
	a = sum.Sum(a, b)
	a = sub.Sub(a, b)
	a = multiply.Multiply(a, b)
	fmt.Println(a)

	//MEDIUM
	// 1. Создаём карту
	products := make(map[string]int)

	// 2. Добавляем продукты с ценами
	products["Bred"] = 100
	products["Coffee"] = 499
	products["Beer"] = 140
	products["Ketchup"] = 78

	/* 3. Находим цену одного из продуктов
	по имени и выводим в терминал */
	fmt.Println(products["Coffee"])
	// 4. Удаляем один из продуктов и выводим измененную мапу
	delete(products, "Bred")
	fmt.Println(products)

	//HARD
	results := make(map[string]int)
	results["Умножение"] = multiply.Multiply(a, b)
	results["Вычитание"] = sub.Sub(a, b)
	results["Сложение"] = sum.Sum(a, b)
	results["Деление"] = div.Div(a, b)
	results["Остаток"] = divRem.DivRem(a, b)
	fmt.Println(results)

	for i, v := range results {
		fmt.Println(i, v)
	}

}
