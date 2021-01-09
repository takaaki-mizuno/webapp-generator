package repositories

import (
	"time"

	"github.com/omiselabs/gin-boilerplate/internal/models"
	"gorm.io/gorm"
)

// Repository ...
type Repository interface {
	Find(scopes ...func(*gorm.DB) *gorm.DB) (*models.Base, error)
}

// ByID ...
func ByID(id uint64) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}

// ByNotificationID ...
func ByNotificationID(notificationID uint64) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("notification_id = ?", notificationID)
	}
}

// ByUserID ...
func ByUserID(userID uint64) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", userID)
	}
}

// ByIDs ...
func ByIDs(ids *[]uint64) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(*ids) == 0 {
			return db
		}
		return db.Where("id IN (?)", *ids)
	}
}

// OrderBy ...
func OrderBy(order string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	}
}

// Limit ...
func Limit(limit int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit)
	}
}

// Offset ...
func Offset(offset int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset)
	}
}

// ByIsEnabled ...
func ByIsEnabled(isEnabled bool) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("is_enabled = ?", isEnabled)
	}
}

// WithinCreatedAt ...
func WithinCreatedAt(seconds int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("created_at > ?", time.Now().Add(time.Duration(-seconds)*time.Second))
	}
}

// WithinUpdatedAt ...
func WithinUpdatedAt(seconds int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("updated_at > ?", time.Now().Add(time.Duration(-seconds)*time.Second))
	}
}

// AfterUpdatedAt ...
func AfterUpdatedAt(seconds int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("updated_at > ?", time.Unix(int64(seconds), 0))
	}
}

// Preload ... preload related models
func Preload(relationName string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(relationName)
	}
}

// PreloadWith ... preload related models with conditions
func PreloadWith(relationName string, conditions string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(relationName, conditions)
	}
}

// JoinPreload ... join preload related models
// Note: JoinPreload works with one-to-one relation. In other cases, use preload.
func JoinPreload(relationName string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Joins(relationName)
	}
}
