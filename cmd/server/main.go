package main

import (
	"log"

	"simrs-golang/config"
	"simrs-golang/internal/database"
	"simrs-golang/internal/handlers"
	"simrs-golang/internal/middleware"
	"simrs-golang/internal/repositories"
	"simrs-golang/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	database.Connect(&cfg.Database)
	database.Migrate()

	middleware.InitAuth(&cfg.JWT)

	// Repositories
	userRepo := repositories.NewUserRepository()
	patientRepo := repositories.NewPatientRepository()
	doctorRepo := repositories.NewDoctorRepository()
	appointmentRepo := repositories.NewAppointmentRepository()
	medicalRecordRepo := repositories.NewMedicalRecordRepository()
	medicationRepo := repositories.NewMedicationRepository()
	billingRepo := repositories.NewBillingRepository()
	departmentRepo := repositories.NewDepartmentRepository()
	roomRepo := repositories.NewRoomRepository()

	// Services
	authService := services.NewAuthService(userRepo)
	patientService := services.NewPatientService(patientRepo)
	doctorService := services.NewDoctorService(doctorRepo, userRepo, departmentRepo)
	appointmentService := services.NewAppointmentService(appointmentRepo)
	medicalRecordService := services.NewMedicalRecordService(medicalRecordRepo, appointmentRepo, medicationRepo)
	medicationService := services.NewMedicationService(medicationRepo)
	billingService := services.NewBillingService(billingRepo, patientRepo)
	roomService := services.NewRoomService(roomRepo, patientRepo, doctorRepo)
	dashboardService := services.NewDashboardService(patientRepo, doctorRepo, appointmentRepo, medicalRecordRepo, billingRepo, roomRepo)

	// Handlers
	authHandler := handlers.NewAuthHandler(authService)
	patientHandler := handlers.NewPatientHandler(patientService)
	doctorHandler := handlers.NewDoctorHandler(doctorService)
	appointmentHandler := handlers.NewAppointmentHandler(appointmentService)
	medicalRecordHandler := handlers.NewMedicalRecordHandler(medicalRecordService)
	medicationHandler := handlers.NewMedicationHandler(medicationService)
	billingHandler := handlers.NewBillingHandler(billingService)
	roomHandler := handlers.NewRoomHandler(roomService)
	departmentHandler := handlers.NewDepartmentHandler(departmentRepo)
	dashboardHandler := handlers.NewDashboardHandler(dashboardService)

	router := gin.Default()
	gin.SetMode(cfg.Server.Mode)
	router.Use(middleware.CORS())

	// Health Check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "SIMRS - Smart Hospital Information System",
			"version": "1.0.0",
		})
	})

	// Public routes
	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/login", authHandler.Login)
	}

	// Protected routes
	api := router.Group("/api/v1")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/profile", authHandler.GetProfile)

		// Departments
		api.GET("/departments", departmentHandler.GetAll)

		// Dashboard
		api.GET("/dashboard", dashboardHandler.GetDashboard)

		// Patients
		patients := api.Group("/patients")
		{
			patients.GET("", patientHandler.GetAll)
			patients.GET("/:id", patientHandler.GetByID)
			patients.POST("", middleware.RoleMiddleware("admin", "doctor"), patientHandler.Create)
			patients.PUT("/:id", middleware.RoleMiddleware("admin", "doctor"), patientHandler.Update)
			patients.DELETE("/:id", middleware.RoleMiddleware("admin"), patientHandler.Delete)
		}

		// Doctors
		doctors := api.Group("/doctors")
		{
			doctors.GET("", doctorHandler.GetAll)
			doctors.GET("/available", doctorHandler.GetAvailable)
			doctors.GET("/:id", doctorHandler.GetByID)
			doctors.GET("/department/:deptId", doctorHandler.GetByDepartment)
			doctors.POST("", middleware.RoleMiddleware("admin"), doctorHandler.Create)
			doctors.DELETE("/:id", middleware.RoleMiddleware("admin"), doctorHandler.Delete)
			doctors.POST("/schedules", middleware.RoleMiddleware("admin", "doctor"), doctorHandler.AddSchedule)
			doctors.GET("/:id/schedules", doctorHandler.GetSchedules)
		}

		// Appointments
		appointments := api.Group("/appointments")
		{
			appointments.GET("", appointmentHandler.GetAll)
			appointments.GET("/:id", appointmentHandler.GetByID)
			appointments.GET("/patient/:patientId", appointmentHandler.GetByPatient)
			appointments.GET("/doctor/:doctorId", appointmentHandler.GetByDoctor)
			appointments.POST("", appointmentHandler.Create)
			appointments.PATCH("/:id/status", appointmentHandler.UpdateStatus)
		}

		// Medical Records
		records := api.Group("/medical-records")
		{
			records.GET("", medicalRecordHandler.GetAll)
			records.GET("/:id", medicalRecordHandler.GetByID)
			records.GET("/patient/:patientId", medicalRecordHandler.GetByPatient)
			records.POST("", middleware.RoleMiddleware("admin", "doctor"), medicalRecordHandler.Create)
			records.POST("/:id/prescriptions", middleware.RoleMiddleware("admin", "doctor"), medicalRecordHandler.AddPrescription)
		}

		// Medications
		medications := api.Group("/medications")
		{
			medications.GET("", medicationHandler.GetAll)
			medications.GET("/:id", medicationHandler.GetByID)
			medications.POST("", middleware.RoleMiddleware("admin"), medicationHandler.Create)
			medications.PUT("/:id", middleware.RoleMiddleware("admin"), medicationHandler.Update)
			medications.DELETE("/:id", middleware.RoleMiddleware("admin"), medicationHandler.Delete)
		}

		// Billing
		billings := api.Group("/billings")
		{
			billings.GET("", billingHandler.GetAll)
			billings.GET("/:id", billingHandler.GetByID)
			billings.GET("/patient/:patientId", billingHandler.GetByPatient)
			billings.POST("", billingHandler.Create)
			billings.POST("/:id/pay", billingHandler.Pay)
		}

		// Rooms & Inpatients
		rooms := api.Group("/rooms")
		{
			rooms.GET("", roomHandler.GetAll)
			rooms.GET("/available", roomHandler.GetAvailable)
		}

		inpatients := api.Group("/inpatients")
		{
			inpatients.GET("/active", roomHandler.GetActiveInpatients)
			inpatients.GET("/:id", roomHandler.GetInpatientByID)
			inpatients.POST("/admit", middleware.RoleMiddleware("admin", "doctor"), roomHandler.AdmitPatient)
			inpatients.POST("/:id/discharge", middleware.RoleMiddleware("admin", "doctor"), roomHandler.DischargePatient)
		}
	}

	log.Printf("SIMRS server starting on port %s", cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
