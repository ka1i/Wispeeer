package info

import "fmt"

func ShowBanner() {
	fmt.Printf(`> %s _                               
 _      __(_)________  ___  ___  ___  _____
| | /| / / / ___/ __ \/ _ \/ _ \/ _ \/ ___/
| |/ |/ / (__  ) /_/ /  __/  __/  __/ /    
|__/|__/_/____/ .___/\___/\___/\___/_/     
             /_/ %s
`, Services, Version.ToString())
}
