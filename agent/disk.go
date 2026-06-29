package main

import "github.com/shirou/gopsutil/v3/disk"

// GetDisk возвращает % заполнения указанного раздела
func GetDisk(path string) (float64, error) {
	d, err := disk.Usage(path)
	if err != nil {
		return 0, err
	}
	return d.UsedPercent, nil
}
