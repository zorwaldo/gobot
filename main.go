package main

import (
	"awesomeProject/database"
	"awesomeProject/model"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
	}
	Db *gorm.DB
}

var BotMap map[string]*tgbotapi.BotAPI
var config Config

func sendMessages() {
	fmt.Sprintf("CRON STARTED")
	var mailings []model.Mailing
	query := fmt.Sprintf(`select m.start_hour, m.mailint_type, c.chat_id, b.token from mailing m
		left outer join public.bots_in_chats bic on bic.id = m.botchat_id
		left join bot b on b.id = bic.botid
		left join chat c on bic.chatid = c.chat_id`)
	//where start_hour = date_part('HOUR', now());
	config.Db.Raw(query).Scan(&mailings)

	for _, i := range mailings {
		fmt.Println(i.ChatId)
		fmt.Println(i.Token)
		fmt.Println(i.MailingType)
		fmt.Println(i.StartHour)

		directory := "img/zior/"

		files, err := ioutil.ReadDir(directory)
		if err != nil {
			fmt.Println("Ошибка при чтении директории:", err)
			return
		}

		var images [][]byte

		fmt.Println("Файлы в директории:")
		for _, file := range files {
			if file.IsDir() {
				fmt.Println("Директория:", file.Name())
			} else {
				fmt.Println("Файл:", file.Name())
			}
			photo, err := ioutil.ReadFile(directory + file.Name())
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			images = append(images, photo)
		}

		photo1, err := ioutil.ReadFile(directory + "4.jpg")
		//photo2, err := ioutil.ReadFile(directory + "2.jpg")
		//photo3, err := ioutil.ReadFile(directory + "3.jpg")
		if err != nil {
			panic(err)
		}
		photo11 := tgbotapi.FileBytes{
			Name:  "picture",
			Bytes: photo1,
		}
		//
		//photo12 := tgbotapi.FileBytes{
		//	Name:  "picture",
		//	Bytes: photo2,
		//}
		//
		//photo13 := tgbotapi.FileBytes{
		//	Name:  "picture",
		//	Bytes: photo3,
		//}
		arr := []tgbotapi.FileBytes{photo11, photo11, photo11}
		rand.Seed(time.Now().UnixNano())
		// Получение случайного индекса
		randomIndex := rand.Intn(len(arr))
		log.Print(randomIndex)
		fmt.Sprintf(strconv.FormatInt(i.ChatId, 10))
		msg := tgbotapi.NewPhoto(i.ChatId, arr[randomIndex])
		msg.Caption = "Хо-хо-хо С новим годом"
		//msg := tgbotapi.NewMessage(-698982981, "Батя сказал пора спать")
		send, err := BotMap[i.Token].Send(msg)
		if err != nil {
			return
		}
		log.Print(send)
	}
}

func Init() {
	// Открываем файл с конфигурацией
	file, err := os.Open("config.yml")
	if err != nil {
		panic(fmt.Errorf("ошибка при открытии файла конфигурации: %s", err))
	}
	defer file.Close()

	// Декодируем содержимое файла YAML в структуру Config

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(fmt.Errorf("ошибка при декодировании файла конфигурации: %s", err))
	}

	// Формируем строку подключения к базе данных PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.Database.Host,
		config.Database.Username,
		config.Database.Password,
		config.Database.Name,
		config.Database.Port)
	config.Db = database.Init(dsn)

	var bots []model.Bot
	query := "SELECT * FROM bot"
	config.Db.Raw(query).Scan(&bots)

	BotMap = make(map[string]*tgbotapi.BotAPI)

	for _, s := range bots {
		bot, err := tgbotapi.NewBotAPI(s.Token)
		if err != nil {
			log.Printf("errror", err)
			panic(err)
		}
		BotMap[s.Token] = bot
	}

	for key := range BotMap {
		fmt.Println(key)
	}

}

func main() {
	Init()

	c := cron.New()
	c.AddFunc("*/1 * * * *", sendMessages)

	c.Start()
	//cron := cron.New()
	//
	//// Добавляем задачу в cron на выполнение каждый час в 5 минут
	//_, err := cron.AddFunc("*/1 * * * *", sendMessages)
	//if err != nil {
	//	fmt.Println("Ошибка при добавлении задачи в cron:", err)
	//	return
	//}
	//
	//cron.Start()
	for {
		time.Sleep(1 * time.Second)
	}
	// Теперь у вас есть подключение к базе данных PostgreSQL через Gorm
	// Можете выполнять запросы или другие операции с базой данных через db
}

func notMain() {
	bot, err := tgbotapi.NewBotAPI("6136660352:AAGxS-sEaNvbux1uQH_PM6EFzh_rwsxWIGY")

	//bot, err := tgbotapi.NewBotAPI("6242637551:AAGgUkDQh8fMOP9mw6z8nCzeckCZZ_SrP9s")
	if err != nil {
		log.Printf("errror", err)
		panic(err)
	}

	bot.Debug = true

	directory := "ziorNg/"
	photo1, err := ioutil.ReadFile(directory + "4.jpg")
	//photo2, err := ioutil.ReadFile(directory + "2.jpg")
	//photo3, err := ioutil.ReadFile(directory + "3.jpg")
	if err != nil {
		panic(err)
	}
	photo11 := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photo1,
	}
	//
	//photo12 := tgbotapi.FileBytes{
	//	Name:  "picture",
	//	Bytes: photo2,
	//}
	//
	//photo13 := tgbotapi.FileBytes{
	//	Name:  "picture",
	//	Bytes: photo3,
	//}
	arr := []tgbotapi.FileBytes{photo11, photo11, photo11}

	// Инициализация генератора случайных чисел
	for {
		rand.Seed(time.Now().UnixNano())
		// Получение случайного индекса
		randomIndex := rand.Intn(len(arr))
		log.Print(randomIndex)
		msg := tgbotapi.NewPhoto(-4005521799, arr[randomIndex])
		msg.Caption = "Хо-хо-хо С новим годом"
		//msg := tgbotapi.NewMessage(-698982981, "Батя сказал пора спать")
		send, err := bot.Send(msg)
		if err != nil {
			return
		}
		time.Sleep(10 * time.Second)
		log.Print(send)
	}

}
