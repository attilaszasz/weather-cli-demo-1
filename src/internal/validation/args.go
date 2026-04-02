package validation

import "fmt"

func ValidateArgs(extraArgs []string) error {
	if len(extraArgs) > 0 {
		return fmt.Errorf("unexpected positional arguments: %v", extraArgs)
	}
	return nil
}
