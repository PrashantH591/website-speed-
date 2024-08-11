// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/reports": {
            "post": {
                "description": "Run a lighthouse audit to generate a report. The field ` + "`" + `raw_json` + "`" + ` contains the\nJSON output returned from lighthouse as a string.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a Lighthouse Report",
                "parameters": [
                    {
                        "description": "Lighthouse parameters to generate the report",
                        "name": "ReportRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.ReportRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.Report"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.AuditResult": {
            "type": "object",
            "properties": {
                "DisplayValue": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "numericUnit": {
                    "type": "string"
                },
                "numericValue": {
                    "type": "number"
                },
                "score": {
                    "type": "number"
                },
                "scoreDisplayMode": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "api.Report": {
            "type": "object",
            "properties": {
                "audit_results": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/api.AuditResult"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "form_factor": {
                    "description": "Optional parameter, possible values are desktop or mobile. If unset will default to desktop",
                    "type": "string",
                    "example": "desktop"
                },
                "id": {
                    "type": "string"
                },
                "location": {
                    "description": "Optional parameter, default location will be used if not set",
                    "type": "string",
                    "example": "australia-southeast1"
                },
                "performance_score": {
                    "type": "number"
                },
                "raw_json": {
                    "description": "RawJSON contains the lighthouse JSON result",
                    "type": "string"
                },
                "throughput_kbps": {
                    "description": "Optional parameter, by default will be set to 1000 if omitted",
                    "type": "integer",
                    "example": 50000
                },
                "url": {
                    "description": "Required parameter the URL of the website",
                    "type": "string",
                    "example": "https://www.google.com"
                }
            }
        },
        "api.ReportRequest": {
            "type": "object",
            "properties": {
                "form_factor": {
                    "description": "Optional parameter, possible values are desktop or mobile. If unset will default to desktop",
                    "type": "string",
                    "example": "desktop"
                },
                "location": {
                    "description": "Optional parameter, default location will be used if not set",
                    "type": "string",
                    "example": "australia-southeast1"
                },
                "throughput_kbps": {
                    "description": "Optional parameter, by default will be set to 1000 if omitted",
                    "type": "integer",
                    "example": 50000
                },
                "url": {
                    "description": "Required parameter the URL of the website",
                    "type": "string",
                    "example": "https://www.google.com"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Websu API",
	Description: "Run lighthouse as a service",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}