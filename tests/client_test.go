package tests

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/javice/vet-clinic-api/internal/handlers"
    "github.com/javice/vet-clinic-api/internal/models"
    "github.com/javice/vet-clinic-api/internal/repositories"
    "github.com/javice/vet-clinic-api/internal/routes"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm" 
    "strconv"
)

func setupTestDB() (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Migrar esquemas
    err = db.AutoMigrate(&models.Client{}, &models.Pet{})
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

func TestClientEndpoints(t *testing.T) {
    router, db, err := setupTestRouter()
    if err != nil {
        t.Fatalf("Error inicializando el router: %v", router)
    }

	
    // Test POST /clients
    t.Run("Create Client", func(t *testing.T) {
        clientData := models.Client{
            Name:    "Test Client",
            Email:   "test@example.com",
            Phone:   "123456789",
            Address: "Test Address",
        }

        jsonData, _ := json.Marshal(clientData)
        req, _ := http.NewRequest("POST", "/api/v1/clients", bytes.NewBuffer(jsonData))
        req.Header.Set("Content-Type", "application/json")

        resp := httptest.NewRecorder()
        router.ServeHTTP(resp, req)

        assert.Equal(t, http.StatusCreated, resp.Code)

        var createdClient models.Client
        err := json.Unmarshal(resp.Body.Bytes(), &createdClient)
        assert.NoError(t, err)
        assert.Equal(t, clientData.Name, createdClient.Name)
        assert.Equal(t, clientData.Email, createdClient.Email)
        assert.NotZero(t, createdClient.ID)
    })

    // Test GET /clients
    t.Run("Get All Clients", func(t *testing.T) {
        req, _ := http.NewRequest("GET", "/api/v1/clients", nil)
        resp := httptest.NewRecorder()
        router.ServeHTTP(resp, req)

        assert.Equal(t, http.StatusOK, resp.Code)

        var clients []models.Client
        err := json.Unmarshal(resp.Body.Bytes(), &clients)
        assert.NoError(t, err)
        assert.GreaterOrEqual(t, len(clients), 1)
    })

    // Test GET /clients/:id
    t.Run("Get Client By ID", func(t *testing.T) {
        // Crear un cliente primero
        //var client models.Client
		// Crear un cliente en la base de datos
		client := models.Client{
			Name:    "Test Client",
			Email:   "test@example.com",
			Phone:   "123456789",
			Address: "Test Address",
		}
        db.First(&client)

        req, _ := http.NewRequest("GET", "/api/v1/clients/"+strconv.FormatUint(uint64(client.ID), 10), nil)
        resp := httptest.NewRecorder()
        router.ServeHTTP(resp, req)

        if client.ID > 0 {
            assert.Equal(t, http.StatusOK, resp.Code)

            var fetchedClient models.Client
            err := json.Unmarshal(resp.Body.Bytes(), &fetchedClient)
            assert.NoError(t, err)
            assert.Equal(t, client.ID, fetchedClient.ID)
        }
    })

    // Test PUT /clients/:id
    t.Run("Update Client", func(t *testing.T) {
        // Crear un cliente primero
        var client models.Client
        db.First(&client)

        if client.ID > 0 {
            updateData := models.Client{
                Name:    "Updated Name",
                Email:   client.Email,
                Phone:   client.Phone,
                Address: "Updated Address",
            }

            jsonData, _ := json.Marshal(updateData)
            req, _ := http.NewRequest("PUT", "/api/v1/clients/"+strconv.FormatUint(uint64(client.ID), 10), bytes.NewBuffer(jsonData))
            req.Header.Set("Content-Type", "application/json")

            resp := httptest.NewRecorder()
            router.ServeHTTP(resp, req)

            assert.Equal(t, http.StatusOK, resp.Code)

            var updatedClient models.Client
            err := json.Unmarshal(resp.Body.Bytes(), &updatedClient)
            assert.NoError(t, err)
            assert.Equal(t, "Updated Name", updatedClient.Name)
            assert.Equal(t, "Updated Address", updatedClient.Address)
        }
    })

    // Test DELETE /clients/:id
    t.Run("Delete Client", func(t *testing.T) {
        // Crear un cliente primero
        client := models.Client{
            Name:    "Client to Delete",
            Email:   "delete@example.com",
            Phone:   "987654321",
            Address: "Delete Address",
        }
        db.Create(&client)

        req, _ := http.NewRequest("DELETE", "/api/v1/clients/"+strconv.FormatUint(uint64(client.ID), 10), nil)
        resp := httptest.NewRecorder()
        router.ServeHTTP(resp, req)

        assert.Equal(t, http.StatusOK, resp.Code)

        // Verificar que el cliente fue eliminado
        var count int64
        db.Model(&models.Client{}).Where("id = ?", client.ID).Count(&count)
        assert.Equal(t, int64(0), count)
    })
}
