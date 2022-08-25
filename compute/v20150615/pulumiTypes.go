


package v20150615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AdditionalUnattendContent struct {
	ComponentName *ComponentNames `pulumi:"componentName"`
	Content       *string         `pulumi:"content"`
	PassName      *PassNames      `pulumi:"passName"`
	SettingName   *SettingNames   `pulumi:"settingName"`
}





type AdditionalUnattendContentInput interface {
	pulumi.Input

	ToAdditionalUnattendContentOutput() AdditionalUnattendContentOutput
	ToAdditionalUnattendContentOutputWithContext(context.Context) AdditionalUnattendContentOutput
}

type AdditionalUnattendContentArgs struct {
	ComponentName ComponentNamesPtrInput `pulumi:"componentName"`
	Content       pulumi.StringPtrInput  `pulumi:"content"`
	PassName      PassNamesPtrInput      `pulumi:"passName"`
	SettingName   SettingNamesPtrInput   `pulumi:"settingName"`
}

func (AdditionalUnattendContentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalUnattendContent)(nil)).Elem()
}

func (i AdditionalUnattendContentArgs) ToAdditionalUnattendContentOutput() AdditionalUnattendContentOutput {
	return i.ToAdditionalUnattendContentOutputWithContext(context.Background())
}

func (i AdditionalUnattendContentArgs) ToAdditionalUnattendContentOutputWithContext(ctx context.Context) AdditionalUnattendContentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalUnattendContentOutput)
}





type AdditionalUnattendContentArrayInput interface {
	pulumi.Input

	ToAdditionalUnattendContentArrayOutput() AdditionalUnattendContentArrayOutput
	ToAdditionalUnattendContentArrayOutputWithContext(context.Context) AdditionalUnattendContentArrayOutput
}

type AdditionalUnattendContentArray []AdditionalUnattendContentInput

func (AdditionalUnattendContentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalUnattendContent)(nil)).Elem()
}

func (i AdditionalUnattendContentArray) ToAdditionalUnattendContentArrayOutput() AdditionalUnattendContentArrayOutput {
	return i.ToAdditionalUnattendContentArrayOutputWithContext(context.Background())
}

func (i AdditionalUnattendContentArray) ToAdditionalUnattendContentArrayOutputWithContext(ctx context.Context) AdditionalUnattendContentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalUnattendContentArrayOutput)
}

type AdditionalUnattendContentOutput struct{ *pulumi.OutputState }

func (AdditionalUnattendContentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalUnattendContent)(nil)).Elem()
}

func (o AdditionalUnattendContentOutput) ToAdditionalUnattendContentOutput() AdditionalUnattendContentOutput {
	return o
}

func (o AdditionalUnattendContentOutput) ToAdditionalUnattendContentOutputWithContext(ctx context.Context) AdditionalUnattendContentOutput {
	return o
}

func (o AdditionalUnattendContentOutput) ComponentName() ComponentNamesPtrOutput {
	return o.ApplyT(func(v AdditionalUnattendContent) *ComponentNames { return v.ComponentName }).(ComponentNamesPtrOutput)
}

func (o AdditionalUnattendContentOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalUnattendContent) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o AdditionalUnattendContentOutput) PassName() PassNamesPtrOutput {
	return o.ApplyT(func(v AdditionalUnattendContent) *PassNames { return v.PassName }).(PassNamesPtrOutput)
}

func (o AdditionalUnattendContentOutput) SettingName() SettingNamesPtrOutput {
	return o.ApplyT(func(v AdditionalUnattendContent) *SettingNames { return v.SettingName }).(SettingNamesPtrOutput)
}

type AdditionalUnattendContentArrayOutput struct{ *pulumi.OutputState }

func (AdditionalUnattendContentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalUnattendContent)(nil)).Elem()
}

func (o AdditionalUnattendContentArrayOutput) ToAdditionalUnattendContentArrayOutput() AdditionalUnattendContentArrayOutput {
	return o
}

func (o AdditionalUnattendContentArrayOutput) ToAdditionalUnattendContentArrayOutputWithContext(ctx context.Context) AdditionalUnattendContentArrayOutput {
	return o
}

func (o AdditionalUnattendContentArrayOutput) Index(i pulumi.IntInput) AdditionalUnattendContentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AdditionalUnattendContent {
		return vs[0].([]AdditionalUnattendContent)[vs[1].(int)]
	}).(AdditionalUnattendContentOutput)
}

type AdditionalUnattendContentResponse struct {
	ComponentName *string `pulumi:"componentName"`
	Content       *string `pulumi:"content"`
	PassName      *string `pulumi:"passName"`
	SettingName   *string `pulumi:"settingName"`
}

type AdditionalUnattendContentResponseOutput struct{ *pulumi.OutputState }

func (AdditionalUnattendContentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalUnattendContentResponse)(nil)).Elem()
}

func (o AdditionalUnattendContentResponseOutput) ToAdditionalUnattendContentResponseOutput() AdditionalUnattendContentResponseOutput {
	return o
}

func (o AdditionalUnattendContentResponseOutput) ToAdditionalUnattendContentResponseOutputWithContext(ctx context.Context) AdditionalUnattendContentResponseOutput {
	return o
}

func (o AdditionalUnattendContentResponseOutput) ComponentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalUnattendContentResponse) *string { return v.ComponentName }).(pulumi.StringPtrOutput)
}

func (o AdditionalUnattendContentResponseOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalUnattendContentResponse) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o AdditionalUnattendContentResponseOutput) PassName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalUnattendContentResponse) *string { return v.PassName }).(pulumi.StringPtrOutput)
}

func (o AdditionalUnattendContentResponseOutput) SettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalUnattendContentResponse) *string { return v.SettingName }).(pulumi.StringPtrOutput)
}

type AdditionalUnattendContentResponseArrayOutput struct{ *pulumi.OutputState }

func (AdditionalUnattendContentResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalUnattendContentResponse)(nil)).Elem()
}

func (o AdditionalUnattendContentResponseArrayOutput) ToAdditionalUnattendContentResponseArrayOutput() AdditionalUnattendContentResponseArrayOutput {
	return o
}

func (o AdditionalUnattendContentResponseArrayOutput) ToAdditionalUnattendContentResponseArrayOutputWithContext(ctx context.Context) AdditionalUnattendContentResponseArrayOutput {
	return o
}

func (o AdditionalUnattendContentResponseArrayOutput) Index(i pulumi.IntInput) AdditionalUnattendContentResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AdditionalUnattendContentResponse {
		return vs[0].([]AdditionalUnattendContentResponse)[vs[1].(int)]
	}).(AdditionalUnattendContentResponseOutput)
}

type ApiEntityReference struct {
	Id *string `pulumi:"id"`
}





type ApiEntityReferenceInput interface {
	pulumi.Input

	ToApiEntityReferenceOutput() ApiEntityReferenceOutput
	ToApiEntityReferenceOutputWithContext(context.Context) ApiEntityReferenceOutput
}

type ApiEntityReferenceArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ApiEntityReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiEntityReference)(nil)).Elem()
}

func (i ApiEntityReferenceArgs) ToApiEntityReferenceOutput() ApiEntityReferenceOutput {
	return i.ToApiEntityReferenceOutputWithContext(context.Background())
}

func (i ApiEntityReferenceArgs) ToApiEntityReferenceOutputWithContext(ctx context.Context) ApiEntityReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiEntityReferenceOutput)
}

type ApiEntityReferenceOutput struct{ *pulumi.OutputState }

func (ApiEntityReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiEntityReference)(nil)).Elem()
}

func (o ApiEntityReferenceOutput) ToApiEntityReferenceOutput() ApiEntityReferenceOutput {
	return o
}

func (o ApiEntityReferenceOutput) ToApiEntityReferenceOutputWithContext(ctx context.Context) ApiEntityReferenceOutput {
	return o
}

func (o ApiEntityReferenceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEntityReference) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ApiEntityReferenceResponse struct {
	Id *string `pulumi:"id"`
}

type ApiEntityReferenceResponseOutput struct{ *pulumi.OutputState }

func (ApiEntityReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiEntityReferenceResponse)(nil)).Elem()
}

func (o ApiEntityReferenceResponseOutput) ToApiEntityReferenceResponseOutput() ApiEntityReferenceResponseOutput {
	return o
}

func (o ApiEntityReferenceResponseOutput) ToApiEntityReferenceResponseOutputWithContext(ctx context.Context) ApiEntityReferenceResponseOutput {
	return o
}

func (o ApiEntityReferenceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEntityReferenceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type BootDiagnostics struct {
	Enabled    *bool   `pulumi:"enabled"`
	StorageUri *string `pulumi:"storageUri"`
}





type BootDiagnosticsInput interface {
	pulumi.Input

	ToBootDiagnosticsOutput() BootDiagnosticsOutput
	ToBootDiagnosticsOutputWithContext(context.Context) BootDiagnosticsOutput
}

type BootDiagnosticsArgs struct {
	Enabled    pulumi.BoolPtrInput   `pulumi:"enabled"`
	StorageUri pulumi.StringPtrInput `pulumi:"storageUri"`
}

func (BootDiagnosticsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BootDiagnostics)(nil)).Elem()
}

func (i BootDiagnosticsArgs) ToBootDiagnosticsOutput() BootDiagnosticsOutput {
	return i.ToBootDiagnosticsOutputWithContext(context.Background())
}

func (i BootDiagnosticsArgs) ToBootDiagnosticsOutputWithContext(ctx context.Context) BootDiagnosticsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootDiagnosticsOutput)
}

func (i BootDiagnosticsArgs) ToBootDiagnosticsPtrOutput() BootDiagnosticsPtrOutput {
	return i.ToBootDiagnosticsPtrOutputWithContext(context.Background())
}

func (i BootDiagnosticsArgs) ToBootDiagnosticsPtrOutputWithContext(ctx context.Context) BootDiagnosticsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootDiagnosticsOutput).ToBootDiagnosticsPtrOutputWithContext(ctx)
}









type BootDiagnosticsPtrInput interface {
	pulumi.Input

	ToBootDiagnosticsPtrOutput() BootDiagnosticsPtrOutput
	ToBootDiagnosticsPtrOutputWithContext(context.Context) BootDiagnosticsPtrOutput
}

type bootDiagnosticsPtrType BootDiagnosticsArgs

func BootDiagnosticsPtr(v *BootDiagnosticsArgs) BootDiagnosticsPtrInput {
	return (*bootDiagnosticsPtrType)(v)
}

func (*bootDiagnosticsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BootDiagnostics)(nil)).Elem()
}

func (i *bootDiagnosticsPtrType) ToBootDiagnosticsPtrOutput() BootDiagnosticsPtrOutput {
	return i.ToBootDiagnosticsPtrOutputWithContext(context.Background())
}

func (i *bootDiagnosticsPtrType) ToBootDiagnosticsPtrOutputWithContext(ctx context.Context) BootDiagnosticsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootDiagnosticsPtrOutput)
}

type BootDiagnosticsOutput struct{ *pulumi.OutputState }

func (BootDiagnosticsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BootDiagnostics)(nil)).Elem()
}

func (o BootDiagnosticsOutput) ToBootDiagnosticsOutput() BootDiagnosticsOutput {
	return o
}

func (o BootDiagnosticsOutput) ToBootDiagnosticsOutputWithContext(ctx context.Context) BootDiagnosticsOutput {
	return o
}

func (o BootDiagnosticsOutput) ToBootDiagnosticsPtrOutput() BootDiagnosticsPtrOutput {
	return o.ToBootDiagnosticsPtrOutputWithContext(context.Background())
}

func (o BootDiagnosticsOutput) ToBootDiagnosticsPtrOutputWithContext(ctx context.Context) BootDiagnosticsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BootDiagnostics) *BootDiagnostics {
		return &v
	}).(BootDiagnosticsPtrOutput)
}

func (o BootDiagnosticsOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BootDiagnostics) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o BootDiagnosticsOutput) StorageUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BootDiagnostics) *string { return v.StorageUri }).(pulumi.StringPtrOutput)
}

type BootDiagnosticsPtrOutput struct{ *pulumi.OutputState }

func (BootDiagnosticsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BootDiagnostics)(nil)).Elem()
}

func (o BootDiagnosticsPtrOutput) ToBootDiagnosticsPtrOutput() BootDiagnosticsPtrOutput {
	return o
}

func (o BootDiagnosticsPtrOutput) ToBootDiagnosticsPtrOutputWithContext(ctx context.Context) BootDiagnosticsPtrOutput {
	return o
}

func (o BootDiagnosticsPtrOutput) Elem() BootDiagnosticsOutput {
	return o.ApplyT(func(v *BootDiagnostics) BootDiagnostics {
		if v != nil {
			return *v
		}
		var ret BootDiagnostics
		return ret
	}).(BootDiagnosticsOutput)
}

func (o BootDiagnosticsPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BootDiagnostics) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o BootDiagnosticsPtrOutput) StorageUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BootDiagnostics) *string {
		if v == nil {
			return nil
		}
		return v.StorageUri
	}).(pulumi.StringPtrOutput)
}

type BootDiagnosticsInstanceViewResponse struct {
	ConsoleScreenshotBlobUri string `pulumi:"consoleScreenshotBlobUri"`
	SerialConsoleLogBlobUri  string `pulumi:"serialConsoleLogBlobUri"`
}

type BootDiagnosticsInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (BootDiagnosticsInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BootDiagnosticsInstanceViewResponse)(nil)).Elem()
}

func (o BootDiagnosticsInstanceViewResponseOutput) ToBootDiagnosticsInstanceViewResponseOutput() BootDiagnosticsInstanceViewResponseOutput {
	return o
}

func (o BootDiagnosticsInstanceViewResponseOutput) ToBootDiagnosticsInstanceViewResponseOutputWithContext(ctx context.Context) BootDiagnosticsInstanceViewResponseOutput {
	return o
}

func (o BootDiagnosticsInstanceViewResponseOutput) ConsoleScreenshotBlobUri() pulumi.StringOutput {
	return o.ApplyT(func(v BootDiagnosticsInstanceViewResponse) string { return v.ConsoleScreenshotBlobUri }).(pulumi.StringOutput)
}

func (o BootDiagnosticsInstanceViewResponseOutput) SerialConsoleLogBlobUri() pulumi.StringOutput {
	return o.ApplyT(func(v BootDiagnosticsInstanceViewResponse) string { return v.SerialConsoleLogBlobUri }).(pulumi.StringOutput)
}

type BootDiagnosticsInstanceViewResponsePtrOutput struct{ *pulumi.OutputState }

func (BootDiagnosticsInstanceViewResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BootDiagnosticsInstanceViewResponse)(nil)).Elem()
}

func (o BootDiagnosticsInstanceViewResponsePtrOutput) ToBootDiagnosticsInstanceViewResponsePtrOutput() BootDiagnosticsInstanceViewResponsePtrOutput {
	return o
}

func (o BootDiagnosticsInstanceViewResponsePtrOutput) ToBootDiagnosticsInstanceViewResponsePtrOutputWithContext(ctx context.Context) BootDiagnosticsInstanceViewResponsePtrOutput {
	return o
}

func (o BootDiagnosticsInstanceViewResponsePtrOutput) Elem() BootDiagnosticsInstanceViewResponseOutput {
	return o.ApplyT(func(v *BootDiagnosticsInstanceViewResponse) BootDiagnosticsInstanceViewResponse {
		if v != nil {
			return *v
		}
		var ret BootDiagnosticsInstanceViewResponse
		return ret
	}).(BootDiagnosticsInstanceViewResponseOutput)
}

func (o BootDiagnosticsInstanceViewResponsePtrOutput) ConsoleScreenshotBlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BootDiagnosticsInstanceViewResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ConsoleScreenshotBlobUri
	}).(pulumi.StringPtrOutput)
}

func (o BootDiagnosticsInstanceViewResponsePtrOutput) SerialConsoleLogBlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BootDiagnosticsInstanceViewResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SerialConsoleLogBlobUri
	}).(pulumi.StringPtrOutput)
}

type BootDiagnosticsResponse struct {
	Enabled    *bool   `pulumi:"enabled"`
	StorageUri *string `pulumi:"storageUri"`
}

type BootDiagnosticsResponseOutput struct{ *pulumi.OutputState }

func (BootDiagnosticsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BootDiagnosticsResponse)(nil)).Elem()
}

func (o BootDiagnosticsResponseOutput) ToBootDiagnosticsResponseOutput() BootDiagnosticsResponseOutput {
	return o
}

func (o BootDiagnosticsResponseOutput) ToBootDiagnosticsResponseOutputWithContext(ctx context.Context) BootDiagnosticsResponseOutput {
	return o
}

func (o BootDiagnosticsResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BootDiagnosticsResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o BootDiagnosticsResponseOutput) StorageUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BootDiagnosticsResponse) *string { return v.StorageUri }).(pulumi.StringPtrOutput)
}

type BootDiagnosticsResponsePtrOutput struct{ *pulumi.OutputState }

func (BootDiagnosticsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BootDiagnosticsResponse)(nil)).Elem()
}

func (o BootDiagnosticsResponsePtrOutput) ToBootDiagnosticsResponsePtrOutput() BootDiagnosticsResponsePtrOutput {
	return o
}

func (o BootDiagnosticsResponsePtrOutput) ToBootDiagnosticsResponsePtrOutputWithContext(ctx context.Context) BootDiagnosticsResponsePtrOutput {
	return o
}

func (o BootDiagnosticsResponsePtrOutput) Elem() BootDiagnosticsResponseOutput {
	return o.ApplyT(func(v *BootDiagnosticsResponse) BootDiagnosticsResponse {
		if v != nil {
			return *v
		}
		var ret BootDiagnosticsResponse
		return ret
	}).(BootDiagnosticsResponseOutput)
}

func (o BootDiagnosticsResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BootDiagnosticsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o BootDiagnosticsResponsePtrOutput) StorageUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BootDiagnosticsResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageUri
	}).(pulumi.StringPtrOutput)
}

type DataDisk struct {
	Caching      *CachingTypes         `pulumi:"caching"`
	CreateOption DiskCreateOptionTypes `pulumi:"createOption"`
	DiskSizeGB   *int                  `pulumi:"diskSizeGB"`
	Image        *VirtualHardDisk      `pulumi:"image"`
	Lun          int                   `pulumi:"lun"`
	Name         string                `pulumi:"name"`
	Vhd          VirtualHardDisk       `pulumi:"vhd"`
}





type DataDiskInput interface {
	pulumi.Input

	ToDataDiskOutput() DataDiskOutput
	ToDataDiskOutputWithContext(context.Context) DataDiskOutput
}

type DataDiskArgs struct {
	Caching      CachingTypesPtrInput       `pulumi:"caching"`
	CreateOption DiskCreateOptionTypesInput `pulumi:"createOption"`
	DiskSizeGB   pulumi.IntPtrInput         `pulumi:"diskSizeGB"`
	Image        VirtualHardDiskPtrInput    `pulumi:"image"`
	Lun          pulumi.IntInput            `pulumi:"lun"`
	Name         pulumi.StringInput         `pulumi:"name"`
	Vhd          VirtualHardDiskInput       `pulumi:"vhd"`
}

func (DataDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDisk)(nil)).Elem()
}

func (i DataDiskArgs) ToDataDiskOutput() DataDiskOutput {
	return i.ToDataDiskOutputWithContext(context.Background())
}

func (i DataDiskArgs) ToDataDiskOutputWithContext(ctx context.Context) DataDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataDiskOutput)
}





type DataDiskArrayInput interface {
	pulumi.Input

	ToDataDiskArrayOutput() DataDiskArrayOutput
	ToDataDiskArrayOutputWithContext(context.Context) DataDiskArrayOutput
}

type DataDiskArray []DataDiskInput

func (DataDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDisk)(nil)).Elem()
}

func (i DataDiskArray) ToDataDiskArrayOutput() DataDiskArrayOutput {
	return i.ToDataDiskArrayOutputWithContext(context.Background())
}

func (i DataDiskArray) ToDataDiskArrayOutputWithContext(ctx context.Context) DataDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataDiskArrayOutput)
}

type DataDiskOutput struct{ *pulumi.OutputState }

func (DataDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDisk)(nil)).Elem()
}

func (o DataDiskOutput) ToDataDiskOutput() DataDiskOutput {
	return o
}

func (o DataDiskOutput) ToDataDiskOutputWithContext(ctx context.Context) DataDiskOutput {
	return o
}

func (o DataDiskOutput) Caching() CachingTypesPtrOutput {
	return o.ApplyT(func(v DataDisk) *CachingTypes { return v.Caching }).(CachingTypesPtrOutput)
}

func (o DataDiskOutput) CreateOption() DiskCreateOptionTypesOutput {
	return o.ApplyT(func(v DataDisk) DiskCreateOptionTypes { return v.CreateOption }).(DiskCreateOptionTypesOutput)
}

func (o DataDiskOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataDisk) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o DataDiskOutput) Image() VirtualHardDiskPtrOutput {
	return o.ApplyT(func(v DataDisk) *VirtualHardDisk { return v.Image }).(VirtualHardDiskPtrOutput)
}

func (o DataDiskOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v DataDisk) int { return v.Lun }).(pulumi.IntOutput)
}

func (o DataDiskOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DataDisk) string { return v.Name }).(pulumi.StringOutput)
}

func (o DataDiskOutput) Vhd() VirtualHardDiskOutput {
	return o.ApplyT(func(v DataDisk) VirtualHardDisk { return v.Vhd }).(VirtualHardDiskOutput)
}

type DataDiskArrayOutput struct{ *pulumi.OutputState }

func (DataDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDisk)(nil)).Elem()
}

func (o DataDiskArrayOutput) ToDataDiskArrayOutput() DataDiskArrayOutput {
	return o
}

func (o DataDiskArrayOutput) ToDataDiskArrayOutputWithContext(ctx context.Context) DataDiskArrayOutput {
	return o
}

func (o DataDiskArrayOutput) Index(i pulumi.IntInput) DataDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataDisk {
		return vs[0].([]DataDisk)[vs[1].(int)]
	}).(DataDiskOutput)
}

type DataDiskResponse struct {
	Caching      *string                  `pulumi:"caching"`
	CreateOption string                   `pulumi:"createOption"`
	DiskSizeGB   *int                     `pulumi:"diskSizeGB"`
	Image        *VirtualHardDiskResponse `pulumi:"image"`
	Lun          int                      `pulumi:"lun"`
	Name         string                   `pulumi:"name"`
	Vhd          VirtualHardDiskResponse  `pulumi:"vhd"`
}

type DataDiskResponseOutput struct{ *pulumi.OutputState }

func (DataDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDiskResponse)(nil)).Elem()
}

func (o DataDiskResponseOutput) ToDataDiskResponseOutput() DataDiskResponseOutput {
	return o
}

func (o DataDiskResponseOutput) ToDataDiskResponseOutputWithContext(ctx context.Context) DataDiskResponseOutput {
	return o
}

func (o DataDiskResponseOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *string { return v.Caching }).(pulumi.StringPtrOutput)
}

func (o DataDiskResponseOutput) CreateOption() pulumi.StringOutput {
	return o.ApplyT(func(v DataDiskResponse) string { return v.CreateOption }).(pulumi.StringOutput)
}

func (o DataDiskResponseOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o DataDiskResponseOutput) Image() VirtualHardDiskResponsePtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *VirtualHardDiskResponse { return v.Image }).(VirtualHardDiskResponsePtrOutput)
}

func (o DataDiskResponseOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v DataDiskResponse) int { return v.Lun }).(pulumi.IntOutput)
}

func (o DataDiskResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DataDiskResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DataDiskResponseOutput) Vhd() VirtualHardDiskResponseOutput {
	return o.ApplyT(func(v DataDiskResponse) VirtualHardDiskResponse { return v.Vhd }).(VirtualHardDiskResponseOutput)
}

type DataDiskResponseArrayOutput struct{ *pulumi.OutputState }

func (DataDiskResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDiskResponse)(nil)).Elem()
}

func (o DataDiskResponseArrayOutput) ToDataDiskResponseArrayOutput() DataDiskResponseArrayOutput {
	return o
}

func (o DataDiskResponseArrayOutput) ToDataDiskResponseArrayOutputWithContext(ctx context.Context) DataDiskResponseArrayOutput {
	return o
}

func (o DataDiskResponseArrayOutput) Index(i pulumi.IntInput) DataDiskResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataDiskResponse {
		return vs[0].([]DataDiskResponse)[vs[1].(int)]
	}).(DataDiskResponseOutput)
}

type DiagnosticsProfile struct {
	BootDiagnostics *BootDiagnostics `pulumi:"bootDiagnostics"`
}





type DiagnosticsProfileInput interface {
	pulumi.Input

	ToDiagnosticsProfileOutput() DiagnosticsProfileOutput
	ToDiagnosticsProfileOutputWithContext(context.Context) DiagnosticsProfileOutput
}

type DiagnosticsProfileArgs struct {
	BootDiagnostics BootDiagnosticsPtrInput `pulumi:"bootDiagnostics"`
}

func (DiagnosticsProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsProfile)(nil)).Elem()
}

func (i DiagnosticsProfileArgs) ToDiagnosticsProfileOutput() DiagnosticsProfileOutput {
	return i.ToDiagnosticsProfileOutputWithContext(context.Background())
}

func (i DiagnosticsProfileArgs) ToDiagnosticsProfileOutputWithContext(ctx context.Context) DiagnosticsProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsProfileOutput)
}

func (i DiagnosticsProfileArgs) ToDiagnosticsProfilePtrOutput() DiagnosticsProfilePtrOutput {
	return i.ToDiagnosticsProfilePtrOutputWithContext(context.Background())
}

func (i DiagnosticsProfileArgs) ToDiagnosticsProfilePtrOutputWithContext(ctx context.Context) DiagnosticsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsProfileOutput).ToDiagnosticsProfilePtrOutputWithContext(ctx)
}









type DiagnosticsProfilePtrInput interface {
	pulumi.Input

	ToDiagnosticsProfilePtrOutput() DiagnosticsProfilePtrOutput
	ToDiagnosticsProfilePtrOutputWithContext(context.Context) DiagnosticsProfilePtrOutput
}

type diagnosticsProfilePtrType DiagnosticsProfileArgs

func DiagnosticsProfilePtr(v *DiagnosticsProfileArgs) DiagnosticsProfilePtrInput {
	return (*diagnosticsProfilePtrType)(v)
}

func (*diagnosticsProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsProfile)(nil)).Elem()
}

func (i *diagnosticsProfilePtrType) ToDiagnosticsProfilePtrOutput() DiagnosticsProfilePtrOutput {
	return i.ToDiagnosticsProfilePtrOutputWithContext(context.Background())
}

func (i *diagnosticsProfilePtrType) ToDiagnosticsProfilePtrOutputWithContext(ctx context.Context) DiagnosticsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsProfilePtrOutput)
}

type DiagnosticsProfileOutput struct{ *pulumi.OutputState }

func (DiagnosticsProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsProfile)(nil)).Elem()
}

func (o DiagnosticsProfileOutput) ToDiagnosticsProfileOutput() DiagnosticsProfileOutput {
	return o
}

func (o DiagnosticsProfileOutput) ToDiagnosticsProfileOutputWithContext(ctx context.Context) DiagnosticsProfileOutput {
	return o
}

func (o DiagnosticsProfileOutput) ToDiagnosticsProfilePtrOutput() DiagnosticsProfilePtrOutput {
	return o.ToDiagnosticsProfilePtrOutputWithContext(context.Background())
}

func (o DiagnosticsProfileOutput) ToDiagnosticsProfilePtrOutputWithContext(ctx context.Context) DiagnosticsProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiagnosticsProfile) *DiagnosticsProfile {
		return &v
	}).(DiagnosticsProfilePtrOutput)
}

func (o DiagnosticsProfileOutput) BootDiagnostics() BootDiagnosticsPtrOutput {
	return o.ApplyT(func(v DiagnosticsProfile) *BootDiagnostics { return v.BootDiagnostics }).(BootDiagnosticsPtrOutput)
}

type DiagnosticsProfilePtrOutput struct{ *pulumi.OutputState }

func (DiagnosticsProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsProfile)(nil)).Elem()
}

func (o DiagnosticsProfilePtrOutput) ToDiagnosticsProfilePtrOutput() DiagnosticsProfilePtrOutput {
	return o
}

func (o DiagnosticsProfilePtrOutput) ToDiagnosticsProfilePtrOutputWithContext(ctx context.Context) DiagnosticsProfilePtrOutput {
	return o
}

func (o DiagnosticsProfilePtrOutput) Elem() DiagnosticsProfileOutput {
	return o.ApplyT(func(v *DiagnosticsProfile) DiagnosticsProfile {
		if v != nil {
			return *v
		}
		var ret DiagnosticsProfile
		return ret
	}).(DiagnosticsProfileOutput)
}

func (o DiagnosticsProfilePtrOutput) BootDiagnostics() BootDiagnosticsPtrOutput {
	return o.ApplyT(func(v *DiagnosticsProfile) *BootDiagnostics {
		if v == nil {
			return nil
		}
		return v.BootDiagnostics
	}).(BootDiagnosticsPtrOutput)
}

type DiagnosticsProfileResponse struct {
	BootDiagnostics *BootDiagnosticsResponse `pulumi:"bootDiagnostics"`
}

type DiagnosticsProfileResponseOutput struct{ *pulumi.OutputState }

func (DiagnosticsProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsProfileResponse)(nil)).Elem()
}

func (o DiagnosticsProfileResponseOutput) ToDiagnosticsProfileResponseOutput() DiagnosticsProfileResponseOutput {
	return o
}

func (o DiagnosticsProfileResponseOutput) ToDiagnosticsProfileResponseOutputWithContext(ctx context.Context) DiagnosticsProfileResponseOutput {
	return o
}

func (o DiagnosticsProfileResponseOutput) BootDiagnostics() BootDiagnosticsResponsePtrOutput {
	return o.ApplyT(func(v DiagnosticsProfileResponse) *BootDiagnosticsResponse { return v.BootDiagnostics }).(BootDiagnosticsResponsePtrOutput)
}

type DiagnosticsProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (DiagnosticsProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsProfileResponse)(nil)).Elem()
}

func (o DiagnosticsProfileResponsePtrOutput) ToDiagnosticsProfileResponsePtrOutput() DiagnosticsProfileResponsePtrOutput {
	return o
}

func (o DiagnosticsProfileResponsePtrOutput) ToDiagnosticsProfileResponsePtrOutputWithContext(ctx context.Context) DiagnosticsProfileResponsePtrOutput {
	return o
}

func (o DiagnosticsProfileResponsePtrOutput) Elem() DiagnosticsProfileResponseOutput {
	return o.ApplyT(func(v *DiagnosticsProfileResponse) DiagnosticsProfileResponse {
		if v != nil {
			return *v
		}
		var ret DiagnosticsProfileResponse
		return ret
	}).(DiagnosticsProfileResponseOutput)
}

func (o DiagnosticsProfileResponsePtrOutput) BootDiagnostics() BootDiagnosticsResponsePtrOutput {
	return o.ApplyT(func(v *DiagnosticsProfileResponse) *BootDiagnosticsResponse {
		if v == nil {
			return nil
		}
		return v.BootDiagnostics
	}).(BootDiagnosticsResponsePtrOutput)
}

type DiskEncryptionSettings struct {
	DiskEncryptionKey KeyVaultSecretReference `pulumi:"diskEncryptionKey"`
	Enabled           *bool                   `pulumi:"enabled"`
	KeyEncryptionKey  *KeyVaultKeyReference   `pulumi:"keyEncryptionKey"`
}





type DiskEncryptionSettingsInput interface {
	pulumi.Input

	ToDiskEncryptionSettingsOutput() DiskEncryptionSettingsOutput
	ToDiskEncryptionSettingsOutputWithContext(context.Context) DiskEncryptionSettingsOutput
}

type DiskEncryptionSettingsArgs struct {
	DiskEncryptionKey KeyVaultSecretReferenceInput `pulumi:"diskEncryptionKey"`
	Enabled           pulumi.BoolPtrInput          `pulumi:"enabled"`
	KeyEncryptionKey  KeyVaultKeyReferencePtrInput `pulumi:"keyEncryptionKey"`
}

func (DiskEncryptionSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSettings)(nil)).Elem()
}

func (i DiskEncryptionSettingsArgs) ToDiskEncryptionSettingsOutput() DiskEncryptionSettingsOutput {
	return i.ToDiskEncryptionSettingsOutputWithContext(context.Background())
}

func (i DiskEncryptionSettingsArgs) ToDiskEncryptionSettingsOutputWithContext(ctx context.Context) DiskEncryptionSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSettingsOutput)
}

func (i DiskEncryptionSettingsArgs) ToDiskEncryptionSettingsPtrOutput() DiskEncryptionSettingsPtrOutput {
	return i.ToDiskEncryptionSettingsPtrOutputWithContext(context.Background())
}

