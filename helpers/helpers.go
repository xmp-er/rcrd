package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"time"

	"github.com/sqweek/dialog"
)

// Get the OS currently being used
func GetOS() string {
	os := runtime.GOOS
	return os
}

func GetCommand(operating_system string, aud_op string, lib string, fps string, outputPath string, output_format string) (*exec.Cmd, error) {
	var cmd *exec.Cmd
	timestamp := time.Now().Unix()
	folder_name := time.Now().Format("02_January_2006")
	file_name := time.Unix(timestamp, 0).Format("15:04:05")
	switch operating_system {
	case "darwin":
		if fps == "" {
			fps = "30"
		}
		if lib == "" {
			lib = "h264"
		}
		if aud_op == "" {
			aud_op = "aac"
		}
		if output_format == "" {
			output_format = "mkv"
		}
		if outputPath == "" {
			temp_path, err := GetDesktopFolder()
			if err != nil {
				fmt.Println("Error in getting Desktop folder ", err)
				return nil, err
			}
			outputPath = temp_path

		}
		err := os.MkdirAll(outputPath+"/"+folder_name, 0755)
		if err != nil {
			fmt.Println("Error creating directory ", err)
			return nil, err
		}
		final_output := outputPath + "/" + folder_name + "/" + file_name + "." + output_format
		cmd = exec.Command("ffmpeg", "-f", "avfoundation", "-i", "1:none", "-f", "avfoundation", "-r", fps, "-i", ":1", "-c:v", lib, "-c:a", aud_op, final_output)
	case "windows":
		if fps == "" {
			fps = "30"
		}
		if lib == "" {
			lib = "libx264"
		}
		if aud_op == "" {
			aud_op = "aac"
		}
		if output_format == "" {
			output_format = "mkv"
		}
		if outputPath == "" {
			temp_path, err := GetDesktopFolder()
			if err != nil {
				fmt.Println("Error in getting Desktop folder ", err)
				return nil, err
			}
			outputPath = temp_path

		}
		final_output := outputPath + "/" + folder_name + "/" + file_name + "." + output_format
		cmd = exec.Command("ffmpeg", "-f", "gdigrab", "-framerate", fps, "-i", "desktop", "-f", "dshow", "-i", "audio=virtual-audio-capturer", "-c:v", lib, "-c:a", aud_op, final_output)
	case "linux":
		if fps == "" {
			fps = "30"
		}
		if lib == "" {
			lib = "libx264"
		}
		if aud_op == "" {
			aud_op = "aac"
		}
		if output_format == "" {
			output_format = "mkv"
		}
		if outputPath == "" {
			temp_path, err := GetDesktopFolder()
			if err != nil {
				fmt.Println("Error in getting Desktop folder ", err)
				return nil, err
			}
			outputPath = temp_path

		}
		final_output := outputPath + "/" + folder_name + "/" + file_name + "." + output_format
		cmd = exec.Command("ffmpeg", "-f", "x11grab", "-framerate", fps, "-i", ":1.0", "-f", "alsa", "-i", "default", "-c:v", lib, "-c:a", aud_op, final_output)
	}
	return cmd, nil
}

func GetDesktopFolder() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	desktopPath := filepath.Join(usr.HomeDir, "Desktop")
	return desktopPath, nil
}

func SelectWorkingDirectory() (string, error) {
	selectedDir, err := dialog.Directory().Title("Select Location").Browse()
	if err != nil {
		fmt.Println("Error selecting directory ", err)
		return "", err
	}
	return selectedDir, nil
}
