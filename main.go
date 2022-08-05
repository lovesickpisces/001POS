package main

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

type Employee struct {
	name     string
	pin      string
	username string
	sales    float32
	ccTips   float32
	clockIn  time.Time
	clockOut time.Time
	decTips  float32
	job      string
}

func loadingScreen() {
	fmt.Printf("\n\n\n\n\n\n\n\n")
	fmt.Println("\tLucas POS System :)")
	fmt.Println("----------------------------------------")
	fmt.Printf("\n\n\n\n\n\n\n\n")

}
func leaveCheck(input string) {
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
	}
}

func clockInPrintout(name string) {
	dt := time.Now()
	fmt.Printf("\n\n\n")
	fmt.Println("\tEmployee Clock In")
	fmt.Println("----------------------------------------")
	fmt.Printf("Username: %s\t%s\n", name, dt.Format("January 2, 2006"))
	fmt.Printf("\nClocked in at: \t\t%s\n", dt.Format("3:04:05 PM"))
	fmt.Printf("Job: \t\t\t%s\n", "insert job")
	fmt.Printf("\n")
	fmt.Println("----------------------------------------")
}

func findEmployee(employeePin string) {
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
		if strings.Compare(employeePin, record[1]) == 0 {
			//if isClockedIn(employeePin) {
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("do you want to clock in %s? (y/n)", record[2])
			fmt.Print("-> ")
			input, _ := reader.ReadString('\n')
			input = strings.Replace(input, "\n", "", -1)
			if strings.Compare(input, "y") == 0 {
				var e = logClockIn(record[0], record[1], record[2])
				storeEmployee(e)
				clockInPrintout(record[2])
				break
			} else if strings.Compare(input, "n") == 0 {
				fmt.Println("Okay :)")
				break
			} else {
				leaveCheck(input)
				break
			}
			//}
		}
		//  else {
		// 	fmt.Printf("do you want to clock out %s? (y/n)", record[2])
		// }
	}
}

func getDailyLog() string {
	todayDt := time.Now().Format("2006-01-02")
	todayDt = todayDt + "-log.csv"
	return todayDt
}
func logClockIn(newName, newPin, newUser string) Employee {
	dt := time.Now()
	todayDt := getDailyLog()
	employee := Employee{
		name:     newName,
		pin:      newPin,
		username: newUser,
		sales:    0,
		ccTips:   0,
		clockIn:  dt,
		decTips:  0,
		job:      "insert job",
	}
	outFile, err := os.OpenFile(todayDt, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := outFile.Write([]byte(employee.pin + "," + employee.username + "," + employee.clockIn.Format("15:04:05") + ",Clocked IN" + "\n")); err != nil {
		outFile.Close()
		log.Fatal(err)
	}
	if err := outFile.Close(); err != nil {
		log.Fatal(err)
	}
	return employee
}

func isClockedIn(emloyeePin string) bool {
	todayDt := getDailyLog()
	infile, err := os.Open(todayDt)
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
		if strings.Compare(emloyeePin, record[0]) == 0 {
			return false
		}

	}
	return true
}
func storeEmployee(e Employee) {

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	loadingScreen()
	for {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		if strings.Compare(input, "") != 0 {
			leaveCheck(input)
			findEmployee(input)
		}
	}
}

/*
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
