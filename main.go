package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/xmp-er/rcrd/helpers"
)

func main() {

	//Defining the flag variables
	var frameRates int
	var library string
	var audio_output string
	var select_path string
	var op_format string

	flag.IntVar(&frameRates, "r", 30, "Specify the framerates of the video output")
	flag.StringVar(&library, "c:v", "", "Specify the encoding library")
	flag.StringVar(&audio_output, "c:a", "", "Specify the audio output format")
	flag.StringVar(&select_path, "p", "n", "Specify if you want to specify location on startup as y/n")
	flag.StringVar(&op_format, "f", "mkv", "Specify the format of the output video")

	flag.Parse()

	//Converting framerates to string
	frames := strconv.Itoa(frameRates)

	var cmd *exec.Cmd

	//Check the local file if there is a savepath or if flag to prompt for path is enable, if not prompt the user to select so, in the other case, move along with the default path
	loc_Path, err := os.ReadFile("save_path.txt")
	if err != nil || select_path == "y" {
		if select_path == "y" {
			temp_path, err_temp := helpers.SelectWorkingDirectory()
			if err_temp != nil {
				return
			}
			err_temp = os.WriteFile("save_path.txt", []byte(temp_path), 0644)
			if err_temp != nil {
				fmt.Println("Error while writing the path to file ", err)
				return
			}
			loc_Path = []byte(temp_path)
		} else if t_e, ok := err.(*os.PathError); ok {
			fmt.Println("No file selected ", t_e)
			temp_path, err_temp := helpers.SelectWorkingDirectory()
			if err_temp != nil {
				return
			}
			err_temp = os.WriteFile("save_path.txt", []byte(temp_path), 0644)
			if err_temp != nil {
				fmt.Println("Error while writing the path to file ", err)
				return
			}
			loc_Path = []byte(temp_path)
		}
	}
	local_path := string(loc_Path)

	//Get the FFMpeg command according to the specified Operating System
	operating_system := helpers.GetOS()
	cmd, err = helpers.GetCommand(operating_system, audio_output, library, frames, local_path, op_format)

	fmt.Println("The command is ", cmd)
	if err != nil {
		return
	}

	//Execute the command and listen for cancellation signal

	var stderr bytes.Buffer //So we are aware if we get a error in between
	cmd.Stderr = &stderr

	err = cmd.Start()
	if err != nil {
		fmt.Println("Error while starting the command", err)
	}

	cmd.Wait()
	if len(stderr.String()) > 0 {
		fmt.Println(stderr.String())
	}

}
