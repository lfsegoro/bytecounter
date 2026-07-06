package collector

type InterfaceStat struct {
	Name    string
	RxBytes uint64
	TxBytes uint64
}

type Collector interface {
	Collect() ([]InterfaceStat, error)
}