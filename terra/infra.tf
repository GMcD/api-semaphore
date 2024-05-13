
resource "google_sql_database_instance" "instance" {
  name                = var.app_db_instance
  region              = var.region
  database_version    = "POSTGRES_16"
  deletion_protection = false
  settings {
    tier = "db-f1-micro"
  }
}
resource "google_sql_database" "database" {
  name     = var.app_db_name
  instance = google_sql_database_instance.instance.name
}
resource "google_sql_user" "database-user" {
  name     = var.app_db_username
  instance = google_sql_database_instance.instance.name
  password = var.app_db_password
}
resource "google_cloud_run_service" "run" {
  name     = var.app_service
  location = var.region
  template {
    spec {
      containers {
        # image = "gcr.io/${var.project_id}/api-semaphore/${var.app_service}:1"
        image = "gcr.io/${var.project_id}/${var.app_service}:latest"
        resources {
          limits = {
            memory = "2G"
          }
        }
        ports {
          container_port = 8100
        }
        # env {
        #   name  = "JWT_KEY"
        #   value = var.jwt_key
        # }
        env {
          name  = "DB_URL"
          value = "postgresql://${var.app_db_username}:${var.app_db_password}@/${var.app_db_name}?host=/cloudsql/${google_sql_database_instance.instance.connection_name}"
        }
      }
    }
    metadata {
      annotations = {
        "run.googleapis.com/cloudsql-instances" = google_sql_database_instance.instance.connection_name
      }
    }
  }
}

resource "google_cloud_run_service_iam_member" "member" {
  service  = google_cloud_run_service.run.name
  location = google_cloud_run_service.run.location
  role     = "roles/run.invoker"
  member   = "allUsers"
}
