package elastis

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/nurimansyah89/terraform-elastis-provider/elastis/client"
	"github.com/nurimansyah89/terraform-elastis-provider/elastis/client/handlers"
)

type vmResource struct {
	client *client.ElastisClient
}

var (
	_ resource.Resource              = &vmResource{}
	_ resource.ResourceWithConfigure = &vmResource{}
)

func NewVMResource() resource.Resource {
	return &vmResource{}
}

func (r *vmResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vm"
}

func (r *vmResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"location": schema.StringAttribute{
				Required: true,
			},
			"vm": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					// Request
					"backup": schema.BoolAttribute{
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"billing_account_id": schema.Int64Attribute{
						Optional: true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"description": schema.StringAttribute{
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"name": schema.StringAttribute{
						Required: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"os_name": schema.StringAttribute{
						Required: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"os_version": schema.StringAttribute{
						Required: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"username": schema.StringAttribute{
						Required: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"password": schema.StringAttribute{
						Required:  true,
						Sensitive: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"disks": schema.Int64Attribute{
						Required: true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"memory": schema.Int64Attribute{
						Required: true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"vcpu": schema.Int64Attribute{
						Required: true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"reserve_public_ip": schema.BoolAttribute{
						Optional: true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"network_uuid": schema.StringAttribute{
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"cloud_init": schema.StringAttribute{
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"public_key": schema.StringAttribute{
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},

					// Computed Response
					"uuid": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"user_id": schema.Int64Attribute{
						Computed: true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"status": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"private_ipv4": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"hostname": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"billing_account": schema.Int64Attribute{
						Computed: true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"created_at": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"updated_at": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"storage": schema.ListNestedAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"uuid": schema.StringAttribute{
									Computed: true,
								},
								"user_id": schema.Int64Attribute{
									Computed: true,
								},
								"created_at": schema.StringAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"primary": schema.BoolAttribute{
									Computed: true,
								},
								"size": schema.Int64Attribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *vmResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.ElastisClient)
}

func (r *vmResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan VMModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	payload := client.VMPayload{
		Backup:           plan.VM.Backup.ValueBool(),
		BillingAccountID: int(plan.VM.BillingAccountID.ValueInt64()),
		Description:      plan.VM.Description.ValueString(),
		Name:             plan.VM.Name.ValueString(),
		OSName:           plan.VM.OSName.ValueString(),
		OSVersion:        plan.VM.OSVersion.ValueString(),
		Username:         plan.VM.Username.ValueString(),
		Password:         plan.VM.Password.ValueString(),
		Disks:            int(plan.VM.Disks.ValueInt64()),
		Memory:           int(plan.VM.Memory.ValueInt64()),
		VCPU:             int(plan.VM.VCPU.ValueInt64()),
		ReservePublicIP:  plan.VM.ReservePublicIP.ValueBool(),
		NetworkUUID:      plan.VM.NetworkUUID.ValueString(),
		CloudInit:        plan.VM.CloudInit.ValueString(),
		PublicKey:        plan.VM.PublicKey.ValueString(),
	}

	// Create new VM
	vmHandler := handlers.NewVM(r.client)
	vm, err := vmHandler.CreateVM(plan.Location.ValueString(), payload)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating VM",
			"There was error when creating VM: \n"+err.Error(),
		)
		return
	}

	// Storage
	var objectStorages basetypes.ObjectValue
	for _, storage := range vm.Storage {
		obj, _ := types.ObjectValue(
			map[string]attr.Type{
				"uuid":       types.StringType,
				"user_id":    types.Int64Type,
				"created_at": types.StringType,
				"name":       types.StringType,
				"primary":    types.BoolType,
				"size":       types.Int64Type,
			},
			map[string]attr.Value{
				"uuid":       types.StringValue(storage.UUID),
				"user_id":    types.Int64Value(int64(storage.UserID)),
				"created_at": types.StringValue(storage.CreatedAt),
				"name":       types.StringValue(storage.Name),
				"primary":    types.BoolValue(storage.Primary),
				"size":       types.Int64Value(int64(storage.Size)),
			},
		)
		objectStorages = obj
	}
	storages, _ := types.ListValue(types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"uuid":       types.StringType,
			"user_id":    types.Int64Type,
			"created_at": types.StringType,
			"name":       types.StringType,
			"primary":    types.BoolType,
			"size":       types.Int64Type,
		},
	}, []attr.Value{objectStorages})

	// Populate computed values
	plan.VM.Backup = types.BoolValue(vm.Backup)
	plan.VM.UUID = types.StringValue(vm.UUID)
	plan.VM.UserID = types.Int64Value(int64(vm.UserID))
	plan.VM.Status = types.StringValue(vm.Status)
	plan.VM.PrivateIPV4 = types.StringValue(vm.PrivateIPV4)
	plan.VM.Hostname = types.StringValue(vm.Hostname)
	plan.VM.BillingAccount = types.Int64Value(int64(vm.BillingAccount))
	plan.VM.CreatedAt = types.StringValue(vm.CreatedAt)
	plan.VM.UpdatedAt = types.StringValue(vm.UpdatedAt)
	plan.VM.Storage = storages

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
func (r *vmResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state VMModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get VMS
	vmHandler := handlers.NewVM(r.client)
	vm, err := vmHandler.GetVM(state.Location.ValueString(), state.VM.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading VM",
			"There was error when reading VM: "+err.Error(),
		)
		return
	}

	// Storage
	var objectStorages basetypes.ObjectValue
	for _, storage := range vm.Storage {
		obj, _ := types.ObjectValue(
			map[string]attr.Type{
				"uuid":       types.StringType,
				"user_id":    types.Int64Type,
				"created_at": types.StringType,
				"name":       types.StringType,
				"primary":    types.BoolType,
				"size":       types.Int64Type,
			},
			map[string]attr.Value{
				"uuid":       types.StringValue(storage.UUID),
				"user_id":    types.Int64Value(int64(storage.UserID)),
				"created_at": types.StringValue(storage.CreatedAt),
				"name":       types.StringValue(storage.Name),
				"primary":    types.BoolValue(storage.Primary),
				"size":       types.Int64Value(int64(storage.Size)),
			},
		)
		objectStorages = obj
	}
	storages, _ := types.ListValue(types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"uuid":       types.StringType,
			"user_id":    types.Int64Type,
			"created_at": types.StringType,
			"name":       types.StringType,
			"primary":    types.BoolType,
			"size":       types.Int64Type,
		},
	}, []attr.Value{objectStorages})

	// General
	state.VM = VMPayloadModel{
		UUID:           types.StringValue(vm.UUID),
		UserID:         types.Int64Value(int64(vm.UserID)),
		CreatedAt:      types.StringValue(vm.CreatedAt),
		UpdatedAt:      types.StringValue(vm.UpdatedAt),
		Backup:         types.BoolValue(false),
		Description:    types.StringValue(vm.Description),
		Name:           types.StringValue(vm.Name),
		OSName:         types.StringValue(vm.OSName),
		OSVersion:      types.StringValue(vm.OSVersion),
		Username:       types.StringValue(vm.Username),
		Memory:         types.Int64Value(int64(vm.Memory)),
		VCPU:           types.Int64Value(int64(vm.VCPU)),
		Status:         types.StringValue(vm.Status),
		PrivateIPV4:    types.StringValue(vm.PrivateIPV4),
		Hostname:       types.StringValue(vm.Hostname),
		BillingAccount: types.Int64Value(int64(vm.BillingAccount)),
		Storage:        storages,
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
func (r *vmResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan VMModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get current state
	var state VMModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	payload := client.VMPayloadUpdate{
		UUID:   state.VM.UUID.ValueString(),
		Name:   plan.VM.Name.ValueString(),
		Memory: int(plan.VM.Memory.ValueInt64()),
		VCPU:   int(plan.VM.VCPU.ValueInt64()),
	}

	// Update/Modify VM
	vmHandler := handlers.NewVM(r.client)
	vm, err := vmHandler.UpdateVM(plan.Location.ValueString(), payload, int(state.VM.Memory.ValueInt64()), int(state.VM.VCPU.ValueInt64()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating VM",
			"There was error when updating VM: \n"+err.Error(),
		)
		return
	}

	// Populate
	if &vm.Errors == nil {
		plan.VM.Name = types.StringValue(vm.Name)
		plan.VM.Memory = types.Int64Value(int64(vm.Memory))
		plan.VM.VCPU = types.Int64Value(int64(vm.VCPU))
	} else {
		plan.VM.Name = types.StringValue(payload.Name)
	}

	// Set State
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
func (r *vmResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Get current state
	var state VMModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	payload := client.VMPayloadDelete{
		UUID: state.VM.UUID.ValueString(),
	}

	// Delete
	vmHandler := handlers.NewVM(r.client)
	err := vmHandler.DeleteVM(state.Location.ValueString(), payload)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting VM",
			"There was error when deleting VM: \n"+err.Error(),
		)
		return
	}
}
