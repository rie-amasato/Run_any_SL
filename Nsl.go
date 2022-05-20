package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"unicode/utf8"
)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	signal.Ignore(syscall.SIGINT, syscall.SIGHUP)
	clearScreen()

	data, _ := os.Open("AA.txt")
	defer data.Close()
	scanner := bufio.NewScanner(data)
	lines := make([]string, 0, 100)

	mLen := 100
	//yohaku := strings.Repeat(" ",mLen)

	maxLen := 0
	for scanner.Scan() {
		lines = append(lines, (strings.Repeat(" ", mLen) + (scanner.Text())))
		if maxLen < utf8.RuneCountInString(scanner.Text()) {
			maxLen = utf8.RuneCountInString(scanner.Text())
		}
		//fmt.Println(lines)
	}
	for i := 0; i < len(lines); i += 1 {
		lines[i] += strings.Repeat(" ", maxLen) + strings.Repeat(" ", mLen)
	}

	for i := 1; i < maxLen+mLen; i += 1 {

		for _, s := range lines {
			//L := len(s)
			sUni := []rune(s)

			prline := string(sUni[i : i+mLen])

			fmt.Printf("%s\n", prline)
		}

		time.Sleep(40 * time.Millisecond)
		clearScreen()
	}

	//clearScreen()

}
