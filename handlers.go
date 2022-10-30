package main

import (
	"net/http"
	"log"
	"strconv"
	"fmt"
)

func rollerHandler(w http.ResponseWriter, req *http.Request) {
	dStr := req.Header.Get("diceValue")
	dVal, err := strconv.Atoi(dStr);
	if err != nil {
		log.Panic("Error retrieving int from string in call Atoi()")
	}

	dRoll, err := roll(dVal)
	if err != nil {
		log.Panic(err);
	}
	fmt.Printf("You rolled: %v\n", dRoll)
}

func rollerMinHandler(w http.ResponseWriter, req *http.Request) {
	dStr := req.Header.Get("diceValue")
	dMin := req.Header.Get("minValue")
	dVal, err := strconv.Atoi(dStr);
	if err != nil {
		log.Panic("Error retrieving int from string in call Atoi()")
	}
	min, err := strconv.Atoi(dMin);
	if err != nil {
		log.Panic("Error retrieving int from string in call Atoi()")
	}

	dRoll, err := minRoll(dVal, min)
	if err != nil {
		log.Panic(err);
	}
	fmt.Printf("You rolled: %v\n", dRoll)
}

func rollerAddHandler(w http.ResponseWriter, req *http.Request) {
	dStr := req.Header.Get("diceValue")
	dAdd := req.Header.Get("add")
	dVal, err := strconv.Atoi(dStr);
	if err != nil {
		log.Panic("Error retrieving int from string in call Atoi()")
	}
	add, err := strconv.Atoi(dAdd);
	if err != nil {
		log.Panic("Error retrieving int from string in call Atoi()")
	}

	dRoll, err := addRoll(dVal, add)
	if err != nil {
		log.Panic(err);
	}
	fmt.Printf("You rolled: %v\n", dRoll)
}