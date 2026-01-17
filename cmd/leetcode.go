package main

import (
	"context"
	"fmt"
	"github.com/mayankpal100/leetgo/internal/browser"
	"github.com/mayankpal100/leetgo/internal/scraper"
	    "github.com/chromedp/chromedp"
		"time"

)


func main() {
	ctx, cancel := browser.NewBrowserContext(context.Background())
	defer cancel()

	// 🔁 First try loading cookies
	if err := browser.LoadCookies(ctx); err != nil {
		fmt.Println("No cookies found. Need manual login.")

		// Manual login
		if err := browser.Login(ctx); err != nil {
			panic(err)
		}

		// Save session
		if err := browser.SaveCookies(ctx); err != nil {
			panic(err)
		}

		fmt.Println("🍪 Session saved!")
	}

	fmt.Println("🚀 Logged in using saved session")

fmt.Println("🚀 Logged in using saved session")

problemURL := "https://leetcode.com/problems/two-sum/"
fmt.Println("➡️ Navigating to problem:", problemURL)

if err := chromedp.Run(ctx,
	chromedp.Navigate(problemURL),
	chromedp.Sleep(6*time.Second), // allow full React hydration
); err != nil {
	panic(err)
}

code, err := scraper.FetchGoSolution(ctx)
if err != nil {
	panic(err)
}

fmt.Println("===== GO SOLUTION =====")
fmt.Println(code)


}
