package generator

import (
	kapi "github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/kubectl"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/runtime"

	"github.com/openshift/origin/pkg/route/api"
)

// RouteGenerator implements the kubectl.Generator interface for routes
type RouteGenerator struct{}

// ParamNames returns the parameters required for generating a route
func (RouteGenerator) ParamNames() []kubectl.GeneratorParam {
	return []kubectl.GeneratorParam{
		{"labels", false},
		{"name", true},
	}
}

// Generate accepts a set of parameters and maps them into a new route
func (RouteGenerator) Generate(params map[string]string) (runtime.Object, error) {
	labels, err := kubectl.ParseLabels(params["labels"])
	if err != nil {
		return nil, err
	}

	return &api.Route{
		ObjectMeta: kapi.ObjectMeta{
			Name:   params["name"],
			Labels: labels,
		},
		ServiceName: params["name"],
	}, nil
}

// Useful pattern for validating that RouteGenerator implements
// the Generator interface
var _ kubectl.Generator = RouteGenerator{}