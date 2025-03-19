package tester

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Глобальные переменные для хранения данных
var keywords map[string][]string
var predictionsMap map[string][]string

// Функция для тестирования загрузки данных
func TestLoadData() {
	// Загружаем данные
	loadKeywords()
	loadPredictions()
}

// Открытие JSON-файла для ключевых слов
func loadKeywords() {
	file, err := os.Open("/Users/valerijiskandaran/telegram-bot/predictions/keywords.json")
	if err != nil {
		log.Printf("Ошибка открытия файла с ключевыми словами: %v\n", err)
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&keywords)
	if err != nil {
		log.Printf("Ошибка декодирования JSON с ключевыми словами: %v\n", err)
		return
	}

	fmt.Printf("Ключевые слова успешно загружены: %+v\n", keywords)
}

// Открытие JSON-файла для предсказаний
func loadPredictions() {
	file, err := os.Open("/Users/valerijiskandaran/telegram-bot/predictions/predictions.json")
	if err != nil {
		log.Printf("Ошибка открытия файла с предсказаниями: %v\n", err)
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&predictionsMap)
	if err != nil {
		log.Printf("Ошибка декодирования JSON с предсказаниями: %v\n", err)
		return
	}

	fmt.Printf("Предсказания успешно загружены: %+v\n", predictionsMap)
}