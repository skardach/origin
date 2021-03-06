package v1beta1

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
)

// CustomPodDeploymentStrategy describes a deployment carried out by a custom pod.
type CustomPodDeploymentStrategy struct {
	Image       string       `json:"image,omitempty" yaml:"image,omitempty"`
	Environment []api.EnvVar `json:"environment,omitempty" yaml:"environment,omitempty"`
}

// DeploymentStrategy describes how to perform a deployment.
type DeploymentStrategy struct {
	Type      string                       `json:"type,omitempty" yaml:"type,omitempty"`
	CustomPod *CustomPodDeploymentStrategy `json:"customPod,omitempty" yaml:"customPod,omitempty"`
}

// DeploymentTemplate contains all the necessary information to create a Deployment from a
// DeploymentStrategy.
type DeploymentTemplate struct {
	Strategy           DeploymentStrategy             `json:"strategy,omitempty" yaml:"strategy,omitempty"`
	ControllerTemplate api.ReplicationControllerState `json:"controllerTemplate,omitempty" yaml:"controllerTemplate,omitempty"`
}

// DeploymentState decribes the possible states a Deployment can be in.
type DeploymentState string

const (
	DeploymentNew      DeploymentState = "new"
	DeploymentPending  DeploymentState = "pending"
	DeploymentRunning  DeploymentState = "running"
	DeploymentComplete DeploymentState = "complete"
	DeploymentFailed   DeploymentState = "failed"
)

// A Deployment represents a single unique realization of a DeploymentConfig.
type Deployment struct {
	api.JSONBase       `json:",inline" yaml:",inline"`
	Labels             map[string]string              `json:"labels,omitempty" yaml:"labels,omitempty"`
	Strategy           DeploymentStrategy             `json:"strategy,omitempty" yaml:"strategy,omitempty"`
	ControllerTemplate api.ReplicationControllerState `json:"controllerTemplate,omitempty" yaml:"controllerTemplate,omitempty"`
	State              DeploymentState                `json:"state,omitempty" yaml:"state,omitempty"`
	ConfigID           string                         `json:"configId,omitempty" yaml:"configId,omitempty"`
}

// DeploymentTriggerPolicy describes the possible triggers that result in a new Deployment.
type DeploymentTriggerPolicy struct {
	Type DeploymentTriggerType `json:"type,omitempty" yaml:"type,omitempty"`
}

type DeploymentTriggerType string

const (
	DeploymentTriggerOnImageChange  DeploymentTriggerType = "image-change"
	DeploymentTriggerOnConfigChange DeploymentTriggerType = "config-change"
	DeploymentTriggerManual         DeploymentTriggerType = "manual"
)

// DeploymentConfig represents a configuration for a single deployment of a replication controller:
// what the template for the deployment, how new deployments are triggered, what the current
// deployed state is.
type DeploymentConfig struct {
	api.JSONBase  `json:",inline" yaml:",inline"`
	Labels        map[string]string              `json:"labels,omitempty" yaml:"labels,omitempty"`
	TriggerPolicy DeploymentTriggerPolicy        `json:"triggerPolicy,omitempty" yaml:"triggerPolicy,omitempty"`
	Template      DeploymentTemplate             `json:"template,omitempty" yaml:"template,omitempty"`
	CurrentState  api.ReplicationControllerState `json:"currentState" yaml:"currentState,omitempty"`
}

// A DeploymentConfigList is a collection of deployment configs
type DeploymentConfigList struct {
	api.JSONBase `json:",inline" yaml:",inline"`
	Items        []DeploymentConfig `json:"items,omitempty" yaml:"items,omitempty"`
}

// A DeploymentList is a collection of deployments.
type DeploymentList struct {
	api.JSONBase `json:",inline" yaml:",inline"`
	Items        []Deployment `json:"items,omitempty" yaml:"items,omitempty"`
}
