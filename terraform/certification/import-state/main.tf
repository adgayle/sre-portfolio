resource "random_password" "db_password" {
  length      = var.password_length
  min_special = var.min_special
  min_numeric = var.min_numeric
  min_upper   = var.min_upper
}
