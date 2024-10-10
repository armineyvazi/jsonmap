package utils

import "github.com/armineyvazi/jsonmap/dto"

func CheckLaptopIsNotEmpty(details []dto.LaptopDetail) bool {
	for _, detail := range details {
		if detail.Brand == "" ||
			detail.Model == "" ||
			detail.BatteryStatus == "" ||
			detail.RamCapacity == "" ||
			detail.StorageCapacity == "" ||
			detail.RamType == "" ||
			detail.Processor == "" {
			return false
		}
	}
	return true
}
