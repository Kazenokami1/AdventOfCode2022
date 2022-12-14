package days

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Day1() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day1Sample.txt"
	} else {
		fileName = "inputfiles/Day1.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var elves []Elf
	var meals []int
	for scanner.Scan() {
		if scanner.Text() != "" {
			meal, err := strconv.Atoi(scanner.Text())
			Check(err)
			meals = append(meals, meal)
		} else {
			var elf Elf
			elf.Meals = meals
			elf.calcTotalCalories()
			elves = append(elves, elf)
			meals = []int{}
		}
	}
	//Should figure out a better way to get the last elf added to my array
	var elf Elf
	elf.Meals = meals
	elf.calcTotalCalories()
	elves = append(elves, elf)
	var highestCalories int
	//Part 1 Answer
	for _, elf := range elves {
		if elf.TotalCalories > highestCalories {
			highestCalories = elf.TotalCalories
		}
	}
	fmt.Println("Day1 Puzzle Solutions:")
	fmt.Printf("Highest Elf has %d calories", highestCalories)
	fmt.Println()
	var totalCaloriesPerElf []int
	for _, elf := range elves {
		totalCaloriesPerElf = append(totalCaloriesPerElf, elf.TotalCalories)
	}
	sort.Ints(totalCaloriesPerElf)
	var top3Calories int
	for i := len(totalCaloriesPerElf); i > len(totalCaloriesPerElf)-3; i-- {
		top3Calories += totalCaloriesPerElf[i-1]
	}
	fmt.Printf("Total Calories carried by top 3 elves is %d", top3Calories)
	fmt.Println()
}
