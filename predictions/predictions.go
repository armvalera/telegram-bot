package predictions

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var Keywords map[string][]string
var PredictionsMap map[string][]string

func LoadData() {
	log.Println("Загрузка данных...")
	loadKeywords()
	loadPredictions()
}

func loadKeywords() {
	file, err := os.Open("./predictions/keywords.json")
	if err != nil {
		log.Printf("Ошибка открытия файла с ключевыми словами: %v\n", err)
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&Keywords)
	if err != nil {
		log.Printf("Ошибка декодирования JSON с ключевыми словами: %v\n", err)
		return
	}

	log.Printf("Ключевые слова успешно загружены: %+v\n", Keywords)
}

func loadPredictions() {
	file, err := os.Open("./predictions/predictions.json")
	if err != nil {
		log.Printf("Ошибка открытия файла с предсказаниями: %v\n", err)
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&PredictionsMap)
	if err != nil {
		log.Printf("Ошибка декодирования JSON с предсказаниями: %v\n", err)
		return
	}

	log.Printf("Предсказания успешно загружены: %+v\n", PredictionsMap)
}

func ProcessMessage(msg string) (string, error) {
	msg = strings.ToLower(msg)
	log.Printf("Обрабатываем сообщение: %s\n", msg)

	for topic, words := range Keywords {
		for _, word := range words {
			if strings.Contains(msg, word) {
				log.Printf("Найдено ключевое слово: %s для темы: %s\n", word, topic)
				return GetPredictionForTopic(topic)
			}
		}
	}

	log.Println("Ключевое слово не найдено.")
	return "", fmt.Errorf("тема не найдена")
}

func GetPredictionForTopic(topic string) (string, error) {
	predictions, exists := PredictionsMap[topic]
	if !exists {
		return "", fmt.Errorf("тема '%s' не найдена", topic)
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(predictions))

	return predictions[randomIndex], nil
}

func GetRandomPrediction() string {
	var allPredictions []string
	for _, predictions := range PredictionsMap {
		allPredictions = append(allPredictions, predictions...)
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(allPredictions))

	return allPredictions[randomIndex]
}