source "virtualbox-iso" "centos8-vbox" {
  vm_name              = var.vm_name
  guest_os_type        = "RedHat_64"
  guest_additions_mode = "attach"

  iso_url        = var.iso_url
  iso_checksum   = var.iso_checksum
  http_directory = var.http_directory

  communicator = "ssh"
  ssh_timeout  = "30m"
  ssh_username = var.ssh_username
  ssh_password = var.ssh_password

  cpus                 = var.cpus
  memory               = var.memory
  disk_size            = var.disk_size
  hard_drive_interface = "sata"

  boot_wait = "5s"
  boot_command = [
    "<tab> text ks=http://{{ .HTTPIP }}:{{ .HTTPPort }}/${var.ks_path}<enter><wait>"
  ]

  shutdown_command = "echo ${var.ssh_password} | sudo -S shutdown -P now"
  shutdown_timeout = "5m"
  skip_export      = true
  keep_registered  = true
}

build {
  sources = ["sources.virtualbox-iso.centos8-vbox"]

  provisioner "shell" {
    inline = [
      "echo ${var.ssh_password} | sudo -S dnf install -y epel-release",
      "echo ${var.ssh_password} | sudo -S dnf install -y ansible"
    ]
  }

  provisioner "ansible-local" {
    playbook_dir            = "../../ansible"
    playbook_file           = "../../ansible/workstation-playbook.yml"
    extra_arguments         = ["--extra-vars", "\"ansible_become_password=${var.ssh_password}\""]
    clean_staging_directory = true
  }
}
