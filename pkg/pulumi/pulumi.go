package pulumi

import (
	"context"

	"github.com/darchlabs/infra/internal/env"
	"github.com/darchlabs/infra/pkg/project"
)

func Run(ctx context.Context, p *project.Project, env env.Env) error {
	// first, we need to mount the cloud infrastructure
	if p.Cloud.Provider != project.CloudProviderK8s {
		err := PulumiCloud(ctx, p)
		if err != nil {
			return err
		}
	}

	// second, we need to mount darchs services on cluster
	err := PulumiKubernetes(ctx, p, env)
	if err != nil {
		return err
	}

	return nil
}
