package crane

// add
// set by (id)
// delete
// list
type Limits struct {
	Total     int // limit of the instances
	EachTag   int
	EachLabel int
}

type Requires struct {
	Memory int // at least
}

type Job struct {
	Id       string
	Name     string
	Tags     []string
	Labels   map[string]string
	Envs     map[string]string
	Image    string
	Nodes    []string // #ALL#, #Tag# #Label:value#, Id
	Limits   Limits
	Requires Requires
	Volumes  map[string]string
}

// Scheduler arranges jobs into an
// appropriate sequence and proper host.
type Scheduler struct {

}

func (k *Scheduler) Add() {
	
}