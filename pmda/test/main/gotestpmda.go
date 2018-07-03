package main

import "github.com/performancecopilot/pcp-go/pmda"

func main()  {
	gotest_pmda := pmda.New(pmda.PmdaInterfaceVLatest, "gotest", 50, "gotest.log", "/root/somehelp")
	gotest_pmda.PmdaConnect()
	gotest_pmda.PmdaMain()

}