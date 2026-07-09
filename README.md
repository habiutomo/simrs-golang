# SIMRS - Smart Hospital Information System

Sistem Informasi Manajemen Rumah Sakit berbasis **Go** (Golang) menggunakan framework **Gin**, **GORM**, dan **MySQL**.

## Fitur

- **Autentikasi & Otorisasi** – Login JWT, middleware role-based (admin, doctor, staff)
- **Manajemen Pasien** – CRUD pasien
- **Manajemen Dokter** – CRUD dokter, jadwal praktik
- **Manajemen Janji Temu** – Booking appointment, update status
- **Rekam Medis** – Catatan medis dan resep obat
- **Farmasi** – Manajemen obat
- **Manajemen Kamar** – Kamar dan rawat inap (admit/discharge)
- **Penagihan** – Billing dan pembayaran
- **Dashboard** – Ringkasan data rumah sakit
- **Departemen** – Manajemen departemen

## Tech Stack

| Teknologi    | Keterangan                      |
|-------------|----------------------------------|
| Go 1.21     | Bahasa pemrograman              |
| Gin         | HTTP framework                  |
| GORM        | ORM database                    |
| MySQL       | Database                        |
| JWT         | Authentication                  |
| Viper       | Konfigurasi                     |

## Struktur Proyek

```
simrs-golang/
├── cmd/server/main.go          # Entry point
├── config/                     # Konfigurasi aplikasi
├── internal/
│   ├── database/               # Koneksi & migrasi database
│   ├── dto/                    # Data Transfer Objects
│   ├── handlers/               # HTTP handlers
│   ├── middleware/              # Auth & CORS middleware
│   ├── models/                 # Model database
│   ├── repositories/           # Layer akses data
│   └── services/               # Business logic
├── migrations/                 # SQL migration
├── pkg/response/               # Response helper
├── config.yaml                 # File konfigurasi
└── go.mod
```

## Instalasi & Menjalankan

### 1. Clone & masuk direktori

```bash
git clone https://github.com/habiutomo/simrs-golang.git
cd simrs-golang
```

### 2. Konfigurasi database

Edit `config.yaml`:

```yaml
server:
  port: 8080
  mode: debug

database:
  host: localhost
  port: 3306
  user: root
  password: ""
  dbname: simrs

jwt:
  secret: simrs-secret-key-2024
  expiry: 24
```

### 3. Jalankan migrasi SQL

Jalankan file `migrations/001_init.sql` di MySQL.

### 4. Jalankan server

```bash
go run cmd/server/main.go
```

Server akan berjalan di `http://localhost:8080`.

## API Endpoints

### Public

| Method | Endpoint         | Deskripsi        |
|--------|------------------|------------------|
| GET    | `/health`        | Health check     |
| POST   | `/api/v1/auth/login` | Login user  |

### Protected (memerlukan token JWT)

#### Profile
| Method | Endpoint               | Role  |
|--------|------------------------|-------|
| GET    | `/api/v1/profile`      | all   |

#### Dashboard
| Method | Endpoint                | Role |
|--------|-------------------------|------|
| GET    | `/api/v1/dashboard`     | all  |

#### Departments
| Method | Endpoint                   | Role |
|--------|----------------------------|------|
| GET    | `/api/v1/departments`      | all  |

#### Patients
| Method | Endpoint                          | Role         |
|--------|-----------------------------------|--------------|
| GET    | `/api/v1/patients`                | all          |
| GET    | `/api/v1/patients/:id`            | all          |
| POST   | `/api/v1/patients`                | admin, doctor|
| PUT    | `/api/v1/patients/:id`            | admin, doctor|
| DELETE | `/api/v1/patients/:id`            | admin        |

#### Doctors
| Method | Endpoint                                    | Role         |
|--------|----------------------------------------------|--------------|
| GET    | `/api/v1/doctors`                            | all          |
| GET    | `/api/v1/doctors/available`                 | all          |
| GET    | `/api/v1/doctors/:id`                        | all          |
| GET    | `/api/v1/doctors/department/:deptId`         | all          |
| POST   | `/api/v1/doctors`                            | admin        |
| DELETE | `/api/v1/doctors/:id`                        | admin        |
| POST   | `/api/v1/doctors/schedules`                  | admin, doctor|
| GET    | `/api/v1/doctors/:id/schedules`              | all          |

#### Appointments
| Method | Endpoint                                 | Role |
|--------|------------------------------------------|------|
| GET    | `/api/v1/appointments`                   | all  |
| GET    | `/api/v1/appointments/:id`               | all  |
| GET    | `/api/v1/appointments/patient/:patientId`| all  |
| GET    | `/api/v1/appointments/doctor/:doctorId`  | all  |
| POST   | `/api/v1/appointments`                   | all  |
| PATCH  | `/api/v1/appointments/:id/status`        | all  |

#### Medical Records
| Method | Endpoint                                          | Role         |
|--------|----------------------------------------------------|--------------|
| GET    | `/api/v1/medical-records`                          | all          |
| GET    | `/api/v1/medical-records/:id`                      | all          |
| GET    | `/api/v1/medical-records/patient/:patientId`       | all          |
| POST   | `/api/v1/medical-records`                          | admin, doctor|
| POST   | `/api/v1/medical-records/:id/prescriptions`        | admin, doctor|

#### Medications
| Method | Endpoint                          | Role  |
|--------|-----------------------------------|-------|
| GET    | `/api/v1/medications`             | all   |
| GET    | `/api/v1/medications/:id`         | all   |
| POST   | `/api/v1/medications`             | admin |
| PUT    | `/api/v1/medications/:id`         | admin |
| DELETE | `/api/v1/medications/:id`         | admin |

#### Billings
| Method | Endpoint                              | Role |
|--------|---------------------------------------|------|
| GET    | `/api/v1/billings`                    | all  |
| GET    | `/api/v1/billings/:id`                | all  |
| GET    | `/api/v1/billings/patient/:patientId` | all  |
| POST   | `/api/v1/billings`                    | all  |
| POST   | `/api/v1/billings/:id/pay`            | all  |

#### Rooms & Inpatients
| Method | Endpoint                                    | Role         |
|--------|----------------------------------------------|--------------|
| GET    | `/api/v1/rooms`                              | all          |
| GET    | `/api/v1/rooms/available`                   | all          |
| GET    | `/api/v1/inpatients/active`                  | all          |
| GET    | `/api/v1/inpatients/:id`                     | all          |
| POST   | `/api/v1/inpatients/admit`                   | admin, doctor|
| POST   | `/api/v1/inpatients/:id/discharge`           | admin, doctor|
