terraform {
  required_providers {
    random = {
      source  = "hashicorp/random"
      version = "3.1.0"
    }
  }

  backend "remote" {
    organization = "wec-acme-inc"

    workspaces {
      name = "certification-migrate-state"
    }
  }

  required_version = "1.0.1"
}
