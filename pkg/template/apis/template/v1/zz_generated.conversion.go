// +build !ignore_autogenerated_openshift

// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	unsafe "unsafe"

	v1 "github.com/openshift/api/template/v1"
	template "github.com/openshift/origin/pkg/template/apis/template"
	api_core_v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/kubernetes/pkg/apis/core"
	core_v1 "k8s.io/kubernetes/pkg/apis/core/v1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_BrokerTemplateInstance_To_template_BrokerTemplateInstance,
		Convert_template_BrokerTemplateInstance_To_v1_BrokerTemplateInstance,
		Convert_v1_BrokerTemplateInstanceList_To_template_BrokerTemplateInstanceList,
		Convert_template_BrokerTemplateInstanceList_To_v1_BrokerTemplateInstanceList,
		Convert_v1_BrokerTemplateInstanceSpec_To_template_BrokerTemplateInstanceSpec,
		Convert_template_BrokerTemplateInstanceSpec_To_v1_BrokerTemplateInstanceSpec,
		Convert_v1_Parameter_To_template_Parameter,
		Convert_template_Parameter_To_v1_Parameter,
		Convert_v1_Template_To_template_Template,
		Convert_template_Template_To_v1_Template,
		Convert_v1_TemplateInstance_To_template_TemplateInstance,
		Convert_template_TemplateInstance_To_v1_TemplateInstance,
		Convert_v1_TemplateInstanceCondition_To_template_TemplateInstanceCondition,
		Convert_template_TemplateInstanceCondition_To_v1_TemplateInstanceCondition,
		Convert_v1_TemplateInstanceList_To_template_TemplateInstanceList,
		Convert_template_TemplateInstanceList_To_v1_TemplateInstanceList,
		Convert_v1_TemplateInstanceObject_To_template_TemplateInstanceObject,
		Convert_template_TemplateInstanceObject_To_v1_TemplateInstanceObject,
		Convert_v1_TemplateInstanceRequester_To_template_TemplateInstanceRequester,
		Convert_template_TemplateInstanceRequester_To_v1_TemplateInstanceRequester,
		Convert_v1_TemplateInstanceSpec_To_template_TemplateInstanceSpec,
		Convert_template_TemplateInstanceSpec_To_v1_TemplateInstanceSpec,
		Convert_v1_TemplateInstanceStatus_To_template_TemplateInstanceStatus,
		Convert_template_TemplateInstanceStatus_To_v1_TemplateInstanceStatus,
		Convert_v1_TemplateList_To_template_TemplateList,
		Convert_template_TemplateList_To_v1_TemplateList,
	)
}

