package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer file.Close()
	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n')
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()
	_, err = fmt.Fprintln(file, line)
	return err
}

const filename string = "data.txt"

func main() {
	line, err := ReadFile(filename)
	if err != nil {
		err = WriteFile(filename, "Hello, World!")
		if err != nil {
			fmt.Println("Failed to write file:", err)
			return
		}
		line, err := ReadFile(filename)
		if err != nil {
			fmt.Println("Failed to read file:", err)
			return
		}
		fmt.Println("Read file:", line)
	}
	fmt.Println("Read file:", line)
}
