package main

import (
	"log"
	"os"
	"telegram-bot/predictions"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

func TgBot() {
	apiKey := os.Getenv("TELEGRAM_API_KEY")
	if apiKey == "" {
		log.Fatal("TELEGRAM_API_KEY не установлен")
	}

	var err error
	bot, err = tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
}

func handleMessage(messageText string, chatID int64) {
	if len(predictions.PredictionsMap) == 0 {
		log.Println("Данные предсказаний не загружены.")
		msg := tgbotapi.NewMessage(chatID, "Ошибка: данные предсказаний не загружены.")
		if _, err := bot.Send(msg); err != nil {
			log.Println("Ошибка при отправке сообщения:", err)
		}
		return
	}

	topic, err := predictions.ProcessMessage(messageText)
	if err != nil {
		log.Println("Ошибка обработки сообщения:", err)
		topic = predictions.GetRandomPrediction()
	}

	response := "Ты спрашивал: " + messageText + "\nОтвет от оракула: " + topic
	msg := tgbotapi.NewMessage(chatID, response)
	if _, err := bot.Send(msg); err != nil {
		log.Println("Ошибка при отправке сообщения:", err)
	}
}

func main() {
	log.Println("Запуск бота...")

	// Загружаем данные для предсказаний
	log.Println("Загрузка данных...")
	predictions.LoadData()

	// Инициализируем бота
	log.Println("Инициализация бота...")
	TgBot()

	log.Println("Бот запущен и готов к работе.")

	// Создаем канал для получения обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// Обработка сообщений
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text == "/start" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID,
				"Приветствую, искатель ответов! Я — твой личный оракул. 🔮 Ты готов узнать, что ждёт тебя в будущем? Задай свой вопрос, и я подскажу, что приготовил для тебя мир.")
			if _, err := bot.Send(msg); err != nil {
				log.Println("Ошибка при отправке приветственного сообщения:", err)
			}
		} else if update.Message.Text == "/help" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID,
				"Доступные команды:\n/start - начать работу с ботом\n/help - получить справку")
			if _, err := bot.Send(msg); err != nil {
				log.Println("Ошибка при отправке сообщения:", err)
			}
		} else {
			handleMessage(update.Message.Text, update.Message.Chat.ID)
		}
	}
}