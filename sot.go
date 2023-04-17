package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func sot(file string) {
	data, err := os.ReadFile(file)
	arr := strings.Split(string(data), "\n")

	if err == nil {
		for i := 0; i < len(arr); i++ {
			_, currentDomain, _ := strings.Cut(arr[i], "://")
			currentDomain, _, _ = strings.Cut(currentDomain, "/")
			currentDomain, _, _ = strings.Cut(currentDomain, ":")

			if i < len(arr) {
				for j := 1; j < len(arr); j++ {
					if i+j < len(arr) {
						_, nextDomain, _ := strings.Cut(arr[i+j], "://")
						nextDomain, _, _ = strings.Cut(nextDomain, "/")
						nextDomain, _, _ = strings.Cut(nextDomain, ":")

						if currentDomain != nextDomain {
							arr[i+1], arr[i+j] = arr[i+j], arr[i+1]
						}
					}
				}
			}
			fmt.Println(arr[i])
		}
	} else {
		fmt.Println("Usage: sot -l [File Path]")
	}
}

func main() {
	var file string
	flag.StringVar(&file, "l", "", "List of urls")
	flag.Parse()

	sot(file)
}
