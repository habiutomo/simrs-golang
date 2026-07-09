package database

import (
	"fmt"
	"log"

	"simrs-golang/config"
	"simrs-golang/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(cfg *config.DatabaseConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully")
}

func Migrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Department{},
		&models.Doctor{},
		&models.DoctorSchedule{},
		&models.Patient{},
		&models.Appointment{},
		&models.MedicalRecord{},
		&models.MedicalPrescription{},
		&models.Medication{},
		&models.Billing{},
		&models.BillingItem{},
		&models.Room{},
		&models.Inpatient{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	seedData()
	log.Println("Database migrated successfully")
}

func seedData() {
	var count int64
	DB.Model(&models.Department{}).Count(&count)
	if count > 0 {
		return
	}

	departments := []models.Department{
		{Name: "Kardiologi", Code: "KAR", Description: "Jantung & Pembuluh Darah"},
		{Name: "Pediatri", Code: "PED", Description: "Kesehatan Anak"},
		{Name: "Ortopedi", Code: "ORT", Description: "Tulang & Sendi"},
		{Name: "Neurologi", Code: "NEU", Description: "Syaraf"},
		{Name: "Obstetri & Ginekologi", Code: "OBG", Description: "Kandungan"},
		{Name: "Mata", Code: "MAT", Description: "Kesehatan Mata"},
		{Name: "THT", Code: "THT", Description: "Telinga, Hidung, Tenggorokan"},
		{Name: "Umum", Code: "UMU", Description: "Poli Umum"},
	}
	DB.Create(&departments)

	medications := []models.Medication{
		{Name: "Paracetamol 500mg", Category: "Analgesik", Stock: 500, Price: 5000, Unit: "tablet"},
		{Name: "Amoxicillin 500mg", Category: "Antibiotik", Stock: 300, Price: 12000, Unit: "kapsul"},
		{Name: "Ibuprofen 400mg", Category: "Anti-inflamasi", Stock: 400, Price: 8000, Unit: "tablet"},
		{Name: "Omeprazole 20mg", Category: "Obat Lambung", Stock: 200, Price: 15000, Unit: "kapsul"},
		{Name: "Ceftriaxone 1g", Category: "Antibiotik", Stock: 100, Price: 45000, Unit: "vial"},
		{Name: "Dextrose 5%", Category: "Cairan Infus", Stock: 150, Price: 25000, Unit: "botol"},
		{Name: "Ranitidine 50mg", Category: "Antasida", Stock: 250, Price: 3000, Unit: "ampul"},
		{Name: "Diazepam 5mg", Category: "Sedatif", Stock: 100, Price: 10000, Unit: "tablet"},
	}
	DB.Create(&medications)

	rooms := []models.Room{
		{Number: "101", Type: "VIP", Class: "VIP", Capacity: 1, PricePerDay: 500000, Status: "available"},
		{Number: "102", Type: "VIP", Class: "VIP", Capacity: 1, PricePerDay: 500000, Status: "available"},
		{Number: "201", Type: "Kelas 1", Class: "1", Capacity: 2, PricePerDay: 250000, Status: "available"},
		{Number: "202", Type: "Kelas 1", Class: "1", Capacity: 2, PricePerDay: 250000, Status: "available"},
		{Number: "301", Type: "Kelas 2", Class: "2", Capacity: 4, PricePerDay: 150000, Status: "available"},
		{Number: "302", Type: "Kelas 2", Class: "2", Capacity: 4, PricePerDay: 150000, Status: "available"},
		{Number: "401", Type: "Kelas 3", Class: "3", Capacity: 6, PricePerDay: 75000, Status: "available"},
		{Number: "402", Type: "Kelas 3", Class: "3", Capacity: 6, PricePerDay: 75000, Status: "available"},
	}
	DB.Create(&rooms)

	hasher, _ := NewPassword("admin123")
	admin := models.User{
		Username: "admin", Password: hasher,
		FullName: "Administrator", Email: "admin@simrs.com",
		Role: "admin", Phone: "081234567890",
	}
	DB.Create(&admin)

	hasher2, _ := NewPassword("doctor123")
	doctorUser := models.User{
		Username: "dr.andini", Password: hasher2,
		FullName: "dr. Andini Putri", Email: "andini@simrs.com",
		Role: "doctor", Phone: "081234567891",
	}
	DB.Create(&doctorUser)

	hasher3, _ := NewPassword("doctor123")
	doctorUser2 := models.User{
		Username: "dr.bambang", Password: hasher3,
		FullName: "dr. Bambang Susilo", Email: "bambang@simrs.com",
		Role: "doctor", Phone: "081234567892",
	}
	DB.Create(&doctorUser2)

	var dept1, dept2 models.Department
	DB.Where("code = ?", "KAR").First(&dept1)
	DB.Where("code = ?", "UMU").First(&dept2)

	doctors := []models.Doctor{
		{UserID: doctorUser.ID, DepartmentID: dept1.ID, LicenseNumber: "SIP-2024-001", Specialization: "Kardiologi Intervensi"},
		{UserID: doctorUser2.ID, DepartmentID: dept2.ID, LicenseNumber: "SIP-2024-002", Specialization: "Dokter Umum"},
	}
	DB.Create(&doctors)

	patients := []models.Patient{
		{MedicalRecordNumber: "RM-20240001", Name: "Siti Rahmawati", NIK: "3174015208900001", PlaceOfBirth: "Jakarta", DateOfBirth: "1980-08-15", Gender: "female", Address: "Jl. Merdeka No. 1", Phone: "081298765432", BloodType: "A", Allergies: "Paracetamol"},
		{MedicalRecordNumber: "RM-20240002", Name: "Ahmad Fauzi", NIK: "3174011203850002", PlaceOfBirth: "Bandung", DateOfBirth: "1975-03-12", Gender: "male", Address: "Jl. Sudirman No. 5", Phone: "081298765433", BloodType: "B", Allergies: ""},
		{MedicalRecordNumber: "RM-20240003", Name: "Dewi Sartika", NIK: "3174016708900003", PlaceOfBirth: "Surabaya", DateOfBirth: "1990-08-27", Gender: "female", Address: "Jl. Thamrin No. 10", Phone: "081298765434", BloodType: "O", Allergies: "Seafood"},
	}
	DB.Create(&patients)
}

func NewPassword(password string) (string, error) {
	bytes, err := HashPassword(password)
	return string(bytes), err
}

var _ = NewPassword
