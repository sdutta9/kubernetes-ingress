// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/nginx/kubernetes-ingress/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// GlobalConfigurations returns a GlobalConfigurationInformer.
	GlobalConfigurations() GlobalConfigurationInformer
	// Policies returns a PolicyInformer.
	Policies() PolicyInformer
	// TransportServers returns a TransportServerInformer.
	TransportServers() TransportServerInformer
	// VirtualServers returns a VirtualServerInformer.
	VirtualServers() VirtualServerInformer
	// VirtualServerRoutes returns a VirtualServerRouteInformer.
	VirtualServerRoutes() VirtualServerRouteInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// GlobalConfigurations returns a GlobalConfigurationInformer.
func (v *version) GlobalConfigurations() GlobalConfigurationInformer {
	return &globalConfigurationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Policies returns a PolicyInformer.
func (v *version) Policies() PolicyInformer {
	return &policyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TransportServers returns a TransportServerInformer.
func (v *version) TransportServers() TransportServerInformer {
	return &transportServerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VirtualServers returns a VirtualServerInformer.
func (v *version) VirtualServers() VirtualServerInformer {
	return &virtualServerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VirtualServerRoutes returns a VirtualServerRouteInformer.
func (v *version) VirtualServerRoutes() VirtualServerRouteInformer {
	return &virtualServerRouteInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
