package commands

import (
	"fmt"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type GameMode string
func (GameMode) Type() string { return "GameMode" }
func (GameMode) Options(src cmd.Source) []string {
	return []string{"survival", "s", "creative", "c", "adventure", "a", "supectator"}
}

func stog(g GameMode) (gameMode world.GameMode, err error) {
	gameMode, err = nil, &GameModeArgumentError{g}
	switch g {
	case "survival", "s":
		gameMode, err = world.GameModeSurvival, nil
	case "creative", "c":
		gameMode, err = world.GameModeCreative, nil
	case "adventure", "a":
		gameMode, err = world.GameModeAdventure, nil
	case "spectator":
		gameMode, err = world.GameModeSpectator, nil
	}
	return
}

func itog(g int) (gameMode world.GameMode, err error) {
	gameMode, err = nil, &GameModeArgumentError{g}
	switch g {
	case 0:
		gameMode, err = world.GameModeSurvival, nil
	case 1:
		gameMode, err = world.GameModeCreative, nil
	case 2:
		gameMode, err = world.GameModeAdventure, nil
	case 3:
		gameMode, err = world.GameModeSpectator, nil
	}
	return
}

type SetMyGameModeFromStringCommand struct {
	GameMode GameMode `cmd:"gameMode"`
}

func (c SetMyGameModeFromStringCommand) Run(src cmd.Source, o *cmd.Output) {

	p, ok := src.(*player.Player)

	if !ok {
		o.Error("Execute as a player")
		return
	}

	g, err := stog(c.GameMode)

	if err != nil {
		o.Error("Enter correct gamemode value")
		return
	}

	p.SetGameMode(g)
}

type SetMyGameModeFromIntCommand struct {
	GameMode int `cmd:"gameMode"`
}

func (c SetMyGameModeFromIntCommand) Run(src cmd.Source, o *cmd.Output) {

	p, ok := src.(*player.Player)

	if !ok {
		o.Error("Execute as a player")
		return
	}

	g, err := itog(c.GameMode)

	if err != nil {
		o.Error("Enter correct gamemode value")
		return
	}

	p.SetGameMode(g)
}

type SetTargetGameModeFromStringCommand struct {
	GameMode GameMode     `cmd:"gameMode"`
	Target   []cmd.Target `cmd:"target"`
}

func (c SetTargetGameModeFromStringCommand) Run(src cmd.Source, o *cmd.Output) {
	g, err := stog(c.GameMode)

	if err != nil {
		o.Error("Enter correct gamemode value")
		o.Error(err)
		return
	}

	var players []*player.Player

	for _, t := range c.Target {
		p, ok := t.(*player.Player)

		if !ok {
			o.Error("Enter correct targets")
			return
		}

		players = append(players, p)
	}

	for _, p := range players {
		p.SetGameMode(g)
	}
}

type SetTargetGameModeFromIntCommand struct {
	GameMode int          `cmd:"gameMode"`
	Target   []cmd.Target `cmd:"target"`
}

func (c SetTargetGameModeFromIntCommand) Run(src cmd.Source, o *cmd.Output) {
	g, err := itog(c.GameMode)

	if err != nil {
		o.Error("Enter correct gamemode value")
		return
	}

	var players []*player.Player

	for _, t := range c.Target {
		p, ok := t.(*player.Player)

		if !ok {
			o.Error("Enter correct targets")
			return
		}

		players = append(players, p)
	}

	for _, p := range players {
		p.SetGameMode(g)
	}
}

type GameModeArgumentError struct {
	arg interface{}
}

func (e *GameModeArgumentError) Error() string {
	return fmt.Sprintf("%s is illegal value", e.arg)
}
