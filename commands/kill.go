package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/entity/damage"
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
		p := src.(*player.Player)
		p.SetGameMode(world.GameModeSurvival)
		p.Hurt(p.MaxHealth(), damage.SourceVoid{})
		return
	}

	src.World().Entities()

	// kill other
	for _, target := range targets {
		t := target.(entity.Living)
		t.Hurt(t.MaxHealth(), damage.SourceVoid{})
	}
} 