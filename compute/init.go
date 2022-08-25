package compute

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
	case "azure-native:compute:AvailabilitySet":
		r = &AvailabilitySet{}
	case "azure-native:compute:CapacityReservation":
		r = &CapacityReservation{}
	case "azure-native:compute:CapacityReservationGroup":
		r = &CapacityReservationGroup{}
	case "azure-native:compute:CloudService":
		r = &CloudService{}
	case "azure-native:compute:DedicatedHost":
		r = &DedicatedHost{}
	case "azure-native:compute:DedicatedHostGroup":
		r = &DedicatedHostGroup{}
	case "azure-native:compute:Disk":
		r = &Disk{}
	case "azure-native:compute:DiskAccess":
		r = &DiskAccess{}
	case "azure-native:compute:DiskAccessAPrivateEndpointConnection":
		r = &DiskAccessAPrivateEndpointConnection{}
	case "azure-native:compute:DiskEncryptionSet":
		r = &DiskEncryptionSet{}
	case "azure-native:compute:Gallery":
		r = &Gallery{}
	case "azure-native:compute:GalleryApplication":
		r = &GalleryApplication{}
	case "azure-native:compute:GalleryApplicationVersion":
		r = &GalleryApplicationVersion{}
	case "azure-native:compute:GalleryImage":
		r = &GalleryImage{}
	case "azure-native:compute:GalleryImageVersion":
		r = &GalleryImageVersion{}
	case "azure-native:compute:Image":
		r = &Image{}
	case "azure-native:compute:ProximityPlacementGroup":
		r = &ProximityPlacementGroup{}
	case "azure-native:compute:RestorePoint":
		r = &RestorePoint{}
	case "azure-native:compute:RestorePointCollection":
		r = &RestorePointCollection{}
	case "azure-native:compute:Snapshot":
		r = &Snapshot{}
	case "azure-native:compute:SshPublicKey":
		r = &SshPublicKey{}
	case "azure-native:compute:VirtualMachine":
		r = &VirtualMachine{}
	case "azure-native:compute:VirtualMachineExtension":
		r = &VirtualMachineExtension{}
	case "azure-native:compute:VirtualMachineRunCommandByVirtualMachine":
		r = &VirtualMachineRunCommandByVirtualMachine{}
	case "azure-native:compute:VirtualMachineScaleSet":
		r = &VirtualMachineScaleSet{}
	case "azure-native:compute:VirtualMachineScaleSetExtension":
		r = &VirtualMachineScaleSetExtension{}
	case "azure-native:compute:VirtualMachineScaleSetVM":
		r = &VirtualMachineScaleSetVM{}
	case "azure-native:compute:VirtualMachineScaleSetVMExtension":
		r = &VirtualMachineScaleSetVMExtension{}
	case "azure-native:compute:VirtualMachineScaleSetVMRunCommand":
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
		"compute",
		&module{version},
	)
}
