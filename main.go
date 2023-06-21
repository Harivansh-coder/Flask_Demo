package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"harry/get-pract/core"
	"harry/get-pract/model"
	"harry/get-pract/utils"
)

func main() {

	// array containing all the subjects
	available_subjects := []string{"amdl", "eiot", "wm", "nlp"}

	// help code
	for _, val := range os.Args {
		if val == "help" {
			fmt.Println(utils.GetHelp())
			return
		}
	}

	// getmyexe code
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments")
		return
	}

	//  extracting the subject name and practical number
	subject_name := os.Args[1]

	// check if subject is available
	subject_available := utils.CheckSubject(subject_name)

	if !subject_available {
		fmt.Println("Subject not available or not correct, these are the available subjects: ")
		fmt.Println(available_subjects)
		return

	}

	if len(os.Args) == 3 {

		//  extracting the practical number

		pract_no := os.Args[2]

		// make an api call to get the file
		URL := "https://www.cspracs.tech/api"

		// get the file url from the api

		response_object, err := http.Get(fmt.Sprintf("%s/%s/%s", URL, subject_name, pract_no))

		if err != nil {
			fmt.Println("Error in making api call")
			return
		}

		defer response_object.Body.Close()

		// check if the response is 200
		if response_object.StatusCode != http.StatusOK {
			fmt.Println("Error in making api call")
			return
		}

		// Read the response body
		response_body, err := io.ReadAll(response_object.Body)

		if err != nil {
			fmt.Println("Error in reading response body")
			return
		}

		// creating an instane of the response struct
		var response model.Response

		// unmarshalling the response body

		err = json.Unmarshal(response_body, &response)

		if err != nil {
			fmt.Println("Error in unmarshalling response body")
			return
		}

		// download the source code
		file := core.GetFileFromURL(response.Txt)

		//  print the file data
		fmt.Println("Subject Name: ", subject_name)
		fmt.Println("File downloaded: ", file)
		fmt.Println("File Name: ", response.Name)
		fmt.Println("File Aim: ", response.Aim)
		fmt.Println("File URL: ", response.Txt)

	} else if len(os.Args) == 2 {

		// make an api call to get the file
		URL := "https://www.cspracs.tech/api"

		// get the file url from the api
		response_object, err := http.Get(fmt.Sprintf("%s/%s", URL, subject_name))

		if err != nil {
			fmt.Println("Error in making api call")
			return
		}

		defer response_object.Body.Close()

		// check if the response is 200
		if response_object.StatusCode != http.StatusOK {
			fmt.Println("Error in making api call")
			return
		}

		// Read the response body
		response_body, err := io.ReadAll(response_object.Body)

		if err != nil {
			fmt.Println("Error in reading response body")
			return
		}

		// creating a slice of response structs
		var responses []model.Response
		fmt.Println("Subject Name: ", subject_name)

		// unmarshalling the response body
		err = json.Unmarshal(response_body, &responses)
		if err != nil {
			fmt.Println("Error in unmarshalling response body")
			return
		}

		// iterate over the responses
		for _, response := range responses {
			// print the file data
			fmt.Println("File Name:", response.Name)
			fmt.Println("Aim:", response.Aim)
		}

	}
}
