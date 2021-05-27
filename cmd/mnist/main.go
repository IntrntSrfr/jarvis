package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/intrntsrfr/jarvis"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {

	var ans int
	n := jarvis.NewNetwork(784, 256, 10, 0.1)

	for {
		fmt.Println("1. TRAIN\n2. GUESS\n3. LOAD\n4. SAVE\n5. EXIT")
		fmt.Scan(&ans)
		if ans == 1 {
			fmt.Println("select training criteria:\n1. epochs\n2. total error\n3. average error")
			fmt.Scan(&ans)
			if ans == 1 {
				train(n, 100, 0, 0)
				fmt.Println(1, 1)
			} else if ans == 2 {
				train(n, 0, 20, 0)
				fmt.Println(1, 2)
			} else if ans == 3 {
				train(n, 0, 0, 0.1)
				fmt.Println(1, 3)
			}
		} else if ans == 2 {
			guess(n)
			fmt.Println(2)
		} else if ans == 3 {
			fmt.Println(3)
		} else if ans == 4 {
			saveWeights(n)
			fmt.Println(4)
		} else if ans == 5 {
			break
		}
	}
}

func guess(n *jarvis.Network) {
	f, err := os.Open("./datasets/test_10.csv")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	var dataset []data

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		splits := strings.Split(t, ",")

		labels := make([]float64, 10)
		hot, _ := strconv.Atoi(splits[0])
		labels[hot] = 1.0

		d := make([]float64, 784)
		for i := range d {
			p, _ := strconv.ParseFloat(splits[i+1], 64)
			d[i] = p
		}

		tf := make(jarvis.Matrix, len(labels))
		for l := range labels {
			tf[l] = jarvis.Vector{labels[l]}
		}

		ds := make(jarvis.Matrix, len(d))
		for da := range d {
			ds[da] = jarvis.Vector{d[da]}
		}

		dataset = append(dataset, data{tf, ds, hot})
	}

	right := 0.0

	for _, d := range dataset {
		g := n.Guess(d.data)
		fmt.Println(fmt.Sprintf("GUESS: %v\nGOT: %v", d.labels, g))

		largest := -1
		val := 0.0
		for v := range g {
			if g[v][0] > val {
				val = g[v][0]
				largest = v
			}
		}
		fmt.Println("GUESSED: ", largest)
		fmt.Println("WANTED", d.label)
		fmt.Println()

		if largest == d.label {
			right++
		}
	}

	fmt.Println("GUESSES:", right/float64(len(dataset)))
}

func train(n *jarvis.Network, epochs int, totalErr, avgErr float64) {
	fmt.Println("lol")
	f, err := os.Open("./datasets/train_large.csv")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	var dataset []data

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		splits := strings.Split(t, ",")

		labels := make([]float64, 10)
		hot, _ := strconv.Atoi(splits[0])
		labels[hot] = 1.0

		d := make([]float64, 784)
		for i := range d {
			p, _ := strconv.ParseFloat(splits[i+1], 64)
			d[i] = p
		}

		tf := make(jarvis.Matrix, len(labels))
		for l := range labels {
			tf[l] = jarvis.Vector{labels[l]}
		}

		ds := make(jarvis.Matrix, len(d))
		for da := range d {
			ds[da] = jarvis.Vector{d[da]}
		}

		dataset = append(dataset, data{tf, ds, hot})
	}

	for i := 0; ; i++ {
		start := time.Now()
		t := 0.0
		for _, d := range dataset {
			t += n.Train(d.data, d.labels)
		}

		fmt.Println(fmt.Sprintf("> EPOCH %v - TOTAL ERROR: %v - TOTAL AVERAGE ERROR: %v - TIME TAKEN: %v", i, t, t/float64(len(dataset)), time.Since(start)))

		if avgErr != 0 && t/float64(len(dataset)) < avgErr {
			break
		} else if totalErr != 0 && t < totalErr {
			break
		} else if epochs != 0 && epochs == i {
			break
		}
	}
}

func saveWeights(n *jarvis.Network) {
	d, _ := json.MarshalIndent(n, "", "\t")
	os.WriteFile("./network.wgt", d, 0644)
}

type data struct {
	labels jarvis.Matrix
	data   jarvis.Matrix
	label  int
}
