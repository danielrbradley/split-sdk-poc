package v20210801

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/danielrbradley/split-sdk-poc"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure-native:compute/v20210801:Disk":
		r = &Disk{}
	case "azure-native:compute/v20210801:DiskAccess":
		r = &DiskAccess{}
	case "azure-native:compute/v20210801:DiskAccessAPrivateEndpointConnection":
		r = &DiskAccessAPrivateEndpointConnection{}
	case "azure-native:compute/v20210801:DiskEncryptionSet":
		r = &DiskEncryptionSet{}
	case "azure-native:compute/v20210801:Snapshot":
		r = &Snapshot{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"compute/v20210801",
		&module{version},
	)
}
