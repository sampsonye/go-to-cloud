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
        "/api/configure/artifact": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "制品仓库配置",
                "tags": [
                    "Configure"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/artifact.Artifact"
                            }
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "制品仓库配置",
                "tags": [
                    "Configure"
                ],
                "parameters": [
                    {
                        "description": "Request",
                        "name": "ContentBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/artifact.Artifact"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/configure/artifact/bind": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "制品仓库配置",
                "tags": [
                    "Configure"
                ],
                "parameters": [
                    {
                        "description": "Request",
                        "name": "ContentBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/artifact.Artifact"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/configure/artifact/testing": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "制品仓库配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Configure"
                ],
                "parameters": [
                    {
                        "description": "Request",
                        "name": "ContentBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/artifact.Testing"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/configure/artifact/{id}": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "制品仓库配置",
                "tags": [
                    "Configure"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ArtifactRepo.ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/artifact.Artifact"
                            }
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "制品仓库配置",
                "tags": [
                    "Configure"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ArtifactRepo.ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/configure/coderepo": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "代码仓库配置",
                "tags": [
                    "Configure"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/scm.Scm"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "代码仓库配置",
                "tags": [
                    "Configure"
                ],
                "parameters": [
                    {
                        "description": "Request",
                        "name": "ContentBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/scm.Scm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/configure/coderepo/bind": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "代码仓库配置",
                "tags": [
                    "Configure"
                ],
                "parameters": [
                    {
                        "description": "Request",
                        "name": "ContentBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/scm.Scm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/configure/coderepo/testing": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "代码仓库配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Configure"
                ],
                "parameters": [
                    {
                        "description": "Request",
                        "name": "ContentBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/scm.Testing"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/configure/coderepo/{id}": {
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "代码仓库配置",
                "tags": [
                    "Configure"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "CodeRepo.ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/configure/deploy/k8s": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "k8s环境配置",
                "tags": [
                    "Configure"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/k8s.K8s"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "k8s仓库配置",
                "tags": [
                    "Configure"
                ],
                "parameters": [
                    {
                        "description": "Request",
                        "name": "ContentBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/k8s.K8s"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/configure/deploy/k8s/bind": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "k8s仓库配置",
                "tags": [
                    "Configure"
                ],
                "parameters": [
                    {
                        "description": "Request",
                        "name": "ContentBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/k8s.K8s"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/configure/deploy/k8s/testing": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "k8s仓库配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Configure"
                ],
                "parameters": [
                    {
                        "description": "Request",
                        "name": "ContentBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/k8s.Testing"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/configure/deploy/k8s/{id}": {
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "k8s仓库配置",
                "tags": [
                    "Configure"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "K8sRepo.ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/login": {
            "post": {
                "description": "登录",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "description": "Login Model",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/projects/coderepo": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "列出当前账户已绑定的SCM平台及可见的代码仓库",
                "tags": [
                    "Projects"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/project.CodeRepoGroup"
                            }
                        }
                    }
                }
            }
        },
        "/api/projects/list": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "查看项目信息",
                "tags": [
                    "Projects"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/project.DataModel"
                            }
                        }
                    }
                }
            }
        },
        "/api/user/info": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "查看用户信息",
                "tags": [
                    "User"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/user/logout": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "注销登录",
                "tags": [
                    "User"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "artifact.Artifact": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "isSecurity": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "orgLites": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/artifact.OrgLite"
                    }
                },
                "orgs": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "password": {
                    "type": "string"
                },
                "remark": {
                    "type": "string"
                },
                "type": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "artifact.OrgLite": {
            "type": "object",
            "properties": {
                "orgId": {
                    "type": "integer"
                },
                "orgName": {
                    "type": "string"
                }
            }
        },
        "artifact.Testing": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "isSecurity": {
                    "type": "boolean"
                },
                "password": {
                    "type": "string"
                },
                "type": {
                    "type": "integer"
                },
                "url": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "k8s.K8s": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "kubeconfig": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "orgLites": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/k8s.OrgLite"
                    }
                },
                "orgs": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "remark": {
                    "type": "string"
                },
                "serverVersion": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "k8s.OrgLite": {
            "type": "object",
            "properties": {
                "orgId": {
                    "type": "integer"
                },
                "orgName": {
                    "type": "string"
                }
            }
        },
        "k8s.Testing": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "kubeconfig": {
                    "type": "string"
                }
            }
        },
        "models.LoginModel": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "project.CodeRepoGroup": {
            "type": "object",
            "properties": {
                "host": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "label": {
                    "type": "string"
                },
                "options": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/project.GitSources"
                    }
                }
            }
        },
        "project.DataModel": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "remark": {
                    "type": "string"
                }
            }
        },
        "project.GitSources": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "label": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "scm.OrgLite": {
            "type": "object",
            "properties": {
                "orgId": {
                    "type": "integer"
                },
                "orgName": {
                    "type": "string"
                }
            }
        },
        "scm.Scm": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "isPublic": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "orgLites": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/scm.OrgLite"
                    }
                },
                "orgs": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "origin": {
                    "type": "integer"
                },
                "remark": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "scm.Testing": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "isPublic": {
                    "type": "boolean"
                },
                "origin": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
