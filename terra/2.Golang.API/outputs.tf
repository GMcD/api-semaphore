
# Output Cloud Run Url
output "cloud-run-url" {
  value = google_cloud_run_service.run.status[0].url
}
