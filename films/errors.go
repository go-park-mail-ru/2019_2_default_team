package films

import "fmt"

type FilmNotFoundError struct {
	Field string
}

func (e FilmNotFoundError) Error() string {
	return fmt.Sprintf("no film with this %v found", e.Field)
}
