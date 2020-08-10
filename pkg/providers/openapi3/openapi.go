package openapi

// Object represents the version 3 open api specification
type Object struct {
	Version    string               `json:"openapi,omitempty"`
	Servers    []*ServerRef         `json:"servers,omitempty"`
	Paths      map[string]*PathItem `json:"paths,omitempty"`
	Components *Components          `json:"components,omitempty"`
}

// ServerRef represents a server reference
type ServerRef struct {
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
}

// Info provides metadata about the API.
// The metadata MAY be used by the clients if needed,
// and MAY be presented in editing or documentation
// generation tools for convenience.
type Info struct {
	Description string   `json:"description,omitempty"`
	Version     string   `json:"version,omitempty"`
	Title       string   `json:"title,omitempty"`
	Contact     *Contact `json:"contact,omitempty"`
}

// Contact information for the exposed API.
type Contact struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

// Components Holds a set of reusable objects for different
// aspects of the OAS. All objects defined within the
// components object will have no effect on the API unless
// they are explicitly referenced from properties outside
// the components object.
type Components struct {
	Schemas map[string]*Schema `json:"schemas,omitempty"`
}

// Schema Object allows the definition of input and output
// data types. These types can be objects, but also
// primitives and arrays.
type Schema struct {
	Type        string             `json:"type,omitempty"`
	Description string             `json:"description,omitempty"`
	Required    []string           `json:"required,omitempty"`
	Properties  map[string]*Schema `json:"properties,omitempty"`
	Default     interface{}        `json:"default,omitempty"`
	Enum        []interface{}      `json:"enum,omitempty"`
}

// PathItem describes the operations available on a single path.
// A Path Item MAY be empty, due to ACL constraints.
// The path itself is still exposed to the documentation
// viewer but they will not know which operations and
// parameters are available.
type PathItem struct {
	Ref         string     `json:"ref,omitempty"`
	Description string     `json:"description,omitempty"`
	Get         *Operation `json:"get,omitempty"`
	Put         *Operation `json:"put,omitempty"`
	Post        *Operation `json:"post,omitempty"`
	Delete      *Operation `json:"delete,omitempty"`
	Options     *Operation `json:"options,omitempty"`
	Head        *Operation `json:"head,omitempty"`
	Patch       *Operation `json:"patch,omitempty"`
	Trace       *Operation `json:"trace,omitempty"`
}

// Operation describes a single API operation on a path.
type Operation struct {
	Tags        []string             `json:"tags,omitempty"`
	Description string               `json:"description,omitempty"`
	OperationID string               `json:"operationId,omitempty"`
	Parameters  []*Parameter         `json:"parameters,omitempty"`
	RequestBody *RequestBody         `json:"requestBody,omitempty"`
	Responses   map[string]*Response `json:"responses,omitempty"`
}

// Parameter a list of parameters that are applicable for this operation.
type Parameter struct {
}

// RequestBody the request body applicable for this operation.
// The requestBody is only supported in HTTP methods where
// the HTTP 1.1 specification RFC7231 has explicitly defined
// semantics for request bodies
type RequestBody struct{}

// Response the list of possible responses as they are returned
// from executing this operation.
type Response struct{}