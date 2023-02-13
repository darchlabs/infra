package pulumi

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/darchlabs/infra/pkg/project"
	"github.com/darchlabs/infra/pkg/pulumi/provider/ssh"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optup"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	pulumiaws "github.com/darchlabs/infra/pkg/pulumi/provider/aws"
)

func PulumiCloud(ctx context.Context, p *project.Project) error {
	// create stack if provider is not kubernetes
	stackName := fmt.Sprintf("%s-cloud", string(p.Environment))

	// create or select pulumi stack
	stack, err := auto.UpsertStackInlineSource(ctx, stackName, p.Name, func(ctx *pulumi.Context) error {
		var err error

		// run plan depending cloud
		switch p.Cloud.Provider {
		case project.CloudProviderAWS:
			err = pulumiaws.Run(ctx, p.Name, p.Cloud.AWS.AccessKeyId, p.Cloud.AWS.SecretAccessKey, p.Cloud.AWS.Region)
			break
			// TODO(ca): Digital Ocean
			// TODO(ca): Azure
		}

		// check pulumi cloud error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	// set stack configuration if cloud is AWS
	if p.Cloud.Provider == project.CloudProviderAWS {
		stack.SetConfig(ctx, "aws:accessKey", auto.ConfigValue{Value: string(p.Cloud.AWS.AccessKeyId)})
		stack.SetConfig(ctx, "aws:secretKey", auto.ConfigValue{Value: string(p.Cloud.AWS.SecretAccessKey)})
		stack.SetConfig(ctx, "aws:region", auto.ConfigValue{Value: string(p.Cloud.AWS.Region)})
	}

	// pulumi config set digitalocean:token XXXXXXXXXXXXXX --secret
	if p.Cloud.Provider == project.CloudProviderDO {
		stack.SetConfig(ctx, "digitalocean:token", auto.ConfigValue{Value: string(p.Cloud.DO.Token)})
		stack.SetConfig(ctx, "digitalocean:region", auto.ConfigValue{Value: string(p.Cloud.DO.Region)})
	}

	// refresh stack
	_, err = stack.Refresh(ctx)
	if err != nil {
		return err
	}

	// update stack
	stdout := optup.ProgressStreams(os.Stdout)
	response, err := stack.Up(ctx, stdout)
	if err != nil {
		return err
	}

	// get and set cloud_private_key from pulumi response
	privateKey, ok := response.Outputs["cloud_private_key"].Value.(string)
	if !ok {
		return errors.New("failed to unmarshall output cloud_private_key")
	}
	p.PrivateKey = privateKey

	// get and set cloud_public_key from pulumi response
	publicKey, ok := response.Outputs["cloud_public_key"].Value.(string)
	if !ok {
		return errors.New("failed to unmarshall output cloud_public_key")
	}
	p.PublicKey = publicKey

	// get and set cloud_public_hostname from pulumi response
	url, ok := response.Outputs["cloud_public_hostname"].Value.(string)
	if !ok {
		return errors.New("failed to unmarshall output cloud_public_hostname")
	}
	p.URL = url

	// get and set cloud_public_ip from pulumi response
	ip, ok := response.Outputs["cloud_public_ip"].Value.(string)
	if !ok {
		return errors.New("failed to unmarshall output cloud_public_ip")
	}
	p.IP = ip

	// get and set cloud_instance_id from pulumi response
	instanceId, ok := response.Outputs["cloud_instance_id"].Value.(string)
	if !ok {
		return errors.New("failed to unmarshall output cloud_instance_id")
	}
	p.InstanceId = instanceId

	// get and set cloud_ssh_user from pulumi response
	sshUser, ok := response.Outputs["cloud_ssh_user"].Value.(string)
	if !ok {
		return errors.New("failed to unmarshall output cloud_ssh_user")
	}
	p.SshUser = sshUser

	// check if instance is ready
	svc := ec2.New(ec2.Options{
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(p.Cloud.AWS.AccessKeyId, p.Cloud.AWS.SecretAccessKey, "")),
		Region:      string(p.Cloud.AWS.Region),
	})

	// get instance from aws sdk
	result, err := svc.DescribeInstances(context.Background(), &ec2.DescribeInstancesInput{
		InstanceIds: []string{instanceId},
	})
	if err != nil {
		return err
	}

	// wait for instance status "running"
	ready := false
	for !ready {
		for _, reservation := range result.Reservations {
			for _, instance := range reservation.Instances {
				if *instance.InstanceId == instanceId && instance.State.Name == "running" {
					ready = true
				}
			}
		}

		// wait 5 seconds
		if !ready {
			log.Println("waiting for 5 seconds")
			time.Sleep(5 * time.Second)
		}
	}

	// install cluster and dependencies on instance
	kubeConfig, err := ssh.InstallDependencies(ip, sshUser, privateKey)
	if err != nil {
		return err
	}

	// replace localhost for public ip
	p.KubeConfig = strings.Replace(kubeConfig, "127.0.0.1", ip, 1)

	return nil
}
