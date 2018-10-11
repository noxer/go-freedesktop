package basedir

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func homeDir() (string, error) {

	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	if usr.HomeDir == "" {
		return "", errors.New("unable to find home directory")
	}

	return usr.HomeDir, nil

}

func lookupOrUser(env, dir string) (string, error) {

	e, _ := os.LookupEnv(env)
	if e != "" {
		return e, nil
	}

	d, err := homeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(d, dir), nil

}

// DataHome retrieves the XDG_DATA_HOME variable as defined in https://specifications.freedesktop.org/basedir-spec/latest/ar01s03.html
func DataHome() (string, error) {
	return lookupOrUser("XDG_DATA_HOME", ".local/share")
}

// ConfigHome retrieves the XDG_CONFIG_HOME variable as defined in https://specifications.freedesktop.org/basedir-spec/latest/ar01s03.html
func ConfigHome() (string, error) {
	return lookupOrUser("XDG_CONFIG_HOME", ".config")
}

// DataDirs retrieves the XDG_DATA_DIRS variable as defined in https://specifications.freedesktop.org/basedir-spec/latest/ar01s03.html
func DataDirs() []string {

	env, _ := os.LookupEnv("XDG_DATA_DIRS")
	if env != "" {
		return strings.Split(env, ":")
	}

	return []string{"/usr/local/share", "/usr/share"}

}

// ConfigDirs retrieves the XDG_CONFIG_DIRS variable as defined in https://specifications.freedesktop.org/basedir-spec/latest/ar01s03.html
func ConfigDirs() []string {

	env, _ := os.LookupEnv("XDG_CONFIG_DIRS")
	if env != "" {
		return strings.Split(env, ":")
	}

	return []string{"/etc/xdg"}

}

// CacheHome retrieves the XDG_CACHE_HOME variable as defined in https://specifications.freedesktop.org/basedir-spec/latest/ar01s03.html
func CacheHome() (string, error) {
	return lookupOrUser("XDG_CACHE_HOME", ".cache")
}

// RuntimeDir retrieves the XDG_RUNTIME_DIR variable as defined in https://specifications.freedesktop.org/basedir-spec/latest/ar01s03.html
func RuntimeDir() (string, error) {

	env, _ := os.LookupEnv("XDG_RUNTIME_DIR")
	if env != "" {
		return env, nil
	}

	return "", errors.New("unable to find XDG_RUNTIME_DIR")

}
