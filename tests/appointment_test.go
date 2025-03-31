// tests/appointment_test.go
package tests

import (
    "testing"
    /* "time"
    "bytes"
    "encoding/json" */
    "net/http"
    "net/http/httptest"
    "github.com/stretchr/testify/assert"
    "github.com/javice/vet-clinic-api/internal/models"
)

func TestAppointmentEndpoints(t *testing.T) {
    router, db, err := setupTestRouter()
    assert.NoError(t, err)

    // Creamos primero un cliente y una mascota para asociar la cita
    client := models.Client{Name: "Test Client", Email: "test@test.com", Phone: "123"}
    db.Create(&client)
    
    pet := models.Pet{Name: "Test Pet", Species: "Dog", ClientID: client.ID}
    db.Create(&pet)

    /* t.Run("Create Appointment", func(t *testing.T) {
        appt := models.Appointment{
            PetID: pet.ID,
            Date: time.Now().Add(24 * time.Hour),
            Reason: "Checkup",
        }

        jsonData, _ := json.Marshal(appt)
        req, _ := http.NewRequest("POST", "/api/v1/appointments", bytes.NewBuffer(jsonData))
        req.Header.Set("Content-Type", "application/json")
        
        resp := httptest.NewRecorder()
        router.ServeHTTP(resp, req)
        
        assert.Equal(t, http.StatusCreated, resp.Code)
    }) */

    t.Run("Get All Appointments", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/appointments", nil)
		req.Header.Set("Content-Type", "application/json")
		
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		
		assert.Equal(t, http.StatusOK, resp.Code)
	})

	/* t.Run("Get Appointment By ID", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/appointments/1", nil)
		req.Header.Set("Content-Type", "application/json")
		
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		
		assert.Equal(t, http.StatusOK, resp.Code)
	}) */

	/* t.Run("Update Appointment", func(t *testing.T) {
		appt := models.Appointment{
			PetID: pet.ID,
			Date: time.Now().Add(24 * time.Hour),
			Reason: "Checkup",
		}

		jsonData, _ := json.Marshal(appt)
		req, _ := http.NewRequest("PUT", "/api/v1/appointments/1", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		
		assert.Equal(t, http.StatusOK, resp.Code)
	}) */

	t.Run("Delete Appointment", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/api/v1/appointments/1", nil)
		req.Header.Set("Content-Type", "application/json")
		
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		
		assert.Equal(t, http.StatusOK, resp.Code)
	})
}