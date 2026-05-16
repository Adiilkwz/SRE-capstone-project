variable "aws_region" {
  description = "The AWS region to deploy into"
  default     = "eu-central-1"
}

variable "cluster_name" {
  description = "Name of the Kubernetes cluster"
  default     = "prr-cluster"
}