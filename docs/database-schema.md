# 📄 DATABASE SCHEMA DOCUMENT
## WebifyLab Portal — Phase 1 (MVP)

---

### 1. DOCUMENT INFORMATION

| Field | Detail |
|---|---|
| **Product** | WebifyLab Portal |
| **Document Type** | Database Schema Design |
| **Version** | 1.0.0 |
| **Date** | 05 September 2026 |
| **Database** | PostgreSQL 15 |
| **ORM** | GORM (Golang) |
| **Author** | Rizal (System Analyst & Fullstack Developer) |
| **Status** | Draft — Ready for Implementation |

---

### 2. DESIGN PRINCIPLES

#### 2.1 Core Principles

| Principle | Implementation | Reason |
|---|---|---|
| **UUID Primary Keys** | `UUID` type untuk semua PK | Secure, scalable, SSO-ready, no sequential guessing |
| **Timezone-Aware Timestamps** | `TIMESTAMPTZ` untuk semua timestamp | Consistent across timezones, avoid DST issues |
| **Soft Delete** | `deleted_at` column untuk critical tables | Data recovery, audit trail, prevent accidental loss |
| **Audit Trail** | `created_at`, `updated_at` di semua tables | Track changes, debugging, analytics |
| **Slug-Based URLs** | Unique `slug` column untuk public content | SEO-friendly URLs, human-readable |
| **JSON for Flexible Data** | `JSONB` untuk arrays/complex structures | Flexible schema, no extra tables needed |
| **Explicit Table Names** | Snake_case, plural (e.g., `blog_posts`) | GORM convention, consistency |
| **Foreign Key Constraints** | `ON DELETE CASCADE` atau `SET NULL` | Data integrity, prevent orphaned records |
| **Indexing Strategy** | Index on FK, status, slug, created_at | Query performance, fast filtering/sorting |

#### 2.2 Naming Conventions

| Element | Convention | Example |
|---|---|---|
| Table names | snake_case, plural | `blog_posts`, `case_studies` |
| Column names | snake_case | `created_at`, `client_name` |
| Primary key | `id` | `id UUID PRIMARY KEY` |
| Foreign key | `{table_singular}_id` | `author_id`, `category_id` |
| Index names | `idx_{table}_{column}` | `idx_blog_posts_slug` |
| Unique constraint | `uq_{table}_{column}` | `uq_users_email` |
| Foreign key constraint | `fk_{table}_{column}` | `fk_blog_posts_author_id` |

#### 2.3 Common Columns

Semua tabel memiliki kolom standar berikut:

```sql
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
deleted_at TIMESTAMPTZ NULL  -- Soft delete (optional per table)
```

---

### 3. COMPLETE ERD (Entity Relationship Diagram)

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                              USERS & AUTH                                   │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ┌──────────────────┐         ┌──────────────────────┐                     │
│  │     users        │         │  refresh_tokens      │                     │
│  ├──────────────────┤         ├──────────────────────┤                     │
│  │ id (PK, UUID)    │◄────────│ user_id (FK)         │                     │
│  │ name             │    1:N  │ token (unique)       │                     │
│  │ email (unique)   │         │ expires_at           │
│  │ password_hash    │         │ created_at           │
│  │ role (enum)      │         └──────────────────────┘                     │
│  │ created_at       │                                                      │
│  │ updated_at       │         ┌──────────────────────┐                     │
│  └────────┬─────────┘         │ password_resets      │                     │
│           │                   ├──────────────────────┤                     │
│           │ 1:N               │ id (PK, UUID)        │                     │
│           │                   │ user_id (FK)         │                     │
│           ▼                   │ token (unique)       │                     │
│  ┌──────────────────┐         │ expires_at           │                     │
│  │   blog_posts     │         │ used_at              │                     │
│  │   (author_id FK) │         │ created_at           │                     │
│  └──────────────────┘         └──────────────────────┘                     │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────────────┐
│                           CONTENT MANAGEMENT                                │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ┌──────────────────┐         ┌──────────────────┐                         │
│  │  categories      │         │     tags         │                         │
│  ├──────────────────┤         ├──────────────────┤                         │
│  │ id (PK, UUID)    │         │ id (PK, UUID)    │                         │
│  │ name             │         │ name             │                         │
│  │ slug (unique)    │         │ slug (unique)    │                         │
│  │ type (enum)      │         │ created_at       │                         │
│  │ created_at       │         └────────┬─────────┘                         │
│  └────────┬─────────┘                 │                                   │
│           │ 1:N                       │ M:N                               │
│           ▼                           ▼                                   │
│  ┌──────────────────┐    ┌──────────────────────┐  ┌──────────────────┐   │
│  │  blog_posts      │    │    post_tags         │  │ case_study_tags  │   │
│  ├──────────────────┤    ├──────────────────────┤  ├──────────────────┤   │
│  │ id (PK, UUID)    │◄───│ post_id (FK)         │  │ case_study_id(FK)│   │
│  │ title            │    │ tag_id (FK)          │  │ tag_id (FK)      │   │
│  │ slug (unique)    │    └──────────────────────┘  └──────────────────┘   │
│  │ content (text)   │                                                      │
│  │ excerpt          │                                                      │
│  │ cover_image      │                                                      │
│  │ author_id (FK)   │                                                      │
│  │ category_id (FK) │                                                      │
│  │ status (enum)    │                                                      │
│  │ meta_title       │                                                      │
│  │ meta_description │                                                      │
│  │ published_at     │                                                      │
│  │ read_time        │                                                      │
│  │ created_at       │                                                      │
│  │ updated_at       │                                                      │
│  └──────────────────┘                                                      │
│                                                                             │
│  ┌──────────────────────────────────────────────────────────────────────┐  │
│  │                         case_studies                                 │  │
│  ├──────────────────────────────────────────────────────────────────────┤  │
│  │ id (PK, UUID)                                                        │  │
│  │ title                                                                │  │
│  │ slug (unique)                                                        │  │
│  │ client_name                                                          │  │
│  │ industry                                                             │  │
│  │ challenge (text)                                                     │  │
│  │ approach (text)                                                      │  │
│  │ solution (text)                                                      │  │
│  │ ai_type (enum array)                                                 │  │
│  │ results (text)                                                       │  │
│  │ lessons (text)                                                       │  │
│  │ testimonial (text)                                                   │  │
│  │ metrics (JSONB)                                                      │  │
│  │ cover_image                                                          │  │
│  │ images (JSONB array)                                                 │  │
│  │ status (enum)                                                        │  │
│  │ meta_title                                                           │  │
│  │ meta_description                                                     │  │
│  │ published_at                                                         │  │
│  │ created_at                                                           │  │
│  │ updated_at                                                           │  │
│  └──────────────────────────────────────────────────────────────────────┘  │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────────────┐
│                         PORTFOLIO & SERVICES                                │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ┌──────────────────────────────────────────────────────────────────────┐  │
│  │                          portfolios                                  │  │
│  ├──────────────────────────────────────────────────────────────────────┤  │
│  │ id (PK, UUID)                                                        │  │
│  │ title                                                                │  │
│  │ slug (unique)                                                        │  │
│  │ description (text)                                                   │  │
│  │ category                                                             │  │
│  │ tech_stack (JSONB array)                                             │  │
│  │ live_url                                                             │  │
│  │ thumbnail                                                            │  │
│  │ images (JSONB array)                                                 │  │
│  │ lessons (text)                                                       │  │
│  │ status (enum)                                                        │  │
│  │ order_index                                                          │  │
│  │ meta_title                                                           │  │
│  │ meta_description                                                     │  │
│  │ published_at                                                         │  │
│  │ created_at                                                           │  │
│  │ updated_at                                                           │  │
│  └──────────────────────────────────────────────────────────────────────┘  │
│                                                                             │
│  ┌──────────────────────────────────────────────────────────────────────┐  │
│  │                           services                                   │  │
│  ├──────────────────────────────────────────────────────────────────────┤  │
│  │ id (PK, UUID)                                                        │  │
│  │ name                                                                 │  │
│  │ slug (unique)                                                        │  │
│  │ description (text)                                                   │  │
│  │ icon                                                                 │  │
│  │ use_cases (JSONB array)                                              │  │
│  │ timeline_estimate                                                    │  │
│  │ tech_stack (JSONB array)                                             │  │
│  │ learning_points (text)                                               │  │
│  │ is_active                                                            │  │
│  │ order_index                                                          │  │
│  │ created_at                                                           │  │
│  │ updated_at                                                           │  │
│  └──────────────────────────────────────────────────────────────────────┘  │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────────────┐
│                         LEADS & MEDIA                                       │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ┌──────────────────────────────────────────────────────────────────────┐  │
│  │                             leads                                    │  │
│  ├──────────────────────────────────────────────────────────────────────┤  │
│  │ id (PK, UUID)                                                        │  │
│  │ name                                                                 │  │
│  │ email                                                                │  │
│  │ phone (WhatsApp)                                                     │  │
│  │ service_type                                                         │  │
│  │ budget_range                                                         │  │
│  │ description (text)                                                   │  │
│  │ status (enum)                                                        │  │
│  │ notes (text)                                                         │  │
│  │ ip_address                                                           │  │
│  │ user_agent                                                           │  │
│  │ created_at                                                           │  │
│  │ updated_at                                                           │  │
│  └──────────────────────────────────────────────────────────────────────┘  │
│                                                                             │
│  ┌──────────────────────────────────────────────────────────────────────┐  │
│  │                            media                                     │  │
│  ├──────────────────────────────────────────────────────────────────────┤  │
│  │ id (PK, UUID)                                                        │  │
│  │ filename                                                             │  │
│  │ original_name                                                        │  │
│  │ url                                                                  │  │
│  │ mime_type                                                            │  │
│  │ size (bytes)                                                         │  │
│  │ width (nullable)                                                     │  │
│  │ height (nullable)                                                    │  │
│  │ alt_text                                                             │  │
│  │ uploaded_by (FK → users)                                             │  │
│  │ created_at                                                           │  │
│  └──────────────────────────────────────────────────────────────────────┘  │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────────────┐
│                         SETTINGS                                            │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ┌──────────────────────────────────────────────────────────────────────┐  │
│  │                         site_settings                                │  │
│  ├──────────────────────────────────────────────────────────────────────┤  │
│  │ id (PK, UUID)                                                        │  │
│  │ key (unique)                                                         │  │
│  │ value (JSONB)                                                        │  │
│  │ description                                                          │  │
│  │ updated_at                                                           │  │
│  └──────────────────────────────────────────────────────────────────────┘  │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

