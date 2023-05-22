package client

// Virtual Machine Model
type VMStorageInfo struct {
	UUID      string `json:"uuid"`
	UserID    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
	Name      string `json:"name"`
	Primary   bool   `json:"primary"`
	Size      int    `json:"size"`
}

type VMPayload struct {
	Backup           bool   `json:"backup"`
	BillingAccountID int    `json:"billing_account_id"`
	Description      string `json:"description,omitempty"`
	Name             string `json:"name"`
	OSName           string `json:"os_name"`
	OSVersion        string `json:"os_version"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Disks            int    `json:"disks"`
	Memory           int    `json:"ram"`
	VCPU             int    `json:"vcpu"`
	ReservePublicIP  bool   `json:"reserve_public_ip,omitempty"`
	NetworkUUID      string `json:"network_uuid,omitempty"`
	CloudInit        string `json:"cloud_init,omitempty"`
	PublicKey        string `json:"public_key,omitempty"`
}

type VMPayloadUpdate struct {
	UUID   string `json:"uuid"`
	Name   string `json:"name"`
	Memory int    `json:"ram"`
	VCPU   int    `json:"vcpu"`
}

type VMPayloadDelete struct {
	UUID string `json:"uuid"`
}

type VMResponse struct {
	UUID           string                 `json:"uuid"`
	UserID         int                    `json:"user_id"`
	Backup         bool                   `json:"backup"`
	BillingAccount int                    `json:"billing_account"`
	CreatedAt      string                 `json:"created_at"`
	UpdatedAt      string                 `json:"updated_at"`
	Description    string                 `json:"description,omitempty"`
	Hostname       string                 `json:"hostname"`
	Name           string                 `json:"name"`
	OSName         string                 `json:"os_name"`
	OSVersion      string                 `json:"os_version"`
	Username       string                 `json:"username"`
	Memory         int                    `json:"ram"`
	VCPU           int                    `json:"vcpu"`
	Storage        []VMStorageInfo        `json:"storage,omitempty"`
	PrivateIPV4    string                 `json:"private_ipv4"`
	Status         string                 `json:"status"`
	Errors         map[string]interface{} `json:"errors"`
}

// Location Model
type LocationInfo struct {
	DisplayName string `json:"display_name"`
}

// Floating IP Model
type FloatingIPInfo struct {
	AssignedTo *string `json:"assigned_to"`
	Address    string  `json:"address"`
}
