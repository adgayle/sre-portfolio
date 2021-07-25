# Create an EC2 instance in AWS
## Overview
Utilizes terraform to create and Amazon Linux 2 instance in AWS with supporting security group in the default VPC (sufficient for the purpose of the exercise but not recommended otherwise)

## Usage
Variables and secrets are stored in Terraform Cloud
1. Run `terraform init` to initialize the environment
1. Run `terraform plan` to see the build plan of intent
1. Run `terraform apply'` to deploy the changes
1. Run `terraform destroy` to remove all the changes from AWS

## Notes
We will do the later ones in a GItHub action getting tired of typing terraform commands
