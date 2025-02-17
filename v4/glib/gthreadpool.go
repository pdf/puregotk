// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// The #GThreadPool struct represents a thread pool. It has three
// public read-only members, but the underlying struct is bigger,
// so you must not copy this struct.
type ThreadPool struct {
	Func Func

	UserData uintptr

	Exclusive bool
}

var xThreadPoolGetMaxIdleTime func() uint

// This function will return the maximum @interval that a
// thread will wait in the thread pool for new tasks before
// being stopped.
//
// If this function returns 0, threads waiting in the thread
// pool for new work are not stopped.
func ThreadPoolGetMaxIdleTime() uint {

	cret := xThreadPoolGetMaxIdleTime()
	return cret
}

var xThreadPoolGetMaxUnusedThreads func() int

// Returns the maximal allowed number of unused threads.
func ThreadPoolGetMaxUnusedThreads() int {

	cret := xThreadPoolGetMaxUnusedThreads()
	return cret
}

var xThreadPoolGetNumUnusedThreads func() uint

// Returns the number of currently unused threads.
func ThreadPoolGetNumUnusedThreads() uint {

	cret := xThreadPoolGetNumUnusedThreads()
	return cret
}

var xThreadPoolSetMaxIdleTime func(uint)

// This function will set the maximum @interval that a thread
// waiting in the pool for new tasks can be idle for before
// being stopped. This function is similar to calling
// g_thread_pool_stop_unused_threads() on a regular timeout,
// except this is done on a per thread basis.
//
// By setting @interval to 0, idle threads will not be stopped.
//
// The default value is 15000 (15 seconds).
func ThreadPoolSetMaxIdleTime(IntervalVar uint) {

	xThreadPoolSetMaxIdleTime(IntervalVar)

}

var xThreadPoolSetMaxUnusedThreads func(int)

// Sets the maximal number of unused threads to @max_threads.
// If @max_threads is -1, no limit is imposed on the number
// of unused threads.
//
// The default value is 2.
func ThreadPoolSetMaxUnusedThreads(MaxThreadsVar int) {

	xThreadPoolSetMaxUnusedThreads(MaxThreadsVar)

}

var xThreadPoolStopUnusedThreads func()

// Stops all currently unused threads. This does not change the
// maximal number of unused threads. This function can be used to
// regularly stop all unused threads e.g. from g_timeout_add().
func ThreadPoolStopUnusedThreads() {

	xThreadPoolStopUnusedThreads()

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xThreadPoolGetMaxIdleTime, lib, "g_thread_pool_get_max_idle_time")
	core.PuregoSafeRegister(&xThreadPoolGetMaxUnusedThreads, lib, "g_thread_pool_get_max_unused_threads")
	core.PuregoSafeRegister(&xThreadPoolGetNumUnusedThreads, lib, "g_thread_pool_get_num_unused_threads")
	core.PuregoSafeRegister(&xThreadPoolSetMaxIdleTime, lib, "g_thread_pool_set_max_idle_time")
	core.PuregoSafeRegister(&xThreadPoolSetMaxUnusedThreads, lib, "g_thread_pool_set_max_unused_threads")
	core.PuregoSafeRegister(&xThreadPoolStopUnusedThreads, lib, "g_thread_pool_stop_unused_threads")

}
