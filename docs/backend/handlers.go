package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"sort"
	"strings"
	"time"
)

// Deprecated
func ssiHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "<p>SSI handler (%s).</p>", r.URL.Path[1:])
	_, _ = fmt.Fprintf(w, inspectRequestHTML(r))
}

// Get some status info
func statusHandler(w http.ResponseWriter, r *http.Request) {
	var msg []string
	status := "ok"

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	err := updateReleasesStatus()
	if err != nil {
		msg = append(msg, err.Error())
		status = "error"
	}

	_ = json.NewEncoder(w).Encode(
		ApiStatusResponseType{
			Status:         status,
			Msg:            strings.Join(msg, " "),
			RootVersion:    getRootReleaseVersion(),
			RootVersionURL: VersionToURL(getRootReleaseVersion()),
			Multiwerf:      ReleasesStatus.Releases,
		})
}

// Redirect to default documentation version
func rootDocumentationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Accel-Redirect", fmt.Sprintf("/%v%v", VersionToURL(getRootReleaseVersion()), r.URL.RequestURI()))
}

// Handles request to /v<group>-<channel>/. E.g. /v1.2-beta/
func groupChannelHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_ = updateReleasesStatus()
	var version, URLToRedirect string
	var err error
	err, version = getVersionFromChannelAndGroup(&ReleasesStatus, vars["channel"], vars["group"])

	requestURI := r.URL.RequestURI()
	items := strings.Split(requestURI, "/")
	if len(items) > 1 {
		requestURI = strings.Join(items[2:], "/")
	} else {
		err = errors.New("can't construct URI for redirect")
	}

	if err == nil {
		URLToRedirect = fmt.Sprintf("/%v/%v", VersionToURL(version), requestURI)
		err = validateURL(fmt.Sprintf("https://%s%s", r.Host, URLToRedirect))
	}

	if err != nil {
		log.Printf("Error validating URL: %v, (validated - https://%s/%v/%v)", err.Error(), r.Host, VersionToURL(version), r.URL.RequestURI())
		URLToRedirect = fmt.Sprintf("/404.html")
	}
	w.Header().Set("X-Accel-Redirect", URLToRedirect)
}

func validateURL(URL string) (err error) {
	var resp *http.Response
	allowedStatusCodes := []int{200, 401}
	tries := 3
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 10 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       10 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	for {
		resp, err = client.Get(URL)
		if err == nil && (resp.StatusCode == 301 || resp.StatusCode == 302) {
			if len(resp.Header.Get("Location")) > 0 {
				URL = resp.Header.Get("Location")
			} else {
				tries = 0
			}
			tries--
		} else {
			tries = 0
		}
		if tries < 1 {
			break
		}
	}

	if err == nil {
		place := sort.SearchInts(allowedStatusCodes, resp.StatusCode)
		if place >= len(allowedStatusCodes) {
			err = errors.New(fmt.Sprintf("URL %s is not valid", URL))
		}
	}
	return
}

// Healthcheck handler
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// Get HTML content for /includes/topnav.html request
func topnavHandler(w http.ResponseWriter, r *http.Request) {
	_ = updateReleasesStatus()

	versionMenu := versionMenuType{
		VersionItems:          []versionMenuItems{},
		HTMLContent:           "",
		CurrentGroup:          "",
		CurrentChannel:        "",
		CurrentVersion:        "",
		CurrentVersionURL:     "",
		CurrentPageURL:        "",
		MenuDocumentationLink: "",
	}

	_ = versionMenu.getVersionMenuData(r, &ReleasesStatus)

	tplPath := "./root/"
	if strings.HasPrefix(r.Host, "ru.") {
		tplPath += "ru"
	} else {
		tplPath += "main"
	}

	tplPath += r.RequestURI
	tpl := template.Must(template.ParseFiles(tplPath))
	err := tpl.Execute(w, versionMenu)
	if err != nil {
		// Log error or maybe make some magic?
		log.Printf("Internal Server Error (template error), %v ", err.Error())
		http.Error(w, "Internal Server Error (template error)", 500)
	}
}

func serveFilesHandler(fs http.FileSystem) http.Handler {
	fsh := http.FileServer(fs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upath := r.URL.Path
		if !strings.HasPrefix(upath, "/") {
			upath = "/" + upath
			r.URL.Path = upath
		}
		upath = path.Clean(upath)
		if _, err := os.Stat(fmt.Sprintf("%v%s", fs, upath)); err != nil {
			if os.IsNotExist(err) {
				notFoundHandler(w, r)
				return
			}
		}
		fsh.ServeHTTP(w, r)
	})
}

// Redirect to root documentation if request not matches any location (override 404 response)
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Accel-Redirect", fmt.Sprintf("/404.html"))
}