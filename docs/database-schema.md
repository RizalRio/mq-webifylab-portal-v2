# 🗄️ Database Schema Document
## Webifylab Ecosystem — Version 1.0

| Metadata         | Detail                                      |
|------------------|---------------------------------------------|
| **Product**      | Webifylab Ecosystem                         |
| **Version**      | 1.0                                         |
| **Author**       | Rizal                                       |
| **Created**      | 17 September 2026                           |
| **Status**       | Draft (Design Phase)                        |
| **Database**     | PostgreSQL 16+                              |
| **Implementation** | V1.5 (Contact + Analytics), V2 (SaaS)    |

---

## 1. Design Principles

### 1.1 Core Principles

| Principle | Deskripsi |
|-----------|-----------|
| **Normalized (3NF)** | Hindari data redundancy, pastikan integritas |
| **Audit Trail** | Setiap tabel punya `created_at`, `updated_at` |
| **Soft Delete** | Gunakan `deleted_at` daripada hard delete |
| **Multi-tenant Ready** | Schema siap untuk SaaS multi-tenant di V2 |
| **AI/Vector Ready** | Support pgvector untuk embeddings & RAG |
| **Performance First** | Indexes strategis, query optimization |
| **Security** | Row-Level Security (RLS) untuk multi-tenant |

### 1.2 Naming Conventions

| Element | Convention | Example |
|---------|-----------|---------|
| **Table names** | snake_case, plural | `users`, `contact_submissions` |
| **Column names** | snake_case | `first_name`, `created_at` |
| **Primary keys** | `id` (UUID) | `id UUID PRIMARY KEY` |
| **Foreign keys** | `{table}_id` | `user_id`, `project_id` |
| **Timestamps** | `created_at`, `updated_at`, `deleted_at` | `TIMESTAMPTZ DEFAULT NOW()` |
| **Booleans** | `is_*` atau `has_*` | `is_active`, `has_subscription` |
| **Enums** | snake_case | `app_development`, `web_design` |
| **Indexes** | `idx_{table}_{column}` | `idx_users_email` |

### 1.3 Data Types Strategy

| Use Case | PostgreSQL Type | Alasan |
|----------|----------------|--------|
| **Primary Key** | `UUID` | Globally unique, aman untuk API |
| **Timestamps** | `TIMESTAMPTZ` | Timezone-aware |
| **JSON Data** | `JSONB` | Flexible, indexable, fast |
| **Text (long)** | `TEXT` | No length limit |
| **Text (short)** | `VARCHAR(n)` | With length constraint |
| **Enum-like** | `VARCHAR` + CHECK | More flexible than ENUM type |
| **Money** | `DECIMAL(10,2)` | Precise, no floating point issues |
| **Boolean** | `BOOLEAN` | Native type |
| **Arrays** | `TEXT[]` atau `UUID[]` | Native array support |
| **Vectors (AI)** | `VECTOR(1536)` | pgvector extension |

---

## 2. Entity Relationship Diagram (ERD)

