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
	"time"
)

// Structures go here
type EmployeeShift struct {
	pin      string
	username string
	sales    float32
	ccTips   float32
	clockIn  time.Time
	job      string
}

//Pretty printouts and responses go here
func loadingScreen() {
	fmt.Printf("\n\n\n\n\n\n\n\n")
	fmt.Println("\tLucas POS System :)")
	fmt.Println("----------------------------------------")
	fmt.Printf("\n\n\n\n\n\n")
}
func leaveCheck(input string) bool {
	upperInput := strings.ToUpper(input)
	switch upperInput {
	case "Q":
		os.Exit(0)
	case "QUIT":
		os.Exit(0)
	case "E":
		os.Exit(0)
	case "EXIT":
		os.Exit(0)
	case "B":
		return false
	case "BACK":
		return false
	}
	return true
}
func clockInPrintout(employee EmployeeShift) {
	fmt.Printf("\n\n\n")
	fmt.Println("\tEmployee Clock In")
	fmt.Println("----------------------------------------")
	fmt.Printf("Username: %s\t%s\n", employee.username, employee.clockIn.Format("January 2, 2006"))
	fmt.Printf("\nClocked in at: \t\t%s\n", employee.clockIn.Format("3:04:05 PM"))
	fmt.Printf("Job: \t\t\t%s\n", employee.job)
	fmt.Printf("\n")
	fmt.Println("----------------------------------------")
	fmt.Printf("\nPress 'Enter/Return' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
func clockOutPrintout(employee EmployeeShift) {
	fmt.Printf("\n\n\n")
	fmt.Println("\tEmployee Clock out")
	fmt.Println("----------------------------------------")
	fmt.Printf("Username: %s\t%s\n", employee.username, employee.clockIn.Format("January 2, 2006"))
	fmt.Printf("\nClocked out at: \t\t%s\n", employee.clockIn.Format("3:04:05 PM"))
	fmt.Printf("Job: \t\t\t%s\n", employee.job)
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

// Tools go here (or the 'get' funtions section)
func getDailyLog() string {
	todayDt := time.Now().Format("2006-01-02")
	todayDt = "dailyLogs/" + todayDt + "-log.csv"
	return todayDt
}
func getFeedback(prompt, prompt2 string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nType (q) to quit this program.\n")
	if strings.Compare("", prompt2) == 0 {
		fmt.Printf("%s\n", prompt)
	} else {
		fmt.Printf("%s\n", prompt2)
		fmt.Printf("%s\n", prompt)
	}
	fmt.Print("-> ")
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	return input
}
func getEmployeeShift(newName, newPin, newUser, employeeJob string) EmployeeShift {
	dt := time.Now()
	employee := EmployeeShift{
		pin:      newPin,      // String
		username: newUser,     // String
		sales:    0,           // float32
		ccTips:   0,           // float32
		clockIn:  dt,          // time.Time
		job:      employeeJob, // String
	}
	return employee
}
func getEmployeeJobs(employeePin string) [7]string {
	var returnedJobs [7]string
	var jobs [7]string
	jobs[0] = "Manager"
	jobs[1] = "Server"
	jobs[2] = "Bartender"
	jobs[3] = "Kitchen"
	jobs[4] = "Salary Manager"
	jobs[5] = "Salary Kitchen"
	jobs[6] = "FOH Support"
	count := 0
	infile, err := os.Open("jobs.csv")
	if err != nil {
		log.Fatal(err)
	}
	jobsCSV := csv.NewReader(infile)
	fmt.Println("\n----------------------------------------")
	for {
		record, err := jobsCSV.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if strings.Compare(employeePin, record[0]) == 0 {
			for i := 2; i <= 8; i++ {
				if strings.Compare(record[i], "y") == 0 {
					fmt.Printf("(%d) %s \n", count+1, jobs[i-2])
					returnedJobs[count] = jobs[i-2]
					count++
				} else {
					returnedJobs[i-2] = "n"
				}
			}
			break
		}
	}
	fmt.Println("----------------------------------------")
	return returnedJobs
}

// functions for logging employee information
func logClockIn(employee EmployeeShift) {
	filename := getDailyLog()
	outFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	employeeInfo := employee.pin + "," + employee.username +
		"," + employee.clockIn.Format("15:04:05") + ",Clocked IN" + "," + employee.job + "\n"

	if _, err := outFile.Write([]byte(employeeInfo)); err != nil {
		outFile.Close()
		log.Fatal(err)
	}
	if err := outFile.Close(); err != nil {
		log.Fatal(err)
	}
}
func logClockOut(employee EmployeeShift) {
	filename := getDailyLog()
	outFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	employeeInfo := employee.pin + "," + employee.username +
		"," + employee.clockIn.Format("15:04:05") + ",Clocked Out" + "," + employee.job + "\n"

	if _, err := outFile.Write([]byte(employeeInfo)); err != nil {
		outFile.Close()
		log.Fatal(err)
	}
	if err := outFile.Close(); err != nil {
		log.Fatal(err)
	}
}

// 'Action' functions...
func clockInEmployee(employeePin string) {

	infile, err := os.Open("Employee.csv")
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
				input := getFeedback("do you want to clock out "+record[2]+" (y/n)", "type (b) to go back")
				leaveCheck(input)
				if strings.Compare(input, "y") == 0 {
					employee := getEmployeeShift(record[0], record[1], record[2], "job")
					logClockOut(employee)
					clockOutPrintout(employee)

					//getClockOut(employeePin)
				} else if strings.Compare(input, "n") == 0 {
					fmt.Println("Okay :)")
					break
				}

			} else {
				input := getFeedback("do you want to clock in "+record[2]+" (y/n)", "type (b) to go back")
				leaveCheck(input)
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
	jobsList := getEmployeeJobs(employeePin)
	jobInput := getFeedback("Please type the number of the job you are working today", "Type (b) to go back")
	if !leaveCheck(jobInput) {
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

	employee := getEmployeeShift(name, employeePin, username, thisEmployeeJob)
	logClockIn(employee)
	clockInPrintout(employee)
}

// questionable...
func isClockedIn(employeePin string) bool {
	var clockInStatus bool
	filename := getDailyLog()
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

func main() {
	for {
		loadingScreen()
		input := getFeedback("Enter an employee pin to clock in/out", "")
		leaveCheck(input)
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
