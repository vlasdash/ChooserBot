package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/vlasdash/dating_bot/config"
	"log"
	"net/http"
)

const (
	startCommand     = "/start"
	findCommand      = "/find"
	findMovieCommand = "Фильм"
	findBookCommand  = "Книга"
	psychology       = "Психология"
	history          = "История"
	prose            = "Проза"
	poetry           = "Поэзия"
	horrors          = "Ужасы"
	adventure        = "Приключение"
	detectives       = "Детективы"
	fantastic        = "Фантастика"
)

type Handler struct {
	botToken        string
	commandHandlers map[string]func(*Update) error
}

type KeyboardMarkup struct {
	Keyboard [][]KeyboardButton `json:"keyboard"`
}

type KeyboardButton struct {
	Text string `json:"text"`
}

type ResponseBody struct {
	ChatID      int64           `json:"chat_id"`
	Text        string          `json:"text"`
	ReplyMarkup *KeyboardMarkup `json:"reply_markup,omitempty"`
}

type Update struct {
	Message struct {
		MessageID int64  `json:"message_id"`
		Text      string `json:"text"`
		Chat      struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

func NewHandler(token string) *Handler {
	h := &Handler{
		botToken: token,
	}

	h.commandHandlers = map[string]func(*Update) error{
		startCommand:     h.StartCommand,
		findCommand:      h.FindCommand,
		findMovieCommand: h.FindMovieCommand,
		findBookCommand:  h.FindBookCommand,
		psychology:       h.FindBookByGenreCommand,
		history:          h.FindBookByGenreCommand,
		prose:            h.FindBookByGenreCommand,
		poetry:           h.FindBookByGenreCommand,
		horrors:          h.FindMovieByGenreCommand,
		adventure:        h.FindMovieByGenreCommand,
		detectives:       h.FindMovieByGenreCommand,
		fantastic:        h.FindMovieByGenreCommand,
	}

	return h
}

type sendWebhookURL struct {
	URL string `json:"url"`
}

func (h *Handler) SetWebhook() error {
	webhookURL := config.C.App.WebhookURL
	body := &sendWebhookURL{
		URL: webhookURL,
	}
	reqBytes, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("could not marshal body: %v\n", err)
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/setWebhook", h.botToken)
	_, err = http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return fmt.Errorf("could not set webhook: %v\n", err)
	}

	return nil
}

func (h *Handler) StartCommand(update *Update) error {
	body := &ResponseBody{
		ChatID: update.Message.Chat.ID,
		Text:   "Добро пожаловать! Я помогу Вам выбрать занятие на вечер. Для этого введите команду /find!",
	}

	reqBytes, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error at start command in marshal response: %v\n", err)
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", h.botToken)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return fmt.Errorf("error at start command in http request: %v\n", err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error at start command unexpected http status: %d\n", res.StatusCode)
	}

	return nil
}

func (h *Handler) UnknownCommand(update *Update) error {
	body := &ResponseBody{
		ChatID: update.Message.Chat.ID,
		Text:   "Я Вас не понимаю",
	}

	reqBytes, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error at unknown command in marshal response: %v\n", err)
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", h.botToken)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return fmt.Errorf("error at unknown command in http request: %v\n", err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error at unknown command unexpected http status: %d\n", res.StatusCode)
	}

	return nil
}

func (h *Handler) FindCommand(update *Update) error {
	body := &ResponseBody{
		ChatID: update.Message.Chat.ID,
		Text:   "Выберите, чем вы хотите занять себя сейчас!",
		ReplyMarkup: &KeyboardMarkup{
			Keyboard: [][]KeyboardButton{
				{
					{
						Text: "Книга",
					},
					{
						Text: "Фильм",
					},
				},
			},
		},
	}

	reqBytes, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error at find command in marshal response: %v\n", err)
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", h.botToken)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return fmt.Errorf("error at find command in http request: %v\n", err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error at find command unexpected http status: %d\n", res.StatusCode)
	}

	return nil
}

func (h *Handler) FindBookCommand(update *Update) error {
	body := &ResponseBody{
		ChatID: update.Message.Chat.ID,
		Text:   "Выберите жанр книги!",
		ReplyMarkup: &KeyboardMarkup{
			Keyboard: [][]KeyboardButton{
				{
					{
						Text: "Психология",
					},
					{
						Text: "История",
					},
					{
						Text: "Проза",
					},
					{
						Text: "Поэзия",
					},
				},
			},
		},
	}

	reqBytes, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error at find book command in marshal response: %v\n", err)
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", h.botToken)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return fmt.Errorf("error at find book command in http request: %v\n", err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error at find book command unexpected http status: %d\n", res.StatusCode)
	}

	return nil
}

func (h *Handler) FindMovieCommand(update *Update) error {
	body := &ResponseBody{
		ChatID: update.Message.Chat.ID,
		Text:   "Выберите жанр фильма!",
		ReplyMarkup: &KeyboardMarkup{
			Keyboard: [][]KeyboardButton{
				{
					{
						Text: "Ужасы",
					},
					{
						Text: "Приключение",
					},
					{
						Text: "Детективы",
					},
					{
						Text: "Фантастика",
					},
				},
			},
		},
	}

	reqBytes, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error at find book command in marshal response: %v\n", err)
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", h.botToken)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return fmt.Errorf("error at find book command in http request: %v\n", err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error at find book command unexpected http status: %d\n", res.StatusCode)
	}

	return nil
}

func (h *Handler) FindMovieByGenreCommand(update *Update) error {
	movie := ""
	switch update.Message.Text {
	case horrors:
		movie = "Астрал"
	case adventure:
		movie = "Чупа"
	case detectives:
		movie = "Убийство в Париже"
	case fantastic:
		movie = "Самаритянин"
	}

	body := &ResponseBody{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("Посмотрите %s", movie),
	}

	reqBytes, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error at find book command in marshal response: %v\n", err)
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", h.botToken)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return fmt.Errorf("error at find book command in http request: %v\n", err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error at find book command unexpected http status: %d\n", res.StatusCode)
	}

	return nil
}

func (h *Handler) FindBookByGenreCommand(update *Update) error {
	book := ""
	switch update.Message.Text {
	case psychology:
		book = "З. Фрейд. Введение в психоанализ"
	case history:
		book = "А. В. Яценко. Из варяг в греки"
	case prose:
		book = "Х. Мулиш. Расплата"
	case poetry:
		book = "М. Винтерс. Леонардо да Винчи"
	}

	body := &ResponseBody{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("Почитайте %s", book),
	}

	reqBytes, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error at find book command in marshal response: %v\n", err)
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", h.botToken)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return fmt.Errorf("error at find book command in http request: %v\n", err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error at find book command unexpected http status: %d\n", res.StatusCode)
	}

	return nil
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	body := &Update{}
	err := json.NewDecoder(req.Body).Decode(body)
	if err != nil {
		log.Printf("could not decode request body: %v\n", err)
		return
	}

	handler, ok := h.commandHandlers[body.Message.Text]
	if !ok {
		err = h.UnknownCommand(body)
		if err != nil {
			log.Printf("could not execute command: %v\n", err)
		}
		return
	}

	err = handler(body)
	if err != nil {
		log.Printf("could not execute command: %v\n", err)
	}
}
