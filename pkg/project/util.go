package project

import (
	"errors"
)

func ValidateProject(p *Project) error {
	// validate params of project input
	if p == nil {
		return errors.New("invalid project")
	}

	// validate name
	if p.Name == "" {
		return errors.New("invalid name")
	}

	// validate password
	if p.Password == "" {
		return errors.New("invalid password")
	}

	// validate environment
	if p.Environment == "" || !ValidEnvironment(p.Environment) {
		return errors.New("invalid environment")
	}

	// validate wallet
	if len(p.Wallet.Address) > 0 && len(p.Wallet.Provider) > 0 && !ValidEnvironment(p.Environment) {
		return errors.New("invalid wallet")
	}

	// validate cloud provider
	if p.Cloud == nil || p.Cloud.Provider == "" || !ValidCloudProvider(p.Cloud.Provider) {
		return errors.New("invalid cloud")
	}

	// validate cloud provider: k8s
	// TODO(ca): should validate valid k8s config
	if p.Cloud.Provider == CloudProviderK8s {
		if p.Cloud.K8s == nil || p.Cloud.K8s.Config == "" {
			return errors.New("invalid k8s config")
		}
	}

	// validate cloud provider: awd
	if p.Cloud.Provider == CloudProviderAWS {
		// TODO(ca): should check if credentials are valid
		if p.Cloud.AWS.AccessKeyId == "" || p.Cloud.AWS.SecretAccessKey == "" {
			return errors.New("invalid aws access_key or secret_access")
		}

		if p.Cloud.AWS.Region == "" || !ValidAWSRegion(p.Cloud.AWS.Region) {
			return errors.New("invalid aws region")
		}
	}

	// valida cloud provider: do
	if p.Cloud.Provider == CloudProviderDO {
		if p.Cloud.DO.Token == "" {
			// TODO(ca): should check if token is valid
			return errors.New("invalid do token")
		}

		if p.Cloud.DO.Region == "" || ValidDORegion(p.Cloud.DO.Region) {
			return errors.New("invalid do region")
		}
	}

	return nil
}

func To(p *ProjectInput) *Project {
	return &Project{
		Name:        p.Name,
		Password:    p.Password,
		Environment: p.Environment,
		Wallet: &Wallet{
			Provider: WalletProvider(p.WalletProvider),
			Address:  p.WalletAddress,
		},
		Cloud: &Cloud{
			Provider: p.Cloud,
			K8s: &K8s{
				Config: p.CredentialsK8sConfig,
			},
			AWS: &AWS{
				AccessKeyId:     p.CredentialsAwsAccessKeyId,
				SecretAccessKey: p.CredentialsAwsSecretAccessKey,
				Region:          AWSRegion(p.CredentialsAwsRegion),
			},
			DO: &DO{
				Token:  p.CredentialsDoToken,
				Region: DORegion(p.CredentialsDoRegion),
			},
		},
	}
}
