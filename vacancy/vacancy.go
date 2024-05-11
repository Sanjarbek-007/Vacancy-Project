package vacancy

import (
	"encoding/json"
	"fmt"
	"os"
)

type Vacancy struct {
	Id           int
	CompanyName  string
	Requirements string
	Salary       string
}

func VacancyMain(AorU int) {
	var vacancies []Vacancy
	var buyruq int

	if AorU == 1 {
		fmt.Print("1. Vacancies\n2. Add vacancy\n3. Delete vacancy\n")
		fmt.Scan(&buyruq)
		switch buyruq {
		case 1:
			Vacancies(vacancies)
		case 2:
			AddVacancy(&vacancies)
		case 3:
			DeleteVacancy(&vacancies)
		}
	} else {
		Vacancies(vacancies)
	}

}

func AddVacancy(vacancies *[]Vacancy) {
	ReadJson(vacancies)
	var vacancy Vacancy
	var Id int
	var CompanyName, Salary string
	var Requirments string
	fmt.Print("Vacansy Id: ")
	fmt.Scan(&Id)
	fmt.Print("Company name: ")
	fmt.Scan(&CompanyName)
	fmt.Print("Requirments: ")
	fmt.Scan(&Requirments)
	fmt.Print("Salary: ")
	fmt.Scan(&Salary)

	vacancy.Id = Id
	vacancy.CompanyName = CompanyName
	vacancy.Requirements = Requirments
	vacancy.Salary = Salary
	*vacancies = append(*vacancies, vacancy)

	file, err := os.Create("vacancy/vacancy.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := json.Marshal(*vacancies)
	if err != nil {
		panic(err)
	}

	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Vacancy addes succesfully.")

}

func DeleteVacancy(vacancies *[]Vacancy) {
	var id int
	fmt.Print("Id: ")
	fmt.Scan(&id)

	ReadJson(vacancies)

	var slc []Vacancy
	tek := false
	for _, i := range *vacancies {
		if i.Id != id {
			slc = append(slc, i)
		} else {
			tek = true
		}
	}
	*vacancies = slc

	file, err := os.Create("vacancy/vacancy.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := json.Marshal(*vacancies)
	if err != nil {
		panic(err)
	}

	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
	if tek {
		fmt.Println("Vakansy removed.")
	} else {
		fmt.Println("Not found.")
	}
}

func Vacancies(vacansies []Vacancy) {
	ReadJson(&vacansies)
	for _, i := range vacansies {
		fmt.Printf("%d. Company name: %s\n   Requirements: %s\n   Salary: %s\n\n", i.Id, i.CompanyName, i.Requirements, i.Salary)
	}
}

func ReadJson(vacancies *[]Vacancy) {
	file, err := os.OpenFile("vacancy/vacancy.json", os.O_RDONLY, 0666)
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

	err = json.Unmarshal(data, vacancies)
	if err != nil {
		panic(err)
	}

}

func SearchVacansies(compName string) {
	var vacansies []Vacancy
	ReadJson(&vacansies)
	for _, i := range vacansies {
		if i.CompanyName == compName {
			fmt.Printf("%d. Companyc name: %s\n   Requirements: %s\n   Salary: %s\n\n", i.Id, i.CompanyName, i.Requirements, i.Salary)
		}
	}
}
