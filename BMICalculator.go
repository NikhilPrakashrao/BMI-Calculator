/*

 */

//This defines the package name of the program.
package main

/*This is a preprocessor command that tells the Go
compiler to include all files */
import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

//We declared the main() method from where the program execution begins
func main() {

	/*Here, we are registering the checkBMI function with the /bmi URL pattern
	using HandleFunc of the net/http package*/
	http.HandleFunc("/bmi", checkBMI)

	// A simple	print function to know that server has started
	fmt.Printf("your server is started at localhost:8080 port...\n")

	/*Here, we are calling http.ListenAndServe to serve HTTP requests*/
	err := http.ListenAndServe(":8080", nil)

	/*we check whether there is a problem starting the server.
	If there is, then log the error and exit code.*/
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

/*Here checkBMI gets executed.*/
func checkBMI(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		http.ServeFile(w, r, "bmi.html")

	}
	if r.Method == "POST" {

		/*This is a Go function that accepts ResponseWriter and Request
		as input parameters, parses bmi.html*/
		template, _ := template.ParseFiles("bmi.html")

		/*Here we parse the request body as a form.*/
		r.ParseForm()

		// These are the form values which we get after parsing the form
		fmt.Println(r.FormValue("height"))
		fmt.Println(r.FormValue("weight"))

		// Initialisation and assigning the variables to the form values
		Height := r.FormValue("height")
		Weight := r.FormValue("weight")

		// Converting the form values to float type.
		height, _ := strconv.ParseFloat(Height, 64)
		weight, _ := strconv.ParseFloat(Weight, 64)

		//Calculating the bmi
		bmi := weight / (height * height)
		if height <= 0 || weight <= 0 {
			fmt.Println("Values must be above 0")
		} else if bmi < (18.5) {
			fmt.Println("The person is underweight,And BMI is", bmi)
		} else if bmi > (18.5) && bmi < (25) {
			fmt.Println("The person is normal,And BMI is", bmi)
		} else if bmi > 25 && bmi < 30 {
			fmt.Println("The person is overweight,And BMI is", bmi)
		} else if bmi > 30 {
			fmt.Println("The person is obese,And BMI is", bmi)
		}

		//mybmi := map[string]interface{}{"bmi": bmi}
		// t.Execute(w, mybmi)
		template.Execute(w, bmi)

	}
}
