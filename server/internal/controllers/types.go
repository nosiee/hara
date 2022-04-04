package controllers

import "time"

var JWTExp = time.Now().Add(1 * 365 * 24 * time.Hour).Unix()
