# Create a CentOS Stream 8 image
## Overview
Utilizes Packer and Ansible to create a CentOS Stream 8 image. 

## Usage
1. The entire cloned respository must be available as the `centos.pkr.hcl` references the ansible folder for provisioning customizations.
1. As the CentOS Stream ISO names change (rolling release) you will need to identify the currently available release at the mirror of your preference and update the `iso_url` variable
1. Update the `rootpw` in the `ks_centos8.cfg` file to your preference as well as the `user` definition and password
1. Build the template with with `packer build -force -var 'ssh_password=YourUserPassword' -var 'ssh_username=YouUserName' .`
  * ssh_username is the `user` you created in the `ks_centos8.cfg`
  * ssh_password is the password for the `user` you created in `ks_centos8.cfg`

## Notes
Encrypting the password for the Kickstart file `ks_centos8.cfg` can be done using the example Python script `encryptme.py`