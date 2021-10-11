module github.com/allinbits/cosmos-cash-resolver

go 1.16

// UNCOMMENT THE COSMOS-SDK TO GENERATE THE PROTO FILES
// require github.com/cosmos/cosmos-sdk v0.44.0 // indirect
// replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/labstack/echo/v4 v4.5.0
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5 // indirect
	golang.org/x/sys v0.0.0-20210603081109-ebe580a85c40 // indirect
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
