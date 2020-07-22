package mic

import "fmt"

// 以下几个全局变量可在编译期间控制
var (
	GitCommit = "what"
	BuildTime = "when"
	Who       = "who"
)

// BuildVersion 主版本号由外界传入
func BuildVersion(mainVersion string) string {
	return fmt.Sprintf("%s|%s|%s|%s", mainVersion, GitCommit, BuildTime, Who)
}
