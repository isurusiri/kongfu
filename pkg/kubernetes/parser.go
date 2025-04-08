package kubernetes

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
)

// Resource represents a Kubernetes resource with metadata
type Resource struct {
	Kind       string
	Name       string
	Namespace  string
	APIVersion string
	Spec       interface{}
}

// ParseIngressFile reads and parses Kubernetes Ingress resources from a YAML file
func ParseIngressFile(filePath string) ([]Resource, error) {
	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	// Parse YAML documents from the file
	var resources []Resource
	decoder := yaml.NewDecoder(bytes.NewReader(data))

	// Create a new Kubernetes scheme
	s := runtime.NewScheme()
	_ = scheme.AddToScheme(s)

	for {
		var obj map[string]interface{}
		if err := decoder.Decode(&obj); err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("failed to decode YAML: %w", err)
		}

		// Skip if not an Ingress resource
		kind, ok := obj["kind"].(string)
		if !ok || kind != "Ingress" {
			continue
		}

		// Extract basic metadata
		metadata, ok := obj["metadata"].(map[string]interface{})
		if !ok {
			continue
		}

		name, _ := metadata["name"].(string)
		namespace, _ := metadata["namespace"].(string)
		apiVersion, _ := obj["apiVersion"].(string)

		// Create a resource
		resource := Resource{
			Kind:       kind,
			Name:       name,
			Namespace:  namespace,
			APIVersion: apiVersion,
			Spec:       obj["spec"],
		}

		resources = append(resources, resource)
	}

	if len(resources) == 0 {
		return nil, fmt.Errorf("no Ingress resources found in file %s", filePath)
	}

	return resources, nil
}

// For future implementation:
// - Connecting to a Kubernetes cluster using client-go
// - Fetching resources from a live cluster
// - Parsing annotations specific to Kong
