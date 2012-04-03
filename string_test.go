package genetic

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	p := Init(NewGA)
	p.Run()
	if p[0].GetFitness() != 0 {
		t.Fatalf("Failed to get target text: %s", p[0].String())
	}
}

var (
	GA_TARGET = "Hello world!"
	// use genetic.RAND_MAX
)

func NewGA() Interface {
	return new(GA)
}

func init() {
	rand.Seed(time.Now().Unix())
}

type GA struct {
	str     string
	fitness uint
}

func (ga *GA) GetFitness() uint {
	return ga.fitness
}

func (ga *GA) Init() {
	if len(ga.str) != 0 {
		ga.str = ""
	}
	for j := 0; j < len(GA_TARGET); j++ {
		ga.str += string((rand.Intn(RAND_MAX) % 90) + 32)
	}
}

func (ga *GA) Fitness() {
	var fitness uint
	for j := 0; j < len(GA_TARGET); j++ {
		fitness += uint(abs(ga.str[j] - GA_TARGET[j]))
	}
	ga.fitness = fitness
}

func abs(value byte) byte {
	if value < 0 {
		return 0 - value
	}
	return value
}

func (ga *GA) Mutate() {
	ipos := rand.Intn(RAND_MAX) % len(GA_TARGET)
	delta := uint8((rand.Intn(RAND_MAX) % 90) + 32)
	temp := []byte(ga.str)
	temp[ipos] = ((ga.str[ipos] + delta) % 122)
	ga.str = string(temp)
}

func (ga *GA) Mate(a, b Interface) {
	o1, _ := a.(*GA)
	o2, _ := b.(*GA)
	spos := rand.Intn(RAND_MAX) % len(GA_TARGET)

	target := []byte(ga.str)
	copy(target[:spos], o1.str[:spos])
	copy(target[spos:], o2.str[spos:])
	ga.str = string(target)
}

func (ga *GA) String() string {
	return fmt.Sprintf("%s (%d)", ga.str, ga.fitness)
}
