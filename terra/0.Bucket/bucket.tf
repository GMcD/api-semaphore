
# Create a GCS Bucket for TF state
resource "google_storage_bucket" "tf_state_bucket" {
  name     = var.tf_bucket_name
  location = var.region
  project  = var.project_id
}
