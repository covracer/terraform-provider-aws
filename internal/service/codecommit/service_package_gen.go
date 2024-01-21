// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package codecommit

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	codecommit_sdkv2 "github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceApprovalRuleTemplate,
			TypeName: "aws_codecommit_approval_rule_template",
			Name:     "Approval Rule Template",
		},
		{
			Factory:  DataSourceRepository,
			TypeName: "aws_codecommit_repository",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceApprovalRuleTemplate,
			TypeName: "aws_codecommit_approval_rule_template",
			Name:     "Approval Rule Template",
		},
		{
			Factory:  resourceApprovalRuleTemplateAssociation,
			TypeName: "aws_codecommit_approval_rule_template_association",
			Name:     "Approval Rule Template Association",
		},
		{
			Factory:  ResourceRepository,
			TypeName: "aws_codecommit_repository",
			Name:     "Repository",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceTrigger,
			TypeName: "aws_codecommit_trigger",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.CodeCommit
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*codecommit_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return codecommit_sdkv2.NewFromConfig(cfg, func(o *codecommit_sdkv2.Options) {
		if endpoint := config["endpoint"].(string); endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
