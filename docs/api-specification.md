# 📄 API SPECIFICATION DOCUMENT
## WebifyLab Portal — Phase 1 (MVP)

---

### 1. DOCUMENT INFORMATION

| Field | Detail |
|---|---|
| **Product** | WebifyLab Portal |
| **Document Type** | REST API Specification |
| **Version** | 1.0.0 |
| **Date** | 05 September 2026 |
| **Base URL** | `https://api.webifylab.my.id/api/v1` |
| **Protocol** | HTTPS (mandatory) |
| **Format** | JSON |
| **Author** | Rizal (System Analyst & Fullstack Developer) |
| **Status** | Approved — Ready for Implementation |

---

### 2. OVERVIEW & CONVENTIONS

#### 2.1 Base URL

| Environment | URL |
|---|---|
| **Production** | `https://api.webifylab.my.id/api/v1` |
| **Development** | `http://localhost:8080/api/v1` |

#### 2.2 Authentication

| Type | Implementation |
|---|---|
| **Method** | Bearer Token (JWT) |
| **Header** | `Authorization: Bearer <access_token>` |
| **Token Expiry** | Access token: 24 hours, Refresh token: 7 days |
| **Public Endpoints** | Tidak perlu header Authorization |
| **Admin Endpoints** | Wajib Authorization header + role check |

#### 2.3 Content Type

- **Request:** `Content-Type: application/json` (kecuali upload file: `multipart/form-data`)
- **Response:** `application/json`

#### 2.4 Standard Response Format

**Success Response:**
```json
{
  "success": true,
  "message": "Operation completed successfully",
  "data": { ... }
}
```

**Success Response with Pagination:**
```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": [ ... ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 50,
    "total_pages": 5,
    "has_next": true,
    "has_prev": false
  }
}
```

**Error Response:**
```json
{
  "success": false,
  "message": "Validation failed",
  "errors": [
    {
      "field": "email",
      "message": "Email is required"
    },
    {
      "field": "password",
      "message": "Password must be at least 8 characters"
    }
  ]
}
```

**Single Error Response:**
```json
{
  "success": false,
  "message": "Resource not found",
  "errors": []
}
```

#### 2.5 HTTP Status Codes

| Code | Meaning | Usage |
|---|---|---|
| `200` | OK | Successful GET, PUT, DELETE |
| `201` | Created | Successful POST (resource created) |
| `400` | Bad Request | Invalid request body/params |
| `401` | Unauthorized | Missing or invalid token |
| `403` | Forbidden | Token valid but insufficient permissions |
| `404` | Not Found | Resource doesn't exist |
| `409` | Conflict | Duplicate resource (e.g., email exists) |
| `422` | Unprocessable Entity | Validation errors |
| `429` | Too Many Requests | Rate limit exceeded |
| `500` | Internal Server Error | Server-side error |

#### 2.6 Naming Conventions

| Element | Convention | Example |
|---|---|---|
| Endpoint paths | kebab-case, plural | `/case-studies`, `/blog-posts` |
| Query params | snake_case | `?page=1&per_page=10` |
| JSON keys | snake_case | `created_at`, `client_name` |
| Timestamps | ISO 8601 with timezone | `2026-09-05T10:30:00Z` |

#### 2.7 Common Query Parameters

| Parameter | Type | Default | Description |
|---|---|---|---|
| `page` | integer | 1 | Page number (1-indexed) |
| `limit` | integer | 10 | Items per page (max: 100) |
| `search` | string | - | Search keyword |
| `sort_by` | string | `created_at` | Sort field |
| `sort_order` | string | `desc` | `asc` or `desc` |
| `status` | string | - | Filter by status |
| `category` | string | - | Filter by category slug |
| `tag` | string | - | Filter by tag slug |

---

### 3. AUTHENTICATION FLOW

#### 3.1 JWT Flow Diagram

```
┌─────────┐                          ┌─────────┐
│ Client  │                          │ Backend │
│(Next.js)│                          │ (Gin)   │
└────┬────┘                          └────┬────┘
     │                                    │
     │  1. POST /auth/login               │
     │  { email, password }               │
     │ ─────────────────────────────────► │
     │                                    │
     │  2. { access_token, refresh_token }│
     │ ◄───────────────────────────────── │
     │                                    │
     │  3. Store tokens (httpOnly cookie) │
     │                                    │
     │  4. GET /blog (protected)          │
     │  Authorization: Bearer <token>     │
     │ ─────────────────────────────────► │
     │                                    │
     │  5. Validate JWT                   │
     │  Return data                       │
     │ ◄───────────────────────────────── │
     │                                    │
     │  6. Token expired?                 │
     │  POST /auth/refresh                │
     │  { refresh_token }                 │
     │ ─────────────────────────────────► │
     │                                    │
     │  7. { new_access_token }           │
     │ ◄───────────────────────────────── │
     │                                    │
```

#### 3.2 Token Storage Strategy (Frontend)

```typescript
// frontend/src/lib/auth.ts
// Access token: httpOnly cookie (secure, tidak bisa diakses JS)
// Refresh token: httpOnly cookie (secure, tidak bisa diakses JS)
// Token dikirim otomatis oleh browser ke backend
```

#### 3.3 Role-Based Access Control

| Role | Permissions |
|---|---|
| `super_admin` | Full access (manage users, settings, all content) |
| `admin` | Manage all content, leads, media |
| `editor` | Manage blog, case studies, portfolios (no settings/users) |

---

### 4. ENDPOINTS REFERENCE

---

#### 4.1 AUTHENTICATION ENDPOINTS

##### 🔐 `POST /auth/login`

Login dan dapatkan JWT tokens.

**Auth:** Public  
**Rate Limit:** 5 requests/minute/IP

**Request Body:**
```json
{
  "email": "admin@webifylab.my.id",
  "password": "password123"
}
```

**Validation:**
- `email`: required, valid email format
- `password`: required, min 8 characters

