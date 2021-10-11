package resources

import (
	"net"
	"net/http"
)

// Heartbeat dials a non-existant IP using UDP and returns the outbound
// IP if no error occurs. Though the target IP does not need to be an
// existing target, there does need to be a internet connection.
func (env *Env) Heartbeat(w http.ResponseWriter, r *http.Request) {
	conn, err := net.Dial("udp", "192.168.3.1:80")
	if err != nil {
		JSONIFY(w, http.StatusInternalServerError, JSON{"error": err.Error()})
	} else {
		defer conn.Close()
		local := conn.LocalAddr().(*net.UDPAddr)
		JSONIFY(w, http.StatusOK, JSON{"ip": local.IP})
	}
}