### 4. TABLE DEFINITIONS (Detail)

#### 4.1 users

**Purpose:** Admin users untuk CMS dashboard

| Column | Type | Constraints | Default | Description |
|---|---|---|---|---|
| `id` | UUID | PRIMARY KEY | `gen_random_uuid()` | Unique identifier |
| `name` | VARCHAR(100) | NOT NULL | - | Full name |
| `email` | VARCHAR(255) | NOT NULL, UNIQUE | - | Login email |
| `password_hash` | VARCHAR(255) | NOT NULL | - | Bcrypt hashed password |
| `role` | VARCHAR(20) | NOT NULL | `'editor'` | User role (enum) |
| `created_at` | TIMESTAMPTZ | NOT NULL | `NOW()` | Record creation time |
| `updated_at` | TIMESTAMPTZ | NOT NULL | `NOW()` | Last update time |

**Indexes:**
- `uq_users_email` UNIQUE on `email`
- `idx_users_role` on `role`

**Constraints:**
- `role` must be one of: `super_admin`, `admin`, `editor`

---

#### 4.2 refresh_tokens

**Purpose:** JWT refresh tokens untuk authentication

| Column | Type | Constraints | Default | Description |
|---|---|---|---|---|
| `id` | UUID | PRIMARY KEY | `gen_random_uuid()` | Unique identifier |
| `user_id` | UUID | NOT NULL, FK → users(id) | - | Token owner |
| `token` | VARCHAR(500) | NOT NULL, UNIQUE | - | Refresh token string |
| `expires_at` | TIMESTAMPTZ | NOT NULL | - | Token expiration |
| `created_at` | TIMESTAMPTZ | NOT NULL | `NOW()` | Token creation time |

**Indexes:**
- `uq_refresh_tokens_token` UNIQUE on `token`
- `idx_refresh_tokens_user_id` on `user_id`
- `idx_refresh_tokens_expires_at` on `expires_at`

**Foreign Keys:**
- `fk_refresh_tokens_user_id` → `users(id)` ON DELETE CASCADE

---

#### 4.3 password_resets

**Purpose:** Token untuk forgot password flow

| Column | Type | Constraints | Default | Description |
|---|---|---|---|---|
| `id` | UUID | PRIMARY KEY | `gen_random_uuid()` | Unique identifier |
| `user_id` | UUID | NOT NULL, FK → users(id) | - | User requesting reset |
| `token` | VARCHAR(255) | NOT NULL, UNIQUE | - | Reset token |
| `expires_at` | TIMESTAMPTZ | NOT NULL | - | Token expiration (1 hour) |
| `used_at` | TIMESTAMPTZ | NULL | - | When token was used |
| `created_at` | TIMESTAMPTZ | NOT NULL | `NOW()` | Request time |

**Indexes:**
- `uq_password_resets_token` UNIQUE on `token`
- `idx_password_resets_user_id` on `user_id`

**Foreign Keys:**
- `fk_password_resets_user_id` → `users(id)` ON DELETE CASCADE

---

#### 4.4 categories

**Purpose:** Categories untuk blog posts dan portfolios

| Column | Type | Constraints | Default | Description |
|---|---|---|---|---|
| `id` | UUID | PRIMARY KEY | `gen_random_uuid()` | Unique identifier |
| `name` | VARCHAR(100) | NOT NULL | - | Category name |
| `slug` | VARCHAR(100) | NOT NULL, UNIQUE | - | URL-friendly slug |
| `type` | VARCHAR(20) | NOT NULL | - | Category type (enum) |
| `created_at` | TIMESTAMPTZ | NOT NULL | `NOW()` | Record creation time |

**Indexes:**
- `uq_categories_slug` UNIQUE on `slug`
- `idx_categories_type` on `type`

**Constraints:**
- `type` must be one of: `blog`, `portfolio`

---

#### 4.5 tags

**Purpose:** Tags untuk blog posts dan case studies

| Column | Type | Constraints | Default | Description |
|---|---|---|---|---|
| `id` | UUID | PRIMARY KEY | `gen_random_uuid()` | Unique identifier |
| `name` | VARCHAR(50) | NOT NULL | - | Tag name |
| `slug` | VARCHAR(50) | NOT NULL, UNIQUE | - | URL-friendly slug |
| `created_at` | TIMESTAMPTZ | NOT NULL | `NOW()` | Record creation time |

**Indexes:**
- `uq_tags_slug` UNIQUE on `slug`
- `uq_tags_name` UNIQUE on `name`

---

#### 4.6 blog_posts

**Purpose:** Blog articles untuk content marketing

| Column | Type | Constraints | Default | Description |
|---|---|---|---|---|
| `id` | UUID | PRIMARY KEY | `gen_random_uuid()` | Unique identifier |
| `title` | VARCHAR(255) | NOT NULL | - | Article title |
| `slug` | VARCHAR(255) | NOT NULL, UNIQUE | - | URL-friendly slug |
| `content` | TEXT | NOT NULL | - | Rich text content (HTML/JSON) |
| `excerpt` | TEXT | NULL | - | Short summary (max 300 chars) |
| `cover_image` | VARCHAR(500) | NULL | - | Cover image URL |
| `author_id` | UUID | NOT NULL, FK → users(id) | - | Article author |
| `category_id` | UUID | NULL, FK → categories(id) | - | Article category |
| `status` | VARCHAR(20) | NOT NULL | `'draft'` | Publication status (enum) |
| `meta_title` | VARCHAR(255) | NULL | - | SEO meta title |
| `meta_description` | VARCHAR(500) | NULL | - | SEO meta description |
| `published_at` | TIMESTAMPTZ | NULL | - | When published |
| `read_time` | INTEGER | NULL | - | Estimated read time (minutes) |
| `created_at` | TIMESTAMPTZ | NOT NULL | `NOW()` | Record creation time |
| `updated_at` | TIMESTAMPTZ | NOT NULL | `NOW()` | Last update time |
| `deleted_at` | TIMESTAMPTZ | NULL | - | Soft delete timestamp |

