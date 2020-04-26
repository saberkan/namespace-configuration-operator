// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	apis "github.com/redhat-cop/operator-utils/pkg/util/apis"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupConfig) DeepCopyInto(out *GroupConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupConfig.
func (in *GroupConfig) DeepCopy() *GroupConfig {
	if in == nil {
		return nil
	}
	out := new(GroupConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GroupConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupConfigList) DeepCopyInto(out *GroupConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GroupConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupConfigList.
func (in *GroupConfigList) DeepCopy() *GroupConfigList {
	if in == nil {
		return nil
	}
	out := new(GroupConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GroupConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupConfigSpec) DeepCopyInto(out *GroupConfigSpec) {
	*out = *in
	in.LabelSelector.DeepCopyInto(&out.LabelSelector)
	in.AnnotationSelector.DeepCopyInto(&out.AnnotationSelector)
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make([]apis.LockedResourceTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupConfigSpec.
func (in *GroupConfigSpec) DeepCopy() *GroupConfigSpec {
	if in == nil {
		return nil
	}
	out := new(GroupConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupConfigStatus) DeepCopyInto(out *GroupConfigStatus) {
	*out = *in
	in.EnforcingReconcileStatus.DeepCopyInto(&out.EnforcingReconcileStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupConfigStatus.
func (in *GroupConfigStatus) DeepCopy() *GroupConfigStatus {
	if in == nil {
		return nil
	}
	out := new(GroupConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceConfig) DeepCopyInto(out *NamespaceConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceConfig.
func (in *NamespaceConfig) DeepCopy() *NamespaceConfig {
	if in == nil {
		return nil
	}
	out := new(NamespaceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespaceConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceConfigList) DeepCopyInto(out *NamespaceConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NamespaceConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceConfigList.
func (in *NamespaceConfigList) DeepCopy() *NamespaceConfigList {
	if in == nil {
		return nil
	}
	out := new(NamespaceConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespaceConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceConfigSpec) DeepCopyInto(out *NamespaceConfigSpec) {
	*out = *in
	in.LabelSelector.DeepCopyInto(&out.LabelSelector)
	in.AnnotationSelector.DeepCopyInto(&out.AnnotationSelector)
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make([]apis.LockedResourceTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceConfigSpec.
func (in *NamespaceConfigSpec) DeepCopy() *NamespaceConfigSpec {
	if in == nil {
		return nil
	}
	out := new(NamespaceConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceConfigStatus) DeepCopyInto(out *NamespaceConfigStatus) {
	*out = *in
	in.EnforcingReconcileStatus.DeepCopyInto(&out.EnforcingReconcileStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceConfigStatus.
func (in *NamespaceConfigStatus) DeepCopy() *NamespaceConfigStatus {
	if in == nil {
		return nil
	}
	out := new(NamespaceConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserConfig) DeepCopyInto(out *UserConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserConfig.
func (in *UserConfig) DeepCopy() *UserConfig {
	if in == nil {
		return nil
	}
	out := new(UserConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserConfigList) DeepCopyInto(out *UserConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UserConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserConfigList.
func (in *UserConfigList) DeepCopy() *UserConfigList {
	if in == nil {
		return nil
	}
	out := new(UserConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserConfigSpec) DeepCopyInto(out *UserConfigSpec) {
	*out = *in
	in.LabelSelector.DeepCopyInto(&out.LabelSelector)
	in.AnnotationSelector.DeepCopyInto(&out.AnnotationSelector)
	in.IdentityExtraFieldSelector.DeepCopyInto(&out.IdentityExtraFieldSelector)
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make([]apis.LockedResourceTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserConfigSpec.
func (in *UserConfigSpec) DeepCopy() *UserConfigSpec {
	if in == nil {
		return nil
	}
	out := new(UserConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserConfigStatus) DeepCopyInto(out *UserConfigStatus) {
	*out = *in
	in.EnforcingReconcileStatus.DeepCopyInto(&out.EnforcingReconcileStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserConfigStatus.
func (in *UserConfigStatus) DeepCopy() *UserConfigStatus {
	if in == nil {
		return nil
	}
	out := new(UserConfigStatus)
	in.DeepCopyInto(out)
	return out
}
