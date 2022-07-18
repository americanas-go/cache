package cache

type OptionSet func() func(opt *Option)

func SaveEmpty() func(opt *Option) {
	return func(opt *Option) {
		opt.SaveEmpty = true
	}
}

type Option struct {
	SaveEmpty bool
}
