package main

import (
	"flag"
	"fmt"
	godotenvgenerator "github.com/czhujer/go-dotenv-generator"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-password/password"
	"log"
)

var oldItems map[string]string
var newItems map[string]string

func main() {
	var err error

	var showHelp bool
	flag.BoolVar(&showHelp, "h", false, "show help")
	var envFile string
	flag.StringVar(&envFile, "f", ".env", "name of .env file (with path)")

	flag.Parse()

	usage := `
Generate var(s) for a .env file

godotenv [-f ENV_FILE] KEYS

ENV_FILE: name of .env file (with path)
KEYS: keys you want to generate

example
  godotenvgenerator -f path/.env DB_PASSWORD
`
	// if no args or -h flag
	// print usage and return
	args := flag.Args()
	if showHelp || len(args) == 0 {
		fmt.Println(usage)
		return
	}

	// take rest of args
	keysItems := args[0:]

	// load existing items
	oldItems, err = godotenv.Read(envFile)
	if err != nil {
		log.Fatal(err)
	}

	// generate password for new items
	for _, key := range keysItems {
		if _, found := oldItems[key]; found {
			break
		}
		password, err := password.Generate(10, 2, 0, true, false)
		if err != nil {
			log.Fatal(err)
		}
		genItems := key + "=" + password + "\n"
		newItems, err = godotenv.Unmarshal(genItems)
		if err != nil {
			log.Fatal(err)
		}
		oldItems[key] = newItems[key]
	}

	// write keys to file
	err = godotenvgenerator.Write(oldItems, envFile)
	if err != nil {
		log.Fatal(err)
	}
}
