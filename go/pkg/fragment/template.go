package fragment

type Template struct {
	FragmentID ID
	Params     map[string]ID
	Slots      map[string]ID
}
