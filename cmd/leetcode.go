package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mayankpal100/leetgo/internal/browser"
	"github.com/mayankpal100/leetgo/internal/scraper"
)

func main() {
	parent := context.Background()

	ctx, cancel := browser.NewBrowserContext(parent)
	defer cancel()

	email := os.Getenv("LC_EMAIL")
	pass := os.Getenv("LC_PASS")

	if email == "" || pass == "" {
		fmt.Println("Please set LC_EMAIL and LC_PASS")
		return
	}

	fmt.Println("Logging in...")
	if err := browser.Login(ctx, email, pass); err != nil {
		panic(err)
	}

	fmt.Println("Logged in!")

	code, err := scraper.ScrapeSolution(ctx, "https://leetcode.com/problem-of-the-day/")
	if err != nil {
		panic(err)
	}

	fmt.Println("Fetched solution:")
	fmt.Println(code)
}
