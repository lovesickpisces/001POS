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

type EmployeeShift struct {
	pin      string
	username string
	sales    float32
	ccTips   float32
	clockIn  time.Time
	job      string
}

func loadingScreen() {
	fmt.Printf("\n\n\n\n\n\n\n\n")
	fmt.Println("\tLucas POS System :)")
	fmt.Println("----------------------------------------")
	fmt.Printf("\n\n\n\n\n\n")
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

func clockInPrintout(employee EmployeeShift) {
	fmt.Printf("\n\n\n")
	fmt.Println("\tEmployee Clock In")
	fmt.Println("----------------------------------------")
	fmt.Printf("Username: %s\t%s\n", employee.username, employee.clockIn.Format("January 2, 2006"))
	fmt.Printf("\nClocked in at: \t\t%s\n", employee.clockIn.Format("3:04:05 PM"))
	fmt.Printf("Job: \t\t\t%s\n", employee.job)
	fmt.Printf("\n")
	fmt.Println("----------------------------------------")
}

func findEmployee(employeePin string) {
	/* We start by looking for the Employee.csv file,
	 * and then we check to make sure the file was
	 * found and that everything is working				*/
	infile, err := os.Open("Employee.csv")
	if err != nil {
		log.Fatal(err)
	}

	/* Here we read in the file as a *csv.Reader 		*/
	employeeCSV := csv.NewReader(infile)

	/* This is where we begin to go line by line
	 * Looking through a string[] of each word,
	 * we also check to make sure it isnt reading
	 * past the end of file and make sure there
	 * arent any errors :)							*/
	for {
		record, err := employeeCSV.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		/* Here we compare the employee pin that is
		 * given (employeePin) and compare it with the
		 * list of pins in Employee.csv					*/
		if strings.Compare(employeePin, record[1]) == 0 {

			/* Now we read in if the user wants to clock
			 * in (or out). We use 'reader'	as a
			 * *bufio.Reader type for the user input and
			 * we go through the checks 				*/
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("do you want to clock in %s? (y/n)", record[2])
			fmt.Print("-> ")
			input, _ := reader.ReadString('\n')
			input = strings.Replace(input, "\n", "", -1)

			/* If the User wants to clock in we initalize
			 * 'e' as an EmployeeShift  type and fill it
			 * with the necessary information. Then we
			 * storethe employee while they are clocked
			 * in and give the user a clock in printout.*/
			if strings.Compare(input, "y") == 0 {
				var employee = logClockIn(record[0], record[1], record[2])
				storeEmployee(employee)
				clockInPrintout(employee)
				break
			} else if strings.Compare(input, "n") == 0 {
				fmt.Println("Okay :)")
				break
			} else {
				leaveCheck(input)
				break
			}

		}
	}
}

func getDailyLog() string {
	todayDt := time.Now().Format("2006-01-02")
	todayDt = todayDt + "-log.csv"
	return todayDt
}
func logClockIn(newName, newPin, newUser string) EmployeeShift {
	/*													*/
	/* This starts off pretty simple with getting the
	 * log file name  then filling up the variable
	 * 'employee' of type EmployeeShift. 				*/
	dt := time.Now()
	filename := getDailyLog()
	employee := EmployeeShift{
		pin:      newPin,     // String
		username: newUser,    // String
		sales:    0,          // float32
		ccTips:   0,          // float32
		clockIn:  dt,         // time.Time
		job:      "Your Job", // String
	}

	/* We check if the log file exists or need to create
	 * a new log file.				 					*/
	outFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	/* We write pin,username,hh:mm:ss,clockIn/clockOut
	 * into the log file. The pin is going to be used as
	 * A key later which is why we aere logging it here	*/

	var employeeInfo string
	employeeInfo = employee.pin + "," + employee.username +
		"," + employee.clockIn.Format("15:04:05") + ",Clocked IN" + "\n"

	if _, err := outFile.Write([]byte(employeeInfo)); err != nil {
		outFile.Close()
		log.Fatal(err)
	}
	if err := outFile.Close(); err != nil {
		log.Fatal(err)
	}
	return employee
}

func storeEmployee(e EmployeeShift) {

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	loadingScreen()
	for {
		fmt.Print("\n\n\nType (q) to quit.\n")
		fmt.Print("Enter an employee pin to clock in/out\n")
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
