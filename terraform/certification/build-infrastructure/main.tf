resource "aws_security_group" "allow_ssh" {
  name        = "allow_ssh"
  description = "Allow SSH inbound access"
  vpc_id      = data.aws_vpc.default.id

  ingress {
    description      = "SSH from trusted"
    from_port        = 22
    to_port          = 22
    protocol         = "tcp"
    ipv6_cidr_blocks = ["::/0"]
  }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = -1
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  tags = {
    Name        = "Allow SSH"
    Environment = "Development"
    Lesson      = "Terraform Build Infrastructure"
  }
}

resource "aws_instance" "web" {
  ami           = data.aws_ami.amazon_linux.id
  instance_type = "t3.micro"
  subnet_id     = sort(data.aws_subnet_ids.default.ids)[0]
  key_name      = "terraform_automation"
  vpc_security_group_ids = [aws_security_group.allow_ssh.id]

  tags = {
    Name        = "web-01"
    Lesson      = "Terraform Build Infrastructure"
    Environment = "Development"
  }
}