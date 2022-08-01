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

// type employee struct {
// 	name        string
// 	pin         string
// 	displayName string
// 	phoneNumber string
// 	clockIn     time.Time
// 	clockout    time.Time
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
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("do you want to clock in %s? (y/n)", record[2])
			fmt.Print("-> ")
			input, _ := reader.ReadString('\n')
			input = strings.Replace(input, "\n", "", -1)
			if strings.Compare(input, "y") == 0 {
				logClockIn(record[1] + "," + record[2])
				break
			} else if strings.Compare(input, "n") == 0 {
				fmt.Println("Okay :)")
				break
			} else {
				break
			}
		}
	}
}
func logClockIn(e string) {
	dt := time.Now()
	todayDt := dt.Format("2006-01-02")
	clockInTime := dt.Format("15:04:05")
	todayDt = todayDt + "-log.txt"

	outFile, err := os.OpenFile(todayDt, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := outFile.Write([]byte(e + "," + clockInTime + "\n")); err != nil {
		outFile.Close()
		log.Fatal(err)
	}
	if err := outFile.Close(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Lucas POS System :)")
	fmt.Println("--------------------")
	for {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		if strings.Compare(input, "") != 0 {
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
