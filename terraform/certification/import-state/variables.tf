variable "password_length" {
  type        = number
  description = "Length of password to generate"
  default     = 24
}

variable "min_special" {
  type        = number
  description = "Minimum number of special characters"
  default     = 3
}

variable "min_upper" {
  type        = number
  description = "Minimum number of uppercase characters"
  default     = 3
}

variable "min_numeric" {
  type        = number
  description = "Minimun number of numeric values"
  default     = 2
}
