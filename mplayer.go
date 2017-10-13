package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	
	"mlib"
	"mp"
)

var lib * library.MusicManager
var id int = 1
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i:=0; i<lib.Len() ; i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 7 {
			id++
			lib.Add(&library.MusicEntry{strconv.Itoa(id), tokens[2], tokens[3], tokens[4], tokens[5], tokens[6]})
		}else {
			fmt.Println("Usage: lib add <name><artist><genre><source><type> (7 argv)")
		}
	case "remove":
		if len(tokens) == 3{
			lib.RemoveByName(tokens[2])
		}else {
			fmt.Println("Usage: lib remove <name>")
		}
	default:
		fmt.Println("Unrecongnied lib command:", tokens[1])
	}
}

func handlePlayerCommands(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("Usage: play <name>")
		return
	}

	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The music not found", tokens[1])
		return
	}

	mp.Play(e.Source, e.Type)
}

func main() {
	fmt.Println(
		`
		lib list --View the existing music lib
		lib add <name><artist><genre><source><type> -- Add a music to the music lib
		lib remove <name> --Remove the specified music from the lib
		play <name> -- Play the specified music
		`)
	lib = library.NewMusicManager()
	r := bufio.NewReader(os.Stdin)

	for i:=0; i<100; i++ {
		fmt.Println("Enter command->")
		rawLine,_,_ := r.ReadLine()
		line := string(rawLine)

		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")

		if tokens[0] == "lib"{
			handleLibCommands(tokens)
		}else if tokens[0] == "play" {
			handlePlayerCommands(tokens)
		}else{
			fmt.Println("Unreconginzed command:", tokens[0])
		}
	}
}