package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main(){


	invertCase("Lorem Ipsum Dolor Sit Amet")

	countLetters("Team Engineering PT. Raksasa Laju Lintang")
	
	wordStatistics(`Go​, also known as ​Golang​,[​ 14]​ is a ​statically typed​, ​compiled​ ​programming language​ designed at Google[​ 15]​ by Robert Griesemer, ​Rob Pike​, and ​Ken Thompson​.[​ 12]​ Go is ​syntactically​ similar to ​C​, but with ​memory safety​, ​garbage collection​, ​structural typing​,[​ 6]​ and ​CSP​-style ​concurrency​.[​ 16],
	Go was designed at ​Google​ in 2007 to improve ​programming productivity​ in an era of ​multicore​, networked​ ​machines​ and large ​codebases​.[​ 23]​ The designers wanted to address criticism of other languages in use at ​Google​, but keep their useful characteristics:[​ 24],
	● Static typing​ and ​run-time​ efficiency (like ​C++​)
	● Readability​ and ​usability​ (like ​Python​ or ​JavaScript​)[​ 25]
	● High-performance ​networking​ and ​multiprocessing
	The designers were primarily motivated by their shared ​dislike of C++​.[​ 26]​[27]​[28]
	Go was publicly announced in November 2009,[​ 29]​ and version 1.0 was released in March 2012.[​ 30]​[31]​ Go is widely used in production at Google[​ 32]​ and in many other organizations and
	open-source projects.
	In November 2016, the Go and Go Mono fonts which are ​sans-serif​ and ​monospaced​ respectively were released by type designers ​Charles Bigelow​ and ​Kris Holmes​. Both fonts adhere to ​WGL4​ and were designed to be legible with a large x-height and distinct letterforms by conforming to the DIN
	1450 standard.[​ 33]​[34]
	In April 2018, the original logo was replaced with a stylized GO slanting right with trailing
	streamlines. However, the ​Gopher​ ​mascot​ remained the same.[​ 35]
	In August 2018, the Go principal contributors published two ′′draft designs′′ for new language
	features, ​Generics​ and ​error handling​, and asked Go users to submit feedback on them.[​ 36]​[37]​ Lack of support for generic programming and the verbosity of error handling in Go 1.x had drawn considerable ​criticism​.`)
	
}

func countLetters(text string){
	data := strings.ReplaceAll(text, " ", "")
	
	var checked map[string]string = make(map[string]string)

	var output string = ""
	for _, v := range data {
		s := string(v)
		
		_, upperExists := checked[strings.ToUpper(s)]
		_, lowerExists := checked[strings.ToLower(s)]
		if upperExists || lowerExists || s == "."{
			continue
		} else {
			checked[s] = s
		}
		c := strings.Count(data, strings.ToUpper(s))+ strings.Count(data, strings.ToLower(s))
		count := ""
		if c != 1 {
			count = strconv.Itoa(c)
		}

		output = output + s + count
	}

	fmt.Println(output)
}

func invertCase(text string){

	var output string = ""

	for _, v := range text {
		s := string(v)
		if s == strings.ToUpper(s){
			output = output + strings.ToLower(s)
		} else {
			output = output + strings.ToUpper(s)
		}
	}

	fmt.Println(output)

}

func wordStatistics(text string){

	var totalCount int = 0
	var highestValue int = 0
	var highestValueWord string = ""

	var lowestValue int = 99
	var lowestValueWord string = ""

	words := strings.Split(text, " ")

	var checkedWords map[string]int = make(map[string]int)

	fmt.Println("======================================")
	fmt.Println("Word count for every word")

	for _, v := range words {
		v = strings.TrimSpace(v)
		totalCount++
		_, exists := checkedWords[v]

		if exists {
			continue
		}

		checkedWords[v] = strings.Count(text, v)

		if checkedWords[v] > highestValue {
			highestValue = checkedWords[v]
			highestValueWord = v
		} 

		if checkedWords[v] < lowestValue {
			lowestValue = checkedWords[v]
			lowestValueWord = v
		} 
		
		fmt.Println(v, ": ", checkedWords[v])
	}
	
	fmt.Println("======================================")
	fmt.Printf("Total word count is: %v\n", totalCount)
	fmt.Println("======================================")
	fmt.Println("Highest count word is:", highestValueWord, "with count:", highestValue)
	fmt.Println("======================================")
	fmt.Println("Lowest count word is:", lowestValueWord, " with count:", lowestValue)
	fmt.Println("======================================")


}

