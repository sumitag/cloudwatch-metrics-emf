package metrics

import (
	"testing"
)

func TestCollectMetrics(t *testing.T) {
	metrics := CollectMetrics()
	if len(metrics.Metrics) == 0 {
		t.Errorf("expected non-empty metrics, got %v", metrics.Metrics)
	}
}

func TestToEMF(t *testing.T) {
	metrics := CollectMetrics()
	emf, err := metrics.ToEMF()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(emf) == 0 {
		t.Errorf("expected non-empty EMF string, got %v", emf)
	}
}
