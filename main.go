package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	total, values := readFile()
	pos := compute(total, values)
	writeToFile(pos)
}

func readFile() (totalCount int, values []int) {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for index, line := range lines {
		tt := strings.Split(line, " ")
		if index == 0 {
			totalCount, err = strconv.Atoi(tt[0])
			if err != nil {
				log.Fatalln(err)
			}
			continue
		}

		values = make([]int, len(tt))
		for i := range values {
			values[i], err = strconv.Atoi(tt[i])
			if err != nil {
				log.Fatalln("Could not convert value", err)
			}
		}
	}

	return
}

func compute(total int, values []int) (pos []int) {
	totalComputed := values[len(values)-1]
	pos = append(pos, len(values)-1)

	for index, value := range values {
		totalComputed += value
		if totalComputed >= total || index == len(values) {
			return
		}
		pos = append(pos, index)
	}
	return
}

func writeToFile(pos []int) {
	file, err := os.Create("soln_" + os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	_, err = file.WriteString(fmt.Sprintf("%d\n", len(pos)))
	if err != nil {
		log.Fatalln(err)
	}

	for _, value := range pos {
		if value == len(pos)-1 {
			_, err = file.WriteString(fmt.Sprintf("%d", value))
		} else {
			_, err = file.WriteString(fmt.Sprintf("%d ", value))

		}
		if err != nil {
			log.Fatalln(err)
		}
	}

	err = file.Sync()
	if err != nil {
		log.Fatalln(err)
	}
}
