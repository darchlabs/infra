package aws

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ElasticIP(ctx *pulumi.Context, projectName string, instanceId pulumi.IDOutput) error {
	// create elastic ip
	eip, err := ec2.NewEip(ctx, "myEip", nil)
	if err != nil {
		return err
	}

	// assign elastic ip to instance
	eipName := fmt.Sprintf("dl-%s-eip", projectName)
	_, err = ec2.NewEipAssociation(ctx, eipName, &ec2.EipAssociationArgs{
		InstanceId: instanceId,
		PublicIp:   eip.PublicIp,
	})
	if err != nil {
		return err
	}

	// export pulumi values
	ctx.Export("cloud_public_ip", eip.PublicIp)
	ctx.Export("cloud_public_hostname", eip.PublicDns)

	return nil
}
