package main

import (
	"progressbar/progressbar"
	"time"
)

func main() {
	var bar progressbar.Bar
	bar.NewOption(0, 100)
	for i := 0; i <= 100; i++ {
		time.Sleep(100 * time.Millisecond)
		bar.Play(int64(i))
	}
	bar.Finish()
}
