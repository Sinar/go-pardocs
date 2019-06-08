package pardocs // import "github.com/Sinar/go-pardocs"
import (
	"log"

	"github.com/Sinar/go-pardocs/internal/hansard"
)

type ParliamentDocs struct {
	conf Configuration
}

// CommandMode specifies the operation being executed.
type CommandMode int

// The available commands.
const (
	PLAN CommandMode = iota
	SPLIT
	RESET
)

// Configuration of a Context.
type Configuration struct {

	// Parliament Session Label
	ParliamentSession string

	// Hansard Type
	HansardType hansard.HansardType

	// ./raw + ./data folders are assumed to be relative to this dir
	WorkingDir string

	// Source PDF can be anywhere; maybe make it a Reader to be read direct from S3?
	SourcePDFPath string

	// Command being executed.
	Cmd CommandMode
}

func (pd *ParliamentDocs) Plan() {
	log.Println("In Plan ..")
}

func (pd *ParliamentDocs) Split() {
	log.Println("In Split ..")
}

func (pd *ParliamentDocs) Reset() {
	log.Println("In Reset ...")
	// Clean up plan
	// Clean up split pages folder
	// Clean up merged pages location
}
