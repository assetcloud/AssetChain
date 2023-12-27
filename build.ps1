echo "building assetchain.exe assetchain-cli.exe"
$commitid = git rev-parse --short=8 HEAD
echo $commitid

$BUILDTIME=get-date -format "yyyy-MM-dd/HH:mm:ss"
echo $BUILDTIME

$BUILD_FLAGS='''-X "github.com/assetcloud/assetchain/version.GitCommit={0}"  -X "github.com/assetcloud/chain/common/version.GitCommit={1}" -X "github.com/assetcloud/assetchain/version.BuildTime={2}" -w -s''' -f $commitid,$commitid,$BUILDTIME
echo $BUILD_FLAGS


go env -w CGO_ENABLED=1
go build  -ldflags  $BUILD_FLAGS  -v -o assetchain.exe github.com/assetcloud/assetchain
go build  -ldflags  $BUILD_FLAGS  -v -o assetchain-cli.exe github.com/assetcloud/assetchain/cli

echo "build end"
