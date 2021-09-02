package main

import (
	"net/http"
	"strings"

	"github.com/allinbits/cosmos-cash-resolver/resolver"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

const (
	MaxReqPerSecond = 5
	DidPrefix       = "did:cosmos:"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Use(middleware.RateLimiter(
		middleware.NewRateLimiterMemoryStore(MaxReqPerSecond),
	))

	e.GET("/identifier/:did", func(c echo.Context) error {
		did := c.Param("did")
		accept := strings.Split(c.Request().Header.Get("accept"), ";")[0]
		opt := resolver.ResolutionOption{Accept: accept}
		rr := resolver.ResolveRepresentation(clientCtx, did, opt)

		// add universal resolver specific data:
		rr.ResolutionMetadata.DidProperties = map[string]string{
			"method":           "cosmos",
			"methodSpecificId": strings.TrimPrefix(rr.Document.Id, DidPrefix),
		}

		

	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":2109"))
}

/*
// identifierHandler
	didH := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		did := vars["did"]
		// parse mime
		accept := strings.Split(r.Header.Get("accept"), ";")[0]
		// add universal resolver specific data:
		rr.ResolutionMetadata.DidProperties = map[string]string{
			"method":           "cosmos",
			"methodSpecificId": strings.TrimPrefix(rr.Document.Id, types.DidPrefix),
		}
		// cors
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("origin"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "content-type,accept")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(rr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	rtr.PathPrefix("/identifier/{did}").HandlerFunc(didH).Methods(http.MethodGet, http.MethodOptions)
	rtr.PathPrefix("/1.0/identifiers/{did}").HandlerFunc(didH).Methods(http.MethodGet, http.MethodOptions)
*/
