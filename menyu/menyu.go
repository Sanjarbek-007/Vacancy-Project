package menyu

import (
	"fmt"
	"hh/company"
	"hh/interview"
	user "hh/users"
	"hh/vacancy"
	"os"
	"os/exec"
	"runtime"
)

func Menyu(AorU int) {
	var choice int
	var buyruq int
	fmt.Print("1. Companies\n2. Vacancies\n3. Interviews\n4. Users\n5. Exit\n")
	fmt.Scan(&buyruq)
	switch buyruq {
	case 1:
		company.CompanyMain(AorU)
		fmt.Println("0. Back\n1. Exit")
		fmt.Scan(&choice)
		if choice == 1 {
			return
		}
	case 2:
		vacancy.VacancyMain(AorU)
		fmt.Println("0. Back\n1. Exit")
		fmt.Scan(&choice)
		if choice == 1 {
			return
		}
	case 3:
		interview.ChooseOption(AorU)
		fmt.Println("0. Back\n1. Exit")
		fmt.Scan(&choice)
		if choice == 1 {
			return
		}
	case 4:
		user.UserMain(AorU)
		fmt.Println("0. Back\n1. Exit")
		fmt.Scan(&choice)
		if choice == 1 {
			return
		}
	case 5:
		return
	}
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()

	Menyu(AorU)
}
