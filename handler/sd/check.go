package sd

import (
	"github.com/shirou/gopsutil/disk"
	"net/http"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"goStudyProject/util"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// HealthCheck shows "ok" as the disk usage.
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	msg := "ok"
	util.ResponseJSON(w, http.StatusOK, msg)
}

// DiskCheck checks the disk usage.
func DiskCheck(w http.ResponseWriter, r *http.Request) {
	u, _ := disk.Usage("/")

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		status = http.StatusOK
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusOK
		text = "WARNING"
	}
	msg := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %DMB (%dGB) | Used: %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	util.ResponseJSON(w, status, msg)
}

// CPUCheck check the cpu usage
func CPUCheck(w http.ResponseWriter, r *http.Request) {
	// get physicalcpu
	cores, _ := cpu.Counts(false)

	a,_ := load.Avg()
	l1 := a.Load1
	l5 := a.Load5
	l15 := a.Load15

	status := http.StatusOK
	text := "OK"

	if l5 >= float64(cores -1) {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if l5 >= float64(cores -2) {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	msg := fmt.Sprintf("%s - Load average: %.2f, %.2f, %.2f | Cores: %d", text, l1, l5, l15, cores)
	util.ResponseJSON(w, status, msg)
}


// RAMCheck checks the ram usage
func RAMCheck(w http.ResponseWriter, r *http.Request) {
	u,_ := mem.VirtualMemory()

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		status = http.StatusOK
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusOK
		text = "WARNING"
	}
	msg := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	util.ResponseJSON(w, status, msg)
}
