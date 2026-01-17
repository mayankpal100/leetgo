package browser

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

// NewBrowserContext creates a chromedp browser context with logging
func NewBrowserContext(parent context.Context) (context.Context, context.CancelFunc) {
	opts := []chromedp.ContextOption{
		chromedp.WithLogf(log.Printf),
	}

	ctx, cancel := chromedp.NewContext(parent, opts...)
	return ctx, cancel
}
