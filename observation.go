package data

// TODO: Consider what package this belongs in

// Observation is a normalized structure which should be able to represent any
// format of GNSS Observable
type Observation struct {
	// Could put this into SatelliteData and have each constellation nested under
	// the same "Observation" which is unique for <Epoch + ReferenceStationId>
	Constellation string
	// This does not seem to be correctly implemented anywhere at the moment -
	// could use the station name instead (otherwise have the ID link to a table
	// of ID + station name)
	// Still can't assume that the ReferenceStationId from RTCM is correct
	ReferenceStationID uint16
	// TODO: normalize constellation epochs as timestamp
	Epoch uint32
	// This can be normalized to a type - spec says 0-4 is not applied, applied,
	// unknown, and reseverd
	ClockSteeringIndicator uint8
	// This can be normalized to a type - spec says 0-4 is internal, external
	// (locked), external (not locked), and unknown
	ExternalClockIndicator uint8
	// This could probably be normalized to SmoothingType - spec says true means
	// divergence-free smoothing and false means any other smoothing type
	SmoothingTypeIndicator bool
	// Could be normalized to seconds (or null for no smoothing)
	SmoothingInterval uint8
	SatelliteData     []SatelliteData
}

type SatelliteData struct {
	SatelliteID   int
	// Can probably be int or some time type
	RoughRangeMilliseconds uint8
	// This is specific for each constellation
	Extended uint8
	// Same comment as RangeMilliseconds
	RoughRanges     uint16
	PhaseRangeRates int16
	SignalData      []SignalData
}

type SignalData struct {
	SignalID        int
	Pseudoranges    int32
	PhaseRanges     int32
	PhaseRangeLocks uint16
	HalfCycles      bool
	CNRs            uint16
	PhaseRangeRates int16
}
