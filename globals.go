package capsule

import "github.com/michaelquigley/cf"

func CfOptions() *cf.Options {
	if cfOpt == nil {
		cfOpt = cf.DefaultOptions()
	}
	return cfOpt
}

var cfOpt *cf.Options
