package models

// ============================================
// User Roles
// ============================================
const (
	RoleSuperAdmin = "super_admin"
	RoleAdmin      = "admin"
	RoleEditor     = "editor"
)

// ============================================
// Content Status (Blog, Case Study, Portfolio)
// ============================================
const (
	StatusDraft     = "draft"
	StatusPublished = "published"
	StatusArchived  = "archived"
)

// ============================================
// Lead Status (Pipeline)
// ============================================
const (
	LeadStatusNew          = "new"
	LeadStatusContacted    = "contacted"
	LeadStatusNegotiation  = "negotiation"
	LeadStatusProposalSent = "proposal_sent"
	LeadStatusWon          = "won"
	LeadStatusLost         = "lost"
)

// ============================================
// Service Types (untuk Lead)
// ============================================
const (
	ServiceWebDevelopment = "web_development"
	ServiceMobileApp      = "mobile_app"
	ServiceUIUXDesign     = "uiux_design"
	ServiceAIIntegration  = "ai_integration"
	ServiceOther          = "other"
)

// ============================================
// Budget Ranges
// ============================================
const (
	BudgetUnder5M   = "<5jt"
	Budget5to15M    = "5-15jt"
	Budget15to50M   = "15-50jt"
	Budget50to100M  = "50-100jt"
	BudgetOver100M  = ">100jt"
)

// ============================================
// AI Types (untuk Case Study)
// ============================================
const (
	AISPK             = "spk"
	AISistemPakar     = "sistem_pakar"
	AIDataAnalytics   = "data_analytics"
	AIMachineLearning = "machine_learning"
)

// ============================================
// Portfolio Categories
// ============================================
const (
	CategoryWeb       = "web"
	CategoryMobile    = "mobile"
	CategoryAI        = "ai"
	CategoryEcommerce = "ecommerce"
	CategoryOther     = "other"
)

// ============================================
// Category Types (untuk Category model)
// ============================================
const (
	CategoryTypeBlog      = "blog"
	CategoryTypePortfolio = "portfolio"
)

// ============================================
// Allowed MIME Types (untuk Media)
// ============================================
const (
	MimeTypeJPEG = "image/jpeg"
	MimeTypePNG  = "image/png"
	MimeTypeWebP = "image/webp"
)

// ============================================
// Max Upload Size (5MB in bytes)
// ============================================
const MaxUploadSize = 5242880