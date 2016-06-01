package load_blance

type LoadBlance interface {
	Exist(node string) bool
	Add(node string) error
	Remove(node string) error
	Get(key string) (string, error) /// key only for consistent hash, others ignore it~
	Nodes() []string
	Clear()
}

/// config e.g.  -> {
///   virutalMultis `brief: only for consistent hash, set 0 if use default`
/// }
func LoadBlanceFactory(name string, virutalMultis int) LoadBlance {
	switch name {
	case "consistent hash":
		return NewConsistentHash(virutalMultis)
	case "random selector":
		return &RandomSelector{}
	case "round robin":
		return &RoundRobin{}
	default:
		return nil
	}
	return nil
}
