module github.com/allinbits/cosmos-cash-resolver

go 1.16

// UNCOMMENT THE COSMOS-SDK TO GENERATE THE PROTO FILES
// require github.com/cosmos/cosmos-sdk v0.44.0 // indirect
// replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4

require (
	github.com/labstack/echo/v4 v4.5.0 // indirect
	github.com/cosmos/cosmos-sdk v0.44.0
)
