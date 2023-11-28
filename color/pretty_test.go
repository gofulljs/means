package color

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestColorLogo(t *testing.T) {

	var meansText = fmt.Sprintf(`	__________________________________
    ____ ___  ___  ____ _____  _____
   / __ %c__ \/ _ \/ __ %c/ __ \/ ___/
  / / / / / /  __/ /_/ / / / (__  ) 
 /_/ /_/ /_/\___/\__,_/_/ /_/____/
___________________________________
`, '`', '`')

	err := ColorLogo(meansText, os.Stdout)
	assert.NoError(t, err)

}
