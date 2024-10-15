package metrics

import (
	"fmt"

	"github.com/onsi/gomega/gmeasure"
	"github.com/prometheus/common/model"
)

func ContainerCPU(job string, duration model.Duration, annotation gmeasure.Annotation) Measurement {
	return Measurement{
		Name: fmt.Sprintf("%s Sum of Container CPU Usage", job),
		Query: fmt.Sprintf(
			`sum(avg_over_time(pod:container_cpu_usage:sum{pod=~".*%s.*"}[%s])) * %d`,
			job, duration, CoresToMillicores,
		),
		Unit:       MillicoresUnit,
		Annotation: annotation,
	}
}

func ContainerMemoryWorkingSetBytes(job string, duration model.Duration, annotation gmeasure.Annotation) Measurement {
	return Measurement{
		Name: fmt.Sprintf("%s Sum of Container WorkingSet Memory", job),
		Query: fmt.Sprintf(
			`sum(avg_over_time(container_memory_working_set_bytes{pod=~".*%s.*", container=""}[%s]) / %d)`,
			job, duration, BytesToGigabytesMultiplier,
		),
		Unit:       GigabytesUnit,
		Annotation: annotation,
	}
}

func ContainerGoMemstatsHeapInuse(job string, _ model.Duration, annotation gmeasure.Annotation) Measurement {
	return Measurement{
		Name: fmt.Sprintf("%s Sum of Container Go Memstats Heap Inuse", job),
		Query: fmt.Sprintf(
			`sum(go_memstats_heap_inuse_bytes{pod=~".*%s.*"}) / %d`,
			job, BytesToGigabytesMultiplier,
		),
		Unit:       GigabytesUnit,
		Annotation: annotation,
	}
}

func PersistentVolumeUsedBytes(job string, duration model.Duration, annotation gmeasure.Annotation) Measurement {
	return Measurement{
		Name: fmt.Sprintf("%s Sum of Persistent Volume Used Bytes", job),
		Query: fmt.Sprintf(
			`sum(avg_over_time(kubelet_volume_stats_used_bytes{persistentvolumeclaim=~".*%s.*"}[%s]) / %d)`,
			job, duration, BytesToGigabytesMultiplier,
		),
		Unit:       GigabytesUnit,
		Annotation: annotation,
	}
}
