package main

import (
	"fmt"
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User struct với các trường ID và Name
type User struct {
	gorm.Model
	ID   uint
	Name string
}

// Hàm insertRecord: thực hiện việc chèn một bản ghi User vào DB
func insertRecord(b *testing.B, db *gorm.DB) {
	user := User{Name: "Tipjs"}

	// Thực hiện tạo bản ghi trong DB và kiểm tra lỗi
	if err := db.Create(&user).Error; err != nil {
		b.Fatal(err)
	}
}

func BenchmarkMaxOpenConns1(b *testing.B) {
	// Chuỗi kết nối MySQL
	dsn := "root:root@tcp(127.0.0.1:3307)/shopdevgo?charset=utf8mb4&parseTime=True"

	// Kết nối đến cơ sở dữ liệu với GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Kiểm tra xem bảng User đã tồn tại chưa
	if db.Migrator().HasTable(&User{}) {
		// Xóa bảng nếu đã tồn tại
		if err := db.Migrator().DropTable(&User{}); err != nil {
			// Xử lý lỗi nếu có
			fmt.Println("Error dropping table:", err)
		}
	}

	// Tạo bảng User nếu chưa có
	db.AutoMigrate(&User{})

	// Lấy đối tượng sql.DB từ GORM để cấu hình kết nối
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB from gorm.DB: %v", err)
	}

	// Thiết lập các tham số kết nối
	sqlDB.SetMaxOpenConns(1) // Giới hạn số kết nối mở tối đa là 1
	defer sqlDB.Close()       // Đóng kết nối khi không sử dụng

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})	
}

// go test -bench=. -benchmem

func BenchmarkMaxOpenConns10(b *testing.B) {
	// Chuỗi kết nối MySQL
	dsn := "root:root@tcp(127.0.0.1:3307)/shopdevgo?charset=utf8mb4&parseTime=True&loc=Local"

	// Kết nối đến cơ sở dữ liệu bằng GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		// Logger: false, // Có thể bật logger nếu cần
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Kiểm tra xem bảng User đã tồn tại chưa
	if db.Migrator().HasTable(&User{}) {
		// Xóa bảng nếu đã tồn tại
		if err := db.Migrator().DropTable(&User{}); err != nil {
			// Xử lý lỗi nếu có
			fmt.Println("Error dropping table:", err)
		}
	}

	// Tạo bảng nếu chưa tồn tại
	db.AutoMigrate(&User{})

	// Lấy sql.DB từ GORM để cấu hình kết nối
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB from gorm.DB: %v", err)
	}

	// Thiết lập số lượng kết nối tối đa
	sqlDB.SetMaxOpenConns(10) // Giới hạn tối đa 10 kết nối mở
	defer sqlDB.Close()        // Đóng kết nối khi không sử dụng

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}