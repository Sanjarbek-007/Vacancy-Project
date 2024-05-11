package company

import (
	"encoding/json"
	"fmt"
	"hh/vacancy"
	"os"
)

type Company struct {
	Name    string            `json:"name"`
	Rating  float64           `json:"rating"`
	Vacancy []vacancy.Vacancy `json:"vacancy"`
}

func Info() Company {
	var companies []Company
	data, err := os.ReadFile("company/companies.json")
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return Company{}
	}
	err = json.Unmarshal(data, &companies)
	if err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return Company{}
	}

	for i, n := range companies {
		fmt.Printf("%d - %s\n%s's Rating: %.2f\n", i+1, n.Name, n.Name, n.Rating)
		for _, v := range n.Vacancy {
			fmt.Printf("Vacancy ID: %d\nRequirements: %s\nSalary: %s\n\n", v.Id, v.Requirements, v.Salary)
		}
	}

	var idCompany int
	fmt.Println("Enter the ID of the company you want to select:")
	fmt.Scan(&idCompany)
	if idCompany < 1 || idCompany > len(companies) {
		fmt.Println("Invalid company ID")
		return Company{}
	}

	return companies[idCompany-1]
}

func PrintCompany() {
	var companies []Company
	data, err := os.ReadFile("company/companies.json")
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}
	err = json.Unmarshal(data, &companies)
	if err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return
	}

	for i, n := range companies {
		fmt.Printf("%d-%s\n  %s's Rating: %.2f\n", i+1, n.Name, n.Name, n.Rating)
		for _, v := range n.Vacancy {
			fmt.Printf("  Vacancy ID: %d\n  Requirements: %s\n  Salary: %s\n\n", v.Id, v.Requirements, v.Salary)
		}
	}
}

func WriteJsonCompanies(comp Company) {
	var companies []Company
	data, err := os.ReadFile("company/companies.json")
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}
	

	err = json.Unmarshal(data, &companies)
	if err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return
	}

	companies = append(companies, comp)

	data, err = json.Marshal(companies)
	if err != nil {
		fmt.Println("Failed to marshal JSON:", err)
		return
	}

	err = os.WriteFile("company/companies.json", data, 0666)
	if err != nil {
		fmt.Println("Failed to write file:", err)
	}
}

func CompanyMain(AorU int) {
	comp := Company{}
	var vacan []vacancy.Vacancy
	if AorU == 1 {
		fmt.Println("1. wiew companies\n2. Add company")
		var choice int
		fmt.Scan(&choice)
		if choice == 1 {
			PrintCompany()
		} else {
			fmt.Print("Company's rating: ")
			fmt.Scan(&comp.Rating)
			vacancy.AddVacancy(&vacan)
			comp.Name = vacan[len(vacan)-1].CompanyName
			comp.Vacancy = []vacancy.Vacancy{vacan[len(vacan)-1]}
			WriteJsonCompanies(comp)
		}
	} else if AorU == 2 {
		choice := Info()

		vacancy.SearchVacansies(choice.Name)
	}
}