```
┌─────────────────────────────────────────────────────────────────────────┐
│                          WEBIFYLAB ECOSYSTEM                            │
└─────────────────────────────────────────────────────────────────────────┘

┌──────────────┐       ┌──────────────────────┐       ┌──────────────┐
│    users     │       │ contact_submissions  │       │  portfolios  │
├──────────────┤       ├──────────────────────┤       ├──────────────┤
│ id (PK)      │       │ id (PK)              │       │ id (PK)      │
│ email        │       │ name                 │       │ title        │
│ password_hash│       │ email                │       │ category     │
│ full_name    │       │ service_type         │       │ description  │
│ role         │       │ message              │       │ thumbnail_url│
│ avatar_url   │       │ status               │       │ project_url  │
│ is_active    │       │ source               │       │ tech_stack   │
│ created_at   │       │ ip_address           │       │ is_published │
│ updated_at   │       │ user_agent           │       │ created_at   │
│ deleted_at   │       │ created_at           │       │ updated_at   │
└──────┬───────┘       │ updated_at           │       └──────────────┘
       │               └──────────────────────┘
       │
       │ 1:N
       ▼
┌──────────────┐       ┌──────────────────────┐       ┌──────────────┐
│   sessions   │       │  analytics_events    │       │   projects   │
├──────────────┤       ├──────────────────────┤       ├──────────────┤
│ id (PK)      │       │ id (PK)              │       │ id (PK)      │
│ user_id (FK) │       │ event_type           │       │ client_id    │
│ token        │       │ user_id (FK, null)   │       │ user_id (FK) │
│ ip_address   │       │ session_id           │       │ title        │
│ user_agent   │       │ page_url             │       │ description  │
│ expires_at   │       │ metadata (JSONB)     │       │ status       │
│ created_at   │       │ ip_address           │       │ start_date   │
│ revoked_at   │       │ created_at           │       │ end_date     │
└──────────────┘       └──────────────────────┘       │ budget       │
                                                      │ created_at   │
                                                      │ updated_at   │
                                                      └──────┬───────┘
                                                             │
                                                             │ 1:N
                                                             ▼
┌──────────────┐       ┌──────────────────────┐       ┌──────────────┐
│  tenants     │       │   subscriptions      │       │  invoices    │
├──────────────┤       ├──────────────────────┤       ├──────────────┤
│ id (PK)      │       │ id (PK)              │       │ id (PK)      │
│ name         │       │ tenant_id (FK)       │       │ subscription_│
│ slug         │       │ plan                 │       │  id (FK)     │
│ owner_id(FK) │       │ status               │       │ amount       │
│ settings     │       │ current_period_start │       │ currency     │
│ is_active    │       │ current_period_end   │       │ status       │
│ created_at   │       │ cancel_at_period_end │       │ paid_at      │
│ updated_at   │       │ created_at           │       │ due_date     │
│ deleted_at   │       │ updated_at           │       │ created_at   │
└──────────────┘       └──────────────────────┘       └──────────────┘

┌──────────────┐       ┌──────────────────────┐       ┌──────────────┐
│  saas_products│      │   ai_conversations   │       │  embeddings  │
├──────────────┤       ├──────────────────────┤       ├──────────────┤
│ id (PK)      │       │ id (PK)              │       │ id (PK)      │
│ tenant_id(FK)│       │ user_id (FK)         │       │ content_type │
│ name         │       │ title                │       │ content_id   │
│ description  │       │ model                │       │ chunk_text   │
│ version      │       │ total_tokens         │       │ embedding    │
│ config       │       │ created_at           │       │  (vector)    │
│ is_active    │       │ updated_at           │       │ metadata     │
│ created_at   │       └──────────┬───────────┘       │ created_at   │
│ updated_at   │                  │                   └──────────────┘
└──────────────┘                  │ 1:N
                                  ▼
                       ┌──────────────────────┐
                       │   ai_messages        │
                       ├──────────────────────┤
                       │ id (PK)              │
                       │ conversation_id (FK) │
                       │ role                 │
                       │ content              │
                       │ tokens_used          │
                       │ model                │
                       │ created_at           │
                       └──────────────────────┘
```

---

## 3. Table Definitions

### 3.1 V1.5 Tables (Contact Form + Analytics)

#### Table: `contact_submissions`

**Purpose:** Menyimpan semua pesan kontak dari form landing page (menggantikan Formspree)

```sql
CREATE TABLE contact_submissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Contact info
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL,
    service_type VARCHAR(50) NOT NULL,
    message TEXT NOT NULL,
    
    -- Status tracking
    status VARCHAR(20) NOT NULL DEFAULT 'new'
        CHECK (status IN ('new', 'read', 'replied', 'archived', 'spam')),
    
    -- Metadata
    source VARCHAR(50) DEFAULT 'website',  -- website, whatsapp, email
    ip_address INET,
    user_agent TEXT,
    referrer_url TEXT,
    
    -- Admin notes
    admin_notes TEXT,
    assigned_to UUID REFERENCES users(id),
    
    -- Timestamps
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Indexes
CREATE INDEX idx_contact_submissions_status ON contact_submissions(status);
CREATE INDEX idx_contact_submissions_email ON contact_submissions(email);
CREATE INDEX idx_contact_submissions_created_at ON contact_submissions(created_at DESC);
CREATE INDEX idx_contact_submissions_service_type ON contact_submissions(service_type);

-- Full-text search index
CREATE INDEX idx_contact_submissions_search ON contact_submissions 
    USING GIN(to_tsvector('indonesian', name || ' ' || message));

-- Trigger for updated_at
CREATE TRIGGER update_contact_submissions_updated_at
    BEFORE UPDATE ON contact_submissions
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
```

