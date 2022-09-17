package tools

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type EmployeeShift struct {
	Pin      string
	Username string
	Sales    float32
	CcTips   float32
	ClockIn  time.Time
	Job      string
}
type Employee struct {
	Shift       EmployeeShift
	Phone       float32
	Jobs        [7]string
	PersonalLog DailyLog
}
type DailyLog struct {
	Pin      float32
	User     string
	ClockIn  time.Time
	ClockOut time.Time
	Job      string
}

// Tools
func LeaveCheck(input string) bool {
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

/* changed */
func GetDailyLog() string {
	todayDt := time.Now().Format("2006-01-02")
	todayDt = "../docs/dailyLogs/" + todayDt + "-log.csv"
	return todayDt
}
func GetFeedback(prompt1, prompt2 string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nType (q) to quit this program.\n")
	if strings.Compare("", prompt2) == 0 {
		fmt.Printf("%s\n", prompt1)
	} else {
		fmt.Printf("%s\n", prompt2)
		fmt.Printf("%s\n", prompt1)
	}
	fmt.Print("-> ")
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	return input
}
func CreateEmployeeShift(newName, newPin, newUser, employeeJob string) EmployeeShift {
	dt := time.Now()
	employee := EmployeeShift{
		Pin:      newPin,      // String
		Username: newUser,     // String
		Sales:    0,           // float32
		CcTips:   0,           // float32
		ClockIn:  dt,          // time.Time
		Job:      employeeJob, // String
	}
	return employee
}

/* changed */
func GetEmployeeJobs(employeePin string) [7]string {
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
	infile, err := os.Open("../docs/jobs.csv")
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
func LogClockIn(employee EmployeeShift) {
	filename := GetDailyLog()
	outFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	employeeInfo := employee.Pin + "," + employee.Username +
		"," + employee.ClockIn.Format("15:04:05") + ",Clocked IN" + "," + employee.Job + "\n"

	if _, err := outFile.Write([]byte(employeeInfo)); err != nil {
		outFile.Close()
		log.Fatal(err)
	}
	if err := outFile.Close(); err != nil {
		log.Fatal(err)
	}
}
func LogClockOut(employee EmployeeShift) {
	filename := GetDailyLog()
	outFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	employeeInfo := employee.Pin + "," + employee.Username +
		"," + employee.ClockIn.Format("15:04:05") + ",Clocked Out" + "," + employee.Job + "\n"

	if _, err := outFile.Write([]byte(employeeInfo)); err != nil {
		outFile.Close()
		log.Fatal(err)
	}
	if err := outFile.Close(); err != nil {
		log.Fatal(err)
	}
}
