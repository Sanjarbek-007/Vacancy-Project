package login

import (
	"encoding/json"
	"fmt"
	"os"
)

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func LoginFunc() int {
	adminLogin := "golang"
	adminPassword := "1111"

	var adminORuser = 0

	var users []Login
	ReadJson(&users)

	fmt.Print("1. LogIn\n2. SignIn\n")
	var buyruq int
	fmt.Scan(&buyruq)

	if buyruq == 1 {
		LogIn(&adminORuser, adminLogin, adminPassword, users)
	} else if buyruq == 2 {
		SignIn(&users, &adminORuser)
	} else {
		fmt.Println("Xato buyruq kiritdingiz")
		panic(0)
	}

	return adminORuser
}

func ReadJson(users *[]Login) {
	file, err := os.OpenFile("login/logins.json", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	user := make([]byte, stat.Size())

	_, err = file.Read(user)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(user, &users)
	if err != nil {
		panic(err)
	}
}

func WriteJson(LogPas Login, users *[]Login, AorU *int) bool {
	ReadJson(users)
	for _, i := range *users {
		if LogPas.Login == i.Login {
			fmt.Println("Bunday logenli foydalanuvchi mavjud.")
			return false
		}
	}
	*users = append(*users, LogPas)
	file, err := os.Create("login/logins.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Muvaffaqiyatli ro'yxatdan o'tdingiz.")
	*AorU = 2
	return true
}

func LogIn(AorU *int, adminLogin, adminPassword string, users []Login) {
	var login, password string

	for i := 0; i < 3; i++ {
		fmt.Print("Login: ")
		fmt.Scan(&login)
		fmt.Print("Password: ")
		fmt.Scan(&password)

		if login == adminLogin && password == adminPassword {
			*AorU = 1 // admin
			break
		}

		for _, j := range users {
			if j.Login == login && j.Password == password {
				*AorU = 2 // user
				return
			}
		}

		if i == 2 {
			fmt.Print("Ko'p marta xato urunish.")
			return
		}
		fmt.Println("Login yoki password xato")
	}

}

func SignIn(users *[]Login, AorU *int) {
	var LogPas Login

	fmt.Print("Login: ")
	fmt.Scan(&LogPas.Login)
	fmt.Print("Password: ")
	fmt.Scan(&LogPas.Password)

	for i := 0; i < 4; i++ {
		if WriteJson(LogPas, users, AorU) {
			break
		} else {
			if i == 3 {
				fmt.Print("Ko'p marta xato urunish.")
				return
			}
			fmt.Print("Login: ")
			fmt.Scan(&LogPas.Login)
			fmt.Print("Password: ")
			fmt.Scan(&LogPas.Password)
		}
	}
}
