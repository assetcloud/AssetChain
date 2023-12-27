@echo on

set BUILDTIME=%date:~3%-%time:~0,2%:%time:~3,2%:%time:~6,2%
echo %BUILDTIME%

for /F "delims=" %%i in ('git rev-parse --short HEAD') do ( set commitid=%%i)

set BUILD_FLAGS="-X github.com/assetcloud/assetchain/version.GitCommit=%commitid%  -X github.com/assetcloud/chain/common/version.GitCommit=%commitid% -X github.com/assetcloud/assetchain/version.BuildTime=%BUILDTIME% -w -s"

go env -w CGO_ENABLED=1
go build  -ldflags  %BUILD_FLAGS% -v -o assetchain.exe github.com/assetcloud/assetchain
go build  -ldflags  %BUILD_FLAGS% -v -o assetchain-cli.exe github.com/assetcloud/assetchain/cli
