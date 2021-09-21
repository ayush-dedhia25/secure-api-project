package middleware

import (
   "log"
   "time"
   "strings"
   "net/http"
   "github.com/justinas/alice"
   "github.com/ayush/secure-api/server/middleware/myJWT"
   "github.com/ayush/secure-api/db"
   "github.com/ayush/secure-api/server/templates"
)

func recoverHandler(next http.Handler) http.Handler {
   fn := func(res http.ResponseWriter, req *http.Request) {
      defer func() {
         if err := recover(); err != nil {
            log.Panic("Recovered Panic: %+v", err)
            http.Error(res, http.StatusText(500), 500)
         }
      }()
      next.ServerHTTP(res, req)
   }
   return http.HandlerFunc(fn)
}

func authHandler(next http.Handler) http.Handler {
   fn := func(res http.ResponseWriter, req *http.Request) {
      switch req.URL.Path {
         case "/restricted", "/logout", "/deleteUser":
         default:
      }
   }
}

func logicHandler(res http.ResponseWriter, req *http.Request) {
   switch req.URL.Path {
      case "/restricted":
         csrfSecret := grabCsrfFromRequest(req)
         templates.RenderTemplate(res, "restricted", &templates.Dashboard{csrfSecret, "Hello, Ayush!"})
      case "/login":
         switch req.Method {
            case "GET":
            case "POST":
            default:
         }
      case "/register":
         switch req.Method {
            case "GET":
            case "POST":
            default:
         }
      case "/logout":
      case "/deleteUser":
      default: 
   }
}

func nullifyTokenCookies(res *http.ResponseWriter, req *http.Request) {
   authCookie := http.Cookie{
      Name: "AuthToken",
      Value: "",
      Expires: time.Now().Add(-1000 * time.Hour),
      HttpOnly: true,
   }
   http.SetCookie(*res, &authCookie)
   
   refreshCookie := http.Cookie{
      Name: "RefreshToken",
      Value: "",
      Expires: time.Now().Add(-1000 * time.Hour),
      HttpOnly: true,
   }
   http.SetCookie(*res, &refreshCookie)
   
   refreshCookie, refreshErr := req.Cookie("RefreshToken")
   if refreshErr == http.ErrNoCookie {
      return
   } else if refreshErr != nil {
      log.Panic("Error %+v", refreshErr)
      http.Error(*res, http.StatusText(500), 500)
   }
   
   myJWT.RevokeRefreshToken(refreshCookie.Value)
}

func setAuthAndRefreshCookies(res *http.ResponseWriter, authTokenString string, refreshTokenString string) {
   authCookie := http.Cookie{
      Name: "AuthToken",
      Value: authTokenString,
      HttpOnly: true,
   }
   http.SetCookie(*res, &authCookie)
   
   refreshCookie := http.Cookie{
      Name: "RefreshToken",
      Value: refreshTokenString,
      HttpOnly: true,
   }
   http.SetCookie(*res, &refreshCookie)
}

func grabCsrfFromRequest(req *http.Request) string {
   csrfFromForm := req.FormValue("X-CSRF-Token")
   if csrfFromForm != "" {
      return csrfFromForm
   } else {
      return req.Header.Get("X-CSRF-Token")
   }
}

func NewHandler() http.Handler {
   return alice.New(recoverHandler, authHandler).ThenFunc(logicHandler)
}