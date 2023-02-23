package k8s

import (
	"github.com/darchlabs/infra/internal/env"
	"github.com/darchlabs/infra/pkg/project"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Run(ctx *pulumi.Context, p *project.Project, env env.Env) error {
	// create namespace
	namespace, err := Namespace(ctx, p.Name, p.Environment)
	if err != nil {
		return err
	}

	// create synchronizers
	err = Synchonizers(ctx, env, namespace)
	if err != nil {
		return err
	}

	// create jobs
	// err = Jobs(ctx, env, namespace)
	// if err != nil {
	// 	return err
	// }

	// TODO(ca): create jobs reporter

	// create nodes
	// err = Nodes(ctx, env, namespace)
	// if err != nil {
	// 	return err
	// }

	// TODO(ca): create synchronizers reporter

	// create jobs reporter

	// TODO(ca): create nodes reporter

	// create web app
	err = webapp(ctx, env, namespace)
	if err != nil {
		return err
	}

	// create ingress
	err = Ingress(ctx, env, namespace, p.URL)
	if err != nil {
		return err
	}

	return nil
}