func (i DiskEncryptionSettingsArgs) ToDiskEncryptionSettingsPtrOutputWithContext(ctx context.Context) DiskEncryptionSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSettingsOutput).ToDiskEncryptionSettingsPtrOutputWithContext(ctx)
}









type DiskEncryptionSettingsPtrInput interface {
	pulumi.Input

	ToDiskEncryptionSettingsPtrOutput() DiskEncryptionSettingsPtrOutput
	ToDiskEncryptionSettingsPtrOutputWithContext(context.Context) DiskEncryptionSettingsPtrOutput
}

type diskEncryptionSettingsPtrType DiskEncryptionSettingsArgs

func DiskEncryptionSettingsPtr(v *DiskEncryptionSettingsArgs) DiskEncryptionSettingsPtrInput {
	return (*diskEncryptionSettingsPtrType)(v)
}

func (*diskEncryptionSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionSettings)(nil)).Elem()
}

func (i *diskEncryptionSettingsPtrType) ToDiskEncryptionSettingsPtrOutput() DiskEncryptionSettingsPtrOutput {
	return i.ToDiskEncryptionSettingsPtrOutputWithContext(context.Background())
}

func (i *diskEncryptionSettingsPtrType) ToDiskEncryptionSettingsPtrOutputWithContext(ctx context.Context) DiskEncryptionSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSettingsPtrOutput)
}

type DiskEncryptionSettingsOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSettings)(nil)).Elem()
}

func (o DiskEncryptionSettingsOutput) ToDiskEncryptionSettingsOutput() DiskEncryptionSettingsOutput {
	return o
}

func (o DiskEncryptionSettingsOutput) ToDiskEncryptionSettingsOutputWithContext(ctx context.Context) DiskEncryptionSettingsOutput {
	return o
}

func (o DiskEncryptionSettingsOutput) ToDiskEncryptionSettingsPtrOutput() DiskEncryptionSettingsPtrOutput {
	return o.ToDiskEncryptionSettingsPtrOutputWithContext(context.Background())
}

func (o DiskEncryptionSettingsOutput) ToDiskEncryptionSettingsPtrOutputWithContext(ctx context.Context) DiskEncryptionSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskEncryptionSettings) *DiskEncryptionSettings {
		return &v
	}).(DiskEncryptionSettingsPtrOutput)
}

func (o DiskEncryptionSettingsOutput) DiskEncryptionKey() KeyVaultSecretReferenceOutput {
	return o.ApplyT(func(v DiskEncryptionSettings) KeyVaultSecretReference { return v.DiskEncryptionKey }).(KeyVaultSecretReferenceOutput)
}

func (o DiskEncryptionSettingsOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DiskEncryptionSettings) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o DiskEncryptionSettingsOutput) KeyEncryptionKey() KeyVaultKeyReferencePtrOutput {
	return o.ApplyT(func(v DiskEncryptionSettings) *KeyVaultKeyReference { return v.KeyEncryptionKey }).(KeyVaultKeyReferencePtrOutput)
}

type DiskEncryptionSettingsPtrOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionSettings)(nil)).Elem()
}

func (o DiskEncryptionSettingsPtrOutput) ToDiskEncryptionSettingsPtrOutput() DiskEncryptionSettingsPtrOutput {
	return o
}

func (o DiskEncryptionSettingsPtrOutput) ToDiskEncryptionSettingsPtrOutputWithContext(ctx context.Context) DiskEncryptionSettingsPtrOutput {
	return o
}

func (o DiskEncryptionSettingsPtrOutput) Elem() DiskEncryptionSettingsOutput {
	return o.ApplyT(func(v *DiskEncryptionSettings) DiskEncryptionSettings {
		if v != nil {
			return *v
		}
		var ret DiskEncryptionSettings
		return ret
	}).(DiskEncryptionSettingsOutput)
}

func (o DiskEncryptionSettingsPtrOutput) DiskEncryptionKey() KeyVaultSecretReferencePtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSettings) *KeyVaultSecretReference {
		if v == nil {
			return nil
		}
		return &v.DiskEncryptionKey
	}).(KeyVaultSecretReferencePtrOutput)
}

func (o DiskEncryptionSettingsPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSettings) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o DiskEncryptionSettingsPtrOutput) KeyEncryptionKey() KeyVaultKeyReferencePtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSettings) *KeyVaultKeyReference {
		if v == nil {
			return nil
		}
		return v.KeyEncryptionKey
	}).(KeyVaultKeyReferencePtrOutput)
}

type DiskEncryptionSettingsResponse struct {
	DiskEncryptionKey KeyVaultSecretReferenceResponse `pulumi:"diskEncryptionKey"`
	Enabled           *bool                           `pulumi:"enabled"`
	KeyEncryptionKey  *KeyVaultKeyReferenceResponse   `pulumi:"keyEncryptionKey"`
}

type DiskEncryptionSettingsResponseOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSettingsResponse)(nil)).Elem()
}

func (o DiskEncryptionSettingsResponseOutput) ToDiskEncryptionSettingsResponseOutput() DiskEncryptionSettingsResponseOutput {
	return o
}

func (o DiskEncryptionSettingsResponseOutput) ToDiskEncryptionSettingsResponseOutputWithContext(ctx context.Context) DiskEncryptionSettingsResponseOutput {
	return o
}

func (o DiskEncryptionSettingsResponseOutput) DiskEncryptionKey() KeyVaultSecretReferenceResponseOutput {
	return o.ApplyT(func(v DiskEncryptionSettingsResponse) KeyVaultSecretReferenceResponse { return v.DiskEncryptionKey }).(KeyVaultSecretReferenceResponseOutput)
}

func (o DiskEncryptionSettingsResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DiskEncryptionSettingsResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o DiskEncryptionSettingsResponseOutput) KeyEncryptionKey() KeyVaultKeyReferenceResponsePtrOutput {
	return o.ApplyT(func(v DiskEncryptionSettingsResponse) *KeyVaultKeyReferenceResponse { return v.KeyEncryptionKey }).(KeyVaultKeyReferenceResponsePtrOutput)
}

type DiskEncryptionSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionSettingsResponse)(nil)).Elem()
}

func (o DiskEncryptionSettingsResponsePtrOutput) ToDiskEncryptionSettingsResponsePtrOutput() DiskEncryptionSettingsResponsePtrOutput {
	return o
}

func (o DiskEncryptionSettingsResponsePtrOutput) ToDiskEncryptionSettingsResponsePtrOutputWithContext(ctx context.Context) DiskEncryptionSettingsResponsePtrOutput {
	return o
}

func (o DiskEncryptionSettingsResponsePtrOutput) Elem() DiskEncryptionSettingsResponseOutput {
	return o.ApplyT(func(v *DiskEncryptionSettingsResponse) DiskEncryptionSettingsResponse {
		if v != nil {
			return *v
		}
		var ret DiskEncryptionSettingsResponse
		return ret
	}).(DiskEncryptionSettingsResponseOutput)
}

func (o DiskEncryptionSettingsResponsePtrOutput) DiskEncryptionKey() KeyVaultSecretReferenceResponsePtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSettingsResponse) *KeyVaultSecretReferenceResponse {
		if v == nil {
			return nil
		}
		return &v.DiskEncryptionKey
	}).(KeyVaultSecretReferenceResponsePtrOutput)
}

func (o DiskEncryptionSettingsResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o DiskEncryptionSettingsResponsePtrOutput) KeyEncryptionKey() KeyVaultKeyReferenceResponsePtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSettingsResponse) *KeyVaultKeyReferenceResponse {
		if v == nil {
			return nil
		}
		return v.KeyEncryptionKey
	}).(KeyVaultKeyReferenceResponsePtrOutput)
}

type DiskInstanceViewResponse struct {
	Name     *string                      `pulumi:"name"`
	Statuses []InstanceViewStatusResponse `pulumi:"statuses"`
}

type DiskInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (DiskInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskInstanceViewResponse)(nil)).Elem()
}

func (o DiskInstanceViewResponseOutput) ToDiskInstanceViewResponseOutput() DiskInstanceViewResponseOutput {
	return o
}

func (o DiskInstanceViewResponseOutput) ToDiskInstanceViewResponseOutputWithContext(ctx context.Context) DiskInstanceViewResponseOutput {
	return o
}

func (o DiskInstanceViewResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskInstanceViewResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DiskInstanceViewResponseOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v DiskInstanceViewResponse) []InstanceViewStatusResponse { return v.Statuses }).(InstanceViewStatusResponseArrayOutput)
}

type DiskInstanceViewResponseArrayOutput struct{ *pulumi.OutputState }

func (DiskInstanceViewResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiskInstanceViewResponse)(nil)).Elem()
}

func (o DiskInstanceViewResponseArrayOutput) ToDiskInstanceViewResponseArrayOutput() DiskInstanceViewResponseArrayOutput {
	return o
}

func (o DiskInstanceViewResponseArrayOutput) ToDiskInstanceViewResponseArrayOutputWithContext(ctx context.Context) DiskInstanceViewResponseArrayOutput {
	return o
}

func (o DiskInstanceViewResponseArrayOutput) Index(i pulumi.IntInput) DiskInstanceViewResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DiskInstanceViewResponse {
		return vs[0].([]DiskInstanceViewResponse)[vs[1].(int)]
	}).(DiskInstanceViewResponseOutput)
}

type HardwareProfile struct {
	VmSize *string `pulumi:"vmSize"`
}





type HardwareProfileInput interface {
	pulumi.Input

	ToHardwareProfileOutput() HardwareProfileOutput
	ToHardwareProfileOutputWithContext(context.Context) HardwareProfileOutput
}

type HardwareProfileArgs struct {
	VmSize pulumi.StringPtrInput `pulumi:"vmSize"`
}

func (HardwareProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HardwareProfile)(nil)).Elem()
}

func (i HardwareProfileArgs) ToHardwareProfileOutput() HardwareProfileOutput {
	return i.ToHardwareProfileOutputWithContext(context.Background())
}

func (i HardwareProfileArgs) ToHardwareProfileOutputWithContext(ctx context.Context) HardwareProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardwareProfileOutput)
}

func (i HardwareProfileArgs) ToHardwareProfilePtrOutput() HardwareProfilePtrOutput {
	return i.ToHardwareProfilePtrOutputWithContext(context.Background())
}

func (i HardwareProfileArgs) ToHardwareProfilePtrOutputWithContext(ctx context.Context) HardwareProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardwareProfileOutput).ToHardwareProfilePtrOutputWithContext(ctx)
}









type HardwareProfilePtrInput interface {
	pulumi.Input

	ToHardwareProfilePtrOutput() HardwareProfilePtrOutput
	ToHardwareProfilePtrOutputWithContext(context.Context) HardwareProfilePtrOutput
}

type hardwareProfilePtrType HardwareProfileArgs

func HardwareProfilePtr(v *HardwareProfileArgs) HardwareProfilePtrInput {
	return (*hardwareProfilePtrType)(v)
}

func (*hardwareProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HardwareProfile)(nil)).Elem()
}

func (i *hardwareProfilePtrType) ToHardwareProfilePtrOutput() HardwareProfilePtrOutput {
	return i.ToHardwareProfilePtrOutputWithContext(context.Background())
}

func (i *hardwareProfilePtrType) ToHardwareProfilePtrOutputWithContext(ctx context.Context) HardwareProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardwareProfilePtrOutput)
}

type HardwareProfileOutput struct{ *pulumi.OutputState }

func (HardwareProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HardwareProfile)(nil)).Elem()
}

func (o HardwareProfileOutput) ToHardwareProfileOutput() HardwareProfileOutput {
	return o
}

func (o HardwareProfileOutput) ToHardwareProfileOutputWithContext(ctx context.Context) HardwareProfileOutput {
	return o
}

func (o HardwareProfileOutput) ToHardwareProfilePtrOutput() HardwareProfilePtrOutput {
	return o.ToHardwareProfilePtrOutputWithContext(context.Background())
}

func (o HardwareProfileOutput) ToHardwareProfilePtrOutputWithContext(ctx context.Context) HardwareProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HardwareProfile) *HardwareProfile {
		return &v
	}).(HardwareProfilePtrOutput)
}

func (o HardwareProfileOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HardwareProfile) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type HardwareProfilePtrOutput struct{ *pulumi.OutputState }

func (HardwareProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HardwareProfile)(nil)).Elem()
}

func (o HardwareProfilePtrOutput) ToHardwareProfilePtrOutput() HardwareProfilePtrOutput {
	return o
}

func (o HardwareProfilePtrOutput) ToHardwareProfilePtrOutputWithContext(ctx context.Context) HardwareProfilePtrOutput {
	return o
}

func (o HardwareProfilePtrOutput) Elem() HardwareProfileOutput {
	return o.ApplyT(func(v *HardwareProfile) HardwareProfile {
		if v != nil {
			return *v
		}
		var ret HardwareProfile
		return ret
	}).(HardwareProfileOutput)
}

func (o HardwareProfilePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HardwareProfile) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

type HardwareProfileResponse struct {
	VmSize *string `pulumi:"vmSize"`
}

type HardwareProfileResponseOutput struct{ *pulumi.OutputState }

func (HardwareProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HardwareProfileResponse)(nil)).Elem()
}

func (o HardwareProfileResponseOutput) ToHardwareProfileResponseOutput() HardwareProfileResponseOutput {
	return o
}

func (o HardwareProfileResponseOutput) ToHardwareProfileResponseOutputWithContext(ctx context.Context) HardwareProfileResponseOutput {
	return o
}

func (o HardwareProfileResponseOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HardwareProfileResponse) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type HardwareProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (HardwareProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HardwareProfileResponse)(nil)).Elem()
}

func (o HardwareProfileResponsePtrOutput) ToHardwareProfileResponsePtrOutput() HardwareProfileResponsePtrOutput {
	return o
}

func (o HardwareProfileResponsePtrOutput) ToHardwareProfileResponsePtrOutputWithContext(ctx context.Context) HardwareProfileResponsePtrOutput {
	return o
}

func (o HardwareProfileResponsePtrOutput) Elem() HardwareProfileResponseOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) HardwareProfileResponse {
		if v != nil {
			return *v
		}
		var ret HardwareProfileResponse
		return ret
	}).(HardwareProfileResponseOutput)
}

func (o HardwareProfileResponsePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

type ImageReference struct {
	Offer     *string `pulumi:"offer"`
	Publisher *string `pulumi:"publisher"`
	Sku       *string `pulumi:"sku"`
	Version   *string `pulumi:"version"`
}





type ImageReferenceInput interface {
	pulumi.Input

	ToImageReferenceOutput() ImageReferenceOutput
	ToImageReferenceOutputWithContext(context.Context) ImageReferenceOutput
}

type ImageReferenceArgs struct {
	Offer     pulumi.StringPtrInput `pulumi:"offer"`
	Publisher pulumi.StringPtrInput `pulumi:"publisher"`
	Sku       pulumi.StringPtrInput `pulumi:"sku"`
	Version   pulumi.StringPtrInput `pulumi:"version"`
}

func (ImageReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageReference)(nil)).Elem()
}

func (i ImageReferenceArgs) ToImageReferenceOutput() ImageReferenceOutput {
	return i.ToImageReferenceOutputWithContext(context.Background())
}

func (i ImageReferenceArgs) ToImageReferenceOutputWithContext(ctx context.Context) ImageReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageReferenceOutput)
}

func (i ImageReferenceArgs) ToImageReferencePtrOutput() ImageReferencePtrOutput {
	return i.ToImageReferencePtrOutputWithContext(context.Background())
}

func (i ImageReferenceArgs) ToImageReferencePtrOutputWithContext(ctx context.Context) ImageReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageReferenceOutput).ToImageReferencePtrOutputWithContext(ctx)
}









type ImageReferencePtrInput interface {
	pulumi.Input

	ToImageReferencePtrOutput() ImageReferencePtrOutput
	ToImageReferencePtrOutputWithContext(context.Context) ImageReferencePtrOutput
}

type imageReferencePtrType ImageReferenceArgs

func ImageReferencePtr(v *ImageReferenceArgs) ImageReferencePtrInput {
	return (*imageReferencePtrType)(v)
}

func (*imageReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageReference)(nil)).Elem()
}

func (i *imageReferencePtrType) ToImageReferencePtrOutput() ImageReferencePtrOutput {
	return i.ToImageReferencePtrOutputWithContext(context.Background())
}

func (i *imageReferencePtrType) ToImageReferencePtrOutputWithContext(ctx context.Context) ImageReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageReferencePtrOutput)
}

type ImageReferenceOutput struct{ *pulumi.OutputState }

func (ImageReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageReference)(nil)).Elem()
}

func (o ImageReferenceOutput) ToImageReferenceOutput() ImageReferenceOutput {
	return o
}

func (o ImageReferenceOutput) ToImageReferenceOutputWithContext(ctx context.Context) ImageReferenceOutput {
	return o
}

func (o ImageReferenceOutput) ToImageReferencePtrOutput() ImageReferencePtrOutput {
	return o.ToImageReferencePtrOutputWithContext(context.Background())
}

func (o ImageReferenceOutput) ToImageReferencePtrOutputWithContext(ctx context.Context) ImageReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageReference) *ImageReference {
		return &v
	}).(ImageReferencePtrOutput)
}

func (o ImageReferenceOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ImageReferencePtrOutput struct{ *pulumi.OutputState }

func (ImageReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageReference)(nil)).Elem()
}

func (o ImageReferencePtrOutput) ToImageReferencePtrOutput() ImageReferencePtrOutput {
	return o
}

func (o ImageReferencePtrOutput) ToImageReferencePtrOutputWithContext(ctx context.Context) ImageReferencePtrOutput {
	return o
}

func (o ImageReferencePtrOutput) Elem() ImageReferenceOutput {
	return o.ApplyT(func(v *ImageReference) ImageReference {
		if v != nil {
			return *v
		}
		var ret ImageReference
		return ret
	}).(ImageReferenceOutput)
}

func (o ImageReferencePtrOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Offer
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferencePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferencePtrOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferencePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type ImageReferenceResponse struct {
	Offer     *string `pulumi:"offer"`
	Publisher *string `pulumi:"publisher"`
	Sku       *string `pulumi:"sku"`
	Version   *string `pulumi:"version"`
}

type ImageReferenceResponseOutput struct{ *pulumi.OutputState }

func (ImageReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageReferenceResponse)(nil)).Elem()
}

func (o ImageReferenceResponseOutput) ToImageReferenceResponseOutput() ImageReferenceResponseOutput {
	return o
}

func (o ImageReferenceResponseOutput) ToImageReferenceResponseOutputWithContext(ctx context.Context) ImageReferenceResponseOutput {
	return o
}

func (o ImageReferenceResponseOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponseOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ImageReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (ImageReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageReferenceResponse)(nil)).Elem()
}

func (o ImageReferenceResponsePtrOutput) ToImageReferenceResponsePtrOutput() ImageReferenceResponsePtrOutput {
	return o
}

func (o ImageReferenceResponsePtrOutput) ToImageReferenceResponsePtrOutputWithContext(ctx context.Context) ImageReferenceResponsePtrOutput {
	return o
}

func (o ImageReferenceResponsePtrOutput) Elem() ImageReferenceResponseOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) ImageReferenceResponse {
		if v != nil {
			return *v
		}
		var ret ImageReferenceResponse
		return ret
	}).(ImageReferenceResponseOutput)
}

func (o ImageReferenceResponsePtrOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Offer
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponsePtrOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type InstanceViewStatus struct {
	Code          *string           `pulumi:"code"`
	DisplayStatus *string           `pulumi:"displayStatus"`
	Level         *StatusLevelTypes `pulumi:"level"`
	Message       *string           `pulumi:"message"`
	Time          *string           `pulumi:"time"`
}





type InstanceViewStatusInput interface {
	pulumi.Input

	ToInstanceViewStatusOutput() InstanceViewStatusOutput
	ToInstanceViewStatusOutputWithContext(context.Context) InstanceViewStatusOutput
}

type InstanceViewStatusArgs struct {
	Code          pulumi.StringPtrInput    `pulumi:"code"`
	DisplayStatus pulumi.StringPtrInput    `pulumi:"displayStatus"`
	Level         StatusLevelTypesPtrInput `pulumi:"level"`
	Message       pulumi.StringPtrInput    `pulumi:"message"`
	Time          pulumi.StringPtrInput    `pulumi:"time"`
}

func (InstanceViewStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceViewStatus)(nil)).Elem()
}

func (i InstanceViewStatusArgs) ToInstanceViewStatusOutput() InstanceViewStatusOutput {
	return i.ToInstanceViewStatusOutputWithContext(context.Background())
}

func (i InstanceViewStatusArgs) ToInstanceViewStatusOutputWithContext(ctx context.Context) InstanceViewStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceViewStatusOutput)
}





type InstanceViewStatusArrayInput interface {
	pulumi.Input

	ToInstanceViewStatusArrayOutput() InstanceViewStatusArrayOutput
	ToInstanceViewStatusArrayOutputWithContext(context.Context) InstanceViewStatusArrayOutput
}

type InstanceViewStatusArray []InstanceViewStatusInput

func (InstanceViewStatusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceViewStatus)(nil)).Elem()
}

func (i InstanceViewStatusArray) ToInstanceViewStatusArrayOutput() InstanceViewStatusArrayOutput {
	return i.ToInstanceViewStatusArrayOutputWithContext(context.Background())
}

func (i InstanceViewStatusArray) ToInstanceViewStatusArrayOutputWithContext(ctx context.Context) InstanceViewStatusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceViewStatusArrayOutput)
}

type InstanceViewStatusOutput struct{ *pulumi.OutputState }

func (InstanceViewStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceViewStatus)(nil)).Elem()
}

func (o InstanceViewStatusOutput) ToInstanceViewStatusOutput() InstanceViewStatusOutput {
	return o
}

func (o InstanceViewStatusOutput) ToInstanceViewStatusOutputWithContext(ctx context.Context) InstanceViewStatusOutput {
	return o
}

func (o InstanceViewStatusOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatus) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatus) *string { return v.DisplayStatus }).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusOutput) Level() StatusLevelTypesPtrOutput {
	return o.ApplyT(func(v InstanceViewStatus) *StatusLevelTypes { return v.Level }).(StatusLevelTypesPtrOutput)
}

func (o InstanceViewStatusOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatus) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatus) *string { return v.Time }).(pulumi.StringPtrOutput)
}

type InstanceViewStatusArrayOutput struct{ *pulumi.OutputState }

func (InstanceViewStatusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceViewStatus)(nil)).Elem()
}

func (o InstanceViewStatusArrayOutput) ToInstanceViewStatusArrayOutput() InstanceViewStatusArrayOutput {
	return o
}

func (o InstanceViewStatusArrayOutput) ToInstanceViewStatusArrayOutputWithContext(ctx context.Context) InstanceViewStatusArrayOutput {
	return o
}

func (o InstanceViewStatusArrayOutput) Index(i pulumi.IntInput) InstanceViewStatusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceViewStatus {
		return vs[0].([]InstanceViewStatus)[vs[1].(int)]
	}).(InstanceViewStatusOutput)
}

type InstanceViewStatusResponse struct {
	Code          *string `pulumi:"code"`
	DisplayStatus *string `pulumi:"displayStatus"`
	Level         *string `pulumi:"level"`
	Message       *string `pulumi:"message"`
	Time          *string `pulumi:"time"`
}

type InstanceViewStatusResponseOutput struct{ *pulumi.OutputState }

func (InstanceViewStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceViewStatusResponse)(nil)).Elem()
}

func (o InstanceViewStatusResponseOutput) ToInstanceViewStatusResponseOutput() InstanceViewStatusResponseOutput {
	return o
}

func (o InstanceViewStatusResponseOutput) ToInstanceViewStatusResponseOutputWithContext(ctx context.Context) InstanceViewStatusResponseOutput {
	return o
}

func (o InstanceViewStatusResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatusResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusResponseOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatusResponse) *string { return v.DisplayStatus }).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusResponseOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatusResponse) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatusResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusResponseOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatusResponse) *string { return v.Time }).(pulumi.StringPtrOutput)
}

type InstanceViewStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (InstanceViewStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceViewStatusResponse)(nil)).Elem()
}

func (o InstanceViewStatusResponsePtrOutput) ToInstanceViewStatusResponsePtrOutput() InstanceViewStatusResponsePtrOutput {
	return o
}

func (o InstanceViewStatusResponsePtrOutput) ToInstanceViewStatusResponsePtrOutputWithContext(ctx context.Context) InstanceViewStatusResponsePtrOutput {
	return o
}

func (o InstanceViewStatusResponsePtrOutput) Elem() InstanceViewStatusResponseOutput {
	return o.ApplyT(func(v *InstanceViewStatusResponse) InstanceViewStatusResponse {
		if v != nil {
			return *v
		}
		var ret InstanceViewStatusResponse
		return ret
	}).(InstanceViewStatusResponseOutput)
}

func (o InstanceViewStatusResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceViewStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusResponsePtrOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceViewStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.DisplayStatus
	}).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusResponsePtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceViewStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.Level
	}).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceViewStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusResponsePtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceViewStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.Time
	}).(pulumi.StringPtrOutput)
}

type InstanceViewStatusResponseArrayOutput struct{ *pulumi.OutputState }

func (InstanceViewStatusResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceViewStatusResponse)(nil)).Elem()
}

func (o InstanceViewStatusResponseArrayOutput) ToInstanceViewStatusResponseArrayOutput() InstanceViewStatusResponseArrayOutput {
	return o
}

func (o InstanceViewStatusResponseArrayOutput) ToInstanceViewStatusResponseArrayOutputWithContext(ctx context.Context) InstanceViewStatusResponseArrayOutput {
	return o
}

func (o InstanceViewStatusResponseArrayOutput) Index(i pulumi.IntInput) InstanceViewStatusResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceViewStatusResponse {
		return vs[0].([]InstanceViewStatusResponse)[vs[1].(int)]
	}).(InstanceViewStatusResponseOutput)
}

type KeyVaultKeyReference struct {
	KeyUrl      string      `pulumi:"keyUrl"`
	SourceVault SubResource `pulumi:"sourceVault"`
}





type KeyVaultKeyReferenceInput interface {
	pulumi.Input

	ToKeyVaultKeyReferenceOutput() KeyVaultKeyReferenceOutput
	ToKeyVaultKeyReferenceOutputWithContext(context.Context) KeyVaultKeyReferenceOutput
}

type KeyVaultKeyReferenceArgs struct {
	KeyUrl      pulumi.StringInput `pulumi:"keyUrl"`
	SourceVault SubResourceInput   `pulumi:"sourceVault"`
}

func (KeyVaultKeyReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReference)(nil)).Elem()
}

func (i KeyVaultKeyReferenceArgs) ToKeyVaultKeyReferenceOutput() KeyVaultKeyReferenceOutput {
	return i.ToKeyVaultKeyReferenceOutputWithContext(context.Background())
}

func (i KeyVaultKeyReferenceArgs) ToKeyVaultKeyReferenceOutputWithContext(ctx context.Context) KeyVaultKeyReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceOutput)
}

func (i KeyVaultKeyReferenceArgs) ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput {
	return i.ToKeyVaultKeyReferencePtrOutputWithContext(context.Background())
}

func (i KeyVaultKeyReferenceArgs) ToKeyVaultKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceOutput).ToKeyVaultKeyReferencePtrOutputWithContext(ctx)
}









type KeyVaultKeyReferencePtrInput interface {
	pulumi.Input

	ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput
	ToKeyVaultKeyReferencePtrOutputWithContext(context.Context) KeyVaultKeyReferencePtrOutput
}

type keyVaultKeyReferencePtrType KeyVaultKeyReferenceArgs

func KeyVaultKeyReferencePtr(v *KeyVaultKeyReferenceArgs) KeyVaultKeyReferencePtrInput {
	return (*keyVaultKeyReferencePtrType)(v)
}

func (*keyVaultKeyReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReference)(nil)).Elem()
}

func (i *keyVaultKeyReferencePtrType) ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput {
	return i.ToKeyVaultKeyReferencePtrOutputWithContext(context.Background())
}

func (i *keyVaultKeyReferencePtrType) ToKeyVaultKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferencePtrOutput)
}

type KeyVaultKeyReferenceOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReference)(nil)).Elem()
}

func (o KeyVaultKeyReferenceOutput) ToKeyVaultKeyReferenceOutput() KeyVaultKeyReferenceOutput {
	return o
}

func (o KeyVaultKeyReferenceOutput) ToKeyVaultKeyReferenceOutputWithContext(ctx context.Context) KeyVaultKeyReferenceOutput {
	return o
}

func (o KeyVaultKeyReferenceOutput) ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput {
	return o.ToKeyVaultKeyReferencePtrOutputWithContext(context.Background())
}

func (o KeyVaultKeyReferenceOutput) ToKeyVaultKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultKeyReference) *KeyVaultKeyReference {
		return &v
	}).(KeyVaultKeyReferencePtrOutput)
}

func (o KeyVaultKeyReferenceOutput) KeyUrl() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultKeyReference) string { return v.KeyUrl }).(pulumi.StringOutput)
}

func (o KeyVaultKeyReferenceOutput) SourceVault() SubResourceOutput {
	return o.ApplyT(func(v KeyVaultKeyReference) SubResource { return v.SourceVault }).(SubResourceOutput)
}

type KeyVaultKeyReferencePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReference)(nil)).Elem()
}

func (o KeyVaultKeyReferencePtrOutput) ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput {
	return o
}

func (o KeyVaultKeyReferencePtrOutput) ToKeyVaultKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferencePtrOutput {
	return o
}

func (o KeyVaultKeyReferencePtrOutput) Elem() KeyVaultKeyReferenceOutput {
	return o.ApplyT(func(v *KeyVaultKeyReference) KeyVaultKeyReference {
		if v != nil {
			return *v
		}
		var ret KeyVaultKeyReference
		return ret
	}).(KeyVaultKeyReferenceOutput)
}

func (o KeyVaultKeyReferencePtrOutput) KeyUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReference) *string {
		if v == nil {
			return nil
		}
		return &v.KeyUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultKeyReferencePtrOutput) SourceVault() SubResourcePtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReference) *SubResource {
		if v == nil {
			return nil
		}
		return &v.SourceVault
	}).(SubResourcePtrOutput)
}

type KeyVaultKeyReferenceResponse struct {
	KeyUrl      string              `pulumi:"keyUrl"`
	SourceVault SubResourceResponse `pulumi:"sourceVault"`
}

type KeyVaultKeyReferenceResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReferenceResponse)(nil)).Elem()
}

func (o KeyVaultKeyReferenceResponseOutput) ToKeyVaultKeyReferenceResponseOutput() KeyVaultKeyReferenceResponseOutput {
	return o
}

func (o KeyVaultKeyReferenceResponseOutput) ToKeyVaultKeyReferenceResponseOutputWithContext(ctx context.Context) KeyVaultKeyReferenceResponseOutput {
	return o
}

func (o KeyVaultKeyReferenceResponseOutput) KeyUrl() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultKeyReferenceResponse) string { return v.KeyUrl }).(pulumi.StringOutput)
}

func (o KeyVaultKeyReferenceResponseOutput) SourceVault() SubResourceResponseOutput {
	return o.ApplyT(func(v KeyVaultKeyReferenceResponse) SubResourceResponse { return v.SourceVault }).(SubResourceResponseOutput)
}

type KeyVaultKeyReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReferenceResponse)(nil)).Elem()
}

func (o KeyVaultKeyReferenceResponsePtrOutput) ToKeyVaultKeyReferenceResponsePtrOutput() KeyVaultKeyReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultKeyReferenceResponsePtrOutput) ToKeyVaultKeyReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultKeyReferenceResponsePtrOutput) Elem() KeyVaultKeyReferenceResponseOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponse) KeyVaultKeyReferenceResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultKeyReferenceResponse
		return ret
	}).(KeyVaultKeyReferenceResponseOutput)
}

func (o KeyVaultKeyReferenceResponsePtrOutput) KeyUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultKeyReferenceResponsePtrOutput) SourceVault() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponse) *SubResourceResponse {
		if v == nil {
			return nil
		}
		return &v.SourceVault
	}).(SubResourceResponsePtrOutput)
}

type KeyVaultSecretReference struct {
	SecretUrl   string      `pulumi:"secretUrl"`
	SourceVault SubResource `pulumi:"sourceVault"`
}





type KeyVaultSecretReferenceInput interface {
	pulumi.Input

	ToKeyVaultSecretReferenceOutput() KeyVaultSecretReferenceOutput
	ToKeyVaultSecretReferenceOutputWithContext(context.Context) KeyVaultSecretReferenceOutput
}

