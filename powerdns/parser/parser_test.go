package parser

import (
	"testing"

	"github.com/czerwonk/testutils/assert"
)


func TestQtypes(t *testing.T) {
	data := "SOA\t7895\nTXT\t7912\n\n"

	q := ParseQtypes([]byte(data))

	assert.Int64Equal("qtype SOA", 7895, q["SOA"], t)
	assert.Int64Equal("qtype TXT", 7912, q["TXT"], t)
}

func TestRespsize(t *testing.T) {
	data := "20\t0\n40\t0\n60\t8013\n80\t7995\n100\t0\n150\t0\n200\t0\n400\t0\n600\t0\n800\t0\n1000\t0\n1200\t0\n1400\t0\n1600\t0\n1800\t0\n2000\t0\n2200\t0\n2400\t0\n2600\t0\n2800\t0\n3000\t0\n3200\t0\n3400\t0\n3600\t0\n3800\t0\n4000\t0\n4200\t0\n4400\t0\n4600\t0\n4800\t0\n5000\t0\n5200\t0\n5400\t0\n5600\t0\n5800\t0\n6000\t0\n6200\t0\n6400\t0\n6600\t0\n6800\t0\n7000\t0\n7200\t0\n7400\t0\n7600\t0\n7800\t0\n8000\t0\n8200\t0\n8400\t0\n8600\t0\n8800\t0\n9000\t0\n9200\t0\n9400\t0\n9600\t0\n9800\t0\n10000\t0\n10200\t0\n10400\t0\n10600\t0\n10800\t0\n11000\t0\n11200\t0\n11400\t0\n11600\t0\n11800\t0\n12000\t0\n12200\t0\n12400\t0\n12600\t0\n12800\t0\n13000\t0\n13200\t0\n13400\t0\n13600\t0\n13800\t0\n14000\t0\n14200\t0\n14400\t0\n14600\t0\n14800\t0\n15000\t0\n15200\t0\n15400\t0\n15600\t0\n15800\t0\n16000\t0\n16200\t0\n16400\t0\n16600\t0\n16800\t0\n17000\t0\n17200\t0\n17400\t0\n17600\t0\n17800\t0\n18000\t0\n18200\t0\n18400\t0\n18600\t0\n18800\t0\n19000\t0\n19200\t0\n19400\t0\n19600\t0\n19800\t0\n20000\t0\n20200\t0\n20400\t0\n20600\t0\n20800\t0\n21000\t0\n21200\t0\n21400\t0\n21600\t0\n21800\t0\n22000\t0\n22200\t0\n22400\t0\n22600\t0\n22800\t0\n23000\t0\n23200\t0\n23400\t0\n23600\t0\n23800\t0\n24000\t0\n24200\t0\n24400\t0\n24600\t0\n24800\t0\n25000\t0\n25200\t0\n25400\t0\n25600\t0\n25800\t0\n26000\t0\n26200\t0\n26400\t0\n26600\t0\n26800\t0\n27000\t0\n27200\t0\n27400\t0\n27600\t0\n27800\t0\n28000\t0\n28200\t0\n28400\t0\n28600\t0\n28800\t0\n29000\t0\n29200\t0\n29400\t0\n29600\t0\n29800\t0\n30000\t0\n30200\t0\n30400\t0\n30600\t0\n30800\t0\n31000\t0\n31200\t0\n31400\t0\n31600\t0\n31800\t0\n32000\t0\n32200\t0\n32400\t0\n32600\t0\n32800\t0\n33000\t0\n33200\t0\n33400\t0\n33600\t0\n33800\t0\n34000\t0\n34200\t0\n34400\t0\n34600\t0\n34800\t0\n35000\t0\n35200\t0\n35400\t0\n35600\t0\n35800\t0\n36000\t0\n36200\t0\n36400\t0\n36600\t0\n36800\t0\n37000\t0\n37200\t0\n37400\t0\n37600\t0\n37800\t0\n38000\t0\n38200\t0\n38400\t0\n38600\t0\n38800\t0\n39000\t0\n39200\t0\n39400\t0\n39600\t0\n39800\t0\n40000\t0\n40200\t0\n40400\t0\n40600\t0\n40800\t0\n41000\t0\n41200\t0\n41400\t0\n41600\t0\n41800\t0\n42000\t0\n42200\t0\n42400\t0\n42600\t0\n42800\t0\n43000\t0\n43200\t0\n43400\t0\n43600\t0\n43800\t0\n44000\t0\n44200\t0\n44400\t0\n44600\t0\n44800\t0\n45000\t0\n45200\t0\n45400\t0\n45600\t0\n45800\t0\n46000\t0\n46200\t0\n46400\t0\n46600\t0\n46800\t0\n47000\t0\n47200\t0\n47400\t0\n47600\t0\n47800\t0\n48000\t0\n48200\t0\n48400\t0\n48600\t0\n48800\t0\n49000\t0\n49200\t0\n49400\t0\n49600\t0\n49800\t0\n50000\t0\n50200\t0\n50400\t0\n50600\t0\n50800\t0\n51000\t0\n51200\t0\n51400\t0\n51600\t0\n51800\t0\n52000\t0\n52200\t0\n52400\t0\n52600\t0\n52800\t0\n53000\t0\n53200\t0\n53400\t0\n53600\t0\n53800\t0\n54000\t0\n54200\t0\n54400\t0\n54600\t0\n54800\t0\n55000\t0\n55200\t0\n55400\t0\n55600\t0\n55800\t0\n56000\t0\n56200\t0\n56400\t0\n56600\t0\n56800\t0\n57000\t0\n57200\t0\n57400\t0\n57600\t0\n57800\t0\n58000\t0\n58200\t0\n58400\t0\n58600\t0\n58800\t0\n59000\t0\n59200\t0\n59400\t0\n59600\t0\n59800\t0\n60000\t0\n60200\t0\n60400\t0\n60600\t0\n60800\t0\n61000\t0\n61200\t0\n61400\t0\n61600\t0\n61800\t0\n62000\t0\n62200\t0\n62400\t0\n62600\t0\n62800\t0\n63000\t0\n63200\t0\n63400\t0\n63600\t0\n63800\t0\n64000\t0\n64200\t0\n64400\t0\n64600\t0\n64800\t0\n65535\t0\n"

	rs := ParseRespsizes([]byte(data))

	assert.Int64Equal("response size", 8013, rs["60"], t)
	assert.Int64Equal("response size", 7995, rs["80"], t)
}

