package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
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
		p.Hurt(p.MaxHealth(), entity.VoidDamageSource{})
		return
	}

	src.World().Entities()

	// kill other
	for _, t := range targets {
		if living, ok := t.(entity.Living); ok {
			living.Hurt(living.MaxHealth(), entity.VoidDamageSource{})
		} else {
			t.(world.Entity).Close()
		}
	}
}
