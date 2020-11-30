package chapter8

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// the-gpl du -dir=$HOME/gocode -v=true
// Prints, size of files in B, KB, MB and GB:
// 723464 files: 17730797665.0 B, 17730797.7 KB, 17730.8 MB, 17.7 GB
// It uses, three types of go-routines in a cooperatively to calculate  this:
//   walkDir go-routine: spawns MaxGoRoutines parallel, and each spawn pipe file-sizes to 'sizes' channel.
//   anon go-routine-1: closes the 'sizes' channel when walkDir routines are done
//	 anon go-routine-2:
//				adds all sizes by reading each filesize from channel 'sizes'
//  			prints output every second if verbose flag -v is set
//	WARNING: When should a channel be closed when there are multiple writers.
//		Solution 1: Use Wait group, i.e. close a channel when all writers are done.
//    Solution 2: Implement cancel mechanism using separate  cancel or done channel.
//			The sender go-routines monitor it and stop further processing if 'done' is closed.

// MaxGoRoutines number of parallel go-routines
const MaxOpenFiles = 1024 // Reduce number of open system files
const MaxGoRoutines = MaxOpenFiles / 2

var sema = make(chan struct{}, MaxGoRoutines) // sema counting semaphore to  run
var wg sync.WaitGroup                         // Wait for all walkDir go calls to  finish

// walkDir traverses dir tree and  puts size of each file in channel 'sizes'
func walkDir(dir string, sizes chan<- int64) {
	defer wg.Done()
	for _, entry := range dirEntry(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			go walkDir(subDir, sizes)
		}
		sizes <- entry.Size()
	}
}

// dirEntry reads files in a directory and returns entries for  all.
//	Runs a max  of MaxGoRoutines of these goroutines  in parallel
func dirEntry(dir string) []os.FileInfo {
	defer func() { <-sema }() // Release  semaphore,  when done
	sema <- struct{}{}        // Acquire semaphore
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("du: %v", err)
		return nil
	}
	return entries
}

// DU Disk Usage returns size of a directory
func DU(dir string, verbose bool) int64 {
	sizes := make(chan int64, MaxOpenFiles) // Create c channel that holds sizes for MaxOpenFiles files max
	wg.Add(1)
	// go-routine 1: Walks the directory structure recursively
	go walkDir(dir, sizes)

	// go-routine 2: Closes the sizes channel when done
	go func() {
		wg.Wait()
		close(sizes)
	}()
	// go-routine 3 (this): wait and read sizes from channel
	ticker := time.NewTicker(1 * time.Second)
	if !verbose {
		ticker.Stop()
	}
	var du, nFiles int64
waitLoop:
	for {
		select { // No  mutex needed, select does only one of the channel operations
		case sz, open := <-sizes:
			if !open { //  channel is closed
				break waitLoop
			}
			du += sz
			nFiles++
		case <-ticker.C:
			printDU(nFiles, du)
		}
	}
	// go-routine 2:
	printDU(nFiles, du)
	if verbose {
		ticker.Stop()
	}
	return du
}

// printDU prints disc usage
func printDU(nFiles, nBytes int64) {
	fmt.Printf("%d files: %.1f B, %.1f KB, %.1f MB, %.1f GB\n", nFiles,
		float64(nBytes), float64(nBytes)/1e3,
		float64(nBytes)/1e6, float64(nBytes)/1e9)
}
