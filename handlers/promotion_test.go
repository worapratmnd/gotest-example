package handlers_test

import (
	"fmt"
	"gotest/handlers"
	"gotest/services"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		amount := 100
		expected := 80

		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(expected, nil)

		promoHandler := handlers.NewPromotionHandler(promoService)

		app := fiber.New()
		// http://localhost:8000/calculate?amount=100
		app.Get("/calculate", promoHandler.CalculateDiscount)

		// Act
		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)
		res, _ := app.Test(req)
		defer res.Body.Close()

		// Assert
		assert.Equal(t, fiber.StatusOK, res.StatusCode)

		ioutil.ReadAll(res.Body)
	})
}
