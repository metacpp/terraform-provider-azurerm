package util

import (
	"log"

	oz "github.com/openzipkin/zipkin-go"
	ozHTTP "github.com/openzipkin/zipkin-go/reporter/http"
	"go.opencensus.io/exporter/zipkin"
	"go.opencensus.io/trace"
)

func InitZipkin(svcName string) {
	localEndpoint, err := oz.NewEndpoint(svcName, "192.168.1.5:5454")
	if err != nil {
		log.Fatalf("Failed to create the local zipkin endpoint: %v", err)
	}
	reporter := ozHTTP.NewReporter("http://localhost:9411/api/v2/spans")
	ze := zipkin.NewExporter(reporter, localEndpoint)
	// Register the Zipkin exporter.
	// This step is needed so that traces can be exported.
	trace.RegisterExporter(ze)

	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
}
