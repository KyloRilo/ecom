terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
    }
  }
  required_version = ">= 0.14"

  backend "remote" {
    organization = "gmriley-projects"

    workspaces {
      name = "ecom-gh-actions"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

data "aws_availability_zones" "azs" {
    state = "available"
}