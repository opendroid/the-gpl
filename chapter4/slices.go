package chapter4

import (
	"fmt"
)

type BodyIndex int
type Bodies []string

const (
	// Sun source of energy
	Sun BodyIndex = iota // 0
	// Mercury first planet
	Mercury
	// Venus second planet
	Venus
	// Earth third planet
	Earth
	// Mars forth planet
	Mars
	// Jupiter fifth planet
	Jupiter
	// Saturn sixth planet
	Saturn
	// Uranus seventh planet
	Uranus
	// Neptune eighth planet
	Neptune
	// Pluto ninth bpdy
	Pluto // 9th body
	// HauMea 10th body
	HauMea
	// MakeMake 11th body
	MakeMake
	// GongGong 12th body
	GongGong
	// LastBody: Nothing here
	LastBody
)

var bodiesInSolarSystem = [...]string{Sun: "Sun", Mercury: "Mercury", Venus: "Venus",
	Earth: "Earth", Mars: "Mars", Jupiter: "Jupiter", Saturn: "Saturn",
	Uranus: "Uranus", Neptune: "Neptune", Pluto: "Pluto", HauMea: "HauMea",
	MakeMake: "MakeMake", GongGong: "GongGong"}

// InnerPlanets slice of inner planets
func InnerPlanets() []string {
	return bodiesInSolarSystem[Mercury:Jupiter] // Jupiter not included
}

// GasGiants slices of gas giants
func GasGiants() Bodies {
	return bodiesInSolarSystem[Jupiter:Pluto] // Pluto not included
}

// TNO returns Trans-Neptune Objects
func TNO() Bodies {
	return bodiesInSolarSystem[Pluto:LastBody]
}

// printLenAndCap test function
func PrintLenAndCaOfAllBodies() {
	inner := InnerPlanets()
	gg := GasGiants()
	tno := TNO()
	fmt.Printf("(Len, Cap) Inner: (%d, %d), (%d, %d), (%d, %d)",
		len(inner), cap(inner), len(gg), cap(gg), len(tno), cap(tno))
}
