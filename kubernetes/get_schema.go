package kubernetes

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/kube-openapi/pkg/util/proto"
)

func (r *resources) getSchema() ([]proto.Schema, error) {
	resourcesList, err := r.getServerPreferredResources()
	if err != nil {
		return nil, err
	}
	clientSet, err := getKubeClientSet(r.kubeconfig)
	if err != nil {
		return nil, err
	}

	openAPIResources, err := getOpenAPIResources(clientSet)
	if err != nil {
		return nil, err
	}

	schemas := make([]proto.Schema, 0)
	for _, resource := range resourcesList {
		// resource.APIResources
		gv, err := schema.ParseGroupVersion(resource.GroupVersion)
		if err != nil {
			return nil, err
		}
		for _, apiResource := range resource.APIResources {
			gvk := schema.GroupVersionKind{Group: gv.Group, Version: gv.Version, Kind: apiResource.Kind}
			schema := openAPIResources.LookupResource(gvk)
			// TODO replace with log
			// fmt.Printf("Resource is %+v GVK is %+v schema is %+v\n", apiResource, gvk, schema)
			schemas = append(schemas, schema)
		}
	}
	return schemas, nil
}