**Success Response (200):**
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "dGhpcyBpcyBhIHJlZnJl...",
    "expires_in": 86400,
    "user": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "name": "Rizal",
      "email": "admin@webifylab.my.id",
      "role": "super_admin"
    }
  }
}
```

**Error Responses:**

| Code | Message |
|---|---|
| 400 | `{"message": "Email and password are required"}` |
| 401 | `{"message": "Invalid email or password"}` |
| 429 | `{"message": "Too many login attempts. Try again in 1 minute"}` |

---

##### 🔐 `POST /auth/refresh`

Refresh access token menggunakan refresh token.

**Auth:** Public  
**Rate Limit:** 30 requests/minute/IP

**Request Body:**
```json
{
  "refresh_token": "dGhpcyBpcyBhIHJlZnJl..."
}
```

**Success Response (200):**
```json
{
  "success": true,
  "message": "Token refreshed successfully",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "expires_in": 86400
  }
}
```

**Error Responses:**

| Code | Message |
|---|---|
| 400 | `{"message": "Refresh token is required"}` |
| 401 | `{"message": "Invalid or expired refresh token"}` |

---

##### 🔐 `POST /auth/logout`

Invalidate refresh token (optional, client-side sufficient).

**Auth:** Required

**Success Response (200):**
```json
{
  "success": true,
  "message": "Logout successful",
  "data": null
}
```

---

##### 🔐 `GET /auth/me`

Get current user profile.

**Auth:** Required

**Success Response (200):**
```json
{
  "success": true,
  "message": "User profile retrieved",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "Rizal",
    "email": "admin@webifylab.my.id",
    "role": "super_admin",
    "created_at": "2026-09-05T10:00:00Z"
  }
}
```

---

##### 🔐 `POST /auth/forgot-password`

Request password reset email.

**Auth:** Public  
**Rate Limit:** 3 requests/hour/IP

**Request Body:**
```json
{
  "email": "admin@webifylab.my.id"
}
```

**Success Response (200):**
```json
{
  "success": true,
  "message": "If your email is registered, you will receive a password reset link",
  "data": null
}
```

**Note:** Response selalu success untuk mencegah email enumeration.

---

##### 🔐 `POST /auth/reset-password`

Reset password menggunakan token dari email.

**Auth:** Public

**Request Body:**
```json
{
  "token": "reset-token-from-email",
  "password": "newpassword123",
  "password_confirmation": "newpassword123"
}
```

**Validation:**
- `token`: required
- `password`: required, min 8 characters
- `password_confirmation`: required, must match password

**Success Response (200):**
```json
{
  "success": true,
  "message": "Password reset successful. Please login with your new password.",
  "data": null
}
```

**Error Responses:**

| Code | Message |
|---|---|
| 400 | `{"message": "Token is required"}` |
| 422 | `{"message": "Invalid or expired token"}` |
| 422 | `{"message": "Password confirmation does not match"}` |

---

#### 4.2 BLOG ENDPOINTS

##### 📝 `GET /blog`

List blog posts (public).

**Auth:** Public  
**Cache:** 5 minutes

**Query Parameters:**
| Param | Type | Default | Description |
|---|---|---|---|
| `page` | integer | 1 | Page number |
| `limit` | integer | 10 | Items per page (max 100) |
| `search` | string | - | Search in title, excerpt |
| `category` | string | - | Filter by category slug |
| `tag` | string | - | Filter by tag slug |
| `status` | string | `published` | Only `published` for public |

**Example Request:**
```
GET /blog?page=1&limit=10&category=tutorial&tag=nextjs
```

**Success Response (200):**
```json
{
  "success": true,
  "message": "Blog posts retrieved",
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440001",
      "title": "Membuat REST API dengan Golang dan Gin",
      "slug": "membuat-rest-api-golang-gin",
      "excerpt": "Panduan lengkap membuat REST API...",
      "cover_image": "https://api.webifylab.my.id/uploads/blog/cover-1.jpg",
      "read_time": 8,
      "published_at": "2026-09-01T10:00:00Z",
      "author": {
        "id": "550e8400-e29b-41d4-a716-446655440000",
        "name": "Rizal"
      },
      "category": {
        "id": "...",
        "name": "Tutorial",
        "slug": "tutorial"
      },
      "tags": [
        { "id": "...", "name": "Golang", "slug": "golang" },
        { "id": "...", "name": "Next.js", "slug": "nextjs" }
      ]
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 25,
    "total_pages": 3,
    "has_next": true,
    "has_prev": false
  }
}
```

---

##### 📝 `GET /blog/:slug`

Get single blog post by slug (public).

**Auth:** Public

**Example Request:**
```
GET /blog/membuat-rest-api-golang-gin
```

**Success Response (200):**
```json
{
  "success": true,
  "message": "Blog post retrieved",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "title": "Membuat REST API dengan Golang dan Gin",
    "slug": "membuat-rest-api-golang-gin",
    "content": "<h2>Pendahuluan</h2><p>...</p>",
    "excerpt": "Panduan lengkap membuat REST API...",
    "cover_image": "https://api.webifylab.my.id/uploads/blog/cover-1.jpg",
    "read_time": 8,
    "status": "published",
    "meta_title": "Membuat REST API dengan Golang dan Gin | WebifyLab",
    "meta_description": "Panduan lengkap...",
    "published_at": "2026-09-01T10:00:00Z",
    "created_at": "2026-08-30T10:00:00Z",
    "updated_at": "2026-09-01T09:00:00Z",
    "author": {
      "id": "...",
      "name": "Rizal",
      "email": "admin@webifylab.my.id"
    },
    "category": {
      "id": "...",
      "name": "Tutorial",
      "slug": "tutorial"
    },
    "tags": [
      { "id": "...", "name": "Golang", "slug": "golang" },
      { "id": "...", "name": "Next.js", "slug": "nextjs" }
    ],
    "related_posts": [
      {
        "id": "...",
        "title": "...",
        "slug": "...",
        "cover_image": "...",
        "published_at": "..."
      }
    ]
  }
}
```

**Error Responses:**

| Code | Message |
|---|---|
| 404 | `{"message": "Blog post not found"}` |

---

##### 📝 `POST /blog`

Create new blog post.

**Auth:** Required (admin, editor)

**Request Body:**
```json
{
  "title": "Membuat REST API dengan Golang dan Gin",
  "slug": "membuat-rest-api-golang-gin",
  "content": "<h2>Pendahuluan</h2><p>...</p>",
  "excerpt": "Panduan lengkap membuat REST API...",
  "cover_image": "https://api.webifylab.my.id/uploads/blog/cover-1.jpg",
  "category_id": "550e8400-e29b-41d4-a716-446655440010",
  "tags": ["golang", "nextjs"],
  "status": "draft",
  "meta_title": "Membuat REST API dengan Golang dan Gin | WebifyLab",
  "meta_description": "Panduan lengkap...",
  "read_time": 8
}
```

**Validation:**
- `title`: required, max 255 chars
- `slug`: auto-generated from title if not provided, must be unique
- `content`: required
- `status`: must be `draft`, `published`, or `archived`
- `tags`: array of tag slugs (optional)
- If `status` = `published`, `published_at` auto-set to NOW()

**Success Response (201):**
```json
{
  "success": true,
  "message": "Blog post created successfully",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "title": "Membuat REST API dengan Golang dan Gin",
    "slug": "membuat-rest-api-golang-gin",
    "status": "draft",
    "created_at": "2026-09-05T10:00:00Z"
  }
}
```

---

##### 📝 `PUT /blog/:id`

Update blog post.

**Auth:** Required (admin, editor)

**Request Body:** (same as POST, all fields optional)
```json
{
  "title": "Updated Title",
  "status": "published"
}
```

**Success Response (200):**
```json
{
  "success": true,
  "message": "Blog post updated successfully",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "title": "Updated Title",
    "status": "published",
    "published_at": "2026-09-05T11:00:00Z",
    "updated_at": "2026-09-05T11:00:00Z"
  }
}
```

---

##### 📝 `DELETE /blog/:id`

Soft delete blog post.

**Auth:** Required (admin, editor)

**Success Response (200):**
```json
{
  "success": true,
  "message": "Blog post deleted successfully",
  "data": null
}
```

---

##### 📝 `GET /blog/admin`

List all blog posts (admin view, includes drafts).

**Auth:** Required (admin, editor)

**Query Parameters:** Same as `GET /blog` + `status` can be `draft`, `published`, `archived`, or `all`

**Response:** Same structure as `GET /blog`

---

#### 4.3 CASE STUDY ENDPOINTS

##### 📊 `GET /case-studies`

List case studies (public).

**Auth:** Public

**Query Parameters:**
| Param | Type | Default | Description |
|---|---|---|---|
| `page` | integer | 1 | Page number |
| `limit` | integer | 10 | Items per page |
| `search` | string | - | Search in title, client_name |
| `industry` | string | - | Filter by industry |
| `ai_type` | string | - | Filter by AI type (spk, sistem_pakar, dll) |
| `tag` | string | - | Filter by tag slug |

**Example Request:**
```
GET /case-studies?industry=retail&ai_type=spk
```

**Success Response (200):**
```json
{
  "success": true,
  "message": "Case studies retrieved",
  "data": [
    {
      "id": "...",
      "title": "Optimasi Stok UMKM dengan SPK",
      "slug": "optimasi-stok-umkm-spk",
      "client_name": "Toko Sejahtera",
      "industry": "Retail",
      "cover_image": "...",
      "metrics": {
        "stock_accuracy": "95%",
        "waste_reduction": "40%"
      },
      "ai_type": ["spk", "data_analytics"],
      "published_at": "2026-08-15T10:00:00Z",
      "tags": [...]
    }
  ],
  "pagination": { ... }
}
```

---

##### 📊 `GET /case-studies/:slug`

Get single case study (public).

**Auth:** Public

**Success Response (200):**
```json
{
  "success": true,
  "message": "Case study retrieved",
  "data": {
    "id": "...",
    "title": "Optimasi Stok UMKM dengan SPK",
    "slug": "optimasi-stok-umkm-spk",
    "client_name": "Toko Sejahtera",
    "industry": "Retail",
    "challenge": "Toko mengalami kesulitan mengelola stok...",
    "approach": "Kami menggunakan metode SPK dengan algoritma...",
    "solution": "Sistem yang kami bangun menganalisis data penjualan...",
    "ai_type": ["spk", "data_analytics"],
    "results": "Setelah 3 bulan implementasi, akurasi stok meningkat...",
    "lessons": "Pelajaran utama adalah pentingnya data yang bersih...",
    "testimonial": "WebifyLab membantu kami...",
    "metrics": {
      "stock_accuracy": "95%",
      "waste_reduction": "40%",
      "time_saved": "10 hours/week"
    },
    "cover_image": "...",
    "images": [
      "https://api.webifylab.my.id/uploads/case-studies/img-1.jpg",
      "https://api.webifylab.my.id/uploads/case-studies/img-2.jpg"
    ],
    "meta_title": "...",
    "meta_description": "...",
    "published_at": "2026-08-15T10:00:00Z",
    "tags": [...]
  }
}
```

---

##### 📊 `POST /case-studies`

Create new case study.

**Auth:** Required (admin, editor)

**Request Body:**
```json
{
  "title": "Optimasi Stok UMKM dengan SPK",
  "slug": "optimasi-stok-umkm-spk",
  "client_name": "Toko Sejahtera",
  "industry": "Retail",
  "challenge": "Toko mengalami kesulitan...",
  "approach": "Kami menggunakan metode...",
  "solution": "Sistem yang kami bangun...",
  "ai_type": ["spk", "data_analytics"],
  "results": "Setelah 3 bulan...",
  "lessons": "Pelajaran utama...",
  "testimonial": "WebifyLab membantu...",
  "metrics": {
    "stock_accuracy": "95%",
    "waste_reduction": "40%"
  },
  "cover_image": "...",
  "images": ["url1", "url2"],
  "tags": ["spk", "umkm"],
  "status": "draft",
  "meta_title": "...",
  "meta_description": "..."
}
```

**Validation:**
- `title`, `client_name`, `industry`, `challenge`, `approach`, `solution`, `results`: required
- `ai_type`: array, values must be `spk`, `sistem_pakar`, `data_analytics`, `machine_learning`
- `metrics`: JSON object (optional)
- `images`: array of URLs (optional)

**Success Response (201):**
```json
{
  "success": true,
  "message": "Case study created successfully",
  "data": { "id": "...", "slug": "..." }
}
```

---

##### 📊 `PUT /case-studies/:id`

Update case study.

**Auth:** Required (admin, editor)

**Request Body:** (same as POST, all fields optional)

**Success Response (200):**
```json
{
  "success": true,
  "message": "Case study updated successfully",
  "data": { "id": "...", "updated_at": "..." }
}
```

---

##### 📊 `DELETE /case-studies/:id`

Soft delete case study.

**Auth:** Required (admin, editor)

**Success Response (200):**
```json
{
  "success": true,
  "message": "Case study deleted successfully",
  "data": null
}
```

---

#### 4.4 PORTFOLIO ENDPOINTS

##### 💼 `GET /portfolio`

List portfolio items (public).

**Auth:** Public

**Query Parameters:**
| Param | Type | Default | Description |
|---|---|---|---|
| `page` | integer | 1 | Page number |
| `limit` | integer | 12 | Items per page |
| `category` | string | - | Filter by category (web, mobile, ai, ecommerce, other) |
| `search` | string | - | Search in title, description |

**Success Response (200):**
```json
{
  "success": true,
  "message": "Portfolio items retrieved",
  "data": [
    {
      "id": "...",
      "title": "E-Commerce Platform untuk UMKM",
      "slug": "ecommerce-platform-umkm",
      "description": "Platform e-commerce...",
      "category": "ecommerce",
      "tech_stack": ["Next.js", "Golang", "PostgreSQL"],
      "thumbnail": "...",
      "live_url": "https://example.com",
      "published_at": "2026-08-20T10:00:00Z"
    }
  ],
  "pagination": { ... }
}
```

---

##### 💼 `GET /portfolio/:slug`

Get single portfolio item (public).

**Auth:** Public

**Success Response (200):**
```json
{
  "success": true,
  "message": "Portfolio item retrieved",
  "data": {
    "id": "...",
    "title": "E-Commerce Platform untuk UMKM",
    "slug": "ecommerce-platform-umkm",
    "description": "Platform e-commerce lengkap...",
    "category": "ecommerce",
    "tech_stack": ["Next.js", "Golang", "PostgreSQL", "Docker"],
    "live_url": "https://example.com",
    "thumbnail": "...",
    "images": ["url1", "url2", "url3"],
    "lessons": "Pelajaran utama dalam project ini...",
    "meta_title": "...",
    "meta_description": "...",
    "published_at": "2026-08-20T10:00:00Z",
    "related_portfolios": [ ... ]
  }
}
```

---

##### 💼 `POST /portfolio`

Create new portfolio item.

**Auth:** Required (admin, editor)

**Request Body:**
```json
{
  "title": "E-Commerce Platform untuk UMKM",
  "slug": "ecommerce-platform-umkm",
  "description": "Platform e-commerce...",
  "category": "ecommerce",
  "tech_stack": ["Next.js", "Golang", "PostgreSQL"],
  "live_url": "https://example.com",
  "thumbnail": "...",
  "images": ["url1", "url2"],
  "lessons": "Pelajaran utama...",
  "order_index": 1,
  "status": "published",
  "meta_title": "...",
  "meta_description": "..."
}
```

**Validation:**
- `title`, `description`, `category`: required
- `category`: must be `web`, `mobile`, `ai`, `ecommerce`, `other`
- `tech_stack`, `images`: array of strings (optional)

**Success Response (201):**
```json
{
  "success": true,
  "message": "Portfolio item created successfully",
  "data": { "id": "...", "slug": "..." }
}
```

---

##### 💼 `PUT /portfolio/:id`

Update portfolio item.

**Auth:** Required (admin, editor)

**Request Body:** (same as POST, all fields optional)

**Success Response (200):**
```json
{
  "success": true,
  "message": "Portfolio item updated successfully",
  "data": { "id": "...", "updated_at": "..." }
}
```

---

##### 💼 `PUT /portfolio/reorder`

Reorder portfolio items.

**Auth:** Required (admin, editor)

**Request Body:**
```json
{
  "items": [
    { "id": "uuid-1", "order_index": 1 },
    { "id": "uuid-2", "order_index": 2 },
    { "id": "uuid-3", "order_index": 3 }
  ]
}
```

**Success Response (200):**
```json
{
  "success": true,
  "message": "Portfolio items reordered successfully",
  "data": null
}
```

---

##### 💼 `DELETE /portfolio/:id`

Soft delete portfolio item.

**Auth:** Required (admin, editor)

**Success Response (200):**
```json
{
  "success": true,
  "message": "Portfolio item deleted successfully",
  "data": null
}
```

---

#### 4.5 SERVICE ENDPOINTS

##### 🛠️ `GET /services`

List all services (public, only active).

**Auth:** Public

**Success Response (200):**
```json
{
  "success": true,
  "message": "Services retrieved",
  "data": [
    {
      "id": "...",
      "name": "Web Development",
      "slug": "web-development",
      "description": "Pembuatan website profesional...",
      "icon": "globe",
      "use_cases": [
        "Company profile untuk UMKM",
        "Landing page untuk produk digital"
      ],
      "timeline_estimate": "2-6 minggu",
      "tech_stack": ["Next.js", "Golang", "PostgreSQL"],
      "learning_points": "Mari kita eksperimen bareng...",
      "order_index": 1
    }
  ]
}
```

---

##### 🛠️ `GET /services/:slug`

Get single service (public).

**Auth:** Public

**Success Response (200):**
```json
{
  "success": true,
  "message": "Service retrieved",
  "data": {
    "id": "...",
    "name": "Web Development",
    "slug": "web-development",
    "description": "Pembuatan website profesional...",
    "icon": "globe",
    "use_cases": [...],
    "timeline_estimate": "2-6 minggu",
    "tech_stack": [...],
    "learning_points": "...",
    "is_active": true,
    "order_index": 1
  }
}
```

---

##### 🛠️ `POST /services`

Create new service.

**Auth:** Required (admin)

**Request Body:**
```json
{
  "name": "AI Integration",
  "slug": "ai-integration",
  "description": "Integrasi algoritma AI...",
  "icon": "brain",
  "use_cases": ["Sistem Pendukung Keputusan", "Sistem Pakar"],
  "timeline_estimate": "3-8 minggu",
  "tech_stack": ["Python", "Golang", "TensorFlow"],
  "learning_points": "Mari kita eksperimen...",
  "is_active": true,
  "order_index": 4
}
```

**Success Response (201):**
```json
{
  "success": true,
  "message": "Service created successfully",
  "data": { "id": "...", "slug": "..." }
}
```

---

##### 🛠️ `PUT /services/:id`

Update service.

**Auth:** Required (admin)

**Request Body:** (same as POST, all fields optional)

**Success Response (200):**
```json
{
  "success": true,
  "message": "Service updated successfully",
  "data": { "id": "...", "updated_at": "..." }
}
```

---

##### 🛠️ `DELETE /services/:id`

Delete service (hard delete).

**Auth:** Required (admin)

**Success Response (200):**
```json
{
  "success": true,
  "message": "Service deleted successfully",
  "data": null
}
```

---

#### 4.6 LEAD ENDPOINTS

##### 📬 `POST /leads`

Submit new inquiry (public).

**Auth:** Public  
**Rate Limit:** 5 submissions/hour/IP

**Request Body:**
```json
{
  "name": "Budi Santoso",
  "email": "budi@example.com",
  "phone": "6281234567890",
  "service_type": "web_development",
  "budget_range": "5-15jt",
  "description": "Saya ingin membuat website company profile untuk toko saya..."
}
```

**Validation:**
- `name`: required, max 100 chars
- `email`: required, valid email
- `phone`: required, valid phone format (E.164 recommended)
- `service_type`: required, must be `web_development`, `mobile_app`, `uiux_design`, `ai_integration`, `other`
- `budget_range`: required, must be `<5jt`, `5-15jt`, `15-50jt`, `50-100jt`, `>100jt`
- `description`: required, min 20 chars, max 2000 chars

**Success Response (201):**
```json
{
  "success": true,
  "message": "Inquiry submitted successfully. We'll contact you soon via WhatsApp.",
  "data": {
    "id": "...",
    "status": "new"
  }
}
```

**Side Effects:**
1. Save lead to database
2. Send auto-reply email to user via Resend
3. Send notification email to admin via Resend

**Error Responses:**

| Code | Message |
|---|---|
| 400 | `{"message": "All fields are required"}` |
| 422 | `{"errors": [{"field": "email", "message": "Invalid email format"}]}` |
| 429 | `{"message": "Too many submissions. Please try again later."}` |

---

##### 📬 `GET /leads`

List all leads (admin).

**Auth:** Required (admin)

**Query Parameters:**
| Param | Type | Default | Description |
|---|---|---|---|
| `page` | integer | 1 | Page number |
| `limit` | integer | 20 | Items per page |
| `status` | string | - | Filter by status |
| `service_type` | string | - | Filter by service type |
| `search` | string | - | Search in name, email, description |
| `sort_by` | string | `created_at` | Sort field |
| `sort_order` | string | `desc` | `asc` or `desc` |

**Success Response (200):**
```json
{
  "success": true,
  "message": "Leads retrieved",
  "data": [
    {
      "id": "...",
      "name": "Budi Santoso",
      "email": "budi@example.com",
      "phone": "6281234567890",
      "service_type": "web_development",
      "budget_range": "5-15jt",
      "description": "Saya ingin membuat...",
      "status": "new",
      "notes": null,
      "created_at": "2026-09-05T10:00:00Z"
    }
  ],
  "pagination": { ... }
}
```

---

##### 📬 `GET /leads/:id`

Get lead detail (admin).

**Auth:** Required (admin)

**Success Response (200):**
```json
{
  "success": true,
  "message": "Lead retrieved",
  "data": {
    "id": "...",
    "name": "Budi Santoso",
    "email": "budi@example.com",
    "phone": "6281234567890",
    "service_type": "web_development",
    "budget_range": "5-15jt",
    "description": "Saya ingin membuat website...",
    "status": "new",
    "notes": "Sudah dihubungi via WA, tertarik dengan paket basic",
    "ip_address": "103.45.67.89",
    "user_agent": "Mozilla/5.0...",
    "created_at": "2026-09-05T10:00:00Z",
    "updated_at": "2026-09-05T11:00:00Z"
  }
}
```

---

##### 📬 `PUT /leads/:id`

Update lead status/notes (admin).

**Auth:** Required (admin)

**Request Body:**
```json
{
  "status": "contacted",
  "notes": "Sudah dihubungi via WA, tertarik dengan paket basic"
}
```

**Validation:**
- `status`: must be `new`, `contacted`, `negotiation`, `proposal_sent`, `won`, `lost`

**Success Response (200):**
```json
{
  "success": true,
  "message": "Lead updated successfully",
  "data": {
    "id": "...",
    "status": "contacted",
    "notes": "...",
    "updated_at": "2026-09-05T11:00:00Z"
  }
}
```

---

##### 📬 `DELETE /leads/:id`

Delete lead (admin).

**Auth:** Required (admin)

**Success Response (200):**
```json
{
  "success": true,
  "message": "Lead deleted successfully",
  "data": null
}
```

---

##### 📬 `GET /leads/export`

Export leads to CSV.

**Auth:** Required (admin)

**Query Parameters:**
| Param | Type | Default | Description |
|---|---|---|---|
| `status` | string | - | Filter by status |
| `from` | string (date) | - | Filter from date (YYYY-MM-DD) |
| `to` | string (date) | - | Filter to date (YYYY-MM-DD) |

**Response:**
- Content-Type: `text/csv`
- Filename: `leads-YYYY-MM-DD.csv`

**CSV Columns:**
```
id,name,email,phone,service_type,budget_range,description,status,notes,created_at
```

---

#### 4.7 MEDIA ENDPOINTS

##### 🖼️ `POST /media/upload`

Upload image file.

**Auth:** Required (admin, editor)  
**Content-Type:** `multipart/form-data`

**Request Body:**
```
file: <binary> (required)
alt_text: string (optional)
```

**Validation:**
- `file`: required, max 5MB
- `file type`: must be `image/jpeg`, `image/png`, `image/webp`
- `alt_text`: optional, max 255 chars

**Success Response (201):**
```json
{
  "success": true,
  "message": "File uploaded successfully",
  "data": {
    "id": "...",
    "filename": "abc123-def456.jpg",
    "original_name": "photo.jpg",
    "url": "https://api.webifylab.my.id/uploads/media/abc123-def456.jpg",
    "mime_type": "image/jpeg",
    "size": 123456,
    "width": 1920,
    "height": 1080,
    "alt_text": "Photo description",
    "created_at": "2026-09-05T10:00:00Z"
  }
}
```

**Error Responses:**

| Code | Message |
|---|---|
| 400 | `{"message": "File is required"}` |
| 413 | `{"message": "File size exceeds 5MB limit"}` |
| 415 | `{"message": "Invalid file type. Only JPG, PNG, WebP allowed"}` |

---

##### 🖼️ `GET /media`

List uploaded media (admin).

**Auth:** Required (admin, editor)

**Query Parameters:**
| Param | Type | Default | Description |
|---|---|---|---|
| `page` | integer | 1 | Page number |
| `limit` | integer | 20 | Items per page |
| `mime_type` | string | - | Filter by MIME type |
| `search` | string | - | Search in filename, original_name, alt_text |

**Success Response (200):**
```json
{
  "success": true,
  "message": "Media retrieved",
  "data": [
    {
      "id": "...",
      "filename": "abc123.jpg",
      "original_name": "photo.jpg",
      "url": "https://api.webifylab.my.id/uploads/media/abc123.jpg",
      "mime_type": "image/jpeg",
      "size": 123456,
      "width": 1920,
      "height": 1080,
      "alt_text": "...",
      "created_at": "..."
    }
  ],
  "pagination": { ... }
}
```

---

##### 🖼️ `DELETE /media/:id`

Delete media file.

**Auth:** Required (admin, editor)

**Success Response (200):**
```json
{
  "success": true,
  "message": "Media deleted successfully",
  "data": null
}
```

**Note:** File will be deleted from storage. If media is referenced in content, URL will become broken.

---

#### 4.8 CATEGORY ENDPOINTS

##### 🏷️ `GET /categories`

List categories.

**Auth:** Public (for filtering)

**Query Parameters:**
| Param | Type | Default | Description |
|---|---|---|---|
| `type` | string | - | Filter by type (`blog`, `portfolio`) |

**Success Response (200):**
```json
{
  "success": true,
  "message": "Categories retrieved",
  "data": [
    {
      "id": "...",
      "name": "Tutorial",
      "slug": "tutorial",
      "type": "blog"
    },
    {
      "id": "...",
      "name": "Web Development",
      "slug": "web-development",
      "type": "portfolio"
    }
  ]
}
```

---

##### 🏷️ `POST /categories`

Create new category.

**Auth:** Required (admin)

**Request Body:**
```json
{
  "name": "AI & Data",
  "slug": "ai-data",
  "type": "blog"
}
```

**Validation:**
- `name`: required, max 100 chars
- `slug`: auto-generated from name if not provided, must be unique
- `type`: required, must be `blog` or `portfolio`

**Success Response (201):**
```json
{
  "success": true,
  "message": "Category created successfully",
  "data": { "id": "...", "name": "...", "slug": "...", "type": "..." }
}
```

---

##### 🏷️ `PUT /categories/:id`

Update category.

**Auth:** Required (admin)

**Request Body:** (same as POST, all fields optional)

**Success Response (200):**
```json
{
  "success": true,
  "message": "Category updated successfully",
  "data": { "id": "...", "updated_at": "..." }
}
```

---

##### 🏷️ `DELETE /categories/:id`

Delete category.

**Auth:** Required (admin)

**Success Response (200):**
```json
{
  "success": true,
  "message": "Category deleted successfully",
  "data": null
}
```

**Note:** Blog posts with this category will have `category_id` set to NULL.

---

#### 4.9 TAG ENDPOINTS

##### 🏷️ `GET /tags`

List all tags.

**Auth:** Public

**Success Response (200):**
```json
{
  "success": true,
  "message": "Tags retrieved",
  "data": [
    { "id": "...", "name": "Golang", "slug": "golang" },
    { "id": "...", "name": "SPK", "slug": "spk" }
  ]
}
```

---

##### 🏷️ `POST /tags`

Create new tag.

**Auth:** Required (admin, editor)

**Request Body:**
```json
{
  "name": "Machine Learning",
  "slug": "machine-learning"
}
```

**Success Response (201):**
```json
{
  "success": true,
  "message": "Tag created successfully",
  "data": { "id": "...", "name": "...", "slug": "..." }
}
```

---

##### 🏷️ `DELETE /tags/:id`

Delete tag.

**Auth:** Required (admin)

**Success Response (200):**
```json
{
  "success": true,
  "message": "Tag deleted successfully",
  "data": null
}
```

**Note:** Tag will be removed from all posts/case studies (cascade).

---

#### 4.10 DASHBOARD ENDPOINTS

##### 📊 `GET /dashboard/stats`

Get dashboard overview stats.

**Auth:** Required (admin, editor)

**Success Response (200):**
```json
{
  "success": true,
  "message": "Dashboard stats retrieved",
  "data": {
    "total_blog_posts": 25,
    "published_blog_posts": 20,
    "draft_blog_posts": 5,
    "total_case_studies": 8,
    "published_case_studies": 6,
    "total_portfolios": 15,
    "published_portfolios": 12,
    "total_leads": 45,
    "leads_this_month": 12,
    "leads_by_status": {
      "new": 5,
      "contacted": 3,
      "negotiation": 2,
      "proposal_sent": 1,
      "won": 1,
      "lost": 0
    },
    "total_media": 150
  }
}
```

---

##### 📊 `GET /dashboard/leads-chart`

Get leads per month (last 12 months).

**Auth:** Required (admin)

**Success Response (200):**
```json
{
  "success": true,
  "message": "Leads chart data retrieved",
  "data": [
    { "month": "2025-10", "count": 5 },
    { "month": "2025-11", "count": 8 },
    { "month": "2025-12", "count": 12 },
    { "month": "2026-01", "count": 10 },
    { "month": "2026-02", "count": 15 },
    { "month": "2026-03", "count": 18 },
    { "month": "2026-04", "count": 14 },
    { "month": "2026-05", "count": 20 },
    { "month": "2026-06", "count": 22 },
    { "month": "2026-07", "count": 25 },
    { "month": "2026-08", "count": 30 },
    { "month": "2026-09", "count": 12 }
  ]
}
```

---

##### 📊 `GET /dashboard/recent-leads`

Get 5 most recent leads.

**Auth:** Required (admin)

**Success Response (200):**
```json
{
  "success": true,
  "message": "Recent leads retrieved",
  "data": [
    {
      "id": "...",
      "name": "Budi Santoso",
      "email": "budi@example.com",
      "phone": "6281234567890",
      "service_type": "web_development",
      "status": "new",
      "created_at": "2026-09-05T10:00:00Z"
    }
  ]
}
```

---

##### 📊 `GET /dashboard/recent-posts`

Get 5 most recent blog posts.

**Auth:** Required (admin, editor)

**Success Response (200):**
```json
{
  "success": true,
  "message": "Recent posts retrieved",
  "data": [
    {
      "id": "...",
      "title": "Membuat REST API dengan Golang",
      "slug": "...",
      "status": "published",
      "published_at": "2026-09-01T10:00:00Z"
    }
  ]
}
```

---

#### 4.11 SETTINGS ENDPOINTS

##### ⚙️ `GET /settings`

Get site settings (public, for site-wide config).

**Auth:** Public

**Success Response (200):**
```json
{
  "success": true,
  "message": "Settings retrieved",
  "data": {
    "site_name": "WebifyLab",
    "tagline": "We Experiment, You Grow",
    "contact_email": "hello@webifylab.my.id",
    "contact_whatsapp": "6281234567890",
    "social_media": {
      "instagram": "https://instagram.com/webifylab",
      "linkedin": "https://linkedin.com/company/webifylab",
      "github": "https://github.com/webifylab"
    },
    "seo_default_title": "WebifyLab - Digital Learning Lab",
    "seo_default_description": "WebifyLab adalah digital learning lab...",
    "seo_og_image": "/images/og-default.jpg"
  }
}
```

---

##### ⚙️ `GET /settings/admin`

Get all settings with metadata (admin).

**Auth:** Required (admin)

**Success Response (200):**
```json
{
  "success": true,
  "message": "Settings retrieved",
  "data": [
    {
      "id": "...",
      "key": "site_name",
      "value": "WebifyLab",
      "description": "Nama website",
      "updated_at": "2026-09-01T10:00:00Z"
    }
  ]
}
```

---

##### ⚙️ `PUT /settings`

Update multiple settings at once.

**Auth:** Required (admin)

**Request Body:**
```json
{
  "settings": [
    { "key": "site_name", "value": "WebifyLab" },
    { "key": "tagline", "value": "We Experiment, You Grow" },
    { "key": "contact_email", "value": "hello@webifylab.my.id" },
    { "key": "contact_whatsapp", "value": "6281234567890" },
    { "key": "social_media", "value": { "instagram": "...", "linkedin": "..." } }
  ]
}
```

**Success Response (200):**
```json
{
  "success": true,
  "message": "Settings updated successfully",
  "data": null
}
```

---

#### 4.12 HEALTH CHECK

##### 🏥 `GET /health`

Health check endpoint for monitoring.

**Auth:** Public

**Success Response (200):**
```json
{
  "success": true,
  "message": "Service is healthy",
  "data": {
    "status": "ok",
    "timestamp": "2026-09-05T10:00:00Z",
    "version": "1.0.0",
    "environment": "production",
    "database": "connected",
    "uptime_seconds": 86400
  }
}
```

**Error Response (503):**
```json
{
  "success": false,
  "message": "Service is unhealthy",
  "data": {
    "status": "error",
    "database": "disconnected"
  }
}
```

---

### 5. ERROR CODES REFERENCE

#### 5.1 Application Error Codes

| Code | HTTP Status | Description |
|---|---|---|
| `AUTH_001` | 400 | Missing email or password |
| `AUTH_002` | 401 | Invalid credentials |
| `AUTH_003` | 401 | Invalid or expired token |
| `AUTH_004` | 401 | Refresh token expired |
| `AUTH_005` | 403 | Insufficient permissions |
| `AUTH_006` | 422 | Invalid or expired reset token |
| `AUTH_007` | 409 | Email already registered |
| `VALID_001` | 422 | Validation failed |
| `VALID_002` | 422 | Invalid field format |
| `RES_001` | 404 | Resource not found |
| `RES_002` | 409 | Resource already exists (duplicate slug) |
| `MEDIA_001` | 413 | File too large |
| `MEDIA_002` | 415 | Invalid file type |
| `RATE_001` | 429 | Rate limit exceeded |
| `SYS_001` | 500 | Internal server error |
| `SYS_002` | 503 | Service unavailable |

#### 5.2 Validation Error Examples

**Email validation:**
```json
{
  "success": false,
  "message": "Validation failed",
  "errors": [
    { "field": "email", "message": "Invalid email format" }
  ]
}
```

**Password validation:**
```json
{
  "success": false,
  "message": "Validation failed",
  "errors": [
    { "field": "password", "message": "Password must be at least 8 characters" }
  ]
}
```

**Multiple field errors:**
```json
{
  "success": false,
  "message": "Validation failed",
  "errors": [
    { "field": "name", "message": "Name is required" },
    { "field": "email", "message": "Invalid email format" },
    { "field": "phone", "message": "Phone is required" }
  ]
}
```

---

### 6. RATE LIMITING

#### 6.1 Rate Limit Headers

Setiap response akan menyertakan header:
```
X-RateLimit-Limit: 100
X-RateLimit-Remaining: 95
X-RateLimit-Reset: 1693910400
```

#### 6.2 Rate Limit per Endpoint

| Endpoint | Limit | Window |
|---|---|---|
| `POST /auth/login` | 5 requests | 1 minute / IP |
| `POST /auth/refresh` | 30 requests | 1 minute / IP |
| `POST /auth/forgot-password` | 3 requests | 1 hour / IP |
| `POST /leads` | 5 requests | 1 hour / IP |
| `POST /media/upload` | 30 requests | 1 minute / User |
| Other endpoints | 100 requests | 1 minute / IP |

#### 6.3 Rate Limit Exceeded Response (429)

```json
{
  "success": false,
  "message": "Too many requests. Please try again later.",
  "errors": [],
  "retry_after": 60
}
```

---

### 7. CORS CONFIGURATION

#### 7.1 Allowed Origins

| Origin | Allowed |
|---|---|
| `https://webifylab.my.id` | ✅ Yes |
| `http://localhost:3000` | ✅ Yes (development) |
| Others | ❌ No |

