package wotoRaw

// all users permission.
const (
	PermissionNormalUser UserPermission = iota
	PermissionSpecial
	PermissionAdmin
	PermissionDeveloper
	PermissionOwner
)

// prefixes used to generate unique-id for different types.
const (
	MediaGenreElementPrefix = "gEl-"
	UniqueIdInnerSeparator  = "_"
)
