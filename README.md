# ChronoZinc
Small wrapper for `mzn-fzn` to benchmark MiniZinc models

## Usage
```
chronozinc [config]* [models]* [data]*
```
Files are matched on extension:
- Models: `.mzn`
- Data: `.dzn`
- Configuration: `.json`, `.toml`, `.yml` (viper extensions)

## Installation

`go get -u github.com/jjdekker/chronozinc`

## Sample Configuration

The following configuration can be used to time a model using `flatzinc`,
`fzn-gcode` and `fzn-or-tools`. Running is as easy as: `chronozinc config.toml
model.mzn data1.dzn data2.dzn`.

Tip: Remove the parameters line and place it in
`$HOME/.config/chronozinc/settings.toml` or `/etc/chronozinc/settings.toml`. You
can now make project specific settings files while using the solvers settings as
declared in this file.

```toml
parameters = ["time", "time_ms", "solvetime", "solvetime_ms", "status"]
processes = 2

[solvers.flatzinc]
	binary = "flatzinc"

[solvers.gecode]
	binary = "fzn-gecode"
	globals = "gecode"
	override_flags = "-s"
	[solvers.gecode.extractors]
		time = "runtime:\\s*(?P<result>\\d+.\\d+)\\s"
		solvetime = "solvetime:\\s*(?P<result>\\d+.\\d+)\\s"
		solutions = "solutions:\\s*(?P<result>\\d+)"
		variables = "variables:\\s*(?P<result>\\d+)"
		propagators = "propagators:\\s*(?P<result>\\d+)"
		propagations = "propagations:\\s*(?P<result>\\d+)"
		nodes = "nodes:\\s*(?P<result>\\d+)"
		failures = "failures:\\s*(?P<result>\\d+)"
		restarts = "restarts:\\s*(?P<result>\\d+)"
		peak_depth = "peak depth:\\s*(?P<result>\\d+)"

[solvers.or_tools]
	binary = "fzn-or-tools"
	globals = "or_tools_cp"
	[solvers.or_tools.extractors]
		time_ms = "total runtime:\\s*(?P<result>\\d+)\\s"
		buildtime_ms = "build time:\\s*(?P<result>\\d+)\\s"
		solvetime_ms = "solve time:\\s*(?P<result>\\d+)\\s"
		solutions = "solutions:\\s*(?P<result>\\d+)"
		constraints = "constraints:\\s*(?P<result>\\d+)"
		propagations = "normal propagations:\\s*(?P<result>\\d+)"
		delayed_propagations = "delayed propagations:\\s*(?P<result>\\d+)"
		branches = "branches:\\s*(?P<result>\\d+)"
		failures = "failures:\\s*(?P<result>\\d+)"
		memory = "memory:\\s*(?P<result>\\d+.\\d+)\\s" # in MB


[extractors]
	time = "overall time\\s*(?P<result>\\d+\\.\\d+)\\s*s"

[matchers.status]
	COMPLETE = "=========="
	UNSAT = "=====UNSATISFIABLE====="
	TIMEOUT = "TIMEOUT"
	UNKNOWN = "=====UNKNOWN====="
	SAT = "----------"
```

## TODO
- Remove all usage of `viper` from parsing and runtime packages
- Provide separate the `run` and `parse` commands
- Internal data transformation
- MORE DOCUMENTATION!!!
