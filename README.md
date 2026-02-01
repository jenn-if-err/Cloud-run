# Cloud-run

## Repo Structure

- `go-app/` — Cloud Run service (containerized Go app)
- `go-function/` — Google Cloud Function (Go HTTP function)
- `key.json` — Service account key (for local gcloud auth)
- `Makefile` (in go-app) — Automates build and deploy for Cloud Run

---

## Prerequisites
- Google Cloud SDK (`gcloud`)
- Project owner or editor permissions
- Service account with required IAM roles
- Docker (for Cloud Run)

---

## Cloud Run (go-app)

### Build & Deploy
1. **Auth locally (if needed):**
   ```
   gcloud auth activate-service-account <SERVICE_ACCOUNT_EMAIL> --key-file=key.json
   gcloud config set project <PROJECT_ID>
   ```
2. **Build container:**
   ```
   cd go-app
   docker build -t gcr.io/<PROJECT_ID>/go-app .
   docker push gcr.io/<PROJECT_ID>/go-app
   ```
3. **Deploy to Cloud Run:**
   ```
   gcloud run deploy go-app \
     --image=gcr.io/<PROJECT_ID>/go-app \
     --region=<REGION> \
     --platform=managed \
     --allow-unauthenticated
   ```

### Automate with Makefile
- Use the provided `Makefile` in `go-app` to automate build and deployment:
  ```
  cd go-app
  make deploy
  ```
  (See Makefile for available targets)

---

## Cloud Functions (go-function)

### Deploy (Gen2, Go)
1. **Auth locally (if needed):**
   ```
   gcloud auth activate-service-account <SERVICE_ACCOUNT_EMAIL> --key-file=key.json
   gcloud config set project <PROJECT_ID>
   ```
2. **Deploy function:**
   ```
   cd go-function
   gcloud functions deploy HelloWorld \
     --gen2 \
     --runtime=go121 \
     --region=<REGION> \
     --source=. \
     --entry-point=HelloWorld \
     --trigger-http \
     --allow-unauthenticated
   ```
   - If you need a custom service account:
     ```
     --service-account=<SERVICE_ACCOUNT_EMAIL>
     ```

### Redeploy Cloud Function (after code change)
1. **Auth locally (if needed):**
   ```
   gcloud auth activate-service-account <SERVICE_ACCOUNT_EMAIL> --key-file=key.json
   gcloud config set project <PROJECT_ID>
   ```
2. **Redeploy function:**
   ```
   cd go-function
   gcloud functions deploy HelloWorld \
     --gen2 \
     --runtime=go121 \
     --region=<REGION> \
     --source=. \
     --entry-point=HelloWorld \
     --trigger-http \
     --allow-unauthenticated
   ```
   - If you need a custom service account:
     ```
     --service-account=<SERVICE_ACCOUNT_EMAIL>
     ```
3. **Test your function:**
   ```
   curl "<FUNCTION_URL>?name=jen"
   ```
   - Replace `<FUNCTION_URL>` with the URL from the deployment output or:
     ```
     gcloud functions describe HelloWorld --gen2 --region=<REGION> --format="value(serviceConfig.uri)"
     ```
   - Expected output: `Hello jen!`

---

## IAM Roles Needed
- `roles/cloudbuild.builds.builder`
- `roles/cloudfunctions.developer`
- `roles/storage.admin` or `roles/storage.objectAdmin`
- `roles/artifactregistry.writer`
- `roles/logging.logWriter`

Grant to:
- `[PROJECT_NUMBER]@cloudbuild.gserviceaccount.com`
- `[PROJECT_NUMBER]-compute@developer.gserviceaccount.com`

---

## Notes
- Cloud Run: Containerized, any language/runtime, HTTP endpoint.
- Cloud Functions: Source-based, event-driven, HTTP trigger, no main() needed.
- Use Gen2 for new functions.
- Update Go runtime as needed (Go 1.21 deprecation warning).

---

## Troubleshooting
- Check build logs in Cloud Console if deploy fails.
- Ensure IAM roles are granted to build service accounts.
- For permission errors, see: https://cloud.google.com/functions/docs/troubleshooting#build-service-account

---

## TL;DR
- Build, push, deploy for Cloud Run (or just `make deploy`).
- Deploy from source for Cloud Functions.
- Make sure IAM is set up. If it fails, check logs and fix permissions.