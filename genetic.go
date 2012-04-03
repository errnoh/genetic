// Generic Genetic Algorithm
package genetic

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

var (
	GA_POPSIZE      = 2048
	GA_MAXITER      = 16384
	GA_ELITRATE     = 0.10
	GA_MUTATIONRATE = 0.25
	GA_MUTATION     = int(float64(math.MaxUint16) * GA_MUTATIONRATE)
	RAND_MAX        = math.MaxUint16
)

type Interface interface {
	GetFitness() uint
	Fitness()
	Init()
	Mutate()
	Mate(Interface, Interface)
	String() string
}

type Population []Interface

func (p Population) Init(f func() Interface) Population {
	var citizen Interface
	for i := 0; i < GA_POPSIZE; i++ {
		citizen = f()
		citizen.Init()

		p = append(p, citizen)
	}
	return p
}

func (p Population) CalcFitness() {
	for i := 0; i < GA_POPSIZE; i++ {
		p[i].Fitness()
	}
}

func (p Population) Mate() {
	esize := int(float64(GA_POPSIZE) * GA_ELITRATE)

	for i := esize; i < GA_POPSIZE; i++ {
		// /2 is to prefer better candidates
		o1 := rand.Intn(RAND_MAX) % (GA_POPSIZE / 2)
		o2 := rand.Intn(RAND_MAX) % (GA_POPSIZE / 2)
		p[i].Mate(p[o1], p[o2])
		if rand.Intn(RAND_MAX) < GA_MUTATION {
			p[i].Mutate()
		}
	}
}

func (p Population) Len() int {
	return len(p)
}

func (p Population) Less(i, j int) bool {
	return p[i].GetFitness() < p[j].GetFitness()
}

func (p Population) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// params:: f - constructor for the used data type.
func Init(f func() Interface) Population {
	p := make(Population, 0, GA_POPSIZE)
	p = p.Init(f)
	return p
}

func (p Population) Run() {
	for i := 0; i < GA_MAXITER; i++ {
		p.CalcFitness()
		sort.Sort(p)
		fmt.Printf("i: %d - %s\n", i, p[0])
		if p[0].GetFitness() == 0 {
			return
		}
		p.Mate()
	}
}
