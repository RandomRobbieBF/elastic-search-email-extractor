package main

import (
    "crypto/tls"
    "io/ioutil"
    "net/http"
    "bufio"
    "fmt"
    "os"
    "time"
    "regexp"

)

func searchURLs(url string) {

	// define a struct to hold the JSON data
	type Data struct {
    // other fields
    Email string `json:"email"`
	}
	
    // create an HTTP client that ignores bad SSL certificates
    tr := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
    Transport: tr,
    Timeout:   45 * time.Second,
	}

   

    // create and execute the HTTP request
    req, err := http.NewRequest("GET", url + "/_search?q=text&size=5000", nil)
    if err != nil {
        // handle error
    }
    req.Header.Set("Content-Type", "application/json")
    res, err := client.Do(req)
    if err != nil {
    fmt.Println(err)
	return
    }
    defer res.Body.Close()

    // read and print the response body
    responseBody, err := ioutil.ReadAll(res.Body)
    if err != nil {
         fmt.Println(err)
         return
    }
	

	// Compile the regular expression to extract email addresses
	re := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	// Extract all email addresses
	emails := re.FindAllString(string(responseBody), -1)
	// Remove duplicate email addresses
	emails = removeDuplicates(emails)
	
	// Extract Hostname
	
	// Print the extracted email addresses
	if len(emails) == 0 {
		fmt.Println("No email addresses found.")
	} else {
	
		
	
	
	
		fmt.Println("Found via: "+url+"")
		for _, email := range emails {
			
			fmt.Println(email)
	
	// Open or Create the text file with Append flag
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write the URL to the file
	_, err = file.WriteString(""+email+"   ---   "+url+" \n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
			
		}
	}
	}



// removeDuplicates removes duplicate elements from a string slice.
func removeDuplicates(s []string) []string {
	// Use a map to track unique elements
	m := make(map[string]bool)
	for _, item := range s {
		if _, ok := m[item]; !ok {
			m[item] = true
		}
	}

	// Convert the map keys back to a slice
	result := make([]string, 0, len(m))
	for item := range m {
		result = append(result, item)
	}
	return result
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a URL or file as an argument")
		return
	}

	input := os.Args[1]

	// Check if the argument is a file
	if _, err := os.Stat(input); err == nil {
		// Open the file
		f, err := os.Open(input)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		// Read the file line by line
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			// Set the URL variable as the scanned line
			url := scanner.Text()

			// Send the URL to the grabber function
			searchURLs(url)
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	} else {
		// Set the URL variable as the input argument
		url := input

		// Send the URL to the grabber function
		searchURLs(url)
	}
}
