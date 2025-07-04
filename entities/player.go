package entities

import (
	"qflux/animations"
	"qflux/camera"
	"qflux/components"
)

type PlayerState uint8

const (
	Down PlayerState = iota
	Up
	Left
	Right
)

type Player struct {
	*Entity
	*camera.Camera
	Speed        float64
	LateralSpeed float64
	Health       uint
	Animations   map[PlayerState]*animations.Animation
	CombatComp   *components.BasicCombat
}

func (p *Player) ActiveAnimation(dx, dy int) *animations.Animation {
	if dx > 0 {
		return p.Animations[Right]
	}
	if dx < 0 {
		return p.Animations[Left]
	}
	if dy > 0 {
		return p.Animations[Down]
	}
	if dy < 0 {
		return p.Animations[Up]
	}
	return nil
}
