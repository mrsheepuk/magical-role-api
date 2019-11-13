package roles

/*
RolePuller pulls roles from an underlying service.
*/
type RolePuller interface {
	ByName(name string) (error, string[])
	ByNames(names []string) (error, string[])
}
