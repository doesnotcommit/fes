package fareestimator

import (
	"context"
	"fes"
	"fes/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Estimator(t *testing.T) {
	rrf := mocks.RawRideFetcherMock{
		FetchRawRideFunc: func() (fes.RawRide, bool, error) {
			return fes.RawRide{
				ID: "9",
				Points: []fes.RawPoint{
					{Lat: "37.953066", Lng: "23.735606", Timestamp: "1405587697"},
					{Lat: "37.953009", Lng: "23.735593", Timestamp: "1405587707"},
					{Lat: "37.953195", Lng: "23.736224", Timestamp: "1405587717"},
					{Lat: "37.953433", Lng: "23.736926", Timestamp: "1405587727"},
					{Lat: "37.953450", Lng: "23.737670", Timestamp: "1405587738"},
					{Lat: "37.953265", Lng: "23.738654", Timestamp: "1405587748"},
					{Lat: "37.953527", Lng: "23.738746", Timestamp: "1405587758"},
					{Lat: "37.953308", Lng: "23.738820", Timestamp: "1405587768"},
					{Lat: "37.953319", Lng: "23.738782", Timestamp: "1405587778"},
					{Lat: "37.953520", Lng: "23.738728", Timestamp: "1405587787"},
					{Lat: "37.953681", Lng: "23.739022", Timestamp: "1405587798"},
					{Lat: "37.954530", Lng: "23.739160", Timestamp: "1405587808"},
					{Lat: "37.956386", Lng: "23.738891", Timestamp: "1405587818"},
					{Lat: "37.957037", Lng: "23.738483", Timestamp: "1405587828"},
					{Lat: "37.956751", Lng: "23.738623", Timestamp: "1405587838"},
					{Lat: "37.957109", Lng: "23.738579", Timestamp: "1405587848"},
					{Lat: "37.958090", Lng: "23.738218", Timestamp: "1405587858"},
					{Lat: "37.958981", Lng: "23.737785", Timestamp: "1405587868"},
					{Lat: "37.959762", Lng: "23.737327", Timestamp: "1405587878"},
					{Lat: "37.960813", Lng: "23.736862", Timestamp: "1405587888"},
					{Lat: "37.962113", Lng: "23.735860", Timestamp: "1405587898"},
					{Lat: "37.963503", Lng: "23.734914", Timestamp: "1405587908"},
					{Lat: "37.964694", Lng: "23.733980", Timestamp: "1405587918"},
					{Lat: "37.965675", Lng: "23.733457", Timestamp: "1405587928"},
					{Lat: "37.965588", Lng: "23.733301", Timestamp: "1405587938"},
					{Lat: "37.965579", Lng: "23.733270", Timestamp: "1405587948"},
					{Lat: "37.965566", Lng: "23.733313", Timestamp: "1405587958"},
					{Lat: "37.965564", Lng: "23.733315", Timestamp: "1405587968"},
					{Lat: "37.965563", Lng: "23.733322", Timestamp: "1405587978"},
					{Lat: "37.965526", Lng: "23.733276", Timestamp: "1405587989"},
					{Lat: "37.965432", Lng: "23.733125", Timestamp: "1405587999"},
					{Lat: "37.965294", Lng: "23.732922", Timestamp: "1405588009"},
					{Lat: "37.965330", Lng: "23.732851", Timestamp: "1405588019"},
					{Lat: "37.965344", Lng: "23.732843", Timestamp: "1405588029"},
					{Lat: "37.965293", Lng: "23.732778", Timestamp: "1405588039"},
					{Lat: "37.965296", Lng: "23.732795", Timestamp: "1405588049"},
					{Lat: "37.965296", Lng: "23.732795", Timestamp: "1405588059"},
					{Lat: "37.965296", Lng: "23.732795", Timestamp: "1405588069"},
					{Lat: "37.965296", Lng: "23.732795", Timestamp: "1405588079"},
					{Lat: "37.965296", Lng: "23.732795", Timestamp: "1405588089"},
					{Lat: "37.965296", Lng: "23.732795", Timestamp: "1405588099"},
					{Lat: "37.965345", Lng: "23.732741", Timestamp: "1405588109"},
					{Lat: "37.965360", Lng: "23.732716", Timestamp: "1405588119"},
					{Lat: "37.965359", Lng: "23.732706", Timestamp: "1405588129"},
					{Lat: "37.965362", Lng: "23.732707", Timestamp: "1405588139"},
					{Lat: "37.965363", Lng: "23.732677", Timestamp: "1405588150"},
					{Lat: "37.965369", Lng: "23.732666", Timestamp: "1405588160"},
					{Lat: "37.965367", Lng: "23.732665", Timestamp: "1405588170"},
					{Lat: "37.965366", Lng: "23.732664", Timestamp: "1405588180"},
					{Lat: "37.965366", Lng: "23.732664", Timestamp: "1405588190"},
					{Lat: "37.965367", Lng: "23.732664", Timestamp: "1405588200"},
					{Lat: "37.965367", Lng: "23.732664", Timestamp: "1405588210"},
					{Lat: "37.965367", Lng: "23.732664", Timestamp: "1405588220"},
					{Lat: "37.965367", Lng: "23.732664", Timestamp: "1405588230"},
					{Lat: "37.965367", Lng: "23.732664", Timestamp: "1405588240"},
					{Lat: "37.965367", Lng: "23.732664", Timestamp: "1405588250"},
					{Lat: "37.965367", Lng: "23.732664", Timestamp: "1405588260"},
					{Lat: "37.965367", Lng: "23.732664", Timestamp: "1405588270"},
					{Lat: "37.965367", Lng: "23.732664", Timestamp: "1405588280"},
					{Lat: "37.965367", Lng: "23.732664", Timestamp: "1405588290"},
					{Lat: "37.965367", Lng: "23.732664", Timestamp: "1405588300"},
					{Lat: "37.965367", Lng: "23.732664", Timestamp: "1405588310"},
					{Lat: "37.965367", Lng: "23.732664", Timestamp: "1405588320"},
					{Lat: "37.965367", Lng: "23.732664", Timestamp: "1405588330"},
					{Lat: "37.965366", Lng: "23.732662", Timestamp: "1405588340"},
					{Lat: "37.966104", Lng: "23.731993", Timestamp: "1405588350"},
					{Lat: "37.966193", Lng: "23.731925", Timestamp: "1405588360"},
					{Lat: "37.966014", Lng: "23.731937", Timestamp: "1405588370"},
					{Lat: "37.966100", Lng: "23.731828", Timestamp: "1405588380"},
					{Lat: "37.966136", Lng: "23.731916", Timestamp: "1405588390"},
					{Lat: "37.965943", Lng: "23.732123", Timestamp: "1405588400"},
					{Lat: "37.965927", Lng: "23.732022", Timestamp: "1405588410"},
					{Lat: "37.966823", Lng: "23.731732", Timestamp: "1405588420"},
					{Lat: "37.967018", Lng: "23.731652", Timestamp: "1405588430"},
					{Lat: "37.966201", Lng: "23.731989", Timestamp: "1405588441"},
					{Lat: "37.966027", Lng: "23.731923", Timestamp: "1405588451"},
					{Lat: "37.963705", Lng: "23.732530", Timestamp: "1405588461"},
					{Lat: "37.963705", Lng: "23.732530", Timestamp: "1405588471"},
					{Lat: "37.963705", Lng: "23.732530", Timestamp: "1405588481"},
					{Lat: "37.963705", Lng: "23.732530", Timestamp: "1405588491"},
					{Lat: "37.963705", Lng: "23.732530", Timestamp: "1405588501"},
					{Lat: "37.963705", Lng: "23.732530", Timestamp: "1405588511"},
					{Lat: "37.963705", Lng: "23.732530", Timestamp: "1405588521"},
					{Lat: "37.968053", Lng: "23.730544", Timestamp: "1405588531"},
					{Lat: "37.967438", Lng: "23.730404", Timestamp: "1405588541"},
					{Lat: "37.967376", Lng: "23.730307", Timestamp: "1405588551"},
					{Lat: "37.967349", Lng: "23.730243", Timestamp: "1405588561"},
					{Lat: "37.967349", Lng: "23.730235", Timestamp: "1405588571"},
					{Lat: "37.967349", Lng: "23.730235", Timestamp: "1405588581"},
					{Lat: "37.967348", Lng: "23.730235", Timestamp: "1405588591"},
					{Lat: "37.967345", Lng: "23.730240", Timestamp: "1405588601"},
					{Lat: "37.967445", Lng: "23.729832", Timestamp: "1405588611"},
					{Lat: "37.967666", Lng: "23.729056", Timestamp: "1405588621"},
					{Lat: "37.967987", Lng: "23.728256", Timestamp: "1405588632"},
					{Lat: "37.968190", Lng: "23.727493", Timestamp: "1405588641"},
					{Lat: "37.968579", Lng: "23.726776", Timestamp: "1405588651"},
					{Lat: "37.968840", Lng: "23.725827", Timestamp: "1405588661"},
					{Lat: "37.969206", Lng: "23.724863", Timestamp: "1405588671"},
					{Lat: "37.969592", Lng: "23.723868", Timestamp: "1405588682"},
					{Lat: "37.969731", Lng: "23.723192", Timestamp: "1405588692"},
					{Lat: "37.969754", Lng: "23.723101", Timestamp: "1405588701"},
					{Lat: "37.969754", Lng: "23.723109", Timestamp: "1405588711"},
					{Lat: "37.969751", Lng: "23.723106", Timestamp: "1405588722"},
					{Lat: "37.969752", Lng: "23.723106", Timestamp: "1405588732"},
					{Lat: "37.969752", Lng: "23.723105", Timestamp: "1405588742"},
					{Lat: "37.969873", Lng: "23.722723", Timestamp: "1405588752"},
					{Lat: "37.969949", Lng: "23.722461", Timestamp: "1405588762"},
					{Lat: "37.969687", Lng: "23.722401", Timestamp: "1405588772"},
					{Lat: "37.969647", Lng: "23.722431", Timestamp: "1405588782"},
				},
			}, false, nil
		},
	}
	few := mocks.FareEstimateWriterMock{
		WriteFareEstimateFunc: func(id int, estimate float64) error {
			return nil
		},
	}
	fe := New(&rrf, &few, 10)
	ctx := context.Background()
	if err := fe.GetFareEstimates(ctx); err != nil {
		t.Fatal(err)
	}
	wantWriteCalls := []struct {
		ID       int
		Estimate float64
	}{
		{
			ID:       9,
			Estimate: 6.3471241655565,
		},
	}
	assert.Equal(t, wantWriteCalls, few.WriteFareEstimateCalls(), "write fare estimate calls")
}
