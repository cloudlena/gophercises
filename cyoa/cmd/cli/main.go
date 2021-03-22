package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/mastertinner/gophercises/cyoa"
)

func main() {
	arcsFilePath := flag.String(
		"arcs-file",
		"gopher.json",
		"the path to the JSON file containing the arcs of the story",
	)
	flag.Parse()

	arcs, err := cyoa.ArcsFromFile(*arcsFilePath)
	if err != nil {
		log.Fatal(fmt.Errorf("error getting arcs from file: %w", err))
	}

	printArc(arcs, "intro")
}

// printArc prints the requested story arc to the terminal.
func printArc(arcs map[string]cyoa.Arc, name string) {
	clearScreen()

	arc, ok := arcs[name]
	if !ok {
		log.Fatal(fmt.Sprintf("arc %s not found", name))
	}

	fmt.Println(arc.Title)
	fmt.Println("")
	fmt.Println("")
	for _, s := range arc.Story {
		fmt.Println(s)
		fmt.Println("")
	}
	fmt.Println("")
	if len(arc.Options) == 0 {
		fmt.Println("That's all, folks! Thanks for playing!")
		os.Exit(0)
		return
	}
	fmt.Println("Press a number key to choose your option:")
	fmt.Println("")
	for i, o := range arc.Options {
		fmt.Println(fmt.Sprintf("%v: %s", i+1, o.Text))
	}

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	index, err := strconv.ParseInt(input.Text(), 10, 64)
	if err != nil {
		log.Fatal("invalid option")
	}
	nextArcName := arc.Options[int(index-1)].Arc
	printArc(arcs, nextArcName)
}

// clearScreen clears the terminal screen for different
// operating systems.
func clearScreen() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
