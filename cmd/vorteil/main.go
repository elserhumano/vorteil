package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/sisatech/tablewriter"
	"github.com/spf13/pflag"
	"github.com/vorteil/vorteil/pkg/ext"
	"github.com/vorteil/vorteil/pkg/vcfg"
	"github.com/vorteil/vorteil/pkg/vdecompiler"
	"github.com/vorteil/vorteil/pkg/vdisk"
	"github.com/vorteil/vorteil/pkg/vio"
	"github.com/vorteil/vorteil/pkg/vpkg"
	"github.com/vorteil/vorteil/pkg/vproj"
)

var (
	release = "0.0.0"
	commit  = ""
	date    = "Thu, 01 Jan 1970 00:00:00 +0000"
)

// Each command sets a code so we can capture the exit if an error has been encountered.
var errorStatusCode int
var errorStatusMessage error

// setError for the command to exit with a certain status and message
func setError(err error, code int) {
	errorStatusCode = code
	errorStatusMessage = err
}

func main() {

	commandInit()

	err := rootCmd.Execute()
	if err != nil {
		setError(err, 1)
	}

	// If the global status code for errors hasn't been set os.exit with non-zero number
	if errorStatusCode != 0 {
		handleCommandError()
	}
}

func isEmptyDir(path string) bool {

	fis, err := ioutil.ReadDir(path)
	if err != nil {
		return false
	}

	if len(fis) > 0 {
		return false
	}

	return true
}

func isNotExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

func checkValidNewDirOutput(path string, force bool, dest, flag string) error {
	if !isEmptyDir(path) {
		return checkValidNewFileOutput(path, force, dest, flag)
	}

	return nil
}

func checkValidNewFileOutput(path string, force bool, dest, flag string) error {
	if !isNotExist(path) {
		if force {
			err := os.RemoveAll(path)
			if err != nil {
				return fmt.Errorf("failed to delete existing %s '%s': %w", dest, path, err)
			}

			dir := filepath.Dir(path)
			err = os.MkdirAll(dir, 0777)
			if err != nil {
				return fmt.Errorf("failed to create parent directory for %s '%s': %w", dest, path, err)
			}
		} else {
			return fmt.Errorf("%s '%s' already exists (you can use '%s' to force an overwrite)", dest, path, flag)
		}
	}

	return nil
}

func parseImageFormat(s string) (vdisk.Format, error) {
	format, err := vdisk.ParseFormat(s)
	if err != nil {
		return format, fmt.Errorf("%w -- try one of these: %s", err, strings.Join(vdisk.AllFormatStrings(), ", "))
	}
	return format, nil
}

type sourceType string

const (
	sourceURL     sourceType = "URL"
	sourceFile               = "File"
	sourceDir                = "Dir"
	sourceINVALID            = "INVALID"
)

func getSourceType(src string) (sourceType, error) {
	var err error
	var fi os.FileInfo

	// Check if Source is a URL
	if _, err := url.ParseRequestURI(src); err == nil {
		if u, uErr := url.Parse(src); uErr == nil && u.Scheme != "" && u.Host != "" && u.Path != "" {
			return sourceURL, nil
		}
	}

	// Check if Source is a file or dir
	fi, err = os.Stat(src)
	if !os.IsNotExist(err) && (fi != nil && !fi.IsDir()) {
		return sourceFile, nil
	} else if !os.IsNotExist(err) && (fi != nil && fi.IsDir()) {
		return sourceDir, nil
	}

	// Source is unknown and thus is invalid
	return sourceINVALID, err
}

func getReaderURL(src string) (vpkg.Reader, error) {

	resp, err := http.Get(src)
	if err != nil {
		resp.Body.Close()
		return nil, err
	}
	p := log.NewProgress("Downloading package", "KiB", resp.ContentLength)

	pkgr, err := vpkg.Load(p.ProxyReader(resp.Body))
	if err != nil {
		return nil, err
	}
	return pkgr, nil
}

func getBuilderURL(argName, src string) (vpkg.Builder, error) {
	pkgr, err := getReaderURL(src)
	if err != nil {
		return nil, err
	}
	pkgb, err := vpkg.NewBuilderFromReader(pkgr)
	if err != nil {
		pkgr.Close()
		return nil, err
	}
	return pkgb, nil
}

func getReaderFile(src string) (vpkg.Reader, error) {
	f, err := os.Open(src)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
		return nil, fmt.Errorf("failed to resolve %s ", src)
	}

	pkgr, err := vpkg.Load(f)
	if err != nil {
		f.Close()
		return nil, err
	}
	return pkgr, nil
}
func getBuilderFile(argName, src string) (vpkg.Builder, error) {

	pkgr, err := getReaderFile(src)
	if err != nil {
		pkgr.Close()
		return nil, err
	}
	pkgb, err := vpkg.NewBuilderFromReader(pkgr)
	if err != nil {
		pkgr.Close()
		return nil, err
	}
	return pkgb, err
}

