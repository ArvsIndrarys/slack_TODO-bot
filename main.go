package main

import (
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

	Version := "0.0.1"

	slack.Command("shout <word>", func(conv hanu.ConversationInterface) {
		str, _ := conv.String("word")
		conv.Reply(strings.ToUpper(str))
	})

	slack.Command("whisper <word>", func(conv hanu.ConversationInterface) {
		str, _ := conv.String("word")
		conv.Reply(strings.ToLower(str))
	})

	slack.Command("version", func(conv hanu.ConversationInterface) {
		conv.Reply("Thanks for asking! I'm running `%s`", Version)
	})

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
