package runtime

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/jjdekker/chronozinc/settings"
	"github.com/spf13/viper"
)

// RunAll runs every instance on every solver
func RunAll(solvers []settings.Solver, instances []settings.Instance) {
	work := make(chan func())
	wait := govern(work)

	for i := range solvers {
		for j := range instances {
			work <- func() { RunInstance(&solvers[i], &instances[j]) }
		}
	}
	close(work)

	wait.Wait()
}

func govern(work <-chan func()) *sync.WaitGroup {
	var wg sync.WaitGroup
	procs := viper.GetInt("processes")

	wg.Add(procs)
	for i := 0; i < procs; i++ {
		go func() {
			for f := range work {
				f()
			}
			wg.Done()
		}()
	}

	return &wg
}

// RunInstance runs an instance on a Solver using mzn-fzn
func RunInstance(solver *settings.Solver, instance *settings.Instance) {
	args := []string{
		"--solver", solver.Binary,
		"--flatzinc-flag", solver.Flags,
		instance.Model,
	}
	if solver.Globals != "" {
		args = append(args, "--globals-dir", solver.Globals)
	}
	if instance.Data != "" {
		args = append(args, "--data", instance.Data)
	}
	args = append(args, strings.Split(viper.GetString("flags"), " ")...)
	proc := exec.Command(viper.GetString("mznfzn"), args...)

	if out, err := proc.CombinedOutput(); err != nil {
		log.Printf("Instance %s ended with error %s", instance, err)
	} else {
		path := instance.OutPath(solver.Name)
		os.MkdirAll(filepath.Dir(path), os.ModePerm)
		err := ioutil.WriteFile(path, out, 0644)
		if err != nil {
			log.Printf("Saving results for instance %s ended with error %s",
				instance, err)
		}
	}
}
