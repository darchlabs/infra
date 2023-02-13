package project

import "github.com/darchlabs/infra/internal/env"

type Project struct {
	Name        string      `json:"name"`
	Password    string      `json:"password"`
	Environment Environment `json:"environment"`
	Wallet      *Wallet     `json:"wallet"`
	Cloud       *Cloud      `json:"cloudProvider"`

	Error        string `json:"error"`
	Provisioning bool   `json:"provisioning"`
	Ready        bool   `json:"ready"`

	URL        string `json:"url"`
	IP         string `json:"ip"`
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
	SshUser    string `json:"sshUser"`
	InstanceId string `json:"instanceId"`
	KubeConfig string `json:"kubeConfig"`
}

type ProjectInput struct {
	Name                          string        `json:"name"`
	Password                      string        `json:"password"`
	Environment                   Environment   `json:"environment"`
	WalletProvider                string        `json:"walletProvider"`
	WalletAddress                 string        `json:"walletAddress"`
	Cloud                         CloudProvider `json:"cloudProvider"`
	CredentialsK8sConfig          string        `json:"credentialsK8sConfig"`
	CredentialsAwsAccessKeyId     string        `json:"credentialsAwsAccessKeyId"`
	CredentialsAwsSecretAccessKey string        `json:"credentialsAwsSecretAccessKey"`
	CredentialsAwsRegion          string        `json:"credentialsAwsRegion"`
	CredentialsDoToken            string        `json:"credentialsDoToken"`
	CredentialsDoRegion           string        `json:"credentialsDoRegion"`
}

// Service ..
type Service interface {
	Create(p *Project, env env.Env) (*Project, error)
}
