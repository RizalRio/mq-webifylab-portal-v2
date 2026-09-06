package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// ============================================
// 1. USER MODEL
// ============================================
type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name"`
	Email        string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	PasswordHash string    `gorm:"type:varchar(255);not null" json:"-"` // "-" = tidak di-serialize ke JSON
	Role         string    `gorm:"type:varchar(20);not null;default:'editor'" json:"role"`
	CreatedAt    time.Time `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt    time.Time `gorm:"not null;default:now()" json:"updated_at"`

	// Relationships
	BlogPosts     []BlogPost     `gorm:"foreignKey:AuthorID" json:"blog_posts,omitempty"`
	Media         []Media        `gorm:"foreignKey:UploadedBy" json:"media,omitempty"`
	RefreshTokens []RefreshToken `gorm:"foreignKey:UserID" json:"-"`
}

func (User) TableName() string {
	return "users"
}

// ============================================
// 2. REFRESH TOKEN MODEL
// ============================================
type RefreshToken struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	Token     string    `gorm:"type:varchar(500);uniqueIndex;not null" json:"-"`
	ExpiresAt time.Time `gorm:"not null;index" json:"expires_at"`
	CreatedAt time.Time `gorm:"not null;default:now()" json:"created_at"`

	// Relationships
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
}

func (RefreshToken) TableName() string {
	return "refresh_tokens"
}

// ============================================
// 3. PASSWORD RESET MODEL
// ============================================
type PasswordReset struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	Token     string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"-"`
	ExpiresAt time.Time  `gorm:"not null" json:"expires_at"`
	UsedAt    *time.Time `json:"used_at"`
	CreatedAt time.Time  `gorm:"not null;default:now()" json:"created_at"`

	// Relationships
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
}

func (PasswordReset) TableName() string {
	return "password_resets"
}

// ============================================
// 4. CATEGORY MODEL
// ============================================
type Category struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Slug      string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"slug"`
	Type      string    `gorm:"type:varchar(20);not null;index" json:"type"` // blog or portfolio
	CreatedAt time.Time `gorm:"not null;default:now()" json:"created_at"`

	// Relationships
	BlogPosts []BlogPost `gorm:"foreignKey:CategoryID" json:"blog_posts,omitempty"`
}

func (Category) TableName() string {
	return "categories"
}

// ============================================
// 5. TAG MODEL
// ============================================
type Tag struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"`
	Slug      string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"slug"`
	CreatedAt time.Time `gorm:"not null;default:now()" json:"created_at"`

	// Relationships (Many-to-Many)
	BlogPosts   []BlogPost  `gorm:"many2many:post_tags;" json:"blog_posts,omitempty"`
	CaseStudies []CaseStudy `gorm:"many2many:case_study_tags;" json:"case_studies,omitempty"`
}

func (Tag) TableName() string {
	return "tags"
}

// ============================================
// 6. BLOG POST MODEL
// ============================================
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
	Author   User      `gorm:"foreignKey:AuthorID;constraint:OnDelete:RESTRICT" json:"author"`
	Category *Category `gorm:"foreignKey:CategoryID;constraint:OnDelete:SET NULL" json:"category"`
	Tags     []Tag     `gorm:"many2many:post_tags;" json:"tags"`
}

func (BlogPost) TableName() string {
	return "blog_posts"
}

// ============================================
// 7. CASE STUDY MODEL
// ============================================
type CaseStudy struct {
	ID              uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Title           string         `gorm:"type:varchar(255);not null" json:"title"`
	Slug            string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"slug"`
	ClientName      string         `gorm:"type:varchar(100);not null" json:"client_name"`
	Industry        string         `gorm:"type:varchar(100);not null;index" json:"industry"`
	Challenge       string         `gorm:"type:text;not null" json:"challenge"`
	Approach        string         `gorm:"type:text;not null" json:"approach"`
	Solution        string         `gorm:"type:text;not null" json:"solution"`
	AIType          datatypes.JSON `gorm:"type:jsonb" json:"ai_type"` // Array of AI types
	Results         string         `gorm:"type:text;not null" json:"results"`
	Lessons         *string        `gorm:"type:text" json:"lessons"`
	Testimonial     *string        `gorm:"type:text" json:"testimonial"`
	Metrics         datatypes.JSON `gorm:"type:jsonb" json:"metrics"` // Flexible JSON
	CoverImage      *string        `gorm:"type:varchar(500)" json:"cover_image"`
	Images          datatypes.JSON `gorm:"type:jsonb" json:"images"` // Array of image URLs
	Status          string         `gorm:"type:varchar(20);not null;default:'draft';index" json:"status"`
	MetaTitle       *string        `gorm:"type:varchar(255)" json:"meta_title"`
	MetaDescription *string        `gorm:"type:varchar(500)" json:"meta_description"`
	PublishedAt     *time.Time     `gorm:"index" json:"published_at"`
	CreatedAt       time.Time      `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"not null;default:now()" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Tags []Tag `gorm:"many2many:case_study_tags;" json:"tags"`
}

func (CaseStudy) TableName() string {
	return "case_studies"
}

// ============================================
// 8. PORTFOLIO MODEL
// ============================================
type Portfolio struct {
	ID              uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Title           string         `gorm:"type:varchar(255);not null" json:"title"`
	Slug            string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"slug"`
	Description     string         `gorm:"type:text;not null" json:"description"`
	Category        string         `gorm:"type:varchar(50);not null;index" json:"category"` // web, mobile, ai, ecommerce, other
	TechStack       datatypes.JSON `gorm:"type:jsonb" json:"tech_stack"`                    // Array of technologies
	LiveURL         *string        `gorm:"type:varchar(500)" json:"live_url"`
	Thumbnail       *string        `gorm:"type:varchar(500)" json:"thumbnail"`
	Images          datatypes.JSON `gorm:"type:jsonb" json:"images"` // Array of image URLs
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

