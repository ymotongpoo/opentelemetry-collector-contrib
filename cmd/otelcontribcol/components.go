// Copyright 2019 OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenterror"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/service/defaultcomponents"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/alibabacloudlogserviceexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuremonitorexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/carbonexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/honeycombexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/jaegerthrifthttpexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kinesisexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/lightstepexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/newrelicexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sapmexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/stackdriverexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/k8sobserver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor"
)

func components() (config.Factories, error) {
	errs := []error{}
	factories, err := defaultcomponents.Components()
	if err != nil {
		return config.Factories{}, err
	}

	extensions := []component.ExtensionFactory{
		k8sobserver.NewFactory(),
	}

	for _, ext := range factories.Extensions {
		extensions = append(extensions, ext)
	}

	factories.Extensions, err = component.MakeExtensionFactoryMap(extensions...)
	if err != nil {
		errs = append(errs, err)
	}

	receivers := allReceivers()
	for _, rcv := range factories.Receivers {
		receivers = append(receivers, rcv)
	}
	factories.Receivers, err = component.MakeReceiverFactoryMap(receivers...)
	if err != nil {
		errs = append(errs, err)
	}

	exporters := []component.ExporterFactoryBase{
		&stackdriverexporter.Factory{},
		&azuremonitorexporter.Factory{},
		&signalfxexporter.Factory{},
		&sapmexporter.Factory{},
		&kinesisexporter.Factory{},
		&awsxrayexporter.Factory{},
		&carbonexporter.Factory{},
		&honeycombexporter.Factory{},
		&jaegerthrifthttpexporter.Factory{},
		&lightstepexporter.Factory{},
		&newrelicexporter.Factory{},
		&splunkhecexporter.Factory{},
		&elasticexporter.Factory{},
		&alibabacloudlogserviceexporter.Factory{},
	}
	for _, exp := range factories.Exporters {
		exporters = append(exporters, exp)
	}
	factories.Exporters, err = component.MakeExporterFactoryMap(exporters...)
	if err != nil {
		errs = append(errs, err)
	}

	processors := []component.ProcessorFactoryBase{
		&k8sprocessor.Factory{},
		resourcedetectionprocessor.NewFactory(),
		&metricstransformprocessor.Factory{},
	}
	for _, pr := range factories.Processors {
		processors = append(processors, pr)
	}
	factories.Processors, err = component.MakeProcessorFactoryMap(processors...)
	if err != nil {
		errs = append(errs, err)
	}

	return factories, componenterror.CombineErrors(errs)
}
