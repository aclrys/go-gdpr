package vendorlist

import (
	"github.com/aclrys/go-gdpr/api"
)

// Copying from API for backwards compatibility

// VendorList is a backward compatible interface of api.VendorList
type VendorList interface {
	api.VendorList
}

// Vendor is a backward compatible interface of api.Vendor
type Vendor interface {
	api.Vendor
}
