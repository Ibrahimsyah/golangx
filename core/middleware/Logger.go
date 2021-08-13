package middleware

import echoMiddleware "github.com/labstack/echo/v4/middleware"

var config = echoMiddleware.LoggerConfig{
	Format:           "[${time_custom}] ${method}\t${uri}\t${remote_ip}\t${status}\n",
	CustomTimeFormat: "2006-01-02 15:04:05",
}
var Logger = echoMiddleware.LoggerWithConfig(config)
