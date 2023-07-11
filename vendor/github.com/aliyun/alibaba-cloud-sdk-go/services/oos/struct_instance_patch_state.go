package oos

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// InstancePatchState is a nested struct in oos response
type InstancePatchState struct {
	InstanceId                  string `json:"InstanceId" xml:"InstanceId"`
	PatchGroup                  string `json:"PatchGroup" xml:"PatchGroup"`
	BaselineId                  string `json:"BaselineId" xml:"BaselineId"`
	OwnerInformation            string `json:"OwnerInformation" xml:"OwnerInformation"`
	InstalledCount              string `json:"InstalledCount" xml:"InstalledCount"`
	InstalledOtherCount         string `json:"InstalledOtherCount" xml:"InstalledOtherCount"`
	InstalledPendingRebootCount string `json:"InstalledPendingRebootCount" xml:"InstalledPendingRebootCount"`
	InstalledRejectedCount      string `json:"InstalledRejectedCount" xml:"InstalledRejectedCount"`
	MissingCount                string `json:"MissingCount" xml:"MissingCount"`
	FailedCount                 string `json:"FailedCount" xml:"FailedCount"`
	OperationType               string `json:"OperationType" xml:"OperationType"`
	OperationStartTime          string `json:"OperationStartTime" xml:"OperationStartTime"`
	OperationEndTime            string `json:"OperationEndTime" xml:"OperationEndTime"`
}
