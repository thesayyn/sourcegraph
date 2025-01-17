package config

import (
	"fmt"

	"github.com/sourcegraph/sourcegraph/lib/errors"
	"github.com/sourcegraph/sourcegraph/lib/pointers"
)

// NewDefaultConfig
//
// Warning: never extract `ptr.To(thing)` into a package-level variable! If you
// do this, reconciling a config that overrides a default value for that
// pointer, will affect the subsequent _default_ for all future resources
// reconciled.
func NewDefaultConfig() Sourcegraph {
	return Sourcegraph{
		Spec: SourcegraphSpec{
			// Global config
			ImageRepository: "index.docker.io/sourcegraph",

			// Service-specific config
			Blobstore: BlobstoreSpec{
				StandardConfig: StandardConfig{
					PersistentVolumeConfig: PersistentVolumeConfig{
						StorageSize: "100Gi",
					},
				},
			},
			RepoUpdater: RepoUpdaterSpec{
				StandardConfig: StandardConfig{
					PrometheusPort: pointers.Ptr(6060),
				},
			},
			Symbols: SymbolsSpec{
				StandardConfig: StandardConfig{
					PrometheusPort: pointers.Ptr(6060),
					PersistentVolumeConfig: PersistentVolumeConfig{
						StorageSize: "12Gi",
					},
				},
				Replicas: 1,
			},
			GitServer: GitServerSpec{
				StandardConfig: StandardConfig{
					PrometheusPort: pointers.Ptr(6060),
					PersistentVolumeConfig: PersistentVolumeConfig{
						StorageSize: "200Gi",
					},
				},
				Replicas: 1,
			},
			PGSQL: PGSQLSpec{
				StandardConfig: StandardConfig{
					PrometheusPort: pointers.Ptr(9187),
					PersistentVolumeConfig: PersistentVolumeConfig{
						StorageSize: "200Gi",
					},
				},
				DatabaseConnection: &DatabaseConnectionSpec{
					Host:     "pgsql",
					Port:     "5432",
					User:     "sg",
					Password: "password",
					Database: "sg",
				},
			},
			RedisCache: RedisSpec{
				StandardConfig: StandardConfig{
					PrometheusPort: pointers.Ptr(9121),
					PersistentVolumeConfig: PersistentVolumeConfig{
						StorageSize: "100Gi",
					},
				},
			},
			RedisStore: RedisSpec{
				StandardConfig: StandardConfig{
					PrometheusPort: pointers.Ptr(9121),
					PersistentVolumeConfig: PersistentVolumeConfig{
						StorageSize: "100Gi",
					},
				},
			},
			SyntectServer: SyntectServerSpec{
				StandardConfig: StandardConfig{
					PrometheusPort: pointers.Ptr(6060),
				},
				Replicas: 1,
			},
			PreciseCodeIntel: PreciseCodeIntelSpec{
				StandardConfig: StandardConfig{
					PrometheusPort: pointers.Ptr(6060),
				},
				NumWorkers: 4,
				Replicas:   2,
			},
			CodeInsights: CodeDBSpec{
				StandardConfig: StandardConfig{
					PrometheusPort: pointers.Ptr(9187),
					PersistentVolumeConfig: PersistentVolumeConfig{
						StorageSize: "200Gi",
					},
				},
				DatabaseConnection: &DatabaseConnectionSpec{
					Host:     "codeinsights-db",
					Port:     "5432",
					User:     "postgres",
					Password: "password",
					Database: "postgres",
				},
			},
			CodeIntel: CodeDBSpec{
				StandardConfig: StandardConfig{
					PrometheusPort: pointers.Ptr(9187),
					PersistentVolumeConfig: PersistentVolumeConfig{
						StorageSize: "200Gi",
					},
				},
				DatabaseConnection: &DatabaseConnectionSpec{
					Host:     "codeintel-db",
					Port:     "5432",
					User:     "sg",
					Password: "password",
					Database: "sg",
				},
			},
			Prometheus: PrometheusSpec{
				StandardConfig: StandardConfig{
					PersistentVolumeConfig: PersistentVolumeConfig{
						StorageSize: "200Gi",
					},
				},
			},
			Cadvisor: CadvisorSpec{
				StandardConfig: StandardConfig{
					// cadvisor is opt-in due to the privilege requirements
					Disabled:       true,
					PrometheusPort: pointers.Ptr(48080),
				},
			},
			Worker: WorkerSpec{
				StandardConfig: StandardConfig{
					PrometheusPort: pointers.Ptr(6060),
				},
				Replicas: 1,
			},
		},
	}
}

