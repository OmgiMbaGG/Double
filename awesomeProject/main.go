package main

import (
	"fmt"
	"strings"
	"time"
)

var global float64 = 2.28 // Глобальная переменная для всех функций (вне main)

func main() {
	fmt.Println("Версия Даблятора", global)
	Time := time.Now()
	formattedTime := Time.Format("15:04")
	fmt.Println("Текущее время:", formattedTime)

	Name := ""        //Введите ваше имя
	x := 0            //Введи целое число, которое хочешь удвоить
	a := 0            //Введи целое число, которое хочешь удвоить
	Age := 0          //Проверка на совершеннолетие, введите ваш возраст
	Continue := "нет" //Продолжить...сосал? да/нет

	Continue = strings.ToLower(Continue)

	var opa bool
	if Continue == "да" {
		opa = true
	} else if Continue == "нет" {
		opa = false
	} else {
		fmt.Println("Ответь на вопрос в строке Continue")
		return
	}

	if opa != true {
		fmt.Println("Если хотите продолжить, введите да в строке Continue")
		return
	}

	hour := time.Hour //Проверка времени, не сработает после 23:00
	if hour < 23 {
		fmt.Println("После 23:00 не работаем")
	}

	defer func() {
		fmt.Println("Если хотите продолжить, введите да в строке Continue")
	}()
	defer func() { fmt.Println("За выбор нашего Даблятора!") }() // defer в стеках
	defer func() { fmt.Println(Name) }()
	defer func() { fmt.Println("Спасибо") }() // Анонимные функции, будут выполняться в конце main (стеками)
	defer func() { fmt.Println() }()

	if Age < 18 {
		fmt.Println("Извините, наш Даблятор только для совершеннолетних")
		return
	}

	if Name == "" {
		hi(Name)
		return
	}

	hi(Name)

	if x == 0 || a == 0 {
		fmt.Println("Нечего даблить, введите 2 целых числа в строки х =  a =")
		return
	}

	e, r, s := dobble(x, a)
	fmt.Println("Дабл x =", e)
	fmt.Println("Дабл a =", r)
	fmt.Println("Вдруг нужна сумма чисел =", s)
}
func dobble(x int, a int) (int, int, int) { // возвратная функция
	e := x * 2
	r := a * 2
	s := x + a
	return e, r, s
}

func hi(Name string) { // не возвратная функция
	if Name == "" {
		fmt.Println("Введите имя для пользования Даблятора")
		return
	}
	fmt.Println("Приветствую вас в Дабляторе", Name, "!")
}
