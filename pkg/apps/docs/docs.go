// Code generated by swaggo/swag. DO NOT EDIT
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
        "/ping": {
            "get": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Other"
                ],
                "summary": "ping",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/accounts": {
            "post": {
                "description": "create account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "create account",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateAccountReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/resp.successResp"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/resp.errorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/resp.errorResp"
                        }
                    }
                }
            }
        },
        "/accounts/verify": {
            "post": {
                "description": "verify account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "verify account",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.VerifyAccountReq"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/resp.successResp"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/resp.errorResp"
                        }
                    },
                    "404": {
                        "description": "Resource Not Found",
                        "schema": {
                            "$ref": "#/definitions/resp.errorResp"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/resp.errorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/resp.errorResp"
                        }
                    }
                }
            }
        },
    },
    "definitions": {
        "models.CreateAccountReq": {
            "type": "object",
            "required": [
                "name",
                "password"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 1
                },
                "password": {
                    "type": "string",
                    "pattern": "^(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])[a-zA-z0-9]{8,32}$"
                },
            }
        },
        "models.VerifyAccountReq": {
            "type": "object",
            "required": [
                "name",
                "password"
            ],
            "properties": {
                "name": {
                    "type": "string",
                },
                "password": {
                    "type": "string",
                },
            }
        },
        "resp.errorResp": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean",
                    "default": false
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "resp.successResp": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean",
                    "default": true
                },
                "message": {
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
	Title:            "Senao Prestest API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
