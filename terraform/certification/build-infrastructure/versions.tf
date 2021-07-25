terraform {
  backend "remote" {
    organization = "wec-acme-inc"

    workspaces {
      name = "certification-remote-state"
    }
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 3.0"
    }
  }

  required_version = "~> 1.0"
}
