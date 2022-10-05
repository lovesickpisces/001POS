package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	tools "timekeeping/tools"
)

//global map to store employees and log files:)
var clockedInEmployees = make(map[string]tools.EmployeeShift)
var dailyLogger = make(map[string]tools.DailyLog)

// Printouts
func loadingScreen() {
	fmt.Printf("\n\n\n\n\n\n\n\n")
	fmt.Println("\tLucas POS System :)")
	fmt.Println("-------------------------------------------")
	fmt.Printf("\n\n\n\n\n\n")
}
func clockInPrintout(employee tools.EmployeeShift) {
	fmt.Printf("\n\n\n")
	fmt.Println("\tEmployee Clock In")
	fmt.Println("-------------------------------------------")
	fmt.Printf("Username: \t\t%s\n", employee.Username)
	fmt.Printf("Date: \t\t\t%s\n", employee.ClockIn.Format("January 2, 2006"))
	fmt.Printf("Clocked in at: \t\t%s\n", employee.ClockIn.Format("3:04:05 PM"))
	fmt.Printf("Job: \t\t\t%s\n", employee.Job)
	fmt.Println("-------------------------------------------")
	fmt.Printf("Press 'Enter/Return' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
func clockOutPrintout(employee tools.EmployeeShift) {
	currTime := time.Now()
	fmt.Printf("\n\n\n")
	fmt.Println("\tEmployee Clock out")
	fmt.Println("-------------------------------------------")
	fmt.Printf("Username: \t\t%s\n", employee.Username)
	fmt.Printf("Date: \t\t\t%s\n", employee.ClockIn.Format("January 2, 2006"))
	fmt.Printf("Clock in: \t\t%s\n", employee.ClockIn.Format("3:04:05 PM"))
	fmt.Printf("Clock out: \t\t%s\n", currTime.Format("3:04:05 PM"))
	diff := currTime.Sub(employee.ClockIn)
	shiftTime := time.Time{}.Add(diff)
	fmt.Println("Shift Hours:\t\t", shiftTime.Format("15:04:05"))
	fmt.Printf("Job: \t\t\t%s\n", employee.Job)
	fmt.Println("-------------------------------------------")
	fmt.Printf("\nPress 'Enter/Return' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
func printBadInput(prompt string) {
	fmt.Printf("\n\n\n----------------------------------------\n")
	fmt.Printf("%s", prompt)
	fmt.Printf("----------------------------------------\n\n\n")
	fmt.Printf("\nPress 'Enter/Return' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
func printMaps() {
	fmt.Println("-------------------------------------------")
	fmt.Println("clockedInEmployees:", clockedInEmployees)
	fmt.Println()
	fmt.Println("dailyLogger:", dailyLogger)
	fmt.Println()
	// fmt.Println("employeeTracker:", employeeTracker)
	// fmt.Println()
	fmt.Println("-------------------------------------------")
	fmt.Printf("\nPress 'Enter/Return' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
func logInMenu(employee tools.EmployeeShift) {
	fmt.Printf("\n\n\n\n\n\n\n\n")
	fmt.Println("\tWelcome", employee.Username, " :)")
	fmt.Println("-------------------------------------------")
	fmt.Println("(1) Clock Out")
	switch employee.Job {
	case "Manager":
		fmt.Println("(2) Add New Employee")
		fmt.Println("(3) Delete Employee")
		fmt.Println("(4) Clear Log File")
		fmt.Println("(5) Print Maps")
	case "Server":
	case "Bartender":
	case "Kitchen":
	case "Salary Manager":
		fmt.Println("(2) Add New Employee")
		fmt.Println("(3) Delete Employee")
		fmt.Println("(4) Clear Log File")
	case "Salary Kitchen":
		fmt.Println("(2) Add New Employee")
		fmt.Println("(3) Delete Employee")
		fmt.Println("(4) Clear Log File")
	case "FOH Support":
	}
	fmt.Printf("\n(b) To go Back")

}

//Clock in/out stuff
func clockInEmployee(employeePin string) {
	infile, err := os.Open("../docs/Employee.csv")
	if err != nil {
		log.Fatal(err)
	}
	employeeCSV := csv.NewReader(infile)
	for {
		record, err := employeeCSV.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if strings.Compare(employeePin, record[1]) == 0 {
			input := tools.GetFeedback("(b) to go back", "\n\n\ndo you want to clock in "+record[2]+" (y/n)")
			if !tools.LeaveCheck(input) {
				break
			}
			if strings.Compare(input, "y") == 0 {
				employeejob, noJobError := employeeClockInJob(employeePin)
				if noJobError == nil {
					employee := tools.CreateEmployeeShift(record[0], employeePin, record[2], employeejob)
					dl := tools.CreateDailyLog(employee)
					storeDailyLog(dl)
					storeEmployeeShift(employee)
					clockInPrintout(employee)
				} else {
					break
				}
			} else if strings.Compare(input, "n") == 0 {
				fmt.Println("Okay :)")
				break
			}
		}
	}
}
func clockOutEmployee(clockedInEmployee tools.EmployeeShift) {
	input := tools.GetFeedback("(b) to go back", "\n\n\ndo you want to clock out "+clockedInEmployee.Username+" (y/n)")
	if !tools.LeaveCheck(input) {
		return
	}
	if strings.Compare(input, "y") == 0 {
		dl, err := fetchDailyLog(clockedInEmployee.Pin)
		if err == nil {
			tools.LogClockOut(dl)
			removeDailyLog(clockedInEmployee.Pin)
		} else {
			fmt.Println("Unable to find employee!")
			return
		}

		clockOutPrintout(clockedInEmployee)
		removeEmployeeShift(clockedInEmployee.Pin)
	} else if strings.Compare(input, "n") == 0 {
		printBadInput("\t\tOkay :)\n")
		return
	}
}
func employeeClockInJob(employeePin string) (jobReturn string, err error) {
	count := 0
	jobsList := tools.GetEmployeeJobs(employeePin)
	jobInput := tools.GetFeedback("(b) to go back", "Please type the number of the job you are working today")
	if !tools.LeaveCheck(jobInput) {
		return "", errors.New("left function to go back")
	}
	thisEmployeeJob := ""
	jobInputInt, err := strconv.Atoi(jobInput)
	if err != nil {
		printBadInput("\tThis job is not listed.\n")
	} else {
		if strings.Compare(jobInput, "") != 0 {
			for i := 0; i < 7; i++ {
				if strings.Compare(jobsList[i], "n") == 0 {
					i++
				} else {
					count++
					if count == jobInputInt {
						thisEmployeeJob = jobsList[i]
					}
				}
			}
		}
	}
	if thisEmployeeJob == "" {
		printBadInput("\tThis job is not listed.\n")
		return "", errors.New("picked a job that doesnt exist for this employee")
	}

	return thisEmployeeJob, nil
}

//Shift Map stuff
func fetchEmployeeShift(key string) (eReturn tools.EmployeeShift, err error) {
	var e tools.EmployeeShift
	e, prs := clockedInEmployees[key]
	if prs {
		return e, nil
	} else {
		return e, errors.New("unable to find in map with given key")
	}
}
func storeEmployeeShift(e tools.EmployeeShift) {
	clockedInEmployees[e.Pin] = e
}
func removeEmployeeShift(pin string) {
	delete(clockedInEmployees, pin)
}

//Daily Log stuff
func fetchDailyLog(key string) (DLReturn tools.DailyLog, err error) {
	var dl tools.DailyLog
	dl, prs := dailyLogger[key]
	if prs {
		return dl, nil
	} else {
		return dl, errors.New("unable to find a DailyLog in map with given key")
	}
}
func storeDailyLog(dl tools.DailyLog) {
	dailyLogger[dl.Pin] = dl
}
func removeDailyLog(pin string) {
	delete(dailyLogger, pin)
}

func createEmployee() {
	fmt.Printf("\n\n\n\n\n\n\n\n")
	fmt.Println("\tLucas POS System :)")
	fmt.Println("-------------------------------------------")
	name := tools.GetFeedback("Please enter your new employee's full name (first last):", "")
	prompt := "Please enter " + name + "'s"
	fmt.Printf("\n\n\n\n\n\n\n\n")
	fmt.Println("\tLucas POS System :)")
	fmt.Println("-------------------------------------------")
	username := tools.GetFeedback(prompt+" username:", "")
	fmt.Printf("\n\n\n\n\n\n\n\n")
	fmt.Println("\tLucas POS System :)")
	fmt.Println("-------------------------------------------")
	pin := tools.GetFeedback(prompt+" pin:", "")
	fmt.Printf("\n\n\n\n\n\n\n\n")
	fmt.Println("\tLucas POS System :)")
	fmt.Println("-------------------------------------------")
	phone := tools.GetFeedback(prompt+" phone number:", "")
	filename := "../docs/Employee.csv"
	outFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	employeeInfo := "\n" + name + "," + pin + "," + username + "," + phone
	if _, err := outFile.Write([]byte(employeeInfo)); err != nil {
		outFile.Close()
		log.Fatal(err)
	}
	if err := outFile.Close(); err != nil {
		log.Fatal(err)
	}
}

func logIn(key string) {
	employee, err := fetchEmployeeShift(key)
	if err == nil {

		logInMenu(employee)
		input := tools.GetFeedback("", "")
		if !tools.LeaveCheck(input) {
			return
		}
		switch input {
		case "1":
			clockOutEmployee(employee)
		case "2":
			//add employee
			createEmployee()
		case "3":
			//delete employee
		case "4":
			tools.ClearLog()
		case "5":
			printMaps()
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		case "q":
			os.Exit(0)
		case "b":
			break
		}

	} else {
		clockInEmployee(key)
	}
}
func main() {
	for {
		loadingScreen()
		input := tools.GetFeedback("Enter an employee pin", "")
		tools.LeaveCheck(input)
		if strings.Compare(input, "") != 0 {
			logIn(input)
		}
	}
}

/*
 *	Helpful Notes:
 *		- Comments in main code are 14 tabs in.
 *		- More notes here
 *
 *	Helpful Links:
 *	Reading Console input:
 *		https://tutorialedge.net/golang/reading-console-input-golang/
 *	Structs:
 *		https://gobyexample.com/structs
 *	Reading CSV files:
 *		https://zetcode.com/golang/csv/
 *	Date/Time:
 *		https://tecadmin.net/get-current-date-time-golang/
 *
 *
 *
 *
 */
