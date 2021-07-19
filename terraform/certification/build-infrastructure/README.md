# Create an EC2 instance in AWS
## Overview
Utilizes terraform to create and Amazon Linux 2 instance in AWS with supporting security group in the default VPC (sufficient for the purpose of the exercise but not recommended otherwise)

## Usage
1. Run `terraform init` to initialize the environment
1. Run `terraform plan -var 'ipv6_mgmt_network=<Your comma separated IPv6 CIDR blocks here>' -var 'ipv4_mgmt_network=<Your comma separated IPv4 CIDR blocks here>'` to see the build plan of intent
1. Run `terraform apply -var 'ipv6_mgmt_network=<Your comma separated IPv6 CIDR blocks here>' -var 'ipv4_mgmt_network=<Your comma separated IPv4 CIDR blocks here>'` to deploy the changes
1. Run `terraform destroy -var 'ipv6_mgmt_network=<Your comma separated IPv6 CIDR blocks here>'  -var 'ipv4_mgmt_network=<Your comma separated IPv4 CIDR blocks here>'` to remove all the changes from AWS

## Notes
We will do the later ones in a GItHub action getting tired of typing terraform commands