#### 7.2 Allowed Methods

```
GET, POST, PUT, DELETE, OPTIONS
```

#### 7.3 Allowed Headers

```
Content-Type, Authorization, X-Requested-With
```

#### 7.4 Exposed Headers

```
X-RateLimit-Limit, X-RateLimit-Remaining, X-RateLimit-Reset
```

#### 7.5 Credentials

```
Access-Control-Allow-Credentials: true
```

---

### 8. FILE UPLOAD SPECIFICATION

#### 8.1 Upload Flow

```
┌─────────┐                          ┌─────────┐           ┌─────────┐
│ Client  │                          │ Backend │           │ Storage │
│(Next.js)│                          │ (Gin)   │           │  (VPS)  │
└────┬────┘                          └────┬────┘           └────┬────┘
     │                                    │                     │
     │  1. POST /media/upload             │                     │
     │  multipart/form-data               │                     │
     │ ─────────────────────────────────► │                     │
     │                                    │                     │
     │                                    │  2. Validate file   │
     │                                    │  3. Generate unique │
     │                                    │     filename        │
     │                                    │  4. Compress image  │
     │                                    │  5. Generate WebP   │
     │                                    │ ──────────────────► │
     │                                    │                     │
     │                                    │  6. Save to /uploads│
     │                                    │ ◄────────────────── │
     │                                    │                     │
     │  7. { url: "/uploads/..." }        │                     │
     │ ◄───────────────────────────────── │                     │
     │                                    │                     │
```

