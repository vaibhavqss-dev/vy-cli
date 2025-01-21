package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func PrintRainbowGlowLargeText(text string) {
	colors := []string{
		"\033[31m", // Red
		"\033[33m", // Yellow
		"\033[32m", // Green
		"\033[36m", // Cyan
		"\033[34m", // Blue
		"\033[35m", // Magenta
	}

	cmd := exec.Command("figlet", text)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running figlet:", err)
		os.Exit(1)
	}

	// Clear the screen
	fmt.Print("\033[H\033[2J")

	// Loop through the generated ASCII art and print each line with rainbow color
	for i, line := range strings.Split(string(output), "\n") {
		color := colors[i%len(colors)] // Cycle through colors for each line
		fmt.Printf("%s%s\n", color, line)
	}
    
	fmt.Printf("\033[0m")
}


// Date and time
func Date() string {
    loc, err := time.LoadLocation("Asia/Kolkata")
    if err != nil {
        return ("Error loading location:")
    }
    return time.Now().In(loc).Format("Monday, 02 January 2006, 03:04:05 PM IST")
}