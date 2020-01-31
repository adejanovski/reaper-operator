// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthProvider) DeepCopyInto(out *AuthProvider) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthProvider.
func (in *AuthProvider) DeepCopy() *AuthProvider {
	if in == nil {
		return nil
	}
	out := new(AuthProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScheduling) DeepCopyInto(out *AutoScheduling) {
	*out = *in
	if in.ExcludedKeyspaces != nil {
		in, out := &in.ExcludedKeyspaces, &out.ExcludedKeyspaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScheduling.
func (in *AutoScheduling) DeepCopy() *AutoScheduling {
	if in == nil {
		return nil
	}
	out := new(AutoScheduling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraBackend) DeepCopyInto(out *CassandraBackend) {
	*out = *in
	if in.ContactPoints != nil {
		in, out := &in.ContactPoints, &out.ContactPoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Replication.DeepCopyInto(&out.Replication)
	out.AuthProvider = in.AuthProvider
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraBackend.
func (in *CassandraBackend) DeepCopy() *CassandraBackend {
	if in == nil {
		return nil
	}
	out := new(CassandraBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Reaper) DeepCopyInto(out *Reaper) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Reaper.
func (in *Reaper) DeepCopy() *Reaper {
	if in == nil {
		return nil
	}
	out := new(Reaper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Reaper) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReaperCondition) DeepCopyInto(out *ReaperCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReaperCondition.
func (in *ReaperCondition) DeepCopy() *ReaperCondition {
	if in == nil {
		return nil
	}
	out := new(ReaperCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReaperList) DeepCopyInto(out *ReaperList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Reaper, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReaperList.
func (in *ReaperList) DeepCopy() *ReaperList {
	if in == nil {
		return nil
	}
	out := new(ReaperList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReaperList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReaperSpec) DeepCopyInto(out *ReaperSpec) {
	*out = *in
	in.ServerConfig.DeepCopyInto(&out.ServerConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReaperSpec.
func (in *ReaperSpec) DeepCopy() *ReaperSpec {
	if in == nil {
		return nil
	}
	out := new(ReaperSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReaperStatus) DeepCopyInto(out *ReaperStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ReaperCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReaperStatus.
func (in *ReaperStatus) DeepCopy() *ReaperStatus {
	if in == nil {
		return nil
	}
	out := new(ReaperStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationConfig) DeepCopyInto(out *ReplicationConfig) {
	*out = *in
	if in.SimpleStrategy != nil {
		in, out := &in.SimpleStrategy, &out.SimpleStrategy
		*out = new(int32)
		**out = **in
	}
	if in.NetworkTopologyStrategy != nil {
		in, out := &in.NetworkTopologyStrategy, &out.NetworkTopologyStrategy
		*out = new(map[string]int32)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]int32, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationConfig.
func (in *ReplicationConfig) DeepCopy() *ReplicationConfig {
	if in == nil {
		return nil
	}
	out := new(ReplicationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerConfig) DeepCopyInto(out *ServerConfig) {
	*out = *in
	in.AutoScheduling.DeepCopyInto(&out.AutoScheduling)
	if in.CassandraBackend != nil {
		in, out := &in.CassandraBackend, &out.CassandraBackend
		*out = new(CassandraBackend)
		(*in).DeepCopyInto(*out)
	}
	if in.HangingRepairTimeoutMins != nil {
		in, out := &in.HangingRepairTimeoutMins, &out.HangingRepairTimeoutMins
		*out = new(int32)
		**out = **in
	}
	if in.RepairRunThreadCount != nil {
		in, out := &in.RepairRunThreadCount, &out.RepairRunThreadCount
		*out = new(int32)
		**out = **in
	}
	if in.ScheduleDaysBetween != nil {
		in, out := &in.ScheduleDaysBetween, &out.ScheduleDaysBetween
		*out = new(int32)
		**out = **in
	}
	if in.EnableCrossOrigin != nil {
		in, out := &in.EnableCrossOrigin, &out.EnableCrossOrigin
		*out = new(bool)
		**out = **in
	}
	if in.EnableDynamicSeedList != nil {
		in, out := &in.EnableDynamicSeedList, &out.EnableDynamicSeedList
		*out = new(bool)
		**out = **in
	}
	if in.JmxConnectionTimeoutInSeconds != nil {
		in, out := &in.JmxConnectionTimeoutInSeconds, &out.JmxConnectionTimeoutInSeconds
		*out = new(int32)
		**out = **in
	}
	if in.SegmentCountPerNode != nil {
		in, out := &in.SegmentCountPerNode, &out.SegmentCountPerNode
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerConfig.
func (in *ServerConfig) DeepCopy() *ServerConfig {
	if in == nil {
		return nil
	}
	out := new(ServerConfig)
	in.DeepCopyInto(out)
	return out
}
