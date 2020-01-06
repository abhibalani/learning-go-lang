package main

import (
	"bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
  "text/tabwriter"
)

const getNoOfStudents = 1
const getStudentId = 2
const getStudentName = 3
const getMidTermScore = 4
const getSemScore = 5
const getAttendanceScore = 6
const invalidScoreMessage = "Invalid input for score. Value should be between 0 to 100"

func main(){


	numberOfStudents := getUserInput(getNoOfStudents).(int)

	var students []student =make([]student, numberOfStudents)

	for i:= 0; i < numberOfStudents; i++ {
		fmt.Println("Enter details of student", i+1)
		
		students[i].num = i+1
		students[i].id = getUserInput(getStudentId).(int32)
		students[i].name = getUserInput(getStudentName).(string)
		students[i].midTermScore = getUserInput(getMidTermScore).(float64)
		students[i].semesterScore = getUserInput(getSemScore).(float64)
		students[i].attendanceScore = getUserInput(getAttendanceScore).(float64)
		students[i].finalScore = students[i].getStudentFinalScore()
		students[i].grade = students[i].getStudentGrade(students[i].finalScore)
	}
	
	printResults(students)
	printInformation(numberOfStudents, students)

}

type student struct{
	num int
	id int32
	name string
	midTermScore float64
	semesterScore float64
	finalScore float64
	grade string
	attendanceScore float64

}

func printResults(students []student){
	w := tabwriter.NewWriter(os.Stdout, 12, 10, 1, ' ', tabwriter.Debug)
	fmt.Println("==========================================================")
	fmt.Fprintln(w, "No.\tStudent ID\tName\tFinal Score\t Grade")
	fmt.Fprintln(w, "==========================================================")

	for i:= 0; i < len(students); i++ {
		fmt.Fprintln(w, strconv.Itoa(students[i].num) + "\t" + strconv.Itoa(int(students[i].id)) + "\t" + students[i].name + "\t" + strconv.FormatFloat(students[i].finalScore, 'f', 2, 64)  + "\t" + students[i].grade)
	}
	fmt.Fprintln(w, "==========================================================")
	w.Flush()
}

func printInformation(numberOfStudents int, students []student){
	failedStudentCount := countFailedStudents(students)
	passedStudentCount := numberOfStudents - failedStudentCount
	w := tabwriter.NewWriter(os.Stdout, 20, 10, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "Number of Students\t:\t" + strconv.Itoa(numberOfStudents))
	fmt.Fprintln(w, "Number of Passing Students\t:\t" + strconv.Itoa(passedStudentCount))
	fmt.Fprintln(w, "Failed Students\t:\t" + strconv.Itoa(failedStudentCount))
	w.Flush()

}

func (s student) getStudentFinalScore() float64{
	return (0.2 * s.attendanceScore) + (0.4 * s.midTermScore) + (0.4 * s.semesterScore)
}

func (s student) getStudentGrade( finalScore float64) string{

	var grade string
	switch {
	case finalScore >= 85 && finalScore <= 100:
		grade = "A"
	
	case finalScore >= 76 && finalScore <= 84:
		grade = "B"
	
	case finalScore >= 61 && finalScore <= 75:
		grade = "C"
	
	case finalScore >= 46 && finalScore <= 60:
		grade = "D"
	
	case finalScore >= 0 && finalScore <= 45:
		grade = "E"
	
	default:
		grade = "E"
	}
	return grade
}

func countFailedStudents(students []student) int{
	count := 0
	for i:= 0; i < len(students); i++ {
		if students[i].grade == "D" || students[i].grade == "E" {
			count+=1
		}
	}
	return count
}

func getInput() string{
	reader := bufio.NewReader(os.Stdin)
	ip, _ := reader.ReadString('\n')
	ip = strings.Replace(ip, "\n", "", -1)
	return ip
}

func isScoreValid(score float64) bool{

	if score >= 0 && score <= 100 {
		return true
	} else {
		return false
	}

}
func getUserInput(ipType int) interface{}{
	var inputData interface{}
	switch ipType {

		case getNoOfStudents: 
			fmt.Println("Enter the number of students:")
			data, err := strconv.ParseInt(getInput(), 10, 8)
			if  err != nil {
				fmt.Println("Invalid input for number of students. Please enter an integer value.")
				getUserInput(getNoOfStudents)
			} 
			inputData = int(data)
			
		case getStudentId: 
			fmt.Println("Enter the Student ID:")
			data, err := strconv.ParseInt(getInput(), 10, 32)
			if  err != nil {
				fmt.Println("Invalid input for Student ID. Please enter an integer value.")
				getUserInput(getStudentId)
			} 
			inputData = int32(data)	
			

		case getStudentName: 
			fmt.Println("Enter the Student Name:")
			data := getInput()
			inputData = data

		case getMidTermScore: 
			fmt.Println("Enter Midterm Score:")
			data, err := strconv.ParseFloat(getInput(), 64)
			
			if  err != nil || !isScoreValid(float64(data)){
				fmt.Println(invalidScoreMessage)
				getUserInput(getMidTermScore)
			} 
			inputData = float64(data)	
		
		case getSemScore: 
			fmt.Println("Enter Semester Score:")
			data, err := strconv.ParseFloat(getInput(), 64)
			if  err != nil || !isScoreValid(float64(data)) {
				fmt.Println(invalidScoreMessage)
				getUserInput(getSemScore)
			} 
			inputData = float64(data)			
			
		case getAttendanceScore: 
				fmt.Println("Enter Attendance Score:")
				data, err := strconv.ParseFloat(getInput(), 64)
				if  err != nil || !isScoreValid(float64(data)) {
					fmt.Println(invalidScoreMessage)
					getUserInput(getSemScore)
				} 
				inputData = float64(data)	
					
	}
	return inputData
}

