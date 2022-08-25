package v20170330

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
	case "azure-native:compute/v20170330:AvailabilitySet":
		r = &AvailabilitySet{}
	case "azure-native:compute/v20170330:Disk":
		r = &Disk{}
	case "azure-native:compute/v20170330:Image":
		r = &Image{}
	case "azure-native:compute/v20170330:Snapshot":
		r = &Snapshot{}
	case "azure-native:compute/v20170330:VirtualMachine":
		r = &VirtualMachine{}
	case "azure-native:compute/v20170330:VirtualMachineExtension":
		r = &VirtualMachineExtension{}
	case "azure-native:compute/v20170330:VirtualMachineScaleSet":
		r = &VirtualMachineScaleSet{}
	case "azure-native:compute/v20170330:VirtualMachineScaleSetExtension":
		r = &VirtualMachineScaleSetExtension{}
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
		"compute/v20170330",
		&module{version},
	)
}
