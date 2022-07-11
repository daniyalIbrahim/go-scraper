package helpers

import (
	"bytes"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Process struct {
	pid int
	cpu float64
}

func GetImageFromPath(path string) image.Image {
	log.Println("GetImageFromPath")
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func ImageToBytesBuffer(imagePath string) []byte {
	log.Println("ImageToBytesBuffer")
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read the file as a slice of bytes
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	return buffer
}
func getCPUSample() (idle, total uint64) {
	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}
				total += val // tally up all the numbers to get total ticks
				if i == 4 {  // idle is the 5th field in the cpu line
					idle = val
				}
			}
			return
		}
	}
	return
}
func LogCPUInfo() {
	idle0, total0 := getCPUSample()
	time.Sleep(3 * time.Second)
	idle1, total1 := getCPUSample()

	idleTicks := float64(idle1 - idle0)
	totalTicks := float64(total1 - total0)
	cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks

	GetProcessInfo()
	log.Printf("CPU usage is %f%% [busy: %f, total: %f]\n", cpuUsage, totalTicks-idleTicks, totalTicks)

}
func GetProcessInfo() {
	cmd := exec.Command("ps", "aux")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	processes := make([]*Process, 0)
	for {
		line, err := out.ReadString('\n')
		if err != nil {
			break
		}
		tokens := strings.Split(line, " ")
		ft := make([]string, 0)
		for _, t := range tokens {
			if t != "" && t != "\t" {
				ft = append(ft, t)
			}
		}
		log.Println(len(ft), ft)
		pid, err := strconv.Atoi(ft[1])
		if err != nil {
			continue
		}
		cpu, err := strconv.ParseFloat(ft[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		processes = append(processes, &Process{pid, cpu})
	}
	for _, p := range processes {
		log.Println("Process ", p.pid, " takes ", p.cpu, " % of the CPU")
	}
}
func GetCPUInformation() {
	log.Printf("Getting CPU Information")
	log.Printf("OS: %s", runtime.GOOS)
	//Get MAX CPU CORES
	log.Printf("Max CPU Cores: %v", runtime.NumCPU())
	//Get MAX CPU FREQUENCY
	log.Printf("Runtime GOARCH: %v", runtime.GOARCH)
	//Get MAX CPU THREADS
	log.Printf("Max CPU Threads: %v", runtime.GOMAXPROCS(0))
	GetRamUsage(runtime.GOOS)
}
func GetRamUsage(osName string) {

	if osName == "linux" {
		cmd := exec.Command("free", "-m")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		for _, line := range strings.Split(out.String(), "\n") {
			log.Printf("%s", line)
		}
		LogCPUInfo()

	} else if osName == "windows" {
		cmd := exec.Command("wmic", "memorychip", "get")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		for _, line := range strings.Split(out.String(), "\t") {
			if line != "" {
				log.Printf("%s", line)
			}
		}

	}

}
