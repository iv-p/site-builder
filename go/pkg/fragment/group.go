package fragment

type Outlet struct {
	ID         InstanceID
	Name       string
	FragmentID InstanceID
	SlotName   string
}

type Group struct {
	RootFragment InstanceID
	Outlets      map[string]Outlet
	Data         map[string]interface{}
}
