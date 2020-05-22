package exporter

import (
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Exporter struct {
	gauge_total_msgs prometheus.Gauge
	gauge_last_sent  prometheus.Gauge
	//gaugeVec prometheus.GaugeVec
}

func Run_Exporter_Server() {
	log.Println(`
  This is a prometheus exporter for nats-streaming
  Access: http://127.0.0.1:8081
  `)

	metricsPath := "/metrics"
	listenAddress := ":8081"
	metricsPrefix := "siang"
	exporters := NewExporter(metricsPrefix)
	/*
	   	registry := prometheus.NewRegistry()
	       registry.MustRegister(metrics)
	*/
	prometheus.MustRegister(exporters)

	// Launch http service

	http.Handle(metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
		 <head><title>Dummy Exporter</title></head>
		 <body>
		 <h1>Dummy Exporter</h1>
		 <p><a href='` + metricsPath + `'>Metrics</a></p>
		 </body>
		 </html>`))
	})
	log.Println(http.ListenAndServe(listenAddress, nil))
}

func NewExporter(metricsPrefix string) *Exporter {
	nats_total_msgs := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "nats_total_msgs",
		Help:      "This is a siang gauge metric"})
	nats_subscriptions_last_sent := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "nats_subscriptions_last_sent",
		Help:      "This is a siang gauge metric"})
	/*
		gaugeVec := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: metricsPrefix,
			Name:      "gauge_vec_metric",
			Help:      "This is a siang gauga vece metric"},
			[]string{"myLabel"})
	*/
	return &Exporter{
		//nats_total_msgs: gauge,
		gauge_total_msgs: nats_total_msgs,
		gauge_last_sent:  nats_subscriptions_last_sent,
		//gaugeVec: gaugeVec,
	}
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	//e.gauge.Set(float64(100))
	nats_ip_env := os.Getenv("NATS_IP")
	nats_port_env := os.Getenv("NATS_PORT")
	e.gauge_total_msgs.Set(float64(GetQueueworkerTotalMessage("http://" + nats_ip_env + ":" + nats_port_env + "/streaming/serverz")))
	e.gauge_last_sent.Set(float64(GetQueueworkerLastsent("http://"+nats_ip_env+":"+nats_port_env+"/streaming/serverz", "http://"+nats_ip_env+":"+nats_port_env+"/streaming/channelsz?channel=faas-request&subs=1")))
	//e.gaugeVec.WithLabelValues("hello").Set(float64(0))
	e.gauge_total_msgs.Collect(ch)
	e.gauge_last_sent.Collect(ch)
	//e.gaugeVec.Collect(ch)
}

// 讓exporter的prometheus屬性呼叫Describe方法

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.gauge_total_msgs.Describe(ch)
	e.gauge_last_sent.Describe(ch)
}
