package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	CreatedBy uint
	UpdatedAt time.Time
	UpdatedBy uint
	DeletedAt gorm.DeletedAt `gorm:"index"`
	DeletedBy uint
}

// BeforeCreate, yeni kayıt oluşturulmadan önce tetiklenir.
func (bm *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	currentTime := time.Now()
	if bm.CreatedAt.IsZero() {
		bm.CreatedAt = currentTime
	}
	if bm.UpdatedAt.IsZero() {
		bm.UpdatedAt = currentTime
	}
	if bm.CreatedBy == 0 {
		// Burada mevcut kullanıcıyı almak veya belirlemek gerekebilir.
		// Örneğin, kimlik doğrulama ve yetkilendirme sonucu elde edilen kullanıcı ID'si kullanılabilir.
		bm.CreatedBy = 1 // Örnek olarak 1 kullanıcısı
	}
	return
}

// BeforeUpdate, kayıt güncellenmeden önce tetiklenir.
func (bm *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	bm.UpdatedAt = time.Now()
	if bm.UpdatedBy == 0 {
		// Aynı şekilde güncelleyen kullanıcıyı belirlemek gerekebilir.
		bm.UpdatedBy = 1 // Örnek olarak 1 kullanıcısı
	}
	return
}

// BeforeDelete, kayıt silinmeden önce tetiklenir.
func (bm *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	bm.DeletedBy = 1 // Örnek olarak 1 kullanıcısı
	return
}
