package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {
	bc := context.Background()
	t := 30 * time.Second
	ctx, cancel := context.WithTimeout(bc, t)
	defer cancel()
	ch := input(os.Stdin)
	questions := []string{"Gorilla", "Monkey", "Human", "GiantGorilla", "Orangutan"}
	rand.Seed(time.Now().UnixNano())
	var correctNum int
	for {
		question := questions[rand.Intn(len(questions))]
		fmt.Println(">" + question)
		select {
		case <-ctx.Done():
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
