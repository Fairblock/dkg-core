module github.com/fairblock/dkg-core

go 1.16

require (
	cosmossdk.io/math v1.0.0
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/armon/go-metrics v0.3.9
	github.com/axelarnetwork/tm-events v0.0.0-20211102054919-baac5ca41de2
	github.com/axelarnetwork/utils v0.0.0-20211102051812-9d1028976360
	github.com/btcsuite/btcd v0.22.0-beta
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/cosmos/cosmos-sdk v0.44.7
	github.com/cosmos/ibc-go v1.2.0
	github.com/drand/kyber-bls12381 v0.2.6
	github.com/ethereum/go-ethereum v1.10.11
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.2.0 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/klauspost/compress v1.11.9 // indirect
	github.com/matryer/moq v0.2.3
	github.com/miguelmota/go-ethereum-hdwallet v0.1.1
	github.com/pkg/errors v0.9.1
	github.com/rakyll/statik v0.1.7
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/rs/zerolog v1.23.0
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.8.1
	github.com/stretchr/testify v1.8.2
	github.com/tendermint/tendermint v0.34.14
	github.com/tendermint/tm-db v0.6.4
	github.com/tklauser/go-sysconf v0.3.7 // indirect
	github.com/tyler-smith/go-bip39 v1.1.0 // indirect
	golang.org/x/tools v0.10.0 // indirect
	google.golang.org/genproto v0.0.0-20210828152312-66f60bf46e71
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4 // https://github.com/fairblock/dkg-core/issues/36

// Fix upstream GHSA-h395-qcrw-5vmq vulnerability.
// TODO Remove it: https://github.com/cosmos/cosmos-sdk/issues/10409
replace github.com/gin-gonic/gin => github.com/gin-gonic/gin v1.7.0
