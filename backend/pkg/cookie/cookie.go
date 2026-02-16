package cookie

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

func AddCookie(c *echo.Context, name, value string) {
	addCookie(c, 24*60*60, name, value) // 86400 секунд = 24 часа
}

func DeleteCookie(c *echo.Context, name, value string) {
	addCookie(c, -1, name, value)
}

func addCookie(c *echo.Context, maxAge int, name, value string) {
	expires := time.Now().Add(time.Duration(maxAge) * time.Second)

	c.SetCookie(
		&http.Cookie{
			Name:     name,
			Value:    value,
			Path:     "/",
			Expires:  expires,
			HttpOnly: true,
			MaxAge:   maxAge,
		},
	)
}
