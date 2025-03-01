package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

type PokemonStat string

const (
	Physical PokemonStat = "Physical"
	Special  PokemonStat = "Special"
)

type PokemonType string

const (
	Fire   PokemonType = "Fire"
	Water  PokemonType = "Water"
	Grass  PokemonType = "Grass"
	Normal PokemonType = "Normal"
)

var typeEffectiveness = map[PokemonType]map[PokemonType]float64{
	Fire: {
		Grass:  2.0,
		Water:  0.5,
		Fire:   0.5,
		Normal: 1.0,
	},
	Grass: {
		Grass:  0.5,
		Water:  2.0,
		Fire:   0.5,
		Normal: 1.0,
	},
	Water: {
		Grass:  0.5,
		Water:  0.5,
		Fire:   2.0,
		Normal: 1.0,
	},
	Normal: {
		Grass:  1.0,
		Water:  1.0,
		Fire:   1.0,
		Normal: 1.0,
	},
}

type Pokemon struct {
	Name      string
	Attack    int
	Defense   int
	Health    int
	Protected bool
	Type      PokemonType
	Move      [3]customMove
	MegaEvo   bool
}

type customMove struct {
	Name     string
	Attack   int
	Type     PokemonType
	MoveStat PokemonStat
}

func (p *Pokemon) AttackPokemon(target *Pokemon, move customMove) {
	var STAB float64
	customRand := 0.85 + (r.Float64() * 0.15)
	multiplier := typeEffectiveness[move.Type][target.Type]
	if move.Type == p.Type {
		STAB = 1.5
	} else {
		STAB = 1.0
	}

	attack := float64(p.Attack * move.Attack)
	defense := float64(target.Defense)
	if defense == 0 {
		defense = 1.0
	}
	intermediate := (((2*50/5 + 2) * attack / defense) / 50) + 2
	finaldamage := intermediate * STAB * multiplier * customRand
	if math.IsInf(finaldamage, 0) {
		fmt.Println("Warning: Infinite damage detected! Fixing now...")
		finaldamage = 50
	}
	damage := int(math.Round(finaldamage))
	if damage < 1 {
		damage = 1
	}
	target.Health -= damage
	if multiplier > 1.0 {
		fmt.Println("It was super effective!")
	} else if multiplier < 1.0 {
		fmt.Println("It wasn't very effective...")
	}

	// 🔍 Print Debug Info
	fmt.Printf("\n%s used %s!\n", p.Name, move.Name)
	fmt.Printf("Move Type: %s | Target Type: %s\n", move.Type, target.Type)
	fmt.Printf("Attack Stat: %d | Move Power: %d | Target Defense: %d\n", p.Attack, move.Attack, target.Defense)
	fmt.Printf("STAB: %.2f | Type Effectiveness: %.2f | Random Factor: %.2f\n", STAB, multiplier, customRand)
	fmt.Printf("Base Damage: %.2f | Final Damage Before Rounding: %.2f\n", intermediate, finaldamage)
	fmt.Printf("Final Damage Applied: %d\n", damage)

	fmt.Printf("It did %d damage.\n", damage)
	if target.Health <= 0 {
		time.Sleep(1 * time.Second)
		fmt.Printf("%s has fainted!\n", target.Name)
	}
}

func (p *Pokemon) Protect() bool {
	if p.Protected {
		p.Protected = false
		return true
	}
	p.Protected = true
	return false
}

