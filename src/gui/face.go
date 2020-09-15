package gui

import (
	"context"
	"fmt"
	"time"

	"github.com/webview/webview"
)

func Start(ctx context.Context) {

	fmt.Println("start")
	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			fmt.Println("убили")
			return
		default:
			v := ctx.Value("key").(string)
			fmt.Println(v)
			fmt.Println(ctx.Err())
			time.Sleep(time.Second)
			i++
			fmt.Println(i)
		}
	}
	debug := true
	w := webview.New(debug)
	defer w.Destroy()

	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("https://en.m.wikipedia.org/wiki/Main_Page")
	w.Run()
	fmt.Println("end")

}
