package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	PodMissingMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: argoNamespace,
			Subsystem: workflowsSubsystem,
			Name:      "pod_missing",
			Help:      "Incidents of pod missing. https://argoproj.github.io/argo/metrics/#pod_missing",
		},
		[]string{"recently_started", "node_phase"},
	)
)
