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

// Tools
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
func getDailyLog() string {
	todayDt := time.Now().Format("2006-01-02")
	todayDt = "dailyLogs/" + todayDt + "-log.csv"
	return todayDt
}
func getFeedback(prompt1, prompt2 string) string {
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
