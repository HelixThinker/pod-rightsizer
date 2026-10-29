package knowledge

import (
	"strings"
	"testing"

	coremetrics "github.com/nilx6874/pod-rightsizer/pkg/metrics"
	"github.com/nilx6874/pod-rightsizer/pkg/recommender"
)

func TestEvaluateUsesIndependentEvidenceCount(t *testing.T) {
	rawPolls := make([]coremetrics.ResourceMetrics, 10)
	advice := Evaluate(rawPolls, recommender.Recommendations{
		Observed: recommender.ObservedStatistics{IndependentSamples: 2},
	})
	if !strings.Contains(strings.Join(advice, " "), "Few metric samples") {
		t.Fatalf("advice = %#v, want low-evidence warning", advice)
	}
}
