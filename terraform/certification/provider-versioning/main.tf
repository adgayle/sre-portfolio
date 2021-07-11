resource "random_string" "bucket_name" {
  length  = 18
  special = false
  lower   = true
  upper   = false
  number  = false
}

resource "aws_s3_bucket" "random_bucket" {
  bucket = random_string.bucket_name.id
  acl    = "private"

  tags = {
    Name        = "Random Bucket"
    Environment = "Development"
    Lesson      = "Terraform"
  }
}