#### 8.2 File Constraints

| Constraint | Value |
|---|---|
| Max file size | 5 MB (5,242,880 bytes) |
| Allowed types | `image/jpeg`, `image/png`, `image/webp` |
| Filename format | `{uuid}.{ext}` (e.g., `abc123-def456.jpg`) |
| Storage location | `/app/uploads/media/` (container) → `/var/www/webifylab/uploads/media/` (host) |
| Public URL | `https://api.webifylab.my.id/uploads/media/{filename}` |

#### 8.3 Image Processing

Saat upload, backend akan:
1. **Validate** file type dan size
2. **Generate** unique filename dengan UUID
3. **Compress** image (quality 80%)
4. **Generate** WebP version (optional, untuk optimasi)
5. **Extract** dimensions (width, height)
6. **Save** ke storage
7. **Create** media record di database

---

### 9. ENDPOINTS SUMMARY

#### 9.1 All Endpoints List

| # | Method | Endpoint | Auth | Description |
|---|---|---|---|---|
| **Authentication** | | | | |
| 1 | POST | `/auth/login` | Public | Login |
| 2 | POST | `/auth/refresh` | Public | Refresh token |
| 3 | POST | `/auth/logout` | Required | Logout |
| 4 | GET | `/auth/me` | Required | Get current user |
| 5 | POST | `/auth/forgot-password` | Public | Request reset |
| 6 | POST | `/auth/reset-password` | Public | Reset password |
| **Blog** | | | | |
| 7 | GET | `/blog` | Public | List posts |
| 8 | GET | `/blog/:slug` | Public | Get post |
| 9 | GET | `/blog/admin` | Admin | List all posts |
| 10 | POST | `/blog` | Admin | Create post |
| 11 | PUT | `/blog/:id` | Admin | Update post |
| 12 | DELETE | `/blog/:id` | Admin | Delete post |
| **Case Studies** | | | | |
| 13 | GET | `/case-studies` | Public | List case studies |
| 14 | GET | `/case-studies/:slug` | Public | Get case study |
| 15 | POST | `/case-studies` | Admin | Create case study |
| 16 | PUT | `/case-studies/:id` | Admin | Update case study |
| 17 | DELETE | `/case-studies/:id` | Admin | Delete case study |
| **Portfolio** | | | | |
| 18 | GET | `/portfolio` | Public | List portfolios |
| 19 | GET | `/portfolio/:slug` | Public | Get portfolio |
| 20 | POST | `/portfolio` | Admin | Create portfolio |
| 21 | PUT | `/portfolio/:id` | Admin | Update portfolio |
| 22 | PUT | `/portfolio/reorder` | Admin | Reorder portfolios |
| 23 | DELETE | `/portfolio/:id` | Admin | Delete portfolio |
| **Services** | | | | |
| 24 | GET | `/services` | Public | List services |
| 25 | GET | `/services/:slug` | Public | Get service |
| 26 | POST | `/services` | Admin | Create service |
| 27 | PUT | `/services/:id` | Admin | Update service |
| 28 | DELETE | `/services/:id` | Admin | Delete service |
| **Leads** | | | | |
| 29 | POST | `/leads` | Public | Submit inquiry |
| 30 | GET | `/leads` | Admin | List leads |
| 31 | GET | `/leads/export` | Admin | Export CSV |
| 32 | GET | `/leads/:id` | Admin | Get lead detail |
| 33 | PUT | `/leads/:id` | Admin | Update lead |
| 34 | DELETE | `/leads/:id` | Admin | Delete lead |
| **Media** | | | | |
| 35 | POST | `/media/upload` | Admin | Upload file |
| 36 | GET | `/media` | Admin | List media |
| 37 | DELETE | `/media/:id` | Admin | Delete media |
| **Categories** | | | | |
| 38 | GET | `/categories` | Public | List categories |
| 39 | POST | `/categories` | Admin | Create category |
| 40 | PUT | `/categories/:id` | Admin | Update category |
| 41 | DELETE | `/categories/:id` | Admin | Delete category |
| **Tags** | | | | |
| 42 | GET | `/tags` | Public | List tags |
| 43 | POST | `/tags` | Admin | Create tag |
| 44 | DELETE | `/tags/:id` | Admin | Delete tag |
| **Dashboard** | | | | |
| 45 | GET | `/dashboard/stats` | Admin | Get stats |
| 46 | GET | `/dashboard/leads-chart` | Admin | Leads chart |
| 47 | GET | `/dashboard/recent-leads` | Admin | Recent leads |
| 48 | GET | `/dashboard/recent-posts` | Admin | Recent posts |
| **Settings** | | | | |
| 49 | GET | `/settings` | Public | Get settings |
| 50 | GET | `/settings/admin` | Admin | Get all settings |
| 51 | PUT | `/settings` | Admin | Update settings |
| **Health** | | | | |
| 52 | GET | `/health` | Public | Health check |

