package main

import (
	sigar "github.com/cloudfoundry/gosigar"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/", index)
	e.GET("/ip",clientIP)
	e.GET("/env/:key", getEnv)
	e.GET("/usage", sysUsage)

	e.Logger.Fatal(e.Start("0.0.0.0:1323"))
}

func index(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"/":         "This Page",
		"/ip":       "Show the client IP",
		"/env/:key": "Your can get any environment variable",
		"/usage":    "Show system usage",
	})
}

func clientIP(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"IP": c.RealIP(),
	})
}

func getEnv(c echo.Context) error {
	key := c.Param("key")
	return c.JSON(http.StatusOK, map[string]string{
		"Key":   key,
		"Value": os.Getenv(key),
	})
}

func sysUsage(c echo.Context) error {
	uptime := sigar.Uptime{}
	if err := uptime.Get(); err != nil {
		return err
	}
	avg := sigar.LoadAverage{}
	if err := avg.Get(); err != nil {
		return err
	}
	mem := sigar.Mem{}
	if err := mem.Get(); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"UpTime":        uptime.Format(),
		"LoadAverage1":  avg.One,
		"LoadAverage5":  avg.Five,
		"LoadAverage15": avg.Fifteen,
		"Mem(MB)": map[string]uint64{
			"Total":      ByteToMB(mem.Total),
			"Used":       ByteToMB(mem.Used),
			"ActualUsed": ByteToMB(mem.ActualUsed),
			"Free":       ByteToMB(mem.Free),
			"ActualFree": ByteToMB(mem.ActualFree),
		},
	})
}

func ByteToMB(val uint64) uint64 {
	return val / 1024 / 1014
}
