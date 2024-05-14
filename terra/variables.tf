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

variable "system-email" {
  description = "System Email"
  default     = "system@projectscapa.com"
}

//
// DNS
//
variable "root-zone" {
  description = "Domain Name for Godaddy"
  default     = ""
}

variable "project-zone" {
  description = "Environnment for Project"
  default     = ""
}

variable "project_name" {
  description = "Project Name"
  default     = ""
}

variable "env-name" {
  description = "Environment Name for Resource Tagging."
  default     = "dev"
}

variable "gd_customer" {
  description = "GoDaddy Customer Account"
  default     = "35481312"
}
variable "gd_name" {
  description = "GoDaddy API key Name"
  default     = ""
}
variable "gd_apikey" {
  description = "GoDaddy API key"
  default     = ""
}
variable "gd_secret" {
  description = "GoDaddy API secret"
  default     = ""
}

variable "repository" {
  description = "Artifactory Repository"
  default     = ""
}