// Images

// Map of version to map of service to image tag
var defaultImages = map[string]map[string]string{
	"5.3.9104": defaultImagesForVersion_5_3_9104,
}

var defaultImagesForVersion_5_3_9104 = map[string]string{
	"alpine":                    "alpine-3.14:5.3.2@sha256:982220e0fd8ce55a73798fa7e814a482c4807c412f054c8440c5970b610239b7",
	"blobstore":                 "blobstore:5.3.2@sha256:d625be1eefe61cc42f94498e3c588bf212c4159c8b20c519db84eae4ff715efa",
	"cadvisor":                  "cadvisor:5.3.2@sha256:3860cce1f7ef0278c0d785f66baf69dd2bece19610a2fd6eaa54c03095f2f105",
	"codeinsights-db":           "codeinsights-db:5.3.2@sha256:c4a1bd3908658e1c09558a638e378e5570d5f669d27f9f867eeda25fe60cb88f",
	"codeintel-db":              "codeintel-db:5.3.2@sha256:1e0e93661a65c832b9697048c797f9894dfb502e2e1da2b8209f0018a6632b79",
	"gitserver":                 "gitserver:5.3.2@sha256:6c6042cf3e5f3f16de9b82e3d4ab1647f8bb924cd315245bd7a3162f5489e8c4",
	"pgsql":                     "postgres-12-alpine:5.3.2@sha256:1e0e93661a65c832b9697048c797f9894dfb502e2e1da2b8209f0018a6632b79",
	"pgsql-exporter":            "postgres_exporter:5.3.2@sha256:b9fa66fbcb4cc2d466487259db4ae2deacd7651dac4a9e28c9c7fc36523699d0",
	"precise-code-intel-worker": "precise-code-intel-worker:5.3.2@sha256:6142093097f5757afe772cffd131c1be54bb77335232011254733f51ffb2d6c6",
	"prometheus":                "prometheus:5.3.2@sha256:1b5c003fb39628f79e7655ba33f9ca119ddc4be021602ede3cc1674ef99fcdad",
	"redis-cache":               "redis-cache:5.3.2@sha256:ed79dada4d1a2bd85fb8450dffe227283ab6ae0e7ce56dc5056fbb8202d95624",
	"redis-exporter":            "redis_exporter:5.3.2@sha256:21a9dd9214483a42b11d58bf99e4f268f44257a4f67acd436d458797a31b7786",
	"redis-store":               "redis-store:5.3.2@sha256:0e3270a5eb293c158093f41145810eb5a154f61a74c9a896690dfdecd1b98b39",
	"repo-updater":              "repo-updater:5.3.2@sha256:5a414aa030c7e0922700664a43b449ee5f3fafa68834abef93988c5992c747c6",
	"symbols":                   "symbols:5.3.2@sha256:dd7f923bdbd5dbd231b749a7483110d40d59159084477b9fff84afaf58aad98e",
	"syntect-server":            "syntax-highlighter:5.3.2@sha256:3d16ab2a0203fea85063dcfe2e9d476540ef3274c28881dc4bbd5ca77933d8e8",
	"worker":                    "worker:5.3.2@sha256:776168bb53a0b094f51bfec3d0d38e2938a07bb840b665b645ccf2637f0e779f",
}

func GetDefaultImage(sg *Sourcegraph, component string) (string, error) {
	images, ok := defaultImages[sg.Spec.RequestedVersion]
	if !ok {
		return "", errors.Newf("no default images found for version %s", sg.Spec.RequestedVersion)
	}
	image, ok := images[component]
	if !ok {
		return "", errors.Newf("no default image found for service %s", component)
	}
	return fmt.Sprintf("%s/%s", sg.Spec.ImageRepository, image), nil
}
