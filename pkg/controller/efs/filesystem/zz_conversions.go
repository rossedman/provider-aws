/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package filesystem

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/efs"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/crossplane/provider-aws/apis/efs/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateDescribeFileSystemsInput returns input for read
// operation.
func GenerateDescribeFileSystemsInput(cr *svcapitypes.FileSystem) *svcsdk.DescribeFileSystemsInput {
	res := &svcsdk.DescribeFileSystemsInput{}

	if cr.Status.AtProvider.CreationToken != nil {
		res.SetCreationToken(*cr.Status.AtProvider.CreationToken)
	}
	if cr.Status.AtProvider.FileSystemID != nil {
		res.SetFileSystemId(*cr.Status.AtProvider.FileSystemID)
	}

	return res
}

// GenerateFileSystem returns the current state in the form of *svcapitypes.FileSystem.
func GenerateFileSystem(resp *svcsdk.DescribeFileSystemsOutput) *svcapitypes.FileSystem {
	cr := &svcapitypes.FileSystem{}

	found := false
	for _, elem := range resp.FileSystems {
		if elem.AvailabilityZoneId != nil {
			cr.Status.AtProvider.AvailabilityZoneID = elem.AvailabilityZoneId
		} else {
			cr.Status.AtProvider.AvailabilityZoneID = nil
		}
		if elem.AvailabilityZoneName != nil {
			cr.Spec.ForProvider.AvailabilityZoneName = elem.AvailabilityZoneName
		} else {
			cr.Spec.ForProvider.AvailabilityZoneName = nil
		}
		if elem.CreationTime != nil {
			cr.Status.AtProvider.CreationTime = &metav1.Time{*elem.CreationTime}
		} else {
			cr.Status.AtProvider.CreationTime = nil
		}
		if elem.CreationToken != nil {
			cr.Status.AtProvider.CreationToken = elem.CreationToken
		} else {
			cr.Status.AtProvider.CreationToken = nil
		}
		if elem.Encrypted != nil {
			cr.Spec.ForProvider.Encrypted = elem.Encrypted
		} else {
			cr.Spec.ForProvider.Encrypted = nil
		}
		if elem.FileSystemArn != nil {
			cr.Status.AtProvider.FileSystemARN = elem.FileSystemArn
		} else {
			cr.Status.AtProvider.FileSystemARN = nil
		}
		if elem.FileSystemId != nil {
			cr.Status.AtProvider.FileSystemID = elem.FileSystemId
		} else {
			cr.Status.AtProvider.FileSystemID = nil
		}
		if elem.KmsKeyId != nil {
			cr.Spec.ForProvider.KMSKeyID = elem.KmsKeyId
		} else {
			cr.Spec.ForProvider.KMSKeyID = nil
		}
		if elem.LifeCycleState != nil {
			cr.Status.AtProvider.LifeCycleState = elem.LifeCycleState
		} else {
			cr.Status.AtProvider.LifeCycleState = nil
		}
		if elem.Name != nil {
			cr.Status.AtProvider.Name = elem.Name
		} else {
			cr.Status.AtProvider.Name = nil
		}
		if elem.NumberOfMountTargets != nil {
			cr.Status.AtProvider.NumberOfMountTargets = elem.NumberOfMountTargets
		} else {
			cr.Status.AtProvider.NumberOfMountTargets = nil
		}
		if elem.OwnerId != nil {
			cr.Status.AtProvider.OwnerID = elem.OwnerId
		} else {
			cr.Status.AtProvider.OwnerID = nil
		}
		if elem.PerformanceMode != nil {
			cr.Spec.ForProvider.PerformanceMode = elem.PerformanceMode
		} else {
			cr.Spec.ForProvider.PerformanceMode = nil
		}
		if elem.SizeInBytes != nil {
			f13 := &svcapitypes.FileSystemSize{}
			if elem.SizeInBytes.Timestamp != nil {
				f13.Timestamp = &metav1.Time{*elem.SizeInBytes.Timestamp}
			}
			if elem.SizeInBytes.Value != nil {
				f13.Value = elem.SizeInBytes.Value
			}
			if elem.SizeInBytes.ValueInIA != nil {
				f13.ValueInIA = elem.SizeInBytes.ValueInIA
			}
			if elem.SizeInBytes.ValueInStandard != nil {
				f13.ValueInStandard = elem.SizeInBytes.ValueInStandard
			}
			cr.Status.AtProvider.SizeInBytes = f13
		} else {
			cr.Status.AtProvider.SizeInBytes = nil
		}
		if elem.Tags != nil {
			f14 := []*svcapitypes.Tag{}
			for _, f14iter := range elem.Tags {
				f14elem := &svcapitypes.Tag{}
				if f14iter.Key != nil {
					f14elem.Key = f14iter.Key
				}
				if f14iter.Value != nil {
					f14elem.Value = f14iter.Value
				}
				f14 = append(f14, f14elem)
			}
			cr.Spec.ForProvider.Tags = f14
		} else {
			cr.Spec.ForProvider.Tags = nil
		}
		if elem.ThroughputMode != nil {
			cr.Spec.ForProvider.ThroughputMode = elem.ThroughputMode
		} else {
			cr.Spec.ForProvider.ThroughputMode = nil
		}
		found = true
		break
	}
	if !found {
		return cr
	}

	return cr
}

