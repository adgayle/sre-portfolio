variable "vm_name" {
  type = string
}

variable "iso_url" {
  type    = string
  default = "http://ftp.usf.edu/pub/centos/8-stream/isos/x86_64/CentOS-Stream-8-x86_64-20210126-boot.iso"
}

variable "iso_checksum" {
  type    = string
  default = "file:http://ftp.usf.edu/pub/centos/8-stream/isos/x86_64/CHECKSUM"
}

variable "ssh_username" {
  type    = string
  default = "agayle"
}

variable "ssh_password" {
  type      = string
  sensitive = true
}

variable "cpus" {
  type    = number
  default = 1
}

variable "memory" {
  type = number
  // in MB
  default = 1024
}

variable "disk_size" {
  type = number
  // in MB
  default = 20000
}

variable "http_directory" {
  type    = string
  default = "./http"
}

variable "ks_path" {
  type    = string
  default = "ks_centos8.cfg"
}
