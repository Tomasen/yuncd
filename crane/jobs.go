package crane

// add
// set by (id)
// delete 
// list

struct Job {
Id  string
Name string
Tags []string
Labels map[string]string
Envs map[string]string
Image string
Nodes []string // #ALL#, #Tag# #Label:value#, Id
Limits struct{
Total int // limit of the instances
EachTag int
EachLabel int
}
Requires struct {
Memory int // at least
}
Volumes map[string]string
}

struct Manager {
}

func m (*Manager) add () {
	
}

func m (*Manager) set () {
	
}

func m (*Manager) del () {
	
}

func m (*Manager) list () {
	
}