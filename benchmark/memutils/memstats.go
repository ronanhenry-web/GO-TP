package memutils

import (
	"fmt"
	"runtime"
)

func DisplayMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Allocated = %.2f KiB \t", float64(m.Alloc)/1024)
	fmt.Printf("Total Allocated = %.2f KiB \t", float64(m.TotalAlloc)/1024)
	fmt.Printf("System = %.2f KiB \t", float64(m.Sys)/1024)
	fmt.Printf("NumGC = %d \n", m.NumGC)

}
