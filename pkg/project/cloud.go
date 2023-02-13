package project

type CloudProvider string

const (
	CloudProviderK8s   CloudProvider = "k8s"
	CloudProviderAWS   CloudProvider = "aws"
	CloudProviderAzure CloudProvider = "azure"
	CloudProviderDO    CloudProvider = "do"
	CloudProviderGCP   CloudProvider = "gcp"
)

type Cloud struct {
	Provider CloudProvider `json:"provider"`
	K8s      *K8s          `json:"k8s,omitempty"`
	AWS      *AWS          `json:"aws,omitempty"`
	DO       *DO           `json:"do,omitempty"`
}

// ValidCloudProvider method
func ValidCloudProvider(p CloudProvider) bool {
	cloudProviders := []CloudProvider{
		CloudProviderK8s,
		CloudProviderAWS,
		CloudProviderAzure,
		CloudProviderDO,
		CloudProviderGCP,
	}

	for _, cloudProvider := range cloudProviders {
		if cloudProvider == p {
			return true
		}
	}

	return false
}
