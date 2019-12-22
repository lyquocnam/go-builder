package module

import (
	"fmt"
	"os/exec"
)

type Docker interface {
	Build(tagName string)
	RunDev(serviceName string)
	RunProd(tagName string)
	Execute(cm string)
}

type docker struct {
	logger Logger
}

func NewDocker(logger Logger) *docker {
	return &docker{logger: logger}
}

func (s *docker) Build(tagName string) {
	cm := fmt.Sprintf(`docker build \
-f deploy/Dockerfile \
-t %s \
--build-arg GITHUB_TOKEN=d7519779515245e32e346444a458377f501b70b0 \
.`, tagName)
	s.Execute(cm)
}


func (s *docker) RunDev(serviceName string) {
	cm := fmt.Sprintf(`docker-compose -f deploy/docker-compose.yml \
-f deploy/docker-compose.override.yml \
run --name %s \
-d go-oauth`, serviceName)

	s.Execute(cm)
}


func (s *docker) RunProd(serviceName string) {
	cm := fmt.Sprintf(`docker-compose -f deploy/docker-compose.yml \
-f deploy/docker-compose.prod.yml \
run --name %s \
-d go-oauth`, serviceName)

	s.Execute(cm)
}

func (s *docker) Execute(cmdStr string) {
	fmt.Println(cmdStr)

	out, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			s.logger.Warnf(string(exitErr.Stderr))
			return
		}
		s.logger.Warnf(err.Error())
	} else {
		s.logger.Infof(string(out))
	}
}