# Requirements Document: URL Shortener Service

**Version:** 1.0  
**Status:** Draft  
**Purpose:** Learning project — language and framework agnostic

---

## 1. Overview

A URL shortener takes a long URL and produces a short, shareable alias. When a user visits the short URL, they are redirected to the original long URL. Classic examples include bit.ly and tinyurl.com.

This document defines what the system must do (functional requirements), how well it must do it (non-functional requirements), and the data and API contracts it must honor.

---

## 2. Goals

- Learn how to design and build a full-stack web service from scratch.
- Understand REST API design, database modeling, and redirect mechanics.
- Practice encoding/hashing strategies for generating unique short codes.
- Optionally explore caching, analytics, and rate limiting as stretch goals.

---

## 3. Actors

**Anonymous User** — visits a short URL and gets redirected. No account required.

**Registered User** — creates and manages their own short links. Has a personal dashboard.

**Admin** — can view all links, deactivate abusive ones, and see system-wide analytics.

---

## 4. Functional Requirements

### 4.1 URL Shortening

- A user must be able to submit a long URL and receive a unique short code (e.g., `abc123`).
- The short URL must be in the format: `https://<domain>/<code>`.
- The system must validate that the submitted URL is well-formed (has a valid scheme and host).
- If the same long URL is submitted again by the same user, the system may return the existing short code rather than creating a duplicate.
- A registered user may optionally provide a custom alias (e.g., `my-blog`) instead of an auto-generated code.
  - Custom aliases must be unique across the system.
  - Custom aliases must only contain alphanumeric characters, hyphens, and underscores.
  - Maximum alias length: 50 characters.

### 4.2 Redirection

- When a user visits a short URL, the system must redirect them to the original long URL.
- The redirect must use HTTP status code **301** (permanent) or **302** (temporary).
  - Use 302 if you want redirect analytics to work accurately (browsers cache 301s).
- If a short code does not exist or has been deactivated, the system must return a **404** page with a friendly error message.
- If a short URL has expired, the system must return a **410 Gone** response.

### 4.3 Link Management (Registered Users)

- A registered user must be able to:
  - View a list of all their short links.
  - Delete any of their own short links.
  - Set an optional expiration date on a link at creation time.
  - See basic click statistics for each link (total clicks, last clicked at).
- A registered user must NOT be able to view or manage links belonging to other users.

### 4.4 User Accounts

- Users must be able to register with an email address and password.
- Passwords must be stored hashed (never in plain text).
- Users must be able to log in and receive a session token or JWT.
- Users must be able to log out, invalidating their session.
- Anonymous users may still shorten URLs, but their links will not be manageable after the session ends (no ownership).

### 4.5 Analytics (Stretch Goal)

- For each redirect event, the system should optionally record:
  - Timestamp
  - Referrer URL (if present in request headers)
  - Browser / User-Agent string
  - Country derived from IP (using a geolocation lookup)
- A registered user should be able to view this per-link analytics data on their dashboard.

### 4.6 Rate Limiting

- Anonymous users must be limited to a reasonable number of URL creations per hour (e.g., 10 per IP per hour).
- Registered users may have a higher limit (e.g., 100 per hour).
- When the limit is exceeded, the API must return **HTTP 429 Too Many Requests** with a `Retry-After` header.

---

## 5. Non-Functional Requirements

| Concern               | Requirement                                                                                |
| --------------------- | ------------------------------------------------------------------------------------------ |
| **Availability**      | The service should target 99.9% uptime.                                                    |
| **Redirect latency**  | Redirects should complete in under 100ms at the 95th percentile.                           |
| **Short code length** | Auto-generated codes should be 6–8 characters.                                             |
| **Scalability**       | Architecture should allow horizontal scaling of the web layer without redesign.            |
| **Security**          | HTTPS only. All user inputs sanitized. No open redirects to `javascript:` or `data:` URIs. |
| **Data retention**    | Expired or deleted links should be purged after 30 days (soft delete).                     |
| **Observability**     | Errors must be logged with a correlation ID for debugging.                                 |

---

## 6. Short Code Generation

This is a key engineering decision. Two common approaches:

**Option A — Random base62 string**  
Generate a random string of characters from `[a-zA-Z0-9]`. Check for collisions in the database on insert. Simple to implement. Collision probability is low with 6+ characters (62^6 ≈ 56 billion combinations).

**Option B — Encode a database ID**  
Insert the record first to get an auto-increment ID, then base62-encode that ID as the code. No collision risk. Predictable and slightly sequential, which some consider a privacy concern.

For a learning project, **Option A** is recommended for simplicity. For production at scale, **Option B** (or a distributed ID scheme like Snowflake) is more reliable.

---

## 7. Caching Strategy (Stretch Goal)

The hot path (redirect lookup) is read-heavy and should be cached.

- Use an in-memory cache (e.g., Redis, Memcached, or even a local in-process LRU cache).
- Cache key: the short code. Cache value: the original URL.
- TTL: match the link's `expires_at`, or default to 24 hours.
- On link deletion or expiry, invalidate the cache entry.

With a cache, redirect latency can drop from ~10ms (DB round-trip) to under 1ms.

---

## 8. Out of Scope (for v1)

- QR code generation
- A/B testing across multiple destination URLs
- Link click fraud detection
- OAuth / social login
- Bulk link import via CSV
- Browser extension

These can be added in later versions once the core system is working.

---

## 9. Milestones

| Phase | Deliverable                                                                                    |
| ----- | ---------------------------------------------------------------------------------------------- |
| 1     | Core redirect: POST to shorten, GET to redirect. No auth, no DB persistence (in-memory store). |
| 2     | Persistence: swap in-memory store for a real database. Add expiration support.                 |
| 3     | Auth: user registration, login, link ownership.                                                |
| 4     | Dashboard UI: list, delete, view click counts.                                                 |
| 5     | Rate limiting + input validation hardening.                                                    |
| 6     | Caching layer + analytics events.                                                              |

---

_End of document._
