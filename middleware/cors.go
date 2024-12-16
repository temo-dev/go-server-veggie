package middleware

import (
	"errors"
	"log"
	"net/http"
	commonError "server-veggie/common/error"
	"server-veggie/config"

	"github.com/gin-gonic/gin"
)

var (
	errorNoContent = errors.New("no content")
)

func CORSMiddleware() gin.HandlerFunc {
	return func(content *gin.Context) {
		config, err := config.LoadConfig("../")
		if err != nil {
			log.Fatalf("error loading config: %v", err)
		}
		origin := content.Request.Header.Get("Origin")
		// readUserIP(content.Request)
		// fmt.Println("origin", origin)
		// getUserIP(content.Writer, content.Request)
		if origin == config.AllowedOrigins {
			content.Writer.Header().Set("Access-Control-Allow-Origin", config.AllowedOrigins)
			content.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			content.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
			content.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		} else {
			content.JSON(http.StatusInternalServerError, commonError.ErrInternal(errorNoContent))
			content.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		content.Next()
	}
}

// func readUserIP(r *http.Request) string {
// 	ip, port, err := net.SplitHostPort(r.RemoteAddr)
// 	userIp := net.ParseIP(ip)
// 	if userIp == nil {
// 		fmt.Println("err====")
// 	}
// 	forward := r.Header.Get("X-Forwarded-For")
// 	fmt.Println("userIp", userIp)
// 	fmt.Println("port", port)
// 	fmt.Println("err", err)
// 	fmt.Println("forward", forward)
// 	return ""
// }

// Get the IP address of the server's connected user.
// func getUserIP(httpWriter http.ResponseWriter, httpServer *http.Request) {
// 	var userIP string
// 	if len(httpServer.Header.Get("CF-Connecting-IP")) > 1 {
// 		userIP = httpServer.Header.Get("CF-Connecting-IP")
// 		fmt.Println(net.ParseIP(userIP))
// 	} else if len(httpServer.Header.Get("X-Forwarded-For")) > 1 {
// 		userIP = httpServer.Header.Get("X-Forwarded-For")
// 		fmt.Println(net.ParseIP(userIP))
// 	} else if len(httpServer.Header.Get("X-Real-IP")) > 1 {
// 		userIP = httpServer.Header.Get("X-Real-IP")
// 		fmt.Println(net.ParseIP(userIP))
// 	} else {
// 		userIP = httpServer.RemoteAddr
// 		if strings.Contains(userIP, ":") {
// 			fmt.Println(net.ParseIP(strings.Split(userIP, ":")[0]))
// 		} else {
// 			fmt.Println(net.ParseIP(userIP))
// 		}
// 	}
// }