**Indexes:**
- `uq_blog_posts_slug` UNIQUE on `slug`
- `idx_blog_posts_author_id` on `author_id`
- `idx_blog_posts_category_id` on `category_id`
- `idx_blog_posts_status` on `status`
- `idx_blog_posts_published_at` on `published_at`
- `idx_blog_posts_created_at` on `created_at`
- `idx_blog_posts_deleted_at` on `deleted_at`

**Foreign Keys:**
- `fk_blog_posts_author_id` → `users(id)` ON DELETE RESTRICT
- `fk_blog_posts_category_id` → `categories(id)` ON DELETE SET NULL

**Constraints:**
- `status` must be one of: `draft`, `published`, `archived`
- `published_at` must be set when `status` = `published`

---

#### 4.7 post_tags

**Purpose:** Many-to-many relationship between blog posts and tags

| Column | Type | Constraints | Default | Description |
|---|---|---|---|---|
| `post_id` | UUID | NOT NULL, FK → blog_posts(id) | - | Blog post ID |
| `tag_id` | UUID | NOT NULL, FK → tags(id) | - | Tag ID |

**Primary Key:** Composite key (`post_id`, `tag_id`)

**Indexes:**
- `idx_post_tags_post_id` on `post_id`
- `idx_post_tags_tag_id` on `tag_id`

**Foreign Keys:**
- `fk_post_tags_post_id` → `blog_posts(id)` ON DELETE CASCADE
- `fk_post_tags_tag_id` → `tags(id)` ON DELETE CASCADE

---

#### 4.8 case_studies

**Purpose:** Detailed case studies dengan struktur STAR

| Column | Type | Constraints | Default | Description |
|---|---|---|---|---|
| `id` | UUID | PRIMARY KEY | `gen_random_uuid()` | Unique identifier |
| `title` | VARCHAR(255) | NOT NULL | - | Case study title |
| `slug` | VARCHAR(255) | NOT NULL, UNIQUE | - | URL-friendly slug |
| `client_name` | VARCHAR(100) | NOT NULL | - | Client/company name |
| `industry` | VARCHAR(100) | NOT NULL | - | Client industry |
| `challenge` | TEXT | NOT NULL | - | Problem/challenge faced |
| `approach` | TEXT | NOT NULL | - | Approach/methodology |
| `solution` | TEXT | NOT NULL | - | Solution implemented |
| `ai_type` | JSONB | NULL | - | Array of AI types used |
| `results` | TEXT | NOT NULL | - | Results achieved |
| `lessons` | TEXT | NULL | - | Lessons learned |
| `testimonial` | TEXT | NULL | - | Client testimonial |
| `metrics` | JSONB | NULL | - | Key metrics (JSON) |
| `cover_image` | VARCHAR(500) | NULL | - | Cover image URL |
| `images` | JSONB | NULL | - | Array of image URLs |
| `status` | VARCHAR(20) | NOT NULL | `'draft'` | Publication status |
| `meta_title` | VARCHAR(255) | NULL | - | SEO meta title |
| `meta_description` | VARCHAR(500) | NULL | - | SEO meta description |
| `published_at` | TIMESTAMPTZ | NULL | - | When published |
| `created_at` | TIMESTAMPTZ | NOT NULL | `NOW()` | Record creation time |
| `updated_at` | TIMESTAMPTZ | NOT NULL | `NOW()` | Last update time |
| `deleted_at` | TIMESTAMPTZ | NULL | - | Soft delete timestamp |

**Indexes:**
- `uq_case_studies_slug` UNIQUE on `slug`
- `idx_case_studies_status` on `status`
- `idx_case_studies_industry` on `industry`
- `idx_case_studies_published_at` on `published_at`
- `idx_case_studies_deleted_at` on `deleted_at`
- GIN index on `ai_type` for JSONB queries

**Constraints:**
- `status` must be one of: `draft`, `published`, `archived`
- `ai_type` JSONB array contains values: `spk`, `sistem_pakar`, `data_analytics`, `machine_learning`

**Example `metrics` JSONB:**
```json
{
  "performance_improvement": "40%",
  "cost_reduction": "25%",
  "time_saved": "10 hours/week"
}
```

---

#### 4.9 case_study_tags

**Purpose:** Many-to-many relationship between case studies and tags

| Column | Type | Constraints | Default | Description |
|---|---|---|---|---|
| `case_study_id` | UUID | NOT NULL, FK → case_studies(id) | - | Case study ID |
| `tag_id` | UUID | NOT NULL, FK → tags(id) | - | Tag ID |

**Primary Key:** Composite key (`case_study_id`, `tag_id`)

**Indexes:**
- `idx_case_study_tags_case_study_id` on `case_study_id`
- `idx_case_study_tags_tag_id` on `tag_id`

**Foreign Keys:**
- `fk_case_study_tags_case_study_id` → `case_studies(id)` ON DELETE CASCADE
- `fk_case_study_tags_tag_id` → `tags(id)` ON DELETE CASCADE

---

#### 4.10 portfolios

**Purpose:** Portfolio items untuk showcase projects

| Column | Type | Constraints | Default | Description |
|---|---|---|---|---|
| `id` | UUID | PRIMARY KEY | `gen_random_uuid()` | Unique identifier |
| `title` | VARCHAR(255) | NOT NULL | - | Project title |
| `slug` | VARCHAR(255) | NOT NULL, UNIQUE | - | URL-friendly slug |
| `description` | TEXT | NOT NULL | - | Project description |
| `category` | VARCHAR(50) | NOT NULL | - | Project category |
| `tech_stack` | JSONB | NULL | - | Array of technologies used |
| `live_url` | VARCHAR(500) | NULL | - | Live project URL |
| `thumbnail` | VARCHAR(500) | NULL | - | Thumbnail image URL |
| `images` | JSONB | NULL | - | Array of image URLs |
| `lessons` | TEXT | NULL | - | Lessons learned |
| `status` | VARCHAR(20) | NOT NULL | `'draft'` | Publication status |
| `order_index` | INTEGER | NOT NULL | `0` | Display order |
| `meta_title` | VARCHAR(255) | NULL | - | SEO meta title |
| `meta_description` | VARCHAR(500) | NULL | - | SEO meta description |
| `published_at` | TIMESTAMPTZ | NULL | - | When published |
| `created_at` | TIMESTAMPTZ | NOT NULL | `NOW()` | Record creation time |
| `updated_at` | TIMESTAMPTZ | NOT NULL | `NOW()` | Last update time |
| `deleted_at` | TIMESTAMPTZ | NULL | - | Soft delete timestamp |

**Indexes:**
- `uq_portfolios_slug` UNIQUE on `slug`
- `idx_portfolios_category` on `category`
- `idx_portfolios_status` on `status`
- `idx_portfolios_order_index` on `order_index`
- `idx_portfolios_published_at` on `published_at`
- `idx_portfolios_deleted_at` on `deleted_at`

**Constraints:**
- `status` must be one of: `draft`, `published`, `archived`
- `category` must be one of: `web`, `mobile`, `ai`, `ecommerce`, `other`

**Example `tech_stack` JSONB:**
```json
["Next.js", "Golang", "PostgreSQL", "Docker", "Tailwind CSS"]
```

---

#### 4.11 services

**Purpose:** Services yang ditawarkan WebifyLab

