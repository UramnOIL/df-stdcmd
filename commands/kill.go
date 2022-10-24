package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/entity/damage"
	"github.com/df-mc/dragonfly/server/player"
)

type KillCommand struct {
	Targets cmd.Optional[[]cmd.Target] `cmd:"target"`
}

func (c KillCommand) Run(src cmd.Source, o *cmd.Output) {
	targets, set := c.Targets.Load()

	// suicide
	if set == false {
		p, ok := src.(*player.Player)
		// not player. for example, console
		if !ok {
			o.Error("Execute commands as the player.")
			return
		}
		p.Hurt(p.MaxHealth(), damage.SourceVoid{})
		return
	}

	src.World().Entities()

	// kill other
	for _, t := range targets {
		living, ok := t.(entity.Living)	// FIXME not only Living but also Entity
		if !ok {
			o.Error("Select living targets")
			return
		}
		living.Hurt(living.MaxHealth(), damage.SourceVoid{})
	}
} 