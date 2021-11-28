output "dog_name" {
  description = "The generated dog name from a prefix."
  value       = title(module.dog.pet_name)
}

output "famous_names" {
  description = "The names of famous dogs."
  value       = [for key, value in module.famous_dogs_enhanced : title(value.pet_name)]
}

output "large_dog_names" {
  description = "The names of some large dogs."
  value       = [for key, value in module.large_dogs : title(value.pet_name)]
}
