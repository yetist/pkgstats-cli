package system

import (
	"testing"
)

func TestGetCpuArchitecture(t *testing.T) {
	system := System{}

	cpuArch, err := system.GetCpuArchitecture()

	if err != nil {
		t.Error(err)
	}
	if cpuArch != "riscv64" {
		t.Error(cpuArch)
	}
}
