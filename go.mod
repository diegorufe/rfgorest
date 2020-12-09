module rfgorest

go 1.15

replace rfgocore => E:/trabajo/repos/go/rfgocore

replace rfgodata => E:/trabajo/repos/go/rfgodata

require rfgocore v0.0.1

require (
	github.com/jinzhu/gorm v1.9.16
	github.com/mitchellh/mapstructure v1.4.0
	gorm.io/gorm v1.9.19
	rfgodata v0.0.1
)
