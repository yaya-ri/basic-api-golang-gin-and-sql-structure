package model

//basic struct type for model example

// type ProductModel struct {
// 	ID          uint      `gorm:"primary_key;column:id" json:"id"`
// 	Name        string    `gorm:"column:name" json:"name"`
// 	Description string    `gorm:"column:description" json:"description"`
// 	Enable      *bool     `gorm:"column:enable" json:"enable"`
// 	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
// 	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
// }

// type Category struct {
// 	ID        uint       `gorm:"primary_key;column:id" json:"id"`
// 	Name      string     `gorm:"column:name" json:"name"`
// 	Enable    *bool      `gorm:"column:enable" json:"enable"`
// 	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
// 	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
// 	DeletedAt *time.Time `gorm:"deleted_at" json:"-"`
// }

// type Image struct {
// 	ID        uint       `gorm:"primary_key;column:id" json:"id"`
// 	Name      string     `gorm:"column:name" json:"name"`
// 	File      string     `gorm:"column:file" json:"file"`
// 	Enable    *bool      `gorm:"column:enable" json:"enable"`
// 	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
// 	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
// 	DeletedAt *time.Time `gorm:"deleted_at" json:"-"`
// }

// type CategoryProduct struct {
// 	ID         uint      `gorm:"primary_key;column:id" json:"id"`
// 	ProductID  uint      `gorm:"column:product_id" json:"product_id"`
// 	CategoryID uint      `gorm:"column:category_id" json:"category_id"`
// 	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
// 	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
// }

// type CP struct {
// 	ID         uint       `gorm:"primary_key;column:id" json:"id"`
// 	ProductID  uint       `gorm:"column:product_id" json:"product_id"`
// 	CategoryID uint       `gorm:"column:category_id" json:"category_id"`
// 	CreatedAt  time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
// 	UpdatedAt  time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
// 	DeletedAt  *time.Time `gorm:"deleted_at" json:"-"`
// }

// type ProductImage struct {
// 	ID        uint      `gorm:"primary_key;column:id" json:"id"`
// 	ProductID uint      `gorm:"column:product_id" json:"product_id"`
// 	ImageID   uint      `gorm:"column:image_id" json:"image_id"`
// 	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
// 	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
// }

// type Product struct {
// 	ID          uint       `gorm:"primary_key;column:id" json:"id"`
// 	Name        string     `gorm:"column:name" json:"name"`
// 	Description string     `gorm:"column:description" json:"description"`
// 	Enable      *bool      `gorm:"column:enable" json:"enable"`
// 	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
// 	UpdatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
// 	Category    []Category `gorm:"many2many:category_products;" json:"category,omitempty"`
// 	Image       []Image    `gorm:"many2many:product_images;" json:"image,omitempty"`
// 	DeletedAt   *time.Time `gorm:"deleted_at" json:"-"`
// }

// type Response struct {
// 	Success bool        `json:"success"`
// 	Message string      `json:"message"`
// 	Data    interface{} `json:"data,omitempty"`
// }
