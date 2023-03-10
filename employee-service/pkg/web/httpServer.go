package web

import (
	"crypto/tls"
	"errors"
	"log"
	"net"
	"net/http"
	"time"
)

// Set CORS middleware.
func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		switch origin := r.Header.Get("Origin"); origin {
		case "www.example.com", "example.com":
			w.Header().Set("Access-Control-Allow-Origin", "http://"+origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, ResponseType")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization")

		default:
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "HEADER, GET, POST, PUT, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
		}

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		h.ServeHTTP(w, r)
	})
}

// Run server
func Run(h http.Handler, p string) (err error) {

	httpStart := "http server listening on port:"
	httpFailed := "http server failed to listening on port:"

	log.Printf("%s %s \n", httpStart, p)

	s := &http.Server{
		// Addr optionally specifies the TCP address for the server to listen on,
		// in the form "host:port". If empty, ":http" (port 80) is used.
		// The service names are defined in RFC 6335 and assigned by IANA.
		// See net.Dial for details of the address format.
		Addr: ":" + p,
		// handler to invoke, http.DefaultServeMux if nil
		Handler: cors(h),
		// TLSConfig optionally provides a TLS configuration for use
		// by ServeTLS and ListenAndServeTLS. Note that this value is
		// cloned by ServeTLS and ListenAndServeTLS, so it's not
		// possible to modify the configuration with methods like
		// tls.Config.SetSessionTicketKeys. To use
		// SetSessionTicketKeys, use Server.Serve with a TLS Listener
		// instead.
		TLSConfig: &tls.Config{},
		// ReadTimeout is the maximum duration for reading the entire
		// request, including the body. A zero or negative value means
		// there will be no timeout.
		//
		// Because ReadTimeout does not let Handlers make per-request
		// decisions on each request body's acceptable deadline or
		// upload rate, most users will prefer to use
		// ReadHeaderTimeout. It is valid to use them both.
		ReadTimeout: 10 * time.Second,
		// ReadHeaderTimeout is the amount of time allowed to read
		// request headers. The connection's read deadline is reset
		// after reading the headers and the Handler can decide what
		// is considered too slow for the body. If ReadHeaderTimeout
		// is zero, the value of ReadTimeout is used. If both are
		// zero, there is no timeout.
		ReadHeaderTimeout: 10 * time.Second,
		// WriteTimeout is the maximum duration before timing out
		// writes of the response. It is reset whenever a new
		// request's header is read. Like ReadTimeout, it does not
		// let Handlers make decisions on a per-request basis.
		// A zero or negative value means there will be no timeout.
		WriteTimeout: 10 * time.Second,
		// IdleTimeout is the maximum amount of time to wait for the
		// next request when keep-alives are enabled. If IdleTimeout
		// is zero, the value of ReadTimeout is used. If both are
		// zero, there is no timeout.
		IdleTimeout: 10 * time.Second,
		// MaxHeaderBytes controls the maximum number of bytes the
		// server will read parsing the request header's keys and
		// values, including the request line. It does not limit the
		// size of the request body.
		// If zero, DefaultMaxHeaderBytes is used.
		MaxHeaderBytes: 1 << 20,
		// TLSNextProto optionally specifies a function to take over
		// ownership of the provided TLS connection when an ALPN
		// protocol upgrade has occurred. The map key is the protocol
		// name negotiated. The Handler argument should be used to
		// handle HTTP requests and will initialize the Request's TLS
		// and RemoteAddr if not already set. The connection is
		// automatically closed when the function returns.
		// If TLSNextProto is not nil, HTTP/2 support is not enabled
		// automatically.
		TLSNextProto: map[string]func(*http.Server, *tls.Conn, http.Handler){},
		// ConnState specifies an optional callback function that is
		// called when a client connection changes state. See the
		// ConnState type and associated constants for details.
		ConnState: func(net.Conn, http.ConnState) {
		},
		// ErrorLog specifies an optional logger for errors accepting
		// connections, unexpected behavior from handlers, and
		// underlying FileSystem errors.
		// If nil, logging is done via the log package's standard logger.
		ErrorLog: &log.Logger{},
	}

	if err = s.ListenAndServe(); err != nil {
		return errors.New(httpFailed)
	}

	return nil
}
