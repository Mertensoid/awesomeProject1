// guess - игра, в которой игрок должен угадать случайное число.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func guess() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)
	success := false
	for i := 10; i > 0; i-- {
		fmt.Println("You have", i, "guesses left")
		fmt.Print("Enter your number: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else {
			success = true
			fmt.Println("Congratulations! Your guess was RIGHT!")
			break
		}
	}
	if !success {
		fmt.Println("Sorry! You have no any attempts. The number was:", target)
	}
}
