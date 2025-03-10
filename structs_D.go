/*
Package goysrl is a generated package which contains definitions
of structs which represent a YANG schema. The generated schema can be
compressed by a series of transformations (compression was false
in this case).

This package was generated by /Users/henderiw/CodeProjects/go/pkg/mod/github.com/openconfig/ygot@v0.8.3/genutil/names.go
using the following YANG input files:
	- yang/srl/srl_nokia-aaa-tacacs.yang
	- yang/srl/srl_nokia-aaa-types.yang
	- yang/srl/srl_nokia-aaa.yang
	- yang/srl/srl_nokia-acl.yang
	- yang/srl/srl_nokia-aft.yang
	- yang/srl/srl_nokia-aggregate-routes.yang
	- yang/srl/srl_nokia-app-mgmt.yang
	- yang/srl/srl_nokia-bfd.yang
	- yang/srl/srl_nokia-bgp.yang
	- yang/srl/srl_nokia-boot.yang
	- yang/srl/srl_nokia-bridge-table-mac-duplication-entries.yang
	- yang/srl/srl_nokia-bridge-table-mac-duplication.yang
	- yang/srl/srl_nokia-bridge-table-mac-learning-entries.yang
	- yang/srl/srl_nokia-bridge-table-mac-learning.yang
	- yang/srl/srl_nokia-bridge-table-mac-limit.yang
	- yang/srl/srl_nokia-bridge-table-mac-table.yang
	- yang/srl/srl_nokia-bridge-table-static-mac.yang
	- yang/srl/srl_nokia-bridge-table.yang
	- yang/srl/srl_nokia-common.yang
	- yang/srl/srl_nokia-configuration.yang
	- yang/srl/srl_nokia-dns.yang
	- yang/srl/srl_nokia-features.yang
	- yang/srl/srl_nokia-ftp.yang
	- yang/srl/srl_nokia-gnmi-server.yang
	- yang/srl/srl_nokia-icmp.yang
	- yang/srl/srl_nokia-if-ip.yang
	- yang/srl/srl_nokia-interfaces-bridge-table-mac-duplication-entries.yang
	- yang/srl/srl_nokia-interfaces-bridge-table-mac-learning-entries.yang
	- yang/srl/srl_nokia-interfaces-bridge-table-mac-table.yang
	- yang/srl/srl_nokia-interfaces-bridge-table-statistics.yang
	- yang/srl/srl_nokia-interfaces-bridge-table.yang
	- yang/srl/srl_nokia-interfaces-ip-dhcp-relay.yang
	- yang/srl/srl_nokia-interfaces-ip-dhcp.yang
	- yang/srl/srl_nokia-interfaces-ip-vrrp.yang
	- yang/srl/srl_nokia-interfaces-lag.yang
	- yang/srl/srl_nokia-interfaces-nbr.yang
	- yang/srl/srl_nokia-interfaces-router-adv.yang
	- yang/srl/srl_nokia-interfaces-vlans.yang
	- yang/srl/srl_nokia-interfaces.yang
	- yang/srl/srl_nokia-ip-route-tables.yang
	- yang/srl/srl_nokia-isis.yang
	- yang/srl/srl_nokia-json-rpc.yang
	- yang/srl/srl_nokia-keychains.yang
	- yang/srl/srl_nokia-lacp.yang
	- yang/srl/srl_nokia-linux.yang
	- yang/srl/srl_nokia-lldp-types.yang
	- yang/srl/srl_nokia-lldp.yang
	- yang/srl/srl_nokia-load-balancing.yang
	- yang/srl/srl_nokia-logging.yang
	- yang/srl/srl_nokia-maintenance-mode.yang
	- yang/srl/srl_nokia-micro-bfd.yang
	- yang/srl/srl_nokia-mpls-route-tables.yang
	- yang/srl/srl_nokia-mpls.yang
	- yang/srl/srl_nokia-mtu.yang
	- yang/srl/srl_nokia-network-instance-mtu.yang
	- yang/srl/srl_nokia-network-instance.yang
	- yang/srl/srl_nokia-next-hop-groups.yang
	- yang/srl/srl_nokia-ntp.yang
	- yang/srl/srl_nokia-ospf-lsdb.yang
	- yang/srl/srl_nokia-ospf-types.yang
	- yang/srl/srl_nokia-ospf.yang
	- yang/srl/srl_nokia-ospfv3-lsas.yang
	- yang/srl/srl_nokia-packet-match-types.yang
	- yang/srl/srl_nokia-platform-acl.yang
	- yang/srl/srl_nokia-platform-chassis.yang
	- yang/srl/srl_nokia-platform-control.yang
	- yang/srl/srl_nokia-platform-cpu.yang
	- yang/srl/srl_nokia-platform-disk.yang
	- yang/srl/srl_nokia-platform-fabric.yang
	- yang/srl/srl_nokia-platform-fan.yang
	- yang/srl/srl_nokia-platform-ip-mpls-fwd-resources.yang
	- yang/srl/srl_nokia-platform-lc.yang
	- yang/srl/srl_nokia-platform-memory.yang
	- yang/srl/srl_nokia-platform-mtu.yang
	- yang/srl/srl_nokia-platform-psu.yang
	- yang/srl/srl_nokia-platform-qos.yang
	- yang/srl/srl_nokia-platform-redundancy.yang
	- yang/srl/srl_nokia-platform-resource-monitoring.yang
	- yang/srl/srl_nokia-platform.yang
	- yang/srl/srl_nokia-policy-types.yang
	- yang/srl/srl_nokia-qos.yang
	- yang/srl/srl_nokia-rib-bgp.yang
	- yang/srl/srl_nokia-routing-policy.yang
	- yang/srl/srl_nokia-sflow.yang
	- yang/srl/srl_nokia-snmp-trace.yang
	- yang/srl/srl_nokia-snmp.yang
	- yang/srl/srl_nokia-ssh.yang
	- yang/srl/srl_nokia-static-routes.yang
	- yang/srl/srl_nokia-system-banner.yang
	- yang/srl/srl_nokia-system-bridge-table.yang
	- yang/srl/srl_nokia-system-info.yang
	- yang/srl/srl_nokia-system-name.yang
	- yang/srl/srl_nokia-system.yang
	- yang/srl/srl_nokia-tcp-udp.yang
	- yang/srl/srl_nokia-timezone.yang
	- yang/srl/srl_nokia-tls.yang
Imported modules were sourced from:
	- yang/ietf/...
*/
package goysrl

// Device represents the /device YANG schema element.
type Device struct {
	Acl             *SrlNokiaAcl_Acl                                    `path:"acl" module:"srl_nokia-acl"`
	Bfd             *SrlNokiaBfd_Bfd                                    `path:"bfd" module:"srl_nokia-bfd"`
	Interface       map[string]*SrlNokiaInterfaces_Interface            `path:"interface" module:"srl_nokia-interfaces"`
	NetworkInstance map[string]*SrlNokiaNetworkInstance_NetworkInstance `path:"network-instance" module:"srl_nokia-network-instance"`
	Platform        *SrlNokiaPlatform_Platform                          `path:"platform" module:"srl_nokia-platform"`
	Qos             *SrlNokiaQos_Qos                                    `path:"qos" module:"srl_nokia-qos"`
	RoutingPolicy   *SrlNokiaRoutingPolicy_RoutingPolicy                `path:"routing-policy" module:"srl_nokia-routing-policy"`
	System          *SrlNokiaSystem_System                              `path:"system" module:"srl_nokia-system"`
}

// IsYANGGoStruct ensures that Device implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Device) IsYANGGoStruct() {}
