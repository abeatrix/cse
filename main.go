package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Sourcegraph Support CLI"
	app.Usage = "A CLI tool for querying Sourcegraph config options."
	app.EnableBashCompletion = true

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "option",
			Value: "",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "desc",
			Aliases: []string{"d"},
			Flags:   myFlags,
			Usage:   "Looks Up Site Config description",
			Action: func(c *cli.Context) error {
				// Open schema json
				schemaFile, err := http.Get("https://raw.githubusercontent.com/sourcegraph/sourcegraph/main/schema/github.schema.json")
				if err != nil {
					fmt.Println(err)
				}

				byteValue, _ := ioutil.ReadAll(schemaFile.Body)

				var data map[string]interface{}
				json.Unmarshal([]byte(byteValue), &data)

				result := data["properties"].(map[string]interface{})[c.Args().First()].(map[string]interface{})["description"]
				//fmt.Println("option:", c.String("option"))
				fmt.Println(result)

				return nil
			},
		},
		{
			Name:    "example",
			Aliases: []string{"ex"},
			Usage:   "Looks Up Site Config Example",
			Action: func(c *cli.Context) error {

				// Open schema json
				schemaFile, err := http.Get("https://raw.githubusercontent.com/sourcegraph/sourcegraph/main/schema/github.schema.json")
				if err != nil {
					fmt.Println(err)
				}

				byteValue, _ := ioutil.ReadAll(schemaFile.Body)

				var data map[string]interface{}
				json.Unmarshal([]byte(byteValue), &data)
				result := data["properties"].(map[string]interface{})[c.Args().First()].(map[string]interface{})["example"]

				fmt.Println(result)

				return nil
			},
		},
	}

	//	start app
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
