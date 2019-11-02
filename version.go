package mic

import "fmt"

var (
	// 以下几个全局变量可在编译期间控制
	GitCommit = "what"
	BuildTime = "when"
	Who       = "who"
)

// 主版本号由外界传入
func BuildVersion(mainVersion string) string {
	return fmt.Sprintf("%s|%s|%s|%s", mainVersion, GitCommit, BuildTime, Who)
}
