package commands

import (
	"strings"
)

type CLIArgs struct {
	Flags       map[string]string
	Positionals []string
}

func ParseArgs(args []string) CLIArgs {
	result := CLIArgs{
		Flags:       make(map[string]string),
		Positionals: []string{},
	}

	for i := 0; i < len(args); i++ {
		arg := args[i]

		if strings.HasPrefix(arg, "--") {
			// handle --key=value
			parts := strings.SplitN(arg[2:], "=", 2)
			if len(parts) == 2 {
				result.Flags[parts[0]] = parts[1]
			} else if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
				result.Flags[parts[0]] = args[i+1]
				i++
			} else {
				result.Flags[parts[0]] = "true"
			}
		} else if strings.HasPrefix(arg, "-") {
			// handle -k value
			key := arg[1:]
			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
				result.Flags[key] = args[i+1]
				i++
			} else {
				result.Flags[key] = "true"
			}
		} else {
			// positional arg
			result.Positionals = append(result.Positionals, arg)
		}
	}

	return result
}