| Column | Type | Constraints | Default | Description |
|---|---|---|---|---|
| `id` | UUID | PRIMARY KEY | `gen_random_uuid()` | Unique identifier |
| `name` | VARCHAR(100) | NOT NULL | - | Service name |
| `slug` | VARCHAR(100) | NOT NULL, UNIQUE | - | URL-friendly slug |
| `description` | TEXT | NOT NULL | - | Service description |
| `icon` | VARCHAR(50) | NULL | - | Icon name (Lucide) |
| `use_cases` | JSONB | NULL | - | Array of use cases |
| `timeline_estimate` | VARCHAR(100) | NULL | - | Estimated timeline |
| `tech_stack` | JSONB | NULL | - | Array of technologies |
| `learning_points` | TEXT | NULL | - | What we'll learn together |
| `is_active` | BOOLEAN | NOT NULL | `true` | Whether service is active |
| `order_index` | INTEGER | NOT NULL | `0` | Display order |
| `created_at` | TIMESTAMPTZ | NOT NULL | `NOW()` | Record creation time |
| `updated_at` | TIMESTAMPTZ | NOT NULL | `NOW()` | Last update time |

**Indexes:**
- `uq_services_slug` UNIQUE on `slug`
- `idx_services_is_active` on `is_active`
- `idx_services_order_index` on `order_index`

**Example `use_cases` JSONB:**
```json
[
  "Website company profile untuk UMKM",
  "Landing page untuk produk digital",
  "Web app untuk internal tools"
]
```

---

#### 4.12 leads

**Purpose:** Inquiry dari calon klien (mini CRM)

| Column | Type | Constraints | Default | Description |
|---|---|---|---|---|
| `id` | UUID | PRIMARY KEY | `gen_random_uuid()` | Unique identifier |
| `name` | VARCHAR(100) | NOT NULL | - | Contact name |
| `email` | VARCHAR(255) | NOT NULL | - | Contact email |
| `phone` | VARCHAR(20) | NOT NULL | - | WhatsApp number |
| `service_type` | VARCHAR(50) | NOT NULL | - | Service interested in |
| `budget_range` | VARCHAR(50) | NOT NULL | - | Budget range |
| `description` | TEXT | NOT NULL | - | Project description |
| `status` | VARCHAR(30) | NOT NULL | `'new'` | Lead status (enum) |
| `notes` | TEXT | NULL | - | Internal notes |
| `ip_address` | VARCHAR(45) | NULL | - | Submitter IP address |
| `user_agent` | TEXT | NULL | - | Submitter user agent |
| `created_at` | TIMESTAMPTZ | NOT NULL | `NOW()` | Inquiry time |
| `updated_at` | TIMESTAMPTZ | NOT NULL | `NOW()` | Last update time |

**Indexes:**
- `idx_leads_status` on `status`
- `idx_leads_email` on `email`
- `idx_leads_created_at` on `created_at`
- `idx_leads_service_type` on `service_type`

**Constraints:**
- `status` must be one of: `new`, `contacted`, `negotiation`, `proposal_sent`, `won`, `lost`
- `service_type` must be one of: `web_development`, `mobile_app`, `uiux_design`, `ai_integration`, `other`
- `budget_range` must be one of: `<5jt`, `5-15jt`, `15-50jt`, `50-100jt`, `>100jt`

---

#### 4.13 media

**Purpose:** Uploaded files (images, documents)

| Column | Type | Constraints | Default | Description |
|---|---|---|---|---|
| `id` | UUID | PRIMARY KEY | `gen_random_uuid()` | Unique identifier |
| `filename` | VARCHAR(255) | NOT NULL | - | Stored filename (unique) |
| `original_name` | VARCHAR(255) | NOT NULL | - | Original filename |
| `url` | VARCHAR(500) | NOT NULL | - | Public URL |
| `mime_type` | VARCHAR(100) | NOT NULL | - | MIME type |
| `size` | BIGINT | NOT NULL | - | File size in bytes |
| `width` | INTEGER | NULL | - | Image width (if image) |
| `height` | INTEGER | NULL | - | Image height (if image) |
| `alt_text` | VARCHAR(255) | NULL | - | Alt text for accessibility |
| `uploaded_by` | UUID | NOT NULL, FK → users(id) | - | Uploader user |
| `created_at` | TIMESTAMPTZ | NOT NULL | `NOW()` | Upload time |

**Indexes:**
- `uq_media_filename` UNIQUE on `filename`
- `idx_media_uploaded_by` on `uploaded_by`
- `idx_media_mime_type` on `mime_type`
- `idx_media_created_at` on `created_at`

**Foreign Keys:**
- `fk_media_uploaded_by` → `users(id)` ON DELETE RESTRICT

**Constraints:**
- `mime_type` must be one of: `image/jpeg`, `image/png`, `image/webp`
- `size` must be ≤ 5MB (5242880 bytes)

---

#### 4.14 site_settings

**Purpose:** Key-value store untuk site configuration

| Column | Type | Constraints | Default | Description |
|---|---|---|---|---|
| `id` | UUID | PRIMARY KEY | `gen_random_uuid()` | Unique identifier |
| `key` | VARCHAR(100) | NOT NULL, UNIQUE | - | Setting key |
| `value` | JSONB | NOT NULL | - | Setting value (JSON) |
| `description` | TEXT | NULL | - | Setting description |
| `updated_at` | TIMESTAMPTZ | NOT NULL | `NOW()` | Last update time |

**Indexes:**
- `uq_site_settings_key` UNIQUE on `key`

**Example Settings:**

| Key | Value | Description |
|---|---|---|
| `site_name` | `"WebifyLab"` | Site name |
| `tagline` | `"We Experiment, You Grow"` | Site tagline |
| `contact_email` | `"hello@webifylab.my.id"` | Contact email |
| `contact_whatsapp` | `"6281234567890"` | WhatsApp number |
| `social_media` | `{"instagram": "...", "linkedin": "..."}` | Social media links |
| `seo_default_title` | `"WebifyLab - Digital Learning Lab"` | Default SEO title |
| `seo_default_description` | `"..."` | Default SEO description |
| `seo_og_image` | `"/images/og-default.jpg"` | Default OG image |

---

### 5. ENUMS & CONSTANTS

#### 5.1 User Roles
```go
const (
    RoleSuperAdmin = "super_admin"
    RoleAdmin      = "admin"
    RoleEditor     = "editor"
)
```

#### 5.2 Content Status
```go
const (
    StatusDraft     = "draft"
    StatusPublished = "published"
    StatusArchived  = "archived"
)
```

#### 5.3 Lead Status
```go
const (
    LeadStatusNew           = "new"
    LeadStatusContacted     = "contacted"
    LeadStatusNegotiation   = "negotiation"
    LeadStatusProposalSent  = "proposal_sent"
    LeadStatusWon           = "won"
    LeadStatusLost          = "lost"
)
```

#### 5.4 Service Types (Lead)
```go
const (
    ServiceWebDevelopment  = "web_development"
    ServiceMobileApp       = "mobile_app"
    ServiceUIUXDesign      = "uiux_design"
    ServiceAIIntegration   = "ai_integration"
    ServiceOther           = "other"
)
```

#### 5.5 Budget Ranges
```go
const (
    BudgetUnder5M     = "<5jt"
    Budget5to15M      = "5-15jt"
    Budget15to50M     = "15-50jt"
    Budget50to100M    = "50-100jt"
    BudgetOver100M    = ">100jt"
)
```

#### 5.6 AI Types
```go
const (
    AISPk            = "spk"
    AISistemPakar    = "sistem_pakar"
    AIDataAnalytics  = "data_analytics"
    AIMachineLearning = "machine_learning"
)
```

#### 5.7 Portfolio Categories
```go
const (
    CategoryWeb        = "web"
    CategoryMobile     = "mobile"
    CategoryAI         = "ai"
    CategoryEcommerce  = "ecommerce"
    CategoryOther      = "other"
)
```

#### 5.8 Category Types
```go
const (
    CategoryTypeBlog      = "blog"
    CategoryTypePortfolio = "portfolio"
)
```

#### 5.9 Allowed MIME Types
```go
const (
    MimeTypeJPEG = "image/jpeg"
    MimeTypePNG  = "image/png"
    MimeTypeWebP = "image/webp"
)
```

---

### 6. RELATIONSHIPS SUMMARY

