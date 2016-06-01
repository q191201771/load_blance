package load_blance

import (
	ch "github.com/q191201771/consistent_hash"
)

type ConsistentHash struct {
	core ch.ConsistentHash
}

func NewConsistentHash(virutalMultis int) *ConsistentHash {
	var core ch.ConsistentHash
	if virutalMultis == 0 {
		core = ch.Default()
	} else {
		core = ch.New(virutalMultis)
	}
	return &ConsistentHash{core: core}
}

func (obj *ConsistentHash) Exist(node string) bool {
	return obj.core.Exist(node)
}

func (obj *ConsistentHash) Add(node string) error {
	err := obj.core.Add(node)
	if err == ch.ErrNodeAlreadyExist {
		err = ErrNodeAlreadyExist
	}
	return err
}

func (obj *ConsistentHash) Remove(node string) error {
	err := obj.core.Remove(node)
	if err == ch.ErrNodeNotFound {
		err = ErrNodeNotFound
	}
	return err
}

func (obj *ConsistentHash) Get(key string) (string, error) {
	node, err := obj.core.Get(key)
	if err == ch.ErrNoNode {
		err = ErrNoNode
	}
	return node, err
}

func (obj *ConsistentHash) Nodes() []string {
	return obj.core.Nodes()
}

func (obj *ConsistentHash) Clear() {
	obj.core.Clear()
}
