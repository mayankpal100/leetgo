package browser

import (
    "context"
    "github.com/chromedp/chromedp"
	"time"
)

func Login(ctx context.Context, email, pass string) error {
    tasks := chromedp.Tasks{
        chromedp.Navigate("https://leetcode.com/accounts/login/"),
        chromedp.WaitVisible(`input[name="login"]`),
        chromedp.SendKeys(`input[name="login"]`, email),
        chromedp.SendKeys(`input[name="password"]`, pass),
        chromedp.Click(`button[type="submit"]`),
        chromedp.Sleep(5 * time.Second),
    }
    return chromedp.Run(ctx, tasks)
}
