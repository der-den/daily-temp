package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"syscall"
	"strconv"	
)

var removeCount int
var stayCount int

func deleteOldFiles(directory string, someDaysAgo time.Time) error {

	f, err := os.Open(directory)
    if err != nil {
        fmt.Println(err)
		return err
    }
    files, err := f.Readdir(0)
    if err != nil {
        fmt.Println(err)
		return err
    }

	fmt.Println("Cut datetime: "+someDaysAgo.String())
	
	for _, file := range files {
		
		fileInfo, err := os.Stat(directory+"/"+file.Name())				// can be file or directory
		if err != nil {
			fmt.Println(err)
			return err
		}
		
		stat := fileInfo.Sys().(*syscall.Win32FileAttributeData)
		cTimeSince := time.Unix(0, stat.CreationTime.Nanoseconds())
		
		if cTimeSince.Truncate(24*time.Hour).Before(someDaysAgo) {
			
			if file.IsDir() {
				err := deleteOldFiles(directory+"/"+file.Name(), someDaysAgo)
				if err != nil {
					return err
				}
				err = os.Remove(directory + "/" + file.Name())
				if err != nil {
					return err
				}
				fmt.Println("Deleting directory:", file.Name() + " ("+cTimeSince.String()+")")
			} else {
				
				err := os.Remove(directory + "/" + file.Name())
				if err != nil {
					return err
				}
				fmt.Println("Deleting file:", file.Name() + " ("+cTimeSince.String()+")")
			}
			//fmt.Println("Delete: " + cTimeSince.String() + " @ " + file.Name())
			removeCount = removeCount + 1
		} else if cTimeSince.Truncate(24*time.Hour).After(someDaysAgo) {
			// fmt.Println("After: " + cTimeSince.String() + " @ " + file.Name())
			stayCount = stayCount + 1
		}
		
		
	}

	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: daily-temp <days> <directory>")
		fmt.Println("The first parameter is the days that we wont to keep files. Minimum 1. As single integer. Modification time.")
		log.Fatal("Parameter error, Exit.")
	}

	keepDays := os.Args[1]

	keepDaysInt, err := strconv.ParseInt(keepDays, 10, 64)
	if err != nil {
		fmt.Println("The first parameter is the days that we wont to keep files. Minimum 1. As single integer. Modification timestamp.")
	}

	removeCount = 0
	stayCount = 0
	

	directory := os.Args[2]
	keepDaysIntNegative := -(keepDaysInt)

	someDaysAgo := time.Now().AddDate(0, 0, int(keepDaysIntNegative)).Truncate(24 * time.Hour)

	err = deleteOldFiles(directory, someDaysAgo)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Files/Directorys deleted: "+ strconv.Itoa(removeCount)+", Not touched: "+strconv.Itoa(stayCount))

}
