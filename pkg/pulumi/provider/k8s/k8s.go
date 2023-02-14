package k8s

import (
	"github.com/darchlabs/infra/internal/env"
	"github.com/darchlabs/infra/pkg/project"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Run(ctx *pulumi.Context, p *project.Project, env env.Env) error {
	// 02.- create namespace
	namespace, err := Namespace(ctx, p.Name, p.Environment)
	if err != nil {
		return err
	}

	// 05.- create synchronizers
	err = Synchonizers(ctx, env, namespace)
	if err != nil {
		return err
	}

	return nil
}
