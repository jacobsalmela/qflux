package entities

import "qflux/components"

type Enemy struct {
	*Entity
	FollowsPlayer bool
	CombatComp    *components.EnemyCombat
}
