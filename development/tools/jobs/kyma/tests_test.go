package kyma

import (
	"path"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/jobsuite"
)

var tests = []struct {
	path              string
	image             string
	suite             func(config *jobsuite.Config) jobsuite.Suite
	additionalOptions []jobsuite.Option
}{
	{path: "service-catalog", image: tester.ImageGolangBuildpack1_16, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("tests-generic"),
		},
	},
	{path: "application-connector-tests", image: tester.ImageGolangBuildpack1_16, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
		},
	},
	{path: "application-gateway-tests", image: tester.ImageGolangBuildpack1_16, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
		},
	},
	{path: "application-gateway-legacy-tests", image: tester.ImageGolangBuildpack1_16, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
		},
	},
	{path: "application-operator-tests", image: tester.ImageGolangBuildpack1_16, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
		},
	},
	{path: "application-registry-tests", image: tester.ImageGolangBuildpack1_16, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
		},
	},
	{path: "compass-runtime-agent", image: tester.ImageGolangBuildpack1_16, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("tests-generic"),
		},
	},
	{path: "connection-token-handler-tests", image: tester.ImageGolangBuildpack1_16, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
		},
	},
	{path: "connector-service-tests", image: tester.ImageGolangBuildpack1_16, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
		},
	},
	{path: "console-backend-service", image: tester.ImageGolangBuildpack1_16, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("tests-generic"),
			jobsuite.RunIfChanged("components/console-backend-service/main.go", "scripts/go-dep.mk"),
		},
	},
	{path: "end-to-end/upgrade", image: tester.ImageGolangBuildpack1_16, suite: tester.NewGenericComponentSuite, additionalOptions: []jobsuite.Option{
		jobsuite.RunIfChanged("^tests/end-to-end/upgrade/[^chart]", "tests/end-to-end/upgrade/fix"),
		jobsuite.JobFileSuffix("tests-generic"),
	}},
	{path: "integration/apiserver-proxy", image: tester.ImageGolangBuildpack1_16, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("tests-generic"),
		},
	},
	{path: "rafter", image: tester.ImageGolangBuildpack1_16, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("tests-generic"),
		},
	},
	{path: "integration/api-gateway", image: tester.ImageGolangBuildpack1_16, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("tests-generic"),
		},
	},
	{path: "integration/dex", image: tester.ImageGolangBuildpack1_16, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("tests-generic"),
		},
	},
	{path: "integration/logging", image: tester.ImageGolangBuildpack1_16, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
		},
	},
	{path: "contract/knative-channel-kafka", image: tester.ImageGolangBuildpack1_16, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("tests-generic"),
		},
	},
	{path: "function-controller", image: tester.ImageGolangBuildpack1_16, suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("tests-generic"),
		},
	},
}

func TestTestJobs(t *testing.T) {
	testedConfigurations := make(map[string]struct{})
	repos := map[string]struct{}{}
	for _, test := range tests {
		t.Run(test.path, func(t *testing.T) {
			opts := []jobsuite.Option{
				jobsuite.Test(test.path, test.image),
				jobsuite.KymaRepo(),
				jobsuite.AllReleases(),
			}
			opts = append(opts, test.additionalOptions...)
			cfg := jobsuite.NewConfig(opts...)
			suite := test.suite
			if suite == nil {
				suite = tester.NewComponentSuite
			}
			ts := suite(cfg)
			if pathProvider, ok := ts.(jobsuite.JobConfigPathProvider); ok {
				testedConfigurations[path.Clean(pathProvider.JobConfigPath())] = struct{}{}
			}
			repos[cfg.Repository] = struct{}{}
			ts.Run(t)
		})
	}
	t.Run("All Files covered by test", jobsuite.CheckFilesAreTested(repos, testedConfigurations, jobBasePath, []string{"tests"}))
}
