package v20200501

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
	case "azure-native:compute/v20200501:Disk":
		r = &Disk{}
	case "azure-native:compute/v20200501:DiskAccess":
		r = &DiskAccess{}
	case "azure-native:compute/v20200501:DiskEncryptionSet":
		r = &DiskEncryptionSet{}
	case "azure-native:compute/v20200501:Snapshot":
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
		"compute/v20200501",
		&module{version},
	)
}
