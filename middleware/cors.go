package middleware

import (
	"errors"
	"net/http"
	commonError "server-veggie/common/error"

	"github.com/gin-gonic/gin"
)

var (
	errorNoContent = errors.New("no content")
)

func CORSMiddleware() gin.HandlerFunc {
	return func(content *gin.Context) {
		// readUserIP(content.Request)
		// fmt.Println("origin", origin)

		if true {
			content.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
			content.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			content.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With")
			content.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

			if content.Request.Method == "OPTIONS" {
				content.AbortWithStatus(204)
				return
			}
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
