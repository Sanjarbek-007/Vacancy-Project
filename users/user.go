package user

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name     string
	MyResume Resume
	Field    string
}

type Resume struct {
	Experence int
	Skils     []string
	About     string
	Level     string
}

func UserMain(AorU int) {
	if AorU == 2 {
		var users []User
		var buyruq int
		fmt.Print("1. AddResume\n2. Back\n")
		fmt.Scan(&buyruq)
		switch buyruq {
		case 1:
			AddResume(&users)
		default:
			return
		}
	} else if AorU == 1 {
		var users []User
		var buyruq int
		fmt.Print("1. Resumes\n2. Back\n")
		fmt.Scan(&buyruq)
		switch buyruq {
		case 1:
			UsersInfo(&users)
		default:
			return
		}
	}
}

func AddResume(users *[]User) {
	var user User
	fmt.Print("Name: ")
	fmt.Scan(&user.Name)
	fmt.Print("Field: ")
	fmt.Scan(&user.Field)
	fmt.Print("Experence: ")
	fmt.Scan(&user.MyResume.Experence)
	fmt.Print("About: ")
	fmt.Scan(&user.MyResume.About)
	fmt.Print("Level: ")
	fmt.Scan(&user.MyResume.Level)
	fmt.Println("Eng muhim uchta skillingizni kiriting:")
	for i := 0; i < 3; i++ {
		var skil string
		fmt.Printf("%d-skill: ", i+1)
		fmt.Scan(&skil)
		user.MyResume.Skils = append(user.MyResume.Skils, skil)
	}

	ReadJson(users)
	*users = append(*users, user)

	file, err := os.Create("users/users.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := json.Marshal(*users)
	if err != nil {
		panic(err)
	}

	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}

}

func ReadJson(users *[]User) {
	file, err := os.OpenFile("users/users.json", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	data := make([]byte, stat.Size())

	_, err = file.Read(data)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(data, users)
}

func UsersInfo(users *[]User) {
	ReadJson(users)
	for j, i := range *users {
		fmt.Printf("%d. Name: %s\n   Field: %s\n   Experence: %d\n   About: %s\n   Level: %s\n   Skils: %+v\n\n", j+1, i.Name, i.Field, i.MyResume.Experence, i.MyResume.About, i.MyResume.Level, i.MyResume.Skils)
	}
}
