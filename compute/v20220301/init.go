package v20220301

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
	case "azure-native:compute/v20220301:AvailabilitySet":
		r = &AvailabilitySet{}
	case "azure-native:compute/v20220301:CapacityReservation":
		r = &CapacityReservation{}
	case "azure-native:compute/v20220301:CapacityReservationGroup":
		r = &CapacityReservationGroup{}
	case "azure-native:compute/v20220301:DedicatedHost":
		r = &DedicatedHost{}
	case "azure-native:compute/v20220301:DedicatedHostGroup":
		r = &DedicatedHostGroup{}
	case "azure-native:compute/v20220301:Image":
		r = &Image{}
	case "azure-native:compute/v20220301:ProximityPlacementGroup":
		r = &ProximityPlacementGroup{}
	case "azure-native:compute/v20220301:RestorePoint":
		r = &RestorePoint{}
	case "azure-native:compute/v20220301:RestorePointCollection":
		r = &RestorePointCollection{}
	case "azure-native:compute/v20220301:SshPublicKey":
		r = &SshPublicKey{}
	case "azure-native:compute/v20220301:VirtualMachine":
		r = &VirtualMachine{}
	case "azure-native:compute/v20220301:VirtualMachineExtension":
		r = &VirtualMachineExtension{}
	case "azure-native:compute/v20220301:VirtualMachineRunCommandByVirtualMachine":
		r = &VirtualMachineRunCommandByVirtualMachine{}
	case "azure-native:compute/v20220301:VirtualMachineScaleSet":
		r = &VirtualMachineScaleSet{}
	case "azure-native:compute/v20220301:VirtualMachineScaleSetExtension":
		r = &VirtualMachineScaleSetExtension{}
	case "azure-native:compute/v20220301:VirtualMachineScaleSetVM":
		r = &VirtualMachineScaleSetVM{}
	case "azure-native:compute/v20220301:VirtualMachineScaleSetVMExtension":
		r = &VirtualMachineScaleSetVMExtension{}
	case "azure-native:compute/v20220301:VirtualMachineScaleSetVMRunCommand":
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
		"compute/v20220301",
		&module{version},
	)
}
