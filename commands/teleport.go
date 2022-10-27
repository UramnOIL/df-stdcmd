package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type TeleportToTargetCommand struct {
	Destination []cmd.Target `cmd:"destination"`
}

func (c TeleportToTargetCommand) Run(src cmd.Source, o *cmd.Output) {
	p, ok := src.(*player.Player)
	if !ok {
		o.Error("Execute command as player")
		return
	}

	if len(c.Destination) != 1 {
		o.Error("Enter a correct destination")
	}

	p.Teleport(c.Destination[0].Position())
}

type TeleportToCoordinateCommand struct {
	Destination mgl64.Vec3 `cmd:"destination"`
}

func (c TeleportToCoordinateCommand) Run(src cmd.Source, o *cmd.Output) {
	p, ok := src.(*player.Player)
	if !ok {
		o.Error("Execute command as player")
		return
	}

	p.Teleport(c.Destination)
}

type TeleportVictimToTargetCommand struct {
	Victim      []cmd.Target `cmd:"victim"`
	Destination []cmd.Target `cmd:"destination"`
}

func (c TeleportVictimToTargetCommand) Run(src cmd.Source, o *cmd.Output) {
	if len(c.Destination) != 1 {
		o.Error("Enter a correct destination")
		return
	}

	for _, v := range c.Victim {
		if p, ok := v.(*player.Player); ok {
			p.Teleport(c.Destination[0].Position())
		} else {
			e := v.(world.Entity)
			e.EncodeEntity()
		}
	}
}

type TeleportVictimToCoordinateCommand struct {
	Allower     cmd.Allower
	Victim      []cmd.Target `cmd:"victim"`
	Destination mgl64.Vec3   `cmd:"destination"`
}

func (c TeleportVictimToCoordinateCommand) Run(src cmd.Source, o *cmd.Output) {
	for _, v := range c.Victim {
		if p, ok := v.(*player.Player); ok {
			p.Teleport(c.Destination)
		} else {
			e := v.(world.Entity)
			e.EncodeEntity()
		}
	}
}
