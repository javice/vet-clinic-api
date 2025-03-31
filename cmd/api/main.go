// cmd/api/main.go
package main

import (
    "log"
    "os"

    "github.com/javice/vet-clinic-api/internal/handlers"
    "github.com/javice/vet-clinic-api/internal/repositories"
    "github.com/javice/vet-clinic-api/internal/routes"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "github.com/javice/vet-clinic-api/internal/models"
)

func main() {
    // Configurar la base de datos
    db, err := setupDatabase()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Crear repositorios
    clientRepo := repositories.NewClientRepository(db)
    petRepo := repositories.NewPetRepository(db)
	appointmentRepo := repositories.NewAppointmentRepository(db)

    // Crear handler
    handler := handlers.NewHandler(clientRepo, petRepo, appointmentRepo)

    // Configurar rutas
    router := routes.SetupRouter(handler)

    // Iniciar el servidor
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server running on port %s", port)
    if err := router.Run(":" + port); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}

func setupDatabase() (*gorm.DB, error) {
    // Para desarrollo usamos SQLite
    db, err := gorm.Open(sqlite.Open("vet_clinic.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Migrar esquemas
    err = db.AutoMigrate(&models.Client{}, &models.Pet{}, &models.Appointment{})
    if err != nil {
        return nil, err
    }

    return db, nil
}