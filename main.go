package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"github.com/vaibhavyadav-dev/vy-cli/cmd"
)

//go:embed cmd/cmd.txt
var cmdFile string // Embed the cmd.txt file

func main() {
	if len(os.Args) < 2 {
		if cmdFile == "" {
			fmt.Println("Seems like package is not successfully installed :(") 
			fmt.Println("Please install it with default configuration")
			log.Fatal("cmd.txt content is missing from the binary")
			os.Exit(1)
		}
		cmd.PrintRainbowGlowLargeText("Vaibhav Yadav")
		fmt.Println("Command line made For and By VAIBHAV YADAV")
		fmt.Println(cmdFile)
		
		os.Exit(1)
	}


	command := os.Args[1]

	switch command {
	case "date":
		fmt.Println(cmd.Date())
	case "backup":
		if len(os.Args) == 2 {
			cmd.HandleBackup(false)
		} else if len(os.Args) == 3 && os.Args[2] == "-v" { 
			cmd.HandleBackup(true)
		} else {
			fmt.Println("Invalid usage. Use 'backup' or 'backup -v' for verbose output")
			os.Exit(0)
		}
	case "commit":
		if len(os.Args) < 3{
			fmt.Println("Please provide commit message")
			os.Exit(0)
		}
		// Get the commit message
		msg := cmd.Commit(os.Args[2])
		fmt.Println(msg)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println(cmdFile)
		os.Exit(1)
	}
}
