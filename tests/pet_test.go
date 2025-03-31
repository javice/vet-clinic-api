package tests

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"
	"strconv"

    "github.com/stretchr/testify/assert"
    "github.com/javice/vet-clinic-api/internal/models"
)

func TestPetEndpoints(t *testing.T) {
    // Setup test router
	router, db, err := setupTestRouter()
    if err != nil {
        t.Fatalf("Error setting up test router: %v", err)
    } 
	

    // Crear un cliente para asociar mascotas
    client := models.Client{
        Name:    "Pet Owner",
        Email:   "pet.owner@example.com",
        Phone:   "123456789",
        Address: "Pet Owner Address",
    }
    db.Create(&client)

	

    // Test POST /pets
    t.Run("Create Pet", func(t *testing.T) {
        petData := models.Pet{
            Name:        "Fluffy",
            Species:     "Dog",
            Breed:       "Golden Retriever",
            BirthDate:   time.Now().AddDate(-2, 0, 0),
            ClientID:    client.ID,
            Description: "Friendly dog",
        }

        jsonData, _ := json.Marshal(petData)
        req, _ := http.NewRequest("POST", "/api/v1/pets", bytes.NewBuffer(jsonData))
        req.Header.Set("Content-Type", "application/json")

        resp := httptest.NewRecorder()
        router.ServeHTTP(resp, req)

        assert.Equal(t, http.StatusCreated, resp.Code)

        var createdPet models.Pet
        err := json.Unmarshal(resp.Body.Bytes(), &createdPet)
        assert.NoError(t, err)
        assert.Equal(t, petData.Name, createdPet.Name)
        assert.Equal(t, petData.Species, createdPet.Species)
        assert.NotZero(t, createdPet.ID)
    })

    // Test GET /pets
    t.Run("Get All Pets", func(t *testing.T) {
        req, _ := http.NewRequest("GET", "/api/v1/pets", nil)
        resp := httptest.NewRecorder()
        router.ServeHTTP(resp, req)

        assert.Equal(t, http.StatusOK, resp.Code)

        var pets []models.Pet
        err := json.Unmarshal(resp.Body.Bytes(), &pets)
        assert.NoError(t, err)
        assert.GreaterOrEqual(t, len(pets), 1)
    })

    // Test GET /pets/:id
    t.Run("Get Pet By ID", func(t *testing.T) {
        // Obtener una mascota primero
        var pet models.Pet
        db.First(&pet)

        req, _ := http.NewRequest("GET", "/api/v1/pets/"+strconv.FormatUint(uint64(pet.ID), 10), nil)
        resp := httptest.NewRecorder()
        router.ServeHTTP(resp, req)

        if pet.ID > 0 {
            assert.Equal(t, http.StatusOK, resp.Code)

            var fetchedPet models.Pet
            err := json.Unmarshal(resp.Body.Bytes(), &fetchedPet)
            assert.NoError(t, err)
            assert.Equal(t, pet.ID, fetchedPet.ID)
        }
    })

    // Test GET /pets?client_id=X
    t.Run("Get Pets By Client ID", func(t *testing.T) {
        req, _ := http.NewRequest("GET", "/api/v1/pets?client_id="+strconv.FormatUint(uint64(client.ID), 10), nil)
        resp := httptest.NewRecorder()
        router.ServeHTTP(resp, req)

        assert.Equal(t, http.StatusOK, resp.Code)

        var pets []models.Pet
        err := json.Unmarshal(resp.Body.Bytes(), &pets)
        assert.NoError(t, err)

        for _, pet := range pets {
            assert.Equal(t, client.ID, pet.ClientID)
        }
    })

    // Test PUT /pets/:id
    t.Run("Update Pet", func(t *testing.T) {
        // Obtener una mascota primero
        var pet models.Pet
        db.First(&pet)

        if pet.ID > 0 {
            updateData := models.Pet{
                Name:        "Updated Pet Name",
                Species:     pet.Species,
                Breed:       "Updated Breed",
                ClientID:    pet.ClientID,
                Description: "Updated description",
            }

            jsonData, _ := json.Marshal(updateData)
            req, _ := http.NewRequest("PUT", "/api/v1/pets/"+strconv.FormatUint(uint64(pet.ID), 10), bytes.NewBuffer(jsonData))
            req.Header.Set("Content-Type", "application/json")

            resp := httptest.NewRecorder()
            router.ServeHTTP(resp, req)

            assert.Equal(t, http.StatusOK, resp.Code)

            var updatedPet models.Pet
            err := json.Unmarshal(resp.Body.Bytes(), &updatedPet)
            assert.NoError(t, err)
            assert.Equal(t, "Updated Pet Name", updatedPet.Name)
            assert.Equal(t, "Updated Breed", updatedPet.Breed)
        }
    })

	
    // Test DELETE /pets/:id
    t.Run("Delete Pet", func(t *testing.T) {
        // Crear una mascota primero
        pet := models.Pet{
            Name:        "Pet to Delete",
            Species:     "Cat",
            Breed:       "Siamese",
            ClientID:    client.ID,
			Description: "Friendly cat",
		}
		db.Create(&pet)

		req, _ := http.NewRequest("DELETE", "/api/v1/pets/"+strconv.FormatUint(uint64(pet.ID), 10), nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})
}