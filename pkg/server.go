package pkg

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	proxy "github.com/maxneuvians/go-copilot-proxy/pkg"
)

//go:embed index.html
var content embed.FS

func Start(port string) {
	engine := html.NewFileSystem(http.FS(content), ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	var token = getSessionToken()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Token": "true",
		})
	})

	app.Get("/authenticate/:device_code", func(c *fiber.Ctx) error {
		code := c.Params("device_code")
		accessToken := authenticate(code)
		if accessToken != "" {
			saveToken(accessToken)
			token = getSessionToken()
		}
		return c.JSON(fiber.Map{
			"token": "true",
		})
	})

	app.Post("/chat", func(c *fiber.Ctx) error {
		if token == "" {
			return c.JSON(fiber.Map{
				"error": "Not authenticated",
			})
		}

		var msgs []proxy.Message

		if err := c.BodyParser(&msgs); err != nil {
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		chatResponse, err := chat(token, msgs)

		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"response": chatResponse,
		})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		LoginURL, Code, DeviceCode := login()
		return c.JSON(fiber.Map{
			"LoginURL":   LoginURL,
			"Code":       Code,
			"DeviceCode": DeviceCode,
		})
	})

	app.Listen(":" + port)
}
