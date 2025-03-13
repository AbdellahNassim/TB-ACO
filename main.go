package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Ant struct {
	id        int
	path      []int
	reward    float64
	isKA      bool
}

type Colony struct {
	ants        []*Ant
	trust       float64
	trustMin    float64
	successRate float64
}

const (
	Alpha = 0.1  // Learning rate for trust growth
	Beta  = 0.05 // Decay factor (for alternative version)
	Gamma = 0.02 // Adaptive threshold scaling factor
	MaxIterations = 100
)

func (c *Colony) updateTrust(success bool) {
	if success {
		c.successRate += 1
		c.trust += (1 - c.trust) * (1 - math.Exp(-Alpha * c.successRate))
	} else {
		c.trust *= math.Exp(-Beta)
	}

	// Adaptive minimum trust
	c.trustMin = math.Min(1, 0.1 + Gamma * c.successRate)
	c.trust = math.Max(c.trust, c.trustMin)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	colony := Colony{trust: 0.1, trustMin: 0.1}

	for i := 0; i < MaxIterations; i++ {
		kaSuccess := rand.Float64() < 0.7 // Simulate KA's success rate
		colony.updateTrust(kaSuccess)
		fmt.Printf("Iteration %d: Trust = %.4f, TrustMin = %.4f\n", i+1, colony.trust, colony.trustMin)
	}
}