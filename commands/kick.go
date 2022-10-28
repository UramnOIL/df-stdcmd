package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type KickCommand struct {
	Target []cmd.Target `cmd:"target"`
	Reason cmd.Optional[string] `cmd:"reason"`
}

func (c KickCommand) Run(src cmd.Source, o *cmd.Output) {
	var players []*player.Player

	for _, t := range c.Target {
		p, ok := t.(*player.Player)
		if !ok {
			o.Error("Enter correct targets")
		}
		players = append(players, p)
	}

	reason := "Kicked by an admin"

	if r, ok := c.Reason.Load(); ok {
		reason = r
	}

	for _, p := range players {
		p.Disconnect(reason)
	}
}