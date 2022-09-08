// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

type VirtualNetworkGateway_STATUSARM struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// ExtendedLocation: The extended location of type local virtual network gateway.
	ExtendedLocation *ExtendedLocation_STATUSARM `json:"extendedLocation,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the virtual network gateway.
	Properties *VirtualNetworkGatewayPropertiesFormat_STATUSARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

type VirtualNetworkGatewayPropertiesFormat_STATUSARM struct {
	// ActiveActive: ActiveActive flag.
	ActiveActive *bool `json:"activeActive,omitempty"`

	// BgpSettings: Virtual network gateway's BGP speaker settings.
	BgpSettings *BgpSettings_STATUSARM `json:"bgpSettings,omitempty"`

	// CustomRoutes: The reference to the address space resource which represents the custom routes address space specified by
	// the customer for virtual network gateway and VpnClient.
	CustomRoutes *AddressSpace_STATUSARM `json:"customRoutes,omitempty"`

	// EnableBgp: Whether BGP is enabled for this virtual network gateway or not.
	EnableBgp *bool `json:"enableBgp,omitempty"`

	// EnableDnsForwarding: Whether dns forwarding is enabled or not.
	EnableDnsForwarding *bool `json:"enableDnsForwarding,omitempty"`

	// EnablePrivateIpAddress: Whether private IP needs to be enabled on this gateway for connections or not.
	EnablePrivateIpAddress *bool `json:"enablePrivateIpAddress,omitempty"`

	// GatewayDefaultSite: The reference to the LocalNetworkGateway resource which represents local network site having default
	// routes. Assign Null value in case of removing existing default site setting.
	GatewayDefaultSite *SubResource_STATUSARM `json:"gatewayDefaultSite,omitempty"`

	// GatewayType: The type of this virtual network gateway.
	GatewayType *VirtualNetworkGatewayPropertiesFormat_STATUS_GatewayType `json:"gatewayType,omitempty"`

	// InboundDnsForwardingEndpoint: The IP address allocated by the gateway to which dns requests can be sent.
	InboundDnsForwardingEndpoint *string `json:"inboundDnsForwardingEndpoint,omitempty"`

	// IpConfigurations: IP configurations for virtual network gateway.
	IpConfigurations []VirtualNetworkGatewayIPConfiguration_STATUSARM `json:"ipConfigurations,omitempty"`

	// ProvisioningState: The provisioning state of the virtual network gateway resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// ResourceGuid: The resource GUID property of the virtual network gateway resource.
	ResourceGuid *string `json:"resourceGuid,omitempty"`

	// Sku: The reference to the VirtualNetworkGatewaySku resource which represents the SKU selected for Virtual network
	// gateway.
	Sku *VirtualNetworkGatewaySku_STATUSARM `json:"sku,omitempty"`

	// VNetExtendedLocationResourceId: Customer vnet resource id. VirtualNetworkGateway of type local gateway is associated
	// with the customer vnet.
	VNetExtendedLocationResourceId *string `json:"vNetExtendedLocationResourceId,omitempty"`

	// VpnClientConfiguration: The reference to the VpnClientConfiguration resource which represents the P2S VpnClient
	// configurations.
	VpnClientConfiguration *VpnClientConfiguration_STATUSARM `json:"vpnClientConfiguration,omitempty"`

	// VpnGatewayGeneration: The generation for this VirtualNetworkGateway. Must be None if gatewayType is not VPN.
	VpnGatewayGeneration *VirtualNetworkGatewayPropertiesFormat_STATUS_VpnGatewayGeneration `json:"vpnGatewayGeneration,omitempty"`

	// VpnType: The type of this virtual network gateway.
	VpnType *VirtualNetworkGatewayPropertiesFormat_STATUS_VpnType `json:"vpnType,omitempty"`
}

type AddressSpace_STATUSARM struct {
	// AddressPrefixes: A list of address blocks reserved for this virtual network in CIDR notation.
	AddressPrefixes []string `json:"addressPrefixes,omitempty"`
}

type BgpSettings_STATUSARM struct {
	// Asn: The BGP speaker's ASN.
	Asn *uint32 `json:"asn,omitempty"`

	// BgpPeeringAddress: The BGP peering address and BGP identifier of this BGP speaker.
	BgpPeeringAddress *string `json:"bgpPeeringAddress,omitempty"`

	// BgpPeeringAddresses: BGP peering address with IP configuration ID for virtual network gateway.
	BgpPeeringAddresses []IPConfigurationBgpPeeringAddress_STATUSARM `json:"bgpPeeringAddresses,omitempty"`

	// PeerWeight: The weight added to routes learned from this BGP speaker.
	PeerWeight *int `json:"peerWeight,omitempty"`
}

type VirtualNetworkGatewayIPConfiguration_STATUSARM struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the virtual network gateway ip configuration.
	Properties *VirtualNetworkGatewayIPConfigurationPropertiesFormat_STATUSARM `json:"properties,omitempty"`
}

type VirtualNetworkGatewaySku_STATUSARM struct {
	// Capacity: The capacity.
	Capacity *int `json:"capacity,omitempty"`

	// Name: Gateway SKU name.
	Name *VirtualNetworkGatewaySku_STATUS_Name `json:"name,omitempty"`

	// Tier: Gateway SKU tier.
	Tier *VirtualNetworkGatewaySku_STATUS_Tier `json:"tier,omitempty"`
}

type VpnClientConfiguration_STATUSARM struct {
	// AadAudience: The AADAudience property of the VirtualNetworkGateway resource for vpn client connection used for AAD
	// authentication.
	AadAudience *string `json:"aadAudience,omitempty"`

	// AadIssuer: The AADIssuer property of the VirtualNetworkGateway resource for vpn client connection used for AAD
	// authentication.
	AadIssuer *string `json:"aadIssuer,omitempty"`

	// AadTenant: The AADTenant property of the VirtualNetworkGateway resource for vpn client connection used for AAD
	// authentication.
	AadTenant *string `json:"aadTenant,omitempty"`

	// RadiusServerAddress: The radius server address property of the VirtualNetworkGateway resource for vpn client connection.
	RadiusServerAddress *string `json:"radiusServerAddress,omitempty"`

	// RadiusServerSecret: The radius secret property of the VirtualNetworkGateway resource for vpn client connection.
	RadiusServerSecret *string `json:"radiusServerSecret,omitempty"`

	// RadiusServers: The radiusServers property for multiple radius server configuration.
	RadiusServers []RadiusServer_STATUSARM `json:"radiusServers,omitempty"`

	// VpnAuthenticationTypes: VPN authentication types for the virtual network gateway..
	VpnAuthenticationTypes []VpnClientConfiguration_STATUS_VpnAuthenticationTypes `json:"vpnAuthenticationTypes,omitempty"`

	// VpnClientAddressPool: The reference to the address space resource which represents Address space for P2S VpnClient.
	VpnClientAddressPool *AddressSpace_STATUSARM `json:"vpnClientAddressPool,omitempty"`

	// VpnClientIpsecPolicies: VpnClientIpsecPolicies for virtual network gateway P2S client.
	VpnClientIpsecPolicies []IpsecPolicy_STATUSARM `json:"vpnClientIpsecPolicies,omitempty"`

	// VpnClientProtocols: VpnClientProtocols for Virtual network gateway.
	VpnClientProtocols []VpnClientConfiguration_STATUS_VpnClientProtocols `json:"vpnClientProtocols,omitempty"`

	// VpnClientRevokedCertificates: VpnClientRevokedCertificate for Virtual network gateway.
	VpnClientRevokedCertificates []VpnClientRevokedCertificate_STATUSARM `json:"vpnClientRevokedCertificates,omitempty"`

	// VpnClientRootCertificates: VpnClientRootCertificate for virtual network gateway.
	VpnClientRootCertificates []VpnClientRootCertificate_STATUSARM `json:"vpnClientRootCertificates,omitempty"`
}

type IPConfigurationBgpPeeringAddress_STATUSARM struct {
	// CustomBgpIpAddresses: The list of custom BGP peering addresses which belong to IP configuration.
	CustomBgpIpAddresses []string `json:"customBgpIpAddresses,omitempty"`

	// DefaultBgpIpAddresses: The list of default BGP peering addresses which belong to IP configuration.
	DefaultBgpIpAddresses []string `json:"defaultBgpIpAddresses,omitempty"`

	// IpconfigurationId: The ID of IP configuration which belongs to gateway.
	IpconfigurationId *string `json:"ipconfigurationId,omitempty"`

	// TunnelIpAddresses: The list of tunnel public IP addresses which belong to IP configuration.
	TunnelIpAddresses []string `json:"tunnelIpAddresses,omitempty"`
}

type IpsecPolicy_STATUSARM struct {
	// DhGroup: The DH Group used in IKE Phase 1 for initial SA.
	DhGroup *DhGroup_STATUS `json:"dhGroup,omitempty"`

	// IkeEncryption: The IKE encryption algorithm (IKE phase 2).
	IkeEncryption *IkeEncryption_STATUS `json:"ikeEncryption,omitempty"`

	// IkeIntegrity: The IKE integrity algorithm (IKE phase 2).
	IkeIntegrity *IkeIntegrity_STATUS `json:"ikeIntegrity,omitempty"`

	// IpsecEncryption: The IPSec encryption algorithm (IKE phase 1).
	IpsecEncryption *IpsecEncryption_STATUS `json:"ipsecEncryption,omitempty"`

	// IpsecIntegrity: The IPSec integrity algorithm (IKE phase 1).
	IpsecIntegrity *IpsecIntegrity_STATUS `json:"ipsecIntegrity,omitempty"`

	// PfsGroup: The Pfs Group used in IKE Phase 2 for new child SA.
	PfsGroup *PfsGroup_STATUS `json:"pfsGroup,omitempty"`

	// SaDataSizeKilobytes: The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for a site
	// to site VPN tunnel.
	SaDataSizeKilobytes *int `json:"saDataSizeKilobytes,omitempty"`

	// SaLifeTimeSeconds: The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for a site
	// to site VPN tunnel.
	SaLifeTimeSeconds *int `json:"saLifeTimeSeconds,omitempty"`
}

type RadiusServer_STATUSARM struct {
	// RadiusServerAddress: The address of this radius server.
	RadiusServerAddress *string `json:"radiusServerAddress,omitempty"`

	// RadiusServerScore: The initial score assigned to this radius server.
	RadiusServerScore *int `json:"radiusServerScore,omitempty"`

	// RadiusServerSecret: The secret used for this radius server.
	RadiusServerSecret *string `json:"radiusServerSecret,omitempty"`
}

type VirtualNetworkGatewayIPConfigurationPropertiesFormat_STATUSARM struct {
	// PrivateIPAddress: Private IP Address for this gateway.
	PrivateIPAddress *string `json:"privateIPAddress,omitempty"`

	// PrivateIPAllocationMethod: The private IP address allocation method.
	PrivateIPAllocationMethod *IPAllocationMethod_STATUS `json:"privateIPAllocationMethod,omitempty"`

	// ProvisioningState: The provisioning state of the virtual network gateway IP configuration resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// PublicIPAddress: The reference to the public IP resource.
	PublicIPAddress *SubResource_STATUSARM `json:"publicIPAddress,omitempty"`

	// Subnet: The reference to the subnet resource.
	Subnet *SubResource_STATUSARM `json:"subnet,omitempty"`
}

type VpnClientRevokedCertificate_STATUSARM struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the vpn client revoked certificate.
	Properties *VpnClientRevokedCertificatePropertiesFormat_STATUSARM `json:"properties,omitempty"`
}

type VpnClientRootCertificate_STATUSARM struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the vpn client root certificate.
	Properties *VpnClientRootCertificatePropertiesFormat_STATUSARM `json:"properties,omitempty"`
}

type VpnClientRevokedCertificatePropertiesFormat_STATUSARM struct {
	// ProvisioningState: The provisioning state of the VPN client revoked certificate resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// Thumbprint: The revoked VPN client certificate thumbprint.
	Thumbprint *string `json:"thumbprint,omitempty"`
}

type VpnClientRootCertificatePropertiesFormat_STATUSARM struct {
	// ProvisioningState: The provisioning state of the VPN client root certificate resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// PublicCertData: The certificate public data.
	PublicCertData *string `json:"publicCertData,omitempty"`
}