type KeyVaultSecretReferenceArgs struct {
	SecretUrl   pulumi.StringInput `pulumi:"secretUrl"`
	SourceVault SubResourceInput   `pulumi:"sourceVault"`
}

func (KeyVaultSecretReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultSecretReference)(nil)).Elem()
}

func (i KeyVaultSecretReferenceArgs) ToKeyVaultSecretReferenceOutput() KeyVaultSecretReferenceOutput {
	return i.ToKeyVaultSecretReferenceOutputWithContext(context.Background())
}

func (i KeyVaultSecretReferenceArgs) ToKeyVaultSecretReferenceOutputWithContext(ctx context.Context) KeyVaultSecretReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultSecretReferenceOutput)
}

func (i KeyVaultSecretReferenceArgs) ToKeyVaultSecretReferencePtrOutput() KeyVaultSecretReferencePtrOutput {
	return i.ToKeyVaultSecretReferencePtrOutputWithContext(context.Background())
}

func (i KeyVaultSecretReferenceArgs) ToKeyVaultSecretReferencePtrOutputWithContext(ctx context.Context) KeyVaultSecretReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultSecretReferenceOutput).ToKeyVaultSecretReferencePtrOutputWithContext(ctx)
}









type KeyVaultSecretReferencePtrInput interface {
	pulumi.Input

	ToKeyVaultSecretReferencePtrOutput() KeyVaultSecretReferencePtrOutput
	ToKeyVaultSecretReferencePtrOutputWithContext(context.Context) KeyVaultSecretReferencePtrOutput
}

type keyVaultSecretReferencePtrType KeyVaultSecretReferenceArgs

func KeyVaultSecretReferencePtr(v *KeyVaultSecretReferenceArgs) KeyVaultSecretReferencePtrInput {
	return (*keyVaultSecretReferencePtrType)(v)
}

func (*keyVaultSecretReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultSecretReference)(nil)).Elem()
}

func (i *keyVaultSecretReferencePtrType) ToKeyVaultSecretReferencePtrOutput() KeyVaultSecretReferencePtrOutput {
	return i.ToKeyVaultSecretReferencePtrOutputWithContext(context.Background())
}

func (i *keyVaultSecretReferencePtrType) ToKeyVaultSecretReferencePtrOutputWithContext(ctx context.Context) KeyVaultSecretReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultSecretReferencePtrOutput)
}

type KeyVaultSecretReferenceOutput struct{ *pulumi.OutputState }

func (KeyVaultSecretReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultSecretReference)(nil)).Elem()
}

func (o KeyVaultSecretReferenceOutput) ToKeyVaultSecretReferenceOutput() KeyVaultSecretReferenceOutput {
	return o
}

func (o KeyVaultSecretReferenceOutput) ToKeyVaultSecretReferenceOutputWithContext(ctx context.Context) KeyVaultSecretReferenceOutput {
	return o
}

func (o KeyVaultSecretReferenceOutput) ToKeyVaultSecretReferencePtrOutput() KeyVaultSecretReferencePtrOutput {
	return o.ToKeyVaultSecretReferencePtrOutputWithContext(context.Background())
}

func (o KeyVaultSecretReferenceOutput) ToKeyVaultSecretReferencePtrOutputWithContext(ctx context.Context) KeyVaultSecretReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultSecretReference) *KeyVaultSecretReference {
		return &v
	}).(KeyVaultSecretReferencePtrOutput)
}

func (o KeyVaultSecretReferenceOutput) SecretUrl() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultSecretReference) string { return v.SecretUrl }).(pulumi.StringOutput)
}

func (o KeyVaultSecretReferenceOutput) SourceVault() SubResourceOutput {
	return o.ApplyT(func(v KeyVaultSecretReference) SubResource { return v.SourceVault }).(SubResourceOutput)
}

type KeyVaultSecretReferencePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultSecretReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultSecretReference)(nil)).Elem()
}

func (o KeyVaultSecretReferencePtrOutput) ToKeyVaultSecretReferencePtrOutput() KeyVaultSecretReferencePtrOutput {
	return o
}

func (o KeyVaultSecretReferencePtrOutput) ToKeyVaultSecretReferencePtrOutputWithContext(ctx context.Context) KeyVaultSecretReferencePtrOutput {
	return o
}

func (o KeyVaultSecretReferencePtrOutput) Elem() KeyVaultSecretReferenceOutput {
	return o.ApplyT(func(v *KeyVaultSecretReference) KeyVaultSecretReference {
		if v != nil {
			return *v
		}
		var ret KeyVaultSecretReference
		return ret
	}).(KeyVaultSecretReferenceOutput)
}

func (o KeyVaultSecretReferencePtrOutput) SecretUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultSecretReference) *string {
		if v == nil {
			return nil
		}
		return &v.SecretUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultSecretReferencePtrOutput) SourceVault() SubResourcePtrOutput {
	return o.ApplyT(func(v *KeyVaultSecretReference) *SubResource {
		if v == nil {
			return nil
		}
		return &v.SourceVault
	}).(SubResourcePtrOutput)
}

type KeyVaultSecretReferenceResponse struct {
	SecretUrl   string              `pulumi:"secretUrl"`
	SourceVault SubResourceResponse `pulumi:"sourceVault"`
}

type KeyVaultSecretReferenceResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultSecretReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultSecretReferenceResponse)(nil)).Elem()
}

func (o KeyVaultSecretReferenceResponseOutput) ToKeyVaultSecretReferenceResponseOutput() KeyVaultSecretReferenceResponseOutput {
	return o
}

func (o KeyVaultSecretReferenceResponseOutput) ToKeyVaultSecretReferenceResponseOutputWithContext(ctx context.Context) KeyVaultSecretReferenceResponseOutput {
	return o
}

func (o KeyVaultSecretReferenceResponseOutput) SecretUrl() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultSecretReferenceResponse) string { return v.SecretUrl }).(pulumi.StringOutput)
}

func (o KeyVaultSecretReferenceResponseOutput) SourceVault() SubResourceResponseOutput {
	return o.ApplyT(func(v KeyVaultSecretReferenceResponse) SubResourceResponse { return v.SourceVault }).(SubResourceResponseOutput)
}

type KeyVaultSecretReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultSecretReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultSecretReferenceResponse)(nil)).Elem()
}

func (o KeyVaultSecretReferenceResponsePtrOutput) ToKeyVaultSecretReferenceResponsePtrOutput() KeyVaultSecretReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultSecretReferenceResponsePtrOutput) ToKeyVaultSecretReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultSecretReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultSecretReferenceResponsePtrOutput) Elem() KeyVaultSecretReferenceResponseOutput {
	return o.ApplyT(func(v *KeyVaultSecretReferenceResponse) KeyVaultSecretReferenceResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultSecretReferenceResponse
		return ret
	}).(KeyVaultSecretReferenceResponseOutput)
}

func (o KeyVaultSecretReferenceResponsePtrOutput) SecretUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultSecretReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SecretUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultSecretReferenceResponsePtrOutput) SourceVault() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *KeyVaultSecretReferenceResponse) *SubResourceResponse {
		if v == nil {
			return nil
		}
		return &v.SourceVault
	}).(SubResourceResponsePtrOutput)
}

type LinuxConfiguration struct {
	DisablePasswordAuthentication *bool             `pulumi:"disablePasswordAuthentication"`
	Ssh                           *SshConfiguration `pulumi:"ssh"`
}





type LinuxConfigurationInput interface {
	pulumi.Input

	ToLinuxConfigurationOutput() LinuxConfigurationOutput
	ToLinuxConfigurationOutputWithContext(context.Context) LinuxConfigurationOutput
}

type LinuxConfigurationArgs struct {
	DisablePasswordAuthentication pulumi.BoolPtrInput      `pulumi:"disablePasswordAuthentication"`
	Ssh                           SshConfigurationPtrInput `pulumi:"ssh"`
}

func (LinuxConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxConfiguration)(nil)).Elem()
}

func (i LinuxConfigurationArgs) ToLinuxConfigurationOutput() LinuxConfigurationOutput {
	return i.ToLinuxConfigurationOutputWithContext(context.Background())
}

func (i LinuxConfigurationArgs) ToLinuxConfigurationOutputWithContext(ctx context.Context) LinuxConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxConfigurationOutput)
}

func (i LinuxConfigurationArgs) ToLinuxConfigurationPtrOutput() LinuxConfigurationPtrOutput {
	return i.ToLinuxConfigurationPtrOutputWithContext(context.Background())
}

func (i LinuxConfigurationArgs) ToLinuxConfigurationPtrOutputWithContext(ctx context.Context) LinuxConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxConfigurationOutput).ToLinuxConfigurationPtrOutputWithContext(ctx)
}









type LinuxConfigurationPtrInput interface {
	pulumi.Input

	ToLinuxConfigurationPtrOutput() LinuxConfigurationPtrOutput
	ToLinuxConfigurationPtrOutputWithContext(context.Context) LinuxConfigurationPtrOutput
}

type linuxConfigurationPtrType LinuxConfigurationArgs

func LinuxConfigurationPtr(v *LinuxConfigurationArgs) LinuxConfigurationPtrInput {
	return (*linuxConfigurationPtrType)(v)
}

func (*linuxConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxConfiguration)(nil)).Elem()
}

func (i *linuxConfigurationPtrType) ToLinuxConfigurationPtrOutput() LinuxConfigurationPtrOutput {
	return i.ToLinuxConfigurationPtrOutputWithContext(context.Background())
}

func (i *linuxConfigurationPtrType) ToLinuxConfigurationPtrOutputWithContext(ctx context.Context) LinuxConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxConfigurationPtrOutput)
}

type LinuxConfigurationOutput struct{ *pulumi.OutputState }

func (LinuxConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxConfiguration)(nil)).Elem()
}

func (o LinuxConfigurationOutput) ToLinuxConfigurationOutput() LinuxConfigurationOutput {
	return o
}

func (o LinuxConfigurationOutput) ToLinuxConfigurationOutputWithContext(ctx context.Context) LinuxConfigurationOutput {
	return o
}

func (o LinuxConfigurationOutput) ToLinuxConfigurationPtrOutput() LinuxConfigurationPtrOutput {
	return o.ToLinuxConfigurationPtrOutputWithContext(context.Background())
}

func (o LinuxConfigurationOutput) ToLinuxConfigurationPtrOutputWithContext(ctx context.Context) LinuxConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinuxConfiguration) *LinuxConfiguration {
		return &v
	}).(LinuxConfigurationPtrOutput)
}

func (o LinuxConfigurationOutput) DisablePasswordAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LinuxConfiguration) *bool { return v.DisablePasswordAuthentication }).(pulumi.BoolPtrOutput)
}

func (o LinuxConfigurationOutput) Ssh() SshConfigurationPtrOutput {
	return o.ApplyT(func(v LinuxConfiguration) *SshConfiguration { return v.Ssh }).(SshConfigurationPtrOutput)
}

type LinuxConfigurationPtrOutput struct{ *pulumi.OutputState }

func (LinuxConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxConfiguration)(nil)).Elem()
}

func (o LinuxConfigurationPtrOutput) ToLinuxConfigurationPtrOutput() LinuxConfigurationPtrOutput {
	return o
}

func (o LinuxConfigurationPtrOutput) ToLinuxConfigurationPtrOutputWithContext(ctx context.Context) LinuxConfigurationPtrOutput {
	return o
}

func (o LinuxConfigurationPtrOutput) Elem() LinuxConfigurationOutput {
	return o.ApplyT(func(v *LinuxConfiguration) LinuxConfiguration {
		if v != nil {
			return *v
		}
		var ret LinuxConfiguration
		return ret
	}).(LinuxConfigurationOutput)
}

func (o LinuxConfigurationPtrOutput) DisablePasswordAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LinuxConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.DisablePasswordAuthentication
	}).(pulumi.BoolPtrOutput)
}

func (o LinuxConfigurationPtrOutput) Ssh() SshConfigurationPtrOutput {
	return o.ApplyT(func(v *LinuxConfiguration) *SshConfiguration {
		if v == nil {
			return nil
		}
		return v.Ssh
	}).(SshConfigurationPtrOutput)
}

type LinuxConfigurationResponse struct {
	DisablePasswordAuthentication *bool                     `pulumi:"disablePasswordAuthentication"`
	Ssh                           *SshConfigurationResponse `pulumi:"ssh"`
}

type LinuxConfigurationResponseOutput struct{ *pulumi.OutputState }

func (LinuxConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxConfigurationResponse)(nil)).Elem()
}

func (o LinuxConfigurationResponseOutput) ToLinuxConfigurationResponseOutput() LinuxConfigurationResponseOutput {
	return o
}

func (o LinuxConfigurationResponseOutput) ToLinuxConfigurationResponseOutputWithContext(ctx context.Context) LinuxConfigurationResponseOutput {
	return o
}

func (o LinuxConfigurationResponseOutput) DisablePasswordAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LinuxConfigurationResponse) *bool { return v.DisablePasswordAuthentication }).(pulumi.BoolPtrOutput)
}

func (o LinuxConfigurationResponseOutput) Ssh() SshConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LinuxConfigurationResponse) *SshConfigurationResponse { return v.Ssh }).(SshConfigurationResponsePtrOutput)
}

type LinuxConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (LinuxConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxConfigurationResponse)(nil)).Elem()
}

func (o LinuxConfigurationResponsePtrOutput) ToLinuxConfigurationResponsePtrOutput() LinuxConfigurationResponsePtrOutput {
	return o
}

func (o LinuxConfigurationResponsePtrOutput) ToLinuxConfigurationResponsePtrOutputWithContext(ctx context.Context) LinuxConfigurationResponsePtrOutput {
	return o
}

func (o LinuxConfigurationResponsePtrOutput) Elem() LinuxConfigurationResponseOutput {
	return o.ApplyT(func(v *LinuxConfigurationResponse) LinuxConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret LinuxConfigurationResponse
		return ret
	}).(LinuxConfigurationResponseOutput)
}

func (o LinuxConfigurationResponsePtrOutput) DisablePasswordAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LinuxConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.DisablePasswordAuthentication
	}).(pulumi.BoolPtrOutput)
}

func (o LinuxConfigurationResponsePtrOutput) Ssh() SshConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *LinuxConfigurationResponse) *SshConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.Ssh
	}).(SshConfigurationResponsePtrOutput)
}

type NetworkInterfaceReference struct {
	Id      *string `pulumi:"id"`
	Primary *bool   `pulumi:"primary"`
}





type NetworkInterfaceReferenceInput interface {
	pulumi.Input

	ToNetworkInterfaceReferenceOutput() NetworkInterfaceReferenceOutput
	ToNetworkInterfaceReferenceOutputWithContext(context.Context) NetworkInterfaceReferenceOutput
}

type NetworkInterfaceReferenceArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Primary pulumi.BoolPtrInput   `pulumi:"primary"`
}

func (NetworkInterfaceReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceReference)(nil)).Elem()
}

func (i NetworkInterfaceReferenceArgs) ToNetworkInterfaceReferenceOutput() NetworkInterfaceReferenceOutput {
	return i.ToNetworkInterfaceReferenceOutputWithContext(context.Background())
}

func (i NetworkInterfaceReferenceArgs) ToNetworkInterfaceReferenceOutputWithContext(ctx context.Context) NetworkInterfaceReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceReferenceOutput)
}





type NetworkInterfaceReferenceArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceReferenceArrayOutput() NetworkInterfaceReferenceArrayOutput
	ToNetworkInterfaceReferenceArrayOutputWithContext(context.Context) NetworkInterfaceReferenceArrayOutput
}

type NetworkInterfaceReferenceArray []NetworkInterfaceReferenceInput

func (NetworkInterfaceReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceReference)(nil)).Elem()
}

func (i NetworkInterfaceReferenceArray) ToNetworkInterfaceReferenceArrayOutput() NetworkInterfaceReferenceArrayOutput {
	return i.ToNetworkInterfaceReferenceArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceReferenceArray) ToNetworkInterfaceReferenceArrayOutputWithContext(ctx context.Context) NetworkInterfaceReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceReferenceArrayOutput)
}

type NetworkInterfaceReferenceOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceReference)(nil)).Elem()
}

func (o NetworkInterfaceReferenceOutput) ToNetworkInterfaceReferenceOutput() NetworkInterfaceReferenceOutput {
	return o
}

func (o NetworkInterfaceReferenceOutput) ToNetworkInterfaceReferenceOutputWithContext(ctx context.Context) NetworkInterfaceReferenceOutput {
	return o
}

func (o NetworkInterfaceReferenceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceReference) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceReferenceOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceReference) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

type NetworkInterfaceReferenceArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceReference)(nil)).Elem()
}

func (o NetworkInterfaceReferenceArrayOutput) ToNetworkInterfaceReferenceArrayOutput() NetworkInterfaceReferenceArrayOutput {
	return o
}

func (o NetworkInterfaceReferenceArrayOutput) ToNetworkInterfaceReferenceArrayOutputWithContext(ctx context.Context) NetworkInterfaceReferenceArrayOutput {
	return o
}

func (o NetworkInterfaceReferenceArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfaceReference {
		return vs[0].([]NetworkInterfaceReference)[vs[1].(int)]
	}).(NetworkInterfaceReferenceOutput)
}

type NetworkInterfaceReferenceResponse struct {
	Id      *string `pulumi:"id"`
	Primary *bool   `pulumi:"primary"`
}

type NetworkInterfaceReferenceResponseOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceReferenceResponse)(nil)).Elem()
}

func (o NetworkInterfaceReferenceResponseOutput) ToNetworkInterfaceReferenceResponseOutput() NetworkInterfaceReferenceResponseOutput {
	return o
}

func (o NetworkInterfaceReferenceResponseOutput) ToNetworkInterfaceReferenceResponseOutputWithContext(ctx context.Context) NetworkInterfaceReferenceResponseOutput {
	return o
}

func (o NetworkInterfaceReferenceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceReferenceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceReferenceResponseOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceReferenceResponse) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

type NetworkInterfaceReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceReferenceResponse)(nil)).Elem()
}

func (o NetworkInterfaceReferenceResponseArrayOutput) ToNetworkInterfaceReferenceResponseArrayOutput() NetworkInterfaceReferenceResponseArrayOutput {
	return o
}

func (o NetworkInterfaceReferenceResponseArrayOutput) ToNetworkInterfaceReferenceResponseArrayOutputWithContext(ctx context.Context) NetworkInterfaceReferenceResponseArrayOutput {
	return o
}

func (o NetworkInterfaceReferenceResponseArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfaceReferenceResponse {
		return vs[0].([]NetworkInterfaceReferenceResponse)[vs[1].(int)]
	}).(NetworkInterfaceReferenceResponseOutput)
}

type NetworkProfile struct {
	NetworkInterfaces []NetworkInterfaceReference `pulumi:"networkInterfaces"`
}





type NetworkProfileInput interface {
	pulumi.Input

	ToNetworkProfileOutput() NetworkProfileOutput
	ToNetworkProfileOutputWithContext(context.Context) NetworkProfileOutput
}

type NetworkProfileArgs struct {
	NetworkInterfaces NetworkInterfaceReferenceArrayInput `pulumi:"networkInterfaces"`
}

func (NetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil)).Elem()
}

func (i NetworkProfileArgs) ToNetworkProfileOutput() NetworkProfileOutput {
	return i.ToNetworkProfileOutputWithContext(context.Background())
}

func (i NetworkProfileArgs) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput)
}

func (i NetworkProfileArgs) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return i.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (i NetworkProfileArgs) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput).ToNetworkProfilePtrOutputWithContext(ctx)
}









type NetworkProfilePtrInput interface {
	pulumi.Input

	ToNetworkProfilePtrOutput() NetworkProfilePtrOutput
	ToNetworkProfilePtrOutputWithContext(context.Context) NetworkProfilePtrOutput
}

type networkProfilePtrType NetworkProfileArgs

func NetworkProfilePtr(v *NetworkProfileArgs) NetworkProfilePtrInput {
	return (*networkProfilePtrType)(v)
}

func (*networkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (i *networkProfilePtrType) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return i.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *networkProfilePtrType) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfilePtrOutput)
}

type NetworkProfileOutput struct{ *pulumi.OutputState }

func (NetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil)).Elem()
}

func (o NetworkProfileOutput) ToNetworkProfileOutput() NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return o.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (o NetworkProfileOutput) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkProfile) *NetworkProfile {
		return &v
	}).(NetworkProfilePtrOutput)
}

func (o NetworkProfileOutput) NetworkInterfaces() NetworkInterfaceReferenceArrayOutput {
	return o.ApplyT(func(v NetworkProfile) []NetworkInterfaceReference { return v.NetworkInterfaces }).(NetworkInterfaceReferenceArrayOutput)
}

type NetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (NetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (o NetworkProfilePtrOutput) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return o
}

func (o NetworkProfilePtrOutput) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return o
}

func (o NetworkProfilePtrOutput) Elem() NetworkProfileOutput {
	return o.ApplyT(func(v *NetworkProfile) NetworkProfile {
		if v != nil {
			return *v
		}
		var ret NetworkProfile
		return ret
	}).(NetworkProfileOutput)
}

func (o NetworkProfilePtrOutput) NetworkInterfaces() NetworkInterfaceReferenceArrayOutput {
	return o.ApplyT(func(v *NetworkProfile) []NetworkInterfaceReference {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaces
	}).(NetworkInterfaceReferenceArrayOutput)
}

type NetworkProfileResponse struct {
	NetworkInterfaces []NetworkInterfaceReferenceResponse `pulumi:"networkInterfaces"`
}

type NetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfileResponse)(nil)).Elem()
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponseOutput() NetworkProfileResponseOutput {
	return o
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponseOutputWithContext(ctx context.Context) NetworkProfileResponseOutput {
	return o
}

func (o NetworkProfileResponseOutput) NetworkInterfaces() NetworkInterfaceReferenceResponseArrayOutput {
	return o.ApplyT(func(v NetworkProfileResponse) []NetworkInterfaceReferenceResponse { return v.NetworkInterfaces }).(NetworkInterfaceReferenceResponseArrayOutput)
}

type NetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfileResponse)(nil)).Elem()
}

func (o NetworkProfileResponsePtrOutput) ToNetworkProfileResponsePtrOutput() NetworkProfileResponsePtrOutput {
	return o
}

func (o NetworkProfileResponsePtrOutput) ToNetworkProfileResponsePtrOutputWithContext(ctx context.Context) NetworkProfileResponsePtrOutput {
	return o
}

func (o NetworkProfileResponsePtrOutput) Elem() NetworkProfileResponseOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) NetworkProfileResponse {
		if v != nil {
			return *v
		}
		var ret NetworkProfileResponse
		return ret
	}).(NetworkProfileResponseOutput)
}

func (o NetworkProfileResponsePtrOutput) NetworkInterfaces() NetworkInterfaceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) []NetworkInterfaceReferenceResponse {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaces
	}).(NetworkInterfaceReferenceResponseArrayOutput)
}

type OSDisk struct {
	Caching            *CachingTypes           `pulumi:"caching"`
	CreateOption       DiskCreateOptionTypes   `pulumi:"createOption"`
	DiskSizeGB         *int                    `pulumi:"diskSizeGB"`
	EncryptionSettings *DiskEncryptionSettings `pulumi:"encryptionSettings"`
	Image              *VirtualHardDisk        `pulumi:"image"`
	Name               string                  `pulumi:"name"`
	OsType             *OperatingSystemTypes   `pulumi:"osType"`
	Vhd                VirtualHardDisk         `pulumi:"vhd"`
}





type OSDiskInput interface {
	pulumi.Input

	ToOSDiskOutput() OSDiskOutput
	ToOSDiskOutputWithContext(context.Context) OSDiskOutput
}

type OSDiskArgs struct {
	Caching            CachingTypesPtrInput           `pulumi:"caching"`
	CreateOption       DiskCreateOptionTypesInput     `pulumi:"createOption"`
	DiskSizeGB         pulumi.IntPtrInput             `pulumi:"diskSizeGB"`
	EncryptionSettings DiskEncryptionSettingsPtrInput `pulumi:"encryptionSettings"`
	Image              VirtualHardDiskPtrInput        `pulumi:"image"`
	Name               pulumi.StringInput             `pulumi:"name"`
	OsType             OperatingSystemTypesPtrInput   `pulumi:"osType"`
	Vhd                VirtualHardDiskInput           `pulumi:"vhd"`
}

func (OSDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OSDisk)(nil)).Elem()
}

func (i OSDiskArgs) ToOSDiskOutput() OSDiskOutput {
	return i.ToOSDiskOutputWithContext(context.Background())
}

func (i OSDiskArgs) ToOSDiskOutputWithContext(ctx context.Context) OSDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSDiskOutput)
}

func (i OSDiskArgs) ToOSDiskPtrOutput() OSDiskPtrOutput {
	return i.ToOSDiskPtrOutputWithContext(context.Background())
}

func (i OSDiskArgs) ToOSDiskPtrOutputWithContext(ctx context.Context) OSDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSDiskOutput).ToOSDiskPtrOutputWithContext(ctx)
}









type OSDiskPtrInput interface {
	pulumi.Input

	ToOSDiskPtrOutput() OSDiskPtrOutput
	ToOSDiskPtrOutputWithContext(context.Context) OSDiskPtrOutput
}

type osdiskPtrType OSDiskArgs

func OSDiskPtr(v *OSDiskArgs) OSDiskPtrInput {
	return (*osdiskPtrType)(v)
}

func (*osdiskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OSDisk)(nil)).Elem()
}

func (i *osdiskPtrType) ToOSDiskPtrOutput() OSDiskPtrOutput {
	return i.ToOSDiskPtrOutputWithContext(context.Background())
}

func (i *osdiskPtrType) ToOSDiskPtrOutputWithContext(ctx context.Context) OSDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSDiskPtrOutput)
}

type OSDiskOutput struct{ *pulumi.OutputState }

func (OSDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSDisk)(nil)).Elem()
}

func (o OSDiskOutput) ToOSDiskOutput() OSDiskOutput {
	return o
}

func (o OSDiskOutput) ToOSDiskOutputWithContext(ctx context.Context) OSDiskOutput {
	return o
}

func (o OSDiskOutput) ToOSDiskPtrOutput() OSDiskPtrOutput {
	return o.ToOSDiskPtrOutputWithContext(context.Background())
}

func (o OSDiskOutput) ToOSDiskPtrOutputWithContext(ctx context.Context) OSDiskPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OSDisk) *OSDisk {
		return &v
	}).(OSDiskPtrOutput)
}

func (o OSDiskOutput) Caching() CachingTypesPtrOutput {
	return o.ApplyT(func(v OSDisk) *CachingTypes { return v.Caching }).(CachingTypesPtrOutput)
}

func (o OSDiskOutput) CreateOption() DiskCreateOptionTypesOutput {
	return o.ApplyT(func(v OSDisk) DiskCreateOptionTypes { return v.CreateOption }).(DiskCreateOptionTypesOutput)
}

func (o OSDiskOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v OSDisk) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o OSDiskOutput) EncryptionSettings() DiskEncryptionSettingsPtrOutput {
	return o.ApplyT(func(v OSDisk) *DiskEncryptionSettings { return v.EncryptionSettings }).(DiskEncryptionSettingsPtrOutput)
}

func (o OSDiskOutput) Image() VirtualHardDiskPtrOutput {
	return o.ApplyT(func(v OSDisk) *VirtualHardDisk { return v.Image }).(VirtualHardDiskPtrOutput)
}

func (o OSDiskOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v OSDisk) string { return v.Name }).(pulumi.StringOutput)
}

func (o OSDiskOutput) OsType() OperatingSystemTypesPtrOutput {
	return o.ApplyT(func(v OSDisk) *OperatingSystemTypes { return v.OsType }).(OperatingSystemTypesPtrOutput)
}

func (o OSDiskOutput) Vhd() VirtualHardDiskOutput {
	return o.ApplyT(func(v OSDisk) VirtualHardDisk { return v.Vhd }).(VirtualHardDiskOutput)
}

type OSDiskPtrOutput struct{ *pulumi.OutputState }

func (OSDiskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSDisk)(nil)).Elem()
}

func (o OSDiskPtrOutput) ToOSDiskPtrOutput() OSDiskPtrOutput {
	return o
}

func (o OSDiskPtrOutput) ToOSDiskPtrOutputWithContext(ctx context.Context) OSDiskPtrOutput {
	return o
}

func (o OSDiskPtrOutput) Elem() OSDiskOutput {
	return o.ApplyT(func(v *OSDisk) OSDisk {
		if v != nil {
			return *v
		}
		var ret OSDisk
		return ret
	}).(OSDiskOutput)
}

func (o OSDiskPtrOutput) Caching() CachingTypesPtrOutput {
	return o.ApplyT(func(v *OSDisk) *CachingTypes {
		if v == nil {
			return nil
		}
		return v.Caching
	}).(CachingTypesPtrOutput)
}

func (o OSDiskPtrOutput) CreateOption() DiskCreateOptionTypesPtrOutput {
	return o.ApplyT(func(v *OSDisk) *DiskCreateOptionTypes {
		if v == nil {
			return nil
		}
		return &v.CreateOption
	}).(DiskCreateOptionTypesPtrOutput)
}

func (o OSDiskPtrOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OSDisk) *int {
		if v == nil {
			return nil
		}
		return v.DiskSizeGB
	}).(pulumi.IntPtrOutput)
}

func (o OSDiskPtrOutput) EncryptionSettings() DiskEncryptionSettingsPtrOutput {
	return o.ApplyT(func(v *OSDisk) *DiskEncryptionSettings {
		if v == nil {
			return nil
		}
		return v.EncryptionSettings
	}).(DiskEncryptionSettingsPtrOutput)
}

func (o OSDiskPtrOutput) Image() VirtualHardDiskPtrOutput {
	return o.ApplyT(func(v *OSDisk) *VirtualHardDisk {
		if v == nil {
			return nil
		}
		return v.Image
	}).(VirtualHardDiskPtrOutput)
}

func (o OSDiskPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDisk) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskPtrOutput) OsType() OperatingSystemTypesPtrOutput {
	return o.ApplyT(func(v *OSDisk) *OperatingSystemTypes {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(OperatingSystemTypesPtrOutput)
}

func (o OSDiskPtrOutput) Vhd() VirtualHardDiskPtrOutput {
	return o.ApplyT(func(v *OSDisk) *VirtualHardDisk {
		if v == nil {
			return nil
		}
		return &v.Vhd
	}).(VirtualHardDiskPtrOutput)
}

type OSDiskResponse struct {
	Caching            *string                         `pulumi:"caching"`
	CreateOption       string                          `pulumi:"createOption"`
	DiskSizeGB         *int                            `pulumi:"diskSizeGB"`
	EncryptionSettings *DiskEncryptionSettingsResponse `pulumi:"encryptionSettings"`
	Image              *VirtualHardDiskResponse        `pulumi:"image"`
	Name               string                          `pulumi:"name"`
	OsType             *string                         `pulumi:"osType"`
	Vhd                VirtualHardDiskResponse         `pulumi:"vhd"`
}

type OSDiskResponseOutput struct{ *pulumi.OutputState }

func (OSDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSDiskResponse)(nil)).Elem()
}

func (o OSDiskResponseOutput) ToOSDiskResponseOutput() OSDiskResponseOutput {
	return o
}

func (o OSDiskResponseOutput) ToOSDiskResponseOutputWithContext(ctx context.Context) OSDiskResponseOutput {
	return o
}

func (o OSDiskResponseOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *string { return v.Caching }).(pulumi.StringPtrOutput)
}

func (o OSDiskResponseOutput) CreateOption() pulumi.StringOutput {
	return o.ApplyT(func(v OSDiskResponse) string { return v.CreateOption }).(pulumi.StringOutput)
}

func (o OSDiskResponseOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o OSDiskResponseOutput) EncryptionSettings() DiskEncryptionSettingsResponsePtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *DiskEncryptionSettingsResponse { return v.EncryptionSettings }).(DiskEncryptionSettingsResponsePtrOutput)
}

func (o OSDiskResponseOutput) Image() VirtualHardDiskResponsePtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *VirtualHardDiskResponse { return v.Image }).(VirtualHardDiskResponsePtrOutput)
}

func (o OSDiskResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v OSDiskResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o OSDiskResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o OSDiskResponseOutput) Vhd() VirtualHardDiskResponseOutput {
	return o.ApplyT(func(v OSDiskResponse) VirtualHardDiskResponse { return v.Vhd }).(VirtualHardDiskResponseOutput)
}

type OSDiskResponsePtrOutput struct{ *pulumi.OutputState }

func (OSDiskResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSDiskResponse)(nil)).Elem()
}

func (o OSDiskResponsePtrOutput) ToOSDiskResponsePtrOutput() OSDiskResponsePtrOutput {
	return o
}

