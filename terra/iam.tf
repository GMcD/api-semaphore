# Service Account
resource "google_service_account" "gcr_pull" {
  account_id = "service-account-gcr-pull"
  display_name = "Service Account for Pulling images from GCR"
}

# Storage Viewer for Images
resource "google_project_iam_member" "storage_viewer" {
  count = 1
  project = var.project_id
  role = "roles/storage.objectViewer"
  member = "serviceAccount:${google_service_account.gcr_pull.email}"
}

# Artifact Reader for Manifests
resource "google_project_iam_member" "artifact_reader" {
  count = 1
  project = var.project_id
  role = "roles/artifactregistry.reader"
  member = "serviceAccount:${google_service_account.gcr_pull.email}"
}

# Generate Service Key
resource "google_service_account_key" "gcr_pull_key" {
  service_account_id = google_service_account.gcr_pull.name
  public_key_type    = "TYPE_X509_PEM_FILE"
}


# Output Project Email
output "service_account_email" {
  description = "Service account email (for single use)."
  value       = google_service_account.gcr_pull.email
}

# Output Service Account Private Key
output "service_account_private_key" {
    description = "Service Account Private Key"
    value       = google_service_account_key.gcr_pull_key.private_key
    sensitive   = true
}
