package crane

// add
// set by (id)
// delete
// list

type Host struct {
	Id     string            `db:"host_id, primarykey, autoincrement"`
	Addr   string            `db:"addr"`
	CACert string            `db:"cacert"`
	Cert   string            `db:"cert"`
	Key    string            `db:"key"`
	Tags   []string          `db:"tags"`
	Labels map[string]string `db:"labels"`
	Envs   map[string]string `db:"envs"`
}

// Monitor observing, checking,
// keeping a continuous record of
// host list.
type Monitor struct {
}

func (op *Monitor) Add() {
	
}

func (op *Monitor) Set() {

}

func (op *Monitor) Del() {

}

func (op *Monitor) List() {

}
