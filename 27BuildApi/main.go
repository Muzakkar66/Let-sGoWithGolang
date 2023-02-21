package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//make a model -file
type Course struct {
	CourseId    string `json:"courseid"`
	CourseName  string `json:"coursename"`
	CoursePrice int    `json:"courseprice"`
	Author      string `json:"author"`
}
type Author struct {
	FullName string `json:"fullname"`
	WebSite  string `json: "website`
}

//fake DB
var courses []Course

//midleware , helper  -file
func (c *Course) IsEmpty() bool {
	return c.CourseId == " " && c.CourseName == ""
}
func main() {
	fmt.Println("welcome to the building api in golang")
	// serveHome()
	// createOneCourse()
}

//controllers --file

//serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello serve home func</h1>"))
}

//get all courses

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(courses)

}
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	//lopp through courses , find id and return  the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return

		}
	}
	json.NewEncoder(w).Encode("No course found with this given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create new course")
	w.Header().Set("Content-type", "application/json")
	//what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please enter some data ")
	}

	// what if - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}
	//generate unique id string
	//append course into course

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}
 func updateOneCourse(w http.ResponseWriter, r *http.Request)  {
	
 }