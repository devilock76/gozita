package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"log"
	"encoding/xml"
)

type Card struct {
	HWName		string
	Name		string
	Rates		[]int
	Depths		[]int
}

type JackSession struct {
	Card		string
	Rate		int
	Depth		int
}

var CurrentCards []Card

// Get List of currently available cards
func getCardList () []string {
	cards, err := ioutil.ReadFile("/proc/asound/cards")
	if err != nil {
		log.Fatal(err)
	}
	return cards
}

// Get current jack session information
func getJackSession () JackSession {
	// for now read the conf file
	jackStr, err := ioutil.ReadFile("~/.config/jack/conf.xml")
	if err != nil {
		log.Fatal(err)
	}
}

// Get config file data
func getConfigData () {
	configStr, err := ioutil.ReadFile("~/.config/gozita/default.conf")
	if err != nil {
		log.Fatal(err)
	}
}

// Unites all the other functions to a usable current world view
func buildCardList (jack JackSession) {
	cardList := getCardList()
}

func main () {
	curSession := getJackSession()
	
}	
