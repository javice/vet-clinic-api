package tests

import (
	"github.com/gin-gonic/gin"
    "github.com/javice/vet-clinic-api/internal/handlers"
    "github.com/javice/vet-clinic-api/internal/models"
    "github.com/javice/vet-clinic-api/internal/repositories"
    "github.com/javice/vet-clinic-api/internal/routes"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "testing"
)

func setupTestDB() (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
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

func setupTestRouter() (*gin.Engine, *gorm.DB, error) {
    // Configurar la base de datos en memoria para pruebas
    //db, err := setupTestDB()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    if err != nil {
        return nil, nil, err
    }

	// Ejecutar migraciones
    err = db.AutoMigrate(&models.Client{}, &models.Pet{}, &models.Appointment{})
    if err != nil {
        return nil, nil, err
    }

    // Crear repositorios
    clientRepo := repositories.NewClientRepository(db)
    petRepo := repositories.NewPetRepository(db)
	appointmentRepo := repositories.NewAppointmentRepository(db)

    // Crear handler
    handler := handlers.NewHandler(clientRepo, petRepo, appointmentRepo)

    // Configurar rutas
    router := routes.SetupRouter(handler)

    return router, db, nil
}

func TestMainSuite(t *testing.T) {
    // Ejecutar los tests en el orden deseado
    t.Run("Client Tests", TestClientEndpoints)
    t.Run("Pet Tests", TestPetEndpoints)
    t.Run("Appointment Tests", TestAppointmentEndpoints)
}