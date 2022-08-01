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

func loadingScreen() {
	fmt.Printf("\n\n\n\n\n\n\n\n")
	fmt.Println("\tLucas POS System :)")
	fmt.Println("-----------------------------------")
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
	fmt.Printf("Username: %s\n", name)
	fmt.Printf("Clocked in at: %s\n", dt.Format(time.RFC822))
	println("hours worked xx/xx-xx/xx : TBA")

}

// func clockOutPrintout(name string) {
// 	dt := time.Now()
// 	fmt.Printf("Username: %s\n", name)
// 	fmt.Printf("Clocked in at: %s\n", dt.Format(time.RFC822))
// 	println("hours worked xx/xx-xx/xx : TBA")
// }

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
				logClockIn(record[1] + "," + record[2])
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

func logClockIn(e string) {
	dt := time.Now()
	todayDt := dt.Format("2006-01-02")
	clockInTime := dt.Format("15:04:05")
	todayDt = todayDt + "-log.csv"

	outFile, err := os.OpenFile(todayDt, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := outFile.Write([]byte(e + "," + clockInTime + ",Clocked IN" + "\n")); err != nil {
		outFile.Close()
		log.Fatal(err)
	}
	if err := outFile.Close(); err != nil {
		log.Fatal(err)
	}
}

func isClockedIn(emloyeePin string) bool {
	dt := time.Now()
	todayDt := dt.Format("2006-01-02")
	todayDt = todayDt + "-log.csv"
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