| From | To | Type | On Delete | Description |
|---|---|---|---|---|
| `users` | `blog_posts` | 1:N | RESTRICT | User authors many posts |
| `users` | `media` | 1:N | RESTRICT | User uploads many media |
| `users` | `refresh_tokens` | 1:N | CASCADE | User has many refresh tokens |
| `users` | `password_resets` | 1:N | CASCADE | User has many password resets |
| `categories` | `blog_posts` | 1:N | SET NULL | Category has many posts |
| `blog_posts` | `tags` | M:N | CASCADE | Posts have many tags (via post_tags) |
| `case_studies` | `tags` | M:N | CASCADE | Case studies have many tags |
| `blog_posts` | `categories` | N:1 | - | Post belongs to category |
| `blog_posts` | `users` | N:1 | - | Post belongs to author |

---

### 7. INDEXES SUMMARY

#### 7.1 Unique Indexes
| Table | Column(s) | Index Name |
|---|---|---|
| `users` | `email` | `uq_users_email` |
| `refresh_tokens` | `token` | `uq_refresh_tokens_token` |
| `password_resets` | `token` | `uq_password_resets_token` |
| `categories` | `slug` | `uq_categories_slug` |
| `tags` | `slug`, `name` | `uq_tags_slug`, `uq_tags_name` |
| `blog_posts` | `slug` | `uq_blog_posts_slug` |
| `case_studies` | `slug` | `uq_case_studies_slug` |
| `portfolios` | `slug` | `uq_portfolios_slug` |
| `services` | `slug` | `uq_services_slug` |
| `media` | `filename` | `uq_media_filename` |
| `site_settings` | `key` | `uq_site_settings_key` |

#### 7.2 Performance Indexes
| Table | Column(s) | Index Name | Purpose |
|---|---|---|---|
| `blog_posts` | `status` | `idx_blog_posts_status` | Filter by status |
| `blog_posts` | `published_at` | `idx_blog_posts_published_at` | Sort by publish date |
| `blog_posts` | `author_id` | `idx_blog_posts_author_id` | Filter by author |
| `blog_posts` | `category_id` | `idx_blog_posts_category_id` | Filter by category |
| `case_studies` | `status` | `idx_case_studies_status` | Filter by status |
| `case_studies` | `industry` | `idx_case_studies_industry` | Filter by industry |
| `portfolios` | `category` | `idx_portfolios_category` | Filter by category |
| `portfolios` | `order_index` | `idx_portfolios_order_index` | Sort by order |
| `leads` | `status` | `idx_leads_status` | Filter by status |
| `leads` | `created_at` | `idx_leads_created_at` | Sort by date |
| `media` | `uploaded_by` | `idx_media_uploaded_by` | Filter by uploader |

---

### 8. GORM MODELS (Go Code)

```go
// backend/internal/models/user.go
package models

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type User struct {
    ID            uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    Name          string         `gorm:"type:varchar(100);not null" json:"name"`
    Email         string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
    PasswordHash  string         `gorm:"type:varchar(255);not null" json:"-"`
    Role          string         `gorm:"type:varchar(20);not null;default:'editor'" json:"role"`
    CreatedAt     time.Time      `gorm:"not null;default:now()" json:"created_at"`
    UpdatedAt     time.Time      `gorm:"not null;default:now()" json:"updated_at"`
    
    // Relationships
    BlogPosts     []BlogPost     `gorm:"foreignKey:AuthorID" json:"blog_posts,omitempty"`
    Media         []Media        `gorm:"foreignKey:UploadedBy" json:"media,omitempty"`
    RefreshTokens []RefreshToken `gorm:"foreignKey:UserID" json:"-"`
}

func (User) TableName() string {
    return "users"
}
```

```go
// backend/internal/models/blog_post.go
package models

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/datatypes"
    "gorm.io/gorm"
)

type BlogPost struct {
    ID              uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    Title           string         `gorm:"type:varchar(255);not null" json:"title"`
    Slug            string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"slug"`
    Content         string         `gorm:"type:text;not null" json:"content"`
    Excerpt         *string        `gorm:"type:text" json:"excerpt"`
    CoverImage      *string        `gorm:"type:varchar(500)" json:"cover_image"`
    AuthorID        uuid.UUID      `gorm:"type:uuid;not null;index" json:"author_id"`
    CategoryID      *uuid.UUID     `gorm:"type:uuid;index" json:"category_id"`
    Status          string         `gorm:"type:varchar(20);not null;default:'draft';index" json:"status"`
    MetaTitle       *string        `gorm:"type:varchar(255)" json:"meta_title"`
    MetaDescription *string        `gorm:"type:varchar(500)" json:"meta_description"`
    PublishedAt     *time.Time     `gorm:"index" json:"published_at"`
    ReadTime        *int           `json:"read_time"`
    CreatedAt       time.Time      `gorm:"not null;default:now()" json:"created_at"`
    UpdatedAt       time.Time      `gorm:"not null;default:now()" json:"updated_at"`
    DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
    
    // Relationships
    Author          User           `gorm:"foreignKey:AuthorID" json:"author"`
    Category        *Category      `gorm:"foreignKey:CategoryID" json:"category"`
    Tags            []Tag          `gorm:"many2many:post_tags;" json:"tags"`
}

func (BlogPost) TableName() string {
    return "blog_posts"
}
```

```go
// backend/internal/models/case_study.go
package models

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/datatypes"
    "gorm.io/gorm"
)

type CaseStudy struct {
    ID              uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    Title           string         `gorm:"type:varchar(255);not null" json:"title"`
    Slug            string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"slug"`
    ClientName      string         `gorm:"type:varchar(100);not null" json:"client_name"`
    Industry        string         `gorm:"type:varchar(100);not null;index" json:"industry"`
    Challenge       string         `gorm:"type:text;not null" json:"challenge"`
    Approach        string         `gorm:"type:text;not null" json:"approach"`
    Solution        string         `gorm:"type:text;not null" json:"solution"`
    AIType          datatypes.JSON `gorm:"type:jsonb" json:"ai_type"`
    Results         string         `gorm:"type:text;not null" json:"results"`
    Lessons         *string        `gorm:"type:text" json:"lessons"`
    Testimonial     *string        `gorm:"type:text" json:"testimonial"`
    Metrics         datatypes.JSON `gorm:"type:jsonb" json:"metrics"`
    CoverImage      *string        `gorm:"type:varchar(500)" json:"cover_image"`
    Images          datatypes.JSON `gorm:"type:jsonb" json:"images"`
    Status          string         `gorm:"type:varchar(20);not null;default:'draft';index" json:"status"`
    MetaTitle       *string        `gorm:"type:varchar(255)" json:"meta_title"`
    MetaDescription *string        `gorm:"type:varchar(500)" json:"meta_description"`
    PublishedAt     *time.Time     `gorm:"index" json:"published_at"`
    CreatedAt       time.Time      `gorm:"not null;default:now()" json:"created_at"`
    UpdatedAt       time.Time      `gorm:"not null;default:now()" json:"updated_at"`
    DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
    
    // Relationships
    Tags            []Tag          `gorm:"many2many:case_study_tags;" json:"tags"`
}

func (CaseStudy) TableName() string {
    return "case_studies"
}
```

```go
// backend/internal/models/portfolio.go
package models

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/datatypes"
    "gorm.io/gorm"
)

type Portfolio struct {
    ID              uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    Title           string         `gorm:"type:varchar(255);not null" json:"title"`
    Slug            string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"slug"`
    Description     string         `gorm:"type:text;not null" json:"description"`
    Category        string         `gorm:"type:varchar(50);not null;index" json:"category"`
    TechStack       datatypes.JSON `gorm:"type:jsonb" json:"tech_stack"`
    LiveURL         *string        `gorm:"type:varchar(500)" json:"live_url"`
    Thumbnail       *string        `gorm:"type:varchar(500)" json:"thumbnail"`
    Images          datatypes.JSON `gorm:"type:jsonb" json:"images"`
    Lessons         *string        `gorm:"type:text" json:"lessons"`
    Status          string         `gorm:"type:varchar(20);not null;default:'draft';index" json:"status"`
    OrderIndex      int            `gorm:"not null;default:0;index" json:"order_index"`
    MetaTitle       *string        `gorm:"type:varchar(255)" json:"meta_title"`
    MetaDescription *string        `gorm:"type:varchar(500)" json:"meta_description"`
    PublishedAt     *time.Time     `gorm:"index" json:"published_at"`
    CreatedAt       time.Time      `gorm:"not null;default:now()" json:"created_at"`
    UpdatedAt       time.Time      `gorm:"not null;default:now()" json:"updated_at"`
    DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Portfolio) TableName() string {
    return "portfolios"
}
```

```go
// backend/internal/models/lead.go
package models