func getBuilderDir(argName, src string) (vpkg.Builder, error) {
	var ptgt *vproj.Target
	path, target := vproj.Split(src)
	proj, err := vproj.LoadProject(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
		return nil, fmt.Errorf("failed to resolve %s '%s'", argName, src)
	}

	ptgt, err = proj.Target(target)
	if err != nil {
		return nil, err
	}

	pkgb, err := ptgt.NewBuilder()
	return pkgb, err
}

func getPackagerReader(argName, src string) (vpkg.Reader, error) {
	var err error
	var pkgR vpkg.Reader
	sType, err := getSourceType(src)
	if err != nil {
		return nil, err
	}

	switch sType {
	case sourceURL:
		pkgR, err = getReaderURL(src)
	case sourceFile:
		pkgR, err = getReaderFile(src)
	case sourceINVALID:
		fallthrough
	default:
		err = fmt.Errorf("failed to resolve %s '%s'", argName, src)
	}

	return pkgR, err
}
func getPackageBuilder(argName, src string) (vpkg.Builder, error) {
	var err error
	var pkgB vpkg.Builder
	sType, err := getSourceType(src)
	if err != nil {
		return nil, err
	}

	switch sType {
	case sourceURL:
		pkgB, err = getBuilderURL(argName, src)
	case sourceFile:
		pkgB, err = getBuilderFile(argName, src)
	case sourceDir:
		pkgB, err = getBuilderDir(argName, src)
	case sourceINVALID:
		fallthrough
	default:
		err = fmt.Errorf("failed to resolve %s '%s'", argName, src)
	}

	return pkgB, err

	// TODO: check for vrepo strings

}

var (
	flagIcon             string
	flagVCFG             []string
	flagInfoDate         string
	flagInfoURL          string
	flagSystemFilesystem string
	flagSystemOutputMode string
	flagSysctl           []string
	flagVMDiskSize       string
	flagVMInodes         string
	flagVMRAM            string
	overrideVCFG         vcfg.VCFG
)

func addModifyFlags(f *pflag.FlagSet) {
	vcfgFlags.AddTo(f)
}

// mergeFlagVCFGFiles : Merge values from from VCFG files stored in 'flagVCFG', and then merge vcfg flag values with overrideVCFG
func mergeVCFGFlagValues(b *vpkg.Builder) error {
	var err error
	var f vio.File
	var cfg *vcfg.VCFG

	// Iterate over vcfg paths stored in flagVCFG, read vcfg files and merge into b
	for _, path := range flagVCFG {
		f, err = vio.Open(path)
		if err != nil {
			return err
		}

		cfg, err = vcfg.LoadFile(f)
		if err != nil {
			return err
		}

		err = (*b).MergeVCFG(cfg)
		if err != nil {
			return err
		}
	}

	// Merge overrideVCFG object containing flag values into b
	err = (*b).MergeVCFG(&overrideVCFG)
	return err
}

func modifyPackageBuilder(b vpkg.Builder) error {
	var err error
	var f vio.File

	err = vcfgFlags.Validate()
	if err != nil {
		return err
	}

	if flagIcon != "" {
		f, err = vio.LazyOpen(flagIcon)
		if err != nil {
			return err
		}
		b.SetIcon(f)
	}

	err = mergeVCFGFlagValues(&b)
	if err != nil {
		return err
	}

	err = handleFileInjections(b)
	return err
}

// NumbersMode determines which numbers format a PrintableSize should render to.
var NumbersMode int

// SetNumbersMode parses s and sets NumbersMode accordingly.
func SetNumbersMode(s string) error {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	switch s {
	case "", "short":
		NumbersMode = 0
	case "dec", "decimal":
		NumbersMode = 1
	case "hex", "hexadecimal":
		NumbersMode = 2
	default:
		return fmt.Errorf("numbers mode must be one of 'dec', 'hex', or 'short'")
	}
	return nil
}

// PrintableSize is a wrapper around int to alter its string formatting behaviour.
type PrintableSize int

// String returns a string representation of the PrintableSize, formatted according to the global NumbersMode.
func (c PrintableSize) String() string {
	switch NumbersMode {
	case 0:
		x := int(c)
		if x == 0 {
			return "0"
		}
		var units int
		var suffixes = []string{"", "K", "M", "G"}
		for {
			if x%1024 != 0 {
				break
			}
			x /= 1024
			units++
			if units == len(suffixes)-1 {
				break
			}
		}
		return fmt.Sprintf("%d%s", x, suffixes[units])
	case 1:
		return fmt.Sprintf("%d", int(c))
	case 2:
		return fmt.Sprintf("%#x", int(c))
	default:
		panic("invalid NumbersMode")
	}
}

