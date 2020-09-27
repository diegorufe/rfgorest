module rfgorest

go 1.13

replace rfgocore => E:/trabajo/repos/go/rfgocore

replace rfgodata => E:/trabajo/repos/go/rfgodata

require rfgocore v0.0.1

require (
	github.com/jinzhu/gorm v1.9.16
	gorm.io/gorm v1.9.19
	rfgodata v0.0.1
)
