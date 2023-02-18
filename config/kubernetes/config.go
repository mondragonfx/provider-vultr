/*
Copyright 2021 Upbound Inc.
*/

package kubernetes

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("vultr_kubernetes", func(r *ujconfig.Resource) {
	})
}
