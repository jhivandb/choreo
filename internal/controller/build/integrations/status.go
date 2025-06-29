// Copyright 2025 The OpenChoreo Authors
// SPDX-License-Identifier: Apache-2.0

package integrations

type BuildWorkflowStep string

const (
	CloneStep BuildWorkflowStep = "clone-step"
	BuildStep BuildWorkflowStep = "build-step"
	PushStep  BuildWorkflowStep = "push-step"
)

type StepPhase string

// Workflow and node statuses
const (
	Running   StepPhase = "Running"
	Succeeded StepPhase = "Succeeded"
	Failed    StepPhase = "Failed"
)