**Columns Detail:**

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| `id` | UUID | PK, auto-generated | Unique identifier |
| `name` | VARCHAR(100) | NOT NULL | Nama pengirim |
| `email` | VARCHAR(255) | NOT NULL | Email pengirim |
| `service_type` | VARCHAR(50) | NOT NULL, CHECK | Jenis layanan yang diminati |
| `message` | TEXT | NOT NULL | Pesan dari pengunjung |
| `status` | VARCHAR(20) | NOT NULL, DEFAULT 'new' | Status follow-up |
| `source` | VARCHAR(50) | DEFAULT 'website' | Sumber kontak |
| `ip_address` | INET | - | IP address pengirim |
| `user_agent` | TEXT | - | Browser/device info |
| `referrer_url` | TEXT | - | URL asal pengunjung |
| `admin_notes` | TEXT | - | Catatan internal admin |
| `assigned_to` | UUID | FK → users | Admin yang menangani |
| `created_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | Waktu submit |
| `updated_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | Waktu update terakhir |
| `deleted_at` | TIMESTAMPTZ | NULL | Soft delete timestamp |

**Service Type Enum Values:**
- `app_development` — Pengembangan Aplikasi
- `web_design` — Desain Grafis & Web Design
- `saas_ecosystem` — SaaS & Ekosistem Digital
- `ai_data_consulting` — Konsultasi AI & Data
- `other` — Lainnya

---

#### Table: `analytics_events`

**Purpose:** Self-hosted analytics (menggantikan Google Analytics/Plausible)

```sql
CREATE TABLE analytics_events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Event data
    event_type VARCHAR(50) NOT NULL,
    session_id UUID NOT NULL,
    user_id UUID REFERENCES users(id),  -- NULL for anonymous
    
    -- Page/URL info
    page_url TEXT NOT NULL,
    page_title TEXT,
    referrer_url TEXT,
    
    -- Metadata (flexible JSON)
    metadata JSONB DEFAULT '{}',
    
    -- Device/Location
    ip_address INET,
    user_agent TEXT,
    country_code VARCHAR(2),
    city VARCHAR(100),
    
    -- Timestamp
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Indexes
CREATE INDEX idx_analytics_events_session_id ON analytics_events(session_id);
CREATE INDEX idx_analytics_events_event_type ON analytics_events(event_type);
CREATE INDEX idx_analytics_events_page_url ON analytics_events(page_url);
CREATE INDEX idx_analytics_events_created_at ON analytics_events(created_at DESC);
CREATE INDEX idx_analytics_events_user_id ON analytics_events(user_id);

-- JSONB index for metadata queries
CREATE INDEX idx_analytics_events_metadata ON analytics_events USING GIN(metadata);

-- Partition by month (for performance)
-- CREATE TABLE analytics_events_y2026m09 PARTITION OF analytics_events
--     FOR VALUES FROM ('2026-09-01') TO ('2026-10-01');
```

**Event Types:**
- `page_view` — Pengunjung membuka halaman
- `cta_click` — Klik tombol CTA
- `form_submit` — Submit form kontak
- `form_error` — Error saat submit form
- `scroll_depth` — Scroll mencapai % tertentu
- `session_start` — Sesi dimulai
- `session_end` — Sesi berakhir

**Metadata Examples:**

```json
// page_view
{
  "load_time_ms": 1200,
  "viewport_width": 1920,
  "viewport_height": 1080
}

// cta_click
{
  "button_text": "Konsultasi Gratis",
  "section": "hero",
  "target_url": "/contact"
}

// form_submit
{
  "form_name": "contact",
  "service_type": "app_development",
  "success": true
}
```

---

### 3.2 V2 Tables (SaaS Ecosystem)

#### Table: `users`

**Purpose:** User accounts untuk SaaS platform

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Authentication
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    email_verified_at TIMESTAMPTZ,
    
    -- Profile
    full_name VARCHAR(100) NOT NULL,
    avatar_url TEXT,
    bio TEXT,
    
    -- Role & permissions
    role VARCHAR(20) NOT NULL DEFAULT 'user'
        CHECK (role IN ('user', 'admin', 'super_admin')),
    
    -- Status
    is_active BOOLEAN NOT NULL DEFAULT true,
    last_login_at TIMESTAMPTZ,
    
    -- Timestamps
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Indexes
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role ON users(role);
CREATE INDEX idx_users_is_active ON users(is_active) WHERE is_active = true;
CREATE INDEX idx_users_created_at ON users(created_at DESC);
```

---

#### Table: `tenants`

**Purpose:** Multi-tenant architecture untuk SaaS (setiap klien = 1 tenant)

```sql
CREATE TABLE tenants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Identity
    name VARCHAR(100) NOT NULL,
    slug VARCHAR(50) NOT NULL UNIQUE,
    
    -- Ownership
    owner_id UUID NOT NULL REFERENCES users(id),
    
    -- Settings (flexible JSON)
    settings JSONB DEFAULT '{}',
    
    -- Status
    is_active BOOLEAN NOT NULL DEFAULT true,
    plan VARCHAR(20) NOT NULL DEFAULT 'free'
        CHECK (plan IN ('free', 'starter', 'professional', 'enterprise')),
    
    -- Timestamps
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Indexes
CREATE INDEX idx_tenants_slug ON tenants(slug);
CREATE INDEX idx_tenants_owner_id ON tenants(owner_id);
CREATE INDEX idx_tenants_plan ON tenants(plan);
CREATE INDEX idx_tenants_is_active ON tenants(is_active) WHERE is_active = true;

