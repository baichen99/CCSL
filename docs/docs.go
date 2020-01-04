// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-01-05 01:30:19.795281 +0800 CST m=+1.232747113

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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/users": {
            "get": {
                "description": "get users list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "List users",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "select from page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit number",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "order by field",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "asc",
                            "desc"
                        ],
                        "type": "string",
                        "description": "order by asc or desc",
                        "name": "orderBy",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "super",
                            "admin",
                            "user",
                            "learner"
                        ],
                        "type": "string",
                        "description": "filter type of user",
                        "name": "userType",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "search name of user",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "search username of user",
                        "name": "username",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "active",
                            "inactive"
                        ],
                        "type": "string",
                        "description": "filter state of user",
                        "name": "state",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.GetUsersListResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "create user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.UserCreateForm"
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "put": {
                "description": "update a user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update user",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "updated user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controllers.UserUpdateForm"
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "erroro message"
                },
                "message": {
                    "type": "string",
                    "example": "error title"
                }
            }
        },
        "controllers.GetUsersListResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.User"
                    }
                },
                "limit": {
                    "type": "integer",
                    "example": 10
                },
                "message": {
                    "type": "string",
                    "example": "success"
                },
                "page": {
                    "type": "integer",
                    "example": 1
                },
                "total": {
                    "type": "integer",
                    "example": 100
                }
            }
        },
        "controllers.UserCreateForm": {
            "type": "object",
            "required": [
                "name",
                "state",
                "userType",
                "username"
            ],
            "properties": {
                "avatar": {
                    "type": "string",
                    "example": "https://ccsl.shu.edu.cn/public/assets/default.png"
                },
                "name": {
                    "type": "string",
                    "example": "Adrian Duan"
                },
                "password": {
                    "type": "string",
                    "example": "p@ssw0rd"
                },
                "state": {
                    "type": "string",
                    "enum": [
                        "active",
                        "inactive"
                    ],
                    "example": "active"
                },
                "userType": {
                    "type": "string",
                    "enum": [
                        "super",
                        "admin",
                        "user",
                        "learner"
                    ],
                    "example": "admin"
                },
                "username": {
                    "type": "string",
                    "example": "adrianduan@icloud.com"
                }
            }
        },
        "controllers.UserUpdateForm": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string",
                    "example": "https://ccsl.shu.edu.cn/public/assets/default.png"
                },
                "name": {
                    "type": "string",
                    "example": "Adrian Duan"
                },
                "state": {
                    "type": "string",
                    "enum": [
                        "active",
                        "inactive"
                    ],
                    "example": "active"
                },
                "userType": {
                    "type": "string",
                    "enum": [
                        "super",
                        "admin",
                        "user",
                        "learner"
                    ],
                    "example": "admin"
                },
                "username": {
                    "type": "string",
                    "example": "adrianduan@icloud.com"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string",
                    "example": "https://ccsl.shu.edu.cn/public/assets/default.png"
                },
                "createdAt": {
                    "type": "string",
                    "example": "2000-12-30T00:00:00Z"
                },
                "id": {
                    "type": "string",
                    "format": "uuid",
                    "example": "550e8400-e29b-41d4-a716-446655440000"
                },
                "name": {
                    "type": "string",
                    "example": "Adrian Duan"
                },
                "state": {
                    "type": "string",
                    "example": "active"
                },
                "updatedAt": {
                    "type": "string",
                    "example": "2000-12-30T00:00:00Z"
                },
                "userType": {
                    "type": "string",
                    "example": "admin"
                },
                "username": {
                    "type": "string",
                    "example": "adrianduan@icloud.com"
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
	Host:        "localhost:8888",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "CCSL API",
	Description: "",
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