// PlainTable prints data in a grid, handling alignment automatically.
func PlainTable(vals [][]string) {
	if len(vals) == 0 {
		panic(errors.New("no rows provided"))
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorder(false)
	table.SetColumnSeparator("")
	for i := 1; i < len(vals); i++ {
		table.Append(vals[i])
	}

	table.Render()
}

func createSymlinkCallback(iio *vdecompiler.IO, inode *ext.Inode, dpath string) func() error {
	return func() error {
		rdr, err := iio.InodeReader(inode)
		if err != nil {
			return err
		}
		data, err := ioutil.ReadAll(rdr)
		if err != nil {
			return err
		}

		err = os.Symlink(string(string(data)), dpath)
		if err != nil {
			return err
		}
		return nil
	}
}

func copyInodeToRegularFile(iio *vdecompiler.IO, inode *ext.Inode, dpath string) error {
	var err error
	var f *os.File
	var rdr io.Reader

	err = utilFileNotExists(dpath)
	if err != nil {
		return err
	}

	f, err = os.Create(dpath)
	if err != nil {
		return err
	}
	defer f.Close()

	rdr, err = iio.InodeReader(inode)
	if err != nil {
		return err
	}

	_, err = io.CopyN(f, rdr, int64(vdecompiler.InodeSize(inode)))
	return err
}

func utilFileNotExists(fpath string) error {
	_, err := os.Stat(fpath)
	if !os.IsNotExist(err) {
		if err == nil {
			err = fmt.Errorf("file already exists: %s", fpath)
		}
		return err
	}
	return nil
}

func decompile(srcPath, outPath string) {

	log.Printf("decompiling disk")

	iio, err := vdecompiler.Open(srcPath)
	if err != nil {
		setError(err, 1)
		return
	}
	defer iio.Close()

	fi, err := os.Stat(outPath)
	if err != nil && !os.IsNotExist(err) {
		setError(err, 2)
		return
	}
	var into bool
	if !os.IsNotExist(err) && fi.IsDir() {
		into = true
	}

	fpath := "/"
	dpath := outPath
	if into {
		dpath = filepath.ToSlash(filepath.Join(outPath, filepath.Base(fpath)))
	}

	var counter int

	symlinkCallbacks := make([]func() error, 0)

	var recurse func(int, string, string) error
	recurse = func(ino int, rpath string, dpath string) error {
		var entries []*vdecompiler.DirectoryEntry

		inode, err := iio.ResolveInode(ino)
		if err != nil {
			return err
		}

		if flagTouched && inode.LastAccessTime == 0 && !vdecompiler.InodeIsDirectory(inode) && rpath != "/" {
			log.Printf("skipping untouched object: %s", rpath)
			goto DONE
		}

		counter++

		log.Debugf("copying %s", rpath)

		if vdecompiler.InodeIsRegularFile(inode) {
			err = copyInodeToRegularFile(iio, inode, dpath)
			goto DONE
		}

		if vdecompiler.InodeIsSymlink(inode) {
			symlinkCallbacks = append(symlinkCallbacks, createSymlinkCallback(iio, inode, dpath))
			goto DONE
		}

		if !vdecompiler.InodeIsDirectory(inode) {
			log.Warnf("skipping abnormal file: %s", rpath)
			goto DONE
		}

		// INODE IS DIR
		err = utilFileNotExists(dpath)
		if err == nil {
			err = os.MkdirAll(dpath, 0777)
			if err == nil {
				entries, err = iio.Readdir(inode)
			}
		}

		if err != nil {
			return err
		}

		for _, entry := range entries {
			if entry.Name == "." || entry.Name == ".." {
				continue
			}
			err = recurse(entry.Inode, filepath.ToSlash(filepath.Join(rpath, entry.Name)), filepath.Join(dpath, entry.Name))
			if err != nil {
				return err
			}
		}

	DONE:
		return err
	}

	ino, err := iio.ResolvePathToInodeNo(fpath)
	if err != nil {
		setError(err, 3)
		return
	}
	err = recurse(ino, filepath.ToSlash(filepath.Base(fpath)), dpath)
	if err != nil {
		setError(err, 4)
		return
	}

	for _, fn := range symlinkCallbacks {
		err = fn()
		if err != nil {
			setError(err, 5)
			return
		}
	}

	if flagTouched && counter <= 1 {
		log.Warnf("No touched files detected. Are you sure this disk has been run?")
	}
}
