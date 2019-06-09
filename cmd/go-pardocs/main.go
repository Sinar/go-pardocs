package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Sinar/go-pardocs/internal/hansard"

	"github.com/Sinar/go-pardocs"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/subcommands"
)

func main() {
	log.Println("Welcome to Sinar Project go-pardocs!")

	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&splitCmd{}, "Execute Split")
	subcommands.Register(&planCmd{}, "Planning Split")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}

// ============ PLAN command ===========
type planCmd struct {
	sessionLabel string
	hansardType  string
	workingDir   string
}

func (*planCmd) Name() string     { return "plan" }
func (*planCmd) Synopsis() string { return "Plan the Parliament Doc Splitting .." }
func (*planCmd) Usage() string {
	return `plan -session <name> -type <L|BL> <sourcePDFPath>
Examplse:
	./go-pardocs plan -session par14sesi1 -type L ./raw/Lisan/JDR12032019.pdf
	./go-pardocs plan -session par13sesi3 -type L ./raw/Lisan/JWP DR 161018.pdf
`
}
func (p *planCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.sessionLabel, "session", "", "Parliament Session Name e.g par14sesi1")
	f.StringVar(&p.hansardType, "type", "", "HansardType: [L|BL] for Lisan/BukanLisan")
	//f.StringVar(&p.workingDir, "dir", ".", "Where raw + data stored; e.g. /tmp")
}
func (p *planCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	log.Println("In plan Execute ..")
	//spew.Dump(p)

	// TODO: Refactor? so checks all;  or should it be fast exit?
	if p.sessionLabel == "" {
		// default means it is NOT set correctly!
		fmt.Println("session REQUIRED!!")
		fmt.Println(p.Usage())
		return subcommands.ExitUsageError
	}
	sessionLabel := p.sessionLabel
	// For string, not needed; we can check the default values ..
	if !isFlagPassed(f, "type") {
		fmt.Println("type REQUIRED!!")
		fmt.Println(p.Usage())
		return subcommands.ExitUsageError
	}
	//for _, unsetFlag := range UnsetFlags(f) {
	//	fmt.Println("UNSET: ", unsetFlag.Name)
	//}

	// Check if the values are correct or not
	//if p.hansardType != "L" && p.hansardType != "BL" {
	//	fmt.Println("VALID HANSARDTYPE: L or BL")
	//	fmt.Println(p.Usage())
	//	return subcommands.ExitUsageError
	//}

	var hansardType hansard.HansardType
	switch p.hansardType {
	case "L":
		hansardType = hansard.HANSARD_SPOKEN
	case "BL":
		hansardType = hansard.HANSARD_WRITTEN
	default:
		fmt.Println("VALID HANSARDTYPE: L or BL")
		fmt.Println(p.Usage())
		return subcommands.ExitUsageError
	}

	if f.NArg() == 0 {
		// shoudl validate if real existing file?
		fmt.Println("Need VALID SourcePDFPath!!!")
		fmt.Println(p.Usage())
		return subcommands.ExitUsageError
	}

	if f.NArg() > 1 {
		// shoudl validate if real existing file?
		fmt.Println("Too many Args!!!!")
		fmt.Println(p.Usage())
		return subcommands.ExitUsageError
	}

	sourcePDFPath := f.Args()[0]
	log.Println("SourcePDFPath: ", sourcePDFPath)

	conf := pardocs.Configuration{sessionLabel, hansardType, p.workingDir,
		sourcePDFPath, pardocs.PLAN}
	// DEBUG
	//spew.Dump(conf)
	// Detect the cover page and suggest label names?

	parDoc := pardocs.ParliamentDocs{conf}
	// Execute the plan .. should catch errors  with xerrors :P
	parDoc.Plan()
	// Print out the location of the plan to be reviewed?
	// Suggest any changes; automatic anomaly checks? strange odds rule; not in order?

	return subcommands.ExitSuccess
}

// =============== SPLIT command =================
type splitCmd struct{}

func (*splitCmd) Name() string     { return "split" }
func (*splitCmd) Synopsis() string { return "Splitting the plan .." }
func (*splitCmd) Usage() string {
	return `split
... :( ...
`
}
func (p *splitCmd) SetFlags(f *flag.FlagSet) {
	//f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}
func (p *splitCmd) Execute(_ context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {

	// Check pre-reqs that the plan exists; otherwise suggest to run plan?
	// or maybe even automatically run it ?

	// Setup the configuration ..

	// Slices of the args for the flag ..
	spew.Println(f.Args())

	// Nothign here ..
	//spew.Dump(args)
	return subcommands.ExitFailure
}

// Helper functions; is it really needed?
// https://stackoverflow.com/questions/35809252/check-if-flag-was-provided-in-go
func isFlagPassed(f *flag.FlagSet, name string) bool {
	found := false
	f.Visit(func(f *flag.Flag) {
		//fmt.Println(">>>FLAG: ", f.Name)
		if f.Name == name {
			//fmt.Println("Found ", name, " in ", f.Name)
			found = true
		}
	})
	//fmt.Println("END ISFLAGPASSED")
	return found
}

// https://stackoverflow.com/questions/52914127/is-there-a-way-to-determine-whether-a-flag-was-set-when-using-flag-visitall
func unsetFlags(fs *flag.FlagSet) []*flag.Flag {
	var unset []*flag.Flag
	fs.VisitAll(func(f *flag.Flag) {
		unset = append(unset, f)
	})
	fs.Visit(func(f *flag.Flag) {
		fmt.Println(">>>FLAG: ", f.Name)
		for i, h := range unset {
			if f == h {
				unset = append(unset[:i], unset[i+1:]...)
			}
		}
	})
	return unset
}
