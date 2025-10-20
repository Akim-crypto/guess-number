package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
    green = "\033[32m"
    red   = "\033[31m"
    blue  = "\033[34m"
    reset = "\033[0m"
)

const welcome = `
==============================
  Guess the Number (CLI Game)
==============================
Правила:
• Компьютер загадывает число от 1 до 100.
• Вы выбираете сложность:
  - easy   : 10 попыток
  - medium : 7 попыток
  - hard   : 5 попыток
• После каждого ответа будет подсказка: больше или меньше.
• Игра заканчивается, когда число угадано или кончились попытки.
Удачи!
`


var dificultyAttempts = map[string]int{
	"easy":10,
	"medium":7,
	"hard":5,
}

func askDifficulty(r *bufio.Reader)int{
	for {
		fmt.Print("Выберите сложность (easy/medium/hard):")
		text, _ := r.ReadString('\n')
		choice := strings.ToLower(strings.TrimSpace(text))
		if attempts , ok := dificultyAttempts[choice]; ok {
			return attempts
		}
		fmt.Println("Некорректный выбор. Введите : easy,medium или hard")
	}
}

func askGuess(r *bufio.Reader) int{
	for {
		fmt.Print("Ваш ответ (число 1-100)")
		text,_:=r.ReadString('\n')
		text = strings.TrimSpace(text)
		n,err := strconv.Atoi(text) 
		if err == nil && n >= 1 && n <= 100{
			return n 
		}
		fmt.Println("Пожалуйста, введите целое число от 1 до 100")
	}
}


func play(r *bufio.Reader){
	fmt.Println(welcome)

	rand.Seed(time.Now().UnixNano())
	secret := rand.Intn(100) + 1 

	attemptsLeft := askDifficulty(r)
	attemptsUsed := 0

	for attemptsLeft > 0 {
		fmt.Printf("\nОсталось попыток: %d\n",attemptsLeft)
		quess := askGuess(r)
		attemptsUsed++
		attemptsLeft--

		if quess == secret {
			fmt.Printf("\n🎉 Поздравляю! Вы угадали число %d за %d попыток.\n", secret, attemptsUsed)
			return
		} else if quess < secret {
			fmt.Println(red +"Мое число больше"+ reset)
		} else {
			fmt.Println(red + "Мое число меньше" + reset)
		}
	}
	fmt.Printf("\n😔 Попытки закончились. Загаданное число было: %d\n", secret)
}


func wantReplay(r *bufio.Reader) bool {
	for {
		fmt.Print("\nСыграем еще раз? (y/n):")
		text,_ := r.ReadString('\n')
		ans := strings.ToLower(strings.TrimSpace(text))
		if ans == "y" || ans == "yes" || ans =="д" || ans == "да"{
			return  true
		}
		if ans == "n" || ans =="no" ||ans == "н" || ans =="нет" {
			return false
		}
		fmt.Println("Введите y/n.")
	}
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	for {
		play(reader)
		if !wantReplay(reader){
			fmt.Println("Спасибо за игру")
			return
		}
	}
}