package browser

import (
	"context"
	"encoding/json"
	"os"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

const cookieFile = "cookies.json"

// SaveCookies saves browser cookies to disk
func SaveCookies(ctx context.Context) error {
	var cookies []*network.Cookie

	err := chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			cookies, err = network.GetCookies().Do(ctx)
			return err
		}),
	)
	if err != nil {
		return err
	}

	data, _ := json.MarshalIndent(cookies, "", "  ")
	return os.WriteFile(cookieFile, data, 0644)
}

// LoadCookies loads cookies into browser
func LoadCookies(ctx context.Context) error {
	data, err := os.ReadFile(cookieFile)
	if err != nil {
		return err
	}

	var cookies []*network.Cookie
	if err := json.Unmarshal(data, &cookies); err != nil {
		return err
	}

	return chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			for _, c := range cookies {
				err := network.SetCookie(c.Name, c.Value).
					WithDomain(c.Domain).
					WithPath(c.Path).
					WithHTTPOnly(c.HTTPOnly).
					WithSecure(c.Secure).
					Do(ctx)
				if err != nil {
					return err
				}
			}
			return nil
		}),
	)
}
