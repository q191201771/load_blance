package load_blance

import (
	"log"
	"testing"
)

func create() ([]LoadBlance, []string) {
	names := []string{
		"consistent hash",
		"random selector",
		"round robin",
	}
	var lbs []LoadBlance
	for _, v := range names {
		lb := LoadBlanceFactory(v, 20)
		lbs = append(lbs, lb)
	}
	return lbs, names
}

func TestFactory(t *testing.T) {
	if LoadBlanceFactory("consistent hash", 20) == nil ||
		LoadBlanceFactory("random selector", 20) == nil ||
		LoadBlanceFactory("round robin", 20) == nil ||
		LoadBlanceFactory("asd", 20) != nil {
		t.Fatal("Factory fail.")
	}
}

func TestBasic(t *testing.T) {
	lbs, _ := create()
	for _, v := range lbs {
		if v.Exist("test") {
			t.Fatal("!")
		}
		if node, err := v.Get("test"); node != "" || err != ErrNoNode {
			t.Fatal("!")
		}
		if len(v.Nodes()) != 0 {
			t.Fatal("!")
		}
		if err := v.Remove("test"); err != ErrNodeNotFound {
			t.Fatal("!")
		}
		if err := v.Add("test"); err != nil {
			t.Fatal("!")
		}
		if !v.Exist("test") {
			t.Fatal("!")
		}
		if node, err := v.Get("aaa"); node != "test" || err != nil {
			t.Fatal("!")
		}
		if err := v.Add("test"); err != ErrNodeAlreadyExist {
			t.Fatal("!")
		}
		if err := v.Remove("test"); err != nil {
			t.Fatal("!")
		}
		v.Add("node1")
		v.Add("node2")
		if len(v.Nodes()) != 2 {
			t.Fatal("!")
		}
		v.Clear()
		if len(v.Nodes()) != 0 {
			t.Fatal("!")
		}
	}
}

func TestHuge(t *testing.T) {
	lbs, names := create()
	for i, v := range lbs {
		v.Add("127.0.0.1:8384")
		v.Add("127.0.0.1:10002")
		v.Add("127.0.0.1:9001")

		node2count := map[string]int{
			"127.0.0.1:8384":  0,
			"127.0.0.1:10002": 0,
			"127.0.0.1:9001":  0,
		}
		log.Println("-----Huge", names[i])
		for i := 0; i < 100000; i++ {
			node, err := v.Get("test")
			if err != nil {
				t.Fatal("!")
			}
			node2count[node]++
		}
		for k, v := range node2count {
			log.Println("  ", k, v)
		}
	}
}