// GenerateCreateFileSystemInput returns a create input.
func GenerateCreateFileSystemInput(cr *svcapitypes.FileSystem) *svcsdk.CreateFileSystemInput {
	res := &svcsdk.CreateFileSystemInput{}

	if cr.Spec.ForProvider.AvailabilityZoneName != nil {
		res.SetAvailabilityZoneName(*cr.Spec.ForProvider.AvailabilityZoneName)
	}
	if cr.Spec.ForProvider.Backup != nil {
		res.SetBackup(*cr.Spec.ForProvider.Backup)
	}
	if cr.Spec.ForProvider.Encrypted != nil {
		res.SetEncrypted(*cr.Spec.ForProvider.Encrypted)
	}
	if cr.Spec.ForProvider.KMSKeyID != nil {
		res.SetKmsKeyId(*cr.Spec.ForProvider.KMSKeyID)
	}
	if cr.Spec.ForProvider.PerformanceMode != nil {
		res.SetPerformanceMode(*cr.Spec.ForProvider.PerformanceMode)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f5 := []*svcsdk.Tag{}
		for _, f5iter := range cr.Spec.ForProvider.Tags {
			f5elem := &svcsdk.Tag{}
			if f5iter.Key != nil {
				f5elem.SetKey(*f5iter.Key)
			}
			if f5iter.Value != nil {
				f5elem.SetValue(*f5iter.Value)
			}
			f5 = append(f5, f5elem)
		}
		res.SetTags(f5)
	}
	if cr.Spec.ForProvider.ThroughputMode != nil {
		res.SetThroughputMode(*cr.Spec.ForProvider.ThroughputMode)
	}

	return res
}

// GenerateUpdateFileSystemInput returns an update input.
func GenerateUpdateFileSystemInput(cr *svcapitypes.FileSystem) *svcsdk.UpdateFileSystemInput {
	res := &svcsdk.UpdateFileSystemInput{}

	if cr.Status.AtProvider.FileSystemID != nil {
		res.SetFileSystemId(*cr.Status.AtProvider.FileSystemID)
	}
	if cr.Spec.ForProvider.ThroughputMode != nil {
		res.SetThroughputMode(*cr.Spec.ForProvider.ThroughputMode)
	}

	return res
}

// GenerateDeleteFileSystemInput returns a deletion input.
func GenerateDeleteFileSystemInput(cr *svcapitypes.FileSystem) *svcsdk.DeleteFileSystemInput {
	res := &svcsdk.DeleteFileSystemInput{}

	if cr.Status.AtProvider.FileSystemID != nil {
		res.SetFileSystemId(*cr.Status.AtProvider.FileSystemID)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "FileSystemNotFound"
}
