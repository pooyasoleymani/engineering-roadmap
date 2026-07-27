# go/

```text
go/

basics/

runtime/

concurrency/

networking/

protobuf/

grpc/

reflection/

generics/

unsafe/

projects/
```

---

Structure:

```
go/
│
├── README.md
│
├── 01-go-fundamentals/
│   ├── 01-variables.md
│   ├── 02-functions.md
│   ├── 03-pointers.md
│   ├── 04-structs.md
│   ├── 05-interfaces.md
│   ├── 06-errors.md
│   ├── 07-generics.md
│   ├── 08-packages-modules.md
│   └── 09-testing.md
│
├── 02-concurrency/
│   ├── 01-goroutines.md
│   ├── 02-channels.md
│   ├── 03-select.md
│   ├── 04-context.md
│   ├── 05-worker-pool.md
│   ├── 06-sync-package.md
│   ├── 07-memory-leaks.md
│   └── 08-concurrency-patterns.md
│
├── 03-http/
│   ├── 01-http-server.md
│   ├── 02-handlers.md
│   ├── 03-json.md
│   ├── 04-middleware.md
│   ├── 05-router-chi.md
│   └── 06-http-testing.md
│
├── 04-database/
│   ├── 01-database-sql.md
│   ├── 02-mysql.md
│   ├── 03-repository-pattern.md
│   ├── 04-transactions.md
│   ├── 05-migrations.md
│   └── 06-database-testing.md
│
├── 05-backend-architecture/
│   ├── 01-service-layer.md
│   ├── 02-dependency-injection.md
│   ├── 03-interfaces.md
│   ├── 04-project-layout.md
│   └── 05-configuration.md
│
├── 06-authentication/
│   ├── 01-password-hashing.md
│   ├── 02-jwt.md
│   ├── 03-auth-middleware.md
│   ├── 04-refresh-tokens.md
│   └── 05-role-based-access-control.md
│
├── 07-production/
│   ├── 01-structured-logging.md
│   ├── 02-graceful-shutdown.md
│   ├── 03-docker.md
│   ├── 04-docker-compose.md
│   ├── 05-ci-cd.md
│   ├── 06-observability.md
│   └── 07-deployment.md
│
├── 08-messaging/
│   ├── 01-nats.md
│   ├── 02-rabbitmq.md
│   └── 03-kafka.md
│
├── 09-caching/
│   ├── 01-redis.md
│   └── 02-cache-patterns.md
│
├── 10-grpc/
│   ├── 01-protobuf.md
│   ├── 02-grpc-server.md
│   └── 03-grpc-client.md
│
├── 11-system-design/
│   ├── 01-load-balancing.md
│   ├── 02-database-indexes.md
│   ├── 03-rate-limiting.md
│   ├── 04-message-queues.md
│   └── 05-distributed-systems.md
│
└── exercises/
    ├── book-api/
    ├── bank-api/
    └── chat-server/
```