package slurmfs

import (
	"context"
	"os"
	"net"
	"time"
        "os/exec"
	"net/http"
	"path/filepath"

	"github.com/frobnitzem/go-p9p"
)

type SlurmServer struct {
	client Client
	//jobcache []...
	logfile *os.File
}

// Creating an http client from a UNIX socket
// https://gist.github.com/teknoraver/5ffacb8757330715bcbcc90e6d46ac74
type Client struct {
	http.Client
}

func UnixClient(path string) Client {
	return Client{ http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, _, addr string) (net.Conn, error) {
				return (&net.Dialer{}).DialContext(ctx, "unix", path)
			},
		},
	}}
}

func (s *SlurmServer) log(info string) error {
	_, err := s.logfile.WriteString(info)
	return err
}

func (s *SlurmServer) readlog(b []byte, off int64) (int, error) {
	return s.logfile.ReadAt(b, off)
}

func NewServer(ctx context.Context, path string) (*SlurmServer, error) {
    // may be fs.ModeDir
        err := os.MkdirAll(path, os.ModeDir | 0o700)
        if err != nil {
            return nil, err
        }

        // Create a logfile
	log, err := os.Create(filepath.Join(path,"log"))
        if err != nil {
            return nil, err
        }
        srv := SlurmServer{logfile: log}

        // Start up the slurmrestd on a socket

        sock := filepath.Join(path, "socket")
        // create if not exists (and capture its log messages)
        if _, err = os.Stat(sock); err != nil {
            //os.Remove(sock)
            restd := exec.Command("/usr/sbin/slurmrestd", "unix:" + sock)
            go monitor(restd, srv.log)
            time.Sleep(200*time.Millisecond)
        }

	//client := &http.Client{} (could use tcp...)
	client := UnixClient(sock)
        srv.client = client

	return &srv, nil
}

func (_ *SlurmServer) RequireAuth(_ context.Context) bool {
    return false
}
func (_ *SlurmServer) Auth(ctx context.Context,
    uname, aname string) (p9p.AuthFile, error) {
    return nil, nil
}
/*
func (s *SlurmServer) Attach(ctx context.Context, uname, aname string,
    af p9p.AuthFile) (p9p.Dirent, error) {
    return Root(s), nil
}

type RootEnt struct {
    fs *SlurmServer
    entries []string
}*/
