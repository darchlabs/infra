package aws

import (
	"fmt"

	"github.com/darchlabs/infra/pkg/project"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// create AWS instance
func Instance(ctx *pulumi.Context, projectName string, groupId pulumi.IDOutput, keyPairName string, accessKey string, secretKey string, region project.AWSRegion) (pulumi.IDOutput, error) {
	id := "ami-0aa7d40eeae50c9a9" // Amazon Linux AMI 2 (amd64)

	// create ec2 instance
	instanceName := fmt.Sprintf("dl-%s-instance", projectName)
	instance, err := ec2.NewInstance(ctx, instanceName, &ec2.InstanceArgs{
		Tags:                pulumi.StringMap{"name": pulumi.String(instanceName)},
		InstanceType:        pulumi.String("t3.large"),
		VpcSecurityGroupIds: pulumi.StringArray{groupId},
		Ami:                 pulumi.String(id),
		KeyName:             pulumi.String(keyPairName),
	})
	if err != nil {
		return pulumi.IDOutput{}, err
	}

	// export pulumi values
	ctx.Export("cloud_instance_id", instance.ID())
	ctx.Export("cloud_ssh_user", pulumi.String("ec2-user"))

	return instance.ID(), nil
}
