package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/sbstjn/hanu"
)

type Token struct {
	Value string `yaml:"token"`
}

var token Token

func main() {

	extractToken()

	slack, err := hanu.New(token.Value)

	if err != nil {
		log.Fatal(err)
	}

	Version := "0.0.2"

	slack.Command("version", func(conv hanu.ConversationInterface) {
		conv.Reply("I'm running `%s`", Version)
	})

	add(slack)
	slack.Listen()
}

//extractToken goes in a file ./config.yaml to get the token so as to hide it
//the config.yaml only looks like :
//token : <your_token-here>

func extractToken() {

	tokenFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Println("Couldn't open the config file")
	}
	err = yaml.Unmarshal(tokenFile, &token)
}

//add currently returns all words said by the user without the word add
//future use : get that string, add an ID to it and register it as a task in a map[int]string
func add(slack *hanu.Bot) {
	addCommand := "add"
	for i := 0; i < 10; i++ {
		addCommand += fmt.Sprintf(" <word%d>", i)
		slack.Command(addCommand, func(conv hanu.ConversationInterface) {
			answer := strings.Fields(addCommand)
			yolo := ""
			temp := ""
			for i := range answer {
				temp, _ = conv.String(fmt.Sprintf("word%d", i))
				yolo += fmt.Sprintf("%s ", temp)
			}
			conv.Reply(yolo)
		})
	}
}
