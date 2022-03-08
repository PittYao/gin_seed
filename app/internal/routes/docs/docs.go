// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
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
        "/ping": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    },
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "返回 server 相关信息，可以用于健康检查",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "x"
                ],
                "summary": "默认的 Ping 接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "you can set custom trace id in header",
                        "name": "trace_id",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/record/one": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stream one"
                ],
                "summary": "录制单个流",
                "parameters": [
                    {
                        "description": " ",
                        "name": "recordOneReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dao.RecordOneReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/record/one/page": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stream one"
                ],
                "summary": "分页查询",
                "parameters": [
                    {
                        "description": " ",
                        "name": "recordOneReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dao.RecordOnePageReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dao.RecordOnePageReq": {
            "type": "object",
            "required": [
                "page",
                "pageSize"
            ],
            "properties": {
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "total_count": {
                    "type": "integer"
                }
            }
        },
        "dao.RecordOneReq": {
            "type": "object",
            "required": [
                "rtspUrl"
            ],
            "properties": {
                "rtspUrl": {
                    "type": "string"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "响应码",
                    "type": "integer"
                },
                "data": {
                    "description": "响应数据"
                },
                "msg": {
                    "description": "响应消息",
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