func TestStats(t *testing.T) {
	data := "corrupt-packets=0,deferred-cache-inserts=0,deferred-cache-lookup=0,deferred-packetcache-inserts=0,deferred-packetcache-lookup=0,dnsupdate-answers=0,dnsupdate-changes=0,dnsupdate-queries=0,dnsupdate-refused=0,incoming-notifications=0,overload-drops=0,packetcache-hit=65,packetcache-miss=15956,packetcache-size=2,query-cache-hit=8742,query-cache-miss=31076,query-cache-size=5,rd-queries=15956,recursing-answers=0,recursing-questions=0,recursion-unanswered=0,security-status=0,servfail-packets=0,signatures=0,tcp-answers=0,tcp-answers-bytes=0,tcp-queries=0,tcp4-answers=0,tcp4-answers-bytes=0,tcp4-queries=0,tcp6-answers=0,tcp6-answers-bytes=0,tcp6-queries=0,timedout-packets=0,udp-answers=16021,udp-answers-bytes=1209390,udp-do-queries=0,udp-queries=16021,udp4-answers=16021,udp4-answers-bytes=1209390,udp4-queries=16021,udp6-answers=0,udp6-answers-bytes=0,udp6-queries=0,fd-usage=61,key-cache-size=0,latency=2,meta-cache-size=1,qsize-q=0,real-memory-usage=11440128,signature-cache-size=0,sys-msec=58648,udp-in-errors=0,udp-noport-errors=505,udp-recvbuf-errors=0,udp-sndbuf-errors=0,uptime=270115,user-msec=4796,"

	s := ParseStats([]byte(data))

	assert.Int64Equal("stats udp-answers", 16021, s.UdpAnswers, t)
}
