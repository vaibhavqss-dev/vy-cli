package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func HandleBackup(verbose bool) {
	// Check if rclone is installed and configured
	err := checkRcloneInstallation()
	if err != nil {
		fmt.Println(err)
		return
	}


	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		return
	}

	filesToBackup := make(map[string]string)
	possibleFiles := map[string]string{
		// Terminal and shell settings
		".bashrc":       filepath.Join(homeDir, ".bashrc"),
		".profile":      filepath.Join(homeDir, ".profile"),
		".zshrc":        filepath.Join(homeDir, ".zshrc"),
		".bash_aliases": filepath.Join(homeDir, ".bash_aliases"),
		".zsh_aliases":  filepath.Join(homeDir, ".zsh_aliases"),
		"custom-scripts": filepath.Join(homeDir, "bin"), // Custom scripts directory
		
		// Desktop environment settings
		"dconf-settings": filepath.Join(homeDir, ".config/dconf"),
		"gnome-terminal": filepath.Join(homeDir, ".config/gnome-terminal"),
		"gtk-3.0":       filepath.Join(homeDir, ".config/gtk-3.0"),
		"gtk-4.0":       filepath.Join(homeDir, ".config/gtk-4.0"),
		"nautilus":      filepath.Join(homeDir, ".config/nautilus"),
		"autostart":     filepath.Join(homeDir, ".config/autostart"),
		
		// Theme and appearance
		"themes":         filepath.Join(homeDir, ".themes"),
		"icons":          filepath.Join(homeDir, ".icons"),
		"backgrounds":    filepath.Join(homeDir, ".local/share/backgrounds"),
		"fonts":          filepath.Join(homeDir, ".fonts"),
		"local-fonts":    filepath.Join(homeDir, ".local/share/fonts"),
		
		// Application preferences
		"mimeapps":      filepath.Join(homeDir, ".config/mimeapps.list"),
		"user-dirs":     filepath.Join(homeDir, ".config/user-dirs.dirs"),
		"preferences":   filepath.Join(homeDir, ".local/share/preferences"),
		"custom-launchers": filepath.Join(homeDir, ".local/share/applications"),
		"plank":         filepath.Join(homeDir, ".config/plank"),
		
		// SSH keys and configuration
		"ssh":           filepath.Join(homeDir, ".ssh"),
		
		// Environment and system settings
		"pam-environment": filepath.Join(homeDir, ".pam_environment"),
		"xprofile":        filepath.Join(homeDir, ".xprofile"),
		"xinitrc":         filepath.Join(homeDir, ".xinitrc"),
		
		// Developer tools and editors
		"vimrc":         filepath.Join(homeDir, ".vimrc"),
		"vim":           filepath.Join(homeDir, ".vim"),
		"nvim":          filepath.Join(homeDir, ".config/nvim"),
		"tmux":          filepath.Join(homeDir, ".tmux.conf"),
		"gitconfig":     filepath.Join(homeDir, ".gitconfig"),
		
		// Systemd user services
		"systemd-user":  filepath.Join(homeDir, ".config/systemd/user"),
	}
	
	for name, path := range possibleFiles {
		if _, err := os.Stat(path); err == nil {
			filesToBackup[name] = path
		}
	}
	

	// Define backup directory on OneDrive
	backupDir := "Backups/ubuntu-settings"

	// Track backup status
	successCount := 0
	totalFiles := len(filesToBackup)
	
	fmt.Println("Please wait.... I'm Uploading files to OneDrive.....\nThis Will Take Time Depending Upon Speed of Internet and Size of Folder :) ...")
	
	if(verbose){
		fmt.Println("Starting Ubuntu settings backup...")
		fmt.Printf("Total configurations to backup: %d\n\n", totalFiles)
	}
	
	// Process each file individually
	for name, path := range filesToBackup {
		if _, err := os.Stat(path); err != nil {
			fmt.Printf("  ❌ Skipping: configuration not found\n\n")
			continue
		}

		if(verbose){
			fmt.Printf("📤 Uploading %s to OneDrive... ", name)
		}


		err = rclone(path, filepath.Join(backupDir, name))
		if err != nil {
			fmt.Printf("❌ Failed\n  Error: %v\n\n", err)
			continue
		}

		if(verbose){
			fmt.Printf("✅ Success\n\n")
		}

		successCount++
	}

	fmt.Printf("Backup completed! Successfully backed up %d of %d configurations\n", successCount, totalFiles)
}

func rclone(localFilePath, destinationPath string) error {
	// Check if the destination directory exists, if not create it
	checkCmd := exec.Command("rclone", "lsf", "ubuntuSettings_onedrive:"+filepath.Dir(destinationPath))
	if err := checkCmd.Run(); err != nil {
		mkdirCmd := exec.Command("rclone", "mkdir", "ubuntuSettings_onedrive:"+filepath.Dir(destinationPath))
		if err := mkdirCmd.Run(); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}

	// Proceed with the file copy
	cmd := exec.Command("rclone", "copy", localFilePath, "ubuntuSettings_onedrive:"+destinationPath, "--progress")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("rclone error: %w\nOutput: %s\nError: %s",
			err, stdout.String(), stderr.String())
	}

	return nil
}

func checkRcloneInstallation() error {
	// Check if rclone is installed
	_, err := exec.LookPath("rclone")
	if err != nil {
		return fmt.Errorf("rclone is not installed. Please install it using:\n" +
			"curl https://rclone.org/install.sh | sudo bash\n" +
			"Then authenticate with: rclone config")
	}

	// Check if rclone is configured with ubuntuSettings_onedrive
	cmd := exec.Command("rclone", "listremotes")
	output, err := cmd.Output()
	if err != nil || !bytes.Contains(output, []byte("ubuntuSettings_onedrive:")) {
		return fmt.Errorf("rclone is not configured with ubuntuSettings_onedrive.\n" +
			"Please run 'rclone config' and set up your OneDrive connection")
	}

	return nil
}
