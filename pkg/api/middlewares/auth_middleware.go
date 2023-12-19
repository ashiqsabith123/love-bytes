package middlewares

import (
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/ashiqsabith123/api-gateway/pkg/helper"
	"github.com/ashiqsabith123/api-gateway/pkg/models/responce"
	"github.com/gin-gonic/gin"
)

type client struct {
	rate      int
	reg       bool
	starttime time.Time
}

var (
	mutex   sync.RWMutex
	clients = make(map[string]*client)
)

func ApiRateLimiter(C *gin.Context) {
	mutex.Lock()
	defer mutex.Unlock()

	ip := C.ClientIP()

	go func() {

		for {
			mutex.Lock()

			for ip, client := range clients {
				if time.Since(client.starttime) > 1*time.Hour {
					delete(clients, ip)
				}

			}

			mutex.Unlock()

		}

	}()

	if j, ok := clients[ip]; !ok {
		clients[ip] = &client{
			rate:      1,
			reg:       true,
			starttime: time.Now(),
		}

	} else {
		j.rate++
		if j.rate >= 3 {
			C.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":   "reached maximum otp limit",
				"message": "try after some time",
			})
			return
		}
	}

	C.Next()

}

func Authenticate(C *gin.Context) {

	authHeader := C.GetHeader("Authorization")

	token := strings.Split(authHeader, " ")

	if len(token) < 2 {

		resp := responce.ErrorReposonce(401, "Unauthorized", "token not found")
		C.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		return
	}

	claim, err := helper.ValidateJWTTokens(token[1])

	if err != nil {

		resp := responce.ErrorReposonce(401, "Unauthorized", err.Error())
		C.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		return
	}

	C.Set("userID", claim["id"].(float64))

}
