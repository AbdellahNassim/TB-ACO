package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"gonum.org/v1/gonum/plot"
	"gonum.org/v1/gonum/plot/plotter"
	"gonum.org/v1/gonum/plot/vg"
)

type City struct {
	id   int
	x, y float64
}

type Ant struct {
	path       []int
	distance   float64
	trust      float64
	trustMin   float64
	successRate float64
}

type TBACO struct {
	cities       []City
	nAnts         int
	pheromones    [][]float64
	distances     [][]float64
	alpha, beta  float64
	rho           float64
	ka            *Ant
	trustHistory  []float64
	standardACOHistory []float64
}

func euclideanDistance(a, b City) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

func (aco *TBACO) initialize(numCities, numAnts int) {
	rand.Seed(time.Now().UnixNano())
	aco.nAnts = numAnts
	aco.alpha, aco.beta, aco.rho = 1.0, 2.0, 0.5

	aco.cities = make([]City, numCities)
	aco.distances = make([][]float64, numCities)
	aco.pheromones = make([][]float64, numCities)

	for i := 0; i < numCities; i++ {
		aco.cities[i] = City{i, rand.Float64() * 100, rand.Float64() * 100}
		aco.distances[i] = make([]float64, numCities)
		aco.pheromones[i] = make([]float64, numCities)
		for j := 0; j < numCities; j++ {
			if i != j {
				aco.distances[i][j] = euclideanDistance(aco.cities[i], aco.cities[j])
				aco.pheromones[i][j] = 1.0
			}
		}
	}

	aco.ka = &Ant{trust: 0.1, trustMin: 0.1, successRate: 0.0}
	aco.trustHistory = []float64{}
	aco.standardACOHistory = []float64{}
}

func (aco *TBACO) updateTrust(success bool) {
	alpha := 0.1
	gamma := 0.02

	if success {
		aco.ka.successRate += 1.0
		aco.ka.trust += (1.0 - aco.ka.trust) * alpha
	}

	aco.ka.trustMin = math.Min(0.1+gamma*aco.ka.successRate, 1.0)
	aco.ka.trust = math.Max(aco.ka.trust, aco.ka.trustMin)
	aco.trustHistory = append(aco.trustHistory, aco.ka.trust)
}

func (aco *TBACO) run(iterations int) {
	for i := 0; i < iterations; i++ {
		aco.updateTrust(rand.Float64() < 0.7)
	}
}

func (aco *TBACO) benchmarkStandardACO() {
	fmt.Println("Running standard ACO for comparison...")
	standardTrust := 0.1
	alpha := 0.1
	for i := 0; i < 100; i++ {
		if rand.Float64() < 0.7 {
			standardTrust += (1.0 - standardTrust) * alpha
		}
		aco.standardACOHistory = append(aco.standardACOHistory, standardTrust)
	}
}

func (aco *TBACO) plotResults() {
	p := plot.New()
	p.Title.Text = "TB-ACO vs Standard ACO Trust Growth"
	p.X.Label.Text = "Iterations"
	p.Y.Label.Text = "Trust Level"

	tpoints := make(plotter.XYs, len(aco.trustHistory))
	sapoints := make(plotter.XYs, len(aco.standardACOHistory))

	for i := range aco.trustHistory {
		tpoints[i].X = float64(i)
		tpoints[i].Y = aco.trustHistory[i]
	}
	for i := range aco.standardACOHistory {
		sapoints[i].X = float64(i)
		sapoints[i].Y = aco.standardACOHistory[i]
	}

	trustLine, _ := plotter.NewLine(tpoints)
	standardLine, _ := plotter.NewLine(sapoints)
	trustLine.Color = plotutil.Color(0)
	standardLine.Color = plotutil.Color(1)

	p.Add(trustLine, standardLine)
	p.Legend.Add("TB-ACO", trustLine)
	p.Legend.Add("Standard ACO", standardLine)

	if err := p.Save(8*vg.Inch, 4*vg.Inch, "trust_comparison.png"); err != nil {
		fmt.Println("Failed to save plot:", err)
	} else {
		fmt.Println("Plot saved as trust_comparison.png")
	}
}

func main() {
	fmt.Println("Little kids, are you still sleeping? Say 'Alhamdulillah' before running TB-ACO! 😏")
	aco := TBACO{}
	aco.initialize(10, 5)
	aco.run(100)
	aco.benchmarkStandardACO()
	aco.plotResults()
}
