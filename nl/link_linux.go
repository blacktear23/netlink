package nl

const (
	DEFAULT_CHANGE = 0xFFFFFFFF
)

const (
	IFLA_INFO_UNSPEC = iota
	IFLA_INFO_KIND   = iota
	IFLA_INFO_DATA   = iota
	IFLA_INFO_XSTATS = iota
	IFLA_INFO_MAX    = IFLA_INFO_XSTATS
)

const (
	IFLA_VLAN_UNSPEC      = iota
	IFLA_VLAN_ID          = iota
	IFLA_VLAN_FLAGS       = iota
	IFLA_VLAN_EGRESS_QOS  = iota
	IFLA_VLAN_INGRESS_QOS = iota
	IFLA_VLAN_PROTOCOL    = iota
	IFLA_VLAN_MAX         = IFLA_VLAN_PROTOCOL
)

const (
	VETH_INFO_UNSPEC = iota
	VETH_INFO_PEER   = iota
	VETH_INFO_MAX    = VETH_INFO_PEER
)

const (
	IFLA_VXLAN_UNSPEC     = iota
	IFLA_VXLAN_ID         = iota
	IFLA_VXLAN_GROUP      = iota
	IFLA_VXLAN_LINK       = iota
	IFLA_VXLAN_LOCAL      = iota
	IFLA_VXLAN_TTL        = iota
	IFLA_VXLAN_TOS        = iota
	IFLA_VXLAN_LEARNING   = iota
	IFLA_VXLAN_AGEING     = iota
	IFLA_VXLAN_LIMIT      = iota
	IFLA_VXLAN_PORT_RANGE = iota
	IFLA_VXLAN_PROXY      = iota
	IFLA_VXLAN_RSC        = iota
	IFLA_VXLAN_L2MISS     = iota
	IFLA_VXLAN_L3MISS     = iota
	IFLA_VXLAN_PORT       = iota
	IFLA_VXLAN_GROUP6     = iota
	IFLA_VXLAN_LOCAL6     = iota
	IFLA_VXLAN_MAX        = IFLA_VXLAN_LOCAL6
)

const (
	BRIDGE_MODE_UNSPEC  = iota
	BRIDGE_MODE_HAIRPIN = iota
)

const (
	IFLA_BRPORT_UNSPEC        = iota
	IFLA_BRPORT_STATE         = iota
	IFLA_BRPORT_PRIORITY      = iota
	IFLA_BRPORT_COST          = iota
	IFLA_BRPORT_MODE          = iota
	IFLA_BRPORT_GUARD         = iota
	IFLA_BRPORT_PROTECT       = iota
	IFLA_BRPORT_FAST_LEAVE    = iota
	IFLA_BRPORT_LEARNING      = iota
	IFLA_BRPORT_UNICAST_FLOOD = iota
	IFLA_BRPORT_MAX           = IFLA_BRPORT_UNICAST_FLOOD
)

const (
	IFLA_IPVLAN_UNSPEC = iota
	IFLA_IPVLAN_MODE   = iota
	IFLA_IPVLAN_MAX    = IFLA_IPVLAN_MODE
)

const (
	// not defined in syscall
	IFLA_NET_NS_FD = 28
)
