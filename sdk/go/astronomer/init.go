// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package astronomer

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ryan-pip/pulumi-astronomer/sdk/go/astronomer/internal"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "astronomer:index/apiToken:ApiToken":
		r = &ApiToken{}
	case "astronomer:index/cluster:Cluster":
		r = &Cluster{}
	case "astronomer:index/deployment:Deployment":
		r = &Deployment{}
	case "astronomer:index/hybridClusterWorkspaceAuthorization:HybridClusterWorkspaceAuthorization":
		r = &HybridClusterWorkspaceAuthorization{}
	case "astronomer:index/team:Team":
		r = &Team{}
	case "astronomer:index/teamRoles:TeamRoles":
		r = &TeamRoles{}
	case "astronomer:index/userInvite:UserInvite":
		r = &UserInvite{}
	case "astronomer:index/userRoles:UserRoles":
		r = &UserRoles{}
	case "astronomer:index/workspace:Workspace":
		r = &Workspace{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:astronomer" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"astronomer",
		"index/apiToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"astronomer",
		"index/cluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"astronomer",
		"index/deployment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"astronomer",
		"index/hybridClusterWorkspaceAuthorization",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"astronomer",
		"index/team",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"astronomer",
		"index/teamRoles",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"astronomer",
		"index/userInvite",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"astronomer",
		"index/userRoles",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"astronomer",
		"index/workspace",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"astronomer",
		&pkg{version},
	)
}
