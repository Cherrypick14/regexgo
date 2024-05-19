package main

import (
	"fmt"
	"regexp"
)

func main() {
	logEntry := "2023-05-18 12:35:12 [ERROR] Failed to connect to the database."

	//Define the regular expression
	pattern := `^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2} \[(\w+)\] (.*)`
	re := regexp.MustCompile(pattern)

	// Match the pattern against log entry
	match := re.FindStringSubmatch(logEntry)

	if len(match) > 0 {
		timestamp := match[1]
		severity_level := match[2]
		message := match[3]

		fmt.Printf("Timestamp %s\n Severity %s\n Message %s\n", timestamp, severity_level, message)
	} else {
		fmt.Println("No match found")
	}

}
