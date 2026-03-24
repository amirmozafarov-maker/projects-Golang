package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Visit struct {
	Date           string
	Specialization string
}

type Patient struct {
	Visits []Visit
}

type ClinicJournal struct {
	Records map[string]Patient
}

func NewClinicJournal() *ClinicJournal {
	return &ClinicJournal{
		Records: make(map[string]Patient),
	}
}

type PatientNotFoundError struct{}

func (e PatientNotFoundError) Error() string {
	return "patient not found"
}

// SaveOperation
func (journal *ClinicJournal) SaveOperation(scanner *bufio.Scanner) error {
	if !scanner.Scan() {
		return fmt.Errorf("Invalid input")
	}

	fullName := strings.TrimSpace(scanner.Text())

	if !scanner.Scan() {
		return fmt.Errorf("Invalid input")
	}

	specialization := strings.TrimSpace(scanner.Text())

	if !scanner.Scan() {
		return fmt.Errorf("Invalid input")
	}

	date := strings.TrimSpace(scanner.Text())

	if fullName == "" || specialization == "" || date == "" {
		return fmt.Errorf("Invalid input")
	}

	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return fmt.Errorf("Invalid input")
	}

	visit := Visit{
		Specialization: specialization,
		Date:           date,
	}

	patient, exists := journal.Records[fullName]
	if !exists {
		patient = Patient{
			Visits: make([]Visit, 0),
		}
	}
	patient.Visits = append(patient.Visits, visit)
	journal.Records[fullName] = patient
	return nil
}

// GetHistoryOperation
func (journal *ClinicJournal) GetHistoryOperation(scanner *bufio.Scanner) error {
	if !scanner.Scan() {
		return fmt.Errorf("Invalid input")
	}
	fullName := strings.TrimSpace(scanner.Text())

	if fullName == "" {
		return fmt.Errorf("Invalid input")
	}

	patient, exists := journal.Records[fullName]
	if !exists {
		return PatientNotFoundError{}
	}

	for _, visit := range patient.Visits {
		fmt.Printf("%s %s\n", visit.Specialization, visit.Date)
	}

	return nil
}

// GetLastVisitOperation
func (journal *ClinicJournal) GetLastVisitOperation(scanner *bufio.Scanner) error {
	if !scanner.Scan() {
		return fmt.Errorf("Invalid input")
	}
	fullName := strings.TrimSpace(scanner.Text())

	if !scanner.Scan() {
		return fmt.Errorf("Invalid input")
	}
	specialization := strings.TrimSpace(scanner.Text())

	if fullName == "" || specialization == "" {
		return fmt.Errorf("invalid input")
	}

	patient, exists := journal.Records[fullName]
	if !exists {
		return PatientNotFoundError{}
	}

	var lastDate string
	var lastDateTime time.Time

	for _, visit := range patient.Visits {
		if visit.Specialization == specialization {
			visitDate, _ := time.Parse("2006-01-02", visit.Date)

			if lastDate == "" || visitDate.After(lastDateTime) {
				lastDate = visit.Date
				lastDateTime = visitDate
			}
		}
	}

	if lastDate == "" {
		return nil
	}

	fmt.Println(lastDate)
	return nil
}

func handleError(err error) {
	if err == nil {
		return
	}

	switch err.(type) {
	case PatientNotFoundError:
		fmt.Println(err)
	default:
		fmt.Println("Invalid input")
	}
}

func main() {
	journal := NewClinicJournal()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Система 'Журнал посещений'")
	fmt.Println("Доступные команды: Save, GetHistory, GetLastVisit")
	fmt.Println("Для выхода нажмите Control+C")

	for {
		fmt.Print("\n> ")

		if !scanner.Scan() {
			break
		}

		command := strings.TrimSpace(scanner.Text())

		switch command {
		case "Save":
			err := journal.SaveOperation(scanner)
			handleError(err)

		case "GetHistory":
			err := journal.GetHistoryOperation(scanner)
			handleError(err)

		case "GetLastVisit":
			err := journal.GetLastVisitOperation(scanner)
			handleError(err)

		case "":
			continue

		default:
			fmt.Println("Invalid input")
		}
	}
}
