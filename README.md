## My full stack journey

- focused more heavy on the backend and system design
- lightweight frontend using NextJS and Tailwind
- Microservice baackend
- Using technologies like gRPC
-

## Tech stack

- NextJS and Tailwind for frontend
- Golang for microservices backend
- Postgres for DB
- will try to use redis for caching
- will try to use kafka for messaging
- will use grpc for internal service communication
- will use http/rest for frontend vs backend communication
- will use cloud to hotes gcp or aws
- will use docker to containerize
- will learn k8s for scalability
- will learn datadog or others for logging

## Notes

- Not use AI to write code

## Pre

- I did a lot of research initially with the help of

## Day 1 - 10/22/2025

- I chatted with AI around what I should build this has been my biggest challenge so far even before day 1 as I have been wanting to build something useful but came to a conclusion that I want to start with a custom ecommerce website that's somewhat complex
- I set up my repo but I should've probably done some system architecture design and requirements gathering if I were in a professional setting but since I'm more focused on learnning new technologies, I say let's just jump in.

## Day 2 - 10/23/2025

- Set up gateway service
- Set up basic handler for Hello and GetProducts that will call Products service
- Set up clients where handlers will utilize clients to call Products service
- Use grpc call to products service

## Day 3 - 10/24/2025

- Set up products repo
- created basic template
- need to set up proto and spin up grpc server
- signed up buf registry to maintain proto files

## Day 4 - 10/25/2025

- Containerized gateway using docker file
- Should multi-stage this to make it lean
- https://buf.build/docs/bsr/quickstart/#call-apis-with-client-sdks
  - Makes it easy to import client SDKs to use common protos
  - Make sure to go mod tidy when links don't resolve
- Use common folder to define services
- I was able to set up grpc for products
- using grpcui --plaintext localhost:8081, i was able to make a call
- buf.yaml is the configuration file for your proto module
- buf.gen.yaml is what generates the code locally (creates the gen file)
- Workflow: `buf lint` -> `buf generate` -> `buf push`
  - pushes the proto and i will be able to get the following:
  - "buf.build/gen/go/wassup-chicken/common/protocolbuffers/go/api/v1"
  - "buf.build/gen/go/wassup-chicken/common/grpc/go/api/v1/apiv1grpc"
- From proto to generated Go packages:
  - Message types (data): `buf.build/gen/go/wassup-chicken/common/protocolbuffers/go/api/v1`
    - e.g., `GetProductRequest`, `GetProductResponse`
  - gRPC server interfaces (methods): `buf.build/gen/go/wassup-chicken/common/grpc/go/api/v1/apiv1grpc`
    - e.g., `ProductServiceServer`, `RegisterProductServiceServer`
- When the proto changes, run:
  - `buf lint`
  - `buf generate`
  - `buf push`
  - update/get deps if needed:
    - gRPC Go SDK
    - Protocol Buffers Go
- Local generation with `buf.gen.yaml` (in this repo):

  - Command: `buf generate` (reads `buf.gen.yaml` and writes to `./gen`)
  - What gets generated with current config:
    - `protocolbuffers/go` → `*.pb.go`
      - contains: message types, enums, and service descriptors
    - `grpc/go` → `*_grpc.pb.go`
      - contains: gRPC server/client interfaces and registration functions
  - Optional: to also generate Connect handlers/clients, add plugin:
    - `- remote: buf.build/connectrpc/go` (outputs `*_connect.go`)

- File type differences (quick reference):
  - `*.pb.go` (protocolbuffers): messages, enums, service descriptors; used by both gRPC and Connect
  - `*_grpc.pb.go` (grpc): gRPC server/client interfaces and registration functions
  - `*_connect.go` (connectrpc): Connect handlers/clients for HTTP/2 and HTTP/1.1

## Day 5 - 10/27/2025

- Set up grpc handlers for GetProduct and GetProducts
- Did some research on setting up postgres on the products service
- still need to learn how to query and understand the database url

## Day 6 - 10/30/2025

- Created a new package for store to work on postgres
- Set up basic templates for it
  - Initializing a connection pool, making it an interface to GetProd and GetProds
- For next time:
  - think about environment variables
  - context , how it should be used
  - refine gateway services
- Rules of thumb
  - Use the handler’s ctx for all work tied to the request.
  - Add timeouts at the edge (DB calls, outbound RPC).
  - Never keep a context in a struct field; pass it as a parameter.
  - Root context is for service lifecycle; request context is for per-request work.

## Day 7 - 11/03/2025

- Today, I learned about postgres set up and how i can connect to the PostgresSQL server from golang package.
- PostgreSQL setup:
  - Server: Postgres.app (PostgreSQL 18.0) running on `localhost:5432`
  - Database: `products` database exists
  - Connection string: `postgresql://postgres@localhost:5432/products`
    - Protocol: `postgresql://`
    - Username: `postgres`
    - Host: `localhost`
    - Port: `5432` (default PostgreSQL port)
    - Database: `products`
  - This connects to the Postgres.app server's `products` database
- psql terminal client:
  - `psql` is the PostgreSQL command-line client
  - Available via Homebrew: `/opt/homebrew/opt/postgresql@18/bin/psql`
  - Available via Postgres.app: `/Applications/Postgres.app/Contents/Versions/latest/bin/psql`
  - Both clients work with the Postgres.app server - the version difference (client v14/v18 vs server v18) is fine for development
  - Common psql commands:
    - `\l` - list all databases
    - `\dt` - list tables in current database
    - `\d table_name` - describe table structure
    - `\c database_name` - switch database
    - `\q` - quit/exit
    - Connect: `psql -h localhost -U postgres -d products`
- Client vs Server:
  - **Homebrew `postgresql@18`**: Installs both server AND client tools
  - **Postgres.app**: Server runs on port 5432, client tools bundled in app
  - **libpq** (Homebrew): Client-only package, no server
  - Recommendation: Use Postgres.app server, use either client (Homebrew or Postgres.app's)
- Learn how data is persisted
  - localhost - on mac machine
  - cloud - ec2 public dns would be the host
- Want to explore how to set up a script to write a database table on start up
  - how to create a table if doesn't exists
  - how to persist data

## Day 8 12/5/2025

- Architecture change: Switching from microservices to monolithic architecture
  - As a solo engineer, microservices add unnecessary complexity and operational overhead
  - Goal is to ship a working product and deploy it to the cloud
  - Focus on learning web application fundamentals rather than distributed systems complexity
  - Can always refactor into microservices later if needed (strangler pattern)
  - Future consideration: Break down into services only when there's a clear need (team growth, independent scaling requirements, or well-defined domain boundaries)
- Tech stack:
  - **Frontend**: NextJS + Tailwind CSS
  - **Backend**: Golang (monolithic)
  - **Database**: PostgreSQL
  - **Caching**: Redis
  - **Containerization**: Docker
  - **Orchestration**: Kubernetes (if needed for scaling)
  - **Cloud**: GCP or AWS (for hosting)
  - **API**: HTTP/REST (for frontend-backend communication)
  - **Monitoring/Logging**: Datadog or similar (mentioned earlier)
  - **CI/CD**: GitHub Actions or similar (for automated deployments)
  - **Additional considerations** (to add as needed):
    - Authentication/Authorization (JWT, OAuth, or session-based)
    - Payment processing (Stripe, PayPal, etc.)
    - File storage (S3, Cloud Storage for product images)
    - Email service (SendGrid, SES for order confirmations, notifications)
    - Search (Elasticsearch or Algolia for product search - optional initially)
