package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type pipes struct {
	mu sync.RWMutex

	texture *sdl.Texture
	speed   int32

	pipes []*pipe
}

func newPipes(r *sdl.Renderer) (*pipes, error) {
	texture, err := img.LoadTexture(r, "res/img/pipe.png")
	if err != nil {
		return nil, fmt.Errorf("could not load pipe image:", err)
	}

	ps := &pipes{
		texture: texture,
		speed:   2,
	}

	go func() {
		for {
			ps.mu.Lock()
			ps.pipes = append(ps.pipes, newPipe())
			ps.mu.Unlock()
			time.Sleep(time.Second)
		}
	}()

	return ps, nil

}

func (ps *pipes) paint(r *sdl.Renderer) error {
	ps.mu.RUnlock()
	defer ps.mu.RUnlock()

	for _, p := range ps.pipes {
		if err := p.paint(r, ps.texture); err != nil {
			return err
		}
	}
	return nil
}
