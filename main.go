package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

var loadFile = flag.String("f", "", "target CSV file")

func main() {
	os.Exit(execute())
}

func execute() int {
	flag.Parse()
	if len(*loadFile) < 1 {
		flag.PrintDefaults()
		return 1
	}
	var spells = loadData(*loadFile)

	for i, spell := range spells {
		fmt.Printf("========= %c ==========\n", spell)
		var wg = &sync.WaitGroup{}
		var result []string
		s, r := useSpell(spells, result, i)
		wg.Add(1)
		siritori(wg, s, r)
		wg.Wait()
	}
	return 0
}

func loadData(fileName string) [][]rune {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	data = bytes.TrimRight(data, "\n")
	var runeData = make([][]rune, 0, len(data))
	for _, v := range strings.Split(string(data), ",") {
		runeData = append(runeData, []rune(v))
	}
	fmt.Printf("SpellBook: %d %c\n", len(runeData), runeData)
	return runeData
}
