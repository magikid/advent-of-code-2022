package comms

type Datastream struct {
	DatastreamBuffer             string
	startOfPacketMarkerLocation  int
	startOfMessageMarkerLocation int
}
