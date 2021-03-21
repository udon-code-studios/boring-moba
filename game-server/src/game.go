package main

import (
	"math"
	"time"
)

// Game maintains current game instance data and recieves player input data from the Hub.
type Game struct {
	GameState GameState

	PlayerInputs chan PlayerInput
}

func newGame() *Game {
	return &Game{
		PlayerInputs: make(chan PlayerInput),
	}
}

func (g *Game) run() {
	go g.updateGameState()
	go g.handlePlayerInput()
}

func (g *Game) updateGameState() {
	// loop as long as the game exists
	for {
		// update each player at a time
		for i := 0; i < len(g.GameState.Players); i++ {
			timeNow := int(time.Now().UnixNano() / 1000000)
			timeDif := timeNow - g.GameState.Players[i].LastUpdateTime                           // time difference (ms) between now and the last time player position was updated
			distDif := ((float32(g.GameState.Players[i].MoveSpeed)) / 1000.0) * float32(timeDif) // distance difference between now and the last time player position was updated
			distance := g.GameState.Players[i].CurrentPosition.distance(g.GameState.Players[i].TargetPosition)

			// if next movement will overshoot remaining distance, go to target position
			// else calculate X and Y movement differentials
			if float32(distance) <= distDif {
				g.GameState.Players[i].CurrentPosition = g.GameState.Players[i].TargetPosition
				g.GameState.Players[i].LastUpdateTime = timeNow
			} else {
				direction := g.GameState.Players[i].CurrentPosition.direction(g.GameState.Players[i].TargetPosition)
				xDif := int(float32(math.Cos(direction)) * distDif)
				yDif := int(float32(math.Sin(direction)) * distDif)

				// only update LastUpdateTime if any movement differentials were not zero
				if xDif != 0 || yDif != 0 {
					g.GameState.Players[i].CurrentPosition.X += xDif
					g.GameState.Players[i].CurrentPosition.Y += yDif
					g.GameState.Players[i].LastUpdateTime = timeNow
				}
			}
		}
	}
}

func (g *Game) handlePlayerInput() {
	for {
		select {
		case input := <-(g.PlayerInputs):
			for i := 0; i < len(g.GameState.Players); i++ {
				if g.GameState.Players[i].Id == input.Id {
					g.GameState.Players[i].TargetPosition = input.NewTargetPosition
				}
			}
		}
	}
}
