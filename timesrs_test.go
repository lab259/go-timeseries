package timesrs_test

import (
	"github.com/jamillosantos/macchiato"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"testing"
)

func TestTimesrs(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	macchiato.RunSpecs(t, "go-timeseries")
}
