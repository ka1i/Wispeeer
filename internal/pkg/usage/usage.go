package usage

import (
	"fmt"

	"github.com/wispeeer/wispeeer/pkg/info"
)

// app usage menual
func Usage(extraUsage string) {
	fmt.Printf(`Usage: %s

     ------- < Commands Arguments > -------
argument:
%s
options:
  -h, help          Show help message. 
  -v, version       Show the app version.
For more help, 
you can use '%s help' for the detailed information
or you can check the repo: https://github.com/ka1i/wispeeer.git  
`, info.MicroService, extraUsage, info.MicroService)
}
