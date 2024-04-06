package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"

	"github.com/xmp-er/rcrd/helpers"
)

func main() {

	var cmd *exec.Cmd

	operating_system := helpers.GetOS()

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory: ", err)
		return
	}

	outputFile := "output.mp4"
	outputPath := filepath.Join(currentDir, outputFile)

	switch operating_system {
	case "darwin":
		cmd = exec.Command("ffmpeg", "-f", "avfoundation", "-i", "1:none", "-f", "avfoundation", "-i", ":1", "-c:v", "libx264", "-c:a", "aac", outputPath)
	case "windows":
		cmd = exec.Command("ffmpeg", "-f", "gdigrab", "-framerate", "30", "-i", "desktop", "-f", "dshow", "-i", "audio=virtual-audio-capturer", "-c:v", "libx264", "-c:a", "aac", outputPath)
	case "linux":
		cmd = exec.Command("ffmpeg", "-f", "x11grab", "-framerate", "25", "-i", ":0.0", "-f", "alsa", "-i", "default", "-c:v", "libx264", "-c:a", "aac", outputPath)
	}

	err = cmd.Start()
	if err != nil {
		fmt.Println("Error while starting the command ", err)
	}
	fmt.Println("recording started")
	time.Sleep(5 * time.Second)
	err = cmd.Process.Signal(syscall.SIGINT)
	if err != nil {
		fmt.Println("Error while stopping the process ", err)
	}

}
