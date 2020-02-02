package main

import (
	"fmt"

	img "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type scene struct {
	bg *sdl.Texture
}

func newScene(r *sdl.Renderer) (*scene, error) {
	t, err := img.LoadTexture(r, "res/img/background.png")
	if err != nil {
		return nil, fmt.Errorf("could not load background image: %v", err)
	}

	return &scene{bg: t}, nil

}

func (s *scene) paint(r *sdl.Renderer) error {
	r.Clear()

	if err := r.Copy(s, bg, nil, nil); err != nil {
		return fmt.Errorf("could not copy background: %v", err)
	}

	r.Present()
	return nil
}
