package framework

import "github.com/onsi/ginkgo/v2"

var (
	LabelCAPI                  = ginkgo.Label("capi")
	LabelDisruptive            = ginkgo.Label("disruptive")
	LabelAutoscaler            = ginkgo.Label("autoscaler")
	LabelOperators             = ginkgo.Label("operators")
	LabelPeriodic              = ginkgo.Label("periodic")
	LabelSpot                  = ginkgo.Label("spot-instances")
	LabelMachines              = ginkgo.Label("machines")
	LabelMachineHealthChecks   = ginkgo.Label("machine-health-checks")
	LabelCloudProviderSpecific = ginkgo.Label("cloud-provider-specific")
	LabelProviderAWS           = ginkgo.Label("AWS")
	LabelQeJobOnly             = ginkgo.Label("QeJobOnly")
)
