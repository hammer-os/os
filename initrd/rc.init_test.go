package main

import (
	"reflect"
	"strings"
	"testing"
)

var (
	procCgroups = `#subsys_name	hierarchy	num_cgroups	enabled
cpuset	10	2	1
cpu	2	62	1
cpuacct	2	62	0
blkio	11	62	0
memory	4	164	1
devices	5	63	1
freezer	7	2	1
net_cls	3	2	1
perf_event	6	2	1
net_prio	3	2	1
hugetlb	8	2	1
pids	9	63	1`

	wantSubsystems = []string{
		"cpuset",
		"cpu",
		"memory",
		"devices",
		"freezer",
		"net_cls",
		"perf_event",
		"net_prio",
		"hugetlb",
		"pids",
	}
)

func TestCgroupSubsystems(t *testing.T) {
	r := strings.NewReader(procCgroups)
	subsystems, err := cgroupSubsystems(r)
	if err != nil {
		t.Fatalf("find subsystems failed: %v", err)
	}

	if !reflect.DeepEqual(subsystems, wantSubsystems) {
		t.Fatalf("unexpected subsystems\n%v\n%v",
			wantSubsystems, subsystems)
	}
}
