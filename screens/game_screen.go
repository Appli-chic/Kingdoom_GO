package screens

import (
	"github.com/kingdoom/managers"
	"github.com/veandco/go-sdl2/sdl"
)

type GameScreen struct {
	window          *sdl.Window
	renderer        *sdl.Renderer
	resourceManager managers.ResourceManager
	world           World
}

func NewGameScreen(window *sdl.Window, renderer *sdl.Renderer) *GameScreen {
	resourceManager := managers.NewResourceManager()
	g := &GameScreen{window, renderer, resourceManager, NewWorld(&resourceManager, renderer, 50, 50)}
	return g
}

func (g *GameScreen) HandleEvents() bool {
	running := true

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			println("Quit")
			running = false
			break
		}
	}

	g.world.HandleEvents()

	return running
}

func (g *GameScreen) Update() {
	g.world.Update()
}

func (g *GameScreen) Render() {
	width, height := g.window.GetSize()
	g.renderer.Clear()
	g.renderer.SetDrawColor(255, 0, 0, 255)
	g.renderer.FillRect(&sdl.Rect{W: width, H: height})

	g.world.Render()

	g.renderer.Present()
}