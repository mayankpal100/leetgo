package scraper

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func FetchGoSolution(ctx context.Context) (string, error) {
	var code string

	// ---------- STEP 1: Click Solutions tab ----------
	if err := chromedp.Run(ctx,
		chromedp.Evaluate(`
		(() => {
		  const tabs = document.querySelectorAll('.flexlayout__tab_button');
		  for (const tab of tabs) {
		    if (tab.querySelector('#solutions_tab')) {
		      tab.click();
		      return true;
		    }
		  }
		  return false;
		})()
		`, nil),
	); err != nil {
		return "", err
	}

	// ---------- STEP 2: Wait for solutions list (poll, max 3s) ----------
	pollCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	if err := chromedp.Run(pollCtx,
		chromedp.Poll(`
		(() => document.body.innerText.includes('Solution'))
		`, nil, chromedp.WithPollingInterval(200*time.Millisecond)),
	); err != nil {
		return "", fmt.Errorf("solutions list did not load")
	}

	// ---------- STEP 3: Scroll to force virtualization render ----------
	if err := chromedp.Run(ctx,
		chromedp.Evaluate(`
		(() => {
		  const el = document.querySelector('[data-cy="solutions-list"]');
		  if (!el) return false;
		  el.scrollTop = el.scrollHeight;
		  return true;
		})()
		`, nil),
	); err != nil {
		// non-fatal; continue
	}

	time.Sleep(500 * time.Millisecond)

	// ---------- STEP 4: Click 2nd visible solution ----------
	if err := chromedp.Run(ctx,
		chromedp.Evaluate(`
		(() => {
		  const items = [...document.querySelectorAll('[data-cy]')]
		    .filter(e => e.innerText && e.offsetParent !== null);

		  if (items.length < 2) return false;
		  items[1].click();
		  return true;
		})()
		`, nil),
	); err != nil {
		return "", fmt.Errorf("could not click solution row")
	}

	// ---------- STEP 5: Wait for solution viewer (max 3s) ----------
	viewCtx, cancel2 := context.WithTimeout(ctx, 3*time.Second)
	defer cancel2()

	if err := chromedp.Run(viewCtx,
		chromedp.Poll(`
		(() => document.querySelector('pre code') !== null)
		`, nil, chromedp.WithPollingInterval(200*time.Millisecond)),
	); err != nil {
		return "", fmt.Errorf("solution viewer did not load")
	}

	// ---------- STEP 6: Switch language to Go ----------
	_ = chromedp.Run(ctx,
		chromedp.Evaluate(`
		(() => {
		  const tabs = [...document.querySelectorAll('[role="tab"]')];
		  const goTab = tabs.find(t => t.innerText.trim() === 'Go');
		  if (!goTab) return false;
		  goTab.click();
		  return true;
		})()
		`, nil),
	)

	time.Sleep(500 * time.Millisecond)

	// ---------- STEP 7: Extract Go solution ----------
	if err := chromedp.Run(ctx,
		chromedp.Evaluate(`
		(() => {
		  const block = document.querySelector('pre code');
		  return block ? block.innerText : '';
		})()
		`, &code),
	); err != nil {
		return "", err
	}

	if strings.TrimSpace(code) == "" {
		return "", fmt.Errorf("Go solution code not found")
	}

	return code, nil
}
