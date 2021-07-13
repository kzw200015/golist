package args

import "flag"

var srcPath string
var port int

func GetSrcPath() string {
	return srcPath
}

func GetPort() int {
	return port
}

func init() {
	flag.StringVar(&srcPath, "s", "./", "srcPath")
	flag.IntVar(&port, "p", 8080, "port")
	flag.Parse()
}
