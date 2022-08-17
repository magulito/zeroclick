package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type Configs struct {
	Logs_path string `json:"logs_path"`
}

type Commands struct {
	Commands []Command `json:"commands"`
}

type Command struct {
	Name             string   `json:"name"`
	Command          string   `json:"command"`
	OS               string   `json:"os"`
	Path             string   `json:"patch"`
	Network_elements []string `json:"network_elements"`
}

func main() {

	//Define the version number of the tool
	version := "2.0.1"

	//Get the name of the host
	hostname, _ := os.Hostname()
	currentTime := time.Now()

	//Define the name of the logfile
	log_filename := appConfigs(hostname)

	//Sets log settings
	f, err := os.OpenFile(log_filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err.Error())
		log.Println("Error " + err.Error() + "Please, read the README.txt file for more information.")
	}

	defer f.Close()

	log.SetOutput(f)

	log.Println("zeroClick - Automated Checklist - version " + version)
	fmt.Printf("\n\t\tzeroClick - Automated Checklist - version %v", version)

	log.Println("Hostname: " + hostname + "\t-\tDate and time: " + currentTime.Format("2006-01-02 15:04:05"))
	// log.Println("Date and time: " + currentTime.Format("2006-01-02 15:04:05"))

	// Open the json file
	jsonFile, err := os.Open("../configs/commands.json")

	if err != nil {
		log.Println(err.Error())

		os.Exit(1)

	}

	// Defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// Read the command.json file as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var cmds Commands

	json.Unmarshal(byteValue, &cmds)

	log.Println("Checklist has started running on " + hostname + ".")
	// fmt.Println("Checklist has started running on " + hostname + ".")

	for i := 0; i < len(cmds.Commands); i++ {

		if cmds.Commands[i].OS == runtime.GOOS {

			for n := 0; n < len(cmds.Commands[i].Network_elements); n++ {

				if cmds.Commands[i].Network_elements[n] == "all" {

					log.Println("Started running command " + cmds.Commands[i].Name + ".")

					run_command := run_command(cmds.Commands[i].Command, cmds.Commands[i].Path, cmds.Commands[i].OS)

					log.Println("<commandOutput>" + cmds.Commands[i].Name + ":" + run_command + "</commandOutput>")

					log.Println("Finished running command " + cmds.Commands[i].Name + ".")

				} else {

					if hostname == cmds.Commands[i].Network_elements[n] {

						log.Println("Started running command " + cmds.Commands[i].Name + ".")

						run_command := run_command(cmds.Commands[i].Command, cmds.Commands[i].Path, cmds.Commands[i].OS)

						log.Println("<commandOutput>" + cmds.Commands[i].Name + ":" + run_command + "</commandOutput>")

						log.Println("Finished running command " + cmds.Commands[i].Name + ".")

					}
				}
			}
		}
	}
	log.Println("Finished running the checklist.")

	fmt.Println("\n\t\t  Vodacom Mozambique - PrepaidSystems - 2022")

}

func run_command(command string, path string, os string) string {

	if runtime.GOOS == "linux" {

		cmd, err := exec.Command("bash", "-c", command, path).Output()

		if err != nil {

			return err.Error()
		}

		return string(cmd)
	}

	if runtime.GOOS == "windows" {

		cmd, err := exec.Command("cmd", command, path).Output()

		if err != nil {

			return err.Error()
		}

		return string(cmd)
	}

	fmt.Println("This tool has been designed to run on windows or linux operating systems. Please, contact your systems administrator.")
	return "This tool has been designed to run on windows or linux operating systems. Please, contact your systems administrator."

}

func appConfigs(hostname string) string {

	currentTime := time.Now()
	// Open the json file

	jsonConfigFile, err := os.Open("../configs/configs.json")

	//Check if there was an error. If yes, create the default log path and use it.
	if err != nil {

		os.MkdirAll("logs", os.ModePerm)

		return "../logs/checklist_" + currentTime.Format("20060102150405") + "_" + hostname + ".log"
	}

	// Defer the closing of our jsonFile so that we can parse it later on
	defer jsonConfigFile.Close()

	// Read the command.json file as a byte array.
	byteValue, err := ioutil.ReadAll(jsonConfigFile)

	//Check if there was an error. If yes, create the default log path and use it.
	if err != nil {

		os.MkdirAll("logs", os.ModePerm)

		return "../logs/checklist_" + currentTime.Format("20060102150405") + "_" + hostname + ".log"
	}

	var configs Configs

	json.Unmarshal(byteValue, &configs)

	//Check if the path exists. If not, create it.
	if _, err := os.Stat(configs.Logs_path); os.IsNotExist(err) {

		os.MkdirAll(configs.Logs_path, os.ModePerm)
	}

	return configs.Logs_path + "checklist_" + currentTime.Format("20060102150405") + "_" + hostname + ".log"

}
