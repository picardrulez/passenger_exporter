package main

import (
	"log"
	"strings"
)

func int2float64(myint int) float64 {
	output := float64(myint)
	return output
}

func testArray(input XMLinfo) bool {
	returnbool := true
	name0 := getName(input.Supergroups.Supergroup[0].Name)
	name1 := getName(input.Supergroups.Supergroup[1].Name)
	name2 := getName(input.Supergroups.Supergroup[2].Name)
	name3 := getName(input.Supergroups.Supergroup[3].Name)

	if name0 != "pcms" {
		log.Println("supergroup 0 should be PCMS but value is " + name0)
		returnbool = false
	}
	if name1 != "broker" {
		log.Println("supergroup 1 should be Broker but value is " + name1)
		returnbool = false
	}
	if name2 != "benefits" {
		log.Println("supergroup 2 should be Benefits but value is " + name2)
		returnbool = false
	}
	if name3 != "admin" {
		log.Println("supergroup 3 should be Admin but value is " + name3)
		returnbool = false
	}
	return returnbool
}

func getName(input string) string {
	s := strings.Split(input, "/")
	return s[3]
}
