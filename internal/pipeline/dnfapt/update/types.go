package update

type PipelineData struct {
	localHost bool   // apply to local host or remote host
	hostType  string // Vm or container
	OsFamily  string // Rhel, Debian, etc.
	Err       error  // error if any
}
