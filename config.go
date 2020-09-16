package commons

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/wlbr/common/log"
)

type CommonConfig struct {
	BuildTimeStamp   time.Time
	GitVersion       string
	ShowVersion      bool
	ActiveLogLevel   log.LogLevel
	logLevel         string
	LogFileName      string
	Logger           *log.Logger
	WorkingDirectory string
	colouredLogging  bool
	cleanup          []func() error
}

func (cfg CommonConfig) String() string {
	var logfname string
	if cfg.LogFileName == "" {
		logfname = "STDOUT"
	} else {
		logfname = cfg.LogFileName
	}
	return fmt.Sprintf("\tBuildTimeStamp: %s\n"+
		"\tGitVersion: %s\n"+
		"\tActiveLogLevel: %+v\n"+
		"\tLogFileName: %s\n"+
		"\tLogger: %v\n"+
		"\tWorking Directory: %s\n",
		cfg.BuildTimeStamp, cfg.GitVersion, cfg.ActiveLogLevel.String(),
		logfname, cfg.Logger, cfg.WorkingDirectory)
}

// GetInspectData offers some additional debugging information
func (cfg CommonConfig) GetInspectData() string {
	return fmt.Sprintf("\tGolang compile version: %s \n", runtime.Version()) +
		fmt.Sprintf("\tCompile GOROOT: %s \n", runtime.GOROOT()) +
		fmt.Sprintf("\tCompile GOOS: %s \n", runtime.GOOS) +
		fmt.Sprintf("\tCompile GOARCH: %s \n", runtime.GOARCH) +
		fmt.Sprintf("\tRuntime NumCPU: %d \n", runtime.NumCPU()) +
		fmt.Sprintf("\tRuntime NumGoroutine: %d\n", runtime.NumGoroutine())
}

func (cfg *CommonConfig) FlagDefinition() {
	flag.StringVar(&cfg.logLevel, "loglevel", "Warn", "Determines logging verbosity. [All|Info|Debug|Warn|Error|Fatal|Off].")
	flag.StringVar(&cfg.LogFileName, "logfile", "", "Sets the name of the logfile. Uses STDOUT if empty.")
	flag.BoolVar(&cfg.ShowVersion, "version", false, "Show version info.")
	flag.BoolVar(&cfg.colouredLogging, "logcolour", true, "Use coloured logging (switch of when redirecting log output).")
}

func (cfg *CommonConfig) Initialize(version string, buildtimestamp string) *CommonConfig {
	btime, err := time.Parse("2006-01-02_15:04:05_MST", buildtimestamp)
	if err != nil {
		btime = time.Now()
	}
	cfg.BuildTimeStamp = btime
	cfg.GitVersion = version
	cfg.WorkingDirectory, _ = os.Getwd()

	if !flag.Parsed() {
		flag.Parse()
	}
	// Settig up the logger
	cfg.ActiveLogLevel, err = log.LogLevelString(cfg.logLevel)
	if err != nil {
		cfg.ActiveLogLevel = log.All
	}
	cfg.Logger = log.NewLogger(cfg.LogFileName, cfg.ActiveLogLevel, cfg.colouredLogging)
	cfg.Logger.SetConvenienceLogger()
	log.Debug("Current working directory is '%s'.", cfg.WorkingDirectory)
	if err != nil {
		log.Warn("Error in config, Loglevel '%s' not existing in tools/loglevel.go. Setting LogLevel to 'All'", cfg.logLevel)
	}

	if cfg.ShowVersion {
		v := cfg.GitVersion
		if strings.ToLower(v) == "unknown build" {
			v = "'Unknown build'"
		}

		fmt.Printf("Version %s built on %s using %s.\n", v, cfg.BuildTimeStamp.Format("02.01.2006"), runtime.Version())
		os.Exit(0)
	}
	return cfg
}

func (cfg *CommonConfig) CleanUp() {
	log.Debug("Cleaning up.")
	for _, fun := range cfg.cleanup {
		fun()
	}
}

func (cfg *CommonConfig) AddCleanUpFn(f func() error) {
	cfg.cleanup = append(cfg.cleanup, f)
}

func (cfg *CommonConfig) FatalExit() {
	cfg.CleanUp()
	os.Exit(1)
}
