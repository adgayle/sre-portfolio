terraform {
  required_version = "~> 1.0.0"

  backend "remote" {
    organization = "wec-acme-inc"

    workspaces {
      name = "pets-saga"
    }
  }
}

module "dog" {
  source = "./modules/pet-generator"

  pet_prefix = var.single_pet_prefix
}

module "famous_dogs_enhanced" {
  source = "./modules/pet-generator"

  for_each   = var.famous_dogs_prefix
  pet_prefix = each.key
}

module "large_dogs" {
  source = "./modules/pet-generator"

  for_each   = { for key, value in var.dogs : key => value if value == "large" }
  pet_prefix = each.key
}
