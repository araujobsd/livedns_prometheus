module github.com/araujobsd/livedns_prometheus

go 1.18

require (
	github.com/alecthomas/template v0.0.0-20160405071501-a0175ee3bccc
	github.com/alecthomas/units v0.0.0-20151022065526-2efee857e7cf
	github.com/czerwonk/bird_exporter v0.0.0-20180824071309-d984aa91405c
	github.com/czerwonk/bird_socket v0.0.0-20180126164527-19dd9277893a
	github.com/czerwonk/testutils v0.0.0-20170526233935-dd9dabe360d4
	github.com/golang/protobuf v1.1.1-0.20180803194343-7d1b268556d6
	github.com/matttproud/golang_protobuf_extensions v1.0.1
	github.com/prometheus/client_golang v0.9.0-pre1.0.20180713201052-bcbbc08eb2dd
	github.com/prometheus/client_model v0.0.0-20180712105110-5c3871d89910
	github.com/prometheus/common v0.0.0-20180801064454-c7de2306084e
	github.com/prometheus/procfs v0.0.0-20180725123919-05ee40e3a273
	github.com/sirupsen/logrus v1.0.7-0.20180731161355-d329d24db431
	golang.org/x/crypto v0.0.0-20180802221240-56440b844dfe
	golang.org/x/sys v0.0.0-20180802203216-0ffbfd41fbef
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)

require github.com/beorn7/perks v1.0.1 // indirect
