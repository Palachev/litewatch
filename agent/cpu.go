package main

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

// GetCPU возвращает % загрузки CPU (усреднённо за 1 секунду)
func GetCPU() (float64, error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	if len(percent) == 0 {
		return 0, nil
	}
	return percent[0], nil
}
