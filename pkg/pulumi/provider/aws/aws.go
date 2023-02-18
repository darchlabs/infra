package aws

import (
	"github.com/darchlabs/infra/pkg/project"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Run(ctx *pulumi.Context, projectName string, accessKey string, secretKey string, region project.AWSRegion) error {
	// create key pair on aws
	keyPairName, err := KeyPair(ctx, projectName)
	if err != nil {
		return err
	}

	// create security group on aws
	groupId, err := SecurityGroup(ctx, projectName)
	if err != nil {
		return err
	}

	// create instance on aws
	instanceId, err := Instance(ctx, projectName, groupId, keyPairName, accessKey, secretKey, region)
	if err != nil {
		return err
	}

	// assign elastic ip to instance
	err = ElasticIP(ctx, projectName, instanceId)
	if err != nil {
		return err
	}

	return nil
}
