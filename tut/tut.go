package tut

import (
	log "github.com/sirupsen/logrus"
)

// Calculate returns x + 2.
func Calculate(x int) (result int) {
	log.Info("Run Function Calculate")
	result = x + 2
	return result
}
