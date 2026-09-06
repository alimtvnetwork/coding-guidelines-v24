package appfault

// WithContext returns a new immutable AppError with the attached key-value context.
// The receiver is never mutated in-place, guaranteeing strict immutability.
func (e *AppError) WithContext(key string, value any) *AppError {
	if e == nil {
		return nil
	}

	cloned := e.clone()
	cloned.ctx.Set(key, value)

	return cloned
}

// WithOp returns a new immutable AppError with the operation context.
func (e *AppError) WithOp(op string) *AppError {
	return e.WithContext("Op", op)
}

// WithSeverity returns a new immutable AppError with the severity level.
func (e *AppError) WithSeverity(severity SeverityType) *AppError {
	return e.WithContext("Severity", severity.Name())
}

// WithPriority returns a new immutable AppError with the priority level.
func (e *AppError) WithPriority(priority PriorityType) *AppError {
	return e.WithContext("Priority", priority.Name())
}

// WithUrl returns a new immutable AppError with the request URL context.
func (e *AppError) WithUrl(url string) *AppError {
	return e.WithContext("Url", url)
}

// WithStatusCode returns a new immutable AppError with the specified HTTP status code.
func (e *AppError) WithStatusCode(statusCode int) *AppError {
	if e == nil {
		return nil
	}

	cloned := e.clone()
	cloned.statusCode = statusCode
	cloned.ctx.Set("StatusCode", statusCode)

	return cloned
}

// WithEndpoint returns a new immutable AppError with the API endpoint context.
func (e *AppError) WithEndpoint(endpoint string) *AppError {
	return e.WithContext("Endpoint", endpoint)
}

// WithSiteId returns a new immutable AppError with the target Site ID context.
func (e *AppError) WithSiteId(siteId int64) *AppError {
	return e.WithContext("SiteId", siteId)
}

// WithSnapshotId returns a new immutable AppError with the snapshot identifier.
func (e *AppError) WithSnapshotId(snapshotId string) *AppError {
	return e.WithContext("SnapshotId", snapshotId)
}

// WithSlug returns a new immutable AppError with the entity slug.
func (e *AppError) WithSlug(slug string) *AppError {
	return e.WithContext("Slug", slug)
}

// WithPluginContext returns a new immutable AppError with plugin ID and slug.
func (e *AppError) WithPluginContext(pluginId int64, slug string) *AppError {
	return e.WithContext("PluginId", pluginId).WithSlug(slug)
}

// WithCaller returns a new immutable AppError with the specified caller site metadata.
func (e *AppError) WithCaller(caller CallerInfo) *AppError {
	if e == nil {
		return nil
	}

	cloned := e.clone()
	cloned.caller = caller

	return cloned
}

// WithPath returns a new immutable AppError with the file or directory path context.
// It sets both "Path" and "FilePath" keys so diagnostic tooling can discover either.
func (e *AppError) WithPath(path string) *AppError {
	if e == nil {
		return nil
	}

	cloned := e.clone()
	cloned.ctx.Set("Path", path)
	cloned.ctx.Set("FilePath", path)

	return cloned
}

// WithFilePath is an alias for WithPath emphasizing targeted file destinations.
func (e *AppError) WithFilePath(filePath string) *AppError {
	return e.WithPath(filePath)
}

// WithPaths returns a new immutable AppError recording multiple target paths.
func (e *AppError) WithPaths(paths ...string) *AppError {
	if e == nil {
		return nil
	}

	cloned := e.clone()
	cloned.ctx.Set("Paths", paths)
	if len(paths) > 0 {
		cloned.ctx.Set("Path", paths[0])
		cloned.ctx.Set("FilePath", paths[0])
	}

	return cloned
}

// WithVar returns a new immutable AppError embedding a named variable and its value.
// It sets the direct variable name key as well as tracking under the "Variables" map.
func (e *AppError) WithVar(name string, value any) *AppError {
	if e == nil {
		return nil
	}

	cloned := e.clone()
	cloned.ctx.Set(name, value)

	vars, isFound := cloned.ctx.Get("Variables")
	var varMap map[string]any
	if isFound {
		if existingMap, isMap := vars.(map[string]any); isMap {
			varMap = make(map[string]any, len(existingMap)+1)
			for k, v := range existingMap {
				varMap[k] = v
			}
		}
	}
	if varMap == nil {
		varMap = make(map[string]any, 1)
	}
	varMap[name] = value
	cloned.ctx.Set("Variables", varMap)

	return cloned
}

// WithVars returns a new immutable AppError embedding multiple variables from a map.
func (e *AppError) WithVars(vars map[string]any) *AppError {
	if e == nil || len(vars) == 0 {
		return e
	}

	cloned := e.clone()
	for k, v := range vars {
		cloned.ctx.Set(k, v)
	}

	existingVars, isFound := cloned.ctx.Get("Variables")
	var varMap map[string]any
	if isFound {
		if m, isMap := existingVars.(map[string]any); isMap {
			varMap = make(map[string]any, len(m)+len(vars))
			for k, v := range m {
				varMap[k] = v
			}
		}
	}
	if varMap == nil {
		varMap = make(map[string]any, len(vars))
	}
	for k, v := range vars {
		varMap[k] = v
	}
	cloned.ctx.Set("Variables", varMap)

	return cloned
}

// Context returns a copy of the underlying diagnostic metadata ContextMap.
func (e *AppError) Context() ContextMap {
	if e == nil || e.ctx == nil {
		return NewContextMap()
	}

	return e.ctx.Clone()
}
