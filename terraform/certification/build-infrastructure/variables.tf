variable "ipv4_mgmt_network" {
  description = "Comma separated network subnets in IPv4 CIDR format used for management"
  type        = string
}

variable "ipv6_mgmt_network" {
  description = "Comma separated network subnets in IPv6 CIDR format used for management"
  type        = string
}
