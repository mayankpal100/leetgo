package scraper

import (
    "context"
    "github.com/chromedp/chromedp"
    "strings"
	"time"
	"fmt"
)

func ScrapeSolution(ctx context.Context, url string) (string, error) {
    var html string

    err := chromedp.Run(ctx,
        chromedp.Navigate(url),
        chromedp.Sleep(3*time.Second),
        chromedp.OuterHTML(`div[data-cy="question-content"]`, &html),
    )
    if err != nil {
        return "", err
    }

    // Extract the code block
    // (Simplest approach: find first <pre><code>…</code></pre>)
    start := strings.Index(html, "<pre")
    end   := strings.Index(html, "</pre>")
    if start == -1 || end == -1 {
        return "", fmt.Errorf("no solution block found")
    }
    snippet := html[start:end]
    return snippet, nil
}
