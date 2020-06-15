// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package iso

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer/builder/vsphere/common"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName            *string                  `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType          *string                  `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerDebug                *bool                    `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce                *bool                    `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError              *string                  `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars             map[string]string        `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars        []string                 `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	HTTPDir                    *string                  `mapstructure:"http_directory" cty:"http_directory" hcl:"http_directory"`
	HTTPPortMin                *int                     `mapstructure:"http_port_min" cty:"http_port_min" hcl:"http_port_min"`
	HTTPPortMax                *int                     `mapstructure:"http_port_max" cty:"http_port_max" hcl:"http_port_max"`
	HTTPAddress                *string                  `mapstructure:"http_bind_address" cty:"http_bind_address" hcl:"http_bind_address"`
	VCenterServer              *string                  `mapstructure:"vcenter_server" cty:"vcenter_server" hcl:"vcenter_server"`
	Username                   *string                  `mapstructure:"username" cty:"username" hcl:"username"`
	Password                   *string                  `mapstructure:"password" cty:"password" hcl:"password"`
	InsecureConnection         *bool                    `mapstructure:"insecure_connection" cty:"insecure_connection" hcl:"insecure_connection"`
	Datacenter                 *string                  `mapstructure:"datacenter" cty:"datacenter" hcl:"datacenter"`
	Version                    *uint                    `mapstructure:"vm_version" cty:"vm_version" hcl:"vm_version"`
	GuestOSType                *string                  `mapstructure:"guest_os_type" cty:"guest_os_type" hcl:"guest_os_type"`
	Firmware                   *string                  `mapstructure:"firmware" cty:"firmware" hcl:"firmware"`
	DiskControllerType         *string                  `mapstructure:"disk_controller_type" cty:"disk_controller_type" hcl:"disk_controller_type"`
	Storage                    []FlatDiskConfig         `mapstructure:"storage" cty:"storage" hcl:"storage"`
	NICs                       []FlatNIC                `mapstructure:"network_adapters" cty:"network_adapters" hcl:"network_adapters"`
	USBController              *bool                    `mapstructure:"usb_controller" cty:"usb_controller" hcl:"usb_controller"`
	Notes                      *string                  `mapstructure:"notes" cty:"notes" hcl:"notes"`
	VMName                     *string                  `mapstructure:"vm_name" cty:"vm_name" hcl:"vm_name"`
	Folder                     *string                  `mapstructure:"folder" cty:"folder" hcl:"folder"`
	Cluster                    *string                  `mapstructure:"cluster" cty:"cluster" hcl:"cluster"`
	Host                       *string                  `mapstructure:"host" cty:"host" hcl:"host"`
	ResourcePool               *string                  `mapstructure:"resource_pool" cty:"resource_pool" hcl:"resource_pool"`
	Datastore                  *string                  `mapstructure:"datastore" cty:"datastore" hcl:"datastore"`
	SetHostForDatastoreUploads *bool                    `mapstructure:"set_host_for_datastore_uploads" cty:"set_host_for_datastore_uploads" hcl:"set_host_for_datastore_uploads"`
	CPUs                       *int32                   `mapstructure:"CPUs" cty:"CPUs" hcl:"CPUs"`
	CpuCores                   *int32                   `mapstructure:"cpu_cores" cty:"cpu_cores" hcl:"cpu_cores"`
	CPUReservation             *int64                   `mapstructure:"CPU_reservation" cty:"CPU_reservation" hcl:"CPU_reservation"`
	CPULimit                   *int64                   `mapstructure:"CPU_limit" cty:"CPU_limit" hcl:"CPU_limit"`
	CpuHotAddEnabled           *bool                    `mapstructure:"CPU_hot_plug" cty:"CPU_hot_plug" hcl:"CPU_hot_plug"`
	RAM                        *int64                   `mapstructure:"RAM" cty:"RAM" hcl:"RAM"`
	RAMReservation             *int64                   `mapstructure:"RAM_reservation" cty:"RAM_reservation" hcl:"RAM_reservation"`
	RAMReserveAll              *bool                    `mapstructure:"RAM_reserve_all" cty:"RAM_reserve_all" hcl:"RAM_reserve_all"`
	MemoryHotAddEnabled        *bool                    `mapstructure:"RAM_hot_plug" cty:"RAM_hot_plug" hcl:"RAM_hot_plug"`
	VideoRAM                   *int64                   `mapstructure:"video_ram" cty:"video_ram" hcl:"video_ram"`
	VGPUProfile                *string                  `mapstructure:"vgpu_profile" cty:"vgpu_profile" hcl:"vgpu_profile"`
	NestedHV                   *bool                    `mapstructure:"NestedHV" cty:"NestedHV" hcl:"NestedHV"`
	ConfigParams               map[string]string        `mapstructure:"configuration_parameters" cty:"configuration_parameters" hcl:"configuration_parameters"`
	ToolsSyncTime              *bool                    `mapstructure:"tools_sync_time" cty:"tools_sync_time" hcl:"tools_sync_time"`
	ToolsUpgradePolicy         *bool                    `mapstructure:"tools_upgrade_policy" cty:"tools_upgrade_policy" hcl:"tools_upgrade_policy"`
	ISOChecksum                *string                  `mapstructure:"iso_checksum" required:"true" cty:"iso_checksum" hcl:"iso_checksum"`
	RawSingleISOUrl            *string                  `mapstructure:"iso_url" required:"true" cty:"iso_url" hcl:"iso_url"`
	ISOUrls                    []string                 `mapstructure:"iso_urls" cty:"iso_urls" hcl:"iso_urls"`
	TargetPath                 *string                  `mapstructure:"iso_target_path" cty:"iso_target_path" hcl:"iso_target_path"`
	TargetExtension            *string                  `mapstructure:"iso_target_extension" cty:"iso_target_extension" hcl:"iso_target_extension"`
	CdromType                  *string                  `mapstructure:"cdrom_type" cty:"cdrom_type" hcl:"cdrom_type"`
	ISOPaths                   []string                 `mapstructure:"iso_paths" cty:"iso_paths" hcl:"iso_paths"`
	RemoveCdrom                *bool                    `mapstructure:"remove_cdrom" cty:"remove_cdrom" hcl:"remove_cdrom"`
	FloppyIMGPath              *string                  `mapstructure:"floppy_img_path" cty:"floppy_img_path" hcl:"floppy_img_path"`
	FloppyFiles                []string                 `mapstructure:"floppy_files" cty:"floppy_files" hcl:"floppy_files"`
	FloppyDirectories          []string                 `mapstructure:"floppy_dirs" cty:"floppy_dirs" hcl:"floppy_dirs"`
	FloppyLabel                *string                  `mapstructure:"floppy_label" cty:"floppy_label" hcl:"floppy_label"`
	BootOrder                  *string                  `mapstructure:"boot_order" cty:"boot_order" hcl:"boot_order"`
	BootCommand                []string                 `mapstructure:"boot_command" cty:"boot_command" hcl:"boot_command"`
	BootWait                   *string                  `mapstructure:"boot_wait" cty:"boot_wait" hcl:"boot_wait"`
	HTTPIP                     *string                  `mapstructure:"http_ip" cty:"http_ip" hcl:"http_ip"`
	WaitTimeout                *string                  `mapstructure:"ip_wait_timeout" cty:"ip_wait_timeout" hcl:"ip_wait_timeout"`
	SettleTimeout              *string                  `mapstructure:"ip_settle_timeout" cty:"ip_settle_timeout" hcl:"ip_settle_timeout"`
	WaitAddress                *string                  `mapstructure:"ip_wait_address" cty:"ip_wait_address" hcl:"ip_wait_address"`
	Type                       *string                  `mapstructure:"communicator" cty:"communicator" hcl:"communicator"`
	PauseBeforeConnect         *string                  `mapstructure:"pause_before_connecting" cty:"pause_before_connecting" hcl:"pause_before_connecting"`
	SSHHost                    *string                  `mapstructure:"ssh_host" cty:"ssh_host" hcl:"ssh_host"`
	SSHPort                    *int                     `mapstructure:"ssh_port" cty:"ssh_port" hcl:"ssh_port"`
	SSHUsername                *string                  `mapstructure:"ssh_username" cty:"ssh_username" hcl:"ssh_username"`
	SSHPassword                *string                  `mapstructure:"ssh_password" cty:"ssh_password" hcl:"ssh_password"`
	SSHKeyPairName             *string                  `mapstructure:"ssh_keypair_name" cty:"ssh_keypair_name" hcl:"ssh_keypair_name"`
	SSHTemporaryKeyPairName    *string                  `mapstructure:"temporary_key_pair_name" cty:"temporary_key_pair_name" hcl:"temporary_key_pair_name"`
	SSHClearAuthorizedKeys     *bool                    `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys" hcl:"ssh_clear_authorized_keys"`
	SSHPrivateKeyFile          *string                  `mapstructure:"ssh_private_key_file" cty:"ssh_private_key_file" hcl:"ssh_private_key_file"`
	SSHPty                     *bool                    `mapstructure:"ssh_pty" cty:"ssh_pty" hcl:"ssh_pty"`
	SSHTimeout                 *string                  `mapstructure:"ssh_timeout" cty:"ssh_timeout" hcl:"ssh_timeout"`
	SSHWaitTimeout             *string                  `mapstructure:"ssh_wait_timeout" undocumented:"true" cty:"ssh_wait_timeout" hcl:"ssh_wait_timeout"`
	SSHAgentAuth               *bool                    `mapstructure:"ssh_agent_auth" cty:"ssh_agent_auth" hcl:"ssh_agent_auth"`
	SSHDisableAgentForwarding  *bool                    `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding" hcl:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts       *int                     `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts" hcl:"ssh_handshake_attempts"`
	SSHBastionHost             *string                  `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host" hcl:"ssh_bastion_host"`
	SSHBastionPort             *int                     `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port" hcl:"ssh_bastion_port"`
	SSHBastionAgentAuth        *bool                    `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth" hcl:"ssh_bastion_agent_auth"`
	SSHBastionUsername         *string                  `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username" hcl:"ssh_bastion_username"`
	SSHBastionPassword         *string                  `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password" hcl:"ssh_bastion_password"`
	SSHBastionInteractive      *bool                    `mapstructure:"ssh_bastion_interactive" cty:"ssh_bastion_interactive" hcl:"ssh_bastion_interactive"`
	SSHBastionPrivateKeyFile   *string                  `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file" hcl:"ssh_bastion_private_key_file"`
	SSHFileTransferMethod      *string                  `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method" hcl:"ssh_file_transfer_method"`
	SSHProxyHost               *string                  `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host" hcl:"ssh_proxy_host"`
	SSHProxyPort               *int                     `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port" hcl:"ssh_proxy_port"`
	SSHProxyUsername           *string                  `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username" hcl:"ssh_proxy_username"`
	SSHProxyPassword           *string                  `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password" hcl:"ssh_proxy_password"`
	SSHKeepAliveInterval       *string                  `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval" hcl:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout        *string                  `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout" hcl:"ssh_read_write_timeout"`
	SSHRemoteTunnels           []string                 `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels" hcl:"ssh_remote_tunnels"`
	SSHLocalTunnels            []string                 `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels" hcl:"ssh_local_tunnels"`
	SSHPublicKey               []byte                   `mapstructure:"ssh_public_key" cty:"ssh_public_key" hcl:"ssh_public_key"`
	SSHPrivateKey              []byte                   `mapstructure:"ssh_private_key" cty:"ssh_private_key" hcl:"ssh_private_key"`
	WinRMUser                  *string                  `mapstructure:"winrm_username" cty:"winrm_username" hcl:"winrm_username"`
	WinRMPassword              *string                  `mapstructure:"winrm_password" cty:"winrm_password" hcl:"winrm_password"`
	WinRMHost                  *string                  `mapstructure:"winrm_host" cty:"winrm_host" hcl:"winrm_host"`
	WinRMNoProxy               *bool                    `mapstructure:"winrm_no_proxy" cty:"winrm_no_proxy" hcl:"winrm_no_proxy"`
	WinRMPort                  *int                     `mapstructure:"winrm_port" cty:"winrm_port" hcl:"winrm_port"`
	WinRMTimeout               *string                  `mapstructure:"winrm_timeout" cty:"winrm_timeout" hcl:"winrm_timeout"`
	WinRMUseSSL                *bool                    `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl" hcl:"winrm_use_ssl"`
	WinRMInsecure              *bool                    `mapstructure:"winrm_insecure" cty:"winrm_insecure" hcl:"winrm_insecure"`
	WinRMUseNTLM               *bool                    `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm" hcl:"winrm_use_ntlm"`
	Command                    *string                  `mapstructure:"shutdown_command" cty:"shutdown_command" hcl:"shutdown_command"`
	Timeout                    *string                  `mapstructure:"shutdown_timeout" cty:"shutdown_timeout" hcl:"shutdown_timeout"`
	DisableShutdown            *bool                    `mapstructure:"disable_shutdown" cty:"disable_shutdown" hcl:"disable_shutdown"`
	CreateSnapshot             *bool                    `mapstructure:"create_snapshot" cty:"create_snapshot" hcl:"create_snapshot"`
	ConvertToTemplate          *bool                    `mapstructure:"convert_to_template" cty:"convert_to_template" hcl:"convert_to_template"`
	Export                     *common.FlatExportConfig `mapstructure:"export" cty:"export" hcl:"export"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":              &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":            &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":                   &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                   &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":                &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":          &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables":     &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"http_directory":                 &hcldec.AttrSpec{Name: "http_directory", Type: cty.String, Required: false},
		"http_port_min":                  &hcldec.AttrSpec{Name: "http_port_min", Type: cty.Number, Required: false},
		"http_port_max":                  &hcldec.AttrSpec{Name: "http_port_max", Type: cty.Number, Required: false},
		"http_bind_address":              &hcldec.AttrSpec{Name: "http_bind_address", Type: cty.String, Required: false},
		"vcenter_server":                 &hcldec.AttrSpec{Name: "vcenter_server", Type: cty.String, Required: false},
		"username":                       &hcldec.AttrSpec{Name: "username", Type: cty.String, Required: false},
		"password":                       &hcldec.AttrSpec{Name: "password", Type: cty.String, Required: false},
		"insecure_connection":            &hcldec.AttrSpec{Name: "insecure_connection", Type: cty.Bool, Required: false},
		"datacenter":                     &hcldec.AttrSpec{Name: "datacenter", Type: cty.String, Required: false},
		"vm_version":                     &hcldec.AttrSpec{Name: "vm_version", Type: cty.Number, Required: false},
		"guest_os_type":                  &hcldec.AttrSpec{Name: "guest_os_type", Type: cty.String, Required: false},
		"firmware":                       &hcldec.AttrSpec{Name: "firmware", Type: cty.String, Required: false},
		"disk_controller_type":           &hcldec.AttrSpec{Name: "disk_controller_type", Type: cty.String, Required: false},
		"storage":                        &hcldec.BlockListSpec{TypeName: "storage", Nested: hcldec.ObjectSpec((*FlatDiskConfig)(nil).HCL2Spec())},
		"network_adapters":               &hcldec.BlockListSpec{TypeName: "network_adapters", Nested: hcldec.ObjectSpec((*FlatNIC)(nil).HCL2Spec())},
		"usb_controller":                 &hcldec.AttrSpec{Name: "usb_controller", Type: cty.Bool, Required: false},
		"notes":                          &hcldec.AttrSpec{Name: "notes", Type: cty.String, Required: false},
		"vm_name":                        &hcldec.AttrSpec{Name: "vm_name", Type: cty.String, Required: false},
		"folder":                         &hcldec.AttrSpec{Name: "folder", Type: cty.String, Required: false},
		"cluster":                        &hcldec.AttrSpec{Name: "cluster", Type: cty.String, Required: false},
		"host":                           &hcldec.AttrSpec{Name: "host", Type: cty.String, Required: false},
		"resource_pool":                  &hcldec.AttrSpec{Name: "resource_pool", Type: cty.String, Required: false},
		"datastore":                      &hcldec.AttrSpec{Name: "datastore", Type: cty.String, Required: false},
		"set_host_for_datastore_uploads": &hcldec.AttrSpec{Name: "set_host_for_datastore_uploads", Type: cty.Bool, Required: false},
		"CPUs":                           &hcldec.AttrSpec{Name: "CPUs", Type: cty.Number, Required: false},
		"cpu_cores":                      &hcldec.AttrSpec{Name: "cpu_cores", Type: cty.Number, Required: false},
		"CPU_reservation":                &hcldec.AttrSpec{Name: "CPU_reservation", Type: cty.Number, Required: false},
		"CPU_limit":                      &hcldec.AttrSpec{Name: "CPU_limit", Type: cty.Number, Required: false},
		"CPU_hot_plug":                   &hcldec.AttrSpec{Name: "CPU_hot_plug", Type: cty.Bool, Required: false},
		"RAM":                            &hcldec.AttrSpec{Name: "RAM", Type: cty.Number, Required: false},
		"RAM_reservation":                &hcldec.AttrSpec{Name: "RAM_reservation", Type: cty.Number, Required: false},
		"RAM_reserve_all":                &hcldec.AttrSpec{Name: "RAM_reserve_all", Type: cty.Bool, Required: false},
		"RAM_hot_plug":                   &hcldec.AttrSpec{Name: "RAM_hot_plug", Type: cty.Bool, Required: false},
		"video_ram":                      &hcldec.AttrSpec{Name: "video_ram", Type: cty.Number, Required: false},
		"vgpu_profile":                   &hcldec.AttrSpec{Name: "vgpu_profile", Type: cty.String, Required: false},
		"NestedHV":                       &hcldec.AttrSpec{Name: "NestedHV", Type: cty.Bool, Required: false},
		"configuration_parameters":       &hcldec.AttrSpec{Name: "configuration_parameters", Type: cty.Map(cty.String), Required: false},
		"tools_sync_time":                &hcldec.AttrSpec{Name: "tools_sync_time", Type: cty.Bool, Required: false},
		"tools_upgrade_policy":           &hcldec.AttrSpec{Name: "tools_upgrade_policy", Type: cty.Bool, Required: false},
		"iso_checksum":                   &hcldec.AttrSpec{Name: "iso_checksum", Type: cty.String, Required: false},
		"iso_url":                        &hcldec.AttrSpec{Name: "iso_url", Type: cty.String, Required: false},
		"iso_urls":                       &hcldec.AttrSpec{Name: "iso_urls", Type: cty.List(cty.String), Required: false},
		"iso_target_path":                &hcldec.AttrSpec{Name: "iso_target_path", Type: cty.String, Required: false},
		"iso_target_extension":           &hcldec.AttrSpec{Name: "iso_target_extension", Type: cty.String, Required: false},
		"cdrom_type":                     &hcldec.AttrSpec{Name: "cdrom_type", Type: cty.String, Required: false},
		"iso_paths":                      &hcldec.AttrSpec{Name: "iso_paths", Type: cty.List(cty.String), Required: false},
		"remove_cdrom":                   &hcldec.AttrSpec{Name: "remove_cdrom", Type: cty.Bool, Required: false},
		"floppy_img_path":                &hcldec.AttrSpec{Name: "floppy_img_path", Type: cty.String, Required: false},
		"floppy_files":                   &hcldec.AttrSpec{Name: "floppy_files", Type: cty.List(cty.String), Required: false},
		"floppy_dirs":                    &hcldec.AttrSpec{Name: "floppy_dirs", Type: cty.List(cty.String), Required: false},
		"floppy_label":                   &hcldec.AttrSpec{Name: "floppy_label", Type: cty.String, Required: false},
		"boot_order":                     &hcldec.AttrSpec{Name: "boot_order", Type: cty.String, Required: false},
		"boot_command":                   &hcldec.AttrSpec{Name: "boot_command", Type: cty.List(cty.String), Required: false},
		"boot_wait":                      &hcldec.AttrSpec{Name: "boot_wait", Type: cty.String, Required: false},
		"http_ip":                        &hcldec.AttrSpec{Name: "http_ip", Type: cty.String, Required: false},
		"ip_wait_timeout":                &hcldec.AttrSpec{Name: "ip_wait_timeout", Type: cty.String, Required: false},
		"ip_settle_timeout":              &hcldec.AttrSpec{Name: "ip_settle_timeout", Type: cty.String, Required: false},
		"ip_wait_address":                &hcldec.AttrSpec{Name: "ip_wait_address", Type: cty.String, Required: false},
		"communicator":                   &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":        &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                       &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_port":                       &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                   &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"ssh_password":                   &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_keypair_name":               &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"temporary_key_pair_name":        &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"ssh_clear_authorized_keys":      &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_private_key_file":           &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_pty":                        &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                    &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"ssh_wait_timeout":               &hcldec.AttrSpec{Name: "ssh_wait_timeout", Type: cty.String, Required: false},
		"ssh_agent_auth":                 &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"ssh_disable_agent_forwarding":   &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"ssh_handshake_attempts":         &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"ssh_bastion_host":               &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"ssh_bastion_port":               &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"ssh_bastion_agent_auth":         &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"ssh_bastion_username":           &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"ssh_bastion_password":           &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"ssh_bastion_interactive":        &hcldec.AttrSpec{Name: "ssh_bastion_interactive", Type: cty.Bool, Required: false},
		"ssh_bastion_private_key_file":   &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"ssh_file_transfer_method":       &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"ssh_proxy_host":                 &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"ssh_proxy_port":                 &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"ssh_proxy_username":             &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"ssh_proxy_password":             &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"ssh_keep_alive_interval":        &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"ssh_read_write_timeout":         &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"ssh_remote_tunnels":             &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_local_tunnels":              &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_public_key":                 &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"ssh_private_key":                &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"winrm_username":                 &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"winrm_password":                 &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"winrm_host":                     &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"winrm_no_proxy":                 &hcldec.AttrSpec{Name: "winrm_no_proxy", Type: cty.Bool, Required: false},
		"winrm_port":                     &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                  &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                  &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":                 &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":                 &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
		"shutdown_command":               &hcldec.AttrSpec{Name: "shutdown_command", Type: cty.String, Required: false},
		"shutdown_timeout":               &hcldec.AttrSpec{Name: "shutdown_timeout", Type: cty.String, Required: false},
		"disable_shutdown":               &hcldec.AttrSpec{Name: "disable_shutdown", Type: cty.Bool, Required: false},
		"create_snapshot":                &hcldec.AttrSpec{Name: "create_snapshot", Type: cty.Bool, Required: false},
		"convert_to_template":            &hcldec.AttrSpec{Name: "convert_to_template", Type: cty.Bool, Required: false},
		"export":                         &hcldec.BlockSpec{TypeName: "export", Nested: hcldec.ObjectSpec((*common.FlatExportConfig)(nil).HCL2Spec())},
	}
	return s
}
