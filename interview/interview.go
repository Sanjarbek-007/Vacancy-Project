package interview

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Interview struct {
	Companies     []string
	InterviewDate string
	Recruiter     string
}

func (i *Interview) MakeAppointment(date string) {
	i.InterviewDate = date
}

func ReadJson() []Interview {
	f, err := os.Open("interview/interviews.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	var interviews []Interview

	err = json.Unmarshal(data, &interviews)
	if err != nil {
		panic(err)
	}
	return interviews
}

func WriteJson(rec, date, comName string) {
	oldInterviews := ReadJson()

	newInterview := Interview{
		Companies:     []string{comName},
		InterviewDate: date,
		Recruiter:     rec,
	}

	oldInterviews = append(oldInterviews, newInterview)

	file, err := os.OpenFile("interview/interviews.json", os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(oldInterviews)
	if err != nil {
		panic(err)
	}
}

func ChooseOption(AorU int) {
	if AorU == 1 {
		fmt.Println("Choose an option:")
		fmt.Println("1. view interwiev")
		fmt.Println("2. Write interwiev")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			Print()
		case 2:
			var rec, date, comName string
			fmt.Println("Enter recruiter name:")
			fmt.Scanln(&rec)
			fmt.Println("Enter interview date:")
			fmt.Scanln(&date)
			fmt.Println("Enter company name:")
			fmt.Scanln(&comName)
			WriteJson(rec, date, comName)
		default:
			fmt.Println("Invalid choice")
		}
	} else {
		Print()
	}
}

func Print() {
	interviews := ReadJson()
	for i, v := range interviews {
		fmt.Printf("Interview:      %d\n", i+1)
		fmt.Printf("Company:        %s\n", v.Companies[0]) // Access Companies through v
		fmt.Printf("Interview Date: %s\n", v.InterviewDate)
		fmt.Printf("Recruiter:      %s\n\n", v.Recruiter)
	}
}