import (
    "time"
    "github.com/google/uuid"
)

type Lead struct {
    ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    Name        string    `gorm:"type:varchar(100);not null" json:"name"`
    Email       string    `gorm:"type:varchar(255);not null;index" json:"email"`
    Phone       string    `gorm:"type:varchar(20);not null" json:"phone"`
    ServiceType string    `gorm:"type:varchar(50);not null;index" json:"service_type"`
    BudgetRange string    `gorm:"type:varchar(50);not null" json:"budget_range"`
    Description string    `gorm:"type:text;not null" json:"description"`
    Status      string    `gorm:"type:varchar(30);not null;default:'new';index" json:"status"`
    Notes       *string   `gorm:"type:text" json:"notes"`
    IPAddress   *string   `gorm:"type:varchar(45)" json:"ip_address"`
    UserAgent   *string   `gorm:"type:text" json:"user_agent"`
    CreatedAt   time.Time `gorm:"not null;default:now();index" json:"created_at"`
    UpdatedAt   time.Time `gorm:"not null;default:now()" json:"updated_at"`
}

func (Lead) TableName() string {
    return "leads"
}
```

```go
// backend/internal/models/media.go
package models

import (
    "time"
    "github.com/google/uuid"
)

type Media struct {
    ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    Filename     string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"filename"`
    OriginalName string    `gorm:"type:varchar(255);not null" json:"original_name"`
    URL          string    `gorm:"type:varchar(500);not null" json:"url"`
    MimeType     string    `gorm:"type:varchar(100);not null;index" json:"mime_type"`
    Size         int64     `gorm:"not null" json:"size"`
    Width        *int      `json:"width"`
    Height       *int      `json:"height"`
    AltText      *string   `gorm:"type:varchar(255)" json:"alt_text"`
    UploadedBy   uuid.UUID `gorm:"type:uuid;not null;index" json:"uploaded_by"`
    CreatedAt    time.Time `gorm:"not null;default:now();index" json:"created_at"`
    
    // Relationships
    Uploader     User      `gorm:"foreignKey:UploadedBy" json:"uploader"`
}

func (Media) TableName() string {
    return "media"
}
```

```go
// backend/internal/models/category.go
package models

import (
    "time"
    "github.com/google/uuid"
)

type Category struct {
    ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    Name      string    `gorm:"type:varchar(100);not null" json:"name"`
    Slug      string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"slug"`
    Type      string    `gorm:"type:varchar(20);not null;index" json:"type"`
    CreatedAt time.Time `gorm:"not null;default:now()" json:"created_at"`
    
    // Relationships
    BlogPosts []BlogPost `gorm:"foreignKey:CategoryID" json:"blog_posts,omitempty"`
}

func (Category) TableName() string {
    return "categories"
}
```

```go
// backend/internal/models/tag.go
package models

import (
    "time"
    "github.com/google/uuid"
)

type Tag struct {
    ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    Name      string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"`
    Slug      string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"slug"`
    CreatedAt time.Time `gorm:"not null;default:now()" json:"created_at"`
    
    // Relationships
    BlogPosts   []BlogPost   `gorm:"many2many:post_tags;" json:"blog_posts,omitempty"`
    CaseStudies []CaseStudy  `gorm:"many2many:case_study_tags;" json:"case_studies,omitempty"`
}

func (Tag) TableName() string {
    return "tags"
}
```

```go
// backend/internal/models/service.go
package models

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/datatypes"
)

type Service struct {
    ID              uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    Name            string         `gorm:"type:varchar(100);not null" json:"name"`
    Slug            string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"slug"`
    Description     string         `gorm:"type:text;not null" json:"description"`
    Icon            *string        `gorm:"type:varchar(50)" json:"icon"`
    UseCases        datatypes.JSON `gorm:"type:jsonb" json:"use_cases"`
    TimelineEstimate *string       `gorm:"type:varchar(100)" json:"timeline_estimate"`
    TechStack       datatypes.JSON `gorm:"type:jsonb" json:"tech_stack"`
    LearningPoints  *string        `gorm:"type:text" json:"learning_points"`
    IsActive        bool           `gorm:"not null;default:true;index" json:"is_active"`
    OrderIndex      int            `gorm:"not null;default:0;index" json:"order_index"`
    CreatedAt       time.Time      `gorm:"not null;default:now()" json:"created_at"`
    UpdatedAt       time.Time      `gorm:"not null;default:now()" json:"updated_at"`
}

func (Service) TableName() string {
    return "services"
}
```

```go
// backend/internal/models/site_setting.go
package models

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/datatypes"
)

