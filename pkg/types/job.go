// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"reflect"

	"k8s.io/utils/set"

	"github.com/SlinkyProject/slurm-client/pkg/object"
)

const (
	ObjectTypeJobInfo = "jobInfo"
)

// Represents the JobInfo Job State
// +enum
type JobInfoJobState string

const (
	// Base States: mutually exclusive, must only have one
	JobInfoJobStateBOOTFAIL    JobInfoJobState = "BOOT_FAIL"
	JobInfoJobStateCOMPLETED   JobInfoJobState = "COMPLETED"
	JobInfoJobStateCANCELLED   JobInfoJobState = "CANCELLED"
	JobInfoJobStateDEADLINE    JobInfoJobState = "DEADLINE"
	JobInfoJobStateFAILED      JobInfoJobState = "FAILED"
	JobInfoJobStateNODEFAIL    JobInfoJobState = "NODE_FAIL"
	JobInfoJobStateOUTOFMEMORY JobInfoJobState = "OUT_OF_MEMORY"
	JobInfoJobStatePENDING     JobInfoJobState = "PENDING"
	JobInfoJobStatePREEMPTED   JobInfoJobState = "PREEMPTED"
	JobInfoJobStateRUNNING     JobInfoJobState = "RUNNING"
	JobInfoJobStateSUSPENDED   JobInfoJobState = "SUSPENDED"
	JobInfoJobStateTIMEOUT     JobInfoJobState = "TIMEOUT"

	// Flag States: not mutually exclusive, can have zero or more
	JobInfoJobStateCOMPLETING   JobInfoJobState = "COMPLETING"
	JobInfoJobStateCONFIGURING  JobInfoJobState = "CONFIGURING"
	JobInfoJobStateLAUNCHFAILED JobInfoJobState = "LAUNCH_FAILED"
	JobInfoJobStatePOWERUPNODE  JobInfoJobState = "POWER_UP_NODE"
	JobInfoJobStateRECONFIGFAIL JobInfoJobState = "RECONFIG_FAIL"
	JobInfoJobStateREQUEUED     JobInfoJobState = "REQUEUED"
	JobInfoJobStateREQUEUEFED   JobInfoJobState = "REQUEUE_FED"
	JobInfoJobStateREQUEUEHOLD  JobInfoJobState = "REQUEUE_HOLD"
	JobInfoJobStateRESIZING     JobInfoJobState = "RESIZING"
	JobInfoJobStateRESVDELHOLD  JobInfoJobState = "RESV_DEL_HOLD"
	JobInfoJobStateREVOKED      JobInfoJobState = "REVOKED"
	JobInfoJobStateSIGNALING    JobInfoJobState = "SIGNALING"
	JobInfoJobStateSPECIALEXIT  JobInfoJobState = "SPECIAL_EXIT"
	JobInfoJobStateSTAGEOUT     JobInfoJobState = "STAGE_OUT"
	JobInfoJobStateSTOPPED      JobInfoJobState = "STOPPED"
	JobInfoJobStateUPDATEDB     JobInfoJobState = "UPDATE_DB"
)

// JobInfo represents a Slurm job.
type JobInfo struct {
	JobId        int32
	Partition    string
	JobState     set.Set[JobInfoJobState]
	Uid          int32
	UserName     string
	Cpus         int64
	NodeCount    int64
	Hold         bool
	EligibleTime int64
}

//+kubebuilder:object:root=true
//+kubebuilder:object:generate=false

// GetKey implements Object.
func (j *JobInfo) GetKey() object.ObjectKey {
	return object.ObjectKey(j.JobId)
}

// GetType implements Object.
func (j *JobInfo) GetType() object.ObjectType {
	return ObjectTypeJobInfo
}

// DeepEqualObject implements Object.
func (j *JobInfo) DeepEqualObject(object object.Object) bool {
	return reflect.DeepEqual(j, object)
}

// DeepCopyObject implements Object.
func (j *JobInfo) DeepCopyObject() object.Object {
	return j.DeepCopy()
}

// JobInfoList represents a list of Slurm jobs.
type JobInfoList struct {
	Items []JobInfo
}

// GetType implements ObjectList.
func (j *JobInfoList) GetType() object.ObjectType {
	return ObjectTypeJobInfo
}

// GetItems implements ObjectList.
func (j *JobInfoList) GetItems() []object.Object {
	jobInfos := make([]object.Object, 0)
	for _, jobInfo := range j.Items {
		jobInfos = append(jobInfos, jobInfo.DeepCopyObject())
	}
	return jobInfos
}

// Size implements ObjectList.
func (j *JobInfoList) Size() int {
	return len(j.Items)
}

// AppendItem implements ObjectList.
func (j *JobInfoList) AppendItem(object object.Object) {
	jobInfo := object.(*JobInfo)
	jobInfo = jobInfo.DeepCopy()
	j.Items = append(j.Items, *jobInfo)
}

// DeepCopyObjectList implements ObjectList.
func (j *JobInfoList) DeepCopyObjectList() object.ObjectList {
	return j.DeepCopy()
}
