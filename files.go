package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func createStorage() error {
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		err := os.Mkdir("data", 0777)
		if err != nil {
			return err
		}
	}
	return nil
}

func fileAppend(file string, length int, data string) error {
	f, err := os.OpenFile("data/"+file+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}(f)

	_, err = f.WriteString(fmt.Sprintf("%d: %s\n", length, data))
	if err != nil {
		return err
	}
	return nil
}

func readFileTask(file string) ([]string, error) {
	f, err := os.OpenFile("data/"+file+".txt", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}(f)

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}

	var tasks []string

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.SplitN(line, ": ", 2)
		tasks = append(tasks, parts[1])
	}
	return tasks, nil
}

func removeFileLine(file string, position int) error {
	f, err := os.OpenFile("data/"+file+".txt", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}(f)

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}

	lines = append(lines[:position], lines[position+1:]...)

	for i := position; i < len(lines); i++ {
		lines[i] = strconv.Itoa(i+1) + lines[i][1:]
	}

	joined := strings.Join(lines, "\n")
	joined += "\n"
	return os.WriteFile("data/"+file+".txt", []byte(joined), 0644)
}

func updateFileLine(file string, position int, data string) error {
	f, err := os.OpenFile("data/"+file+".txt", os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}(f)

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() != "\n" {
			lines = append(lines, scanner.Text())
		}
	}

	num := int(lines[position][0]) - 48
	lines[position] = fmt.Sprintf("%d: %s", num, data)
	joined := strings.Join(lines, "\n")
	joined += "\n"

	return os.WriteFile("data/"+file+".txt", []byte(joined), 0644)
}
