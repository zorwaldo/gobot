package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

func main() {
	//bot, err := tgbotapi.NewBotAPI("6136660352:AAGxS-sEaNvbux1uQH_PM6EFzh_rwsxWIGY")
	bot, err := tgbotapi.NewBotAPI("6242637551:AAGgUkDQh8fMOP9mw6z8nCzeckCZZ_SrP9s")
	if err != nil {
		log.Printf("errror", err)
		panic(err)
	}

	bot.Debug = true

	directory := "hz/"
	photo1, err := ioutil.ReadFile(directory + "1.jpg")
	photo2, err := ioutil.ReadFile(directory + "2.jpg")
	photo3, err := ioutil.ReadFile(directory + "3.jpg")
	if err != nil {
		panic(err)
	}
	photo11 := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photo1,
	}

	photo12 := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photo2,
	}

	photo13 := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photo3,
	}
	arr := []tgbotapi.FileBytes{photo13, photo11, photo12}

	// Инициализация генератора случайных чисел
	//-4005521799 c Никой
	//-698982981 с Пацанами
	//-4051045296 с Леной
	for {
		rand.Seed(time.Now().UnixNano())
		// Получение случайного индекса
		randomIndex := rand.Intn(len(arr))
		log.Print(randomIndex)
		msg := tgbotapi.NewPhoto(-4051045296, arr[randomIndex])
		msg.Caption = "Мяу"
		//msg := tgbotapi.NewMessage(-698982981, "Батя сказал пора спать")
		send, err := bot.Send(msg)
		if err != nil {
			return
		}
		time.Sleep(5 * time.Second)
		log.Print(send)
	}

}
