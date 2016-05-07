package netstat

import (
    "fmt"
    "os"
    "testing"
)

func TestNetstat(t *testing.T) {
    netstat, err := NetStatTCP()
    if err != nil {
        fmt.Printf("Error %s\n", err)
        os.Exit(1)
    }
    fmt.Printf("%-6s   %-24s    %-24s    %-14s  %s\n", "PROTO", "LOCAL ADDRESS", "FOREIGN ADDRESS", "STATE", "PID")
    for row := range netstat {
        fmt.Printf("%-6s   %-24s    %-24s    %-14s  %d\n",
            row.Proto,
            fmt.Sprintf("%s:%d", row.LocalAddr, row.LocalPort),
            fmt.Sprintf("%s:%d", row.RemoteAddr, row.RemotePort),
            getState(row.State),
            row.ProcessId)
    }
    netstat, err = NetStatUDP()
    if err != nil {
        fmt.Printf("Error %s\n", err)
        os.Exit(1)
    }
    for row := range netstat {
        fmt.Printf("%-6s   %-24s    %-24s    %-14s  %d\n",
            row.Proto,
            fmt.Sprintf("%s:%d", row.LocalAddr, row.LocalPort),
            "*.*",
            getState(row.State),
            row.ProcessId)
    }
}
