package main

import (
	"os/exec"

	"github.com/xmp-er/rcrd/helpers"
)

func main() {
	var cmd *exec.Cmd

	os := helpers.GetOS()

	switch os {
	case "darwin":
		cmd = exec.Command("ffmpeg", "-f", "avfoundation", "-i", "1:none", "-f", "avfoundation", "-i", ":1", "-c:v", "libx264", "-c:a", "aac", "output.mp4")
	case "windows":
		cmd = exec.Command("ffmpeg", "-f", "gdigrab", "-framerate", "30", "-i", "desktop", "-f", "dshow", "-i", "audio=virtual-audio-capturer", "-c:v", "libx264", "-c:a", "aac", "output.mp4")
	case "linux":
		cmd = exec.Command("ffmpeg", "-f", "x11grab", "-framerate", "25", "-i", ":0.0", "-f", "alsa", "-i", "default", "-c:v", "libx264", "-c:a", "aac", "output.mp4")
	}
}