func (o OSDiskResponsePtrOutput) ToOSDiskResponsePtrOutputWithContext(ctx context.Context) OSDiskResponsePtrOutput {
	return o
}

func (o OSDiskResponsePtrOutput) Elem() OSDiskResponseOutput {
	return o.ApplyT(func(v *OSDiskResponse) OSDiskResponse {
		if v != nil {
			return *v
		}
		var ret OSDiskResponse
		return ret
	}).(OSDiskResponseOutput)
}

func (o OSDiskResponsePtrOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.Caching
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskResponsePtrOutput) CreateOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreateOption
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskResponsePtrOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *int {
		if v == nil {
			return nil
		}
		return v.DiskSizeGB
	}).(pulumi.IntPtrOutput)
}

func (o OSDiskResponsePtrOutput) EncryptionSettings() DiskEncryptionSettingsResponsePtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *DiskEncryptionSettingsResponse {
		if v == nil {
			return nil
		}
		return v.EncryptionSettings
	}).(DiskEncryptionSettingsResponsePtrOutput)
}

func (o OSDiskResponsePtrOutput) Image() VirtualHardDiskResponsePtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *VirtualHardDiskResponse {
		if v == nil {
			return nil
		}
		return v.Image
	}).(VirtualHardDiskResponsePtrOutput)
}

func (o OSDiskResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskResponsePtrOutput) Vhd() VirtualHardDiskResponsePtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *VirtualHardDiskResponse {
		if v == nil {
			return nil
		}
		return &v.Vhd
	}).(VirtualHardDiskResponsePtrOutput)
}

type OSProfile struct {
	AdminPassword        *string               `pulumi:"adminPassword"`
	AdminUsername        *string               `pulumi:"adminUsername"`
	ComputerName         *string               `pulumi:"computerName"`
	CustomData           *string               `pulumi:"customData"`
	LinuxConfiguration   *LinuxConfiguration   `pulumi:"linuxConfiguration"`
	Secrets              []VaultSecretGroup    `pulumi:"secrets"`
	WindowsConfiguration *WindowsConfiguration `pulumi:"windowsConfiguration"`
}





type OSProfileInput interface {
	pulumi.Input

	ToOSProfileOutput() OSProfileOutput
	ToOSProfileOutputWithContext(context.Context) OSProfileOutput
}

type OSProfileArgs struct {
	AdminPassword        pulumi.StringPtrInput        `pulumi:"adminPassword"`
	AdminUsername        pulumi.StringPtrInput        `pulumi:"adminUsername"`
	ComputerName         pulumi.StringPtrInput        `pulumi:"computerName"`
	CustomData           pulumi.StringPtrInput        `pulumi:"customData"`
	LinuxConfiguration   LinuxConfigurationPtrInput   `pulumi:"linuxConfiguration"`
	Secrets              VaultSecretGroupArrayInput   `pulumi:"secrets"`
	WindowsConfiguration WindowsConfigurationPtrInput `pulumi:"windowsConfiguration"`
}

func (OSProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfile)(nil)).Elem()
}

func (i OSProfileArgs) ToOSProfileOutput() OSProfileOutput {
	return i.ToOSProfileOutputWithContext(context.Background())
}

func (i OSProfileArgs) ToOSProfileOutputWithContext(ctx context.Context) OSProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfileOutput)
}

func (i OSProfileArgs) ToOSProfilePtrOutput() OSProfilePtrOutput {
	return i.ToOSProfilePtrOutputWithContext(context.Background())
}

func (i OSProfileArgs) ToOSProfilePtrOutputWithContext(ctx context.Context) OSProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfileOutput).ToOSProfilePtrOutputWithContext(ctx)
}









type OSProfilePtrInput interface {
	pulumi.Input

	ToOSProfilePtrOutput() OSProfilePtrOutput
	ToOSProfilePtrOutputWithContext(context.Context) OSProfilePtrOutput
}

type osprofilePtrType OSProfileArgs

func OSProfilePtr(v *OSProfileArgs) OSProfilePtrInput {
	return (*osprofilePtrType)(v)
}

func (*osprofilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfile)(nil)).Elem()
}

func (i *osprofilePtrType) ToOSProfilePtrOutput() OSProfilePtrOutput {
	return i.ToOSProfilePtrOutputWithContext(context.Background())
}

func (i *osprofilePtrType) ToOSProfilePtrOutputWithContext(ctx context.Context) OSProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfilePtrOutput)
}

type OSProfileOutput struct{ *pulumi.OutputState }

func (OSProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfile)(nil)).Elem()
}

func (o OSProfileOutput) ToOSProfileOutput() OSProfileOutput {
	return o
}

func (o OSProfileOutput) ToOSProfileOutputWithContext(ctx context.Context) OSProfileOutput {
	return o
}

func (o OSProfileOutput) ToOSProfilePtrOutput() OSProfilePtrOutput {
	return o.ToOSProfilePtrOutputWithContext(context.Background())
}

func (o OSProfileOutput) ToOSProfilePtrOutputWithContext(ctx context.Context) OSProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OSProfile) *OSProfile {
		return &v
	}).(OSProfilePtrOutput)
}

func (o OSProfileOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfile) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o OSProfileOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfile) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o OSProfileOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfile) *string { return v.ComputerName }).(pulumi.StringPtrOutput)
}

func (o OSProfileOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfile) *string { return v.CustomData }).(pulumi.StringPtrOutput)
}

func (o OSProfileOutput) LinuxConfiguration() LinuxConfigurationPtrOutput {
	return o.ApplyT(func(v OSProfile) *LinuxConfiguration { return v.LinuxConfiguration }).(LinuxConfigurationPtrOutput)
}

func (o OSProfileOutput) Secrets() VaultSecretGroupArrayOutput {
	return o.ApplyT(func(v OSProfile) []VaultSecretGroup { return v.Secrets }).(VaultSecretGroupArrayOutput)
}

func (o OSProfileOutput) WindowsConfiguration() WindowsConfigurationPtrOutput {
	return o.ApplyT(func(v OSProfile) *WindowsConfiguration { return v.WindowsConfiguration }).(WindowsConfigurationPtrOutput)
}

type OSProfilePtrOutput struct{ *pulumi.OutputState }

func (OSProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfile)(nil)).Elem()
}

func (o OSProfilePtrOutput) ToOSProfilePtrOutput() OSProfilePtrOutput {
	return o
}

func (o OSProfilePtrOutput) ToOSProfilePtrOutputWithContext(ctx context.Context) OSProfilePtrOutput {
	return o
}

func (o OSProfilePtrOutput) Elem() OSProfileOutput {
	return o.ApplyT(func(v *OSProfile) OSProfile {
		if v != nil {
			return *v
		}
		var ret OSProfile
		return ret
	}).(OSProfileOutput)
}

func (o OSProfilePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o OSProfilePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o OSProfilePtrOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfile) *string {
		if v == nil {
			return nil
		}
		return v.ComputerName
	}).(pulumi.StringPtrOutput)
}

func (o OSProfilePtrOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfile) *string {
		if v == nil {
			return nil
		}
		return v.CustomData
	}).(pulumi.StringPtrOutput)
}

func (o OSProfilePtrOutput) LinuxConfiguration() LinuxConfigurationPtrOutput {
	return o.ApplyT(func(v *OSProfile) *LinuxConfiguration {
		if v == nil {
			return nil
		}
		return v.LinuxConfiguration
	}).(LinuxConfigurationPtrOutput)
}

func (o OSProfilePtrOutput) Secrets() VaultSecretGroupArrayOutput {
	return o.ApplyT(func(v *OSProfile) []VaultSecretGroup {
		if v == nil {
			return nil
		}
		return v.Secrets
	}).(VaultSecretGroupArrayOutput)
}

func (o OSProfilePtrOutput) WindowsConfiguration() WindowsConfigurationPtrOutput {
	return o.ApplyT(func(v *OSProfile) *WindowsConfiguration {
		if v == nil {
			return nil
		}
		return v.WindowsConfiguration
	}).(WindowsConfigurationPtrOutput)
}

type OSProfileResponse struct {
	AdminPassword        *string                       `pulumi:"adminPassword"`
	AdminUsername        *string                       `pulumi:"adminUsername"`
	ComputerName         *string                       `pulumi:"computerName"`
	CustomData           *string                       `pulumi:"customData"`
	LinuxConfiguration   *LinuxConfigurationResponse   `pulumi:"linuxConfiguration"`
	Secrets              []VaultSecretGroupResponse    `pulumi:"secrets"`
	WindowsConfiguration *WindowsConfigurationResponse `pulumi:"windowsConfiguration"`
}

type OSProfileResponseOutput struct{ *pulumi.OutputState }

func (OSProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfileResponse)(nil)).Elem()
}

func (o OSProfileResponseOutput) ToOSProfileResponseOutput() OSProfileResponseOutput {
	return o
}

func (o OSProfileResponseOutput) ToOSProfileResponseOutputWithContext(ctx context.Context) OSProfileResponseOutput {
	return o
}

func (o OSProfileResponseOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o OSProfileResponseOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o OSProfileResponseOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *string { return v.ComputerName }).(pulumi.StringPtrOutput)
}

func (o OSProfileResponseOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *string { return v.CustomData }).(pulumi.StringPtrOutput)
}

func (o OSProfileResponseOutput) LinuxConfiguration() LinuxConfigurationResponsePtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *LinuxConfigurationResponse { return v.LinuxConfiguration }).(LinuxConfigurationResponsePtrOutput)
}

func (o OSProfileResponseOutput) Secrets() VaultSecretGroupResponseArrayOutput {
	return o.ApplyT(func(v OSProfileResponse) []VaultSecretGroupResponse { return v.Secrets }).(VaultSecretGroupResponseArrayOutput)
}

func (o OSProfileResponseOutput) WindowsConfiguration() WindowsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *WindowsConfigurationResponse { return v.WindowsConfiguration }).(WindowsConfigurationResponsePtrOutput)
}

type OSProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (OSProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfileResponse)(nil)).Elem()
}

func (o OSProfileResponsePtrOutput) ToOSProfileResponsePtrOutput() OSProfileResponsePtrOutput {
	return o
}

func (o OSProfileResponsePtrOutput) ToOSProfileResponsePtrOutputWithContext(ctx context.Context) OSProfileResponsePtrOutput {
	return o
}

func (o OSProfileResponsePtrOutput) Elem() OSProfileResponseOutput {
	return o.ApplyT(func(v *OSProfileResponse) OSProfileResponse {
		if v != nil {
			return *v
		}
		var ret OSProfileResponse
		return ret
	}).(OSProfileResponseOutput)
}

func (o OSProfileResponsePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o OSProfileResponsePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o OSProfileResponsePtrOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ComputerName
	}).(pulumi.StringPtrOutput)
}

func (o OSProfileResponsePtrOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomData
	}).(pulumi.StringPtrOutput)
}

func (o OSProfileResponsePtrOutput) LinuxConfiguration() LinuxConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *LinuxConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.LinuxConfiguration
	}).(LinuxConfigurationResponsePtrOutput)
}

func (o OSProfileResponsePtrOutput) Secrets() VaultSecretGroupResponseArrayOutput {
	return o.ApplyT(func(v *OSProfileResponse) []VaultSecretGroupResponse {
		if v == nil {
			return nil
		}
		return v.Secrets
	}).(VaultSecretGroupResponseArrayOutput)
}

func (o OSProfileResponsePtrOutput) WindowsConfiguration() WindowsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *WindowsConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.WindowsConfiguration
	}).(WindowsConfigurationResponsePtrOutput)
}

type Plan struct {
	Name          *string `pulumi:"name"`
	Product       *string `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     *string `pulumi:"publisher"`
}





type PlanInput interface {
	pulumi.Input

	ToPlanOutput() PlanOutput
	ToPlanOutputWithContext(context.Context) PlanOutput
}

type PlanArgs struct {
	Name          pulumi.StringPtrInput `pulumi:"name"`
	Product       pulumi.StringPtrInput `pulumi:"product"`
	PromotionCode pulumi.StringPtrInput `pulumi:"promotionCode"`
	Publisher     pulumi.StringPtrInput `pulumi:"publisher"`
}

func (PlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (i PlanArgs) ToPlanOutput() PlanOutput {
	return i.ToPlanOutputWithContext(context.Background())
}

func (i PlanArgs) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanOutput)
}

func (i PlanArgs) ToPlanPtrOutput() PlanPtrOutput {
	return i.ToPlanPtrOutputWithContext(context.Background())
}

func (i PlanArgs) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanOutput).ToPlanPtrOutputWithContext(ctx)
}









type PlanPtrInput interface {
	pulumi.Input

	ToPlanPtrOutput() PlanPtrOutput
	ToPlanPtrOutputWithContext(context.Context) PlanPtrOutput
}

type planPtrType PlanArgs

func PlanPtr(v *PlanArgs) PlanPtrInput {
	return (*planPtrType)(v)
}

func (*planPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Plan)(nil)).Elem()
}

func (i *planPtrType) ToPlanPtrOutput() PlanPtrOutput {
	return i.ToPlanPtrOutputWithContext(context.Background())
}

func (i *planPtrType) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanPtrOutput)
}

type PlanOutput struct{ *pulumi.OutputState }

func (PlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (o PlanOutput) ToPlanOutput() PlanOutput {
	return o
}

func (o PlanOutput) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return o
}

func (o PlanOutput) ToPlanPtrOutput() PlanPtrOutput {
	return o.ToPlanPtrOutputWithContext(context.Background())
}

func (o PlanOutput) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Plan) *Plan {
		return &v
	}).(PlanPtrOutput)
}

func (o PlanOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PlanOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o PlanOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o PlanOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

type PlanPtrOutput struct{ *pulumi.OutputState }

func (PlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Plan)(nil)).Elem()
}

func (o PlanPtrOutput) ToPlanPtrOutput() PlanPtrOutput {
	return o
}

func (o PlanPtrOutput) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return o
}

func (o PlanPtrOutput) Elem() PlanOutput {
	return o.ApplyT(func(v *Plan) Plan {
		if v != nil {
			return *v
		}
		var ret Plan
		return ret
	}).(PlanOutput)
}

func (o PlanPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return v.Product
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

type PlanResponse struct {
	Name          *string `pulumi:"name"`
	Product       *string `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     *string `pulumi:"publisher"`
}

type PlanResponseOutput struct{ *pulumi.OutputState }

func (PlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanResponse)(nil)).Elem()
}

func (o PlanResponseOutput) ToPlanResponseOutput() PlanResponseOutput {
	return o
}

func (o PlanResponseOutput) ToPlanResponseOutputWithContext(ctx context.Context) PlanResponseOutput {
	return o
}

func (o PlanResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PlanResponseOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o PlanResponseOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o PlanResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

type PlanResponsePtrOutput struct{ *pulumi.OutputState }

func (PlanResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlanResponse)(nil)).Elem()
}

func (o PlanResponsePtrOutput) ToPlanResponsePtrOutput() PlanResponsePtrOutput {
	return o
}

func (o PlanResponsePtrOutput) ToPlanResponsePtrOutputWithContext(ctx context.Context) PlanResponsePtrOutput {
	return o
}

func (o PlanResponsePtrOutput) Elem() PlanResponseOutput {
	return o.ApplyT(func(v *PlanResponse) PlanResponse {
		if v != nil {
			return *v
		}
		var ret PlanResponse
		return ret
	}).(PlanResponseOutput)
}

func (o PlanResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Product
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Capacity *float64 `pulumi:"capacity"`
	Name     *string  `pulumi:"name"`
	Tier     *string  `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.Float64PtrInput `pulumi:"capacity"`
	Name     pulumi.StringPtrInput  `pulumi:"name"`
	Tier     pulumi.StringPtrInput  `pulumi:"tier"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v Sku) *float64 { return v.Capacity }).(pulumi.Float64PtrOutput)
}

func (o SkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Sku) *float64 {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.Float64PtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Capacity *float64 `pulumi:"capacity"`
	Name     *string  `pulumi:"name"`
	Tier     *string  `pulumi:"tier"`
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SkuResponse) *float64 { return v.Capacity }).(pulumi.Float64PtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SkuResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.Float64PtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SshConfiguration struct {
	PublicKeys []SshPublicKey `pulumi:"publicKeys"`
}





type SshConfigurationInput interface {
	pulumi.Input

	ToSshConfigurationOutput() SshConfigurationOutput
	ToSshConfigurationOutputWithContext(context.Context) SshConfigurationOutput
}

type SshConfigurationArgs struct {
	PublicKeys SshPublicKeyArrayInput `pulumi:"publicKeys"`
}

func (SshConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SshConfiguration)(nil)).Elem()
}

func (i SshConfigurationArgs) ToSshConfigurationOutput() SshConfigurationOutput {
	return i.ToSshConfigurationOutputWithContext(context.Background())
}

func (i SshConfigurationArgs) ToSshConfigurationOutputWithContext(ctx context.Context) SshConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshConfigurationOutput)
}

func (i SshConfigurationArgs) ToSshConfigurationPtrOutput() SshConfigurationPtrOutput {
	return i.ToSshConfigurationPtrOutputWithContext(context.Background())
}

func (i SshConfigurationArgs) ToSshConfigurationPtrOutputWithContext(ctx context.Context) SshConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshConfigurationOutput).ToSshConfigurationPtrOutputWithContext(ctx)
}









type SshConfigurationPtrInput interface {
	pulumi.Input

	ToSshConfigurationPtrOutput() SshConfigurationPtrOutput
	ToSshConfigurationPtrOutputWithContext(context.Context) SshConfigurationPtrOutput
}

type sshConfigurationPtrType SshConfigurationArgs

func SshConfigurationPtr(v *SshConfigurationArgs) SshConfigurationPtrInput {
	return (*sshConfigurationPtrType)(v)
}

func (*sshConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SshConfiguration)(nil)).Elem()
}

func (i *sshConfigurationPtrType) ToSshConfigurationPtrOutput() SshConfigurationPtrOutput {
	return i.ToSshConfigurationPtrOutputWithContext(context.Background())
}

func (i *sshConfigurationPtrType) ToSshConfigurationPtrOutputWithContext(ctx context.Context) SshConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshConfigurationPtrOutput)
}

type SshConfigurationOutput struct{ *pulumi.OutputState }

func (SshConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshConfiguration)(nil)).Elem()
}

func (o SshConfigurationOutput) ToSshConfigurationOutput() SshConfigurationOutput {
	return o
}

func (o SshConfigurationOutput) ToSshConfigurationOutputWithContext(ctx context.Context) SshConfigurationOutput {
	return o
}

func (o SshConfigurationOutput) ToSshConfigurationPtrOutput() SshConfigurationPtrOutput {
	return o.ToSshConfigurationPtrOutputWithContext(context.Background())
}

func (o SshConfigurationOutput) ToSshConfigurationPtrOutputWithContext(ctx context.Context) SshConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SshConfiguration) *SshConfiguration {
		return &v
	}).(SshConfigurationPtrOutput)
}

func (o SshConfigurationOutput) PublicKeys() SshPublicKeyArrayOutput {
	return o.ApplyT(func(v SshConfiguration) []SshPublicKey { return v.PublicKeys }).(SshPublicKeyArrayOutput)
}

type SshConfigurationPtrOutput struct{ *pulumi.OutputState }

func (SshConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SshConfiguration)(nil)).Elem()
}

func (o SshConfigurationPtrOutput) ToSshConfigurationPtrOutput() SshConfigurationPtrOutput {
	return o
}

func (o SshConfigurationPtrOutput) ToSshConfigurationPtrOutputWithContext(ctx context.Context) SshConfigurationPtrOutput {
	return o
}

func (o SshConfigurationPtrOutput) Elem() SshConfigurationOutput {
	return o.ApplyT(func(v *SshConfiguration) SshConfiguration {
		if v != nil {
			return *v
		}
		var ret SshConfiguration
		return ret
	}).(SshConfigurationOutput)
}

func (o SshConfigurationPtrOutput) PublicKeys() SshPublicKeyArrayOutput {
	return o.ApplyT(func(v *SshConfiguration) []SshPublicKey {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(SshPublicKeyArrayOutput)
}

type SshConfigurationResponse struct {
	PublicKeys []SshPublicKeyResponse `pulumi:"publicKeys"`
}

type SshConfigurationResponseOutput struct{ *pulumi.OutputState }

func (SshConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshConfigurationResponse)(nil)).Elem()
}

func (o SshConfigurationResponseOutput) ToSshConfigurationResponseOutput() SshConfigurationResponseOutput {
	return o
}

func (o SshConfigurationResponseOutput) ToSshConfigurationResponseOutputWithContext(ctx context.Context) SshConfigurationResponseOutput {
	return o
}

func (o SshConfigurationResponseOutput) PublicKeys() SshPublicKeyResponseArrayOutput {
	return o.ApplyT(func(v SshConfigurationResponse) []SshPublicKeyResponse { return v.PublicKeys }).(SshPublicKeyResponseArrayOutput)
}

type SshConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (SshConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SshConfigurationResponse)(nil)).Elem()
}

func (o SshConfigurationResponsePtrOutput) ToSshConfigurationResponsePtrOutput() SshConfigurationResponsePtrOutput {
	return o
}

func (o SshConfigurationResponsePtrOutput) ToSshConfigurationResponsePtrOutputWithContext(ctx context.Context) SshConfigurationResponsePtrOutput {
	return o
}

func (o SshConfigurationResponsePtrOutput) Elem() SshConfigurationResponseOutput {
	return o.ApplyT(func(v *SshConfigurationResponse) SshConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret SshConfigurationResponse
		return ret
	}).(SshConfigurationResponseOutput)
}

func (o SshConfigurationResponsePtrOutput) PublicKeys() SshPublicKeyResponseArrayOutput {
	return o.ApplyT(func(v *SshConfigurationResponse) []SshPublicKeyResponse {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(SshPublicKeyResponseArrayOutput)
}

type SshPublicKey struct {
	KeyData *string `pulumi:"keyData"`
	Path    *string `pulumi:"path"`
}





type SshPublicKeyInput interface {
	pulumi.Input

	ToSshPublicKeyOutput() SshPublicKeyOutput
	ToSshPublicKeyOutputWithContext(context.Context) SshPublicKeyOutput
}

type SshPublicKeyArgs struct {
	KeyData pulumi.StringPtrInput `pulumi:"keyData"`
	Path    pulumi.StringPtrInput `pulumi:"path"`
}

func (SshPublicKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SshPublicKey)(nil)).Elem()
}

func (i SshPublicKeyArgs) ToSshPublicKeyOutput() SshPublicKeyOutput {
	return i.ToSshPublicKeyOutputWithContext(context.Background())
}

func (i SshPublicKeyArgs) ToSshPublicKeyOutputWithContext(ctx context.Context) SshPublicKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshPublicKeyOutput)
}





type SshPublicKeyArrayInput interface {
	pulumi.Input

	ToSshPublicKeyArrayOutput() SshPublicKeyArrayOutput
	ToSshPublicKeyArrayOutputWithContext(context.Context) SshPublicKeyArrayOutput
}

type SshPublicKeyArray []SshPublicKeyInput

func (SshPublicKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SshPublicKey)(nil)).Elem()
}

func (i SshPublicKeyArray) ToSshPublicKeyArrayOutput() SshPublicKeyArrayOutput {
	return i.ToSshPublicKeyArrayOutputWithContext(context.Background())
}

func (i SshPublicKeyArray) ToSshPublicKeyArrayOutputWithContext(ctx context.Context) SshPublicKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshPublicKeyArrayOutput)
}

type SshPublicKeyOutput struct{ *pulumi.OutputState }

func (SshPublicKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshPublicKey)(nil)).Elem()
}

func (o SshPublicKeyOutput) ToSshPublicKeyOutput() SshPublicKeyOutput {
	return o
}

func (o SshPublicKeyOutput) ToSshPublicKeyOutputWithContext(ctx context.Context) SshPublicKeyOutput {
	return o
}

func (o SshPublicKeyOutput) KeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SshPublicKey) *string { return v.KeyData }).(pulumi.StringPtrOutput)
}

func (o SshPublicKeyOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SshPublicKey) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type SshPublicKeyArrayOutput struct{ *pulumi.OutputState }

func (SshPublicKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SshPublicKey)(nil)).Elem()
}

func (o SshPublicKeyArrayOutput) ToSshPublicKeyArrayOutput() SshPublicKeyArrayOutput {
	return o
}

func (o SshPublicKeyArrayOutput) ToSshPublicKeyArrayOutputWithContext(ctx context.Context) SshPublicKeyArrayOutput {
	return o
}

func (o SshPublicKeyArrayOutput) Index(i pulumi.IntInput) SshPublicKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SshPublicKey {
		return vs[0].([]SshPublicKey)[vs[1].(int)]
	}).(SshPublicKeyOutput)
}

type SshPublicKeyResponse struct {
	KeyData *string `pulumi:"keyData"`
	Path    *string `pulumi:"path"`
}

type SshPublicKeyResponseOutput struct{ *pulumi.OutputState }

func (SshPublicKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshPublicKeyResponse)(nil)).Elem()
}

func (o SshPublicKeyResponseOutput) ToSshPublicKeyResponseOutput() SshPublicKeyResponseOutput {
	return o
}

func (o SshPublicKeyResponseOutput) ToSshPublicKeyResponseOutputWithContext(ctx context.Context) SshPublicKeyResponseOutput {
	return o
}

func (o SshPublicKeyResponseOutput) KeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SshPublicKeyResponse) *string { return v.KeyData }).(pulumi.StringPtrOutput)
}

func (o SshPublicKeyResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SshPublicKeyResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type SshPublicKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (SshPublicKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SshPublicKeyResponse)(nil)).Elem()
}

func (o SshPublicKeyResponseArrayOutput) ToSshPublicKeyResponseArrayOutput() SshPublicKeyResponseArrayOutput {
	return o
}

func (o SshPublicKeyResponseArrayOutput) ToSshPublicKeyResponseArrayOutputWithContext(ctx context.Context) SshPublicKeyResponseArrayOutput {
	return o
}

func (o SshPublicKeyResponseArrayOutput) Index(i pulumi.IntInput) SshPublicKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SshPublicKeyResponse {
		return vs[0].([]SshPublicKeyResponse)[vs[1].(int)]
	}).(SshPublicKeyResponseOutput)
}

type StorageProfile struct {
	DataDisks      []DataDisk      `pulumi:"dataDisks"`
	ImageReference *ImageReference `pulumi:"imageReference"`
	OsDisk         *OSDisk         `pulumi:"osDisk"`
}





type StorageProfileInput interface {
	pulumi.Input

	ToStorageProfileOutput() StorageProfileOutput
	ToStorageProfileOutputWithContext(context.Context) StorageProfileOutput
}

type StorageProfileArgs struct {
	DataDisks      DataDiskArrayInput     `pulumi:"dataDisks"`
	ImageReference ImageReferencePtrInput `pulumi:"imageReference"`
	OsDisk         OSDiskPtrInput         `pulumi:"osDisk"`
}

func (StorageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfile)(nil)).Elem()
}

func (i StorageProfileArgs) ToStorageProfileOutput() StorageProfileOutput {
	return i.ToStorageProfileOutputWithContext(context.Background())
}

func (i StorageProfileArgs) ToStorageProfileOutputWithContext(ctx context.Context) StorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileOutput)
}

func (i StorageProfileArgs) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return i.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (i StorageProfileArgs) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileOutput).ToStorageProfilePtrOutputWithContext(ctx)
}









type StorageProfilePtrInput interface {
	pulumi.Input

	ToStorageProfilePtrOutput() StorageProfilePtrOutput
	ToStorageProfilePtrOutputWithContext(context.Context) StorageProfilePtrOutput
}

type storageProfilePtrType StorageProfileArgs

func StorageProfilePtr(v *StorageProfileArgs) StorageProfilePtrInput {
	return (*storageProfilePtrType)(v)
}

func (*storageProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfile)(nil)).Elem()
}

func (i *storageProfilePtrType) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return i.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (i *storageProfilePtrType) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfilePtrOutput)
}

type StorageProfileOutput struct{ *pulumi.OutputState }

func (StorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfile)(nil)).Elem()
}

func (o StorageProfileOutput) ToStorageProfileOutput() StorageProfileOutput {
	return o
}

func (o StorageProfileOutput) ToStorageProfileOutputWithContext(ctx context.Context) StorageProfileOutput {
	return o
}

func (o StorageProfileOutput) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return o.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (o StorageProfileOutput) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageProfile) *StorageProfile {
		return &v
	}).(StorageProfilePtrOutput)
}

func (o StorageProfileOutput) DataDisks() DataDiskArrayOutput {
	return o.ApplyT(func(v StorageProfile) []DataDisk { return v.DataDisks }).(DataDiskArrayOutput)
}

func (o StorageProfileOutput) ImageReference() ImageReferencePtrOutput {
	return o.ApplyT(func(v StorageProfile) *ImageReference { return v.ImageReference }).(ImageReferencePtrOutput)
}

func (o StorageProfileOutput) OsDisk() OSDiskPtrOutput {
	return o.ApplyT(func(v StorageProfile) *OSDisk { return v.OsDisk }).(OSDiskPtrOutput)
}

type StorageProfilePtrOutput struct{ *pulumi.OutputState }

func (StorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfile)(nil)).Elem()
}

func (o StorageProfilePtrOutput) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return o
}

func (o StorageProfilePtrOutput) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return o
}

func (o StorageProfilePtrOutput) Elem() StorageProfileOutput {
	return o.ApplyT(func(v *StorageProfile) StorageProfile {
		if v != nil {
			return *v
		}
		var ret StorageProfile
		return ret
	}).(StorageProfileOutput)
}

func (o StorageProfilePtrOutput) DataDisks() DataDiskArrayOutput {
	return o.ApplyT(func(v *StorageProfile) []DataDisk {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(DataDiskArrayOutput)
}

func (o StorageProfilePtrOutput) ImageReference() ImageReferencePtrOutput {
	return o.ApplyT(func(v *StorageProfile) *ImageReference {
		if v == nil {
			return nil
		}
		return v.ImageReference
	}).(ImageReferencePtrOutput)
}

func (o StorageProfilePtrOutput) OsDisk() OSDiskPtrOutput {
	return o.ApplyT(func(v *StorageProfile) *OSDisk {
		if v == nil {
			return nil
		}
		return v.OsDisk
	}).(OSDiskPtrOutput)
}

type StorageProfileResponse struct {
	DataDisks      []DataDiskResponse      `pulumi:"dataDisks"`
	ImageReference *ImageReferenceResponse `pulumi:"imageReference"`
	OsDisk         *OSDiskResponse         `pulumi:"osDisk"`
}

type StorageProfileResponseOutput struct{ *pulumi.OutputState }

func (StorageProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfileResponse)(nil)).Elem()
}

func (o StorageProfileResponseOutput) ToStorageProfileResponseOutput() StorageProfileResponseOutput {
	return o
}

func (o StorageProfileResponseOutput) ToStorageProfileResponseOutputWithContext(ctx context.Context) StorageProfileResponseOutput {
	return o
}

func (o StorageProfileResponseOutput) DataDisks() DataDiskResponseArrayOutput {
	return o.ApplyT(func(v StorageProfileResponse) []DataDiskResponse { return v.DataDisks }).(DataDiskResponseArrayOutput)
}

func (o StorageProfileResponseOutput) ImageReference() ImageReferenceResponsePtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *ImageReferenceResponse { return v.ImageReference }).(ImageReferenceResponsePtrOutput)
}

func (o StorageProfileResponseOutput) OsDisk() OSDiskResponsePtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *OSDiskResponse { return v.OsDisk }).(OSDiskResponsePtrOutput)
}

type StorageProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfileResponse)(nil)).Elem()
}

func (o StorageProfileResponsePtrOutput) ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput {
	return o
}

func (o StorageProfileResponsePtrOutput) ToStorageProfileResponsePtrOutputWithContext(ctx context.Context) StorageProfileResponsePtrOutput {
	return o
}

func (o StorageProfileResponsePtrOutput) Elem() StorageProfileResponseOutput {
	return o.ApplyT(func(v *StorageProfileResponse) StorageProfileResponse {
		if v != nil {
			return *v
		}
		var ret StorageProfileResponse
		return ret
	}).(StorageProfileResponseOutput)
}

func (o StorageProfileResponsePtrOutput) DataDisks() DataDiskResponseArrayOutput {
	return o.ApplyT(func(v *StorageProfileResponse) []DataDiskResponse {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(DataDiskResponseArrayOutput)
}

func (o StorageProfileResponsePtrOutput) ImageReference() ImageReferenceResponsePtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *ImageReferenceResponse {
		if v == nil {
			return nil
		}
		return v.ImageReference
	}).(ImageReferenceResponsePtrOutput)
}

