package main

import (
	"fmt"
	"github.com/k0kubun/pp"
	"strings"
	"time"
)

var global float64 = 2.28 // Глобальная переменная для всех функций (вне main)

func main() {
	fmt.Println("Версия Даблятора", global)

	Time := time.Now()
	formattedTime := Time.Format("15")

	hour := time.Hour //Проверка времени, не сработает после 22:00
	if hour > 22 {
		pp.Println("Текущее время:", formattedTime, "часов")
	} else {
		pp.Println("После 23:00 не работаем")
		return
	}

	Name := ""       //Введите ваше имя
	x := 0           //Введи целое число, которое хочешь удвоить
	a := 0           //Введи целое число, которое хочешь удвоить
	Age := 20        //Проверка на совершеннолетие, введите ваш возраст
	Continue := "да" //Продолжить...сосал? да/нет

	Continue = strings.ToLower(Continue)

	var opa bool
	if Continue == "да" {
		opa = true
	} else if Continue == "нет" {
		opa = false
	} else {
		pp.Println("Ответь на вопрос в строке Continue")
		return
	}

	if opa != true {
		pp.Println("Если хотите продолжить, введите да в строке №29 Continue")
		return
	}

	defer func() {
		pp.Println("Если хотите продолжить, введите да в строке Continue")
	}()
	defer func() { pp.Println("За выбор нашего Даблятора!") }() // defer в стеках
	defer func() { pp.Println(Name) }()
	defer func() { pp.Println("Спасибо") }() // Анонимные функции, будут выполняться в конце main (стеками)
	defer func() { pp.Println() }()

	if Age < 18 {
		pp.Println("Извините, наш Даблятор только для совершеннолетних")
		return
	}

	if Name == "" {
		hi(Name)
		return
	}

	hi(Name)

	if x == 0 || a == 0 {
		pp.Println("Нечего даблить, введите 2 целых числа в строки х =  a =")
		return
	}

	e, r, s := dobble(x, a)
	pp.Println("Дабл x =", e)
	pp.Println("Дабл a =", r)
	pp.Println("Вдруг нужна сумма чисел =", s)
}

func dobble(x int, a int) (int, int, int) { // возвратная функция
	e := x * 2
	r := a * 2
	s := x + a
	return e, r, s
}

func hi(Name string) { // не возвратная функция
	if Name == "" {
		pp.Println("Введите имя для пользования Даблятора")
		return
	}
	pp.Println("Приветствую вас в Дабляторе", Name, "!")
}
