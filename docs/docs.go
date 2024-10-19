// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/v1/auth/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.LoginResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/landing/{code}": {
            "get": {
                "description": "Landing",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Landing"
                ],
                "summary": "Landing",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Merchant Code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/landing.LandingResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/landing/{code}/banners": {
            "get": {
                "description": "Get Landing Banners",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Landing"
                ],
                "summary": "Get Landing Banners",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/landing.GetLandingBannersResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/merchant/{merchant_id}/additional-fees": {
            "get": {
                "description": "Get merchant active additional fees",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Merchant"
                ],
                "summary": "Get Merchant Active Additional Fees",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Merchant ID",
                        "name": "merchant_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/merchant.GetMerchantActiveAdditionalFeesResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/merchant/{merchant_id}/payment-types": {
            "get": {
                "description": "Get merchant active payment types",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Merchant"
                ],
                "summary": "Get Merchant Active Payment Types",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Merchant ID",
                        "name": "merchant_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/merchant.GetMerchantActivePaymentTypesResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/order": {
            "post": {
                "description": "Create order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Create Order",
                "parameters": [
                    {
                        "description": "Create order request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order.CreateOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/order.CreateOrderResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/order/:order_id/push-payment-event": {
            "get": {
                "description": "Push payment event to client",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Push Payment Event",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order ID",
                        "name": "order_id",
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
        "/api/v1/order/pay": {
            "post": {
                "description": "Pay order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Pay Order",
                "parameters": [
                    {
                        "description": "Pay order request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order.PayOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/order.PayOrderResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/order/payment-callback": {
            "post": {
                "description": "Payment callback from payment gateway",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Payment Callback",
                "parameters": [
                    {
                        "description": "Payment callback request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order.PaymentCallbackRequest"
                        }
                    }
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
        "auth.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "auth.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "landing.Banner": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "end_date": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "start_date": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "landing.GetLandingBannersResponse": {
            "type": "object",
            "properties": {
                "banners": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/landing.Banner"
                    }
                }
            }
        },
        "landing.Item": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "item_variants": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/landing.ItemVariant"
                    }
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "priority": {
                    "type": "integer"
                }
            }
        },
        "landing.ItemCategory": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/landing.Item"
                    }
                },
                "name": {
                    "type": "string"
                },
                "priority": {
                    "type": "integer"
                }
            }
        },
        "landing.ItemVariant": {
            "type": "object",
            "properties": {
                "item_variant_options": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/landing.ItemVariantOption"
                    }
                },
                "max_select": {
                    "type": "integer"
                },
                "min_select": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "landing.ItemVariantOption": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                }
            }
        },
        "landing.LandingResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "item_categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/landing.ItemCategory"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "merchant.AdditionalFee": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "fee": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "integer"
                }
            }
        },
        "merchant.GetMerchantActiveAdditionalFeesResponse": {
            "type": "object",
            "properties": {
                "additional_fees": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/merchant.AdditionalFee"
                    }
                }
            }
        },
        "merchant.GetMerchantActivePaymentTypesResponse": {
            "type": "object",
            "properties": {
                "payment_types": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/merchant.PaymentType"
                    }
                }
            }
        },
        "merchant.PaymentType": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "integer"
                }
            }
        },
        "order.CreateOrderRequest": {
            "type": "object",
            "required": [
                "guest",
                "items",
                "merchant_id",
                "total_price"
            ],
            "properties": {
                "guest": {
                    "$ref": "#/definitions/order.Guest"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/order.Item"
                    }
                },
                "merchant_id": {
                    "type": "integer"
                },
                "order_type": {
                    "type": "integer"
                },
                "payment_method": {
                    "type": "integer"
                },
                "table_id": {
                    "type": "integer"
                },
                "total_price": {
                    "type": "integer"
                }
            }
        },
        "order.CreateOrderResponse": {
            "type": "object",
            "properties": {
                "due_time": {
                    "type": "integer"
                },
                "order_code": {
                    "type": "string"
                },
                "order_id": {
                    "type": "integer"
                },
                "payment_qr": {
                    "type": "string"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "order.Guest": {
            "type": "object",
            "required": [
                "name",
                "total_person"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "total_person": {
                    "type": "integer"
                }
            }
        },
        "order.Item": {
            "type": "object",
            "required": [
                "amount",
                "item_id"
            ],
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "item_id": {
                    "type": "integer"
                },
                "note": {
                    "type": "string"
                },
                "variants": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/order.ItemVariant"
                    }
                }
            }
        },
        "order.ItemVariant": {
            "type": "object",
            "properties": {
                "option_ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "variant_id": {
                    "type": "integer"
                }
            }
        },
        "order.PayOrderRequest": {
            "type": "object",
            "required": [
                "order_id"
            ],
            "properties": {
                "order_id": {
                    "type": "integer"
                }
            }
        },
        "order.PayOrderResponse": {
            "type": "object",
            "properties": {
                "order_code": {
                    "type": "string"
                },
                "order_id": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "order.PaymentCallbackRequest": {
            "type": "object",
            "properties": {
                "acquirer": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "fraud_status": {
                    "type": "string"
                },
                "gross_amount": {
                    "type": "string"
                },
                "issuer": {
                    "type": "string"
                },
                "merchant_id": {
                    "type": "string"
                },
                "order_id": {
                    "type": "string"
                },
                "payment_type": {
                    "type": "string"
                },
                "settlement_time": {
                    "type": "string"
                },
                "signature_key": {
                    "type": "string"
                },
                "status_code": {
                    "type": "string"
                },
                "status_message": {
                    "type": "string"
                },
                "transaction_id": {
                    "type": "string"
                },
                "transaction_status": {
                    "type": "string"
                },
                "transaction_time": {
                    "type": "string"
                },
                "transaction_type": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}