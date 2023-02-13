package projectservice

import (
	"context"
	"errors"
	"strings"

	"github.com/darchlabs/infra/internal/env"
	"github.com/darchlabs/infra/pkg/project"
	"github.com/darchlabs/infra/pkg/pulumi"
)

// ProjectService ...
type ProjectService struct{}

// NewService ...
func NewService() *ProjectService {
	return &ProjectService{}
}

// Create ...
func (ps *ProjectService) Create(p *project.Project, env env.Env) (*project.Project, error) {
	// check if project input are valid
	err := project.ValidateProject(p)
	if err != nil {
		return nil, err
	}

	// define context for use in stack
	ctx := context.Background()

	// run pulumi stack
	p.Provisioning = true
	p.Ready = true

	err = pulumi.Run(ctx, p, env)
	if err != nil {
		p.Provisioning = false
		p.Ready = false

		if strings.Contains(err.Error(), "unable to validate AWS credentials") {
			return nil, errors.New("invalid credentials")
		}

		if strings.Contains(err.Error(), "failed to select stack") {
			return nil, errors.New("invalid pulumi credential")
		}

		if strings.Contains(err.Error(), "UnauthorizedOperation") {
			return nil, errors.New("unauthorized operation")
		}

		return nil, err
	}

	// finish proccess
	p.Provisioning = false
	p.Ready = true

	return p, nil
}
