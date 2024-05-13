# GCS bucket name to store terraform tfstate
variable "tf_bucket_name" {
  description = "Terraform State Bucket"
  type        = string
  default     = "gmcd-414115-tf-state"
}

# Region for Resources
variable "region" {
  description = "Google Cloud region"
  type        = string
  default     = "europe-west2"
}

# Project Id
variable "project_id" {
  description = "Google Project ID"
  type        = string
}

# Service Details
variable "app_service" {
  description = "Cloud Run Service"
  type        = string
}

variable "app_db_instance" {
  description = "Cloud SQL Service Instance"
  type        = string
}

variable "app_db_name" {
  description = "SQl Database Name"
  type        = string
}

variable "app_db_username" {
  description = "SQL Database User Name"
  type        = string
}

variable "app_db_password" {
  description = "SQL User Password - XXX Move to secret store"
  type        = string
}

# variable "jwt_key" {
#   type = string
# }
