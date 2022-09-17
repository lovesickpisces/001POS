package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	tools "timekeeping/tools"
)

// Printouts
func loadingScreen() {
	fmt.Printf("\n\n\n\n\n\n\n\n")
	fmt.Println("\tLucas POS System :)")
	fmt.Println("----------------------------------------")
	fmt.Printf("\n\n\n\n\n\n")
}
func clockInPrintout(employee tools.EmployeeShift) {
	fmt.Printf("\n\n\n")
	fmt.Println("\tEmployee Clock In")
	fmt.Println("----------------------------------------")
	fmt.Printf("Username: %s\t%s\n", employee.Username, employee.ClockIn.Format("January 2, 2006"))
	fmt.Printf("\nClocked in at: \t\t%s\n", employee.ClockIn.Format("3:04:05 PM"))
	fmt.Printf("Job: \t\t\t%s\n", employee.Job)
	fmt.Printf("\n")
	fmt.Println("----------------------------------------")
	fmt.Printf("\nPress 'Enter/Return' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
func clockOutPrintout(employee tools.EmployeeShift) {
	fmt.Printf("\n\n\n")
	fmt.Println("\tEmployee Clock out")
	fmt.Println("----------------------------------------")
	fmt.Printf("Username: %s\t%s\n", employee.Username, employee.ClockIn.Format("January 2, 2006"))
	fmt.Printf("\nClocked out at: \t\t%s\n", employee.ClockIn.Format("3:04:05 PM"))
	fmt.Printf("Job: \t\t\t%s\n", employee.Job)
	fmt.Printf("\n")
	fmt.Println("----------------------------------------")
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
		//clockedInCheck := false
		if strings.Compare(employeePin, record[1]) == 0 {
			if isClockedIn(employeePin) {
				input := tools.GetFeedback("do you want to clock out "+record[2]+" (y/n)", "type (b) to go back")
				tools.LeaveCheck(input)
				if strings.Compare(input, "y") == 0 {
					employee := tools.CreateEmployeeShift(record[0], record[1], record[2], "job")
					tools.LogClockOut(employee)
					clockOutPrintout(employee)

					//getClockOut(employeePin)
				} else if strings.Compare(input, "n") == 0 {
					fmt.Println("Okay :)")
					break
				}

			} else {
				input := tools.GetFeedback("do you want to clock in "+record[2]+" (y/n)", "type (b) to go back")
				tools.LeaveCheck(input)
				if strings.Compare(input, "y") == 0 {
					employeeClockInJob(employeePin, record[0], record[2])
				} else if strings.Compare(input, "n") == 0 {
					fmt.Println("Okay :)")
					break
				}
			}
		}
	}
}
func employeeClockInJob(employeePin, name, username string) {
	count := 0
	jobsList := tools.GetEmployeeJobs(employeePin)
	jobInput := tools.GetFeedback("Please type the number of the job you are working today", "Type (b) to go back")
	if !tools.LeaveCheck(jobInput) {
		return
	}
	jobInputInt, err := strconv.Atoi(jobInput)
	if err != nil {
		printBadInput("\tThis job is not listed.\n")
		return
	}
	thisEmployeeJob := ""
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
	if thisEmployeeJob == "" {
		printBadInput("\tThis job is not listed.\n")
		return
	}

	employee := tools.CreateEmployeeShift(name, employeePin, username, thisEmployeeJob)
	tools.LogClockIn(employee)
	clockInPrintout(employee)
}

// questionable...
func isClockedIn(employeePin string) bool {
	var clockInStatus bool
	filename := tools.GetDailyLog()
	infile, err := os.Open(filename)
	if err != nil {
		return false
	}
	logCSV := csv.NewReader(infile)
	for {
		record, err := logCSV.Read()
		if err == io.EOF {
			break
		}
		if strings.Compare(employeePin, record[0]) == 0 {
			if strings.Compare(record[3], "Clocked IN") == 0 {
				clockInStatus = true
			} else if strings.Compare(record[3], "Clocked OUT") == 0 {
				clockInStatus = false
			}
		}
	}
	return clockInStatus
}

// Todo list:
/*
func hash(key string) string {
}
*/
/*
func createEmployee() Employee {

}
*/
/*
func fetchEmployee() Employee {

}
*/
/*
func storeEmployee(e Employee) {

}
*/
/*
func updateEmployee(e Employee) {

}
*/
/*
getDailyLog() DailyLog {

}
*/

func main() {
	//clockedInEmployees := make(map[string]Employee)
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
