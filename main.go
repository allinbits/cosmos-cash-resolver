package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	"github.com/allinbits/cosmos-cash-resolver/resolver"
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
	serverAddr = flag.String("grpc-server", "localhost:9090", "The target grpc server address in the format of host:port")
	listenAddr = flag.String("listen", "localhost:2109", "The REST server listen address in the format of host:port")
	rpsLimit   = flag.Int("mrps", 10, "Max-Requests-Per-Seconds: define the throttle limit in requests per seconds")
)

func openGRPCConnection(addr string) (conn *grpc.ClientConn, err error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
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

func main() {
	flag.Parse()
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

		return c.JSON(http.StatusOK, rr)
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
