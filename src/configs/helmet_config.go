package configs

import (
	"github.com/gofiber/helmet/v2"
)

func HelmetConfig() helmet.Config {

	// Return de config die gebruikt kan worden in ./middleware/fiber_middleware.go
	// Deze is nu leeg.
	return helmet.Config{}
}