func (o StorageProfileResponsePtrOutput) OsDisk() OSDiskResponsePtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *OSDiskResponse {
		if v == nil {
			return nil
		}
		return v.OsDisk
	}).(OSDiskResponsePtrOutput)
}

type SubResource struct {
	Id *string `pulumi:"id"`
}





type SubResourceInput interface {
	pulumi.Input

	ToSubResourceOutput() SubResourceOutput
	ToSubResourceOutputWithContext(context.Context) SubResourceOutput
}

type SubResourceArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SubResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (i SubResourceArgs) ToSubResourceOutput() SubResourceOutput {
	return i.ToSubResourceOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput)
}

func (i SubResourceArgs) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return i.ToSubResourcePtrOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput).ToSubResourcePtrOutputWithContext(ctx)
}









type SubResourcePtrInput interface {
	pulumi.Input

	ToSubResourcePtrOutput() SubResourcePtrOutput
	ToSubResourcePtrOutputWithContext(context.Context) SubResourcePtrOutput
}

type subResourcePtrType SubResourceArgs

func SubResourcePtr(v *SubResourceArgs) SubResourcePtrInput {
	return (*subResourcePtrType)(v)
}

func (*subResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResource)(nil)).Elem()
}

func (i *subResourcePtrType) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return i.ToSubResourcePtrOutputWithContext(context.Background())
}

func (i *subResourcePtrType) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourcePtrOutput)
}





type SubResourceArrayInput interface {
	pulumi.Input

	ToSubResourceArrayOutput() SubResourceArrayOutput
	ToSubResourceArrayOutputWithContext(context.Context) SubResourceArrayOutput
}

type SubResourceArray []SubResourceInput

func (SubResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResource)(nil)).Elem()
}

func (i SubResourceArray) ToSubResourceArrayOutput() SubResourceArrayOutput {
	return i.ToSubResourceArrayOutputWithContext(context.Background())
}

func (i SubResourceArray) ToSubResourceArrayOutputWithContext(ctx context.Context) SubResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceArrayOutput)
}

type SubResourceOutput struct{ *pulumi.OutputState }

func (SubResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (o SubResourceOutput) ToSubResourceOutput() SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return o.ToSubResourcePtrOutputWithContext(context.Background())
}

func (o SubResourceOutput) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubResource) *SubResource {
		return &v
	}).(SubResourcePtrOutput)
}

func (o SubResourceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResource) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourcePtrOutput struct{ *pulumi.OutputState }

func (SubResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResource)(nil)).Elem()
}

func (o SubResourcePtrOutput) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return o
}

func (o SubResourcePtrOutput) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return o
}

func (o SubResourcePtrOutput) Elem() SubResourceOutput {
	return o.ApplyT(func(v *SubResource) SubResource {
		if v != nil {
			return *v
		}
		var ret SubResource
		return ret
	}).(SubResourceOutput)
}

func (o SubResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResource) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SubResourceArrayOutput struct{ *pulumi.OutputState }

func (SubResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResource)(nil)).Elem()
}

func (o SubResourceArrayOutput) ToSubResourceArrayOutput() SubResourceArrayOutput {
	return o
}

func (o SubResourceArrayOutput) ToSubResourceArrayOutputWithContext(ctx context.Context) SubResourceArrayOutput {
	return o
}

func (o SubResourceArrayOutput) Index(i pulumi.IntInput) SubResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubResource {
		return vs[0].([]SubResource)[vs[1].(int)]
	}).(SubResourceOutput)
}

type SubResourceResponse struct {
	Id *string `pulumi:"id"`
}

type SubResourceResponseOutput struct{ *pulumi.OutputState }

func (SubResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutput() SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutputWithContext(ctx context.Context) SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResourceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourceResponsePtrOutput struct{ *pulumi.OutputState }

func (SubResourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) Elem() SubResourceResponseOutput {
	return o.ApplyT(func(v *SubResourceResponse) SubResourceResponse {
		if v != nil {
			return *v
		}
		var ret SubResourceResponse
		return ret
	}).(SubResourceResponseOutput)
}

func (o SubResourceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SubResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (SubResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseArrayOutput) ToSubResourceResponseArrayOutput() SubResourceResponseArrayOutput {
	return o
}

func (o SubResourceResponseArrayOutput) ToSubResourceResponseArrayOutputWithContext(ctx context.Context) SubResourceResponseArrayOutput {
	return o
}

func (o SubResourceResponseArrayOutput) Index(i pulumi.IntInput) SubResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubResourceResponse {
		return vs[0].([]SubResourceResponse)[vs[1].(int)]
	}).(SubResourceResponseOutput)
}

type UpgradePolicy struct {
	Mode *UpgradeMode `pulumi:"mode"`
}





type UpgradePolicyInput interface {
	pulumi.Input

	ToUpgradePolicyOutput() UpgradePolicyOutput
	ToUpgradePolicyOutputWithContext(context.Context) UpgradePolicyOutput
}

type UpgradePolicyArgs struct {
	Mode UpgradeModePtrInput `pulumi:"mode"`
}

func (UpgradePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UpgradePolicy)(nil)).Elem()
}

func (i UpgradePolicyArgs) ToUpgradePolicyOutput() UpgradePolicyOutput {
	return i.ToUpgradePolicyOutputWithContext(context.Background())
}

func (i UpgradePolicyArgs) ToUpgradePolicyOutputWithContext(ctx context.Context) UpgradePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpgradePolicyOutput)
}

func (i UpgradePolicyArgs) ToUpgradePolicyPtrOutput() UpgradePolicyPtrOutput {
	return i.ToUpgradePolicyPtrOutputWithContext(context.Background())
}

func (i UpgradePolicyArgs) ToUpgradePolicyPtrOutputWithContext(ctx context.Context) UpgradePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpgradePolicyOutput).ToUpgradePolicyPtrOutputWithContext(ctx)
}









type UpgradePolicyPtrInput interface {
	pulumi.Input

	ToUpgradePolicyPtrOutput() UpgradePolicyPtrOutput
	ToUpgradePolicyPtrOutputWithContext(context.Context) UpgradePolicyPtrOutput
}

type upgradePolicyPtrType UpgradePolicyArgs

func UpgradePolicyPtr(v *UpgradePolicyArgs) UpgradePolicyPtrInput {
	return (*upgradePolicyPtrType)(v)
}

func (*upgradePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UpgradePolicy)(nil)).Elem()
}

func (i *upgradePolicyPtrType) ToUpgradePolicyPtrOutput() UpgradePolicyPtrOutput {
	return i.ToUpgradePolicyPtrOutputWithContext(context.Background())
}

func (i *upgradePolicyPtrType) ToUpgradePolicyPtrOutputWithContext(ctx context.Context) UpgradePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpgradePolicyPtrOutput)
}

type UpgradePolicyOutput struct{ *pulumi.OutputState }

func (UpgradePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpgradePolicy)(nil)).Elem()
}

func (o UpgradePolicyOutput) ToUpgradePolicyOutput() UpgradePolicyOutput {
	return o
}

func (o UpgradePolicyOutput) ToUpgradePolicyOutputWithContext(ctx context.Context) UpgradePolicyOutput {
	return o
}

func (o UpgradePolicyOutput) ToUpgradePolicyPtrOutput() UpgradePolicyPtrOutput {
	return o.ToUpgradePolicyPtrOutputWithContext(context.Background())
}

func (o UpgradePolicyOutput) ToUpgradePolicyPtrOutputWithContext(ctx context.Context) UpgradePolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UpgradePolicy) *UpgradePolicy {
		return &v
	}).(UpgradePolicyPtrOutput)
}

func (o UpgradePolicyOutput) Mode() UpgradeModePtrOutput {
	return o.ApplyT(func(v UpgradePolicy) *UpgradeMode { return v.Mode }).(UpgradeModePtrOutput)
}

type UpgradePolicyPtrOutput struct{ *pulumi.OutputState }

func (UpgradePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpgradePolicy)(nil)).Elem()
}

func (o UpgradePolicyPtrOutput) ToUpgradePolicyPtrOutput() UpgradePolicyPtrOutput {
	return o
}

func (o UpgradePolicyPtrOutput) ToUpgradePolicyPtrOutputWithContext(ctx context.Context) UpgradePolicyPtrOutput {
	return o
}

func (o UpgradePolicyPtrOutput) Elem() UpgradePolicyOutput {
	return o.ApplyT(func(v *UpgradePolicy) UpgradePolicy {
		if v != nil {
			return *v
		}
		var ret UpgradePolicy
		return ret
	}).(UpgradePolicyOutput)
}

func (o UpgradePolicyPtrOutput) Mode() UpgradeModePtrOutput {
	return o.ApplyT(func(v *UpgradePolicy) *UpgradeMode {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(UpgradeModePtrOutput)
}

type UpgradePolicyResponse struct {
	Mode *string `pulumi:"mode"`
}

type UpgradePolicyResponseOutput struct{ *pulumi.OutputState }

func (UpgradePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpgradePolicyResponse)(nil)).Elem()
}

func (o UpgradePolicyResponseOutput) ToUpgradePolicyResponseOutput() UpgradePolicyResponseOutput {
	return o
}

func (o UpgradePolicyResponseOutput) ToUpgradePolicyResponseOutputWithContext(ctx context.Context) UpgradePolicyResponseOutput {
	return o
}

func (o UpgradePolicyResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UpgradePolicyResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

type UpgradePolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (UpgradePolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpgradePolicyResponse)(nil)).Elem()
}

func (o UpgradePolicyResponsePtrOutput) ToUpgradePolicyResponsePtrOutput() UpgradePolicyResponsePtrOutput {
	return o
}

func (o UpgradePolicyResponsePtrOutput) ToUpgradePolicyResponsePtrOutputWithContext(ctx context.Context) UpgradePolicyResponsePtrOutput {
	return o
}

func (o UpgradePolicyResponsePtrOutput) Elem() UpgradePolicyResponseOutput {
	return o.ApplyT(func(v *UpgradePolicyResponse) UpgradePolicyResponse {
		if v != nil {
			return *v
		}
		var ret UpgradePolicyResponse
		return ret
	}).(UpgradePolicyResponseOutput)
}

func (o UpgradePolicyResponsePtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

type VaultCertificate struct {
	CertificateStore *string `pulumi:"certificateStore"`
	CertificateUrl   *string `pulumi:"certificateUrl"`
}





type VaultCertificateInput interface {
	pulumi.Input

	ToVaultCertificateOutput() VaultCertificateOutput
	ToVaultCertificateOutputWithContext(context.Context) VaultCertificateOutput
}

type VaultCertificateArgs struct {
	CertificateStore pulumi.StringPtrInput `pulumi:"certificateStore"`
	CertificateUrl   pulumi.StringPtrInput `pulumi:"certificateUrl"`
}

func (VaultCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultCertificate)(nil)).Elem()
}

func (i VaultCertificateArgs) ToVaultCertificateOutput() VaultCertificateOutput {
	return i.ToVaultCertificateOutputWithContext(context.Background())
}

func (i VaultCertificateArgs) ToVaultCertificateOutputWithContext(ctx context.Context) VaultCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultCertificateOutput)
}





type VaultCertificateArrayInput interface {
	pulumi.Input

	ToVaultCertificateArrayOutput() VaultCertificateArrayOutput
	ToVaultCertificateArrayOutputWithContext(context.Context) VaultCertificateArrayOutput
}

type VaultCertificateArray []VaultCertificateInput

func (VaultCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultCertificate)(nil)).Elem()
}

func (i VaultCertificateArray) ToVaultCertificateArrayOutput() VaultCertificateArrayOutput {
	return i.ToVaultCertificateArrayOutputWithContext(context.Background())
}

func (i VaultCertificateArray) ToVaultCertificateArrayOutputWithContext(ctx context.Context) VaultCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultCertificateArrayOutput)
}

type VaultCertificateOutput struct{ *pulumi.OutputState }

func (VaultCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultCertificate)(nil)).Elem()
}

func (o VaultCertificateOutput) ToVaultCertificateOutput() VaultCertificateOutput {
	return o
}

func (o VaultCertificateOutput) ToVaultCertificateOutputWithContext(ctx context.Context) VaultCertificateOutput {
	return o
}

func (o VaultCertificateOutput) CertificateStore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultCertificate) *string { return v.CertificateStore }).(pulumi.StringPtrOutput)
}

func (o VaultCertificateOutput) CertificateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultCertificate) *string { return v.CertificateUrl }).(pulumi.StringPtrOutput)
}

type VaultCertificateArrayOutput struct{ *pulumi.OutputState }

func (VaultCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultCertificate)(nil)).Elem()
}

func (o VaultCertificateArrayOutput) ToVaultCertificateArrayOutput() VaultCertificateArrayOutput {
	return o
}

func (o VaultCertificateArrayOutput) ToVaultCertificateArrayOutputWithContext(ctx context.Context) VaultCertificateArrayOutput {
	return o
}

func (o VaultCertificateArrayOutput) Index(i pulumi.IntInput) VaultCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VaultCertificate {
		return vs[0].([]VaultCertificate)[vs[1].(int)]
	}).(VaultCertificateOutput)
}

type VaultCertificateResponse struct {
	CertificateStore *string `pulumi:"certificateStore"`
	CertificateUrl   *string `pulumi:"certificateUrl"`
}

type VaultCertificateResponseOutput struct{ *pulumi.OutputState }

func (VaultCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultCertificateResponse)(nil)).Elem()
}

func (o VaultCertificateResponseOutput) ToVaultCertificateResponseOutput() VaultCertificateResponseOutput {
	return o
}

func (o VaultCertificateResponseOutput) ToVaultCertificateResponseOutputWithContext(ctx context.Context) VaultCertificateResponseOutput {
	return o
}

func (o VaultCertificateResponseOutput) CertificateStore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultCertificateResponse) *string { return v.CertificateStore }).(pulumi.StringPtrOutput)
}

func (o VaultCertificateResponseOutput) CertificateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultCertificateResponse) *string { return v.CertificateUrl }).(pulumi.StringPtrOutput)
}

type VaultCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VaultCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultCertificateResponse)(nil)).Elem()
}

func (o VaultCertificateResponseArrayOutput) ToVaultCertificateResponseArrayOutput() VaultCertificateResponseArrayOutput {
	return o
}

func (o VaultCertificateResponseArrayOutput) ToVaultCertificateResponseArrayOutputWithContext(ctx context.Context) VaultCertificateResponseArrayOutput {
	return o
}

func (o VaultCertificateResponseArrayOutput) Index(i pulumi.IntInput) VaultCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VaultCertificateResponse {
		return vs[0].([]VaultCertificateResponse)[vs[1].(int)]
	}).(VaultCertificateResponseOutput)
}

type VaultSecretGroup struct {
	SourceVault       *SubResource       `pulumi:"sourceVault"`
	VaultCertificates []VaultCertificate `pulumi:"vaultCertificates"`
}





type VaultSecretGroupInput interface {
	pulumi.Input

	ToVaultSecretGroupOutput() VaultSecretGroupOutput
	ToVaultSecretGroupOutputWithContext(context.Context) VaultSecretGroupOutput
}

type VaultSecretGroupArgs struct {
	SourceVault       SubResourcePtrInput        `pulumi:"sourceVault"`
	VaultCertificates VaultCertificateArrayInput `pulumi:"vaultCertificates"`
}

func (VaultSecretGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultSecretGroup)(nil)).Elem()
}

func (i VaultSecretGroupArgs) ToVaultSecretGroupOutput() VaultSecretGroupOutput {
	return i.ToVaultSecretGroupOutputWithContext(context.Background())
}

func (i VaultSecretGroupArgs) ToVaultSecretGroupOutputWithContext(ctx context.Context) VaultSecretGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultSecretGroupOutput)
}





type VaultSecretGroupArrayInput interface {
	pulumi.Input

	ToVaultSecretGroupArrayOutput() VaultSecretGroupArrayOutput
	ToVaultSecretGroupArrayOutputWithContext(context.Context) VaultSecretGroupArrayOutput
}

type VaultSecretGroupArray []VaultSecretGroupInput

func (VaultSecretGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultSecretGroup)(nil)).Elem()
}

func (i VaultSecretGroupArray) ToVaultSecretGroupArrayOutput() VaultSecretGroupArrayOutput {
	return i.ToVaultSecretGroupArrayOutputWithContext(context.Background())
}

func (i VaultSecretGroupArray) ToVaultSecretGroupArrayOutputWithContext(ctx context.Context) VaultSecretGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultSecretGroupArrayOutput)
}

type VaultSecretGroupOutput struct{ *pulumi.OutputState }

func (VaultSecretGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultSecretGroup)(nil)).Elem()
}

func (o VaultSecretGroupOutput) ToVaultSecretGroupOutput() VaultSecretGroupOutput {
	return o
}

func (o VaultSecretGroupOutput) ToVaultSecretGroupOutputWithContext(ctx context.Context) VaultSecretGroupOutput {
	return o
}

func (o VaultSecretGroupOutput) SourceVault() SubResourcePtrOutput {
	return o.ApplyT(func(v VaultSecretGroup) *SubResource { return v.SourceVault }).(SubResourcePtrOutput)
}

func (o VaultSecretGroupOutput) VaultCertificates() VaultCertificateArrayOutput {
	return o.ApplyT(func(v VaultSecretGroup) []VaultCertificate { return v.VaultCertificates }).(VaultCertificateArrayOutput)
}

type VaultSecretGroupArrayOutput struct{ *pulumi.OutputState }

func (VaultSecretGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultSecretGroup)(nil)).Elem()
}

func (o VaultSecretGroupArrayOutput) ToVaultSecretGroupArrayOutput() VaultSecretGroupArrayOutput {
	return o
}

func (o VaultSecretGroupArrayOutput) ToVaultSecretGroupArrayOutputWithContext(ctx context.Context) VaultSecretGroupArrayOutput {
	return o
}

func (o VaultSecretGroupArrayOutput) Index(i pulumi.IntInput) VaultSecretGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VaultSecretGroup {
		return vs[0].([]VaultSecretGroup)[vs[1].(int)]
	}).(VaultSecretGroupOutput)
}

type VaultSecretGroupResponse struct {
	SourceVault       *SubResourceResponse       `pulumi:"sourceVault"`
	VaultCertificates []VaultCertificateResponse `pulumi:"vaultCertificates"`
}

type VaultSecretGroupResponseOutput struct{ *pulumi.OutputState }

func (VaultSecretGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultSecretGroupResponse)(nil)).Elem()
}

func (o VaultSecretGroupResponseOutput) ToVaultSecretGroupResponseOutput() VaultSecretGroupResponseOutput {
	return o
}

func (o VaultSecretGroupResponseOutput) ToVaultSecretGroupResponseOutputWithContext(ctx context.Context) VaultSecretGroupResponseOutput {
	return o
}

func (o VaultSecretGroupResponseOutput) SourceVault() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VaultSecretGroupResponse) *SubResourceResponse { return v.SourceVault }).(SubResourceResponsePtrOutput)
}

func (o VaultSecretGroupResponseOutput) VaultCertificates() VaultCertificateResponseArrayOutput {
	return o.ApplyT(func(v VaultSecretGroupResponse) []VaultCertificateResponse { return v.VaultCertificates }).(VaultCertificateResponseArrayOutput)
}

type VaultSecretGroupResponseArrayOutput struct{ *pulumi.OutputState }

func (VaultSecretGroupResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultSecretGroupResponse)(nil)).Elem()
}

func (o VaultSecretGroupResponseArrayOutput) ToVaultSecretGroupResponseArrayOutput() VaultSecretGroupResponseArrayOutput {
	return o
}

func (o VaultSecretGroupResponseArrayOutput) ToVaultSecretGroupResponseArrayOutputWithContext(ctx context.Context) VaultSecretGroupResponseArrayOutput {
	return o
}

func (o VaultSecretGroupResponseArrayOutput) Index(i pulumi.IntInput) VaultSecretGroupResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VaultSecretGroupResponse {
		return vs[0].([]VaultSecretGroupResponse)[vs[1].(int)]
	}).(VaultSecretGroupResponseOutput)
}

type VirtualHardDisk struct {
	Uri *string `pulumi:"uri"`
}





type VirtualHardDiskInput interface {
	pulumi.Input

	ToVirtualHardDiskOutput() VirtualHardDiskOutput
	ToVirtualHardDiskOutputWithContext(context.Context) VirtualHardDiskOutput
}

type VirtualHardDiskArgs struct {
	Uri pulumi.StringPtrInput `pulumi:"uri"`
}

func (VirtualHardDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHardDisk)(nil)).Elem()
}

func (i VirtualHardDiskArgs) ToVirtualHardDiskOutput() VirtualHardDiskOutput {
	return i.ToVirtualHardDiskOutputWithContext(context.Background())
}

func (i VirtualHardDiskArgs) ToVirtualHardDiskOutputWithContext(ctx context.Context) VirtualHardDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHardDiskOutput)
}

func (i VirtualHardDiskArgs) ToVirtualHardDiskPtrOutput() VirtualHardDiskPtrOutput {
	return i.ToVirtualHardDiskPtrOutputWithContext(context.Background())
}

func (i VirtualHardDiskArgs) ToVirtualHardDiskPtrOutputWithContext(ctx context.Context) VirtualHardDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHardDiskOutput).ToVirtualHardDiskPtrOutputWithContext(ctx)
}









type VirtualHardDiskPtrInput interface {
	pulumi.Input

	ToVirtualHardDiskPtrOutput() VirtualHardDiskPtrOutput
	ToVirtualHardDiskPtrOutputWithContext(context.Context) VirtualHardDiskPtrOutput
}

type virtualHardDiskPtrType VirtualHardDiskArgs

func VirtualHardDiskPtr(v *VirtualHardDiskArgs) VirtualHardDiskPtrInput {
	return (*virtualHardDiskPtrType)(v)
}

func (*virtualHardDiskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHardDisk)(nil)).Elem()
}

func (i *virtualHardDiskPtrType) ToVirtualHardDiskPtrOutput() VirtualHardDiskPtrOutput {
	return i.ToVirtualHardDiskPtrOutputWithContext(context.Background())
}

func (i *virtualHardDiskPtrType) ToVirtualHardDiskPtrOutputWithContext(ctx context.Context) VirtualHardDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHardDiskPtrOutput)
}

type VirtualHardDiskOutput struct{ *pulumi.OutputState }

func (VirtualHardDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHardDisk)(nil)).Elem()
}

func (o VirtualHardDiskOutput) ToVirtualHardDiskOutput() VirtualHardDiskOutput {
	return o
}

func (o VirtualHardDiskOutput) ToVirtualHardDiskOutputWithContext(ctx context.Context) VirtualHardDiskOutput {
	return o
}

func (o VirtualHardDiskOutput) ToVirtualHardDiskPtrOutput() VirtualHardDiskPtrOutput {
	return o.ToVirtualHardDiskPtrOutputWithContext(context.Background())
}

func (o VirtualHardDiskOutput) ToVirtualHardDiskPtrOutputWithContext(ctx context.Context) VirtualHardDiskPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualHardDisk) *VirtualHardDisk {
		return &v
	}).(VirtualHardDiskPtrOutput)
}

func (o VirtualHardDiskOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHardDisk) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type VirtualHardDiskPtrOutput struct{ *pulumi.OutputState }

func (VirtualHardDiskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHardDisk)(nil)).Elem()
}

func (o VirtualHardDiskPtrOutput) ToVirtualHardDiskPtrOutput() VirtualHardDiskPtrOutput {
	return o
}

func (o VirtualHardDiskPtrOutput) ToVirtualHardDiskPtrOutputWithContext(ctx context.Context) VirtualHardDiskPtrOutput {
	return o
}

func (o VirtualHardDiskPtrOutput) Elem() VirtualHardDiskOutput {
	return o.ApplyT(func(v *VirtualHardDisk) VirtualHardDisk {
		if v != nil {
			return *v
		}
		var ret VirtualHardDisk
		return ret
	}).(VirtualHardDiskOutput)
}

func (o VirtualHardDiskPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHardDisk) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type VirtualHardDiskResponse struct {
	Uri *string `pulumi:"uri"`
}

type VirtualHardDiskResponseOutput struct{ *pulumi.OutputState }

func (VirtualHardDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHardDiskResponse)(nil)).Elem()
}

func (o VirtualHardDiskResponseOutput) ToVirtualHardDiskResponseOutput() VirtualHardDiskResponseOutput {
	return o
}

func (o VirtualHardDiskResponseOutput) ToVirtualHardDiskResponseOutputWithContext(ctx context.Context) VirtualHardDiskResponseOutput {
	return o
}

func (o VirtualHardDiskResponseOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHardDiskResponse) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type VirtualHardDiskResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualHardDiskResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHardDiskResponse)(nil)).Elem()
}

func (o VirtualHardDiskResponsePtrOutput) ToVirtualHardDiskResponsePtrOutput() VirtualHardDiskResponsePtrOutput {
	return o
}

func (o VirtualHardDiskResponsePtrOutput) ToVirtualHardDiskResponsePtrOutputWithContext(ctx context.Context) VirtualHardDiskResponsePtrOutput {
	return o
}

func (o VirtualHardDiskResponsePtrOutput) Elem() VirtualHardDiskResponseOutput {
	return o.ApplyT(func(v *VirtualHardDiskResponse) VirtualHardDiskResponse {
		if v != nil {
			return *v
		}
		var ret VirtualHardDiskResponse
		return ret
	}).(VirtualHardDiskResponseOutput)
}

func (o VirtualHardDiskResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHardDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineAgentInstanceViewResponse struct {
	ExtensionHandlers []VirtualMachineExtensionHandlerInstanceViewResponse `pulumi:"extensionHandlers"`
	Statuses          []InstanceViewStatusResponse                         `pulumi:"statuses"`
	VmAgentVersion    *string                                              `pulumi:"vmAgentVersion"`
}

type VirtualMachineAgentInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineAgentInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineAgentInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineAgentInstanceViewResponseOutput) ToVirtualMachineAgentInstanceViewResponseOutput() VirtualMachineAgentInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineAgentInstanceViewResponseOutput) ToVirtualMachineAgentInstanceViewResponseOutputWithContext(ctx context.Context) VirtualMachineAgentInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineAgentInstanceViewResponseOutput) ExtensionHandlers() VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineAgentInstanceViewResponse) []VirtualMachineExtensionHandlerInstanceViewResponse {
		return v.ExtensionHandlers
	}).(VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput)
}

func (o VirtualMachineAgentInstanceViewResponseOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineAgentInstanceViewResponse) []InstanceViewStatusResponse { return v.Statuses }).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachineAgentInstanceViewResponseOutput) VmAgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineAgentInstanceViewResponse) *string { return v.VmAgentVersion }).(pulumi.StringPtrOutput)
}

type VirtualMachineAgentInstanceViewResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineAgentInstanceViewResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineAgentInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineAgentInstanceViewResponsePtrOutput) ToVirtualMachineAgentInstanceViewResponsePtrOutput() VirtualMachineAgentInstanceViewResponsePtrOutput {
	return o
}

func (o VirtualMachineAgentInstanceViewResponsePtrOutput) ToVirtualMachineAgentInstanceViewResponsePtrOutputWithContext(ctx context.Context) VirtualMachineAgentInstanceViewResponsePtrOutput {
	return o
}

func (o VirtualMachineAgentInstanceViewResponsePtrOutput) Elem() VirtualMachineAgentInstanceViewResponseOutput {
	return o.ApplyT(func(v *VirtualMachineAgentInstanceViewResponse) VirtualMachineAgentInstanceViewResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineAgentInstanceViewResponse
		return ret
	}).(VirtualMachineAgentInstanceViewResponseOutput)
}

func (o VirtualMachineAgentInstanceViewResponsePtrOutput) ExtensionHandlers() VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineAgentInstanceViewResponse) []VirtualMachineExtensionHandlerInstanceViewResponse {
		if v == nil {
			return nil
		}
		return v.ExtensionHandlers
	}).(VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput)
}

func (o VirtualMachineAgentInstanceViewResponsePtrOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineAgentInstanceViewResponse) []InstanceViewStatusResponse {
		if v == nil {
			return nil
		}
		return v.Statuses
	}).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachineAgentInstanceViewResponsePtrOutput) VmAgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineAgentInstanceViewResponse) *string {
		if v == nil {
			return nil
		}
		return v.VmAgentVersion
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineExtensionHandlerInstanceViewResponse struct {
	Status             *InstanceViewStatusResponse `pulumi:"status"`
	Type               *string                     `pulumi:"type"`
	TypeHandlerVersion *string                     `pulumi:"typeHandlerVersion"`
}

type VirtualMachineExtensionHandlerInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionHandlerInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineExtensionHandlerInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineExtensionHandlerInstanceViewResponseOutput) ToVirtualMachineExtensionHandlerInstanceViewResponseOutput() VirtualMachineExtensionHandlerInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineExtensionHandlerInstanceViewResponseOutput) ToVirtualMachineExtensionHandlerInstanceViewResponseOutputWithContext(ctx context.Context) VirtualMachineExtensionHandlerInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineExtensionHandlerInstanceViewResponseOutput) Status() InstanceViewStatusResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionHandlerInstanceViewResponse) *InstanceViewStatusResponse {
		return v.Status
	}).(InstanceViewStatusResponsePtrOutput)
}

func (o VirtualMachineExtensionHandlerInstanceViewResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionHandlerInstanceViewResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionHandlerInstanceViewResponseOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionHandlerInstanceViewResponse) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineExtensionHandlerInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput) ToVirtualMachineExtensionHandlerInstanceViewResponseArrayOutput() VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput {
	return o
}

func (o VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput) ToVirtualMachineExtensionHandlerInstanceViewResponseArrayOutputWithContext(ctx context.Context) VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput {
	return o
}

func (o VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineExtensionHandlerInstanceViewResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineExtensionHandlerInstanceViewResponse {
		return vs[0].([]VirtualMachineExtensionHandlerInstanceViewResponse)[vs[1].(int)]
	}).(VirtualMachineExtensionHandlerInstanceViewResponseOutput)
}

type VirtualMachineExtensionInstanceView struct {
	Name               *string              `pulumi:"name"`
	Statuses           []InstanceViewStatus `pulumi:"statuses"`
	Substatuses        []InstanceViewStatus `pulumi:"substatuses"`
	Type               *string              `pulumi:"type"`
	TypeHandlerVersion *string              `pulumi:"typeHandlerVersion"`
}





type VirtualMachineExtensionInstanceViewInput interface {
	pulumi.Input

	ToVirtualMachineExtensionInstanceViewOutput() VirtualMachineExtensionInstanceViewOutput
	ToVirtualMachineExtensionInstanceViewOutputWithContext(context.Context) VirtualMachineExtensionInstanceViewOutput
}

type VirtualMachineExtensionInstanceViewArgs struct {
	Name               pulumi.StringPtrInput        `pulumi:"name"`
	Statuses           InstanceViewStatusArrayInput `pulumi:"statuses"`
	Substatuses        InstanceViewStatusArrayInput `pulumi:"substatuses"`
	Type               pulumi.StringPtrInput        `pulumi:"type"`
	TypeHandlerVersion pulumi.StringPtrInput        `pulumi:"typeHandlerVersion"`
}

func (VirtualMachineExtensionInstanceViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineExtensionInstanceView)(nil)).Elem()
}

func (i VirtualMachineExtensionInstanceViewArgs) ToVirtualMachineExtensionInstanceViewOutput() VirtualMachineExtensionInstanceViewOutput {
	return i.ToVirtualMachineExtensionInstanceViewOutputWithContext(context.Background())
}

func (i VirtualMachineExtensionInstanceViewArgs) ToVirtualMachineExtensionInstanceViewOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineExtensionInstanceViewOutput)
}

func (i VirtualMachineExtensionInstanceViewArgs) ToVirtualMachineExtensionInstanceViewPtrOutput() VirtualMachineExtensionInstanceViewPtrOutput {
	return i.ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(context.Background())
}

func (i VirtualMachineExtensionInstanceViewArgs) ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineExtensionInstanceViewOutput).ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(ctx)
}









type VirtualMachineExtensionInstanceViewPtrInput interface {
	pulumi.Input

	ToVirtualMachineExtensionInstanceViewPtrOutput() VirtualMachineExtensionInstanceViewPtrOutput
	ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(context.Context) VirtualMachineExtensionInstanceViewPtrOutput
}

type virtualMachineExtensionInstanceViewPtrType VirtualMachineExtensionInstanceViewArgs

func VirtualMachineExtensionInstanceViewPtr(v *VirtualMachineExtensionInstanceViewArgs) VirtualMachineExtensionInstanceViewPtrInput {
	return (*virtualMachineExtensionInstanceViewPtrType)(v)
}

func (*virtualMachineExtensionInstanceViewPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineExtensionInstanceView)(nil)).Elem()
}

func (i *virtualMachineExtensionInstanceViewPtrType) ToVirtualMachineExtensionInstanceViewPtrOutput() VirtualMachineExtensionInstanceViewPtrOutput {
	return i.ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(context.Background())
}

func (i *virtualMachineExtensionInstanceViewPtrType) ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineExtensionInstanceViewPtrOutput)
}

type VirtualMachineExtensionInstanceViewOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionInstanceViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineExtensionInstanceView)(nil)).Elem()
}

func (o VirtualMachineExtensionInstanceViewOutput) ToVirtualMachineExtensionInstanceViewOutput() VirtualMachineExtensionInstanceViewOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewOutput) ToVirtualMachineExtensionInstanceViewOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewOutput) ToVirtualMachineExtensionInstanceViewPtrOutput() VirtualMachineExtensionInstanceViewPtrOutput {
	return o.ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(context.Background())
}

func (o VirtualMachineExtensionInstanceViewOutput) ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineExtensionInstanceView) *VirtualMachineExtensionInstanceView {
		return &v
	}).(VirtualMachineExtensionInstanceViewPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceView) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewOutput) Statuses() InstanceViewStatusArrayOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceView) []InstanceViewStatus { return v.Statuses }).(InstanceViewStatusArrayOutput)
}

func (o VirtualMachineExtensionInstanceViewOutput) Substatuses() InstanceViewStatusArrayOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceView) []InstanceViewStatus { return v.Substatuses }).(InstanceViewStatusArrayOutput)
}

func (o VirtualMachineExtensionInstanceViewOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceView) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceView) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type VirtualMachineExtensionInstanceViewPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionInstanceViewPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineExtensionInstanceView)(nil)).Elem()
}

func (o VirtualMachineExtensionInstanceViewPtrOutput) ToVirtualMachineExtensionInstanceViewPtrOutput() VirtualMachineExtensionInstanceViewPtrOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewPtrOutput) ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewPtrOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewPtrOutput) Elem() VirtualMachineExtensionInstanceViewOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceView) VirtualMachineExtensionInstanceView {
		if v != nil {
			return *v
		}
		var ret VirtualMachineExtensionInstanceView
		return ret
	}).(VirtualMachineExtensionInstanceViewOutput)
}

func (o VirtualMachineExtensionInstanceViewPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceView) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewPtrOutput) Statuses() InstanceViewStatusArrayOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceView) []InstanceViewStatus {
		if v == nil {
			return nil
		}
		return v.Statuses
	}).(InstanceViewStatusArrayOutput)
}

func (o VirtualMachineExtensionInstanceViewPtrOutput) Substatuses() InstanceViewStatusArrayOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceView) []InstanceViewStatus {
		if v == nil {
			return nil
		}
		return v.Substatuses
	}).(InstanceViewStatusArrayOutput)
}

func (o VirtualMachineExtensionInstanceViewPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceView) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewPtrOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceView) *string {
		if v == nil {
			return nil
		}
		return v.TypeHandlerVersion
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineExtensionInstanceViewResponse struct {
	Name               *string                      `pulumi:"name"`
	Statuses           []InstanceViewStatusResponse `pulumi:"statuses"`
	Substatuses        []InstanceViewStatusResponse `pulumi:"substatuses"`
	Type               *string                      `pulumi:"type"`
	TypeHandlerVersion *string                      `pulumi:"typeHandlerVersion"`
}

type VirtualMachineExtensionInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineExtensionInstanceViewResponseOutput) ToVirtualMachineExtensionInstanceViewResponseOutput() VirtualMachineExtensionInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewResponseOutput) ToVirtualMachineExtensionInstanceViewResponseOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceViewResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewResponseOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceViewResponse) []InstanceViewStatusResponse { return v.Statuses }).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachineExtensionInstanceViewResponseOutput) Substatuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceViewResponse) []InstanceViewStatusResponse { return v.Substatuses }).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachineExtensionInstanceViewResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceViewResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewResponseOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceViewResponse) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type VirtualMachineExtensionInstanceViewResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionInstanceViewResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineExtensionInstanceViewResponsePtrOutput) ToVirtualMachineExtensionInstanceViewResponsePtrOutput() VirtualMachineExtensionInstanceViewResponsePtrOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewResponsePtrOutput) ToVirtualMachineExtensionInstanceViewResponsePtrOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewResponsePtrOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewResponsePtrOutput) Elem() VirtualMachineExtensionInstanceViewResponseOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceViewResponse) VirtualMachineExtensionInstanceViewResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineExtensionInstanceViewResponse
		return ret
	}).(VirtualMachineExtensionInstanceViewResponseOutput)
}

func (o VirtualMachineExtensionInstanceViewResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceViewResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewResponsePtrOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceViewResponse) []InstanceViewStatusResponse {
		if v == nil {
			return nil
		}
		return v.Statuses
	}).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachineExtensionInstanceViewResponsePtrOutput) Substatuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceViewResponse) []InstanceViewStatusResponse {
		if v == nil {
			return nil
		}
		return v.Substatuses
	}).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachineExtensionInstanceViewResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceViewResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewResponsePtrOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceViewResponse) *string {
		if v == nil {
			return nil
		}
		return v.TypeHandlerVersion
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineExtensionInstanceViewResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionInstanceViewResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineExtensionInstanceViewResponseArrayOutput) ToVirtualMachineExtensionInstanceViewResponseArrayOutput() VirtualMachineExtensionInstanceViewResponseArrayOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewResponseArrayOutput) ToVirtualMachineExtensionInstanceViewResponseArrayOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewResponseArrayOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineExtensionInstanceViewResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineExtensionInstanceViewResponse {
		return vs[0].([]VirtualMachineExtensionInstanceViewResponse)[vs[1].(int)]
	}).(VirtualMachineExtensionInstanceViewResponseOutput)
}

type VirtualMachineExtensionResponse struct {
	AutoUpgradeMinorVersion *bool                                        `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag          *string                                      `pulumi:"forceUpdateTag"`
	Id                      string                                       `pulumi:"id"`
	InstanceView            *VirtualMachineExtensionInstanceViewResponse `pulumi:"instanceView"`
	Location                string                                       `pulumi:"location"`
	Name                    string                                       `pulumi:"name"`
	ProtectedSettings       interface{}                                  `pulumi:"protectedSettings"`
	ProvisioningState       string                                       `pulumi:"provisioningState"`
	Publisher               *string                                      `pulumi:"publisher"`
	Settings                interface{}                                  `pulumi:"settings"`
	Tags                    map[string]string                            `pulumi:"tags"`
	Type                    string                                       `pulumi:"type"`
	TypeHandlerVersion      *string                                      `pulumi:"typeHandlerVersion"`
}

type VirtualMachineExtensionResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineExtensionResponse)(nil)).Elem()
}

func (o VirtualMachineExtensionResponseOutput) ToVirtualMachineExtensionResponseOutput() VirtualMachineExtensionResponseOutput {
	return o
}

func (o VirtualMachineExtensionResponseOutput) ToVirtualMachineExtensionResponseOutputWithContext(ctx context.Context) VirtualMachineExtensionResponseOutput {
	return o
}

func (o VirtualMachineExtensionResponseOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineExtensionResponseOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o VirtualMachineExtensionResponseOutput) InstanceView() VirtualMachineExtensionInstanceViewResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) *VirtualMachineExtensionInstanceViewResponse {
		return v.InstanceView
	}).(VirtualMachineExtensionInstanceViewResponsePtrOutput)
}

func (o VirtualMachineExtensionResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o VirtualMachineExtensionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineExtensionResponseOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o VirtualMachineExtensionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualMachineExtensionResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionResponseOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

func (o VirtualMachineExtensionResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualMachineExtensionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualMachineExtensionResponseOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type VirtualMachineExtensionResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineExtensionResponse)(nil)).Elem()
}

func (o VirtualMachineExtensionResponseArrayOutput) ToVirtualMachineExtensionResponseArrayOutput() VirtualMachineExtensionResponseArrayOutput {
	return o
}

func (o VirtualMachineExtensionResponseArrayOutput) ToVirtualMachineExtensionResponseArrayOutputWithContext(ctx context.Context) VirtualMachineExtensionResponseArrayOutput {
	return o
}

func (o VirtualMachineExtensionResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineExtensionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineExtensionResponse {
		return vs[0].([]VirtualMachineExtensionResponse)[vs[1].(int)]
	}).(VirtualMachineExtensionResponseOutput)
}

type VirtualMachineInstanceViewResponse struct {
	BootDiagnostics      *BootDiagnosticsInstanceViewResponse          `pulumi:"bootDiagnostics"`
	Disks                []DiskInstanceViewResponse                    `pulumi:"disks"`
	Extensions           []VirtualMachineExtensionInstanceViewResponse `pulumi:"extensions"`
	PlatformFaultDomain  *int                                          `pulumi:"platformFaultDomain"`
	PlatformUpdateDomain *int                                          `pulumi:"platformUpdateDomain"`
	RdpThumbPrint        *string                                       `pulumi:"rdpThumbPrint"`
	Statuses             []InstanceViewStatusResponse                  `pulumi:"statuses"`
	VmAgent              *VirtualMachineAgentInstanceViewResponse      `pulumi:"vmAgent"`
}

type VirtualMachineInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineInstanceViewResponseOutput) ToVirtualMachineInstanceViewResponseOutput() VirtualMachineInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineInstanceViewResponseOutput) ToVirtualMachineInstanceViewResponseOutputWithContext(ctx context.Context) VirtualMachineInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineInstanceViewResponseOutput) BootDiagnostics() BootDiagnosticsInstanceViewResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *BootDiagnosticsInstanceViewResponse {
		return v.BootDiagnostics
	}).(BootDiagnosticsInstanceViewResponsePtrOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) Disks() DiskInstanceViewResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) []DiskInstanceViewResponse { return v.Disks }).(DiskInstanceViewResponseArrayOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) Extensions() VirtualMachineExtensionInstanceViewResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) []VirtualMachineExtensionInstanceViewResponse {
		return v.Extensions
	}).(VirtualMachineExtensionInstanceViewResponseArrayOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) PlatformFaultDomain() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *int { return v.PlatformFaultDomain }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) PlatformUpdateDomain() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *int { return v.PlatformUpdateDomain }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) RdpThumbPrint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *string { return v.RdpThumbPrint }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) []InstanceViewStatusResponse { return v.Statuses }).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) VmAgent() VirtualMachineAgentInstanceViewResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *VirtualMachineAgentInstanceViewResponse { return v.VmAgent }).(VirtualMachineAgentInstanceViewResponsePtrOutput)
}

type VirtualMachineScaleSetExtension struct {
	AutoUpgradeMinorVersion *bool       `pulumi:"autoUpgradeMinorVersion"`
	Id                      *string     `pulumi:"id"`
	Name                    *string     `pulumi:"name"`
	ProtectedSettings       interface{} `pulumi:"protectedSettings"`
	Publisher               *string     `pulumi:"publisher"`
	Settings                interface{} `pulumi:"settings"`
	Type                    *string     `pulumi:"type"`
	TypeHandlerVersion      *string     `pulumi:"typeHandlerVersion"`
}





type VirtualMachineScaleSetExtensionInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetExtensionOutput() VirtualMachineScaleSetExtensionOutput
	ToVirtualMachineScaleSetExtensionOutputWithContext(context.Context) VirtualMachineScaleSetExtensionOutput
}

type VirtualMachineScaleSetExtensionArgs struct {
	AutoUpgradeMinorVersion pulumi.BoolPtrInput   `pulumi:"autoUpgradeMinorVersion"`
	Id                      pulumi.StringPtrInput `pulumi:"id"`
	Name                    pulumi.StringPtrInput `pulumi:"name"`
	ProtectedSettings       pulumi.Input          `pulumi:"protectedSettings"`
	Publisher               pulumi.StringPtrInput `pulumi:"publisher"`
	Settings                pulumi.Input          `pulumi:"settings"`
	Type                    pulumi.StringPtrInput `pulumi:"type"`
	TypeHandlerVersion      pulumi.StringPtrInput `pulumi:"typeHandlerVersion"`
}

func (VirtualMachineScaleSetExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetExtension)(nil)).Elem()
}

func (i VirtualMachineScaleSetExtensionArgs) ToVirtualMachineScaleSetExtensionOutput() VirtualMachineScaleSetExtensionOutput {
	return i.ToVirtualMachineScaleSetExtensionOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetExtensionArgs) ToVirtualMachineScaleSetExtensionOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetExtensionOutput)
}





type VirtualMachineScaleSetExtensionArrayInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetExtensionArrayOutput() VirtualMachineScaleSetExtensionArrayOutput
	ToVirtualMachineScaleSetExtensionArrayOutputWithContext(context.Context) VirtualMachineScaleSetExtensionArrayOutput
}

type VirtualMachineScaleSetExtensionArray []VirtualMachineScaleSetExtensionInput

func (VirtualMachineScaleSetExtensionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetExtension)(nil)).Elem()
}

func (i VirtualMachineScaleSetExtensionArray) ToVirtualMachineScaleSetExtensionArrayOutput() VirtualMachineScaleSetExtensionArrayOutput {
	return i.ToVirtualMachineScaleSetExtensionArrayOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetExtensionArray) ToVirtualMachineScaleSetExtensionArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetExtensionArrayOutput)
}

type VirtualMachineScaleSetExtensionOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetExtension)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionOutput) ToVirtualMachineScaleSetExtensionOutput() VirtualMachineScaleSetExtensionOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionOutput) ToVirtualMachineScaleSetExtensionOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtension) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetExtensionOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtension) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtension) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtension) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o VirtualMachineScaleSetExtensionOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtension) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtension) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

func (o VirtualMachineScaleSetExtensionOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtension) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtension) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetExtensionArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetExtension)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionArrayOutput) ToVirtualMachineScaleSetExtensionArrayOutput() VirtualMachineScaleSetExtensionArrayOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionArrayOutput) ToVirtualMachineScaleSetExtensionArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionArrayOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetExtensionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetExtension {
		return vs[0].([]VirtualMachineScaleSetExtension)[vs[1].(int)]
	}).(VirtualMachineScaleSetExtensionOutput)
}

type VirtualMachineScaleSetExtensionProfile struct {
	Extensions []VirtualMachineScaleSetExtension `pulumi:"extensions"`
}





type VirtualMachineScaleSetExtensionProfileInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetExtensionProfileOutput() VirtualMachineScaleSetExtensionProfileOutput
	ToVirtualMachineScaleSetExtensionProfileOutputWithContext(context.Context) VirtualMachineScaleSetExtensionProfileOutput
}

type VirtualMachineScaleSetExtensionProfileArgs struct {
	Extensions VirtualMachineScaleSetExtensionArrayInput `pulumi:"extensions"`
}

func (VirtualMachineScaleSetExtensionProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetExtensionProfile)(nil)).Elem()
}

func (i VirtualMachineScaleSetExtensionProfileArgs) ToVirtualMachineScaleSetExtensionProfileOutput() VirtualMachineScaleSetExtensionProfileOutput {
	return i.ToVirtualMachineScaleSetExtensionProfileOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetExtensionProfileArgs) ToVirtualMachineScaleSetExtensionProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetExtensionProfileOutput)
}

func (i VirtualMachineScaleSetExtensionProfileArgs) ToVirtualMachineScaleSetExtensionProfilePtrOutput() VirtualMachineScaleSetExtensionProfilePtrOutput {
	return i.ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetExtensionProfileArgs) ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetExtensionProfileOutput).ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetExtensionProfilePtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetExtensionProfilePtrOutput() VirtualMachineScaleSetExtensionProfilePtrOutput
	ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(context.Context) VirtualMachineScaleSetExtensionProfilePtrOutput
}

type virtualMachineScaleSetExtensionProfilePtrType VirtualMachineScaleSetExtensionProfileArgs

func VirtualMachineScaleSetExtensionProfilePtr(v *VirtualMachineScaleSetExtensionProfileArgs) VirtualMachineScaleSetExtensionProfilePtrInput {
	return (*virtualMachineScaleSetExtensionProfilePtrType)(v)
}

func (*virtualMachineScaleSetExtensionProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetExtensionProfile)(nil)).Elem()
}

func (i *virtualMachineScaleSetExtensionProfilePtrType) ToVirtualMachineScaleSetExtensionProfilePtrOutput() VirtualMachineScaleSetExtensionProfilePtrOutput {
	return i.ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetExtensionProfilePtrType) ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetExtensionProfilePtrOutput)
}

type VirtualMachineScaleSetExtensionProfileOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetExtensionProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionProfileOutput) ToVirtualMachineScaleSetExtensionProfileOutput() VirtualMachineScaleSetExtensionProfileOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionProfileOutput) ToVirtualMachineScaleSetExtensionProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionProfileOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionProfileOutput) ToVirtualMachineScaleSetExtensionProfilePtrOutput() VirtualMachineScaleSetExtensionProfilePtrOutput {
	return o.ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetExtensionProfileOutput) ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetExtensionProfile) *VirtualMachineScaleSetExtensionProfile {
		return &v
	}).(VirtualMachineScaleSetExtensionProfilePtrOutput)
}

func (o VirtualMachineScaleSetExtensionProfileOutput) Extensions() VirtualMachineScaleSetExtensionArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionProfile) []VirtualMachineScaleSetExtension { return v.Extensions }).(VirtualMachineScaleSetExtensionArrayOutput)
}

type VirtualMachineScaleSetExtensionProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetExtensionProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionProfilePtrOutput) ToVirtualMachineScaleSetExtensionProfilePtrOutput() VirtualMachineScaleSetExtensionProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionProfilePtrOutput) ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionProfilePtrOutput) Elem() VirtualMachineScaleSetExtensionProfileOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtensionProfile) VirtualMachineScaleSetExtensionProfile {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetExtensionProfile
		return ret
	}).(VirtualMachineScaleSetExtensionProfileOutput)
}

func (o VirtualMachineScaleSetExtensionProfilePtrOutput) Extensions() VirtualMachineScaleSetExtensionArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtensionProfile) []VirtualMachineScaleSetExtension {
		if v == nil {
			return nil
		}
		return v.Extensions
	}).(VirtualMachineScaleSetExtensionArrayOutput)
}

type VirtualMachineScaleSetExtensionProfileResponse struct {
	Extensions []VirtualMachineScaleSetExtensionResponse `pulumi:"extensions"`
}

type VirtualMachineScaleSetExtensionProfileResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetExtensionProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionProfileResponseOutput) ToVirtualMachineScaleSetExtensionProfileResponseOutput() VirtualMachineScaleSetExtensionProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionProfileResponseOutput) ToVirtualMachineScaleSetExtensionProfileResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionProfileResponseOutput) Extensions() VirtualMachineScaleSetExtensionResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionProfileResponse) []VirtualMachineScaleSetExtensionResponse {
		return v.Extensions
	}).(VirtualMachineScaleSetExtensionResponseArrayOutput)
}

type VirtualMachineScaleSetExtensionProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetExtensionProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionProfileResponsePtrOutput) ToVirtualMachineScaleSetExtensionProfileResponsePtrOutput() VirtualMachineScaleSetExtensionProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionProfileResponsePtrOutput) ToVirtualMachineScaleSetExtensionProfileResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionProfileResponsePtrOutput) Elem() VirtualMachineScaleSetExtensionProfileResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtensionProfileResponse) VirtualMachineScaleSetExtensionProfileResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetExtensionProfileResponse
		return ret
	}).(VirtualMachineScaleSetExtensionProfileResponseOutput)
}

func (o VirtualMachineScaleSetExtensionProfileResponsePtrOutput) Extensions() VirtualMachineScaleSetExtensionResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtensionProfileResponse) []VirtualMachineScaleSetExtensionResponse {
		if v == nil {
			return nil
		}
		return v.Extensions
	}).(VirtualMachineScaleSetExtensionResponseArrayOutput)
}

type VirtualMachineScaleSetExtensionResponse struct {
	AutoUpgradeMinorVersion *bool       `pulumi:"autoUpgradeMinorVersion"`
	Id                      *string     `pulumi:"id"`
	Name                    *string     `pulumi:"name"`
	ProtectedSettings       interface{} `pulumi:"protectedSettings"`
	ProvisioningState       string      `pulumi:"provisioningState"`
	Publisher               *string     `pulumi:"publisher"`
	Settings                interface{} `pulumi:"settings"`
	Type                    *string     `pulumi:"type"`
	TypeHandlerVersion      *string     `pulumi:"typeHandlerVersion"`
}

type VirtualMachineScaleSetExtensionResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetExtensionResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionResponseOutput) ToVirtualMachineScaleSetExtensionResponseOutput() VirtualMachineScaleSetExtensionResponseOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionResponseOutput) ToVirtualMachineScaleSetExtensionResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionResponseOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionResponseOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetExtensionResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetExtensionResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionResponseArrayOutput) ToVirtualMachineScaleSetExtensionResponseArrayOutput() VirtualMachineScaleSetExtensionResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionResponseArrayOutput) ToVirtualMachineScaleSetExtensionResponseArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetExtensionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetExtensionResponse {
		return vs[0].([]VirtualMachineScaleSetExtensionResponse)[vs[1].(int)]
	}).(VirtualMachineScaleSetExtensionResponseOutput)
}

type VirtualMachineScaleSetIPConfiguration struct {
	Id                              *string            `pulumi:"id"`
	LoadBalancerBackendAddressPools []SubResource      `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerInboundNatPools     []SubResource      `pulumi:"loadBalancerInboundNatPools"`
	Name                            string             `pulumi:"name"`
	Subnet                          ApiEntityReference `pulumi:"subnet"`
}





type VirtualMachineScaleSetIPConfigurationInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetIPConfigurationOutput() VirtualMachineScaleSetIPConfigurationOutput
	ToVirtualMachineScaleSetIPConfigurationOutputWithContext(context.Context) VirtualMachineScaleSetIPConfigurationOutput
}

type VirtualMachineScaleSetIPConfigurationArgs struct {
	Id                              pulumi.StringPtrInput   `pulumi:"id"`
	LoadBalancerBackendAddressPools SubResourceArrayInput   `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerInboundNatPools     SubResourceArrayInput   `pulumi:"loadBalancerInboundNatPools"`
	Name                            pulumi.StringInput      `pulumi:"name"`
	Subnet                          ApiEntityReferenceInput `pulumi:"subnet"`
}

func (VirtualMachineScaleSetIPConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetIPConfiguration)(nil)).Elem()
}

func (i VirtualMachineScaleSetIPConfigurationArgs) ToVirtualMachineScaleSetIPConfigurationOutput() VirtualMachineScaleSetIPConfigurationOutput {
	return i.ToVirtualMachineScaleSetIPConfigurationOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetIPConfigurationArgs) ToVirtualMachineScaleSetIPConfigurationOutputWithContext(ctx context.Context) VirtualMachineScaleSetIPConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetIPConfigurationOutput)
}





type VirtualMachineScaleSetIPConfigurationArrayInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetIPConfigurationArrayOutput() VirtualMachineScaleSetIPConfigurationArrayOutput
	ToVirtualMachineScaleSetIPConfigurationArrayOutputWithContext(context.Context) VirtualMachineScaleSetIPConfigurationArrayOutput
}

type VirtualMachineScaleSetIPConfigurationArray []VirtualMachineScaleSetIPConfigurationInput

func (VirtualMachineScaleSetIPConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetIPConfiguration)(nil)).Elem()
}

func (i VirtualMachineScaleSetIPConfigurationArray) ToVirtualMachineScaleSetIPConfigurationArrayOutput() VirtualMachineScaleSetIPConfigurationArrayOutput {
	return i.ToVirtualMachineScaleSetIPConfigurationArrayOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetIPConfigurationArray) ToVirtualMachineScaleSetIPConfigurationArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetIPConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetIPConfigurationArrayOutput)
}

type VirtualMachineScaleSetIPConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIPConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetIPConfiguration)(nil)).Elem()
}

func (o VirtualMachineScaleSetIPConfigurationOutput) ToVirtualMachineScaleSetIPConfigurationOutput() VirtualMachineScaleSetIPConfigurationOutput {
	return o
}

func (o VirtualMachineScaleSetIPConfigurationOutput) ToVirtualMachineScaleSetIPConfigurationOutputWithContext(ctx context.Context) VirtualMachineScaleSetIPConfigurationOutput {
	return o
}

func (o VirtualMachineScaleSetIPConfigurationOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfiguration) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetIPConfigurationOutput) LoadBalancerBackendAddressPools() SubResourceArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfiguration) []SubResource { return v.LoadBalancerBackendAddressPools }).(SubResourceArrayOutput)
}

func (o VirtualMachineScaleSetIPConfigurationOutput) LoadBalancerInboundNatPools() SubResourceArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfiguration) []SubResource { return v.LoadBalancerInboundNatPools }).(SubResourceArrayOutput)
}

func (o VirtualMachineScaleSetIPConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfiguration) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetIPConfigurationOutput) Subnet() ApiEntityReferenceOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfiguration) ApiEntityReference { return v.Subnet }).(ApiEntityReferenceOutput)
}

type VirtualMachineScaleSetIPConfigurationArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIPConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetIPConfiguration)(nil)).Elem()
}

func (o VirtualMachineScaleSetIPConfigurationArrayOutput) ToVirtualMachineScaleSetIPConfigurationArrayOutput() VirtualMachineScaleSetIPConfigurationArrayOutput {
	return o
}

func (o VirtualMachineScaleSetIPConfigurationArrayOutput) ToVirtualMachineScaleSetIPConfigurationArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetIPConfigurationArrayOutput {
	return o
}

func (o VirtualMachineScaleSetIPConfigurationArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetIPConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetIPConfiguration {
		return vs[0].([]VirtualMachineScaleSetIPConfiguration)[vs[1].(int)]
	}).(VirtualMachineScaleSetIPConfigurationOutput)
}

type VirtualMachineScaleSetIPConfigurationResponse struct {
	Id                              *string                    `pulumi:"id"`
	LoadBalancerBackendAddressPools []SubResourceResponse      `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerInboundNatPools     []SubResourceResponse      `pulumi:"loadBalancerInboundNatPools"`
	Name                            string                     `pulumi:"name"`
	Subnet                          ApiEntityReferenceResponse `pulumi:"subnet"`
}

type VirtualMachineScaleSetIPConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIPConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetIPConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) ToVirtualMachineScaleSetIPConfigurationResponseOutput() VirtualMachineScaleSetIPConfigurationResponseOutput {
	return o
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) ToVirtualMachineScaleSetIPConfigurationResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetIPConfigurationResponseOutput {
	return o
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) LoadBalancerBackendAddressPools() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfigurationResponse) []SubResourceResponse {
		return v.LoadBalancerBackendAddressPools
	}).(SubResourceResponseArrayOutput)
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) LoadBalancerInboundNatPools() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfigurationResponse) []SubResourceResponse {
		return v.LoadBalancerInboundNatPools
	}).(SubResourceResponseArrayOutput)
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfigurationResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) Subnet() ApiEntityReferenceResponseOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfigurationResponse) ApiEntityReferenceResponse { return v.Subnet }).(ApiEntityReferenceResponseOutput)
}

type VirtualMachineScaleSetIPConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIPConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetIPConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetIPConfigurationResponseArrayOutput) ToVirtualMachineScaleSetIPConfigurationResponseArrayOutput() VirtualMachineScaleSetIPConfigurationResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetIPConfigurationResponseArrayOutput) ToVirtualMachineScaleSetIPConfigurationResponseArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetIPConfigurationResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetIPConfigurationResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetIPConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetIPConfigurationResponse {
		return vs[0].([]VirtualMachineScaleSetIPConfigurationResponse)[vs[1].(int)]
	}).(VirtualMachineScaleSetIPConfigurationResponseOutput)
}

type VirtualMachineScaleSetNetworkConfiguration struct {
	Id               *string                                 `pulumi:"id"`
	IpConfigurations []VirtualMachineScaleSetIPConfiguration `pulumi:"ipConfigurations"`
	Name             string                                  `pulumi:"name"`
	Primary          *bool                                   `pulumi:"primary"`
}





type VirtualMachineScaleSetNetworkConfigurationInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetNetworkConfigurationOutput() VirtualMachineScaleSetNetworkConfigurationOutput
	ToVirtualMachineScaleSetNetworkConfigurationOutputWithContext(context.Context) VirtualMachineScaleSetNetworkConfigurationOutput
}

type VirtualMachineScaleSetNetworkConfigurationArgs struct {
	Id               pulumi.StringPtrInput                           `pulumi:"id"`
	IpConfigurations VirtualMachineScaleSetIPConfigurationArrayInput `pulumi:"ipConfigurations"`
	Name             pulumi.StringInput                              `pulumi:"name"`
	Primary          pulumi.BoolPtrInput                             `pulumi:"primary"`
}

func (VirtualMachineScaleSetNetworkConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetNetworkConfiguration)(nil)).Elem()
}

func (i VirtualMachineScaleSetNetworkConfigurationArgs) ToVirtualMachineScaleSetNetworkConfigurationOutput() VirtualMachineScaleSetNetworkConfigurationOutput {
	return i.ToVirtualMachineScaleSetNetworkConfigurationOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetNetworkConfigurationArgs) ToVirtualMachineScaleSetNetworkConfigurationOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetNetworkConfigurationOutput)
}





type VirtualMachineScaleSetNetworkConfigurationArrayInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetNetworkConfigurationArrayOutput() VirtualMachineScaleSetNetworkConfigurationArrayOutput
	ToVirtualMachineScaleSetNetworkConfigurationArrayOutputWithContext(context.Context) VirtualMachineScaleSetNetworkConfigurationArrayOutput
}

type VirtualMachineScaleSetNetworkConfigurationArray []VirtualMachineScaleSetNetworkConfigurationInput

func (VirtualMachineScaleSetNetworkConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetNetworkConfiguration)(nil)).Elem()
}

func (i VirtualMachineScaleSetNetworkConfigurationArray) ToVirtualMachineScaleSetNetworkConfigurationArrayOutput() VirtualMachineScaleSetNetworkConfigurationArrayOutput {
	return i.ToVirtualMachineScaleSetNetworkConfigurationArrayOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetNetworkConfigurationArray) ToVirtualMachineScaleSetNetworkConfigurationArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetNetworkConfigurationArrayOutput)
}

type VirtualMachineScaleSetNetworkConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetNetworkConfiguration)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) ToVirtualMachineScaleSetNetworkConfigurationOutput() VirtualMachineScaleSetNetworkConfigurationOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) ToVirtualMachineScaleSetNetworkConfigurationOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfiguration) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) IpConfigurations() VirtualMachineScaleSetIPConfigurationArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfiguration) []VirtualMachineScaleSetIPConfiguration {
		return v.IpConfigurations
	}).(VirtualMachineScaleSetIPConfigurationArrayOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfiguration) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfiguration) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

type VirtualMachineScaleSetNetworkConfigurationArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetNetworkConfiguration)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkConfigurationArrayOutput) ToVirtualMachineScaleSetNetworkConfigurationArrayOutput() VirtualMachineScaleSetNetworkConfigurationArrayOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationArrayOutput) ToVirtualMachineScaleSetNetworkConfigurationArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationArrayOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetNetworkConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetNetworkConfiguration {
		return vs[0].([]VirtualMachineScaleSetNetworkConfiguration)[vs[1].(int)]
	}).(VirtualMachineScaleSetNetworkConfigurationOutput)
}

type VirtualMachineScaleSetNetworkConfigurationResponse struct {
	Id               *string                                         `pulumi:"id"`
	IpConfigurations []VirtualMachineScaleSetIPConfigurationResponse `pulumi:"ipConfigurations"`
	Name             string                                          `pulumi:"name"`
	Primary          *bool                                           `pulumi:"primary"`
}

type VirtualMachineScaleSetNetworkConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetNetworkConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) ToVirtualMachineScaleSetNetworkConfigurationResponseOutput() VirtualMachineScaleSetNetworkConfigurationResponseOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) ToVirtualMachineScaleSetNetworkConfigurationResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationResponseOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) IpConfigurations() VirtualMachineScaleSetIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationResponse) []VirtualMachineScaleSetIPConfigurationResponse {
		return v.IpConfigurations
	}).(VirtualMachineScaleSetIPConfigurationResponseArrayOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationResponse) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

type VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetNetworkConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput) ToVirtualMachineScaleSetNetworkConfigurationResponseArrayOutput() VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput) ToVirtualMachineScaleSetNetworkConfigurationResponseArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetNetworkConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetNetworkConfigurationResponse {
		return vs[0].([]VirtualMachineScaleSetNetworkConfigurationResponse)[vs[1].(int)]
	}).(VirtualMachineScaleSetNetworkConfigurationResponseOutput)
}

type VirtualMachineScaleSetNetworkProfile struct {
	NetworkInterfaceConfigurations []VirtualMachineScaleSetNetworkConfiguration `pulumi:"networkInterfaceConfigurations"`
}





type VirtualMachineScaleSetNetworkProfileInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetNetworkProfileOutput() VirtualMachineScaleSetNetworkProfileOutput
	ToVirtualMachineScaleSetNetworkProfileOutputWithContext(context.Context) VirtualMachineScaleSetNetworkProfileOutput
}

type VirtualMachineScaleSetNetworkProfileArgs struct {
	NetworkInterfaceConfigurations VirtualMachineScaleSetNetworkConfigurationArrayInput `pulumi:"networkInterfaceConfigurations"`
}

func (VirtualMachineScaleSetNetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetNetworkProfile)(nil)).Elem()
}

func (i VirtualMachineScaleSetNetworkProfileArgs) ToVirtualMachineScaleSetNetworkProfileOutput() VirtualMachineScaleSetNetworkProfileOutput {
	return i.ToVirtualMachineScaleSetNetworkProfileOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetNetworkProfileArgs) ToVirtualMachineScaleSetNetworkProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetNetworkProfileOutput)
}

func (i VirtualMachineScaleSetNetworkProfileArgs) ToVirtualMachineScaleSetNetworkProfilePtrOutput() VirtualMachineScaleSetNetworkProfilePtrOutput {
	return i.ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetNetworkProfileArgs) ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetNetworkProfileOutput).ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetNetworkProfilePtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetNetworkProfilePtrOutput() VirtualMachineScaleSetNetworkProfilePtrOutput
	ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(context.Context) VirtualMachineScaleSetNetworkProfilePtrOutput
}

type virtualMachineScaleSetNetworkProfilePtrType VirtualMachineScaleSetNetworkProfileArgs

func VirtualMachineScaleSetNetworkProfilePtr(v *VirtualMachineScaleSetNetworkProfileArgs) VirtualMachineScaleSetNetworkProfilePtrInput {
	return (*virtualMachineScaleSetNetworkProfilePtrType)(v)
}

func (*virtualMachineScaleSetNetworkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetNetworkProfile)(nil)).Elem()
}

func (i *virtualMachineScaleSetNetworkProfilePtrType) ToVirtualMachineScaleSetNetworkProfilePtrOutput() VirtualMachineScaleSetNetworkProfilePtrOutput {
	return i.ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetNetworkProfilePtrType) ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetNetworkProfilePtrOutput)
}

type VirtualMachineScaleSetNetworkProfileOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetNetworkProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkProfileOutput) ToVirtualMachineScaleSetNetworkProfileOutput() VirtualMachineScaleSetNetworkProfileOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkProfileOutput) ToVirtualMachineScaleSetNetworkProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkProfileOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkProfileOutput) ToVirtualMachineScaleSetNetworkProfilePtrOutput() VirtualMachineScaleSetNetworkProfilePtrOutput {
	return o.ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetNetworkProfileOutput) ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetNetworkProfile) *VirtualMachineScaleSetNetworkProfile {
		return &v
	}).(VirtualMachineScaleSetNetworkProfilePtrOutput)
}

func (o VirtualMachineScaleSetNetworkProfileOutput) NetworkInterfaceConfigurations() VirtualMachineScaleSetNetworkConfigurationArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkProfile) []VirtualMachineScaleSetNetworkConfiguration {
		return v.NetworkInterfaceConfigurations
	}).(VirtualMachineScaleSetNetworkConfigurationArrayOutput)
}

type VirtualMachineScaleSetNetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetNetworkProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkProfilePtrOutput) ToVirtualMachineScaleSetNetworkProfilePtrOutput() VirtualMachineScaleSetNetworkProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkProfilePtrOutput) ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkProfilePtrOutput) Elem() VirtualMachineScaleSetNetworkProfileOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkProfile) VirtualMachineScaleSetNetworkProfile {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetNetworkProfile
		return ret
	}).(VirtualMachineScaleSetNetworkProfileOutput)
}

func (o VirtualMachineScaleSetNetworkProfilePtrOutput) NetworkInterfaceConfigurations() VirtualMachineScaleSetNetworkConfigurationArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkProfile) []VirtualMachineScaleSetNetworkConfiguration {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaceConfigurations
	}).(VirtualMachineScaleSetNetworkConfigurationArrayOutput)
}

type VirtualMachineScaleSetNetworkProfileResponse struct {
	NetworkInterfaceConfigurations []VirtualMachineScaleSetNetworkConfigurationResponse `pulumi:"networkInterfaceConfigurations"`
}

type VirtualMachineScaleSetNetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetNetworkProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkProfileResponseOutput) ToVirtualMachineScaleSetNetworkProfileResponseOutput() VirtualMachineScaleSetNetworkProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkProfileResponseOutput) ToVirtualMachineScaleSetNetworkProfileResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkProfileResponseOutput) NetworkInterfaceConfigurations() VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkProfileResponse) []VirtualMachineScaleSetNetworkConfigurationResponse {
		return v.NetworkInterfaceConfigurations
	}).(VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput)
}

type VirtualMachineScaleSetNetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetNetworkProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkProfileResponsePtrOutput) ToVirtualMachineScaleSetNetworkProfileResponsePtrOutput() VirtualMachineScaleSetNetworkProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkProfileResponsePtrOutput) ToVirtualMachineScaleSetNetworkProfileResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkProfileResponsePtrOutput) Elem() VirtualMachineScaleSetNetworkProfileResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkProfileResponse) VirtualMachineScaleSetNetworkProfileResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetNetworkProfileResponse
		return ret
	}).(VirtualMachineScaleSetNetworkProfileResponseOutput)
}

func (o VirtualMachineScaleSetNetworkProfileResponsePtrOutput) NetworkInterfaceConfigurations() VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkProfileResponse) []VirtualMachineScaleSetNetworkConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaceConfigurations
	}).(VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput)
}

type VirtualMachineScaleSetOSDisk struct {
	Caching       *CachingTypes         `pulumi:"caching"`
	CreateOption  DiskCreateOptionTypes `pulumi:"createOption"`
	Image         *VirtualHardDisk      `pulumi:"image"`
	Name          string                `pulumi:"name"`
	OsType        *OperatingSystemTypes `pulumi:"osType"`
	VhdContainers []string              `pulumi:"vhdContainers"`
}





type VirtualMachineScaleSetOSDiskInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetOSDiskOutput() VirtualMachineScaleSetOSDiskOutput
	ToVirtualMachineScaleSetOSDiskOutputWithContext(context.Context) VirtualMachineScaleSetOSDiskOutput
}

type VirtualMachineScaleSetOSDiskArgs struct {
	Caching       CachingTypesPtrInput         `pulumi:"caching"`
	CreateOption  DiskCreateOptionTypesInput   `pulumi:"createOption"`
	Image         VirtualHardDiskPtrInput      `pulumi:"image"`
	Name          pulumi.StringInput           `pulumi:"name"`
	OsType        OperatingSystemTypesPtrInput `pulumi:"osType"`
	VhdContainers pulumi.StringArrayInput      `pulumi:"vhdContainers"`
}

func (VirtualMachineScaleSetOSDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetOSDisk)(nil)).Elem()
}

func (i VirtualMachineScaleSetOSDiskArgs) ToVirtualMachineScaleSetOSDiskOutput() VirtualMachineScaleSetOSDiskOutput {
	return i.ToVirtualMachineScaleSetOSDiskOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetOSDiskArgs) ToVirtualMachineScaleSetOSDiskOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetOSDiskOutput)
}

func (i VirtualMachineScaleSetOSDiskArgs) ToVirtualMachineScaleSetOSDiskPtrOutput() VirtualMachineScaleSetOSDiskPtrOutput {
	return i.ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetOSDiskArgs) ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetOSDiskOutput).ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetOSDiskPtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetOSDiskPtrOutput() VirtualMachineScaleSetOSDiskPtrOutput
	ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(context.Context) VirtualMachineScaleSetOSDiskPtrOutput
}

type virtualMachineScaleSetOSDiskPtrType VirtualMachineScaleSetOSDiskArgs

func VirtualMachineScaleSetOSDiskPtr(v *VirtualMachineScaleSetOSDiskArgs) VirtualMachineScaleSetOSDiskPtrInput {
	return (*virtualMachineScaleSetOSDiskPtrType)(v)
}

func (*virtualMachineScaleSetOSDiskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetOSDisk)(nil)).Elem()
}

func (i *virtualMachineScaleSetOSDiskPtrType) ToVirtualMachineScaleSetOSDiskPtrOutput() VirtualMachineScaleSetOSDiskPtrOutput {
	return i.ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetOSDiskPtrType) ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetOSDiskPtrOutput)
}

type VirtualMachineScaleSetOSDiskOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOSDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetOSDisk)(nil)).Elem()
}

func (o VirtualMachineScaleSetOSDiskOutput) ToVirtualMachineScaleSetOSDiskOutput() VirtualMachineScaleSetOSDiskOutput {
	return o
}

func (o VirtualMachineScaleSetOSDiskOutput) ToVirtualMachineScaleSetOSDiskOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSDiskOutput {
	return o
}

func (o VirtualMachineScaleSetOSDiskOutput) ToVirtualMachineScaleSetOSDiskPtrOutput() VirtualMachineScaleSetOSDiskPtrOutput {
	return o.ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetOSDiskOutput) ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSDiskPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetOSDisk) *VirtualMachineScaleSetOSDisk {
		return &v
	}).(VirtualMachineScaleSetOSDiskPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) Caching() CachingTypesPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) *CachingTypes { return v.Caching }).(CachingTypesPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) CreateOption() DiskCreateOptionTypesOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) DiskCreateOptionTypes { return v.CreateOption }).(DiskCreateOptionTypesOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) Image() VirtualHardDiskPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) *VirtualHardDisk { return v.Image }).(VirtualHardDiskPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) OsType() OperatingSystemTypesPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) *OperatingSystemTypes { return v.OsType }).(OperatingSystemTypesPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) VhdContainers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) []string { return v.VhdContainers }).(pulumi.StringArrayOutput)
}

type VirtualMachineScaleSetOSDiskPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOSDiskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetOSDisk)(nil)).Elem()
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) ToVirtualMachineScaleSetOSDiskPtrOutput() VirtualMachineScaleSetOSDiskPtrOutput {
	return o
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSDiskPtrOutput {
	return o
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) Elem() VirtualMachineScaleSetOSDiskOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) VirtualMachineScaleSetOSDisk {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetOSDisk
		return ret
	}).(VirtualMachineScaleSetOSDiskOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) Caching() CachingTypesPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) *CachingTypes {
		if v == nil {
			return nil
		}
		return v.Caching
	}).(CachingTypesPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) CreateOption() DiskCreateOptionTypesPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) *DiskCreateOptionTypes {
		if v == nil {
			return nil
		}
		return &v.CreateOption
	}).(DiskCreateOptionTypesPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) Image() VirtualHardDiskPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) *VirtualHardDisk {
		if v == nil {
			return nil
		}
		return v.Image
	}).(VirtualHardDiskPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) OsType() OperatingSystemTypesPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) *OperatingSystemTypes {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(OperatingSystemTypesPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) VhdContainers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) []string {
		if v == nil {
			return nil
		}
		return v.VhdContainers
	}).(pulumi.StringArrayOutput)
}

type VirtualMachineScaleSetOSDiskResponse struct {
	Caching       *string                  `pulumi:"caching"`
	CreateOption  string                   `pulumi:"createOption"`
	Image         *VirtualHardDiskResponse `pulumi:"image"`
	Name          string                   `pulumi:"name"`
	OsType        *string                  `pulumi:"osType"`
	VhdContainers []string                 `pulumi:"vhdContainers"`
}

type VirtualMachineScaleSetOSDiskResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOSDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetOSDiskResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) ToVirtualMachineScaleSetOSDiskResponseOutput() VirtualMachineScaleSetOSDiskResponseOutput {
	return o
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) ToVirtualMachineScaleSetOSDiskResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSDiskResponseOutput {
	return o
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) *string { return v.Caching }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) CreateOption() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) string { return v.CreateOption }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) Image() VirtualHardDiskResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) *VirtualHardDiskResponse { return v.Image }).(VirtualHardDiskResponsePtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) VhdContainers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) []string { return v.VhdContainers }).(pulumi.StringArrayOutput)
}

type VirtualMachineScaleSetOSDiskResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOSDiskResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetOSDiskResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) ToVirtualMachineScaleSetOSDiskResponsePtrOutput() VirtualMachineScaleSetOSDiskResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) ToVirtualMachineScaleSetOSDiskResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSDiskResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) Elem() VirtualMachineScaleSetOSDiskResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) VirtualMachineScaleSetOSDiskResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetOSDiskResponse
		return ret
	}).(VirtualMachineScaleSetOSDiskResponseOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.Caching
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) CreateOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreateOption
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) Image() VirtualHardDiskResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) *VirtualHardDiskResponse {
		if v == nil {
			return nil
		}
		return v.Image
	}).(VirtualHardDiskResponsePtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) VhdContainers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) []string {
		if v == nil {
			return nil
		}
		return v.VhdContainers
	}).(pulumi.StringArrayOutput)
}

type VirtualMachineScaleSetOSProfile struct {
	AdminPassword        *string               `pulumi:"adminPassword"`
	AdminUsername        *string               `pulumi:"adminUsername"`
	ComputerNamePrefix   *string               `pulumi:"computerNamePrefix"`
	CustomData           *string               `pulumi:"customData"`
	LinuxConfiguration   *LinuxConfiguration   `pulumi:"linuxConfiguration"`
	Secrets              []VaultSecretGroup    `pulumi:"secrets"`
	WindowsConfiguration *WindowsConfiguration `pulumi:"windowsConfiguration"`
}





type VirtualMachineScaleSetOSProfileInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetOSProfileOutput() VirtualMachineScaleSetOSProfileOutput
	ToVirtualMachineScaleSetOSProfileOutputWithContext(context.Context) VirtualMachineScaleSetOSProfileOutput
}

type VirtualMachineScaleSetOSProfileArgs struct {
	AdminPassword        pulumi.StringPtrInput        `pulumi:"adminPassword"`
	AdminUsername        pulumi.StringPtrInput        `pulumi:"adminUsername"`
	ComputerNamePrefix   pulumi.StringPtrInput        `pulumi:"computerNamePrefix"`
	CustomData           pulumi.StringPtrInput        `pulumi:"customData"`
	LinuxConfiguration   LinuxConfigurationPtrInput   `pulumi:"linuxConfiguration"`
	Secrets              VaultSecretGroupArrayInput   `pulumi:"secrets"`
	WindowsConfiguration WindowsConfigurationPtrInput `pulumi:"windowsConfiguration"`
}

func (VirtualMachineScaleSetOSProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetOSProfile)(nil)).Elem()
}

func (i VirtualMachineScaleSetOSProfileArgs) ToVirtualMachineScaleSetOSProfileOutput() VirtualMachineScaleSetOSProfileOutput {
	return i.ToVirtualMachineScaleSetOSProfileOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetOSProfileArgs) ToVirtualMachineScaleSetOSProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetOSProfileOutput)
}

func (i VirtualMachineScaleSetOSProfileArgs) ToVirtualMachineScaleSetOSProfilePtrOutput() VirtualMachineScaleSetOSProfilePtrOutput {
	return i.ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetOSProfileArgs) ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetOSProfileOutput).ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetOSProfilePtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetOSProfilePtrOutput() VirtualMachineScaleSetOSProfilePtrOutput
	ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(context.Context) VirtualMachineScaleSetOSProfilePtrOutput
}

type virtualMachineScaleSetOSProfilePtrType VirtualMachineScaleSetOSProfileArgs

func VirtualMachineScaleSetOSProfilePtr(v *VirtualMachineScaleSetOSProfileArgs) VirtualMachineScaleSetOSProfilePtrInput {
	return (*virtualMachineScaleSetOSProfilePtrType)(v)
}

func (*virtualMachineScaleSetOSProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetOSProfile)(nil)).Elem()
}

func (i *virtualMachineScaleSetOSProfilePtrType) ToVirtualMachineScaleSetOSProfilePtrOutput() VirtualMachineScaleSetOSProfilePtrOutput {
	return i.ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetOSProfilePtrType) ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetOSProfilePtrOutput)
}

type VirtualMachineScaleSetOSProfileOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOSProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetOSProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetOSProfileOutput) ToVirtualMachineScaleSetOSProfileOutput() VirtualMachineScaleSetOSProfileOutput {
	return o
}

func (o VirtualMachineScaleSetOSProfileOutput) ToVirtualMachineScaleSetOSProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSProfileOutput {
	return o
}

func (o VirtualMachineScaleSetOSProfileOutput) ToVirtualMachineScaleSetOSProfilePtrOutput() VirtualMachineScaleSetOSProfilePtrOutput {
	return o.ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetOSProfileOutput) ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetOSProfile) *VirtualMachineScaleSetOSProfile {
		return &v
	}).(VirtualMachineScaleSetOSProfilePtrOutput)
}

func (o VirtualMachineScaleSetOSProfileOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfile) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfile) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileOutput) ComputerNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfile) *string { return v.ComputerNamePrefix }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfile) *string { return v.CustomData }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileOutput) LinuxConfiguration() LinuxConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfile) *LinuxConfiguration { return v.LinuxConfiguration }).(LinuxConfigurationPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileOutput) Secrets() VaultSecretGroupArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfile) []VaultSecretGroup { return v.Secrets }).(VaultSecretGroupArrayOutput)
}

func (o VirtualMachineScaleSetOSProfileOutput) WindowsConfiguration() WindowsConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfile) *WindowsConfiguration { return v.WindowsConfiguration }).(WindowsConfigurationPtrOutput)
}

type VirtualMachineScaleSetOSProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOSProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetOSProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) ToVirtualMachineScaleSetOSProfilePtrOutput() VirtualMachineScaleSetOSProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) Elem() VirtualMachineScaleSetOSProfileOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) VirtualMachineScaleSetOSProfile {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetOSProfile
		return ret
	}).(VirtualMachineScaleSetOSProfileOutput)
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) ComputerNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) *string {
		if v == nil {
			return nil
		}
		return v.ComputerNamePrefix
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) *string {
		if v == nil {
			return nil
		}
		return v.CustomData
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) LinuxConfiguration() LinuxConfigurationPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) *LinuxConfiguration {
		if v == nil {
			return nil
		}
		return v.LinuxConfiguration
	}).(LinuxConfigurationPtrOutput)
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) Secrets() VaultSecretGroupArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) []VaultSecretGroup {
		if v == nil {
			return nil
		}
		return v.Secrets
	}).(VaultSecretGroupArrayOutput)
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) WindowsConfiguration() WindowsConfigurationPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) *WindowsConfiguration {
		if v == nil {
			return nil
		}
		return v.WindowsConfiguration
	}).(WindowsConfigurationPtrOutput)
}

type VirtualMachineScaleSetOSProfileResponse struct {
	AdminPassword        *string                       `pulumi:"adminPassword"`
	AdminUsername        *string                       `pulumi:"adminUsername"`
	ComputerNamePrefix   *string                       `pulumi:"computerNamePrefix"`
	CustomData           *string                       `pulumi:"customData"`
	LinuxConfiguration   *LinuxConfigurationResponse   `pulumi:"linuxConfiguration"`
	Secrets              []VaultSecretGroupResponse    `pulumi:"secrets"`
	WindowsConfiguration *WindowsConfigurationResponse `pulumi:"windowsConfiguration"`
}

type VirtualMachineScaleSetOSProfileResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOSProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetOSProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) ToVirtualMachineScaleSetOSProfileResponseOutput() VirtualMachineScaleSetOSProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) ToVirtualMachineScaleSetOSProfileResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfileResponse) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfileResponse) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) ComputerNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfileResponse) *string { return v.ComputerNamePrefix }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfileResponse) *string { return v.CustomData }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) LinuxConfiguration() LinuxConfigurationResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfileResponse) *LinuxConfigurationResponse {
		return v.LinuxConfiguration
	}).(LinuxConfigurationResponsePtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) Secrets() VaultSecretGroupResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfileResponse) []VaultSecretGroupResponse { return v.Secrets }).(VaultSecretGroupResponseArrayOutput)
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) WindowsConfiguration() WindowsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfileResponse) *WindowsConfigurationResponse {
		return v.WindowsConfiguration
	}).(WindowsConfigurationResponsePtrOutput)
}

type VirtualMachineScaleSetOSProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOSProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetOSProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) ToVirtualMachineScaleSetOSProfileResponsePtrOutput() VirtualMachineScaleSetOSProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) ToVirtualMachineScaleSetOSProfileResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) Elem() VirtualMachineScaleSetOSProfileResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) VirtualMachineScaleSetOSProfileResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetOSProfileResponse
		return ret
	}).(VirtualMachineScaleSetOSProfileResponseOutput)
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) ComputerNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ComputerNamePrefix
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomData
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) LinuxConfiguration() LinuxConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) *LinuxConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.LinuxConfiguration
	}).(LinuxConfigurationResponsePtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) Secrets() VaultSecretGroupResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) []VaultSecretGroupResponse {
		if v == nil {
			return nil
		}
		return v.Secrets
	}).(VaultSecretGroupResponseArrayOutput)
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) WindowsConfiguration() WindowsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) *WindowsConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.WindowsConfiguration
	}).(WindowsConfigurationResponsePtrOutput)
}

type VirtualMachineScaleSetStorageProfile struct {
	ImageReference *ImageReference               `pulumi:"imageReference"`
	OsDisk         *VirtualMachineScaleSetOSDisk `pulumi:"osDisk"`
}





type VirtualMachineScaleSetStorageProfileInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetStorageProfileOutput() VirtualMachineScaleSetStorageProfileOutput
	ToVirtualMachineScaleSetStorageProfileOutputWithContext(context.Context) VirtualMachineScaleSetStorageProfileOutput
}

type VirtualMachineScaleSetStorageProfileArgs struct {
	ImageReference ImageReferencePtrInput               `pulumi:"imageReference"`
	OsDisk         VirtualMachineScaleSetOSDiskPtrInput `pulumi:"osDisk"`
}

func (VirtualMachineScaleSetStorageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetStorageProfile)(nil)).Elem()
}

func (i VirtualMachineScaleSetStorageProfileArgs) ToVirtualMachineScaleSetStorageProfileOutput() VirtualMachineScaleSetStorageProfileOutput {
	return i.ToVirtualMachineScaleSetStorageProfileOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetStorageProfileArgs) ToVirtualMachineScaleSetStorageProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetStorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetStorageProfileOutput)
}

func (i VirtualMachineScaleSetStorageProfileArgs) ToVirtualMachineScaleSetStorageProfilePtrOutput() VirtualMachineScaleSetStorageProfilePtrOutput {
	return i.ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetStorageProfileArgs) ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetStorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetStorageProfileOutput).ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetStorageProfilePtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetStorageProfilePtrOutput() VirtualMachineScaleSetStorageProfilePtrOutput
	ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(context.Context) VirtualMachineScaleSetStorageProfilePtrOutput
}

type virtualMachineScaleSetStorageProfilePtrType VirtualMachineScaleSetStorageProfileArgs

func VirtualMachineScaleSetStorageProfilePtr(v *VirtualMachineScaleSetStorageProfileArgs) VirtualMachineScaleSetStorageProfilePtrInput {
	return (*virtualMachineScaleSetStorageProfilePtrType)(v)
}

func (*virtualMachineScaleSetStorageProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetStorageProfile)(nil)).Elem()
}

func (i *virtualMachineScaleSetStorageProfilePtrType) ToVirtualMachineScaleSetStorageProfilePtrOutput() VirtualMachineScaleSetStorageProfilePtrOutput {
	return i.ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetStorageProfilePtrType) ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetStorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetStorageProfilePtrOutput)
}

type VirtualMachineScaleSetStorageProfileOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetStorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetStorageProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetStorageProfileOutput) ToVirtualMachineScaleSetStorageProfileOutput() VirtualMachineScaleSetStorageProfileOutput {
	return o
}

func (o VirtualMachineScaleSetStorageProfileOutput) ToVirtualMachineScaleSetStorageProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetStorageProfileOutput {
	return o
}

func (o VirtualMachineScaleSetStorageProfileOutput) ToVirtualMachineScaleSetStorageProfilePtrOutput() VirtualMachineScaleSetStorageProfilePtrOutput {
	return o.ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetStorageProfileOutput) ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetStorageProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetStorageProfile) *VirtualMachineScaleSetStorageProfile {
		return &v
	}).(VirtualMachineScaleSetStorageProfilePtrOutput)
}

func (o VirtualMachineScaleSetStorageProfileOutput) ImageReference() ImageReferencePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetStorageProfile) *ImageReference { return v.ImageReference }).(ImageReferencePtrOutput)
}

func (o VirtualMachineScaleSetStorageProfileOutput) OsDisk() VirtualMachineScaleSetOSDiskPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetStorageProfile) *VirtualMachineScaleSetOSDisk { return v.OsDisk }).(VirtualMachineScaleSetOSDiskPtrOutput)
}

type VirtualMachineScaleSetStorageProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetStorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetStorageProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetStorageProfilePtrOutput) ToVirtualMachineScaleSetStorageProfilePtrOutput() VirtualMachineScaleSetStorageProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetStorageProfilePtrOutput) ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetStorageProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetStorageProfilePtrOutput) Elem() VirtualMachineScaleSetStorageProfileOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetStorageProfile) VirtualMachineScaleSetStorageProfile {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetStorageProfile
		return ret
	}).(VirtualMachineScaleSetStorageProfileOutput)
}

func (o VirtualMachineScaleSetStorageProfilePtrOutput) ImageReference() ImageReferencePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetStorageProfile) *ImageReference {
		if v == nil {
			return nil
		}
		return v.ImageReference
	}).(ImageReferencePtrOutput)
}

func (o VirtualMachineScaleSetStorageProfilePtrOutput) OsDisk() VirtualMachineScaleSetOSDiskPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetStorageProfile) *VirtualMachineScaleSetOSDisk {
		if v == nil {
			return nil
		}
		return v.OsDisk
	}).(VirtualMachineScaleSetOSDiskPtrOutput)
}

type VirtualMachineScaleSetStorageProfileResponse struct {
	ImageReference *ImageReferenceResponse               `pulumi:"imageReference"`
	OsDisk         *VirtualMachineScaleSetOSDiskResponse `pulumi:"osDisk"`
}

type VirtualMachineScaleSetStorageProfileResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetStorageProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetStorageProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetStorageProfileResponseOutput) ToVirtualMachineScaleSetStorageProfileResponseOutput() VirtualMachineScaleSetStorageProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetStorageProfileResponseOutput) ToVirtualMachineScaleSetStorageProfileResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetStorageProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetStorageProfileResponseOutput) ImageReference() ImageReferenceResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetStorageProfileResponse) *ImageReferenceResponse { return v.ImageReference }).(ImageReferenceResponsePtrOutput)
}

func (o VirtualMachineScaleSetStorageProfileResponseOutput) OsDisk() VirtualMachineScaleSetOSDiskResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetStorageProfileResponse) *VirtualMachineScaleSetOSDiskResponse {
		return v.OsDisk
	}).(VirtualMachineScaleSetOSDiskResponsePtrOutput)
}

type VirtualMachineScaleSetStorageProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetStorageProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetStorageProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetStorageProfileResponsePtrOutput) ToVirtualMachineScaleSetStorageProfileResponsePtrOutput() VirtualMachineScaleSetStorageProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetStorageProfileResponsePtrOutput) ToVirtualMachineScaleSetStorageProfileResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetStorageProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetStorageProfileResponsePtrOutput) Elem() VirtualMachineScaleSetStorageProfileResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetStorageProfileResponse) VirtualMachineScaleSetStorageProfileResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetStorageProfileResponse
		return ret
	}).(VirtualMachineScaleSetStorageProfileResponseOutput)
}

func (o VirtualMachineScaleSetStorageProfileResponsePtrOutput) ImageReference() ImageReferenceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetStorageProfileResponse) *ImageReferenceResponse {
		if v == nil {
			return nil
		}
		return v.ImageReference
	}).(ImageReferenceResponsePtrOutput)
}

func (o VirtualMachineScaleSetStorageProfileResponsePtrOutput) OsDisk() VirtualMachineScaleSetOSDiskResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetStorageProfileResponse) *VirtualMachineScaleSetOSDiskResponse {
		if v == nil {
			return nil
		}
		return v.OsDisk
	}).(VirtualMachineScaleSetOSDiskResponsePtrOutput)
}

type VirtualMachineScaleSetVMProfile struct {
	ExtensionProfile *VirtualMachineScaleSetExtensionProfile `pulumi:"extensionProfile"`
	NetworkProfile   *VirtualMachineScaleSetNetworkProfile   `pulumi:"networkProfile"`
	OsProfile        *VirtualMachineScaleSetOSProfile        `pulumi:"osProfile"`
	StorageProfile   *VirtualMachineScaleSetStorageProfile   `pulumi:"storageProfile"`
}





type VirtualMachineScaleSetVMProfileInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetVMProfileOutput() VirtualMachineScaleSetVMProfileOutput
	ToVirtualMachineScaleSetVMProfileOutputWithContext(context.Context) VirtualMachineScaleSetVMProfileOutput
}

type VirtualMachineScaleSetVMProfileArgs struct {
	ExtensionProfile VirtualMachineScaleSetExtensionProfilePtrInput `pulumi:"extensionProfile"`
	NetworkProfile   VirtualMachineScaleSetNetworkProfilePtrInput   `pulumi:"networkProfile"`
	OsProfile        VirtualMachineScaleSetOSProfilePtrInput        `pulumi:"osProfile"`
	StorageProfile   VirtualMachineScaleSetStorageProfilePtrInput   `pulumi:"storageProfile"`
}

func (VirtualMachineScaleSetVMProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetVMProfile)(nil)).Elem()
}

func (i VirtualMachineScaleSetVMProfileArgs) ToVirtualMachineScaleSetVMProfileOutput() VirtualMachineScaleSetVMProfileOutput {
	return i.ToVirtualMachineScaleSetVMProfileOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetVMProfileArgs) ToVirtualMachineScaleSetVMProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetVMProfileOutput)
}

func (i VirtualMachineScaleSetVMProfileArgs) ToVirtualMachineScaleSetVMProfilePtrOutput() VirtualMachineScaleSetVMProfilePtrOutput {
	return i.ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetVMProfileArgs) ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetVMProfileOutput).ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetVMProfilePtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetVMProfilePtrOutput() VirtualMachineScaleSetVMProfilePtrOutput
	ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(context.Context) VirtualMachineScaleSetVMProfilePtrOutput
}

type virtualMachineScaleSetVMProfilePtrType VirtualMachineScaleSetVMProfileArgs

func VirtualMachineScaleSetVMProfilePtr(v *VirtualMachineScaleSetVMProfileArgs) VirtualMachineScaleSetVMProfilePtrInput {
	return (*virtualMachineScaleSetVMProfilePtrType)(v)
}

func (*virtualMachineScaleSetVMProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVMProfile)(nil)).Elem()
}

func (i *virtualMachineScaleSetVMProfilePtrType) ToVirtualMachineScaleSetVMProfilePtrOutput() VirtualMachineScaleSetVMProfilePtrOutput {
	return i.ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetVMProfilePtrType) ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetVMProfilePtrOutput)
}

type VirtualMachineScaleSetVMProfileOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetVMProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMProfileOutput) ToVirtualMachineScaleSetVMProfileOutput() VirtualMachineScaleSetVMProfileOutput {
	return o
}

func (o VirtualMachineScaleSetVMProfileOutput) ToVirtualMachineScaleSetVMProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProfileOutput {
	return o
}

func (o VirtualMachineScaleSetVMProfileOutput) ToVirtualMachineScaleSetVMProfilePtrOutput() VirtualMachineScaleSetVMProfilePtrOutput {
	return o.ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetVMProfileOutput) ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetVMProfile {
		return &v
	}).(VirtualMachineScaleSetVMProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) ExtensionProfile() VirtualMachineScaleSetExtensionProfilePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetExtensionProfile {
		return v.ExtensionProfile
	}).(VirtualMachineScaleSetExtensionProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) NetworkProfile() VirtualMachineScaleSetNetworkProfilePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetNetworkProfile { return v.NetworkProfile }).(VirtualMachineScaleSetNetworkProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) OsProfile() VirtualMachineScaleSetOSProfilePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetOSProfile { return v.OsProfile }).(VirtualMachineScaleSetOSProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) StorageProfile() VirtualMachineScaleSetStorageProfilePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetStorageProfile { return v.StorageProfile }).(VirtualMachineScaleSetStorageProfilePtrOutput)
}

type VirtualMachineScaleSetVMProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVMProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) ToVirtualMachineScaleSetVMProfilePtrOutput() VirtualMachineScaleSetVMProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) Elem() VirtualMachineScaleSetVMProfileOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) VirtualMachineScaleSetVMProfile {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetVMProfile
		return ret
	}).(VirtualMachineScaleSetVMProfileOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) ExtensionProfile() VirtualMachineScaleSetExtensionProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetExtensionProfile {
		if v == nil {
			return nil
		}
		return v.ExtensionProfile
	}).(VirtualMachineScaleSetExtensionProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) NetworkProfile() VirtualMachineScaleSetNetworkProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetNetworkProfile {
		if v == nil {
			return nil
		}
		return v.NetworkProfile
	}).(VirtualMachineScaleSetNetworkProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) OsProfile() VirtualMachineScaleSetOSProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetOSProfile {
		if v == nil {
			return nil
		}
		return v.OsProfile
	}).(VirtualMachineScaleSetOSProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) StorageProfile() VirtualMachineScaleSetStorageProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetStorageProfile {
		if v == nil {
			return nil
		}
		return v.StorageProfile
	}).(VirtualMachineScaleSetStorageProfilePtrOutput)
}

type VirtualMachineScaleSetVMProfileResponse struct {
	ExtensionProfile *VirtualMachineScaleSetExtensionProfileResponse `pulumi:"extensionProfile"`
	NetworkProfile   *VirtualMachineScaleSetNetworkProfileResponse   `pulumi:"networkProfile"`
	OsProfile        *VirtualMachineScaleSetOSProfileResponse        `pulumi:"osProfile"`
	StorageProfile   *VirtualMachineScaleSetStorageProfileResponse   `pulumi:"storageProfile"`
}

type VirtualMachineScaleSetVMProfileResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetVMProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) ToVirtualMachineScaleSetVMProfileResponseOutput() VirtualMachineScaleSetVMProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) ToVirtualMachineScaleSetVMProfileResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) ExtensionProfile() VirtualMachineScaleSetExtensionProfileResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetExtensionProfileResponse {
		return v.ExtensionProfile
	}).(VirtualMachineScaleSetExtensionProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) NetworkProfile() VirtualMachineScaleSetNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetNetworkProfileResponse {
		return v.NetworkProfile
	}).(VirtualMachineScaleSetNetworkProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) OsProfile() VirtualMachineScaleSetOSProfileResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetOSProfileResponse {
		return v.OsProfile
	}).(VirtualMachineScaleSetOSProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) StorageProfile() VirtualMachineScaleSetStorageProfileResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetStorageProfileResponse {
		return v.StorageProfile
	}).(VirtualMachineScaleSetStorageProfileResponsePtrOutput)
}

type VirtualMachineScaleSetVMProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVMProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) ToVirtualMachineScaleSetVMProfileResponsePtrOutput() VirtualMachineScaleSetVMProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) ToVirtualMachineScaleSetVMProfileResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) Elem() VirtualMachineScaleSetVMProfileResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) VirtualMachineScaleSetVMProfileResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetVMProfileResponse
		return ret
	}).(VirtualMachineScaleSetVMProfileResponseOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) ExtensionProfile() VirtualMachineScaleSetExtensionProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetExtensionProfileResponse {
		if v == nil {
			return nil
		}
		return v.ExtensionProfile
	}).(VirtualMachineScaleSetExtensionProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) NetworkProfile() VirtualMachineScaleSetNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetNetworkProfileResponse {
		if v == nil {
			return nil
		}
		return v.NetworkProfile
	}).(VirtualMachineScaleSetNetworkProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) OsProfile() VirtualMachineScaleSetOSProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetOSProfileResponse {
		if v == nil {
			return nil
		}
		return v.OsProfile
	}).(VirtualMachineScaleSetOSProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) StorageProfile() VirtualMachineScaleSetStorageProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetStorageProfileResponse {
		if v == nil {
			return nil
		}
		return v.StorageProfile
	}).(VirtualMachineScaleSetStorageProfileResponsePtrOutput)
}

type WinRMConfiguration struct {
	Listeners []WinRMListener `pulumi:"listeners"`
}





type WinRMConfigurationInput interface {
	pulumi.Input

	ToWinRMConfigurationOutput() WinRMConfigurationOutput
	ToWinRMConfigurationOutputWithContext(context.Context) WinRMConfigurationOutput
}

type WinRMConfigurationArgs struct {
	Listeners WinRMListenerArrayInput `pulumi:"listeners"`
}

func (WinRMConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WinRMConfiguration)(nil)).Elem()
}

func (i WinRMConfigurationArgs) ToWinRMConfigurationOutput() WinRMConfigurationOutput {
	return i.ToWinRMConfigurationOutputWithContext(context.Background())
}

func (i WinRMConfigurationArgs) ToWinRMConfigurationOutputWithContext(ctx context.Context) WinRMConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WinRMConfigurationOutput)
}

func (i WinRMConfigurationArgs) ToWinRMConfigurationPtrOutput() WinRMConfigurationPtrOutput {
	return i.ToWinRMConfigurationPtrOutputWithContext(context.Background())
}

func (i WinRMConfigurationArgs) ToWinRMConfigurationPtrOutputWithContext(ctx context.Context) WinRMConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WinRMConfigurationOutput).ToWinRMConfigurationPtrOutputWithContext(ctx)
}









type WinRMConfigurationPtrInput interface {
	pulumi.Input

	ToWinRMConfigurationPtrOutput() WinRMConfigurationPtrOutput
	ToWinRMConfigurationPtrOutputWithContext(context.Context) WinRMConfigurationPtrOutput
}

type winRMConfigurationPtrType WinRMConfigurationArgs

func WinRMConfigurationPtr(v *WinRMConfigurationArgs) WinRMConfigurationPtrInput {
	return (*winRMConfigurationPtrType)(v)
}

func (*winRMConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WinRMConfiguration)(nil)).Elem()
}

func (i *winRMConfigurationPtrType) ToWinRMConfigurationPtrOutput() WinRMConfigurationPtrOutput {
	return i.ToWinRMConfigurationPtrOutputWithContext(context.Background())
}

func (i *winRMConfigurationPtrType) ToWinRMConfigurationPtrOutputWithContext(ctx context.Context) WinRMConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WinRMConfigurationPtrOutput)
}

type WinRMConfigurationOutput struct{ *pulumi.OutputState }

func (WinRMConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WinRMConfiguration)(nil)).Elem()
}

func (o WinRMConfigurationOutput) ToWinRMConfigurationOutput() WinRMConfigurationOutput {
	return o
}

func (o WinRMConfigurationOutput) ToWinRMConfigurationOutputWithContext(ctx context.Context) WinRMConfigurationOutput {
	return o
}

func (o WinRMConfigurationOutput) ToWinRMConfigurationPtrOutput() WinRMConfigurationPtrOutput {
	return o.ToWinRMConfigurationPtrOutputWithContext(context.Background())
}

func (o WinRMConfigurationOutput) ToWinRMConfigurationPtrOutputWithContext(ctx context.Context) WinRMConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WinRMConfiguration) *WinRMConfiguration {
		return &v
	}).(WinRMConfigurationPtrOutput)
}

func (o WinRMConfigurationOutput) Listeners() WinRMListenerArrayOutput {
	return o.ApplyT(func(v WinRMConfiguration) []WinRMListener { return v.Listeners }).(WinRMListenerArrayOutput)
}

type WinRMConfigurationPtrOutput struct{ *pulumi.OutputState }

func (WinRMConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WinRMConfiguration)(nil)).Elem()
}

func (o WinRMConfigurationPtrOutput) ToWinRMConfigurationPtrOutput() WinRMConfigurationPtrOutput {
	return o
}

func (o WinRMConfigurationPtrOutput) ToWinRMConfigurationPtrOutputWithContext(ctx context.Context) WinRMConfigurationPtrOutput {
	return o
}

func (o WinRMConfigurationPtrOutput) Elem() WinRMConfigurationOutput {
	return o.ApplyT(func(v *WinRMConfiguration) WinRMConfiguration {
		if v != nil {
			return *v
		}
		var ret WinRMConfiguration
		return ret
	}).(WinRMConfigurationOutput)
}

func (o WinRMConfigurationPtrOutput) Listeners() WinRMListenerArrayOutput {
	return o.ApplyT(func(v *WinRMConfiguration) []WinRMListener {
		if v == nil {
			return nil
		}
		return v.Listeners
	}).(WinRMListenerArrayOutput)
}

type WinRMConfigurationResponse struct {
	Listeners []WinRMListenerResponse `pulumi:"listeners"`
}

type WinRMConfigurationResponseOutput struct{ *pulumi.OutputState }

func (WinRMConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WinRMConfigurationResponse)(nil)).Elem()
}

func (o WinRMConfigurationResponseOutput) ToWinRMConfigurationResponseOutput() WinRMConfigurationResponseOutput {
	return o
}

func (o WinRMConfigurationResponseOutput) ToWinRMConfigurationResponseOutputWithContext(ctx context.Context) WinRMConfigurationResponseOutput {
	return o
}

func (o WinRMConfigurationResponseOutput) Listeners() WinRMListenerResponseArrayOutput {
	return o.ApplyT(func(v WinRMConfigurationResponse) []WinRMListenerResponse { return v.Listeners }).(WinRMListenerResponseArrayOutput)
}

type WinRMConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (WinRMConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WinRMConfigurationResponse)(nil)).Elem()
}

func (o WinRMConfigurationResponsePtrOutput) ToWinRMConfigurationResponsePtrOutput() WinRMConfigurationResponsePtrOutput {
	return o
}

func (o WinRMConfigurationResponsePtrOutput) ToWinRMConfigurationResponsePtrOutputWithContext(ctx context.Context) WinRMConfigurationResponsePtrOutput {
	return o
}

func (o WinRMConfigurationResponsePtrOutput) Elem() WinRMConfigurationResponseOutput {
	return o.ApplyT(func(v *WinRMConfigurationResponse) WinRMConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret WinRMConfigurationResponse
		return ret
	}).(WinRMConfigurationResponseOutput)
}

func (o WinRMConfigurationResponsePtrOutput) Listeners() WinRMListenerResponseArrayOutput {
	return o.ApplyT(func(v *WinRMConfigurationResponse) []WinRMListenerResponse {
		if v == nil {
			return nil
		}
		return v.Listeners
	}).(WinRMListenerResponseArrayOutput)
}

type WinRMListener struct {
	CertificateUrl *string        `pulumi:"certificateUrl"`
	Protocol       *ProtocolTypes `pulumi:"protocol"`
}





type WinRMListenerInput interface {
	pulumi.Input

	ToWinRMListenerOutput() WinRMListenerOutput
	ToWinRMListenerOutputWithContext(context.Context) WinRMListenerOutput
}

type WinRMListenerArgs struct {
	CertificateUrl pulumi.StringPtrInput `pulumi:"certificateUrl"`
	Protocol       ProtocolTypesPtrInput `pulumi:"protocol"`
}

func (WinRMListenerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WinRMListener)(nil)).Elem()
}

func (i WinRMListenerArgs) ToWinRMListenerOutput() WinRMListenerOutput {
	return i.ToWinRMListenerOutputWithContext(context.Background())
}

func (i WinRMListenerArgs) ToWinRMListenerOutputWithContext(ctx context.Context) WinRMListenerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WinRMListenerOutput)
}





type WinRMListenerArrayInput interface {
	pulumi.Input

	ToWinRMListenerArrayOutput() WinRMListenerArrayOutput
	ToWinRMListenerArrayOutputWithContext(context.Context) WinRMListenerArrayOutput
}

type WinRMListenerArray []WinRMListenerInput

func (WinRMListenerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WinRMListener)(nil)).Elem()
}

func (i WinRMListenerArray) ToWinRMListenerArrayOutput() WinRMListenerArrayOutput {
	return i.ToWinRMListenerArrayOutputWithContext(context.Background())
}

func (i WinRMListenerArray) ToWinRMListenerArrayOutputWithContext(ctx context.Context) WinRMListenerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WinRMListenerArrayOutput)
}

type WinRMListenerOutput struct{ *pulumi.OutputState }

func (WinRMListenerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WinRMListener)(nil)).Elem()
}

func (o WinRMListenerOutput) ToWinRMListenerOutput() WinRMListenerOutput {
	return o
}

func (o WinRMListenerOutput) ToWinRMListenerOutputWithContext(ctx context.Context) WinRMListenerOutput {
	return o
}

func (o WinRMListenerOutput) CertificateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WinRMListener) *string { return v.CertificateUrl }).(pulumi.StringPtrOutput)
}

func (o WinRMListenerOutput) Protocol() ProtocolTypesPtrOutput {
	return o.ApplyT(func(v WinRMListener) *ProtocolTypes { return v.Protocol }).(ProtocolTypesPtrOutput)
}

type WinRMListenerArrayOutput struct{ *pulumi.OutputState }

func (WinRMListenerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WinRMListener)(nil)).Elem()
}

func (o WinRMListenerArrayOutput) ToWinRMListenerArrayOutput() WinRMListenerArrayOutput {
	return o
}

func (o WinRMListenerArrayOutput) ToWinRMListenerArrayOutputWithContext(ctx context.Context) WinRMListenerArrayOutput {
	return o
}

func (o WinRMListenerArrayOutput) Index(i pulumi.IntInput) WinRMListenerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WinRMListener {
		return vs[0].([]WinRMListener)[vs[1].(int)]
	}).(WinRMListenerOutput)
}

type WinRMListenerResponse struct {
	CertificateUrl *string `pulumi:"certificateUrl"`
	Protocol       *string `pulumi:"protocol"`
}

type WinRMListenerResponseOutput struct{ *pulumi.OutputState }

func (WinRMListenerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WinRMListenerResponse)(nil)).Elem()
}

func (o WinRMListenerResponseOutput) ToWinRMListenerResponseOutput() WinRMListenerResponseOutput {
	return o
}

func (o WinRMListenerResponseOutput) ToWinRMListenerResponseOutputWithContext(ctx context.Context) WinRMListenerResponseOutput {
	return o
}

func (o WinRMListenerResponseOutput) CertificateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WinRMListenerResponse) *string { return v.CertificateUrl }).(pulumi.StringPtrOutput)
}

func (o WinRMListenerResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WinRMListenerResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

type WinRMListenerResponseArrayOutput struct{ *pulumi.OutputState }

func (WinRMListenerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WinRMListenerResponse)(nil)).Elem()
}

func (o WinRMListenerResponseArrayOutput) ToWinRMListenerResponseArrayOutput() WinRMListenerResponseArrayOutput {
	return o
}

func (o WinRMListenerResponseArrayOutput) ToWinRMListenerResponseArrayOutputWithContext(ctx context.Context) WinRMListenerResponseArrayOutput {
	return o
}

func (o WinRMListenerResponseArrayOutput) Index(i pulumi.IntInput) WinRMListenerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WinRMListenerResponse {
		return vs[0].([]WinRMListenerResponse)[vs[1].(int)]
	}).(WinRMListenerResponseOutput)
}

type WindowsConfiguration struct {
	AdditionalUnattendContent []AdditionalUnattendContent `pulumi:"additionalUnattendContent"`
	EnableAutomaticUpdates    *bool                       `pulumi:"enableAutomaticUpdates"`
	ProvisionVMAgent          *bool                       `pulumi:"provisionVMAgent"`
	TimeZone                  *string                     `pulumi:"timeZone"`
	WinRM                     *WinRMConfiguration         `pulumi:"winRM"`
}





type WindowsConfigurationInput interface {
	pulumi.Input

	ToWindowsConfigurationOutput() WindowsConfigurationOutput
	ToWindowsConfigurationOutputWithContext(context.Context) WindowsConfigurationOutput
}

type WindowsConfigurationArgs struct {
	AdditionalUnattendContent AdditionalUnattendContentArrayInput `pulumi:"additionalUnattendContent"`
	EnableAutomaticUpdates    pulumi.BoolPtrInput                 `pulumi:"enableAutomaticUpdates"`
	ProvisionVMAgent          pulumi.BoolPtrInput                 `pulumi:"provisionVMAgent"`
	TimeZone                  pulumi.StringPtrInput               `pulumi:"timeZone"`
	WinRM                     WinRMConfigurationPtrInput          `pulumi:"winRM"`
}

func (WindowsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsConfiguration)(nil)).Elem()
}

func (i WindowsConfigurationArgs) ToWindowsConfigurationOutput() WindowsConfigurationOutput {
	return i.ToWindowsConfigurationOutputWithContext(context.Background())
}

func (i WindowsConfigurationArgs) ToWindowsConfigurationOutputWithContext(ctx context.Context) WindowsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsConfigurationOutput)
}

func (i WindowsConfigurationArgs) ToWindowsConfigurationPtrOutput() WindowsConfigurationPtrOutput {
	return i.ToWindowsConfigurationPtrOutputWithContext(context.Background())
}

func (i WindowsConfigurationArgs) ToWindowsConfigurationPtrOutputWithContext(ctx context.Context) WindowsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsConfigurationOutput).ToWindowsConfigurationPtrOutputWithContext(ctx)
}









type WindowsConfigurationPtrInput interface {
	pulumi.Input

	ToWindowsConfigurationPtrOutput() WindowsConfigurationPtrOutput
	ToWindowsConfigurationPtrOutputWithContext(context.Context) WindowsConfigurationPtrOutput
}

type windowsConfigurationPtrType WindowsConfigurationArgs

func WindowsConfigurationPtr(v *WindowsConfigurationArgs) WindowsConfigurationPtrInput {
	return (*windowsConfigurationPtrType)(v)
}

func (*windowsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsConfiguration)(nil)).Elem()
}

func (i *windowsConfigurationPtrType) ToWindowsConfigurationPtrOutput() WindowsConfigurationPtrOutput {
	return i.ToWindowsConfigurationPtrOutputWithContext(context.Background())
}

func (i *windowsConfigurationPtrType) ToWindowsConfigurationPtrOutputWithContext(ctx context.Context) WindowsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsConfigurationPtrOutput)
}

type WindowsConfigurationOutput struct{ *pulumi.OutputState }

func (WindowsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsConfiguration)(nil)).Elem()
}

func (o WindowsConfigurationOutput) ToWindowsConfigurationOutput() WindowsConfigurationOutput {
	return o
}

func (o WindowsConfigurationOutput) ToWindowsConfigurationOutputWithContext(ctx context.Context) WindowsConfigurationOutput {
	return o
}

func (o WindowsConfigurationOutput) ToWindowsConfigurationPtrOutput() WindowsConfigurationPtrOutput {
	return o.ToWindowsConfigurationPtrOutputWithContext(context.Background())
}

func (o WindowsConfigurationOutput) ToWindowsConfigurationPtrOutputWithContext(ctx context.Context) WindowsConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WindowsConfiguration) *WindowsConfiguration {
		return &v
	}).(WindowsConfigurationPtrOutput)
}

func (o WindowsConfigurationOutput) AdditionalUnattendContent() AdditionalUnattendContentArrayOutput {
	return o.ApplyT(func(v WindowsConfiguration) []AdditionalUnattendContent { return v.AdditionalUnattendContent }).(AdditionalUnattendContentArrayOutput)
}

func (o WindowsConfigurationOutput) EnableAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WindowsConfiguration) *bool { return v.EnableAutomaticUpdates }).(pulumi.BoolPtrOutput)
}

func (o WindowsConfigurationOutput) ProvisionVMAgent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WindowsConfiguration) *bool { return v.ProvisionVMAgent }).(pulumi.BoolPtrOutput)
}

func (o WindowsConfigurationOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsConfiguration) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

func (o WindowsConfigurationOutput) WinRM() WinRMConfigurationPtrOutput {
	return o.ApplyT(func(v WindowsConfiguration) *WinRMConfiguration { return v.WinRM }).(WinRMConfigurationPtrOutput)
}

type WindowsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (WindowsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsConfiguration)(nil)).Elem()
}

func (o WindowsConfigurationPtrOutput) ToWindowsConfigurationPtrOutput() WindowsConfigurationPtrOutput {
	return o
}

func (o WindowsConfigurationPtrOutput) ToWindowsConfigurationPtrOutputWithContext(ctx context.Context) WindowsConfigurationPtrOutput {
	return o
}

func (o WindowsConfigurationPtrOutput) Elem() WindowsConfigurationOutput {
	return o.ApplyT(func(v *WindowsConfiguration) WindowsConfiguration {
		if v != nil {
			return *v
		}
		var ret WindowsConfiguration
		return ret
	}).(WindowsConfigurationOutput)
}

func (o WindowsConfigurationPtrOutput) AdditionalUnattendContent() AdditionalUnattendContentArrayOutput {
	return o.ApplyT(func(v *WindowsConfiguration) []AdditionalUnattendContent {
		if v == nil {
			return nil
		}
		return v.AdditionalUnattendContent
	}).(AdditionalUnattendContentArrayOutput)
}

func (o WindowsConfigurationPtrOutput) EnableAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WindowsConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAutomaticUpdates
	}).(pulumi.BoolPtrOutput)
}

func (o WindowsConfigurationPtrOutput) ProvisionVMAgent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WindowsConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.ProvisionVMAgent
	}).(pulumi.BoolPtrOutput)
}

func (o WindowsConfigurationPtrOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WindowsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.TimeZone
	}).(pulumi.StringPtrOutput)
}

func (o WindowsConfigurationPtrOutput) WinRM() WinRMConfigurationPtrOutput {
	return o.ApplyT(func(v *WindowsConfiguration) *WinRMConfiguration {
		if v == nil {
			return nil
		}
		return v.WinRM
	}).(WinRMConfigurationPtrOutput)
}

type WindowsConfigurationResponse struct {
	AdditionalUnattendContent []AdditionalUnattendContentResponse `pulumi:"additionalUnattendContent"`
	EnableAutomaticUpdates    *bool                               `pulumi:"enableAutomaticUpdates"`
	ProvisionVMAgent          *bool                               `pulumi:"provisionVMAgent"`
	TimeZone                  *string                             `pulumi:"timeZone"`
	WinRM                     *WinRMConfigurationResponse         `pulumi:"winRM"`
}

type WindowsConfigurationResponseOutput struct{ *pulumi.OutputState }

func (WindowsConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsConfigurationResponse)(nil)).Elem()
}

func (o WindowsConfigurationResponseOutput) ToWindowsConfigurationResponseOutput() WindowsConfigurationResponseOutput {
	return o
}

func (o WindowsConfigurationResponseOutput) ToWindowsConfigurationResponseOutputWithContext(ctx context.Context) WindowsConfigurationResponseOutput {
	return o
}

func (o WindowsConfigurationResponseOutput) AdditionalUnattendContent() AdditionalUnattendContentResponseArrayOutput {
	return o.ApplyT(func(v WindowsConfigurationResponse) []AdditionalUnattendContentResponse {
		return v.AdditionalUnattendContent
	}).(AdditionalUnattendContentResponseArrayOutput)
}

func (o WindowsConfigurationResponseOutput) EnableAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WindowsConfigurationResponse) *bool { return v.EnableAutomaticUpdates }).(pulumi.BoolPtrOutput)
}

func (o WindowsConfigurationResponseOutput) ProvisionVMAgent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WindowsConfigurationResponse) *bool { return v.ProvisionVMAgent }).(pulumi.BoolPtrOutput)
}

func (o WindowsConfigurationResponseOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsConfigurationResponse) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

func (o WindowsConfigurationResponseOutput) WinRM() WinRMConfigurationResponsePtrOutput {
	return o.ApplyT(func(v WindowsConfigurationResponse) *WinRMConfigurationResponse { return v.WinRM }).(WinRMConfigurationResponsePtrOutput)
}

type WindowsConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (WindowsConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsConfigurationResponse)(nil)).Elem()
}

func (o WindowsConfigurationResponsePtrOutput) ToWindowsConfigurationResponsePtrOutput() WindowsConfigurationResponsePtrOutput {
	return o
}

func (o WindowsConfigurationResponsePtrOutput) ToWindowsConfigurationResponsePtrOutputWithContext(ctx context.Context) WindowsConfigurationResponsePtrOutput {
	return o
}

func (o WindowsConfigurationResponsePtrOutput) Elem() WindowsConfigurationResponseOutput {
	return o.ApplyT(func(v *WindowsConfigurationResponse) WindowsConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret WindowsConfigurationResponse
		return ret
	}).(WindowsConfigurationResponseOutput)
}

func (o WindowsConfigurationResponsePtrOutput) AdditionalUnattendContent() AdditionalUnattendContentResponseArrayOutput {
	return o.ApplyT(func(v *WindowsConfigurationResponse) []AdditionalUnattendContentResponse {
		if v == nil {
			return nil
		}
		return v.AdditionalUnattendContent
	}).(AdditionalUnattendContentResponseArrayOutput)
}

func (o WindowsConfigurationResponsePtrOutput) EnableAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WindowsConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAutomaticUpdates
	}).(pulumi.BoolPtrOutput)
}

func (o WindowsConfigurationResponsePtrOutput) ProvisionVMAgent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WindowsConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ProvisionVMAgent
	}).(pulumi.BoolPtrOutput)
}

func (o WindowsConfigurationResponsePtrOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WindowsConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeZone
	}).(pulumi.StringPtrOutput)
}

func (o WindowsConfigurationResponsePtrOutput) WinRM() WinRMConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *WindowsConfigurationResponse) *WinRMConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.WinRM
	}).(WinRMConfigurationResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AdditionalUnattendContentOutput{})
	pulumi.RegisterOutputType(AdditionalUnattendContentArrayOutput{})
	pulumi.RegisterOutputType(AdditionalUnattendContentResponseOutput{})
	pulumi.RegisterOutputType(AdditionalUnattendContentResponseArrayOutput{})
	pulumi.RegisterOutputType(ApiEntityReferenceOutput{})
	pulumi.RegisterOutputType(ApiEntityReferenceResponseOutput{})
	pulumi.RegisterOutputType(BootDiagnosticsOutput{})
	pulumi.RegisterOutputType(BootDiagnosticsPtrOutput{})
	pulumi.RegisterOutputType(BootDiagnosticsInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(BootDiagnosticsInstanceViewResponsePtrOutput{})
	pulumi.RegisterOutputType(BootDiagnosticsResponseOutput{})
	pulumi.RegisterOutputType(BootDiagnosticsResponsePtrOutput{})
	pulumi.RegisterOutputType(DataDiskOutput{})
	pulumi.RegisterOutputType(DataDiskArrayOutput{})
	pulumi.RegisterOutputType(DataDiskResponseOutput{})
	pulumi.RegisterOutputType(DataDiskResponseArrayOutput{})
	pulumi.RegisterOutputType(DiagnosticsProfileOutput{})
	pulumi.RegisterOutputType(DiagnosticsProfilePtrOutput{})
	pulumi.RegisterOutputType(DiagnosticsProfileResponseOutput{})
	pulumi.RegisterOutputType(DiagnosticsProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSettingsOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSettingsPtrOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSettingsResponseOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(DiskInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(DiskInstanceViewResponseArrayOutput{})
	pulumi.RegisterOutputType(HardwareProfileOutput{})
	pulumi.RegisterOutputType(HardwareProfilePtrOutput{})
	pulumi.RegisterOutputType(HardwareProfileResponseOutput{})
	pulumi.RegisterOutputType(HardwareProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageReferenceOutput{})
	pulumi.RegisterOutputType(ImageReferencePtrOutput{})
	pulumi.RegisterOutputType(ImageReferenceResponseOutput{})
	pulumi.RegisterOutputType(ImageReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(InstanceViewStatusOutput{})
	pulumi.RegisterOutputType(InstanceViewStatusArrayOutput{})
	pulumi.RegisterOutputType(InstanceViewStatusResponseOutput{})
	pulumi.RegisterOutputType(InstanceViewStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(InstanceViewStatusResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferencePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultSecretReferenceOutput{})
	pulumi.RegisterOutputType(KeyVaultSecretReferencePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultSecretReferenceResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultSecretReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(LinuxConfigurationOutput{})
	pulumi.RegisterOutputType(LinuxConfigurationPtrOutput{})
	pulumi.RegisterOutputType(LinuxConfigurationResponseOutput{})
	pulumi.RegisterOutputType(LinuxConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceReferenceOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceReferenceArrayOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceReferenceResponseOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkProfileOutput{})
	pulumi.RegisterOutputType(NetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(OSDiskOutput{})
	pulumi.RegisterOutputType(OSDiskPtrOutput{})
	pulumi.RegisterOutputType(OSDiskResponseOutput{})
	pulumi.RegisterOutputType(OSDiskResponsePtrOutput{})
	pulumi.RegisterOutputType(OSProfileOutput{})
	pulumi.RegisterOutputType(OSProfilePtrOutput{})
	pulumi.RegisterOutputType(OSProfileResponseOutput{})
	pulumi.RegisterOutputType(OSProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(PlanOutput{})
	pulumi.RegisterOutputType(PlanPtrOutput{})
	pulumi.RegisterOutputType(PlanResponseOutput{})
	pulumi.RegisterOutputType(PlanResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SshConfigurationOutput{})
	pulumi.RegisterOutputType(SshConfigurationPtrOutput{})
	pulumi.RegisterOutputType(SshConfigurationResponseOutput{})
	pulumi.RegisterOutputType(SshConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(SshPublicKeyOutput{})
	pulumi.RegisterOutputType(SshPublicKeyArrayOutput{})
	pulumi.RegisterOutputType(SshPublicKeyResponseOutput{})
	pulumi.RegisterOutputType(SshPublicKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(StorageProfileOutput{})
	pulumi.RegisterOutputType(StorageProfilePtrOutput{})
	pulumi.RegisterOutputType(StorageProfileResponseOutput{})
	pulumi.RegisterOutputType(StorageProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(SubResourceOutput{})
	pulumi.RegisterOutputType(SubResourcePtrOutput{})
	pulumi.RegisterOutputType(SubResourceArrayOutput{})
	pulumi.RegisterOutputType(SubResourceResponseOutput{})
	pulumi.RegisterOutputType(SubResourceResponsePtrOutput{})
	pulumi.RegisterOutputType(SubResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(UpgradePolicyOutput{})
	pulumi.RegisterOutputType(UpgradePolicyPtrOutput{})
	pulumi.RegisterOutputType(UpgradePolicyResponseOutput{})
	pulumi.RegisterOutputType(UpgradePolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(VaultCertificateOutput{})
	pulumi.RegisterOutputType(VaultCertificateArrayOutput{})
	pulumi.RegisterOutputType(VaultCertificateResponseOutput{})
	pulumi.RegisterOutputType(VaultCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(VaultSecretGroupOutput{})
	pulumi.RegisterOutputType(VaultSecretGroupArrayOutput{})
	pulumi.RegisterOutputType(VaultSecretGroupResponseOutput{})
	pulumi.RegisterOutputType(VaultSecretGroupResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualHardDiskOutput{})
	pulumi.RegisterOutputType(VirtualHardDiskPtrOutput{})
	pulumi.RegisterOutputType(VirtualHardDiskResponseOutput{})
	pulumi.RegisterOutputType(VirtualHardDiskResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineAgentInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineAgentInstanceViewResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionHandlerInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionInstanceViewOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionInstanceViewPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionInstanceViewResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionInstanceViewResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionProfileOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionProfileResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIPConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIPConfigurationArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIPConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIPConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkConfigurationArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkProfileOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetOSDiskOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetOSDiskPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetOSDiskResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetOSDiskResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetOSProfileOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetOSProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetOSProfileResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetOSProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetStorageProfileOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetStorageProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetStorageProfileResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetStorageProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMProfileOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMProfileResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(WinRMConfigurationOutput{})
	pulumi.RegisterOutputType(WinRMConfigurationPtrOutput{})
	pulumi.RegisterOutputType(WinRMConfigurationResponseOutput{})
	pulumi.RegisterOutputType(WinRMConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(WinRMListenerOutput{})
	pulumi.RegisterOutputType(WinRMListenerArrayOutput{})
	pulumi.RegisterOutputType(WinRMListenerResponseOutput{})
	pulumi.RegisterOutputType(WinRMListenerResponseArrayOutput{})
	pulumi.RegisterOutputType(WindowsConfigurationOutput{})
	pulumi.RegisterOutputType(WindowsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(WindowsConfigurationResponseOutput{})
	pulumi.RegisterOutputType(WindowsConfigurationResponsePtrOutput{})
}
