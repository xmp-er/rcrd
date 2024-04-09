
![for_revered_homie](https://github.com/xmp-er/rcrd/assets/107166230/b96ddf27-e31b-46bc-986a-280df1b66fe5)

# Index
   - [🔴Introduction](#introduction)
   - [🎃Installation](#installation)
   - [🎃Usage](#usage)
   - [🎃Dependencies](#dependencies)
   - [🎃FFMpeg Installation Instructions](#ffmpeg-installation-instructions)
     - [🔺For Darwin (macOS)](#for-darwin-macos)
     - [🔺For Linux (Ubuntu)](#for-linux-ubuntu)
     - [🔺For Windows](#for-windows)

## 🔴Introduction
[Index](#index)

CLI Tool to record video and audio via selected source, uses FFMpeg in Backend
,initially a prompt will appear for save location that can be modified later, file is stored in the format `<date_month_year>/<timestamp>.<output-format>`

## 🎃Installation
[Index](#index)
1. Clone the repository.
2. Ensure you have FFMpeg installed on your system.
3. Run the following commands:
   - `go build` to build the project.
   - `./rcrd` to execute the binary file generated.

## 🎃Usage
[Index](#index)
- Use the following command line flags to specify settings:
   - `-r <framerate>`: Specify the framerates of the video output.
   - `-c:v <library>`: Specify the encoding library.
   - `-c:a <audio_output>`: Specify the audio output format.
   - `-p <select_path>`: Specify if you want to specify location on startup as y/n.
   - `-f <op_format>`: Specify the format of the output video.
   - `-vi <video_input>`: Specify the video input source index.
   - `-ai <audio_input>`: Specify the audio input source index.

## 🎃Dependencies
[Index](#index)
  - FFMpeg
  - Go programming language
  
## 🎃FFMpeg Installation Instructions
[Index](#index)
#### 🔺For Darwin (macOS)
`brew install ffmpeg`
#### 🔺For Linux (Ubuntu)
`sudo apt-get install ffmpeg`
#### 🔺For Windows
 - Download the latest build from the official FFMpeg website: https://ffmpeg.org/download.html
 - Extract the downloaded zip file.
 - Add the FFMpeg bin directory to the system's PATH environment variable.
