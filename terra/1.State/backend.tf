# Backend State Bucket and Project Id
terraform {
  backend "gcs" {
    bucket = "gmcd-414115-tf-state"
    prefix = "gmcd-414115-"
  }
}
