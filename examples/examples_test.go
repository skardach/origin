package examples

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"testing"

	kubeapi "github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/runtime"

	"github.com/openshift/origin/pkg/api/latest"
	buildapi "github.com/openshift/origin/pkg/build/api"
	configapi "github.com/openshift/origin/pkg/config/api"
	templateapi "github.com/openshift/origin/pkg/template/api"
)

func TestExamples(t *testing.T) {
	expected := map[string]runtime.Object{
		"guestbook/template.json":                              &templateapi.Template{},
		"hello-openshift/hello-pod.json":                       &kubeapi.Pod{},
		"simple-ruby-app/buildcfg/buildcfg.json":               &buildapi.BuildConfig{},
		"simple-ruby-app/buildinvoke/pushevent.json":           nil, // Skip.
		"simple-ruby-app/registry_config/registry_config.json": &configapi.Config{},
		"simple-ruby-app/template/template.json":               &templateapi.Template{},
	}

	files := []string{}
	err := filepath.Walk(".", func(path string, f os.FileInfo, err error) error {
		if filepath.Ext(path) == ".json" {
			files = append(files, path)
		}
		return err
	})
	if err != nil {
		t.Errorf("%v", err)
	}

	for _, file := range files {
		expectedObject, ok := expected[file]
		if !ok {
			t.Errorf("No test case defined for example JSON file '%v'", file)
			continue
		}
		if expectedObject == nil {
			continue
		}

		jsonData, _ := ioutil.ReadFile(file)
		if err := latest.Codec.DecodeInto(jsonData, expectedObject); err != nil {
			t.Errorf("Unexpected error while decoding example JSON file '%v': %v", file, err)
		}
	}
}

func TestReadme(t *testing.T) {
	path := "../README.md"
	_, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("Unable to read file: %v", err)
	}
}
