//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azspark

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type BatchJob.
func (b BatchJob) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "appId", b.AppID)
	populate(objectMap, "appInfo", b.AppInfo)
	populate(objectMap, "artifactId", b.ArtifactID)
	populate(objectMap, "errorInfo", b.Errors)
	populate(objectMap, "id", b.ID)
	populate(objectMap, "jobType", b.JobType)
	populate(objectMap, "livyInfo", b.LivyInfo)
	populate(objectMap, "log", b.LogLines)
	populate(objectMap, "name", b.Name)
	populate(objectMap, "pluginInfo", b.Plugin)
	populate(objectMap, "result", b.Result)
	populate(objectMap, "schedulerInfo", b.Scheduler)
	populate(objectMap, "sparkPoolName", b.SparkPoolName)
	populate(objectMap, "state", b.State)
	populate(objectMap, "submitterId", b.SubmitterID)
	populate(objectMap, "submitterName", b.SubmitterName)
	populate(objectMap, "tags", b.Tags)
	populate(objectMap, "workspaceName", b.WorkspaceName)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type BatchJobCollection.
func (b BatchJobCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "from", b.From)
	populate(objectMap, "sessions", b.Sessions)
	populate(objectMap, "total", b.Total)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type BatchJobOptions.
func (b BatchJobOptions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "archives", b.Archives)
	populate(objectMap, "args", b.Arguments)
	populate(objectMap, "artifactId", b.ArtifactID)
	populate(objectMap, "className", b.ClassName)
	populate(objectMap, "conf", b.Configuration)
	populate(objectMap, "driverCores", b.DriverCores)
	populate(objectMap, "driverMemory", b.DriverMemory)
	populate(objectMap, "executorCores", b.ExecutorCores)
	populate(objectMap, "numExecutors", b.ExecutorCount)
	populate(objectMap, "executorMemory", b.ExecutorMemory)
	populate(objectMap, "file", b.File)
	populate(objectMap, "files", b.Files)
	populate(objectMap, "jars", b.Jars)
	populate(objectMap, "name", b.Name)
	populate(objectMap, "pyFiles", b.PythonFiles)
	populate(objectMap, "tags", b.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type BatchJobState.
func (b BatchJobState) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "currentState", b.CurrentState)
	populateTimeRFC3339(objectMap, "deadAt", b.DeadAt)
	populate(objectMap, "jobCreationRequest", b.JobCreationRequest)
	populateTimeRFC3339(objectMap, "notStartedAt", b.NotStartedAt)
	populateTimeRFC3339(objectMap, "recoveringAt", b.RecoveringAt)
	populateTimeRFC3339(objectMap, "runningAt", b.RunningAt)
	populateTimeRFC3339(objectMap, "startingAt", b.StartingAt)
	populateTimeRFC3339(objectMap, "successAt", b.SuccessAt)
	populateTimeRFC3339(objectMap, "killedAt", b.TerminatedAt)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BatchJobState.
func (b *BatchJobState) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "currentState":
			err = unpopulate(val, &b.CurrentState)
			delete(rawMsg, key)
		case "deadAt":
			err = unpopulateTimeRFC3339(val, &b.DeadAt)
			delete(rawMsg, key)
		case "jobCreationRequest":
			err = unpopulate(val, &b.JobCreationRequest)
			delete(rawMsg, key)
		case "notStartedAt":
			err = unpopulateTimeRFC3339(val, &b.NotStartedAt)
			delete(rawMsg, key)
		case "recoveringAt":
			err = unpopulateTimeRFC3339(val, &b.RecoveringAt)
			delete(rawMsg, key)
		case "runningAt":
			err = unpopulateTimeRFC3339(val, &b.RunningAt)
			delete(rawMsg, key)
		case "startingAt":
			err = unpopulateTimeRFC3339(val, &b.StartingAt)
			delete(rawMsg, key)
		case "successAt":
			err = unpopulateTimeRFC3339(val, &b.SuccessAt)
			delete(rawMsg, key)
		case "killedAt":
			err = unpopulateTimeRFC3339(val, &b.TerminatedAt)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Request.
