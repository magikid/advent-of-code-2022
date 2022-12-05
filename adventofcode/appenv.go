package adventofcode

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"sync"
)

const AllDays = 0
const AllParts = 0

type appEnv struct {
	day             int
	part            int
	puzzleInputPath string
}

func (app appEnv) String() string {
	return fmt.Sprintf("day: %v, part: %v, puzzle input path: %v", app.day, app.part, app.puzzleInputPath)
}

func (app *appEnv) fromArgs(args []string) error {
	fl := flag.NewFlagSet("adventofcode", flag.ContinueOnError)
	fl.IntVar(&app.day, "day", AllDays, "Which day should be run?")
	fl.IntVar(&app.part, "part", AllParts, "Which part of the given day should be run?")
	fl.StringVar(&app.puzzleInputPath, "path", "inputs", "The path to the puzzle inputs")

	if err := fl.Parse(args); err != nil {
		fl.Usage()
		return err
	}

	return nil
}

func (app *appEnv) run() error {
	entries, err := os.ReadDir(app.puzzleInputPath)
	if err != nil {
		return err
	}

	puzzleInputMatcher := regexp.MustCompile(`day(\d+)`)
	wg := sync.WaitGroup{}
	inputFileCount := 0

	for _, entry := range entries {
		if entry.Type().IsRegular() && puzzleInputMatcher.MatchString(entry.Name()) {
			matches := puzzleInputMatcher.FindStringSubmatch(entry.Name())
			currentDay, _ := strconv.ParseInt(matches[1], 10, 0)
			puzzleInput, err := app.readPuzzleInput(entry)
			if err != nil {
				return err
			}

			puzzle := Puzzle{currentDay, puzzleInput}
			inputFileCount++
			wg.Add(1)
			go puzzle.Solve(&wg)
		} else {
			log.Printf("Input file %s does not match pattern day(\\d+), skipping", entry.Name())
		}
	}

	if inputFileCount == 0 {
		return errors.New("no input files found in inputs/ directory")
	}

	wg.Wait()

	return nil
}

func (app *appEnv) readPuzzleInput(entry os.DirEntry) (string, error) {
	content, err := ioutil.ReadFile(app.puzzleInputPath + string(os.PathSeparator) + entry.Name())
	if err != nil {
		return "", err
	}

	return string(content), nil
}
