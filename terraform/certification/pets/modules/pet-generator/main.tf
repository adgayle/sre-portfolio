terraform {
  required_version = ">= 1.0"

  required_providers {
    random = {
      source  = "hashicorp/random"
      version = ">= 3.1"
    }
  }
}

resource "random_pet" "pet_name" {
  prefix    = var.pet_prefix
  length    = 1
  separator = " "
}
