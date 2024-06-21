package Keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"master/interface/constants"
)

var ContactKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButtonContact(constants.EMODJI_CALL + "Авторизоваться с номером"),
	),
)
var (
	StartMenu = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Карты")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Кошелек (DC_NEXT)"), tgbotapi.NewKeyboardButton("Кредиты")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Терминалы"), tgbotapi.NewKeyboardButton("Офисы")))

	Maps = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Заказ карты"), tgbotapi.NewKeyboardButton("Информация о картах")),
		tgbotapi.NewKeyboardButtonRow(BackButton))

	WalletDC = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Переводы"), tgbotapi.NewKeyboardButton("Бонусы")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Идентификация"), tgbotapi.NewKeyboardButton(constants.EMOJI_DAP+"Депозиты")),
		tgbotapi.NewKeyboardButtonRow(BackButton))

	Loans = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Дастрас"), tgbotapi.NewKeyboardButton("Автокредиты")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Ипотека"), tgbotapi.NewKeyboardButton("Потребительский кредит")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Ломбард"), tgbotapi.NewKeyboardButton("Связь с оператором")),
		tgbotapi.NewKeyboardButtonRow(BackButton))

	Terminals = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Ошибочное пополнение")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Связь с оператором")),
		tgbotapi.NewKeyboardButtonRow(BackButton))

	Payment = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Ошибочный номер"), tgbotapi.NewKeyboardButton("Паспортные данные")),
		tgbotapi.NewKeyboardButtonRow(BackButton))

	CorrectNumber = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Правильный номер")),
		tgbotapi.NewKeyboardButtonRow(BackButton))

	TerminalNum = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Номер терминала")),
		tgbotapi.NewKeyboardButtonRow(BackButton))

	Ordercard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Как заказать карту?")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Узнать статус заказа")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Связь с оператором")),
		tgbotapi.NewKeyboardButtonRow(BackButton))

	Identification = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Пройти идентификацию")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Информация о идентификации")),
		tgbotapi.NewKeyboardButtonRow(BackButton))

	PassIden = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Паспорт с двух сторон")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Селфи с паспортом")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(" Номер кошелька")),
		tgbotapi.NewKeyboardButtonRow(BackButton))

	Transfers = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("DC на Dc")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Переводы в РФ")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Переводы в Китай")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Переводы в Турцию")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Переводы в Дубай")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Переводы в Корея")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Переводы в Алмата")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Переводы в Сингапур")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Переводы в Франция")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Переводы в Англия")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Переводы в Германия")),
		tgbotapi.NewKeyboardButtonRow(BackButton))

	BackButton = tgbotapi.NewKeyboardButton("Назад")
)
