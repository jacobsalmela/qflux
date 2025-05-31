package entities

import "rpg-tutorial/components"

type Enemy struct {
	*Entity
	FollowsPlayer bool
	CombatComp    *components.EnemyCombat
}
