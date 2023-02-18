package aws

import (
	"fmt"

	"github.com/darchlabs/infra/pkg/util"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func KeyPair(ctx *pulumi.Context, projectName string) (string, error) {
	// generate ssh key
	privateKey, publicKey, err := util.GenerateKey()
	if err != nil {
		return "", err
	}

	// export pulumi values
	ctx.Export("cloud_private_key", pulumi.String(string(privateKey)))
	ctx.Export("cloud_public_key", pulumi.String(string(publicKey)))

	// create ssh key pair
	keyPairName := fmt.Sprintf("dl-%s-key", projectName)
	_, err = ec2.NewKeyPair(ctx, keyPairName, &ec2.KeyPairArgs{
		KeyName:   pulumi.String(keyPairName),
		PublicKey: pulumi.String(string(publicKey)),
	})
	if err != nil {
		return "", err
	}

	return keyPairName, nil
}