func (r Request) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "archives", r.Archives)
	populate(objectMap, "args", r.Arguments)
	populate(objectMap, "className", r.ClassName)
	populate(objectMap, "conf", r.Configuration)
	populate(objectMap, "driverCores", r.DriverCores)
	populate(objectMap, "driverMemory", r.DriverMemory)
	populate(objectMap, "executorCores", r.ExecutorCores)
	populate(objectMap, "numExecutors", r.ExecutorCount)
	populate(objectMap, "executorMemory", r.ExecutorMemory)
	populate(objectMap, "file", r.File)
	populate(objectMap, "files", r.Files)
	populate(objectMap, "jars", r.Jars)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "pyFiles", r.PythonFiles)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Scheduler.
func (s Scheduler) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "cancellationRequestedAt", s.CancellationRequestedAt)
	populate(objectMap, "currentState", s.CurrentState)
	populateTimeRFC3339(objectMap, "endedAt", s.EndedAt)
	populateTimeRFC3339(objectMap, "scheduledAt", s.ScheduledAt)
	populateTimeRFC3339(objectMap, "submittedAt", s.SubmittedAt)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Scheduler.
func (s *Scheduler) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "cancellationRequestedAt":
			err = unpopulateTimeRFC3339(val, &s.CancellationRequestedAt)
			delete(rawMsg, key)
		case "currentState":
			err = unpopulate(val, &s.CurrentState)
			delete(rawMsg, key)
		case "endedAt":
			err = unpopulateTimeRFC3339(val, &s.EndedAt)
			delete(rawMsg, key)
		case "scheduledAt":
			err = unpopulateTimeRFC3339(val, &s.ScheduledAt)
			delete(rawMsg, key)
		case "submittedAt":
			err = unpopulateTimeRFC3339(val, &s.SubmittedAt)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ServicePlugin.
func (s ServicePlugin) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "cleanupStartedAt", s.CleanupStartedAt)
	populate(objectMap, "currentState", s.CurrentState)
	populateTimeRFC3339(objectMap, "monitoringStartedAt", s.MonitoringStartedAt)
	populateTimeRFC3339(objectMap, "preparationStartedAt", s.PreparationStartedAt)
	populateTimeRFC3339(objectMap, "resourceAcquisitionStartedAt", s.ResourceAcquisitionStartedAt)
	populateTimeRFC3339(objectMap, "submissionStartedAt", s.SubmissionStartedAt)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServicePlugin.
func (s *ServicePlugin) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "cleanupStartedAt":
			err = unpopulateTimeRFC3339(val, &s.CleanupStartedAt)
			delete(rawMsg, key)
		case "currentState":
			err = unpopulate(val, &s.CurrentState)
			delete(rawMsg, key)
		case "monitoringStartedAt":
			err = unpopulateTimeRFC3339(val, &s.MonitoringStartedAt)
			delete(rawMsg, key)
		case "preparationStartedAt":
			err = unpopulateTimeRFC3339(val, &s.PreparationStartedAt)
			delete(rawMsg, key)
		case "resourceAcquisitionStartedAt":
			err = unpopulateTimeRFC3339(val, &s.ResourceAcquisitionStartedAt)
			delete(rawMsg, key)
		case "submissionStartedAt":
			err = unpopulateTimeRFC3339(val, &s.SubmissionStartedAt)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Session.
func (s Session) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "appId", s.AppID)
	populate(objectMap, "appInfo", s.AppInfo)
	populate(objectMap, "artifactId", s.ArtifactID)
	populate(objectMap, "errorInfo", s.Errors)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "jobType", s.JobType)
	populate(objectMap, "livyInfo", s.LivyInfo)
	populate(objectMap, "log", s.LogLines)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "pluginInfo", s.Plugin)
	populate(objectMap, "result", s.Result)
	populate(objectMap, "schedulerInfo", s.Scheduler)
	populate(objectMap, "sparkPoolName", s.SparkPoolName)
	populate(objectMap, "state", s.State)
	populate(objectMap, "submitterId", s.SubmitterID)
	populate(objectMap, "submitterName", s.SubmitterName)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "workspaceName", s.WorkspaceName)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SessionCollection.
