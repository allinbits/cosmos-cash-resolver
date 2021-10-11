package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/allinbits/cosmos-cash-resolver/resolver"
	"github.com/allinbits/cosmos-cash-resolver/x/did/types"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
)

const (
	MaxReqPerSecond = 5
	DidPrefix       = "did:cosmos:"
)

var (
	serverAddr = flag.String("grpc-server-address", "localhost:9090", "The target grpc server address in the format of host:port")
	listenAddr = flag.String("listen-address", "0.0.0.0:2109", "The REST server listen address in the format of host:port")
	rpsLimit   = flag.Int("mrps", 10, "Max-Requests-Per-Seconds: define the throttle limit in requests per seconds")
)

// loadSettings load the settings from flags and env vars. Env vars have priority over flags
func loadSettings() {
	flag.Parse()
	if val := os.Getenv("GRPC_SERVER_ADDRESS"); val != "" {
		serverAddr = &val
	}
	if val := os.Getenv("LISTEN_ADDRESS"); val != "" {
		listenAddr = &val
	}
	if val := os.Getenv("MRPS"); val != "" {
		l, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalln("invalid int value for MRPS")
		}
		rpsLimit = &l
	}
}

// openGRPCConnection
func openGRPCConnection(addr string) (conn *grpc.ClientConn, err error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(time.Duration(3) * time.Second),
	}

	conn, err = grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	return
}

type Runtime struct {
	resolves  uint64
	startTime time.Time
}

const (
	didLDJson = "application/did+ld+json"
)

func main() {
	loadSettings()

	// setup server
	e := echo.New()
	e.Use(middleware.Logger())
	e.HideBanner = true
	e.StdLogger.Println("starting cosmos-cash-resolver rest server")
	e.StdLogger.Println("target node is ", *serverAddr)
	// track the curent runtime session
	rt := Runtime{
		resolves:  0,
		startTime: time.Now(),
	}
	// open grpc connection
	conn, err := openGRPCConnection(*serverAddr)
	if err != nil {
		e.StdLogger.Fatal(err)
	}
	defer conn.Close()
	// start the rest server
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Use(middleware.RateLimiter(
		middleware.NewRateLimiterMemoryStore(rate.Limit(*rpsLimit)),
	))

	e.GET("/identifier/:did", func(c echo.Context) error {
		did := c.Param("did")
		// decode the paramater
		did, err = url.QueryUnescape(did)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{})
		}
		accept := strings.Split(c.Request().Header.Get("accept"), ";")[0]
		opt := resolver.ResolutionOption{Accept: accept}
		rr := resolver.ResolveRepresentation(conn, did, opt)

		// add universal resolver specific data:
		rr.ResolutionMetadata.DidProperties = map[string]string{
			"method":           "cosmos",
			"methodSpecificId": strings.TrimPrefix(rr.Document.Id, DidPrefix),
		}

		// track the resolution
		atomic.AddUint64(&rt.resolves, 1)
		c.Response().Header().Set(echo.HeaderContentType, didLDJson)

		return c.JSON(http.StatusOK, rr)
	})

	e.GET("/identifier/aries/:did", func(c echo.Context) error {
		did := c.Param("did")
		// decode the paramater
		did, err = url.QueryUnescape(did)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{})
		}
		accept := strings.Split(c.Request().Header.Get("accept"), ";")[0]
		opt := resolver.ResolutionOption{Accept: accept}
		rr := resolver.ResolveRepresentation(conn, did, opt)

		// add universal resolver specific data:
		rr.ResolutionMetadata.DidProperties = map[string]string{
			"method":           "cosmos",
			"methodSpecificId": strings.TrimPrefix(rr.Document.Id, DidPrefix),
		}

		// this will replace a did document verification material formatted in multibase to hex
		for i, vm := range rr.Document.VerificationMethod {
			if material, ok := vm.VerificationMaterial.(*types.VerificationMethod_PublicKeyMyltibase); ok {
				if material.PublicKeyMyltibase[0:1] != "F" {
					continue
				}
				vmHex := &types.VerificationMethod_PublicKeyHex{
					PublicKeyHex: material.PublicKeyMyltibase[1:],
				}
				rr.Document.VerificationMethod[i].VerificationMaterial = vmHex
			}
		}

		result := map[string]interface{}{
			"didDocument": resolver.AriesDidDocument{
				DidDocument: rr.Document,
				Created:     rr.Metadata.Created,
				Updated:     rr.Metadata.Updated,
			},
		}

		// track the resolution
		atomic.AddUint64(&rt.resolves, 1)

		c.Response().Header().Set(echo.HeaderContentType, didLDJson)
		return c.JSON(http.StatusOK, result)
	})

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, fmt.Sprintf(`
		<html><head></head><body style="font-family:courier; padding:.3rem .8rem"> 
		<h1>Cosmos Cash DID Resolver</h1>
		<p>started on %v</p>
		<p>DIDs resolved since starting:</p>
		<p style="font-size:124;margin:.3rem 0">%v</p>
		<br/><br/>
		Visit <a href="https://github.com/allinbits/cosmos-cash">Cosmos Cash</a> for more info
		</body></html>`, rt.startTime.Format(time.RFC3339), rt.resolves))
	})
	e.StdLogger.Fatal(e.Start(*listenAddr))
}
