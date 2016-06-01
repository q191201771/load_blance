package load_blance

import (
	"math/rand"
	"sync"
	"time"
)

type RandomSelector struct {
	sync.RWMutex
	nodes []string
}

func NewRandomSelector() *RandomSelector {
	return &RandomSelector{}
}

func (rs *RandomSelector) Exist(node string) bool {
	rs.RLock()
	defer rs.RUnlock()
	return arrayExist(rs.nodes, node)
}

func (rs *RandomSelector) Add(node string) error {
	rs.Lock()
	defer rs.Unlock()
	if arrayExist(rs.nodes, node) {
		return ErrNodeAlreadyExist
	}
	rs.nodes = append(rs.nodes, node)
	return nil
}

func (rs *RandomSelector) Remove(node string) error {
	rs.Lock()
	defer rs.Unlock()
	if !arrayExist(rs.nodes, node) {
		return ErrNodeNotFound
	}
	rs.nodes = arraySplice(rs.nodes, node)
	return nil
}

func (rs *RandomSelector) Get(key string) (string, error) {
	rs.RLock()
	defer rs.RUnlock()
	numOfNodes := len(rs.nodes)
	if numOfNodes == 0 {
		return "", ErrNoNode
	}
	return rs.nodes[rand.Intn(numOfNodes)], nil
}

func (rs *RandomSelector) Nodes() []string {
	rs.RLock()
	defer rs.RUnlock()
	return rs.nodes
}

func (rs *RandomSelector) Clear() {
	rs.Lock()
	defer rs.Unlock()
	rs.nodes = arrayClear(rs.nodes)
}

func init() {
	rand.Seed(time.Now().Unix())
}
