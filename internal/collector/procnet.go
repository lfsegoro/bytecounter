package collector

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type ProcCollector struct{}

func NewProcCollector() *ProcCollector {
	return &ProcCollector{}
}

func (c *ProcCollector) Collect() ([]InterfaceStat, error) {

	file, err := os.Open("/proc/net/dev")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var stats []InterfaceStat

	scanner := bufio.NewScanner(file)

	line := 0

	for scanner.Scan() {

		line++

		if line <= 2 {
			continue
		}

		fields := strings.Fields(scanner.Text())

		name := strings.TrimSuffix(fields[0], ":")

		rx, _ := strconv.ParseUint(fields[1], 10, 64)
		tx, _ := strconv.ParseUint(fields[9], 10, 64)

		stats = append(stats, InterfaceStat{
			Name:    name,
			RxBytes: rx,
			TxBytes: tx,
		})
	}

	return stats, scanner.Err()
}
