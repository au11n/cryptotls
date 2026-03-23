package fips140

func CAST(name string, f func() error) {
	if !Enabled {
		return
	}
	if err := f(); err != nil {
		panic("FIPS 140-3 self-test failed: " + name + ": " + err.Error())
	}
}

func PCT(name string, f func() error) {
	if !Enabled {
		return
	}
	if err := f(); err != nil {
		panic("FIPS 140-3 self-test failed: " + name + ": " + err.Error())
	}
}
