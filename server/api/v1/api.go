package v1

import "github.com/sevenzx/eztodo/api/v1/internal"

type _API struct {
	User internal.UserApi
}

var API = new(_API)