**Total: 52 endpoints**

#### 9.2 Endpoints by Auth Type

| Auth Type | Count |
|---|---|
| Public | 20 endpoints |
| Required (any role) | 10 endpoints |
| Admin only | 22 endpoints |

#### 9.3 Endpoints by Module

| Module | Count |
|---|---|
| Authentication | 6 |
| Blog | 6 |
| Case Studies | 5 |
| Portfolio | 6 |
| Services | 5 |
| Leads | 6 |
| Media | 3 |
| Categories | 4 |
| Tags | 3 |
| Dashboard | 4 |
| Settings | 3 |
| Health | 1 |

---

### 10. POSTMAN / INSOMNIA COLLECTION

Untuk testing, import collection ini ke Postman/Insomnia:

**Environment Variables:**
```json
{
  "base_url": "https://api.webifylab.my.id/api/v1",
  "access_token": "",
  "refresh_token": ""
}
```

**Collection Structure:**
```
WebifyLab API
├── Auth
│   ├── Login
│   ├── Refresh Token
│   ├── Logout
│   ├── Get Me
│   ├── Forgot Password
│   └── Reset Password
├── Blog
│   ├── List Posts
│   ├── Get Post
│   ├── Create Post
│   ├── Update Post
│   └── Delete Post
├── Case Studies
│   ├── List Case Studies
│   ├── Get Case Study
│   ├── Create Case Study
│   ├── Update Case Study
│   └── Delete Case Study
├── Portfolio
│   ├── List Portfolios
│   ├── Get Portfolio
│   ├── Create Portfolio
│   ├── Update Portfolio
│   ├── Reorder Portfolios
│   └── Delete Portfolio
├── Services
│   ├── List Services
│   ├── Get Service
│   ├── Create Service
│   ├── Update Service
│   └── Delete Service
├── Leads
│   ├── Submit Inquiry
│   ├── List Leads
│   ├── Export Leads
│   ├── Get Lead
│   ├── Update Lead
│   └── Delete Lead
├── Media
│   ├── Upload File
│   ├── List Media
│   └── Delete Media
├── Categories
│   ├── List Categories
│   ├── Create Category
│   ├── Update Category
│   └── Delete Category
├── Tags
│   ├── List Tags
│   ├── Create Tag
│   └── Delete Tag
├── Dashboard
│   ├── Get Stats
│   ├── Get Leads Chart
│   ├── Get Recent Leads
│   └── Get Recent Posts
├── Settings
│   ├── Get Settings
│   ├── Get Admin Settings
│   └── Update Settings
└── Health
    └── Health Check
```

