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
	f, err := os.Open("./datasets/train_100.csv")
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

		dataset = append(dataset, data{tf, ds})
	}

	fmt.Println(len(dataset))

	n := jarvis.NewNetwork(784, 256, 10, 0.01)

	for i := 0; ; i++ {
		start := time.Now()
		t := 0.0
		for _, d := range dataset {
			t += n.Train(d.data, d.labels)
		}
		fmt.Println(fmt.Sprintf("> EPOCH %v - TOTAL ERROR: %v - TOTAL AVERAGE ERROR: %v - TIME TAKEN: %v", i, t, t/float64(len(dataset)), time.Since(start)))
	}
}

func saveWeights(n *jarvis.Network) {

	f, _ := os.Create("./weights.wgt")
	defer f.Close()

	d, _ := json.MarshalIndent(n, "", "\t")

	os.WriteFile("./network.wgt", d, 0644)
}

type data struct {
	labels jarvis.Matrix
	data   jarvis.Matrix
}
