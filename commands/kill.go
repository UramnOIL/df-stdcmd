package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type SuicideCommand struct {
}

func (c SuicideCommand) Run(src cmd.Source, o *cmd.Output) {

	p, ok := src.(*player.Player)
	// not player. for example, console
	if !ok {
		o.Error("Execute commands as the player.")
		return
	}
	p.Hurt(p.MaxHealth(), entity.VoidDamageSource{})
	return
}

type KillCommand struct {
	Targets []cmd.Target `cmd:"target"`
}

func (c KillCommand) Run(src cmd.Source, o *cmd.Output) {

	src.World().Entities()

	// kill other
	for _, t := range c.Targets {
		if living, ok := t.(entity.Living); ok {
			living.Hurt(living.MaxHealth(), entity.VoidDamageSource{})
		} else {
			t.(world.Entity).Close()
		}
	}
}