func (s SessionCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "from", s.From)
	populate(objectMap, "sessions", s.Sessions)
	populate(objectMap, "total", s.Total)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SessionOptions.
func (s SessionOptions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "archives", s.Archives)
	populate(objectMap, "args", s.Arguments)
	populate(objectMap, "artifactId", s.ArtifactID)
	populate(objectMap, "className", s.ClassName)
	populate(objectMap, "conf", s.Configuration)
	populate(objectMap, "driverCores", s.DriverCores)
	populate(objectMap, "driverMemory", s.DriverMemory)
	populate(objectMap, "executorCores", s.ExecutorCores)
	populate(objectMap, "numExecutors", s.ExecutorCount)
	populate(objectMap, "executorMemory", s.ExecutorMemory)
	populate(objectMap, "file", s.File)
	populate(objectMap, "files", s.Files)
	populate(objectMap, "jars", s.Jars)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "pyFiles", s.PythonFiles)
	populate(objectMap, "tags", s.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SessionState.
func (s SessionState) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "busyAt", s.BusyAt)
	populate(objectMap, "currentState", s.CurrentState)
	populateTimeRFC3339(objectMap, "deadAt", s.DeadAt)
	populateTimeRFC3339(objectMap, "errorAt", s.ErrorAt)
	populateTimeRFC3339(objectMap, "idleAt", s.IdleAt)
	populate(objectMap, "jobCreationRequest", s.JobCreationRequest)
	populateTimeRFC3339(objectMap, "notStartedAt", s.NotStartedAt)
	populateTimeRFC3339(objectMap, "recoveringAt", s.RecoveringAt)
	populateTimeRFC3339(objectMap, "shuttingDownAt", s.ShuttingDownAt)
	populateTimeRFC3339(objectMap, "startingAt", s.StartingAt)
	populateTimeRFC3339(objectMap, "killedAt", s.TerminatedAt)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SessionState.
func (s *SessionState) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "busyAt":
			err = unpopulateTimeRFC3339(val, &s.BusyAt)
			delete(rawMsg, key)
		case "currentState":
			err = unpopulate(val, &s.CurrentState)
			delete(rawMsg, key)
		case "deadAt":
			err = unpopulateTimeRFC3339(val, &s.DeadAt)
			delete(rawMsg, key)
		case "errorAt":
			err = unpopulateTimeRFC3339(val, &s.ErrorAt)
			delete(rawMsg, key)
		case "idleAt":
			err = unpopulateTimeRFC3339(val, &s.IdleAt)
			delete(rawMsg, key)
		case "jobCreationRequest":
			err = unpopulate(val, &s.JobCreationRequest)
			delete(rawMsg, key)
		case "notStartedAt":
			err = unpopulateTimeRFC3339(val, &s.NotStartedAt)
			delete(rawMsg, key)
		case "recoveringAt":
			err = unpopulateTimeRFC3339(val, &s.RecoveringAt)
			delete(rawMsg, key)
		case "shuttingDownAt":
			err = unpopulateTimeRFC3339(val, &s.ShuttingDownAt)
			delete(rawMsg, key)
		case "startingAt":
			err = unpopulateTimeRFC3339(val, &s.StartingAt)
			delete(rawMsg, key)
		case "killedAt":
			err = unpopulateTimeRFC3339(val, &s.TerminatedAt)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StatementCollection.
func (s StatementCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "statements", s.Statements)
	populate(objectMap, "total_statements", s.Total)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type StatementOutput.
func (s StatementOutput) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "data", &s.Data)
	populate(objectMap, "ename", s.ErrorName)
	populate(objectMap, "evalue", s.ErrorValue)
	populate(objectMap, "execution_count", s.ExecutionCount)
	populate(objectMap, "status", s.Status)
	populate(objectMap, "traceback", s.Traceback)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}