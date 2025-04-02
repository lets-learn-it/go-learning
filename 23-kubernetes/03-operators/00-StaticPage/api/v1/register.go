package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "dev.avabodha.in"
const GroupVersion = "v1"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: GroupVersion}

var (
	SchemBuilder = runtime.NewSchemeBuilder(addKnownType)
	AddToScheme  = SchemBuilder.AddToScheme
)

func addKnownType(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&StaticPage{},
		&StaticPageList{},
	)

	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
