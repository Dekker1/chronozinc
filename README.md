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

## TODO
- Remove all usage of `viper` from parsing and runtime packages
- Provide separate the `run` and `parse` commands
- MORE DOCUMENTATION!!!
