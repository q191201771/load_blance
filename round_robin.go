package load_blance

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

type RoundRobin struct {
	sync.RWMutex
	nodes []string
	index int
}

func NewRoundRobin() *RoundRobin {
	return &RoundRobin{}
}

func (rs *RoundRobin) Exist(node string) bool {
	rs.RLock()
	defer rs.RUnlock()
	return arrayExist(rs.nodes, node)
}

func (rs *RoundRobin) Add(node string) error {
	rs.Lock()
	defer rs.Unlock()
	if arrayExist(rs.nodes, node) {
		return ErrNodeAlreadyExist
	}
	rs.nodes = append(rs.nodes, node)
	return nil
}

func (rs *RoundRobin) Remove(node string) error {
	rs.Lock()
	defer rs.Unlock()
	if !arrayExist(rs.nodes, node) {
		return ErrNodeNotFound
	}
	rs.nodes = arraySplice(rs.nodes, node)
	return nil
}

func (rs *RoundRobin) Get(key string) (string, error) {
	rs.RLock()
	defer rs.RUnlock()
	numOfNodes := len(rs.nodes)
	if numOfNodes == 0 {
		return "", ErrNoNode
	}
	node := rs.nodes[rs.index%numOfNodes]
	rs.index++
	if rs.index > math.MaxInt32 {
		rs.index = 0
	}
	return node, nil
}

func (rs *RoundRobin) Nodes() []string {
	rs.RLock()
	defer rs.RUnlock()
	return rs.nodes
}

func (rs *RoundRobin) Clear() {
	rs.Lock()
	defer rs.Unlock()
	rs.nodes = arrayClear(rs.nodes)
}

func init() {
	rand.Seed(time.Now().Unix())
}
