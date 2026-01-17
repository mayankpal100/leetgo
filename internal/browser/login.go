package browser

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

func Login(ctx context.Context) error {
	fmt.Println("➡️ Please complete CAPTCHA + login manually.")
	fmt.Println("➡️ You have 2 minutes...")

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://leetcode.com/accounts/login/"),
	)

	if err != nil {
		return err
	}

	// ⏸️ WAIT for manual login
	time.Sleep(2 * time.Minute)

	fmt.Println("✅ Assuming login is complete")
	return nil
}
