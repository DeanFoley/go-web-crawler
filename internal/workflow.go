package app

import (
	"fmt"
	"os"
	"time"
)

func StartWorkflow(uri string) {
	if _, err := ValidateUrl(uri); err != nil {
		fmt.Println("Please provide a valid URL.")
		return
	}
	webpage, err := GrabWebpage(uri)
	if err != nil {
		fmt.Println(err)
		return
	}
	anchors, err := ExtractAnchors(webpage)
	if err != nil {
		fmt.Println(err)
		return
	}
	anchorBytes := FormatAnchors(anchors)
	workingDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	filename := fmt.Sprintf("%s/anchors-%s.txt", workingDirectory, time.Now().Format(time.RFC3339))
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.Write(anchorBytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Anchors printed to %s! Thank you for using my cool tool!\n", filename)
}