func battle(player, opponent *Pokemon) {
	fmt.Println("Battle Start!")
	fmt.Scanln()
	for player.Health > 0 && opponent.Health > 0 {
		fmt.Printf("\n%s (HP: %d) vs %s (HP: %d)\n", player.Name, player.Health, opponent.Name, opponent.Health)
		fmt.Printf("1. %s\n2. %s\n3. %s\n4. Protect\n", player.Move[0].Name, player.Move[1].Name, player.Move[2].Name)
		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid Input! Please enter a number between 1-4.")
			continue
		}
		if choice < 1 || choice > 4 {
			fmt.Println("Invalid choice! Pick a number between 1-4.")
			continue
		}
		switch choice {
		case 1:
			if !opponent.Protected {
				fmt.Printf("\n%s used %s!\n", player.Name, player.Move[0].Name)
				player.AttackPokemon(opponent, player.Move[0])
			} else {
				fmt.Printf("\n%s used Protect!\n", opponent.Name)
			}
			opponent.Protected = false

		case 2:
			if !opponent.Protected {
				fmt.Printf("\n%s used %s!\n", player.Name, player.Move[1].Name)
				player.AttackPokemon(opponent, player.Move[1])
			} else {
				fmt.Printf("\n%s used Protect!\n", opponent.Name)
			}
			opponent.Protected = false
		case 3:
			if !opponent.Protected {
				fmt.Printf("\n%s used %s!\n", player.Name, player.Move[2].Name)
				player.AttackPokemon(opponent, player.Move[2])
			} else {
				fmt.Printf("\n%s used Protect!\n", opponent.Name)
			}
			opponent.Protected = false

		case 4:
			player.Protected = true
			fmt.Println("The next attack will be nullified!")
		default:
			fmt.Println("Error! please pick an action.")
			continue
		}

		if opponent.Health <= 0 {
			fmt.Println("You won!")
			break
		}

		fmt.Println("Enemy is choosing...")
		time.Sleep(2 * time.Second)
		myrand := r.Float64()
		if myrand >= 0.25 {
			if !player.Protected {
				myrand2 := r.Intn(3)
				if myrand2 == 0 {
					fmt.Printf("\n%s used %s!\n", opponent.Name, opponent.Move[0].Name)
					opponent.AttackPokemon(player, opponent.Move[0])
				} else if myrand2 == 1 {
					fmt.Printf("\n%s used %s!\n", opponent.Name, opponent.Move[1].Name)
					opponent.AttackPokemon(player, opponent.Move[1])
				} else if myrand2 == 2 {
					fmt.Printf("\n%s used %s!\n", opponent.Name, opponent.Move[2].Name)
					opponent.AttackPokemon(player, opponent.Move[2])
				}

			} else {
				fmt.Printf("\n%s used Protect!", player.Name)
			}
			player.Protected = false

		} else if myrand < 0.25 {
			opponent.Protected = true
			fmt.Println("The next attack will be nullified!")
		}

		if player.Health <= 0 {
			fmt.Println("You lost the battle!")
			break
		}
	}
}

func main() {
	bulbasaur := Pokemon{
		Name:    "Bulbasaur",
		Attack:  35,
		Defense: 25,
		Health:  110,
		Move: [3]customMove{
			{Name: "Tackle", Type: Normal, MoveStat: Physical, Attack: 60},
			{Name: "Growl", Type: Normal, MoveStat: Physical, Attack: 50},
			{Name: "Vine Whip", Type: Grass, MoveStat: Physical, Attack: 60},
		},
		MegaEvo: false,
		Type:    Grass,
	}
	squirtle := Pokemon{
		Name:    "Squirtle",
		Attack:  35,
		Defense: 35,
		Health:  100,
		Move: [3]customMove{
			{Name: "Tackle", Type: Normal, MoveStat: Physical, Attack: 60},
			{Name: "Tail Whip", Type: Normal, MoveStat: Physical, Attack: 50},
			{Name: "Water Gun", Type: Water, MoveStat: Special, Attack: 60},
		},
		MegaEvo: false,
		Type:    Water,
	}
	charmander := Pokemon{
		Name:    "Charmander",
		Attack:  45,
		Defense: 20,
		Health:  100,
		Move: [3]customMove{
			{Name: "Scratch", Type: Normal, MoveStat: Physical, Attack: 60},
			{Name: "Growl", Type: Normal, MoveStat: Physical, Attack: 50},
			{Name: "Ember", Type: Fire, MoveStat: Special, Attack: 60},
		},
		MegaEvo: false,
		Type:    Fire,
	}
	var mypokemon int
	var opp int
	fmt.Println("Choose Your Pokemon:")
	fmt.Printf("1. %v\n", bulbasaur.Name)
	fmt.Printf("2. %v\n", charmander.Name)
	fmt.Printf("3. %v\n", squirtle.Name)
	fmt.Scan(&mypokemon)
	fmt.Println("Choose Opposing Pokemon:")
	fmt.Printf("1. %v\n", bulbasaur.Name)
	fmt.Printf("2. %v\n", charmander.Name)
	fmt.Printf("3. %v\n", squirtle.Name)
	fmt.Scan(&opp)
	if mypokemon == opp {
		fmt.Println("Cannot choose same pokemon!")
	} else if mypokemon == 1 && opp == 2 {
		battle(&bulbasaur, &charmander)
	} else if mypokemon == 1 && opp == 3 {
		battle(&bulbasaur, &squirtle)
	} else if mypokemon == 2 && opp == 1 {
		battle(&charmander, &bulbasaur)
	} else if mypokemon == 2 && opp == 3 {
		battle(&charmander, &squirtle)
	} else if mypokemon == 3 && opp == 1 {
		battle(&squirtle, &bulbasaur)
	} else if mypokemon == 3 && opp == 2 {
		battle(&squirtle, &charmander)
	}
}
