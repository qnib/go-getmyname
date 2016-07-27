package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"

  "github.com/docker/engine-api/types/filters"
  "github.com/docker/engine-api/types"
  "github.com/docker/engine-api/client"
  "golang.org/x/net/context"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func getContainerID() (string) {
    file, err := os.Open("/proc/self/cgroup")
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line,":")
        if (len(parts) == 3) && (parts[1] == "cpuset") {
          return strings.Split(parts[2],"/")[2]
        }
    }
    return ""
}

func getContainerName(cId string) (string) {
  defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
  var err error
  var cli *client.Client
  if _, err := os.Stat("/var/run/docker.sock"); err == nil {
    cli, err = client.NewClient("unix:///var/run/docker.sock", "v1.22", nil, defaultHeaders)
  } else {
    dHost := os.Getenv("DOCKER_HOST")
    if ! strings.HasPrefix(dHost, "tcp://") {
      dHost = fmt.Sprintf("tcp://%s", dHost)
    }
    cli, err = client.NewClient(dHost, "v1.22", nil, defaultHeaders)
  }
  if err != nil {
    panic(err)
  }
  cFilter := filters.NewArgs()
  cFilter.Add("id", cId)
  options := types.ContainerListOptions{Filter: cFilter}
  containers, err := cli.ContainerList(context.Background(), options)

  for _,c := range containers {
    return strings.TrimPrefix(c.Names[0], "/")
  }
  return ""
}


func main() {
  cId := getContainerID()
  //fmt.Printf("cId:\t%s\n",cId)
  cName := getContainerName(cId)
  fmt.Println(cName)

}
