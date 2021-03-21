package main

import (
	"math"
)

type Location struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (loc1 Location) distance(loc2 Location) int {
	return int(math.Sqrt(math.Pow(float64(loc2.X-loc1.X), 2) + math.Pow(float64(loc2.Y-loc1.Y), 2)))
}

func (loc1 Location) direction(loc2 Location) float64 {
	return math.Atan(float64(loc2.Y-loc1.Y) / float64(loc2.X-loc1.X))
}

type Player struct {
	Id              int      `json:"id"`
	DisplayName     string   `json:"displayName"`
	LastUpdateTime  int      `json:"lastUpdateTime"` // Unix time (ms) the last time player position was updated
	CurrentPosition Location `json:"currentPosition"`
	TargetPosition  Location `json:"targetPosition"`
	MoveSpeed       int      `json:"moveSpeed"` // units per second
}

type GameState struct {
	Players []Player `json:"players"`
}

type PlayerInput struct {
	Id                int      `json:"id"`
	NewTargetPosition Location `json:"newTargetPosition"`
}
