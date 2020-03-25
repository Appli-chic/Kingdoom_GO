package models

import (
	"github.com/veandco/go-sdl2/sdl"
)

type CharacterInfo struct {
	Key      int
	ImageKey int
	Src      *sdl.Rect
	Width    int32
	Height   int32
}