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
├── Dockerfile               # Docker configuration
├── go.mod                   # Go module dependencies
├── go.sum                   # Dependency checksums
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

See the main [DEPLOYMENT.md](../DEPLOYMENT.md) for comprehensive deployment instructions including:
- Railway
- Render
- Fly.io
- AWS (Elastic Beanstalk, ECS, EC2)
- Google Cloud Platform (Cloud Run, Compute Engine, App Engine)
- Microsoft Azure (App Service, Container Instances, VMs)

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
