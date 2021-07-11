terraform {
  required_providers {
    random = {
      source  = "hashicorp/random"
      version = ">= 3.0.1"
    }

    aws = {
      source  = "hashicorp/aws"
      version = ">= 3.0.0"
    }
  }

  required_version = "~> 1.0"
}
