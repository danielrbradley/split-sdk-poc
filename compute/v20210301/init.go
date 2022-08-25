package v20210301

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
	case "azure-native:compute/v20210301:AvailabilitySet":
		r = &AvailabilitySet{}
	case "azure-native:compute/v20210301:CloudService":
		r = &CloudService{}
	case "azure-native:compute/v20210301:DedicatedHost":
		r = &DedicatedHost{}
	case "azure-native:compute/v20210301:DedicatedHostGroup":
		r = &DedicatedHostGroup{}
	case "azure-native:compute/v20210301:Image":
		r = &Image{}
	case "azure-native:compute/v20210301:ProximityPlacementGroup":
		r = &ProximityPlacementGroup{}
	case "azure-native:compute/v20210301:RestorePoint":
		r = &RestorePoint{}
	case "azure-native:compute/v20210301:RestorePointCollection":
		r = &RestorePointCollection{}
	case "azure-native:compute/v20210301:SshPublicKey":
		r = &SshPublicKey{}
	case "azure-native:compute/v20210301:VirtualMachine":
		r = &VirtualMachine{}
	case "azure-native:compute/v20210301:VirtualMachineExtension":
		r = &VirtualMachineExtension{}
	case "azure-native:compute/v20210301:VirtualMachineRunCommandByVirtualMachine":
		r = &VirtualMachineRunCommandByVirtualMachine{}
	case "azure-native:compute/v20210301:VirtualMachineScaleSet":
		r = &VirtualMachineScaleSet{}
	case "azure-native:compute/v20210301:VirtualMachineScaleSetExtension":
		r = &VirtualMachineScaleSetExtension{}
	case "azure-native:compute/v20210301:VirtualMachineScaleSetVM":
		r = &VirtualMachineScaleSetVM{}
	case "azure-native:compute/v20210301:VirtualMachineScaleSetVMExtension":
		r = &VirtualMachineScaleSetVMExtension{}
	case "azure-native:compute/v20210301:VirtualMachineScaleSetVMRunCommand":
		r = &VirtualMachineScaleSetVMRunCommand{}
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
		"compute/v20210301",
		&module{version},
	)
}
