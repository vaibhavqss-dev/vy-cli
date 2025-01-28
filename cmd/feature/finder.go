package feature

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FindFile(fileName string, fileSize int, hardSearch bool) {

	fmt.Printf("\n\033[34m\033[1m" +
		"███████╗██╗██╗     ███████╗    ███████╗██╗███╗   ██╗██████╗ ███████╗██████╗ \n" +
		"██╔════╝██║██║     ██╔════╝    ██╔════╝██║████╗  ██║██╔══██╗██╔════╝██╔══██╗\n" +
		"█████╗  ██║██║     █████╗      █████╗  ██║██╔██╗ ██║██║  ██║█████╗  ██████╔╝\n" +
		"██╔══╝  ██║██║     ██╔══╝      ██╔══╝  ██║██║╚██╗██║██║  ██║██╔══╝  ██╔══██╗\n" +
		"██║     ██║███████╗███████╗    ██║     ██ ██║ ╚████║██████╔╝███████╗██║  ██║\n" +
		"╚═╝     ╚═╝╚══════╝╚══════╝    ╚═╝     ╚═╝╚═╝  ╚═══╝╚═════╝ ╚══════╝╚═╝  ╚═╝\n\033[0m")

	err := filepath.Walk("/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		// Skip directories
		if !hardSearch {
			// Skip directories
			var excludeFilesFolder = []string{"node_modules", "tmp", "proc", "sys", "dev", "run",
				".snap", ".deb", ".git", ".github", ".vscode", ".idea", ".gitignore", ".gitattributes",
				".gitmodules", ".gitkeep", ".gitlab-ci.yml", ".gitlab", ".nvm", ".cache", ".npm", ".yarn",
				".docker", ".dockerignore", ".dockerfile", ".docker-compose", ".docker-compose.yml",
				".config"}

			baseName := filepath.Base(path)
			for i := range excludeFilesFolder {
				if strings.HasSuffix(baseName, excludeFilesFolder[i]) {
					return filepath.SkipDir
				}
			}
		}

		if strings.Contains(strings.ToLower(filepath.Base(path)), strings.ToLower(fileName)) {
			fmt.Printf("\n\033[32m\033[1m\033[5mFOUND!\033[0m")
			fmt.Printf("\n\033[32mLocation: %s (%d MB)\033[0m \n\n", path, info.Size()/1024/1024)
			os.Exit(0)
		}

		if info.Size()/1024/1024 >= int64(fileSize) {
			if fileSize >= 5 {
				fmt.Printf("\033[33mSize Found: %s (%d bytes) (%d MB)\033[0m \n", path, info.Size(), info.Size()/1024/1024)
			}
		}

		return nil
	})
	if err == nil {
		fmt.Printf("\n\033[31m\033[1mNo File or Directory Found\033[0m\n")
		fmt.Println("Make sure you've typed the correct file or Folder Name")
	}
	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
	}
}
