package http

import (
	"github.com/feiyuw/nux"
	"net/http"
)

func configMemoryRoutes() {
	http.HandleFunc("/page/memory", func(w http.ResponseWriter, r *http.Request) {
		mem, err := nux.MemInfo()
		if err != nil {
			RenderMsgJson(w, err.Error())
			return
		}

		var memFree uint64
		if mem.MemAvailable == 0 {
			memFree = mem.MemFree + mem.Buffers + mem.Cached
		} else {
			memFree = mem.MemAvailable
		}
		memUsed := mem.MemTotal - memFree
		var t uint64 = 1024 * 1024
		RenderDataJson(w, []interface{}{mem.MemTotal / t, memUsed / t, memFree / t})
	})

	http.HandleFunc("/proc/memory", func(w http.ResponseWriter, r *http.Request) {
		mem, err := nux.MemInfo()
		if err != nil {
			RenderMsgJson(w, err.Error())
			return
		}

		var memFree uint64
		if mem.MemAvailable == 0 {
			memFree = mem.MemFree + mem.Buffers + mem.Cached
		} else {
			memFree = mem.MemAvailable
		}
		memUsed := mem.MemTotal - memFree

		RenderDataJson(w, map[string]interface{}{
			"total": mem.MemTotal,
			"free":  memFree,
			"used":  memUsed,
		})
	})
}
