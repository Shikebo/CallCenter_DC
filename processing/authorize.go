package processing

import "fmt"

func Authorize(id int64, user_id int64) bool {
	if id == user_id {
		return true // Пользователь авторизован, если id и user_id совпадают
	} else {
		fmt.Println("Это не ваш номер, пожалуйста, введите свой номер")
		return false // Пользователь не авторизован, если id и user_id не совпадают
	}
}
