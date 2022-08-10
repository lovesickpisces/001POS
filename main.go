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
			 * in (or out).								*/
			input := getFeedback("do you want to clock in " + record[2] + " (y/n)")

			/* If the User wants to clock in we initalize
			 * 'e' as an EmployeeShift  type and fill it
			 * with the necessary information. Then we
			 * store the employee while they are clocked
			 * in and give the user a clock in printout.*/
			if strings.Compare(input, "y") == 0 {

				/* Here is the jobs part*/
				count := 1
				jobsList := getEmployeeJobs(employeePin)
				jobInput := getFeedback("Please type the number of the job you are working today")
				var employeeJob string
				if strings.Compare(jobInput, "") != 0 {
					for i := 0; i < 7; i++ {
						if strings.Compare(jobsList[i], "n") == 0 {
							i++
						} else {
							jobInputInt, err := strconv.Atoi(jobInput)
							if err != nil {
								log.Fatal(err)
							} else if count == jobInputInt {
								employeeJob = jobsList[i]
							}
							count++
						}
					}
				}
				var employee = logClockIn(record[0], record[1], record[2], employeeJob)
				clockInPrintout(employee)
				break
			} else if strings.Compare(input, "n") == 0 {
				fmt.Println("Okay :)")
				break
			}
		}
	}
}

func getDailyLog() string {
	todayDt := time.Now().Format("2006-01-02")
	todayDt = "dailyLogs/" + todayDt + "-log.csv"
	return todayDt
}
func logClockIn(newName, newPin, newUser, employeeJob string) EmployeeShift {

	/* This starts off pretty simple with getting the
	 * log file name  then filling up the variable
	 * 'employee' of type EmployeeShift. 				*/
	dt := time.Now()
	filename := getDailyLog()
	employee := EmployeeShift{
		pin:      newPin,      // String
		username: newUser,     // String
		sales:    0,           // float32
		ccTips:   0,           // float32
		clockIn:  dt,          // time.Time
		job:      employeeJob, // String
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
	employeeInfo := employee.pin + "," + employee.username +
		"," + employee.clockIn.Format("15:04:05") + ",Clocked IN" + "," + employee.job + "\n"

	if _, err := outFile.Write([]byte(employeeInfo)); err != nil {
		outFile.Close()
		log.Fatal(err)
	}
	if err := outFile.Close(); err != nil {
		log.Fatal(err)
	}
	return employee
}
func getEmployeeJobs(employeePin string) [7]string {

	/* This begins with us setting up an array to be
	 * returned and an array with the strings to assign
	 * later to the returned array . There is also a
	 * Count used to keep track of how many to list		*/
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

	/* Here we access the jobs.csv file that keeps track
	 * who has what job. 								*/
	infile, err := os.Open("jobs.csv")
	if err != nil {
		log.Fatal(err)
	}

	/* Kind of just repeating what we did to go through
	 * CSVs in the other functions						*/
	jobsCSV := csv.NewReader(infile)
	for {
		record, err := jobsCSV.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if strings.Compare(employeePin, record[0]) == 0 {
			fmt.Println()
			for i := 2; i <= 8; i++ {
				if strings.Compare(record[i], "y") == 0 {
					fmt.Printf("(%d) %s \n", count+1, jobs[i-2])
					returnedJobs[count] = jobs[i-2]
					count++
				} else {
					returnedJobs[i-2] = "n"
					// count++
				}
			}
			break
		}
	}
	return returnedJobs
}

/*														*/
/* This function just helps minamize the amount of times
 * You need to check input and everything :)			*/
func getFeedback(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nType (q) to quit.\n")
	fmt.Printf("%s\n", prompt)
	fmt.Print("-> ")
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	leaveCheck(input)
	return input
}

func main() {
	loadingScreen()
	for {
		input := getFeedback("Enter an employee pin to clock in/out")
		if strings.Compare(input, "") != 0 {
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
