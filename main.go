package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Flags struct {
	Mean bool
	Median bool
	Mode bool
	Sd bool
}

var (
	MEAN_FLAG = "-mean"
	MEDIAN_FLAG = "-median"
	MODE_FLAG = "-mode"
	SD_FLAG = "-sd"
	HELP = "USAGE:\ngo run main {-mean} {-sd} {-mode} {-median}"
)

func setAllFlagsTrue(flags *Flags) {
	flags.Mean = true
	flags.Sd = true
	flags.Mode = true
	flags.Median = true
}

func main() {
	flags := Flags{}

	arguments := os.Args

	if len(arguments) < 2 {
		setAllFlagsTrue(&flags)
	} else {
		for i := 1; i < len(arguments); i++ {
			if strings.EqualFold(arguments[i], MEAN_FLAG) {
				flags.Mean = true
			} else if strings.EqualFold(arguments[i], MODE_FLAG) {
				flags.Mode = true
			} else if strings.EqualFold(arguments[i], MEDIAN_FLAG) {
				flags.Median = true
			} else if strings.EqualFold(arguments[i], SD_FLAG) {
				flags.Sd = true
			} else {
				fmt.Println(HELP)
				return
			}
		}
	}

	scanner := bufio.NewScanner(os.Stdin)

	sequence := make([]int, 0)

	for scanner.Scan() {
		if numb, err := strconv.Atoi(scanner.Text()); err != nil {
			fmt.Println(err.Error())

			return
		} else if numb <= 100000 && numb >= -100000 {
			sequence = append(sequence, numb)
		} else {
			fmt.Println("Не валидное входящее значение")

			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err.Error())

		return
	}

	if len(sequence) == 0 {
		fmt.Println("Не достаточно элементов в последовательности")

		return
	}

	sort.Ints(sequence)

	if flags.Median {
		fmt.Printf("Median: %.2f\n", calculateMedian(sequence))
	}

	if flags.Mode {
		fmt.Println("Mode:", calculateMode(sequence))
	}

	if flags.Mean {
		fmt.Printf("Mean: %.2f\n", calculateMean(sequence))
	}

	if flags.Sd {
		fmt.Printf("Sd: %.2f\n", calculateSd(sequence))
	}
}

func calculateSd(sequence []int) float64 {
	sum := float64(0)

	mean := calculateMean(sequence)

	for i := range sequence {
		sum += math.Pow(float64(sequence[i]) - float64(mean), 2)
	}

	return math.Sqrt(sum / float64(len(sequence)))
}

func calculateMean(sequence []int) float32 {
	sum := 0

	for i := range sequence {
		sum += sequence[i]
	}

	return float32(sum) / float32(len(sequence))
}

func calculateMode(sequence []int) int {
	sequenceDictionary := make(map[int]int)

	for i := range sequence {
		if _, ok := sequenceDictionary[sequence[i]]; !ok {
			sequenceDictionary[sequence[i]] = 1
		} else {
			sequenceDictionary[sequence[i]]++
		}
	}

	maxValue := -100001
	var rememberValue int

	for _, val := range sequence {
		if sequenceDictionary[val] > maxValue {
			maxValue = sequenceDictionary[val]
			rememberValue = val
		}
	}

	return rememberValue
}

func calculateMedian(sequence []int) float32 {
	if len(sequence) == 0 {
		return 0
	}
	mod := len(sequence) / 2

	if len(sequence) % 2 == 0 {
		return float32(sequence[mod - 1] + sequence[mod]) / 2
	} else {
		return float32(sequence[mod])
	}
}
