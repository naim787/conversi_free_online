package cmd

import "github.com/gofiber/fiber/v2"

func CMD_handler() {
	app := fiber.New()

	// Serve form upload di root
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(`
			<form method="POST" action="/upload" enctype="multipart/form-data">
				<input type="file" name="nmfile">
				<button type="submit">Upload dan Download</button>
			</form>
		`)
	})

	// Handle upload dan kirim langsung file sebagai download
	app.Post("/upload", func(c *fiber.Ctx) error {
		// Ambil file dari form field "nmfile"
		file, err := c.FormFile("nmfile")
		if err != nil {
			return c.Status(400).SendString("Gagal mendapatkan file")
		}

		// Buka file yang diupload
		f, err := file.Open()
		if err != nil {
			return c.Status(500).SendString("Gagal membuka file")
		}
		defer f.Close()

		// Set header agar browser langsung download file-nya
		c.Response().Header.Set("Content-Disposition", "attachment; filename="+file.Filename)
		c.Response().Header.Set("Content-Type", "application/octet-stream")

		// Kirim isi file langsung sebagai response
		return c.SendStream(f)
	})

	app.Listen(":3000")
}