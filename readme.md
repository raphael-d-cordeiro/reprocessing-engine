# ⚙️ Reprocessing Engine

Asynchronous and scalable engine for automatically reprocessing proposals with `waiting-for-integration` status, built with an event-driven architecture using Go, NATS, and Redis.

---

## 📌 Purpose

To automate the reprocessing of proposals that failed integration with external partners by providing:

- Full decoupling between failure source and recovery
- Horizontal scalability via NATS JetStream
- Retry mechanism with attempt control, backoff, and scheduled retries
- Observability through structured logging and metrics
- Idempotency and protection against duplicate processing

---

## 🧱 Architecture Overview

- Scanner reads orders from paginated API
- Redis controls retry delay and attempt count (ZSet + Hash)
- Events are published to `Proposal.Processing.Failed` (NATS JetStream)
- Workers consume the event and call the Integrator
- On success, `Proposal.Reprocessing.Succeeded` event may be published