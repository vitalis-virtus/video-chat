package models

type UriID struct {
	ID uint64 `uri:"id" binding:"required"`
}

type UriIDString struct {
	ID string `uri:"id" binding:"required"`
}
