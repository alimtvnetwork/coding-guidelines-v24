package appfault

// WithContext attaches a generic key-value pair to the AppError context.
func (e *AppError) WithContext(key string, value any) *AppError {
	if e == nil {
		return nil
	}

	if e.Ctx == nil {
		e.Ctx = NewContextMap()
	}

	e.Ctx.Set(key, value)

	return e
}

// WithOp attaches an operation name to the diagnostic context.
func (e *AppError) WithOp(op string) *AppError {
	return e.WithContext("Op", op)
}

// WithSeverity attaches a severity level to the diagnostic context.
func (e *AppError) WithSeverity(severity string) *AppError {
	return e.WithContext("Severity", severity)
}

// WithPriority attaches a priority level (e.g. Critical, High) to the context.
func (e *AppError) WithPriority(priority string) *AppError {
	return e.WithContext("Priority", priority)
}

// WithUrl attaches a request URL to the diagnostic context.
func (e *AppError) WithUrl(url string) *AppError {
	return e.WithContext("Url", url)
}

// WithStatusCode attaches an HTTP status code to the error.
func (e *AppError) WithStatusCode(statusCode int) *AppError {
	if e != nil {
		e.StatusCode = statusCode
	}

	return e.WithContext("StatusCode", statusCode)
}

// WithEndpoint attaches an API endpoint path to the context.
func (e *AppError) WithEndpoint(endpoint string) *AppError {
	return e.WithContext("Endpoint", endpoint)
}

// WithSiteId attaches a target Site ID to the context.
func (e *AppError) WithSiteId(siteId int64) *AppError {
	return e.WithContext("SiteId", siteId)
}

// WithSnapshotId attaches a snapshot identifier to the context.
func (e *AppError) WithSnapshotId(snapshotId string) *AppError {
	return e.WithContext("SnapshotId", snapshotId)
}

// WithSlug attaches a plugin or entity slug to the context.
func (e *AppError) WithSlug(slug string) *AppError {
	return e.WithContext("Slug", slug)
}

// WithPluginContext attaches both plugin ID and slug simultaneously.
func (e *AppError) WithPluginContext(pluginId int64, slug string) *AppError {
	return e.WithContext("PluginId", pluginId).WithSlug(slug)
}

// Context returns a copy of the underlying diagnostic metadata ContextMap.
func (e *AppError) Context() ContextMap {
	if e == nil || e.Ctx == nil {
		return NewContextMap()
	}

	return e.Ctx.Clone()
}
