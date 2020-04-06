module github.com/pingcap/pd/v4

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/chzyer/logex v1.1.10 // indirect
	github.com/chzyer/readline v0.0.0-20171208011716-f6d7a1f6fbf3
	github.com/chzyer/test v0.0.0-20180213035817-a1ea475d72b1 // indirect
	github.com/coreos/go-semver v0.3.0
	github.com/coreos/pkg v0.0.0-20180108230652-97fdf19511ea
	github.com/docker/go-units v0.4.0
	github.com/eknkc/amber v0.0.0-20171010120322-cdade1c07385 // indirect
	github.com/elazarl/go-bindata-assetfs v1.0.0
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32
	github.com/go-openapi/spec v0.19.7 // indirect
	github.com/go-openapi/swag v0.19.8 // indirect
	github.com/go-playground/overalls v0.0.0-20180201144345-22ec1a223b7c
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.2
	github.com/google/btree v1.0.0
	github.com/gorilla/mux v1.7.3
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway v1.13.0
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/juju/ratelimit v1.0.1
	github.com/kevinburke/go-bindata v3.18.0+incompatible
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/mattn/go-shellwords v1.0.5
	github.com/mgechev/revive v1.0.2
	github.com/montanaflynn/stats v0.0.0-20151014174947-eeaced052adb
	github.com/opentracing/opentracing-go v1.1.0
	github.com/phf/go-queue v0.0.0-20170504031614-9abe38d0371d
	github.com/pingcap-incubator/tidb-dashboard v0.0.0-20200326180856-ee5b275f2b40
	github.com/pingcap/advanced-statefulset v0.3.2
	github.com/pingcap/check v0.0.0-20191216031241-8a5a85928f12
	github.com/pingcap/errcode v0.0.0-20180921232412-a1a7271709d9
	github.com/pingcap/errors v0.11.5-0.20190809092503-95897b64e011
	github.com/pingcap/failpoint v0.0.0-20191029060244-12f4ac2fd11d
	github.com/pingcap/kvproto v0.0.0-20200324130106-b8bc94dd8a36
	github.com/pingcap/log v0.0.0-20200117041106-d28c14d3b1cd
	github.com/pingcap/sysutil v0.0.0-20200309085538-962fd285f3bb
	github.com/pingcap/tidb-operator v1.1.0-rc.1
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.0.0
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	github.com/swaggo/http-swagger v0.0.0-20200308142732-58ac5e232fba
	github.com/swaggo/swag v1.6.6-0.20200323071853-8e21f4cefeea
	github.com/syndtr/goleveldb v0.0.0-20180815032940-ae2bd5eed72d
	github.com/unrolled/render v0.0.0-20171102162132-65450fb6b2d3
	github.com/urfave/negroni v1.0.0
	go.etcd.io/etcd v0.5.0-alpha.5.0.20191023171146-3cf2f69b5738
	go.uber.org/goleak v0.10.0
	go.uber.org/zap v1.13.0
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e // indirect
	golang.org/x/tools v0.0.0-20200325010219-a49f79bcc224
	google.golang.org/grpc v1.25.1
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	k8s.io/api v0.18.0
	k8s.io/apimachinery v0.18.0
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/klog v1.0.0
	k8s.io/utils v0.0.0-20200327001022-6496210b90e8
)

replace github.com/renstrom/dedent => github.com/lithammer/dedent v1.1.0

replace k8s.io/api => k8s.io/api v0.0.0-20190918155943-95b840bb6a1f

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190918161926-8f644eb6e783

replace k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190913080033-27d36303b655

replace k8s.io/apiserver => k8s.io/apiserver v0.0.0-20190918160949-bfa5e2e684ad

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20190918162238-f783a3654da8

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190918160344-1fbdaa4c8d90

replace k8s.io/code-generator => k8s.io/code-generator v0.0.0-20190912054826-cd179ad6a269

replace k8s.io/csi-api => k8s.io/csi-api v0.0.0-20190118125032-c4557c74373f

replace k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20190918161219-8c8f079fddc3

replace k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.0.0-20190918162944-7a93a0ddadd8

replace k8s.io/kube-proxy => k8s.io/kube-proxy v0.0.0-20190918162534-de037b596c1e

replace k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.0.0-20190918162820-3b5c1246eb18

replace k8s.io/kubelet => k8s.io/kubelet v0.0.0-20190918162654-250a1838aa2c

replace k8s.io/metrics => k8s.io/metrics v0.0.0-20190918162108-227c654b2546

replace k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.0.0-20190918161442-d4c9c65c82af

replace k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.0.0-20190918162410-e45c26d066f2

replace k8s.io/sample-controller => k8s.io/sample-controller v0.0.0-20190918161628-92eb3cb7496c

replace k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20190918163234-a9c1f33e9fb9

replace k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.0.0-20190918163108-da9fdfce26bb

replace k8s.io/component-base => k8s.io/component-base v0.0.0-20190918160511-547f6c5d7090

replace k8s.io/cri-api => k8s.io/cri-api v0.0.0-20190828162817-608eb1dad4ac

replace k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.0.0-20190918163402-db86a8c7bb21

replace k8s.io/kubectl => k8s.io/kubectl v0.0.0-20190918164019-21692a0861df

replace k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.0.0-20190918163543-cfa506e53441

replace k8s.io/node-api => k8s.io/node-api v0.0.0-20190918163711-2299658ad911

replace github.com/uber-go/atomic => github.com/uber-go/atomic v1.4.0

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.2.0+incompatible

replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2
