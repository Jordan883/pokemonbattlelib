package pokemonbattlelib

import "fmt"

type Pokemon struct {
    NatDex              uint16      // National Pokedex Number
    Nickname            string      // player-given name for the Pokemon
    Level               uint8       // value from 1-100 influencing stats
    Ability             Ability     // name of this Pokemon's ability
    TotalExperience     uint        // the total amount of exp this Pokemon has gained, influencing its level
    Gender              Gender      // this Pokemon's gender
    IVs                 [6]uint8    // values from 0-31 that represents a Pokemon's 'genetic' potential
    EVs                 [6]uint8    // values from 0-255 that represents a Pokemon's training in a particular stat
    Nature              Nature      // represents a Pokemon's disposition and affects stats
    Stats               [6]uint     // the actual stats of a Pokemon determined from the above data
    CurrentHP           uint        // the remaining HP of this Pokemon 
    HeldItem            Item        // the item a Pokemon is holding
    Moves               [4]Move     // the moves the Pokemon currenly knows
    Friendship          uint8       // how close this Pokemon is to its Trainer
    OriginalTrainerID   uint16      // a number associated with the first Trainer who caught this Pokemon
}

func (p *Pokemon) GetName() string {
	return PokemonNames[p.NatDex]
}

func (p *Pokemon) HasValidLevel() bool {
    return p.Level > 1 && p.Level <= 100
}

func (p *Pokemon) HasValidIVs() bool {
    for _, IV := range p.IVs {
        if (IV > 31) {
            return false
        }
    }
    return true
}

func (p *Pokemon) HasValidEVs() bool {
    var totalEVs := 0
    for _, EV := range p.EVs {
        totalEVs += EV
    }
    
    if (totalEVs > 510) {
        return false
    }

    return true
}

// implement Stringer

// display a Pokemon close to how appear in a Pokemon battle
func (p *Pokemon) String() string {
    var displayName string

    if (p.Nickname == "") {
        displayName = p.GetName()
    } else { // the pokemon has a nickname
        displayName = p.Nickname
    }
    
    return fmt.Sprintf("%v%v\tLv%d\nHP: %d/%d\n", displayName, 
        p.Gender, p.Level, p.CurrentHP, p.Stats[0])
}

type Item string

