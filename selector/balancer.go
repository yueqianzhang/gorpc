package selector

import (
	"math/rand"
	"time"
)

type Balancer interface {
	Balance(string, []*Node) *Node
}

var balancerMap = make(map[string]Balancer, 0)

const (
	Random = "random"
	RoundRobin = "roundRobin"
	WeightedRoundRobin = "weightedRoundRobin"
	ConsistentHash = "consistentHash"

	Custom = "custom"
)

func init() {
	RegisterBalancer(Random, DefaultBalancer)
	RegisterBalancer(RoundRobin, RoundRobinBalancer)
}

var DefaultBalancer = &randomBalancer{}
var RoundRobinBalancer = newRoundRobinBalancer()

func RegisterBalancer(name string, balancer Balancer) {
	if balancerMap == nil {
		balancerMap = make(map[string]Balancer)
	}
	balancerMap[name] = balancer
}

func GetBalancer(name string) Balancer {
	if balancer, ok := balancerMap[name]; ok {
		return balancer
	}
	return DefaultBalancer
}

type randomBalancer struct {

}

func (r *randomBalancer) Balance(serviceName string, nodes []*Node) *Node {
	if len(nodes) == 0 {
		return nil
	}
	rand.Seed(time.Now().Unix())
	num := rand.Intn(len(nodes))
	return nodes[num]
}