---

### 11. FRONTEND INTEGRATION GUIDE

#### 11.1 API Client Setup (Next.js)

```typescript
// frontend/src/lib/api.ts
import axios from 'axios';

const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api/v1',
  headers: {
    'Content-Type': 'application/json',
  },
  withCredentials: true, // For cookies
});

// Request interceptor - add auth token
api.interceptors.request.use((config) => {
  // Token auto-sent via httpOnly cookie
  return config;
});

// Response interceptor - handle errors
api.interceptors.response.use(
  (response) => response,
  async (error) => {
    if (error.response?.status === 401) {
      // Try refresh token
      // If refresh fails, redirect to login
    }
    return Promise.reject(error);
  }
);

export default api;
```

#### 11.2 Example API Calls

**Get blog posts:**
```typescript
// frontend/src/lib/services/blog.ts
import api from '../api';

export const blogService = {
  list: async (params?: { page?: number; limit?: number; category?: string }) => {
    const response = await api.get('/blog', { params });
    return response.data;
  },

  getBySlug: async (slug: string) => {
    const response = await api.get(`/blog/${slug}`);
    return response.data;
  },

  create: async (data: CreateBlogInput) => {
    const response = await api.post('/blog', data);
    return response.data;
  },

  update: async (id: string, data: UpdateBlogInput) => {
    const response = await api.put(`/blog/${id}`, data);
    return response.data;
  },

  delete: async (id: string) => {
    const response = await api.delete(`/blog/${id}`);
    return response.data;
  },
};
```

