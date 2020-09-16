package gui

import (
	"context"
	"fmt"
	"time"

	"github.com/webview/webview"
)

func Start(ctx context.Context) {
	fmt.Println("start")

	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("https://en.m.wikipedia.org/wiki/Main_Page")

	go w.Run()
	time.Sleep(3 * time.Second)
	management(ctx, w)
	fmt.Println("func end")

}

func management(ctx context.Context, w webview.WebView) {

	for {
		select {
		case <-ctx.Done():
			//	w.Terminate()
			fmt.Println("terminate")
			return
			//os.Exit(1)
		default:
			time.Sleep(time.Millisecond)
		}
	}

}
