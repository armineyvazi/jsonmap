package dto

type LaptopDetail struct {
	Brand           string `json:"brand"`
	Model           string `json:"model"`
	Processor       string `json:"processor"`
	RamCapacity     string `json:"ram_capacity"`
	RamType         string `json:"ram_type"`
	StorageCapacity string `json:"storage_capacity"`
	BatteryStatus   string `json:"battery_status"`
}
