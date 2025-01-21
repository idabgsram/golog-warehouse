package golog

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

type LogWarehouse struct {
	channel      string
	username     string
	warehouseKey string
}

type LogData struct {
	WorkerChannel  string    `json:"worker_channel"`
	WorkerUsername string    `json:"worker_username"`
	Message        string    `json:"message"`
	Level          string    `json:"level"`
	Color          string    `json:"color"`
	Emoji          string    `json:"emoji"`
	Data           *string   `json:"data"`
	Exception      *string   `json:"exception"`
	StoredAt       time.Time `json:"stored_at"`
}

var Slack *LogWarehouse
var Log LogWarehouse

func connectWarehouse() {
	Log.warehouseKey = os.Getenv("GOLOG_WAREHOUSE_KEY")
	if Log.warehouseKey == "" {
		Log.warehouseKey = "golog_warehouse"
	}

	redisUrl := os.Getenv("GOLOG_REDIS_URL")
	if redisUrl != "" {
		opts, err := redis.ParseURL(redisUrl)
		if err != nil {
			panic(err)
		}

		RedisClient = redis.NewClient(opts)
		return
	}

	host := GetEnv("GOLOG_REDIS_HOST")
	pass := GetEnv("GOLOG_REDIS_PASSWORD")
	port := GetEnv("GOLOG_REDIS_PORT")

	opt := redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: pass,
		DB:       0,
	}

	RedisClient = redis.NewClient(&opt)
}

func New() {
	Log = LogWarehouse{
		channel:  GetEnv("GOLOG_CHANNEL"),
		username: GetEnv("GOLOG_USERNAME"),
	}
	Slack = &Log
	connectWarehouse()
}

func NewCustomInstance(channel string, username string, url string) {
	Log = LogWarehouse{
		username: username,
		channel:  channel,
	}
	Slack = &Log
	connectWarehouse()
}

func GetEnv(key string) string {
	env := os.Getenv(key)

	if env == "" || len(env) < 1 {
		log.Fatalf("Error : %s variable not found on your system, please add to environtment variable.", key)
	}

	return env
}

func (s *LogWarehouse) sendToWarehouse(payload LogData) {

	payload.WorkerUsername = s.username
	payload.WorkerChannel = s.channel
	payload.StoredAt = time.Now()

	pyl, _ := json.Marshal(payload)

	err := RedisClient.RPush(context.Background(), s.warehouseKey, string(pyl)).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func (s *LogWarehouse) compose(message string, messageType string, color string, emoji string, errors error) {

	payload := LogData{
		Message: message,
		Level:   messageType,
		Color:   color,
		Emoji:   emoji,
	}

	if errors != nil {
		exception := fmt.Sprintf("``` %s ```", errors.Error())
		payload.Exception = &exception
	}

	s.sendToWarehouse(payload)
}

func (s *LogWarehouse) composeWithData(message string, messageType string, color string, emoji string, data []byte, e error) {

	payload := LogData{
		Message: message,
		Level:   messageType,
		Color:   color,
		Emoji:   emoji,
	}

	dataMsg := fmt.Sprintf("``` %s ```", string(data))
	payload.Data = &dataMsg

	if e != nil {
		exception := fmt.Sprintf("``` %s ```", e.Error())
		payload.Exception = &exception
	}

	s.sendToWarehouse(payload)
}

func (s *LogWarehouse) Info(message string) {
	s.compose(message, "INFO", "#2eb886", ":ok_hand:", nil)
}

func (s *LogWarehouse) InfoWidthData(message string, data []byte) {
	s.composeWithData(message, "INFO", "#2eb886", ":ok_hand:", data, nil)
}

func (s *LogWarehouse) Error(message string, e error) {
	s.compose(message, "ERROR", "#a30200", ":bomb:", e)
}

func (s *LogWarehouse) ErrorWithData(message string, data []byte, e error) {

	s.composeWithData(message, "ERROR", "#a30200", ":bomb:", data, e)
}

func (s *LogWarehouse) Warning(message string, e error) {
	s.compose(message, "WARNING", "#ffc107", ":warning:", e)
}

func (s *LogWarehouse) WarningWithData(message string, data []byte, e error) {
	s.composeWithData(message, "WARNING", "#ffc107", ":warning:", data, e)
}