**Submit lead:**
```typescript
// frontend/src/lib/services/leads.ts
import api from '../api';

export const leadService = {
  submit: async (data: SubmitLeadInput) => {
    const response = await api.post('/leads', data);
    return response.data;
  },
};
```

---

### 12. BACKEND IMPLEMENTATION GUIDE

#### 12.1 Route Registration (Gin)

```go
// backend/internal/routes/routes.go
package routes

import (
    "github.com/gin-gonic/gin"
    "webifylab/internal/handlers"
    "webifylab/internal/middleware"
)

func SetupRoutes(r *gin.Engine, h *handlers.Handlers) {
    api := r.Group("/api/v1")
    
    // Public routes
    api.GET("/health", h.Health.Check)
    api.GET("/settings", h.Settings.GetPublic)
    api.GET("/categories", h.Categories.List)
    api.GET("/tags", h.Tags.List)
    
    // Public content
    api.GET("/blog", h.Blog.List)
    api.GET("/blog/:slug", h.Blog.GetBySlug)
    api.GET("/case-studies", h.CaseStudy.List)
    api.GET("/case-studies/:slug", h.CaseStudy.GetBySlug)
    api.GET("/portfolio", h.Portfolio.List)
    api.GET("/portfolio/:slug", h.Portfolio.GetBySlug)
    api.GET("/services", h.Service.List)
    api.GET("/services/:slug", h.Service.GetBySlug)
    
    // Public leads
    api.POST("/leads", h.Lead.Submit)
    
    // Auth
    auth := api.Group("/auth")
    {
        auth.POST("/login", h.Auth.Login)
        auth.POST("/refresh", h.Auth.Refresh)
        auth.POST("/forgot-password", h.Auth.ForgotPassword)
        auth.POST("/reset-password", h.Auth.ResetPassword)
    }
    
    // Protected routes
    protected := api.Group("")
    protected.Use(middleware.AuthRequired())
    {
        protected.POST("/auth/logout", h.Auth.Logout)
        protected.GET("/auth/me", h.Auth.Me)
        
        // Admin: Blog
        protected.GET("/blog/admin", h.Blog.ListAdmin)
        protected.POST("/blog", h.Blog.Create)
        protected.PUT("/blog/:id", h.Blog.Update)
        protected.DELETE("/blog/:id", h.Blog.Delete)
        
        // Admin: Case Studies
        protected.POST("/case-studies", h.CaseStudy.Create)
        protected.PUT("/case-studies/:id", h.CaseStudy.Update)
        protected.DELETE("/case-studies/:id", h.CaseStudy.Delete)
        
        // Admin: Portfolio
        protected.POST("/portfolio", h.Portfolio.Create)
        protected.PUT("/portfolio/:id", h.Portfolio.Update)
        protected.PUT("/portfolio/reorder", h.Portfolio.Reorder)
        protected.DELETE("/portfolio/:id", h.Portfolio.Delete)
        
        // Admin: Services
        protected.POST("/services", h.Service.Create)
        protected.PUT("/services/:id", h.Service.Update)
        protected.DELETE("/services/:id", h.Service.Delete)
        
        // Admin: Leads
        protected.GET("/leads", h.Lead.List)
        protected.GET("/leads/export", h.Lead.Export)
        protected.GET("/leads/:id", h.Lead.Get)
        protected.PUT("/leads/:id", h.Lead.Update)
        protected.DELETE("/leads/:id", h.Lead.Delete)
        
        // Admin: Media
        protected.POST("/media/upload", h.Media.Upload)
        protected.GET("/media", h.Media.List)
        protected.DELETE("/media/:id", h.Media.Delete)
        
        // Admin: Categories
        protected.POST("/categories", h.Categories.Create)
        protected.PUT("/categories/:id", h.Categories.Update)
        protected.DELETE("/categories/:id", h.Categories.Delete)
        
        // Admin: Tags
        protected.POST("/tags", h.Tags.Create)
        protected.DELETE("/tags/:id", h.Tags.Delete)
        
        // Admin: Dashboard
        protected.GET("/dashboard/stats", h.Dashboard.Stats)
        protected.GET("/dashboard/leads-chart", h.Dashboard.LeadsChart)
        protected.GET("/dashboard/recent-leads", h.Dashboard.RecentLeads)
        protected.GET("/dashboard/recent-posts", h.Dashboard.RecentPosts)
        
        // Admin: Settings
        protected.GET("/settings/admin", h.Settings.GetAdmin)
        protected.PUT("/settings", h.Settings.Update)
    }
}
```

