# chimu_client

## How to build
~~~~bash
export GOROOT_FINAL=/tmp
export GOOS=windows
export GOARCH=amd64
go build -trimpath -ldflags "-s -w" -o chimu_client

# Optional: Compress exectuable with upx
upx chimu_client
~~~~
