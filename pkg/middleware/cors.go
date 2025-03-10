// Copyright (c) 2022 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/buz/blob/main/LICENSE

package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/silverton-io/buz/pkg/config"
)

func CORS(conf config.Cors) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", strings.Join(conf.AllowOrigin, ", "))
		c.Header("Access-Control-Allow-Credentials", strconv.FormatBool(conf.AllowCredentials))
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Set-Cookie, Cookie")
		c.Header("Access-Control-Allow-Methods", strings.Join(conf.AllowMethods, ", "))
		c.Header("Access-Control-Max-Age", strconv.Itoa(conf.MaxAge))

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
