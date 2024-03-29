// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the License). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the license file accompanying this file. This file is distributed
// on an AS IS BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package engine

import "github.com/blox/blox/daemon-scheduler/pkg/types"

type EventType string

const (
	StartDeploymentEventType            EventType = "StartDeploymentEvent"
	StopTasksEventType                            = "StopTasksEvent"
	UpdateInProgressDeploymentEventType           = "UpdateInProgressDeploymentEvent"
)

type Event interface {
	//GetType returns event-type
	GetType() EventType
}

// StartDeploymentEvent is message used to notify actors to perform a deployment using environment
type StartDeploymentEvent struct {
	Instances   []*string
	Environment types.Environment
}

func (e StartDeploymentEvent) GetType() EventType {
	return StartDeploymentEventType
}

// StopTasksEvent is message used to notify actors to stop tasks
type StopTasksEvent struct {
	Cluster     string
	Tasks       []string
	Environment types.Environment
}

func (e StopTasksEvent) GetType() EventType {
	return StopTasksEventType
}

type UpdateInProgressDeploymentEvent struct {
	Environment types.Environment
}

func (e UpdateInProgressDeploymentEvent) GetType() EventType {
	return UpdateInProgressDeploymentEventType
}
