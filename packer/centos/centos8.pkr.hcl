source "virtualbox-iso" "centos8-vbox" {
  vm_name              = var.vm_name
  guest_os_type        = "RedHat_64"
  guest_additions_mode = "attach"

  iso_url              = var.iso_url
  iso_checksum         = var.iso_checksum
  http_directory       = var.http_directory

  ssh_username         = var.ssh_username
  ssh_password         = var.ssh_password

  cpus                 = var.cpus
  memory               = var.memory
  disk_size            = var.disk_size
  hard_drive_interface = "scsi"

  boot_wait            = "5s"
  boot_command         = [
    "<tab> text ks=http://{{ .HTTPIP }}:{{ .HTTPPort }}/{{ user `ks_path` }}<enter><wait>"
  ]

  shutdown_command     = "echo ${var.ssh_password} | sudo -S shutdown -P now"
  shutdown_timeout     = "5m"
  skip_export          = true
  keep_registered      = true
}

build {
  sources = ["sources.virtualbox-iso.centos8-vbox"]
}