#### 12.2 Standard Response Helper

```go
// backend/pkg/response/response.go
package response

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Response struct {
    Success    bool        `json:"success"`
    Message    string      `json:"message"`
    Data       interface{} `json:"data,omitempty"`
    Errors     []Error     `json:"errors,omitempty"`
    Pagination *Pagination `json:"pagination,omitempty"`
}

type Error struct {
    Field   string `json:"field,omitempty"`
    Message string `json:"message"`
}

type Pagination struct {
    Page       int  `json:"page"`
    Limit      int  `json:"limit"`
    Total      int64 `json:"total"`
    TotalPages int   `json:"total_pages"`
    HasNext    bool  `json:"has_next"`
    HasPrev    bool  `json:"has_prev"`
}

func Success(c *gin.Context, message string, data interface{}) {
    c.JSON(http.StatusOK, Response{
        Success: true,
        Message: message,
        Data:    data,
    })
}

func Created(c *gin.Context, message string, data interface{}) {
    c.JSON(http.StatusCreated, Response{
        Success: true,
        Message: message,
        Data:    data,
    })
}

func Error(c *gin.Context, status int, message string, errors []Error) {
    c.JSON(status, Response{
        Success: false,
        Message: message,
        Errors:  errors,
    })
}

func Paginated(c *gin.Context, message string, data interface{}, pagination Pagination) {
    c.JSON(http.StatusOK, Response{
        Success:    true,
        Message:    message,
        Data:       data,
        Pagination: &pagination,
    })
}
```

---

### 13. TESTING CHECKLIST

#### 13.1 Manual Testing Checklist

**Authentication:**
- [ ] Login dengan valid credentials → 200 + tokens
- [ ] Login dengan invalid credentials → 401
- [ ] Login 6x dalam 1 menit → 429 (rate limit)
- [ ] Refresh token dengan valid refresh token → 200 + new access token
- [ ] Refresh token dengan expired token → 401
- [ ] Access protected endpoint tanpa token → 401
- [ ] Access protected endpoint dengan expired token → 401
- [ ] Forgot password → 200 (selalu success)
- [ ] Reset password dengan valid token → 200
- [ ] Reset password dengan expired token → 422

**Blog:**
- [ ] GET /blog → list published posts only
- [ ] GET /blog?category=tutorial → filter works
- [ ] GET /blog?tag=golang → filter works
- [ ] GET /blog/:slug → full post data
- [ ] GET /blog/non-existent-slug → 404
- [ ] POST /blog (admin) → 201
- [ ] POST /blog tanpa title → 422
- [ ] POST /blog dengan duplicate slug → 409
- [ ] PUT /blog/:id → 200
- [ ] DELETE /blog/:id → 200 (soft delete)
- [ ] GET /blog/non-existent-slug (setelah delete) → 404

**Leads:**
- [ ] POST /leads dengan valid data → 201 + email sent
- [ ] POST /leads dengan invalid email → 422
- [ ] POST /leads 6x dalam 1 jam → 429
- [ ] GET /leads (admin) → list all leads
- [ ] PUT /leads/:id update status → 200
- [ ] GET /leads/export → CSV download

**Media:**
- [ ] POST /media/upload dengan JPG < 5MB → 201
- [ ] POST /media/upload dengan file > 5MB → 413
- [ ] POST /media/upload dengan PDF → 415
- [ ] GET /media (admin) → list uploaded files
- [ ] DELETE /media/:id → 200 + file deleted

**Performance:**
- [ ] GET /blog response time < 500ms
- [ ] GET /case-studies response time < 500ms
- [ ] POST /leads response time < 1s (termasuk email)
- [ ] POST /media/upload response time < 3s

---

### 14. CHANGELOG

| Version | Date | Changes |
|---|---|---|
| 1.0.0 | 2026-09-05 | Initial release - 52 endpoints |