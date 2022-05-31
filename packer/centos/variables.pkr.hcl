variable "vm_name" {
  description = "The name of the virtual machine to be created."
  type        = string
}

variable "iso_url" {
  type        = string
  description = "The URL for the ISO for CentOS 8 Stream."
}

variable "iso_checksum" {
  type        = string
  description = "The file URL for the ISO CHECKSUM file for CentOS 8 Stream."
}

variable "ssh_username" {
  type        = string
  default     = "agayle"
  description = "The username to use for SSH connections."
}

variable "ssh_password" {
  type        = string
  sensitive   = true
  description = "The password of the username to use for the SSH connections."
}

variable "cpus" {
  type        = number
  default     = 1
  description = "The number of virtual CPU to assign."
}

variable "memory" {
  type        = number
  default     = 1024
  description = "The RAM to assign to the virtual machine in MB."
}

variable "disk_size" {
  type        = number
  default     = 20000
  description = "The size of the operating system install target disk in MB."
}

variable "http_directory" {
  type        = string
  default     = "./http"
  description = "The directory to be used as the http source for packer to serve Kickstart files."
}

variable "ks_path" {
  type        = string
  default     = "ks_centos8.cfg"
  description = "The Kickstart file to be used to provision the virtual machine."
}
