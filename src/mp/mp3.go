package mp

import (
	"time"
	"fmt"
)

type MP3Player struct  {
	stat int
	progress int
}

type WAVPlayer struct {
	stat int
	progress int
}

func (p *MP3Player)Play(source string) {
	fmt.Println("Play MP3 music", source)
	
	p.progress = 0

	for p.progress < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(".")
		p.progress += 10
	}
	fmt.Println("\nFinised playing", source)
}

func (p *WAVPlayer)Play(source string) {
	fmt.Println("Play WAV music", source)
	p.progress = 0

	for p.progress <100  {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(".")
		p.progress += 10
	}
	fmt.Println("\nFinised playing", source)
}