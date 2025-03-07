package game

import "github.com/hajimehoshi/ebiten/v2"

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneMap[g.activeSceneId].Draw(screen)
}
