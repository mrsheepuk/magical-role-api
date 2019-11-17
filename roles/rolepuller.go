package roles

/*
RolePuller pulls roles from an underlying service.
*/
type RolePuller interface {
	ByNames(names ...string) (error, string[])
}
