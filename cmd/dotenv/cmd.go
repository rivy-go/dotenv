package main

import (
	"flag"
	"fmt"
	"log"

	"strings"

	"github.com/rivy-go/dotenv"
)

func main() {
	var showHelp bool
	flag.BoolVar(&showHelp, "h", false, "show help")
	var rawEnvFilenames string
	flag.StringVar(&rawEnvFilenames, "f", "", "comma separated paths to .env files")

	flag.Parse()

	usage := `
Run a process with a env setup from a .env file

dotenv [-f ENV_FILE_PATHS] COMMAND_ARGS

ENV_FILE_PATHS: comma separated paths to .env files
COMMAND_ARGS: command and args you want to run

example
  dotenv -f /path/to/something/.env,/another/path/.env fortune
`
	// if no args or -h flag
	// print usage and return
	args := flag.Args()
	if showHelp || len(args) == 0 {
		fmt.Println(usage)
		return
	}

	// load env
	var envFilenames []string
	if rawEnvFilenames != "" {
		envFilenames = strings.Split(rawEnvFilenames, ",")
	}

	// take rest of args and "exec" them
	cmd := args[0]
	cmdArgs := args[1:]

	err := dotenv.Exec(envFilenames, cmd, cmdArgs)
	if err != nil {
		log.Fatal(err)
	}
}
