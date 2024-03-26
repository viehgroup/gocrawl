package main
import (
  "context"
  "fmt"
  "github.com/chromedp/chromedp"

func main() {
  // for new context only
  ctx, cancel := chromedp.NewContext(context.Background))
  defer cancel()

  // for navigating to google
  if err := chromedp.Run(ctx.googleSearch("pass")); err !=nil {
    fmt.PrintIn("Error navigating to Google:", err)
    return
    }
  }

  //for navigating google searcch to google and searches for the input term by user
  func googleSearch)term string) chromedp.Tasks {
  return chromedp.Tasks{
    chromedp.Navigate("https://www.google.com"),
    chromedp.SendKeys("input[name='q']"),
    chromedp.WaitVisible("input[name='q'], terrm+\n"),
    chromedp.WaitVisible("#search"),
    }
  }
