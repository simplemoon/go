/*
* Auth: yuanzp
* Mail: 1034889856@qq.com
* Create Time: 2022-11-01 00:20:25
 */

package test

type args struct {
	a int    // a
	m string // m
}

func (a *args) get() string {
	return a.m
}

func (a *args) set(m string) {
	a.m = m
}
