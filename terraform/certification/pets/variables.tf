variable "single_pet_prefix" {
  description = "Prefix for the name of a single pet."
  type        = string
  default     = "Rover"
}

variable "famous_dogs_prefix" {
  description = "Prefix using famous dog names."
  type        = set(string)
  default = [
    "Clifford",
    "Lassie",
    "Bolt",
    "Beethoven"
  ]
}

variable "dogs" {
  description = "Map of dogs and their sizes."
  type        = map(string)
  default = {
    "German Shepard" = "large"
    "Saint Bernard"  = "large"
    "Fox Terrier"    = "small"
    "Chihuahua"      = "small"
  }
}
