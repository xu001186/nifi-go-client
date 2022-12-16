/*
 * NiFi Rest API
 *
 * The Rest API provides programmatic access to command and control a NiFi instance in real time. Start and                                             stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.17.0
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nifi

type PropertyDescriptor struct {
	// The name of the property key
	Name string `json:"name"`
	// The display name of the property key, if different from the name
	DisplayName string `json:"displayName,omitempty"`
	// The description of what the property does
	Description string `json:"description,omitempty"`
	// A list of the allowable values for the property
	AllowableValues []PropertyAllowableValue `json:"allowableValues,omitempty"`
	// The default value if a user-set value is not specified
	DefaultValue string `json:"defaultValue,omitempty"`
	// Whether or not  the property is required for the component
	Required bool `json:"required,omitempty"`
	// Whether or not  the value of the property is considered sensitive (e.g., passwords and keys)
	Sensitive bool `json:"sensitive,omitempty"`
	// The scope of expression language supported by this property
	ExpressionLanguageScope string `json:"expressionLanguageScope,omitempty"`
	// The description of the expression language scope supported by this property
	ExpressionLanguageScopeDescription string `json:"expressionLanguageScopeDescription,omitempty"`
	TypeProvidedByValue *DefinedType `json:"typeProvidedByValue,omitempty"`
	// A regular expression that can be used to validate the value of this property
	ValidRegex string `json:"validRegex,omitempty"`
	// Name of the validator used for this property descriptor
	Validator string `json:"validator,omitempty"`
	// Whether or not the descriptor is for a dynamically added property
	Dynamic bool `json:"dynamic,omitempty"`
	ResourceDefinition *PropertyResourceDefinition `json:"resourceDefinition,omitempty"`
	// The dependencies that this property has on other properties
	Dependencies []PropertyDependency `json:"dependencies,omitempty"`
}
