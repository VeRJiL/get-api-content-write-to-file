package main

import (
	"fmt"
	"github.com/VeRJiL/get-api-content-write-to-file/cmd/concurrent"
	"github.com/VeRJiL/get-api-content-write-to-file/cmd/normal"
	"github.com/VeRJiL/get-api-content-write-to-file/cmd/search"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Please Provide Program Type: get/search")
	}
	
	switch strings.ToLower(os.Args[1]) {
	case "get":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Please Provide a method: concurrent|normal")
		}
		
		if len(os.Args) < 4 {
			fmt.Fprintln(os.Stderr, "Please provide how many comics do you need")
		}
		
		iteration, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Please provide a Numberic Value for Iteration, got: %s", os.Args[3])
		}
		
		algo := strings.ToLower(os.Args[2])
		
		if algo == "concurrent" {
			concurrent.GetData(iteration)
			return
		}
		
		if algo == "normal" {
			normal.GetData(iteration)
			return
		}
		
		fmt.Fprintln(os.Stderr, "There are only two types of program, normal <-------> concurrent")
	
	case "search":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Pls Pass The File Name")
			os.Exit(-1)
		}
		
		if len(os.Args) < 4 {
			fmt.Fprintln(os.Stderr, "Pls Provide Search Term")
			os.Exit(-1)
		}
		
		search.Search("./outputs/"+os.Args[2]+".json", os.Args[3:])
		return
	default:
		fmt.Fprintln(os.Stderr, "The First Argument should be one of type: get | search")
	}
}
