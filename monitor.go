package slurmfs

import (
    "io"
    "os/exec"
)

// TODO: adapt https://gist.github.com/mxschmitt/6c07b5b97853f05455c3fdaf48b1a8b6
func monitor(cmd *exec.Cmd, cb func(string)error) {
    //restIn, _  := cmd.StdinPipe()
    stdout, err := cmd.StdoutPipe()
    cmd.Stderr = cmd.Stdout
    if err != nil {
        cb(err.Error())
        return
    }
    defer stdout.Close()

    err = cmd.Start()
    if err != nil {
        cb(err.Error())
        return
    }

    chunkSize := 1024
    buf := make([]byte, chunkSize)
    for {
        n, err := stdout.Read(buf)
        if err != nil && err != io.EOF {
            cb(err.Error())
            return
        }
        if err == io.EOF {
            break
        }
        //if n > 0 {
        cb( string(buf[:n]) )
    }
    // successful read of full stdout
}
