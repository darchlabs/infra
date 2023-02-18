package pulumi

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/darchlabs/infra/internal/env"
	"github.com/darchlabs/infra/pkg/project"
	pulumik8s "github.com/darchlabs/infra/pkg/pulumi/provider/k8s"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optup"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func PulumiKubernetes(ctx context.Context, p *project.Project, env env.Env) error {
	// create stack if provider is not kubernetes
	stackName := fmt.Sprintf("%s-kubernetes", string(p.Environment))

	// create or select pulumi stack
	stack, err := auto.UpsertStackInlineSource(ctx, stackName, p.Name, func(ctx *pulumi.Context) error {
		err := pulumik8s.Run(ctx, p, env)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	// set stack configuration for K8s
	stack.SetConfig(ctx, "kubernetes:kubeconfig", auto.ConfigValue{Value: p.KubeConfig})

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

	// get and set k8s_namespace from pulumi response
	namespace, ok := response.Outputs["k8s_namespace"].Value.(string)
	if !ok {
		return errors.New("failed to unmarshall output k8s_namespace")
	}
	// p.Namespace = namespace

	fmt.Println("k8s_namespace: ", namespace)

	return nil
}
