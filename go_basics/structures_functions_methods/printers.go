package main

import (
	"fmt"
	"time"
)

func PrintPersonInfo(id int64, fName string, lName string, birthday time.Time) {
	fmt.Println("ID\tFirst Name\tLast Name\tAge")

	fmt.Printf("%d\t%s\t\t%s\t\t%d\n", id, fName, lName, AgeFromBirthday(birthday))
}

func PrintPersonAge(id int64, birthday time.Time) {
	fmt.Printf("id: %d age:%d\n", id, AgeFromBirthday(birthday))
}

func AgeFromBirthday(birthday time.Time) int64 {
	tNow := time.Now()
	hours := tNow.Sub(birthday).Hours()
	days := hours / 24
	return int64(days / 365)
}
