package main

import (
	"fmt"
	"sync"
)

func siritori(wg *sync.WaitGroup, spells [][]rune, result []string) {
	defer wg.Done()
	var use = false
	for i, spell := range spells {
		var lastSpell = []rune(result[len(result)-1])
		var lastRune = lastSpell[len(lastSpell)-1]
		var firstRune = spell[0]
		if lastRune == firstRune {
			use = true
			s, r := useSpell(spells, result, i)
			wg.Add(1)
			go siritori(wg, s, r)
		}
	}
	if !use {
		fmt.Println(len(result), result)
	}
}

func useSpell(spells [][]rune, result []string, spellIndex int) ([][]rune, []string) {
	var cResult = make([]string, 0, len(result)+1)
	for _, v := range result {
		cResult = append(cResult, v)
	}
	cResult = append(cResult, string(spells[spellIndex]))

	var cSpells = make([][]rune, 0, len(spells)-1)
	for i, v := range spells {
		if i == spellIndex {
			continue
		}
		cSpells = append(cSpells, v)
	}
	return cSpells, cResult
}
