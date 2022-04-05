package controllers

import "time"

const HS512KeySize = 64

var JWTExp = time.Now().Add(1 * 365 * 24 * time.Hour).Unix()
