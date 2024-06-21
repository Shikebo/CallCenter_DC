package processing

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"master/interface/Keyboards"
)

func BackToPreviousMenu(chatID int64, previousMenu string) tgbotapi.ReplyKeyboardMarkup {
	switch previousMenu {
	case "Карты":
		return Keyboards.Maps
	case "КошелекDC_NEXT":
		return Keyboards.WalletDC
	case "Кредиты":
		return Keyboards.Loans
	case "Терминалы":
		return Keyboards.Terminals
	case "Идентификация":
		return Keyboards.Identification
	case "Пройти идентификацию":
		return Keyboards.Maps
	default:
		// В случае, если предыдущее меню не определено, можно вернуть стандартное меню (например, стартовое)
		return Keyboards.StartMenu
	}
}
