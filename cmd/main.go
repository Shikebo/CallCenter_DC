package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"master/interface/Keyboards"
	"master/interface/constants"
	"master/processing"
)

type UserState struct {
	CurrentMenu  string
	PreviousMenu string // Добавляем поле для хранения предыдущего меню
}

var userStates map[int64]*UserState

func main() {
	userStates = make(map[int64]*UserState)
	bot, err := tgbotapi.NewBotAPI(processing.GetToken())
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Sharifova Shikebo: %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		HandleMessage(bot, update)
	}
}
func HandleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r)
			errorMsg := "Эта команда не обработана. Нажмите (Назад) или выберите другие команды."
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, errorMsg)
			_, err := bot.Send(msg)
			if err != nil {
				log.Println("Error sending message:", err)
			}
		}
	}()

	if update.Message == nil {
		return
	}

	username := update.Message.From.UserName
	userID := update.Message.From.ID
	chatID := update.Message.Chat.ID

	log.Printf("UserName: @%s, User_ID: %d, Chat_ID: %d, Message: %s", username, userID, chatID, update.Message.Text)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	firstName := update.Message.From.FirstName
	lastName := update.Message.From.LastName

	if update.Message.Contact != nil {
		contact := update.Message.Contact
		PhoneNumber := contact.PhoneNumber
		log.Printf("Phone Number: %s", PhoneNumber)
	}

	if update.Message.Text != "" {
		msg.Text = update.Message.Text
	}

	if _, ok := userStates[update.Message.Chat.ID]; !ok {
		userStates[update.Message.Chat.ID] = &UserState{}
	}
	userState := userStates[update.Message.Chat.ID]

	switch update.Message.Text {
	case "/start":
		msg.ReplyMarkup = Keyboards.ContactKeyboard
		msg.Text = fmt.Sprintf(constants.WelcomeMessage, lastName, firstName)
	case "\xF0\x9F\x93\x9E Авторизоваться с номером ":
		msg.ReplyMarkup = Keyboards.StartMenu
		msg.Text = fmt.Sprintf(constants.WelcomeMessage, lastName, firstName)
	case "Ошибочное пополнение":
		userState.PreviousMenu = userState.CurrentMenu
		userState.CurrentMenu = "Ошибочное пополнение"
		msg.ReplyMarkup = Keyboards.Payment
	case "Правильный номер":
		userState.PreviousMenu = userState.CurrentMenu
		userState.CurrentMenu = "Правильный номер"
		msg.ReplyMarkup = Keyboards.CorrectNumber
	case "Номер терминала":
		userState.PreviousMenu = userState.CurrentMenu
		userState.CurrentMenu = "Номер терминала"
		msg.ReplyMarkup = Keyboards.TerminalNum
	case "Карты":
		userState.PreviousMenu = userState.CurrentMenu
		userState.CurrentMenu = "Карты"
		msg.ReplyMarkup = Keyboards.Maps
	case "Кошелек (DC_NEXT)":
		userState.PreviousMenu = userState.CurrentMenu
		userState.CurrentMenu = "КошелекDC_NEXT"
		msg.ReplyMarkup = Keyboards.WalletDC
	case "Кредиты":
		userState.PreviousMenu = userState.CurrentMenu
		userState.CurrentMenu = "Кредиты"
		msg.ReplyMarkup = Keyboards.Loans
	case "Терминалы":
		userState.PreviousMenu = userState.CurrentMenu
		userState.CurrentMenu = "Терминалы"
		msg.ReplyMarkup = Keyboards.Terminals
	case "Заказ карты":
		userState.PreviousMenu = userState.CurrentMenu
		userState.CurrentMenu = "Заказ карты"
		msg.ReplyMarkup = Keyboards.Ordercard
	case "Идентификация":
		userState.PreviousMenu = userState.CurrentMenu
		userState.CurrentMenu = "Идентификация"
		msg.ReplyMarkup = Keyboards.Identification
	case "Пройти идентификацию":
		userState.PreviousMenu = userState.CurrentMenu
		userState.CurrentMenu = "Пройти идентификацию"
		msg.ReplyMarkup = Keyboards.PassIden
	case "Переводы":
		userState.PreviousMenu = userState.CurrentMenu
		userState.CurrentMenu = "Переводы"
		msg.ReplyMarkup = Keyboards.Transfers
	case "Назад":
		userState, ok := userStates[update.Message.Chat.ID]
		if ok {
			if userState.PreviousMenu != "" {
				log.Printf("Previous menu found: %s", userState.PreviousMenu)
				msg.ReplyMarkup = processing.BackToPreviousMenu(update.Message.Chat.ID, userState.PreviousMenu)
				userState.CurrentMenu = userState.PreviousMenu
				userState.PreviousMenu = ""
			} else {
				log.Printf("No previous menu found for chat %d", update.Message.Chat.ID)
				msg.ReplyMarkup = Keyboards.StartMenu
				userState.CurrentMenu = ""
			}
		} else {
			log.Printf("No user state found for chat %d", update.Message.Chat.ID)
			msg.ReplyMarkup = Keyboards.StartMenu
		}
	default:
		log.Println("from: ", update.Message.From.ID, " Contact: ", update.Message.Contact.UserID)

		user_id := update.Message.From.ID
		contact_id := update.Message.Contact.UserID

		resp := "Введите свой номер, это не ваш номер "
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, resp)

		if processing.Authorize(contact_id, user_id) {
			resp = "Вы авторизованы"
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, resp)
			msg.ReplyMarkup = Keyboards.StartMenu
		}
		fmt.Println(resp)
	}

	_, err := bot.Send(msg)
	if err != nil {
		log.Println("Error sending message:", err)
	}
}
