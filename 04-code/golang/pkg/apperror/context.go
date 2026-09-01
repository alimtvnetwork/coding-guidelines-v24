package apperror

// WithContext attaches a generic key-value pair to the Fault context.
func (f *Fault) WithContext(key string, value any) *Fault {
	if f == nil {
		return nil
	}

	f.context[key] = value

	return f
}

// WithUrl attaches a request URL to the diagnostic context.
func (f *Fault) WithUrl(url string) *Fault {
	return f.WithContext("url", url)
}

// WithStatusCode attaches an HTTP status code to the error.
func (f *Fault) WithStatusCode(statusCode int) *Fault {
	if f != nil {
		f.statusCode = statusCode
	}

	return f.WithContext("statusCode", statusCode)
}

// WithEndpoint attaches an API endpoint path to the context.
func (f *Fault) WithEndpoint(endpoint string) *Fault {
	return f.WithContext("endpoint", endpoint)
}

// WithSiteId attaches a target Site ID to the context.
func (f *Fault) WithSiteId(siteId int64) *Fault {
	return f.WithContext("siteId", siteId)
}

// WithSnapshotId attaches a snapshot identifier to the context.
func (f *Fault) WithSnapshotId(snapshotId string) *Fault {
	return f.WithContext("snapshotId", snapshotId)
}

// WithSlug attaches a plugin or entity slug to the context.
func (f *Fault) WithSlug(slug string) *Fault {
	return f.WithContext("slug", slug)
}

// WithPluginContext attaches both plugin ID and slug simultaneously.
func (f *Fault) WithPluginContext(pluginId int64, slug string) *Fault {
	return f.WithContext("pluginId", pluginId).WithSlug(slug)
}

// Context returns a copy of the underlying diagnostic metadata map.
func (f *Fault) Context() map[string]any {
	if f == nil {
		return map[string]any{}
	}

	copied := make(map[string]any, len(f.context))
	for k, v := range f.context {
		copied[k] = v
	}

	return copied
}
