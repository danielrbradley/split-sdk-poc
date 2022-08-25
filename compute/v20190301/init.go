package v20190301

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
	case "azure-native:compute/v20190301:AvailabilitySet":
		r = &AvailabilitySet{}
	case "azure-native:compute/v20190301:DedicatedHost":
		r = &DedicatedHost{}
	case "azure-native:compute/v20190301:DedicatedHostGroup":
		r = &DedicatedHostGroup{}
	case "azure-native:compute/v20190301:Disk":
		r = &Disk{}
	case "azure-native:compute/v20190301:Gallery":
		r = &Gallery{}
	case "azure-native:compute/v20190301:GalleryApplication":
		r = &GalleryApplication{}
	case "azure-native:compute/v20190301:GalleryApplicationVersion":
		r = &GalleryApplicationVersion{}
	case "azure-native:compute/v20190301:GalleryImage":
		r = &GalleryImage{}
	case "azure-native:compute/v20190301:GalleryImageVersion":
		r = &GalleryImageVersion{}
	case "azure-native:compute/v20190301:Image":
		r = &Image{}
	case "azure-native:compute/v20190301:ProximityPlacementGroup":
		r = &ProximityPlacementGroup{}
	case "azure-native:compute/v20190301:Snapshot":
		r = &Snapshot{}
	case "azure-native:compute/v20190301:VirtualMachine":
		r = &VirtualMachine{}
	case "azure-native:compute/v20190301:VirtualMachineExtension":
		r = &VirtualMachineExtension{}
	case "azure-native:compute/v20190301:VirtualMachineScaleSet":
		r = &VirtualMachineScaleSet{}
	case "azure-native:compute/v20190301:VirtualMachineScaleSetExtension":
		r = &VirtualMachineScaleSetExtension{}
	case "azure-native:compute/v20190301:VirtualMachineScaleSetVM":
		r = &VirtualMachineScaleSetVM{}
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
		"compute/v20190301",
		&module{version},
	)
}