-- Row-Level Security (RLS) for multi-tenant
ALTER TABLE tenants ENABLE ROW LEVEL SECURITY;

CREATE POLICY tenant_isolation ON tenants
    USING (id = current_setting('app.current_tenant_id')::UUID);
```

---

#### Table: `projects`

**Purpose:** Manajemen proyek klien (untuk portfolio & project tracking)

```sql
CREATE TABLE projects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Client info
    client_id UUID REFERENCES tenants(id),
    user_id UUID NOT NULL REFERENCES users(id),  -- Project owner
    
    -- Project details
    title VARCHAR(200) NOT NULL,
    description TEXT,
    status VARCHAR(30) NOT NULL DEFAULT 'planning'
        CHECK (status IN ('planning', 'in_progress', 'completed', 'on_hold', 'cancelled')),
    
    -- Timeline
    start_date DATE,
    end_date DATE,
    actual_end_date DATE,
    
    -- Budget
    budget DECIMAL(12, 2),
    currency VARCHAR(3) DEFAULT 'IDR',
    
    -- Metadata
    tech_stack TEXT[],  -- Array of technologies
    project_url TEXT,
    thumbnail_url TEXT,
    
    -- Timestamps
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Indexes
CREATE INDEX idx_projects_user_id ON projects(user_id);
CREATE INDEX idx_projects_client_id ON projects(client_id);
CREATE INDEX idx_projects_status ON projects(status);
CREATE INDEX idx_projects_start_date ON projects(start_date);
```

---

#### Table: `portfolios`

**Purpose:** Showcase proyek untuk landing page (public-facing)

```sql
CREATE TABLE portfolios (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Content
    title VARCHAR(200) NOT NULL,
    slug VARCHAR(200) NOT NULL UNIQUE,
    category VARCHAR(50) NOT NULL
        CHECK (category IN ('web_app', 'mobile_app', 'ui_ux_design', 'saas', 'other')),
    description TEXT NOT NULL,
    
    -- Media
    thumbnail_url TEXT NOT NULL,
    images TEXT[],  -- Array of image URLs
    project_url TEXT,
    case_study_url TEXT,
    
    -- Details
    tech_stack TEXT[],
    client_name VARCHAR(100),
    client_industry VARCHAR(50),
    
    -- Publishing
    is_published BOOLEAN NOT NULL DEFAULT false,
    published_at TIMESTAMPTZ,
    featured BOOLEAN NOT NULL DEFAULT false,
    
    -- SEO
    meta_title VARCHAR(60),
    meta_description VARCHAR(160),
    
    -- Timestamps
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Indexes
CREATE INDEX idx_portfolios_slug ON portfolios(slug);
CREATE INDEX idx_portfolios_category ON portfolios(category);
CREATE INDEX idx_portfolios_is_published ON portfolios(is_published) WHERE is_published = true;
CREATE INDEX idx_portfolios_featured ON portfolios(featured) WHERE featured = true;
CREATE INDEX idx_portfolios_published_at ON portfolios(published_at DESC);

-- Full-text search
CREATE INDEX idx_portfolios_search ON portfolios 
    USING GIN(to_tsvector('indonesian', title || ' ' || description));
```

---

#### Table: `subscriptions`

**Purpose:** Manajemen subscription SaaS (billing)

```sql
CREATE TABLE subscriptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Ownership
    tenant_id UUID NOT NULL REFERENCES tenants(id),
    
    -- Plan details
    plan VARCHAR(20) NOT NULL
        CHECK (plan IN ('starter', 'professional', 'enterprise')),
    status VARCHAR(20) NOT NULL DEFAULT 'active'
        CHECK (status IN ('active', 'past_due', 'cancelled', 'trialing')),
    
    -- Billing cycle
    current_period_start TIMESTAMPTZ NOT NULL,
    current_period_end TIMESTAMPTZ NOT NULL,
    cancel_at_period_end BOOLEAN NOT NULL DEFAULT false,
    
    -- Payment
    amount DECIMAL(10, 2) NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'IDR',
    payment_method VARCHAR(50),
    
    -- External references
    payment_gateway_id VARCHAR(100),  -- Midtrans/Xendit ID
    
    -- Timestamps
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Indexes
CREATE INDEX idx_subscriptions_tenant_id ON subscriptions(tenant_id);
CREATE INDEX idx_subscriptions_status ON subscriptions(status);
CREATE INDEX idx_subscriptions_current_period_end ON subscriptions(current_period_end);
```

---

#### Table: `invoices`

**Purpose:** Invoice/billing records

```sql
CREATE TABLE invoices (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- References
    subscription_id UUID NOT NULL REFERENCES subscriptions(id),
    tenant_id UUID NOT NULL REFERENCES tenants(id),
    
    -- Invoice details
    invoice_number VARCHAR(50) NOT NULL UNIQUE,
    amount DECIMAL(10, 2) NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'IDR',
    
    -- Status
    status VARCHAR(20) NOT NULL DEFAULT 'pending'
        CHECK (status IN ('pending', 'paid', 'failed', 'cancelled', 'refunded')),
    
    -- Dates
    issued_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    due_date DATE NOT NULL,
    paid_at TIMESTAMPTZ,
    
    -- PDF
    pdf_url TEXT,
    
    -- Timestamps
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Indexes
CREATE INDEX idx_invoices_subscription_id ON invoices(subscription_id);
CREATE INDEX idx_invoices_tenant_id ON invoices(tenant_id);
CREATE INDEX idx_invoices_status ON invoices(status);
CREATE INDEX idx_invoices_due_date ON invoices(due_date);
CREATE INDEX idx_invoices_invoice_number ON invoices(invoice_number);
```

---

### 3.3 V3 Tables (AI & Data)

#### Table: `ai_conversations`

**Purpose:** Chat history untuk AI features (RAG, assistant)

```sql
CREATE TABLE ai_conversations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Ownership
    user_id UUID NOT NULL REFERENCES users(id),
    tenant_id UUID REFERENCES tenants(id),
    
    -- Conversation details
    title VARCHAR(200),
    model VARCHAR(50) NOT NULL,  -- gpt-4, llama-3, etc
    system_prompt TEXT,
    
    -- Stats
    total_tokens INTEGER DEFAULT 0,
    total_messages INTEGER DEFAULT 0,
    
    -- Timestamps
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Indexes
CREATE INDEX idx_ai_conversations_user_id ON ai_conversations(user_id);
CREATE INDEX idx_ai_conversations_tenant_id ON ai_conversations(tenant_id);
CREATE INDEX idx_ai_conversations_created_at ON ai_conversations(created_at DESC);
```

---

#### Table: `ai_messages`

**Purpose:** Individual messages in AI conversations

```sql
CREATE TABLE ai_messages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- References
    conversation_id UUID NOT NULL REFERENCES ai_conversations(id) ON DELETE CASCADE,
    
    -- Message content
    role VARCHAR(20) NOT NULL
        CHECK (role IN ('system', 'user', 'assistant', 'tool')),
    content TEXT NOT NULL,
    
    -- Metadata
    tokens_used INTEGER,
    model VARCHAR(50),
    finish_reason VARCHAR(50),
    
    -- Embeddings (for RAG retrieval)
    embedding VECTOR(1536),  -- pgvector
    
    -- Timestamp
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Indexes
CREATE INDEX idx_ai_messages_conversation_id ON ai_messages(conversation_id);
CREATE INDEX idx_ai_messages_role ON ai_messages(role);
CREATE INDEX idx_ai_messages_created_at ON ai_messages(created_at);

-- Vector index for similarity search (RAG)
CREATE INDEX idx_ai_messages_embedding ON ai_messages 
    USING ivfflat (embedding vector_cosine_ops) WITH (lists = 100);
```

---

#### Table: `embeddings`

**Purpose:** Vector embeddings untuk knowledge base (RAG)

```sql
CREATE TABLE embeddings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Content reference
    content_type VARCHAR(50) NOT NULL,  -- document, faq, product, etc
    content_id UUID NOT NULL,
    
    -- Chunk data
    chunk_text TEXT NOT NULL,
    chunk_index INTEGER NOT NULL,
    
    -- Embedding vector
    embedding VECTOR(1536) NOT NULL,  -- OpenAI ada-002 dimension
    
    -- Metadata
    metadata JSONB DEFAULT '{}',
    
    -- Timestamp
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Indexes
CREATE INDEX idx_embeddings_content ON embeddings(content_type, content_id);
CREATE INDEX idx_embeddings_chunk_index ON embeddings(chunk_index);

-- Vector index for similarity search
CREATE INDEX idx_embeddings_vector ON embeddings 
    USING ivfflat (embedding vector_cosine_ops) WITH (lists = 100);
```

---

## 4. Database Extensions

```sql
-- Enable required extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";        -- UUID generation
CREATE EXTENSION IF NOT EXISTS "pgcrypto";         -- Cryptographic functions
CREATE EXTENSION IF NOT EXISTS "pg_trgm";          -- Trigram for fuzzy search
CREATE EXTENSION IF NOT EXISTS "vector";           -- pgvector for AI embeddings
CREATE EXTENSION IF NOT EXISTS "pg_stat_statements"; -- Query performance monitoring
```

---

## 5. Helper Functions

### 5.1 Auto-update `updated_at`

```sql
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Apply to all tables with updated_at
CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_tenants_updated_at
    BEFORE UPDATE ON tenants
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- ... (repeat for other tables)
```

### 5.2 Soft Delete Function

```sql
CREATE OR REPLACE FUNCTION soft_delete()
RETURNS TRIGGER AS $$
BEGIN
    NEW.deleted_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
```

### 5.3 Set Current Tenant (for RLS)

```sql
CREATE OR REPLACE FUNCTION set_current_tenant(tenant_uuid UUID)
RETURNS VOID AS $$
BEGIN
    PERFORM set_config('app.current_tenant_id', tenant_uuid::TEXT, TRUE);
END;
$$ LANGUAGE plpgsql;
```

---

## 6. VPS Resource Planning (PostgreSQL di VPS 1GB)

### 6.1 Memory Budget dengan PostgreSQL

```
Total RAM: 1024 MB
├── OS + System Services:        ~300 MB
├── Nginx:                        ~5 MB
├── Golang API:                  ~20-50 MB
├── PostgreSQL:                  ~200-300 MB (tuned)
│   ├── Shared buffers: 128 MB
│   ├── Work mem: 4 MB
│   ├── Maintenance work mem: 16 MB
│   └── Connections: ~50 MB
├── Swap (safety net):           1024 MB (disk)
└── AVAILABLE:                   ~400-500 MB ✅
```

### 6.2 PostgreSQL Tuning untuk VPS 1GB

```ini
# /etc/postgresql/16/main/postgresql.conf

# Memory
shared_buffers = 128MB           # 25% of RAM
effective_cache_size = 384MB     # 50-75% of RAM
work_mem = 4MB                   # Per query
maintenance_work_mem = 16MB      # For VACUUM, CREATE INDEX

# Connections
max_connections = 50             # Keep low for VPS
superuser_reserved_connections = 3

# WAL
wal_buffers = 4MB
min_wal_size = 80MB
max_wal_size = 1GB

# Query Planner
random_page_cost = 1.1           # SSD optimization
effective_io_concurrency = 200   # SSD optimization

# Logging
log_min_duration_statement = 1000  # Log slow queries (>1s)
log_checkpoints = on
log_connections = on
log_disconnections = on

# Performance
checkpoint_completion_target = 0.9
default_statistics_target = 100
```

### 6.3 Connection Pooling (Optional)

Jika koneksi banyak, gunakan **PgBouncer**:

```ini
# /etc/pgbouncer/pgbouncer.ini
[databases]
webifylab = host=127.0.0.1 port=5432 dbname=webifylab

[pgbouncer]
listen_addr = 127.0.0.1
listen_port = 6432
auth_type = md5
auth_file = /etc/pgbouncer/userlist.txt
pool_mode = transaction
max_client_conn = 100
default_pool_size = 20
```

---

## 7. Migration Strategy

### 7.1 Migration Tool: Golang Migrate

```bash
# Install golang-migrate
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Create migrations
migrate create -ext sql -dir db/migrations -seq init_schema
migrate create -ext sql -dir db/migrations -seq add_contact_submissions
migrate create -ext sql -dir db/migrations -seq add_analytics_events
```

### 7.2 Migration Folder Structure

```
apps/api/
└── db/
    ├── migrations/
    │   ├── 000001_init_schema.up.sql
    │   ├── 000001_init_schema.down.sql
    │   ├── 000002_add_contact_submissions.up.sql
    │   ├── 000002_add_contact_submissions.down.sql
    │   ├── 000003_add_analytics_events.up.sql
    │   ├── 000003_add_analytics_events.down.sql
    │   └── ...
    └── seeds/
        ├── seed_users.sql
        └── seed_portfolios.sql
```

### 7.3 Migration Commands

```bash
# Run migrations
migrate -path db/migrations -database "postgres://user:pass@localhost:5432/webifylab?sslmode=disable" up

# Rollback last migration
migrate -path db/migrations -database "postgres://..." down 1

# Check status
migrate -path db/migrations -database "postgres://..." status
```

---

## 8. Golang Integration

### 8.1 Database Connection (Golang)

```go
// internal/config/database.go
package config

import (
    "fmt"
    "os"
    "github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
    SSLMode  string
}

func LoadDatabaseConfig() *DatabaseConfig {
    return &DatabaseConfig{
        Host:     getEnv("DB_HOST", "localhost"),
        Port:     getEnv("DB_PORT", "5432"),
        User:     getEnv("DB_USER", "webifylab"),
        Password: getEnv("DB_PASSWORD", ""),
        DBName:   getEnv("DB_NAME", "webifylab"),
        SSLMode:  getEnv("DB_SSLMODE", "disable"),
    }
}

func (c *DatabaseConfig) ConnectionString() string {
    return fmt.Sprintf(
        "postgres://%s:%s@%s:%s/%s?sslmode=%s",
        c.User, c.Password, c.Host, c.Port, c.DBName, c.SSLMode,
    )
}

func NewDatabasePool(cfg *DatabaseConfig) (*pgxpool.Pool, error) {
    poolConfig, err := pgxpool.ParseConfig(cfg.ConnectionString())
    if err != nil {
        return nil, err
    }
    
    // Connection pool settings
    poolConfig.MaxConns = 20
    poolConfig.MinConns = 5
    
    return pgxpool.NewWithConfig(context.Background(), poolConfig)
}
```

### 8.2 Repository Pattern

```go
// internal/repositories/contact_repository.go
package repositories

import (
    "context"
    "github.com/jackc/pgx/v5/pgxpool"
    "webifylab/api/internal/models"
)

type ContactRepository struct {
    pool *pgxpool.Pool
}

func NewContactRepository(pool *pgxpool.Pool) *ContactRepository {
    return &ContactRepository{pool: pool}
}

func (r *ContactRepository) Create(ctx context.Context, submission *models.ContactSubmission) error {
    query := `
        INSERT INTO contact_submissions (name, email, service_type, message, source, ip_address, user_agent)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id, created_at
    `
    
    return r.pool.QueryRow(ctx, query,
        submission.Name,
        submission.Email,
        submission.ServiceType,
        submission.Message,
        submission.Source,
        submission.IPAddress,
        submission.UserAgent,
    ).Scan(&submission.ID, &submission.CreatedAt)
}

func (r *ContactRepository) GetAll(ctx context.Context, limit, offset int) ([]models.ContactSubmission, error) {
    query := `
        SELECT id, name, email, service_type, message, status, created_at
        FROM contact_submissions
        WHERE deleted_at IS NULL
        ORDER BY created_at DESC
        LIMIT $1 OFFSET $2
    `
    
    rows, err := r.pool.Query(ctx, query, limit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var submissions []models.ContactSubmission
    for rows.Next() {
        var s models.ContactSubmission
        if err := rows.Scan(&s.ID, &s.Name, &s.Email, &s.ServiceType, &s.Message, &s.Status, &s.CreatedAt); err != nil {
            return nil, err
        }
        submissions = append(submissions, s)
    }
    
    return submissions, nil
}
```

---

## 9. Backup & Recovery

### 9.1 Automated Backup Script

```bash
#!/bin/bash
# /opt/scripts/backup-postgres.sh

BACKUP_DIR="/backup/postgres"
DATE=$(date +%Y%m%d_%H%M%S)
DB_NAME="webifylab"
RETENTION_DAYS=7

mkdir -p $BACKUP_DIR

# Dump database
pg_dump -U webifylab -h localhost -F c -b -v -f "$BACKUP_DIR/${DB_NAME}_${DATE}.backup" $DB_NAME

# Compress
gzip "$BACKUP_DIR/${DB_NAME}_${DATE}.backup"

# Delete old backups
find $BACKUP_DIR -type f -mtime +$RETENTION_DAYS -delete

echo "✅ Backup completed: ${DB_NAME}_${DATE}.backup.gz"
```

### 9.2 Cron Job

```bash
# Daily backup at 2 AM
0 2 * * * /opt/scripts/backup-postgres.sh >> /var/log/postgres-backup.log 2>&1
```

### 9.3 Restore Command

```bash
# Restore from backup
pg_restore -U webifylab -h localhost -d webifylab -v /backup/postgres/webifylab_20260917_020000.backup.gz
```

---

## 10. Security Best Practices

### 10.1 Database User Roles

```sql
-- Create application user (limited privileges)
CREATE USER webifylab_app WITH PASSWORD 'strong_password_here';

-- Grant only necessary permissions
GRANT CONNECT ON DATABASE webifylab TO webifylab_app;
GRANT USAGE ON SCHEMA public TO webifylab_app;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO webifylab_app;
GRANT USAGE ON ALL SEQUENCES IN SCHEMA public TO webifylab_app;

-- Create read-only user (for analytics)
CREATE USER webifylab_readonly WITH PASSWORD 'readonly_password';
GRANT CONNECT ON DATABASE webifylab TO webifylab_readonly;
GRANT USAGE ON SCHEMA public TO webifylab_readonly;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO webifylab_readonly;

-- Create backup user
CREATE USER webifylab_backup WITH PASSWORD 'backup_password';
GRANT CONNECT ON DATABASE webifylab TO webifylab_backup;
GRANT USAGE ON SCHEMA public TO webifylab_backup;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO webifylab_backup;
```

### 10.2 Firewall Rules

```bash
# Only allow PostgreSQL from localhost
sudo ufw deny 5432
sudo ufw allow from 127.0.0.1 to any port 5432
```

### 10.3 SSL Connection

```ini
# postgresql.conf
ssl = on
ssl_cert_file = '/etc/ssl/certs/postgres.crt'
ssl_key_file = '/etc/ssl/private/postgres.key'
```

---

## 11. Implementation Roadmap

### Phase 1: V1.5 (Month 2-3)

- [ ] Install PostgreSQL 16 di VPS
- [ ] Tuning konfigurasi untuk VPS 1GB
- [ ] Create database & users
- [ ] Run migrations: `contact_submissions`, `analytics_events`
- [ ] Integrate dengan Golang API
- [ ] Replace Formspree dengan custom form handler
- [ ] Setup backup cron job

### Phase 2: V2 (Month 6+)

- [ ] Run migrations: `users`, `tenants`, `projects`, `portfolios`
- [ ] Implement authentication (JWT)
- [ ] Setup Row-Level Security (RLS)
- [ ] Build admin dashboard
- [ ] Integrate payment gateway (Midtrans/Xendit)
- [ ] Run migrations: `subscriptions`, `invoices`

### Phase 3: V3 (Year 2+)

- [ ] Enable pgvector extension
- [ ] Run migrations: `ai_conversations`, `ai_messages`, `embeddings`
- [ ] Build RAG system
- [ ] Integrate LLM APIs (OpenAI, local models)
- [ ] Build vector search features

---

## 12. Database Schema Checklist

- [ ] All tables have UUID primary keys
- [ ] All tables have `created_at`, `updated_at` timestamps
- [ ] Soft delete implemented (`deleted_at`)
- [ ] Indexes created for frequently queried columns
- [ ] Foreign key constraints defined
- [ ] CHECK constraints for enum-like columns
- [ ] Full-text search indexes for searchable content
- [ ] Vector indexes for AI embeddings (V3)
- [ ] Row-Level Security enabled (V2)
- [ ] Database users with appropriate privileges
- [ ] Backup script created and tested
- [ ] Migration tool configured
- [ ] Connection pooling configured (if needed)
- [ ] SSL enabled for connections
- [ ] Slow query logging enabled

---

## 13. Open Questions

| No | Pertanyaan | Status |
|----|-----------|--------|
| Q1 | Apakah PostgreSQL adalah pilihan yang tepat, atau preferensi SQLite untuk V1.5? | Pending |
| Q2 | Apakah perlu multi-tenant architecture di V2, atau single-tenant dulu? | Pending |
| Q3 | Apakah ada preferensi payment gateway (Midtrans, Xendit, Stripe)? | Pending |
| Q4 | Apakah ingin self-hosted AI models atau pakai API (OpenAI, Anthropic)? | Pending |
| Q5 | Apakah perlu database monitoring tool (pgAdmin, Metabase)? | Pending |

---

*Dokumen ini adalah living document. Versi akan diperbarui seiring perkembangan ekosistem Webifylab.*

**Last Updated:** 17 September 2026
**Next Step:** Setup PostgreSQL di VPS, run migrations V1.5, integrate dengan Golang API.