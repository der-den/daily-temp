package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

func deleteOldFiles(directory string, someDaysAgo time.Time) error {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.ModTime().Truncate(24 * time.Hour).Before(someDaysAgo) {
			if file.IsDir() {
				err := deleteOldFiles(directory+"/"+file.Name(), someDaysAgo)
				if err != nil {
					return err
				}
				err = os.Remove(directory + "/" + file.Name())
				if err != nil {
					return err
				}
				fmt.Println("Deleting directory:", file.Name())
			} else {
				err := os.Remove(directory + "/" + file.Name())
				if err != nil {
					return err
				}
				fmt.Println("Deleting file:", file.Name())
			}
		}
	}

	return nil
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: daily-temp <days> <directory>")
	}
	keepDays := os.Args[1]

	keepDaysInt, err := strconv.ParseInt(keepDays, 10, 64)
	if err != nil {
		fmt.Println("The first parameter is the days that we wont to keep files. Minimum 1. As single integer.")
	}

	directory := os.Args[2]
	keepDaysIntNegative := -(keepDaysInt)

	someDaysAgo := time.Now().AddDate(0, 0, int(keepDaysIntNegative)).Truncate(24 * time.Hour)

	err = deleteOldFiles(directory, someDaysAgo)
	if err != nil {
		log.Fatal(err)
	}
}
