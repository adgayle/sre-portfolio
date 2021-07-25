output "db_password_id" {
  description = "Generated db password id"
  value       = random_password.db_password.id
}

output "db_password" {
  description = "Generated sensitive / no show db password"
  value       = random_password.db_password.result
  sensitive   = true
}
