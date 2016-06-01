package load_blance

func arrayExist(nodes []string, node string) bool {
	return arrayFind(nodes, node) != -1
}

func arrayFind(nodes []string, node string) int {
	for i, v := range nodes {
		if v == node {
			return i
		}
	}
	return -1
}

func arraySplice(nodes []string, node string) []string {
	index := arrayFind(nodes, node)
	if index == -1 {
		return nodes
	}
	return arraySpliceByIndex(nodes, index)
}

func arraySpliceByIndex(nodes []string, i int) []string {
	if i < 0 || i >= len(nodes) {
		return nodes
	}
	return append(nodes[:i], nodes[i+1:]...)
}

func arrayClear(nodes []string) []string {
	return nodes[:0]
}
