query=

publisher_name+eq+PrimaryPublisherName+and+registered+eq+false



data "aws_ami" "web" {
  filter {
    name   = "state"
    values = ["available"]
  }

  filter {
    name   = "tag:Component"
    values = ["web"]
  }

  most_recent = true
}


data "netskope_publishers" "all" {
    publisher_name = "BitBucket-demo-publisher"

}