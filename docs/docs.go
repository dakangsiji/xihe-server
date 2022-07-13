// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/project": {
            "post": {
                "description": "create project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Create",
                "parameters": [
                    {
                        "description": "body of creating project",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.projectCreateModel"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/project/{id}": {
            "put": {
                "description": "update project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Update",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of project",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body of updating project",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.projectUpdateModel"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/user": {
            "put": {
                "description": "update user basic info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update",
                "responses": {}
            }
        }
    },
    "definitions": {
        "controller.projectCreateModel": {
            "type": "object",
            "properties": {
                "cover_id": {
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "inference": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "protocol": {
                    "type": "string"
                },
                "training": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "controller.projectUpdateModel": {
            "type": "object",
            "properties": {
                "cover_id": {
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "tags": {
                    "description": "json [] will be converted to []string",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "type": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
