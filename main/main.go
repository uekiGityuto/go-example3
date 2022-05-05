package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {
	ch := input(os.Stdin)
	questions := []string{"Gorilla", "Monkey", "Human", "GiantGorilla", "Orangutan"}
	rand.Seed(time.Now().UnixNano())
	limit := time.After(30 * time.Second)
	var correctNum int
	for {
		question := questions[rand.Intn(len(questions))]
		fmt.Println(">" + question)
		select {
		case <-limit:
			fmt.Println("終了時間です。")
			fmt.Printf("あなたの正解数は、%dです。", correctNum)
			return
		case word := <-ch:
			if question == word {
				correctNum++
			}
		}
	}
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}
