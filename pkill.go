package pkill

import (
	"fmt"
	"os"
	"regexp"
	"sync"

	"github.com/mitchellh/go-ps"
)

type (
	syncPIDs struct {
		L    sync.Mutex
		PIDs []int
	}

	syncErrors struct {
		L      sync.Mutex
		Errors []error
	}
)

// Pkill functions like `pkill`, sending a signal to each process found by regular expression on their executable name,
// and returning a list of PIDs and any errors during execution. Note: Some processes may have been signaled even if
// this command errors; thus both return values may not be nil.
func Pkill(expr string, sig os.Signal) ([]int, error) {
	return pcmd(expr, sig)
}

// Pgrep functions like `pgrep`, returning a list of PIDs by regular expression on their executable name. Note: Both
// return values may not be nil.
func Pgrep(expr string) ([]int, error) {
	return pcmd(expr, nil)
}

func pcmd(expr string, sig os.Signal) ([]int, error) {
	re, err := regexp.CompilePOSIX(expr)
	if err != nil {
		return nil, err
	}

	procs, err := ps.Processes()
	if err != nil {
		return nil, err
	}

	pids := new(syncPIDs)
	errs := new(syncErrors)
	wg := new(sync.WaitGroup)
	wg.Add(len(procs))
	for _, proc := range procs {
		proc := proc
		go func() {
			defer wg.Done()
			if re.MatchString(proc.Executable()) {
				p, err := os.FindProcess(proc.Pid())
				if err != nil {
					errs.L.Lock()
					errs.Errors = append(errs.Errors, err)
					errs.L.Unlock()
					return
				}
				if sig != nil {
					if err := p.Signal(sig); err != nil {
						errs.L.Lock()
						errs.Errors = append(errs.Errors, fmt.Errorf("%d: %v", p.Pid, err))
						errs.L.Unlock()
						return
					}
				}
				pids.L.Lock()
				pids.PIDs = append(pids.PIDs, p.Pid)
				pids.L.Unlock()
			}
		}()
	}
	wg.Wait()

	if len(errs.Errors) > 0 {
		return pids.PIDs, fmt.Errorf("errors: %v", errs)
	}
	return pids.PIDs, nil
}
