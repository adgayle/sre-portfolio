#!/usr/bin/env python3
'''
Encrypts the entered password for use in the Kickstart configuration

Doc: https://docs.centos.org/en-US/8-docs/advanced-install/assembly_kickstart-commands-and-options-reference/#rootpw-required_kickstart-commands-for-system-configuration
'''

import crypt
import getpass 

password=getpass.getpass()
print(crypt.crypt(password) if (password==getpass.getpass("Confirm: ")) else exit())
