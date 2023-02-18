package project

// import (
// 	// "errors"
// 	// "syscall/js"
// )

// func ToProjectWasm(args []js.Value) (*Project, error) {
// 	// define project struct
// 	p := &Project{}

// 	if len(args) == 1 {
// 		// get first value
// 		input := args[0]

// 		if input.Type() == js.TypeObject {
// 			// define name project value
// 			name := input.Get("name")
// 			if !name.IsUndefined() {
// 				p.Name = name.String()
// 			} else {
// 				return nil, errors.New("invalid name")
// 			}

// 			// define password project value
// 			// TODO(ca): maybe it's neccesary apply hash to password
// 			password := input.Get("pasword")
// 			if !name.IsUndefined() {
// 				p.Password = password.String()
// 			} else {
// 				return nil, errors.New("invalid password")
// 			}

// 			// define environment value
// 			env := input.Get("environment")
// 			if !env.IsUndefined() && ValidEnvironment(Environment(env.String())) {
// 				p.Environment = Environment(env.String())
// 			} else {
// 				return nil, errors.New("invalid environment")
// 			}

// 			// define wallet value
// 			wp := input.Get("walletProvider")
// 			wa := input.Get("walletAddress")
// 			if !wp.IsUndefined() && !wa.IsUndefined() {
// 				if ValidWalletProvider(WalletProvider(wp.String())) {
// 					p.Wallet = &Wallet{
// 						Provider: WalletProvider(wp.String()),
// 						Address:  wa.String(),
// 					}
// 				} else {
// 					return nil, errors.New("invalid wallet provider")
// 				}
// 			}

// 			// define cloud provider values
// 			cp := input.Get("cloudProvider")
// 			if !cp.IsUndefined() && ValidCloudProvider(CloudProvider(cp.String())) {
// 				p.Cloud = &Cloud{
// 					Provider: CloudProvider(cp.String()),
// 				}

// 				// get params if cloud provider is k8s
// 				if p.Cloud.Provider == CloudProviderK8s {
// 					// check if k8s config is defined
// 					config := input.Get("credentialsK8sConfig")
// 					if !config.IsUndefined() {
// 						p.Cloud.K8s = &K8s{
// 							Config: config.String(),
// 						}
// 					} else {
// 						return nil, errors.New("invalid k8s config")
// 					}
// 				}

// 				// get params if cloud provider is aws
// 				if p.Cloud.Provider == CloudProviderAWS {
// 					aki := input.Get("credentialsAwsAccessKeyId")
// 					sak := input.Get("credentialsAwsSecretAccessKey")
// 					r := input.Get("credentialsAwsRegion")

// 					// check if aws values are ok
// 					if !aki.IsUndefined() && !sak.IsUndefined() && !r.IsUndefined() {
// 						if ValidAWSRegion(AWSRegion(r.String())) {
// 							p.Cloud.AWS = &AWS{
// 								AccessKeyId:     aki.String(),
// 								SecretAccessKey: sak.String(),
// 								Region:          AWSRegion(r.String()),
// 							}
// 						} else {
// 							return nil, errors.New("invalid aws region")
// 						}
// 					} else {
// 						return nil, errors.New("invalid aws values")
// 					}
// 				}

// 				// get params if cloud provider is do
// 				if p.Cloud.Provider == CloudProviderDO {
// 					t := input.Get("credentialsDoToken")
// 					r := input.Get("credentialsDoRegion")

// 					// check if do values are ok
// 					if !t.IsUndefined() && !r.IsUndefined() {
// 						if ValidDORegion(DORegion(r.String())) {
// 							p.Cloud.DO = &DO{
// 								Token:  t.String(),
// 								Region: DORegion(r.String()),
// 							}
// 						} else {
// 							return nil, errors.New("invalid do region")
// 						}
// 					} else {
// 						return nil, errors.New("invalid do values")
// 					}
// 				}
// 			} else {
// 				return nil, errors.New("invalid cloud provider")
// 			}
// 		} else {
// 			return nil, errors.New("invalid project object")
// 		}
// 	}

// 	return p, nil
// }
