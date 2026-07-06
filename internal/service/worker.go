package service

import "log"

func (s *Service) DoWork() {

	stats, err := s.collector.Collect()
	if err != nil {
		log.Println(err)
		return
	}

	for _, iface := range stats {
		log.Printf(
			"%-8s RX=%12d TX=%12d",
			iface.Name,
			iface.RxBytes,
			iface.TxBytes,
		)
	}

	log.Println("----------------------------")
}