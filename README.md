# Jobs Backend

A Go backend service for the AI Interview Prep application. Provides authentication via Firebase and AI-powered resume analysis via OpenAI.

## Tech Stack

- **Language:** Go 1.25.3
- **Web Framework:** Chi Router
- **Authentication:** Firebase Admin SDK
- **AI Integration:** OpenAI API
- **Environment:** godotenv for configuration

## Prerequisites

- **Go 1.25.3 or later** - [Download Go](https://golang.org/dl/)
- **Firebase Service Account Key** - `fb-sa-jobs-pk.json` file
- **OpenAI API Key** - For AI-powered features

## Environment Variables

Create a `.env` file in the `jobs-backend` directory with the following variables:

```bash
# Server configuration
PORT=8080                    # Server port (or use HOST=:8080)
# OR
HOST=:8080

# Firebase configuration
FIREBASE=./fb-sa-jobs-pk.json    # Path to Firebase service account JSON file

# OpenAI API
OPENAI_API_KEY=sk-your-openai-api-key

# CORS configuration (optional - defaults to http://localhost:5173 for local dev)
CORS_ORIGINS=http://localhost:5173,https://your-frontend-url.com

# Debug mode (optional - set to "true" to enable CORS debugging)
DEBUG=false
```

## Getting Started

### 1. Clone and Setup

```bash
cd jobs-backend
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Configure Environment

1. Create a `.env` file:
   ```bash
   cp .env.example .env  # If you have an example file
   # Or create manually
   ```

2. Add your Firebase service account key:
   - Download `fb-sa-jobs-pk.json` from Firebase Console
   - Place it in the `jobs-backend` directory
   - Update `FIREBASE` in `.env` to point to this file

3. Set your OpenAI API key in `.env`

### 4. Run Locally

```bash
go run cmd/server/main.go
```

The server will start on `http://localhost:8080` (or the port specified in your `.env` file).

### 5. Verify It's Running

Test the health endpoint:
```bash
curl http://localhost:8080/ping
```

Expected response:
```json
{"message":"pong"}
```

## Development

### Running in Development Mode

```bash
# Run with auto-reload (requires a tool like air or nodemon)
go run cmd/server/main.go

# Or use air for hot-reloading
air
```

### Project Structure

```
jobs-backend/
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
├── internal/
│   ├── clients/
│   │   └── llm.go           # OpenAI client
│   ├── server/
│   │   ├── authenticate.go  # Firebase authentication
│   │   ├── handlers.go      # HTTP handlers
│   │   ├── middleware.go    # CORS and logging middleware
│   │   ├── routes.go        # Route definitions
│   │   └── server.go        # Server initialization
│   └── util/
│       └── utils.go         # Utility functions
├── .ebextensions/           # Elastic Beanstalk configuration
│   └── go.config            # EB environment settings
├── .elasticbeanstalk/       # EB CLI configuration (not in git)
├── bin/                     # Build output (not in git)
│   └── application          # Compiled binary for EB
├── build.sh                 # Build script for Elastic Beanstalk
├── Procfile                 # Tells EB how to run the app
├── Dockerfile               # Docker configuration
├── go.mod                   # Go module dependencies
├── go.sum                   # Dependency checksums
├── fb-sa-jobs-pk.json       # Firebase credentials (not in git)
└── .env                     # Environment variables (not in git)
```

### API Endpoints

- `GET /ping` - Health check (public)
- `POST /upload` - Upload resume for analysis (protected - requires Firebase auth token)

## Building

### Build Binary

Build for your current platform:
```bash
go build -o server ./cmd/server
```

Build for Linux (useful for Docker):
```bash
CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server
```

Build for multiple platforms:
```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o server-linux ./cmd/server

# macOS
GOOS=darwin GOARCH=amd64 go build -o server-macos ./cmd/server

# Windows
GOOS=windows GOARCH=amd64 go build -o server-windows.exe ./cmd/server
```

### Run Built Binary

```bash
./server
```

## Docker

### Build Docker Image

```bash
docker build -t jobs-backend .
```

### Run Docker Container

```bash
docker run -d \
  -p 8080:8080 \
  -e PORT=8080 \
  -e OPENAI_API_KEY=your_key \
  -e FIREBASE=/root/fb-sa-jobs-pk.json \
  -e CORS_ORIGINS=http://localhost:5173 \
  -v $(pwd)/fb-sa-jobs-pk.json:/root/fb-sa-jobs-pk.json:ro \
  --name jobs-backend \
  jobs-backend
```

### Using Docker Compose

See the root `docker-compose.yml` for full stack deployment.

## Testing

### Manual Testing

Test the health endpoint:
```bash
curl http://localhost:8080/ping
```

Test protected endpoint (requires Firebase token):
```bash
curl -X POST http://localhost:8080/upload \
  -H "Authorization: Bearer YOUR_FIREBASE_TOKEN" \
  -F "resume=@path/to/resume.pdf" \
  -F "url=https://example.com/job-posting"
```

## Deployment

### AWS Elastic Beanstalk (Current Production Setup)

This backend is currently deployed to AWS Elastic Beanstalk. Here's how to deploy:

#### Prerequisites

1. **Install AWS EB CLI:**
   ```bash
   pip3 install awsebcli
   # Or via Homebrew (macOS):
   brew install awsebcli
   ```

2. **Configure AWS credentials:**
   ```bash
   aws configure
   ```

#### Initial Setup

1. **Initialize Elastic Beanstalk:**
   ```bash
   cd jobs-backend
   eb init -p go-1.21 jobs-app-backend --region us-east-2
   ```

2. **Create environment:**
   ```bash
   eb create go-env-jobs
   ```

#### Building and Deploying

1. **Build the binary:**
   ```bash
   ./build.sh
   ```
   This creates `bin/application` (Linux binary for EB).

2. **Create deployment package:**
   ```bash
   zip -r deploy.zip . \
     -x "*.git*" \
     -x "*.env*" \
     -x "*.md" \
     -x "!README.md" \
     -x "build.sh" \
     -x ".elasticbeanstalk/*" \
     -x "!.elasticbeanstalk/*.cfg.yml"
   ```
   
   **Important:** Make sure `bin/application` and `Procfile` are included in the zip.

3. **Deploy:**
   ```bash
   eb deploy
   ```
   
   Or deploy from zip:
   ```bash
   # Upload via AWS Console, or commit to git and use:
   eb deploy
   ```

#### Environment Variables

Set environment variables via EB CLI or AWS Console:

```bash
eb setenv \
  OPENAI_API_KEY=your_openai_key \
  FIREBASE=/var/app/current/fb-sa-jobs-pk.json \
  CORS_ORIGINS=https://your-frontend-url.com
```

**Via AWS Console:**
1. Go to Elastic Beanstalk → Your environment
2. Configuration → Software → Environment properties
3. Add/Edit environment variables

#### Firebase Credentials

The Firebase service account file must be included in your deployment:

1. **Option 1: Include in deployment zip**
   - Add `fb-sa-jobs-pk.json` to your deployment zip
   - Set `FIREBASE=/var/app/current/fb-sa-jobs-pk.json`

2. **Option 2: Upload to S3 and use `.ebextensions`**
   - Upload `fb-sa-jobs-pk.json` to S3
   - Create `.ebextensions/firebase.config`:
     ```yaml
     files:
       "/var/app/current/fb-sa-jobs-pk.json":
         mode: "000644"
         owner: root
         group: root
         source: https://your-bucket.s3.us-east-2.amazonaws.com/fb-sa-jobs-pk.json
     ```

#### Required Files

Make sure these files are in your deployment:
- `bin/application` - The compiled binary (created by `build.sh`)
- `Procfile` - Tells EB how to run the app: `web: bin/application`
- `fb-sa-jobs-pk.json` - Firebase credentials (if including directly)

#### CloudFront Setup (HTTPS)

The backend is behind CloudFront for HTTPS:

1. **Create CloudFront Distribution:**
   - Origin: Your Elastic Beanstalk URL
   - Origin protocol: **HTTP Only** (EB uses HTTP, not HTTPS)
   - Viewer protocol: Redirect HTTP to HTTPS

2. **Configure Behaviors:**
   - Cache policy: `CachingDisabled` (for API endpoints)
   - Origin request policy: `AllViewer` (forwards all headers)
   - Response headers policy: None (let backend handle CORS)

3. **Update frontend:**
   - Set `VITE_API_URL` to your CloudFront URL

#### Troubleshooting

- **"no application binary found"**: Make sure `bin/application` exists and is in the deployment zip
- **"no Procfile found"**: Ensure `Procfile` is in the root directory
- **Firebase credential errors**: Verify `FIREBASE` env var path and file exists
- **CORS errors**: Check `CORS_ORIGINS` includes your frontend URL

See [ELASTIC_BEANSTALK_SETUP.md](./ELASTIC_BEANSTALK_SETUP.md) for detailed troubleshooting.

#### Other Deployment Options

See the main [DEPLOYMENT.md](../DEPLOYMENT.md) for other platforms:
- Railway
- Render
- Fly.io
- AWS ECS/EC2
- Google Cloud Platform
- Microsoft Azure

## Troubleshooting

### Server won't start

1. **Check .env file exists:** The server requires a `.env` file in the `jobs-backend` directory
2. **Check Firebase credentials:** Ensure `fb-sa-jobs-pk.json` exists and path in `.env` is correct
3. **Check port availability:** Make sure port 8080 (or your configured port) is not in use

### CORS errors in frontend

1. **Check CORS_ORIGINS:** Ensure your frontend URL is in the `CORS_ORIGINS` environment variable
2. **Enable DEBUG:** Set `DEBUG=true` in `.env` to see CORS logs

### Firebase authentication fails

1. **Verify credentials file:** Check that `fb-sa-jobs-pk.json` is valid
2. **Check file path:** Ensure `FIREBASE` env var points to the correct file path
3. **Verify Firebase project:** Make sure the service account has the correct permissions

## Documentation

- **[Journey.md](./Journey.md)** - Development journey and notes
- **[TDD.md](./TDD.md)** - Technical Design Document
- **[ProjectPlan.md](./ProjectPlan.md)** - Project planning document

## License

[Add your license here]