type SiteSetting struct {
    ID          uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    Key         string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"key"`
    Value       datatypes.JSON `gorm:"type:jsonb;not null" json:"value"`
    Description *string        `gorm:"type:text" json:"description"`
    UpdatedAt   time.Time      `gorm:"not null;default:now()" json:"updated_at"`
}

func (SiteSetting) TableName() string {
    return "site_settings"
}
```

---

### 9. INITIAL MIGRATION SQL

```sql
-- migrations/000001_init_schema.up.sql

-- Enable UUID generation
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ============================================
-- USERS & AUTH
-- ============================================

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'editor',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_users_email UNIQUE (email),
    CONSTRAINT chk_users_role CHECK (role IN ('super_admin', 'admin', 'editor'))
);

CREATE INDEX idx_users_role ON users(role);

CREATE TABLE refresh_tokens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token VARCHAR(500) NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_refresh_tokens_token UNIQUE (token)
);

CREATE INDEX idx_refresh_tokens_user_id ON refresh_tokens(user_id);
CREATE INDEX idx_refresh_tokens_expires_at ON refresh_tokens(expires_at);

CREATE TABLE password_resets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token VARCHAR(255) NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    used_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_password_resets_token UNIQUE (token)
);

CREATE INDEX idx_password_resets_user_id ON password_resets(user_id);

-- ============================================
-- CATEGORIES & TAGS
-- ============================================

CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    slug VARCHAR(100) NOT NULL,
    type VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_categories_slug UNIQUE (slug),
    CONSTRAINT chk_categories_type CHECK (type IN ('blog', 'portfolio'))
);

CREATE INDEX idx_categories_type ON categories(type);

CREATE TABLE tags (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50) NOT NULL,
    slug VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_tags_slug UNIQUE (slug),
    CONSTRAINT uq_tags_name UNIQUE (name)
);

-- ============================================
-- BLOG POSTS
-- ============================================

CREATE TABLE blog_posts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    excerpt TEXT,
    cover_image VARCHAR(500),
    author_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    category_id UUID REFERENCES categories(id) ON DELETE SET NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'draft',
    meta_title VARCHAR(255),
    meta_description VARCHAR(500),
    published_at TIMESTAMPTZ,
    read_time INTEGER,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT uq_blog_posts_slug UNIQUE (slug),
    CONSTRAINT chk_blog_posts_status CHECK (status IN ('draft', 'published', 'archived'))
);

CREATE INDEX idx_blog_posts_author_id ON blog_posts(author_id);
CREATE INDEX idx_blog_posts_category_id ON blog_posts(category_id);
CREATE INDEX idx_blog_posts_status ON blog_posts(status);
CREATE INDEX idx_blog_posts_published_at ON blog_posts(published_at);
CREATE INDEX idx_blog_posts_created_at ON blog_posts(created_at);
CREATE INDEX idx_blog_posts_deleted_at ON blog_posts(deleted_at);

CREATE TABLE post_tags (
    post_id UUID NOT NULL REFERENCES blog_posts(id) ON DELETE CASCADE,
    tag_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    PRIMARY KEY (post_id, tag_id)
);

CREATE INDEX idx_post_tags_post_id ON post_tags(post_id);
CREATE INDEX idx_post_tags_tag_id ON post_tags(tag_id);

-- ============================================
-- CASE STUDIES
-- ============================================

CREATE TABLE case_studies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL,
    client_name VARCHAR(100) NOT NULL,
    industry VARCHAR(100) NOT NULL,
    challenge TEXT NOT NULL,
    approach TEXT NOT NULL,
    solution TEXT NOT NULL,
    ai_type JSONB,
    results TEXT NOT NULL,
    lessons TEXT,
    testimonial TEXT,
    metrics JSONB,
    cover_image VARCHAR(500),
    images JSONB,
    status VARCHAR(20) NOT NULL DEFAULT 'draft',
    meta_title VARCHAR(255),
    meta_description VARCHAR(500),
    published_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT uq_case_studies_slug UNIQUE (slug),
    CONSTRAINT chk_case_studies_status CHECK (status IN ('draft', 'published', 'archived'))
);

CREATE INDEX idx_case_studies_status ON case_studies(status);
CREATE INDEX idx_case_studies_industry ON case_studies(industry);
CREATE INDEX idx_case_studies_published_at ON case_studies(published_at);
CREATE INDEX idx_case_studies_deleted_at ON case_studies(deleted_at);
CREATE INDEX idx_case_studies_ai_type ON case_studies USING GIN (ai_type);

CREATE TABLE case_study_tags (
    case_study_id UUID NOT NULL REFERENCES case_studies(id) ON DELETE CASCADE,
    tag_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    PRIMARY KEY (case_study_id, tag_id)
);

CREATE INDEX idx_case_study_tags_case_study_id ON case_study_tags(case_study_id);
CREATE INDEX idx_case_study_tags_tag_id ON case_study_tags(tag_id);

-- ============================================
-- PORTFOLIOS
-- ============================================

CREATE TABLE portfolios (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    category VARCHAR(50) NOT NULL,
    tech_stack JSONB,
    live_url VARCHAR(500),
    thumbnail VARCHAR(500),
    images JSONB,
    lessons TEXT,
    status VARCHAR(20) NOT NULL DEFAULT 'draft',
    order_index INTEGER NOT NULL DEFAULT 0,
    meta_title VARCHAR(255),
    meta_description VARCHAR(500),
    published_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT uq_portfolios_slug UNIQUE (slug),
    CONSTRAINT chk_portfolios_status CHECK (status IN ('draft', 'published', 'archived')),
    CONSTRAINT chk_portfolios_category CHECK (category IN ('web', 'mobile', 'ai', 'ecommerce', 'other'))
);

CREATE INDEX idx_portfolios_category ON portfolios(category);
CREATE INDEX idx_portfolios_status ON portfolios(status);
CREATE INDEX idx_portfolios_order_index ON portfolios(order_index);
CREATE INDEX idx_portfolios_published_at ON portfolios(published_at);
CREATE INDEX idx_portfolios_deleted_at ON portfolios(deleted_at);

-- ============================================
-- SERVICES
-- ============================================

CREATE TABLE services (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    slug VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    icon VARCHAR(50),
    use_cases JSONB,
    timeline_estimate VARCHAR(100),
    tech_stack JSONB,
    learning_points TEXT,
    is_active BOOLEAN NOT NULL DEFAULT true,
    order_index INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_services_slug UNIQUE (slug)
);

CREATE INDEX idx_services_is_active ON services(is_active);
CREATE INDEX idx_services_order_index ON services(order_index);

-- ============================================
-- LEADS
-- ============================================

CREATE TABLE leads (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    service_type VARCHAR(50) NOT NULL,
    budget_range VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    status VARCHAR(30) NOT NULL DEFAULT 'new',
    notes TEXT,
    ip_address VARCHAR(45),
    user_agent TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_leads_status CHECK (status IN ('new', 'contacted', 'negotiation', 'proposal_sent', 'won', 'lost')),
    CONSTRAINT chk_leads_service_type CHECK (service_type IN ('web_development', 'mobile_app', 'uiux_design', 'ai_integration', 'other')),
    CONSTRAINT chk_leads_budget_range CHECK (budget_range IN ('<5jt', '5-15jt', '15-50jt', '50-100jt', '>100jt'))
);

CREATE INDEX idx_leads_status ON leads(status);
CREATE INDEX idx_leads_email ON leads(email);
CREATE INDEX idx_leads_created_at ON leads(created_at);
CREATE INDEX idx_leads_service_type ON leads(service_type);

-- ============================================
-- MEDIA
-- ============================================

CREATE TABLE media (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    filename VARCHAR(255) NOT NULL,
    original_name VARCHAR(255) NOT NULL,
    url VARCHAR(500) NOT NULL,
    mime_type VARCHAR(100) NOT NULL,
    size BIGINT NOT NULL,
    width INTEGER,
    height INTEGER,
    alt_text VARCHAR(255),
    uploaded_by UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_media_filename UNIQUE (filename),
    CONSTRAINT chk_media_mime_type CHECK (mime_type IN ('image/jpeg', 'image/png', 'image/webp')),
    CONSTRAINT chk_media_size CHECK (size <= 5242880)
);

CREATE INDEX idx_media_uploaded_by ON media(uploaded_by);
CREATE INDEX idx_media_mime_type ON media(mime_type);
CREATE INDEX idx_media_created_at ON media(created_at);

-- ============================================
-- SITE SETTINGS
-- ============================================

CREATE TABLE site_settings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    key VARCHAR(100) NOT NULL,
    value JSONB NOT NULL,
    description TEXT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_site_settings_key UNIQUE (key)
);

-- ============================================
-- UPDATED_AT TRIGGER FUNCTION
-- ============================================

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Apply trigger to all tables with updated_at
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_blog_posts_updated_at BEFORE UPDATE ON blog_posts
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_case_studies_updated_at BEFORE UPDATE ON case_studies
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_portfolios_updated_at BEFORE UPDATE ON portfolios
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_services_updated_at BEFORE UPDATE ON services
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_leads_updated_at BEFORE UPDATE ON leads
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_site_settings_updated_at BEFORE UPDATE ON site_settings
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
```

---

### 10. SEED DATA

```sql
-- migrations/000002_seed_data.up.sql

-- ============================================
-- SEED: Default Admin User
-- ============================================
-- Password: admin123 (bcrypt hash)
INSERT INTO users (id, name, email, password_hash, role) VALUES
('550e8400-e29b-41d4-a716-446655440000', 'Rizal', 'admin@webifylab.my.id', '$2a$10$...', 'super_admin');

-- ============================================
-- SEED: Default Categories
-- ============================================
INSERT INTO categories (id, name, slug, type) VALUES
-- Blog categories
(gen_random_uuid(), 'Tutorial', 'tutorial', 'blog'),
(gen_random_uuid(), 'Business', 'business', 'blog'),
(gen_random_uuid(), 'AI & Data', 'ai-data', 'blog'),
(gen_random_uuid(), 'Eksperimen', 'eksperimen', 'blog'),
(gen_random_uuid(), 'Lessons Learned', 'lessons-learned', 'blog'),
-- Portfolio categories
(gen_random_uuid(), 'Web Development', 'web-development', 'portfolio'),
(gen_random_uuid(), 'Mobile App', 'mobile-app', 'portfolio'),
(gen_random_uuid(), 'AI Integration', 'ai-integration', 'portfolio'),
(gen_random_uuid(), 'E-Commerce', 'e-commerce', 'portfolio');

-- ============================================
-- SEED: Default Tags
-- ============================================
INSERT INTO tags (id, name, slug) VALUES
(gen_random_uuid(), 'Next.js', 'nextjs'),
(gen_random_uuid(), 'Golang', 'golang'),
(gen_random_uuid(), 'PostgreSQL', 'postgresql'),
(gen_random_uuid(), 'Docker', 'docker'),
(gen_random_uuid(), 'SPK', 'spk'),
(gen_random_uuid(), 'Sistem Pakar', 'sistem-pakar'),
(gen_random_uuid(), 'Machine Learning', 'machine-learning'),
(gen_random_uuid(), 'UMKM', 'umkm'),
(gen_random_uuid(), 'Startup', 'startup');

-- ============================================
-- SEED: Default Services
-- ============================================
INSERT INTO services (id, name, slug, description, icon, use_cases, timeline_estimate, tech_stack, learning_points, is_active, order_index) VALUES
(gen_random_uuid(), 'Web Development', 'web-development', 'Pembuatan website profesional untuk UMKM dan Startup', 'globe', 
 '["Company profile", "Landing page", "Web app", "E-commerce"]'::jsonb,
 '2-6 minggu', '["Next.js", "Golang", "PostgreSQL", "Tailwind CSS"]'::jsonb,
 'Mari kita eksperimen bareng bagaimana website bisa jadi motor pertumbuhan bisnis kamu', true, 1),

(gen_random_uuid(), 'Mobile App Development', 'mobile-app', 'Aplikasi mobile cross-platform untuk iOS dan Android', 'smartphone',
 '["Mobile app untuk bisnis", "Aplikasi internal", "Progressive Web App"]'::jsonb,
 '4-12 minggu', '["React Native", "Flutter", "Golang", "PostgreSQL"]'::jsonb,
 'Belajar bareng bagaimana mobile app bisa tingkatkan engagement customer kamu', true, 2),

(gen_random_uuid(), 'UI/UX Design', 'uiux-design', 'Desain interface yang user-friendly dan美观', 'palette',
 '["Redesign website", "Mobile app UI", "Design system"]'::jsonb,
 '1-4 minggu', '["Figma", "Adobe XD", "User Research"]'::jsonb,
 'Eksperimen bareng bagaimana desain yang tepat bisa tingkatkan konversi', true, 3),

(gen_random_uuid(), 'AI Integration', 'ai-integration', 'Integrasi algoritma AI untuk optimasi bisnis', 'brain',
 '["Sistem Pendukung Keputusan (SPK)", "Sistem Pakar", "Data Analytics", "Machine Learning"]'::jsonb,
 '3-8 minggu', '["Python", "Golang", "PostgreSQL", "Scikit-learn", "TensorFlow"]'::jsonb,
 'Mari kita eksperimen bareng bagaimana AI bisa bantu kamu ambil keputusan bisnis lebih cerdas', true, 4);

-- ============================================
-- SEED: Default Site Settings
-- ============================================
INSERT INTO site_settings (key, value, description) VALUES
('site_name', '"WebifyLab"'::jsonb, 'Nama website'),
('tagline', '"We Experiment, You Grow"'::jsonb, 'Tagline website'),
('contact_email', '"hello@webifylab.my.id"'::jsonb, 'Email kontak'),
('contact_whatsapp', '"6281234567890"'::jsonb, 'Nomor WhatsApp (format internasional)'),
('social_media', '{"instagram": "https://instagram.com/webifylab", "linkedin": "https://linkedin.com/company/webifylab", "github": "https://github.com/webifylab"}'::jsonb, 'Link social media'),
('seo_default_title', '"WebifyLab - Digital Learning Lab"'::jsonb, 'Default SEO title'),
('seo_default_description', '"WebifyLab adalah digital learning lab yang membantu UMKM dan Startup tumbuh melalui eksperimen digital."'::jsonb, 'Default SEO description'),
('seo_og_image', '"/images/og-default.jpg"'::jsonb, 'Default Open Graph image');
```

---

### 11. NOTES & CONSIDERATIONS

#### 11.1 GORM Auto-Migration vs Manual Migration

**Development:**
- Gunakan GORM `AutoMigrate()` untuk快速 development
- GORM akan create tables dan indexes otomatis berdasarkan models

**Production:**
- **JANGAN** pakai `AutoMigrate()` di production
- Gunakan `golang-migrate` untuk version-controlled migrations
- Migration files di `/backend/migrations/`
- Run migrations via CLI: `./server migrate`

#### 11.2 JSONB Fields

**When to use JSONB:**
- `tech_stack`, `images`, `use_cases` — arrays of strings
- `metrics`, `ai_type` — flexible structured data
- `value` in `site_settings` — any JSON value

**Querying JSONB:**
```go
// Example: Find case studies with SPK AI type
db.Where("ai_type @> ?", pq.Array([]string{"spk"})).Find(&caseStudies)

// Example: Get metrics from case study
var metrics map[string]interface{}
json.Unmarshal(caseStudy.Metrics, &metrics)
```

#### 11.3 Soft Delete

**Tables with soft delete:**
- `blog_posts`
- `case_studies`
- `portfolios`

**Implementation:**
- GORM automatically handle `DeletedAt` field
- Queries automatically filter out soft-deleted records
- To include soft-deleted: `db.Unscoped().Find(&posts)`
- To restore: `db.Unscoped().Where("id = ?", id).Update("deleted_at", nil)`

#### 11.4 Performance Considerations

**Indexes:**
- All foreign keys have indexes
- Status fields have indexes (for filtering)
- Slug fields have unique indexes (for lookups)
- Created_at fields have indexes (for sorting)

**JSONB Indexes:**
- `case_studies.ai_type` has GIN index for array queries
- Consider GIN indexes untuk other JSONB fields jika perlu complex queries

**Connection Pooling:**
- GORM default connection pool: 10 connections
- Adjust based on VPS RAM: `db.SetMaxOpenConns(20)`

#### 11.5 Data Integrity

**Foreign Key Constraints:**
- `ON DELETE CASCADE` — Delete child records when parent deleted (refresh_tokens, post_tags)
- `ON DELETE SET NULL` — Set child FK to NULL when parent deleted (blog_posts.category_id)
- `ON DELETE RESTRICT` — Prevent parent deletion if children exist (users, media)

**Check Constraints:**
- Enum fields have CHECK constraints
- File size has CHECK constraint (≤ 5MB)
- MIME type has CHECK constraint

#### 11.6 Security Considerations

**Password Hashing:**
- Use bcrypt with cost 10-12
- Never store plain text passwords
- `password_hash` field excluded from JSON responses (`json:"-"`)

**UUID Primary Keys:**
- Prevent sequential ID guessing
- More secure than auto-increment integers
- Slightly larger storage, but worth it for security

**Soft Delete:**
- Prevent accidental data loss
- Audit trail for compliance
- Can restore deleted records if needed

#### 11.7 Future Enhancements

**Phase 2 (SSO):**
- Add `oauth_accounts` table untuk link external providers
- Add `sessions` table untuk session management
- Consider adding `organizations` table untuk multi-tenant

**Phase 3 (Scale):**
- Add `audit_logs` table untuk track all changes
- Add `notifications` table untuk in-app notifications
- Consider partitioning large tables (leads, media)

---

### 12. SUMMARY

#### 12.1 Tables Count
- **Total Tables:** 14
- **Core Tables:** 10 (users, blog_posts, case_studies, portfolios, services, leads, media, categories, tags, site_settings)
- **Junction Tables:** 2 (post_tags, case_study_tags)
- **Auth Tables:** 2 (refresh_tokens, password_resets)

#### 12.2 Relationships
- **One-to-Many:** 8 relationships
- **Many-to-Many:** 2 relationships (via junction tables)
- **Self-referencing:** 0

#### 12.3 Indexes
- **Unique Indexes:** 11
- **Performance Indexes:** 25+
- **GIN Indexes:** 1 (for JSONB)

#### 12.4 Key Features
- ✅ UUID primary keys (secure, SSO-ready)
- ✅ Timezone-aware timestamps
- ✅ Soft delete untuk critical tables
- ✅ JSONB untuk flexible data
- ✅ Comprehensive indexing strategy
- ✅ Foreign key constraints
- ✅ Check constraints untuk enums
- ✅ Auto-updated `updated_at` triggers
- ✅ GORM models ready
- ✅ Migration scripts ready
- ✅ Seed data ready