func autoConvert_v1_BrokerTemplateInstance_To_template_BrokerTemplateInstance(in *v1.BrokerTemplateInstance, out *template.BrokerTemplateInstance, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_BrokerTemplateInstanceSpec_To_template_BrokerTemplateInstanceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_BrokerTemplateInstance_To_template_BrokerTemplateInstance is an autogenerated conversion function.
func Convert_v1_BrokerTemplateInstance_To_template_BrokerTemplateInstance(in *v1.BrokerTemplateInstance, out *template.BrokerTemplateInstance, s conversion.Scope) error {
	return autoConvert_v1_BrokerTemplateInstance_To_template_BrokerTemplateInstance(in, out, s)
}

func autoConvert_template_BrokerTemplateInstance_To_v1_BrokerTemplateInstance(in *template.BrokerTemplateInstance, out *v1.BrokerTemplateInstance, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_template_BrokerTemplateInstanceSpec_To_v1_BrokerTemplateInstanceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_template_BrokerTemplateInstance_To_v1_BrokerTemplateInstance is an autogenerated conversion function.
func Convert_template_BrokerTemplateInstance_To_v1_BrokerTemplateInstance(in *template.BrokerTemplateInstance, out *v1.BrokerTemplateInstance, s conversion.Scope) error {
	return autoConvert_template_BrokerTemplateInstance_To_v1_BrokerTemplateInstance(in, out, s)
}

func autoConvert_v1_BrokerTemplateInstanceList_To_template_BrokerTemplateInstanceList(in *v1.BrokerTemplateInstanceList, out *template.BrokerTemplateInstanceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]template.BrokerTemplateInstance, len(*in))
		for i := range *in {
			if err := Convert_v1_BrokerTemplateInstance_To_template_BrokerTemplateInstance(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1_BrokerTemplateInstanceList_To_template_BrokerTemplateInstanceList is an autogenerated conversion function.
func Convert_v1_BrokerTemplateInstanceList_To_template_BrokerTemplateInstanceList(in *v1.BrokerTemplateInstanceList, out *template.BrokerTemplateInstanceList, s conversion.Scope) error {
	return autoConvert_v1_BrokerTemplateInstanceList_To_template_BrokerTemplateInstanceList(in, out, s)
}

func autoConvert_template_BrokerTemplateInstanceList_To_v1_BrokerTemplateInstanceList(in *template.BrokerTemplateInstanceList, out *v1.BrokerTemplateInstanceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.BrokerTemplateInstance, len(*in))
		for i := range *in {
			if err := Convert_template_BrokerTemplateInstance_To_v1_BrokerTemplateInstance(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_template_BrokerTemplateInstanceList_To_v1_BrokerTemplateInstanceList is an autogenerated conversion function.
func Convert_template_BrokerTemplateInstanceList_To_v1_BrokerTemplateInstanceList(in *template.BrokerTemplateInstanceList, out *v1.BrokerTemplateInstanceList, s conversion.Scope) error {
	return autoConvert_template_BrokerTemplateInstanceList_To_v1_BrokerTemplateInstanceList(in, out, s)
}

func autoConvert_v1_BrokerTemplateInstanceSpec_To_template_BrokerTemplateInstanceSpec(in *v1.BrokerTemplateInstanceSpec, out *template.BrokerTemplateInstanceSpec, s conversion.Scope) error {
	if err := core_v1.Convert_v1_ObjectReference_To_core_ObjectReference(&in.TemplateInstance, &out.TemplateInstance, s); err != nil {
		return err
	}
	if err := core_v1.Convert_v1_ObjectReference_To_core_ObjectReference(&in.Secret, &out.Secret, s); err != nil {
		return err
	}
	out.BindingIDs = *(*[]string)(unsafe.Pointer(&in.BindingIDs))
	return nil
}

// Convert_v1_BrokerTemplateInstanceSpec_To_template_BrokerTemplateInstanceSpec is an autogenerated conversion function.
func Convert_v1_BrokerTemplateInstanceSpec_To_template_BrokerTemplateInstanceSpec(in *v1.BrokerTemplateInstanceSpec, out *template.BrokerTemplateInstanceSpec, s conversion.Scope) error {
	return autoConvert_v1_BrokerTemplateInstanceSpec_To_template_BrokerTemplateInstanceSpec(in, out, s)
}

func autoConvert_template_BrokerTemplateInstanceSpec_To_v1_BrokerTemplateInstanceSpec(in *template.BrokerTemplateInstanceSpec, out *v1.BrokerTemplateInstanceSpec, s conversion.Scope) error {
	if err := core_v1.Convert_core_ObjectReference_To_v1_ObjectReference(&in.TemplateInstance, &out.TemplateInstance, s); err != nil {
		return err
	}
	if err := core_v1.Convert_core_ObjectReference_To_v1_ObjectReference(&in.Secret, &out.Secret, s); err != nil {
		return err
	}
	out.BindingIDs = *(*[]string)(unsafe.Pointer(&in.BindingIDs))
	return nil
}

// Convert_template_BrokerTemplateInstanceSpec_To_v1_BrokerTemplateInstanceSpec is an autogenerated conversion function.
func Convert_template_BrokerTemplateInstanceSpec_To_v1_BrokerTemplateInstanceSpec(in *template.BrokerTemplateInstanceSpec, out *v1.BrokerTemplateInstanceSpec, s conversion.Scope) error {
	return autoConvert_template_BrokerTemplateInstanceSpec_To_v1_BrokerTemplateInstanceSpec(in, out, s)
}

func autoConvert_v1_Parameter_To_template_Parameter(in *v1.Parameter, out *template.Parameter, s conversion.Scope) error {
	out.Name = in.Name
	out.DisplayName = in.DisplayName
	out.Description = in.Description
	out.Value = in.Value
	out.Generate = in.Generate
	out.From = in.From
	out.Required = in.Required
	return nil
}

// Convert_v1_Parameter_To_template_Parameter is an autogenerated conversion function.
func Convert_v1_Parameter_To_template_Parameter(in *v1.Parameter, out *template.Parameter, s conversion.Scope) error {
	return autoConvert_v1_Parameter_To_template_Parameter(in, out, s)
}

func autoConvert_template_Parameter_To_v1_Parameter(in *template.Parameter, out *v1.Parameter, s conversion.Scope) error {
	out.Name = in.Name
	out.DisplayName = in.DisplayName
	out.Description = in.Description
	out.Value = in.Value
	out.Generate = in.Generate
	out.From = in.From
	out.Required = in.Required
	return nil
}

// Convert_template_Parameter_To_v1_Parameter is an autogenerated conversion function.
func Convert_template_Parameter_To_v1_Parameter(in *template.Parameter, out *v1.Parameter, s conversion.Scope) error {
	return autoConvert_template_Parameter_To_v1_Parameter(in, out, s)
}

func autoConvert_v1_Template_To_template_Template(in *v1.Template, out *template.Template, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Message = in.Message
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = make([]runtime.Object, len(*in))
		for i := range *in {
			if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Objects = nil
	}
	out.Parameters = *(*[]template.Parameter)(unsafe.Pointer(&in.Parameters))
	out.ObjectLabels = *(*map[string]string)(unsafe.Pointer(&in.ObjectLabels))
	return nil
}

// Convert_v1_Template_To_template_Template is an autogenerated conversion function.
func Convert_v1_Template_To_template_Template(in *v1.Template, out *template.Template, s conversion.Scope) error {
	return autoConvert_v1_Template_To_template_Template(in, out, s)
}

func autoConvert_template_Template_To_v1_Template(in *template.Template, out *v1.Template, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Message = in.Message
	out.Parameters = *(*[]v1.Parameter)(unsafe.Pointer(&in.Parameters))
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = make([]runtime.RawExtension, len(*in))
		for i := range *in {
			if err := runtime.Convert_runtime_Object_To_runtime_RawExtension(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Objects = nil
	}
	out.ObjectLabels = *(*map[string]string)(unsafe.Pointer(&in.ObjectLabels))
	return nil
}

// Convert_template_Template_To_v1_Template is an autogenerated conversion function.
func Convert_template_Template_To_v1_Template(in *template.Template, out *v1.Template, s conversion.Scope) error {
	return autoConvert_template_Template_To_v1_Template(in, out, s)
}

func autoConvert_v1_TemplateInstance_To_template_TemplateInstance(in *v1.TemplateInstance, out *template.TemplateInstance, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_TemplateInstanceSpec_To_template_TemplateInstanceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_TemplateInstanceStatus_To_template_TemplateInstanceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_TemplateInstance_To_template_TemplateInstance is an autogenerated conversion function.
func Convert_v1_TemplateInstance_To_template_TemplateInstance(in *v1.TemplateInstance, out *template.TemplateInstance, s conversion.Scope) error {
	return autoConvert_v1_TemplateInstance_To_template_TemplateInstance(in, out, s)
}

func autoConvert_template_TemplateInstance_To_v1_TemplateInstance(in *template.TemplateInstance, out *v1.TemplateInstance, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_template_TemplateInstanceSpec_To_v1_TemplateInstanceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_template_TemplateInstanceStatus_To_v1_TemplateInstanceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_template_TemplateInstance_To_v1_TemplateInstance is an autogenerated conversion function.
func Convert_template_TemplateInstance_To_v1_TemplateInstance(in *template.TemplateInstance, out *v1.TemplateInstance, s conversion.Scope) error {
	return autoConvert_template_TemplateInstance_To_v1_TemplateInstance(in, out, s)
}

func autoConvert_v1_TemplateInstanceCondition_To_template_TemplateInstanceCondition(in *v1.TemplateInstanceCondition, out *template.TemplateInstanceCondition, s conversion.Scope) error {
	out.Type = template.TemplateInstanceConditionType(in.Type)
	out.Status = core.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v1_TemplateInstanceCondition_To_template_TemplateInstanceCondition is an autogenerated conversion function.
func Convert_v1_TemplateInstanceCondition_To_template_TemplateInstanceCondition(in *v1.TemplateInstanceCondition, out *template.TemplateInstanceCondition, s conversion.Scope) error {
	return autoConvert_v1_TemplateInstanceCondition_To_template_TemplateInstanceCondition(in, out, s)
}

func autoConvert_template_TemplateInstanceCondition_To_v1_TemplateInstanceCondition(in *template.TemplateInstanceCondition, out *v1.TemplateInstanceCondition, s conversion.Scope) error {
	out.Type = v1.TemplateInstanceConditionType(in.Type)
	out.Status = api_core_v1.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_template_TemplateInstanceCondition_To_v1_TemplateInstanceCondition is an autogenerated conversion function.
func Convert_template_TemplateInstanceCondition_To_v1_TemplateInstanceCondition(in *template.TemplateInstanceCondition, out *v1.TemplateInstanceCondition, s conversion.Scope) error {
	return autoConvert_template_TemplateInstanceCondition_To_v1_TemplateInstanceCondition(in, out, s)
}

func autoConvert_v1_TemplateInstanceList_To_template_TemplateInstanceList(in *v1.TemplateInstanceList, out *template.TemplateInstanceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]template.TemplateInstance, len(*in))
		for i := range *in {
			if err := Convert_v1_TemplateInstance_To_template_TemplateInstance(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1_TemplateInstanceList_To_template_TemplateInstanceList is an autogenerated conversion function.
func Convert_v1_TemplateInstanceList_To_template_TemplateInstanceList(in *v1.TemplateInstanceList, out *template.TemplateInstanceList, s conversion.Scope) error {
	return autoConvert_v1_TemplateInstanceList_To_template_TemplateInstanceList(in, out, s)
}

func autoConvert_template_TemplateInstanceList_To_v1_TemplateInstanceList(in *template.TemplateInstanceList, out *v1.TemplateInstanceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.TemplateInstance, len(*in))
		for i := range *in {
			if err := Convert_template_TemplateInstance_To_v1_TemplateInstance(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_template_TemplateInstanceList_To_v1_TemplateInstanceList is an autogenerated conversion function.
func Convert_template_TemplateInstanceList_To_v1_TemplateInstanceList(in *template.TemplateInstanceList, out *v1.TemplateInstanceList, s conversion.Scope) error {
	return autoConvert_template_TemplateInstanceList_To_v1_TemplateInstanceList(in, out, s)
}

func autoConvert_v1_TemplateInstanceObject_To_template_TemplateInstanceObject(in *v1.TemplateInstanceObject, out *template.TemplateInstanceObject, s conversion.Scope) error {
	if err := core_v1.Convert_v1_ObjectReference_To_core_ObjectReference(&in.Ref, &out.Ref, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_TemplateInstanceObject_To_template_TemplateInstanceObject is an autogenerated conversion function.
func Convert_v1_TemplateInstanceObject_To_template_TemplateInstanceObject(in *v1.TemplateInstanceObject, out *template.TemplateInstanceObject, s conversion.Scope) error {
	return autoConvert_v1_TemplateInstanceObject_To_template_TemplateInstanceObject(in, out, s)
}

func autoConvert_template_TemplateInstanceObject_To_v1_TemplateInstanceObject(in *template.TemplateInstanceObject, out *v1.TemplateInstanceObject, s conversion.Scope) error {
	if err := core_v1.Convert_core_ObjectReference_To_v1_ObjectReference(&in.Ref, &out.Ref, s); err != nil {
		return err
	}
	return nil
}

// Convert_template_TemplateInstanceObject_To_v1_TemplateInstanceObject is an autogenerated conversion function.
func Convert_template_TemplateInstanceObject_To_v1_TemplateInstanceObject(in *template.TemplateInstanceObject, out *v1.TemplateInstanceObject, s conversion.Scope) error {
	return autoConvert_template_TemplateInstanceObject_To_v1_TemplateInstanceObject(in, out, s)
}

func autoConvert_v1_TemplateInstanceRequester_To_template_TemplateInstanceRequester(in *v1.TemplateInstanceRequester, out *template.TemplateInstanceRequester, s conversion.Scope) error {
	out.Username = in.Username
	out.UID = in.UID
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]template.ExtraValue)(unsafe.Pointer(&in.Extra))
	return nil
}

// Convert_v1_TemplateInstanceRequester_To_template_TemplateInstanceRequester is an autogenerated conversion function.
func Convert_v1_TemplateInstanceRequester_To_template_TemplateInstanceRequester(in *v1.TemplateInstanceRequester, out *template.TemplateInstanceRequester, s conversion.Scope) error {
	return autoConvert_v1_TemplateInstanceRequester_To_template_TemplateInstanceRequester(in, out, s)
}

func autoConvert_template_TemplateInstanceRequester_To_v1_TemplateInstanceRequester(in *template.TemplateInstanceRequester, out *v1.TemplateInstanceRequester, s conversion.Scope) error {
	out.Username = in.Username
	out.UID = in.UID
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]v1.ExtraValue)(unsafe.Pointer(&in.Extra))
	return nil
}

// Convert_template_TemplateInstanceRequester_To_v1_TemplateInstanceRequester is an autogenerated conversion function.
func Convert_template_TemplateInstanceRequester_To_v1_TemplateInstanceRequester(in *template.TemplateInstanceRequester, out *v1.TemplateInstanceRequester, s conversion.Scope) error {
	return autoConvert_template_TemplateInstanceRequester_To_v1_TemplateInstanceRequester(in, out, s)
}

func autoConvert_v1_TemplateInstanceSpec_To_template_TemplateInstanceSpec(in *v1.TemplateInstanceSpec, out *template.TemplateInstanceSpec, s conversion.Scope) error {
	if err := Convert_v1_Template_To_template_Template(&in.Template, &out.Template, s); err != nil {
		return err
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(core.LocalObjectReference)
		if err := core_v1.Convert_v1_LocalObjectReference_To_core_LocalObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Secret = nil
	}
	out.Requester = (*template.TemplateInstanceRequester)(unsafe.Pointer(in.Requester))
	return nil
}

// Convert_v1_TemplateInstanceSpec_To_template_TemplateInstanceSpec is an autogenerated conversion function.
func Convert_v1_TemplateInstanceSpec_To_template_TemplateInstanceSpec(in *v1.TemplateInstanceSpec, out *template.TemplateInstanceSpec, s conversion.Scope) error {
	return autoConvert_v1_TemplateInstanceSpec_To_template_TemplateInstanceSpec(in, out, s)
}

func autoConvert_template_TemplateInstanceSpec_To_v1_TemplateInstanceSpec(in *template.TemplateInstanceSpec, out *v1.TemplateInstanceSpec, s conversion.Scope) error {
	if err := Convert_template_Template_To_v1_Template(&in.Template, &out.Template, s); err != nil {
		return err
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(api_core_v1.LocalObjectReference)
		if err := core_v1.Convert_core_LocalObjectReference_To_v1_LocalObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Secret = nil
	}
	out.Requester = (*v1.TemplateInstanceRequester)(unsafe.Pointer(in.Requester))
	return nil
}

// Convert_template_TemplateInstanceSpec_To_v1_TemplateInstanceSpec is an autogenerated conversion function.
func Convert_template_TemplateInstanceSpec_To_v1_TemplateInstanceSpec(in *template.TemplateInstanceSpec, out *v1.TemplateInstanceSpec, s conversion.Scope) error {
	return autoConvert_template_TemplateInstanceSpec_To_v1_TemplateInstanceSpec(in, out, s)
}

func autoConvert_v1_TemplateInstanceStatus_To_template_TemplateInstanceStatus(in *v1.TemplateInstanceStatus, out *template.TemplateInstanceStatus, s conversion.Scope) error {
	out.Conditions = *(*[]template.TemplateInstanceCondition)(unsafe.Pointer(&in.Conditions))
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = make([]template.TemplateInstanceObject, len(*in))
		for i := range *in {
			if err := Convert_v1_TemplateInstanceObject_To_template_TemplateInstanceObject(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Objects = nil
	}
	return nil
}

// Convert_v1_TemplateInstanceStatus_To_template_TemplateInstanceStatus is an autogenerated conversion function.
func Convert_v1_TemplateInstanceStatus_To_template_TemplateInstanceStatus(in *v1.TemplateInstanceStatus, out *template.TemplateInstanceStatus, s conversion.Scope) error {
	return autoConvert_v1_TemplateInstanceStatus_To_template_TemplateInstanceStatus(in, out, s)
}

func autoConvert_template_TemplateInstanceStatus_To_v1_TemplateInstanceStatus(in *template.TemplateInstanceStatus, out *v1.TemplateInstanceStatus, s conversion.Scope) error {
	out.Conditions = *(*[]v1.TemplateInstanceCondition)(unsafe.Pointer(&in.Conditions))
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = make([]v1.TemplateInstanceObject, len(*in))
		for i := range *in {
			if err := Convert_template_TemplateInstanceObject_To_v1_TemplateInstanceObject(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Objects = nil
	}
	return nil
}

// Convert_template_TemplateInstanceStatus_To_v1_TemplateInstanceStatus is an autogenerated conversion function.
func Convert_template_TemplateInstanceStatus_To_v1_TemplateInstanceStatus(in *template.TemplateInstanceStatus, out *v1.TemplateInstanceStatus, s conversion.Scope) error {
	return autoConvert_template_TemplateInstanceStatus_To_v1_TemplateInstanceStatus(in, out, s)
}

func autoConvert_v1_TemplateList_To_template_TemplateList(in *v1.TemplateList, out *template.TemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]template.Template, len(*in))
		for i := range *in {
			if err := Convert_v1_Template_To_template_Template(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1_TemplateList_To_template_TemplateList is an autogenerated conversion function.
func Convert_v1_TemplateList_To_template_TemplateList(in *v1.TemplateList, out *template.TemplateList, s conversion.Scope) error {
	return autoConvert_v1_TemplateList_To_template_TemplateList(in, out, s)
}

func autoConvert_template_TemplateList_To_v1_TemplateList(in *template.TemplateList, out *v1.TemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.Template, len(*in))
		for i := range *in {
			if err := Convert_template_Template_To_v1_Template(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_template_TemplateList_To_v1_TemplateList is an autogenerated conversion function.
func Convert_template_TemplateList_To_v1_TemplateList(in *template.TemplateList, out *v1.TemplateList, s conversion.Scope) error {
	return autoConvert_template_TemplateList_To_v1_TemplateList(in, out, s)
}
