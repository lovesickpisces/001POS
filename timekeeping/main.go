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

//global map to store employees :)
var clockedInEmployees = make(map[string]tools.EmployeeShift)

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
			clockedInEmployee, clockedIn := fetchEmployeeShift(employeePin)
			if clockedIn == nil {
				input := tools.GetFeedback("do you want to clock out "+clockedInEmployee.Username+" (y/n)", "type (b) to go back")
				if !tools.LeaveCheck(input) {
					break
				}
				if strings.Compare(input, "y") == 0 {
					//employee := tools.CreateEmployeeShift(record[0], record[1], record[2], "job")
					tools.LogClockOut(clockedInEmployee)
					clockOutPrintout(clockedInEmployee)
					removeEmployeeShift(employeePin)
					//getClockOut(employeePin)
				} else if strings.Compare(input, "n") == 0 {
					fmt.Println("Okay :)")
					break
				}

			} else {
				input := tools.GetFeedback("do you want to clock in "+record[2]+" (y/n)", "type (b) to go back")
				if !tools.LeaveCheck(input) {
					break
				}
				if strings.Compare(input, "y") == 0 {
					employeejob, noJobError := employeeClockInJob(employeePin)
					if noJobError == nil {
						employee := tools.CreateEmployeeShift(record[0], employeePin, record[2], employeejob)
						tools.LogClockIn(employee)
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
}
func employeeClockInJob(employeePin string) (jobReturn string, err error) {
	count := 0
	jobsList := tools.GetEmployeeJobs(employeePin)
	jobInput := tools.GetFeedback("Please type the number of the job you are working today", "Type (b) to go back")
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

func main() {
	for {
		loadingScreen()
		input := tools.GetFeedback("Enter an employee pin to clock in/out", "")
		tools.LeaveCheck(input)
		if strings.Compare(input, "") != 0 {
			clockInEmployee(input)
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
