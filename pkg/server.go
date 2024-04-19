package pkg

import (
	//"embed"
	//"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

//#go:embed index.html
//var content embed.FS

func Start(port string) {
	//engine := html.NewFileSystem(http.FS(content), ".html")
	engine := html.New("./pkg", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		LoginURL, Code := login()
		return c.JSON(fiber.Map{
			"LoginURL": LoginURL,
			"Code": Code,
		});
	})

	app.Listen(":" + port)
}