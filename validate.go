package appsearch

func (c AppSearch) Validate() bool {
	if c.ApiKey == " " || len(c.ApiKey) <= 10 {
		return false
	} else if c.EngineName == " " {
		return false
	} else if c.Url == " " {
		return false
	}
	return true
}
