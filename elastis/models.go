package elastis

import "github.com/hashicorp/terraform-plugin-framework/types"

// Location Model
type Locations struct {
	DisplayName types.String `tfsdk:"display_name"`
}

type LocationModel struct {
	Locations []Locations `tfsdk:"locations"`
}

// Virtual Machine Model
type VMStorageModel struct {
	UUID      types.String `tfsdk:"uuid"`
	UserID    types.Int64  `tfsdk:"user_id"`
	CreatedAt types.String `tfsdk:"created_at"`
	Name      types.String `tfsdk:"name"`
	Primary   types.Bool   `tfsdk:"primary"`
	Size      types.Int64  `tfsdk:"size"`
}

type VMPayloadModel struct {
	// Request
	Backup           types.Bool   `tfsdk:"backup"`
	BillingAccountID types.Int64  `tfsdk:"billing_account_id"`
	Description      types.String `tfsdk:"description"`
	Name             types.String `tfsdk:"name"`
	OSName           types.String `tfsdk:"os_name"`
	OSVersion        types.String `tfsdk:"os_version"`
	Username         types.String `tfsdk:"username"`
	Password         types.String `tfsdk:"password"`
	Disks            types.Int64  `tfsdk:"disks"`
	Memory           types.Int64  `tfsdk:"memory"`
	VCPU             types.Int64  `tfsdk:"vcpu"`
	ReservePublicIP  types.Bool   `tfsdk:"reserve_public_ip"`
	NetworkUUID      types.String `tfsdk:"network_uuid"`
	CloudInit        types.String `tfsdk:"cloud_init"`
	PublicKey        types.String `tfsdk:"public_key"`

	// Computed Response
	UUID           types.String `tfsdk:"uuid"`
	UserID         types.Int64  `tfsdk:"user_id"`
	Status         types.String `tfsdk:"status"`
	PrivateIPV4    types.String `tfsdk:"private_ipv4"`
	Hostname       types.String `tfsdk:"hostname"`
	BillingAccount types.Int64  `tfsdk:"billing_account"`
	CreatedAt      types.String `tfsdk:"created_at"`
	UpdatedAt      types.String `tfsdk:"updated_at"`
	Storage        types.List   `tfsdk:"storage"`
}

type VMModel struct {
	Location types.String   `tfsdk:"location"`
	VM       VMPayloadModel `tfsdk:"vm"`
}