// ============================================
// 9. SERVICE MODEL
// ============================================
type Service struct {
	ID               uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name             string         `gorm:"type:varchar(100);not null" json:"name"`
	Slug             string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"slug"`
	Description      string         `gorm:"type:text;not null" json:"description"`
	Icon             *string        `gorm:"type:varchar(50)" json:"icon"`
	UseCases         datatypes.JSON `gorm:"type:jsonb" json:"use_cases"` // Array of use cases
	TimelineEstimate *string        `gorm:"type:varchar(100)" json:"timeline_estimate"`
	TechStack        datatypes.JSON `gorm:"type:jsonb" json:"tech_stack"` // Array of technologies
	LearningPoints   *string        `gorm:"type:text" json:"learning_points"`
	IsActive         bool           `gorm:"not null;default:true;index" json:"is_active"`
	OrderIndex       int            `gorm:"not null;default:0;index" json:"order_index"`
	CreatedAt        time.Time      `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"not null;default:now()" json:"updated_at"`
}

func (Service) TableName() string {
	return "services"
}

// ============================================
// 10. LEAD MODEL
// ============================================
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

// ============================================
// 11. MEDIA MODEL
// ============================================
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
	Uploader User `gorm:"foreignKey:UploadedBy;constraint:OnDelete:RESTRICT" json:"uploader"`
}

func (Media) TableName() string {
	return "media"
}

// ============================================
// 12. SITE SETTING MODEL
// ============================================
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

// ============================================
// JUNCTION TABLES (Many-to-Many)
// ============================================

// PostTag - Junction table untuk BlogPost <-> Tag
type PostTag struct {
	PostID uuid.UUID `gorm:"type:uuid;primaryKey" json:"post_id"`
	TagID  uuid.UUID `gorm:"type:uuid;primaryKey" json:"tag_id"`
}

func (PostTag) TableName() string {
	return "post_tags"
}

// CaseStudyTag - Junction table untuk CaseStudy <-> Tag
type CaseStudyTag struct {
	CaseStudyID uuid.UUID `gorm:"type:uuid;primaryKey" json:"case_study_id"`
	TagID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"tag_id"`
}

func (CaseStudyTag) TableName() string {
	return "case_study_tags"
}

// ============================================
// AUTO MIGRATE HELPER
// ============================================

// AllModels returns all models for auto-migration
// Urutan penting karena foreign key dependencies
func AllModels() []interface{} {
	return []interface{}{
		&User{},
		&RefreshToken{},
		&PasswordReset{},
		&Category{},
		&Tag{},
		&BlogPost{},
		&CaseStudy{},
		&Portfolio{},
		&Service{},
		&Lead{},
		&Media{},
		&SiteSetting{},
		&PostTag{},
		&CaseStudyTag{},
	}
}