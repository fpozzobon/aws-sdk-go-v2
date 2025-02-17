// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AlgorithmNameCloudRemoval string

// Enum values for AlgorithmNameCloudRemoval
const (
	// INTERPOLATION
	AlgorithmNameCloudRemovalInterpolation AlgorithmNameCloudRemoval = "INTERPOLATION"
)

// Values returns all known values for AlgorithmNameCloudRemoval. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (AlgorithmNameCloudRemoval) Values() []AlgorithmNameCloudRemoval {
	return []AlgorithmNameCloudRemoval{
		"INTERPOLATION",
	}
}

type AlgorithmNameGeoMosaic string

// Enum values for AlgorithmNameGeoMosaic
const (
	// NEAR
	AlgorithmNameGeoMosaicNear AlgorithmNameGeoMosaic = "NEAR"
	// BILINEAR
	AlgorithmNameGeoMosaicBilinear AlgorithmNameGeoMosaic = "BILINEAR"
	// CUBIC
	AlgorithmNameGeoMosaicCubic AlgorithmNameGeoMosaic = "CUBIC"
	// CUBICSPLINE
	AlgorithmNameGeoMosaicCubicspline AlgorithmNameGeoMosaic = "CUBICSPLINE"
	// LANCZOS
	AlgorithmNameGeoMosaicLanczos AlgorithmNameGeoMosaic = "LANCZOS"
	// AVERAGE
	AlgorithmNameGeoMosaicAverage AlgorithmNameGeoMosaic = "AVERAGE"
	// RMS
	AlgorithmNameGeoMosaicRms AlgorithmNameGeoMosaic = "RMS"
	// MODE
	AlgorithmNameGeoMosaicMode AlgorithmNameGeoMosaic = "MODE"
	// MAX
	AlgorithmNameGeoMosaicMax AlgorithmNameGeoMosaic = "MAX"
	// MIN
	AlgorithmNameGeoMosaicMin AlgorithmNameGeoMosaic = "MIN"
	// MED
	AlgorithmNameGeoMosaicMed AlgorithmNameGeoMosaic = "MED"
	// Q1
	AlgorithmNameGeoMosaicQ1 AlgorithmNameGeoMosaic = "Q1"
	// Q3
	AlgorithmNameGeoMosaicQ3 AlgorithmNameGeoMosaic = "Q3"
	// SUM
	AlgorithmNameGeoMosaicSum AlgorithmNameGeoMosaic = "SUM"
)

// Values returns all known values for AlgorithmNameGeoMosaic. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AlgorithmNameGeoMosaic) Values() []AlgorithmNameGeoMosaic {
	return []AlgorithmNameGeoMosaic{
		"NEAR",
		"BILINEAR",
		"CUBIC",
		"CUBICSPLINE",
		"LANCZOS",
		"AVERAGE",
		"RMS",
		"MODE",
		"MAX",
		"MIN",
		"MED",
		"Q1",
		"Q3",
		"SUM",
	}
}

type AlgorithmNameResampling string

// Enum values for AlgorithmNameResampling
const (
	// NEAR
	AlgorithmNameResamplingNear AlgorithmNameResampling = "NEAR"
	// BILINEAR
	AlgorithmNameResamplingBilinear AlgorithmNameResampling = "BILINEAR"
	// CUBIC
	AlgorithmNameResamplingCubic AlgorithmNameResampling = "CUBIC"
	// CUBICSPLINE
	AlgorithmNameResamplingCubicspline AlgorithmNameResampling = "CUBICSPLINE"
	// LANCZOS
	AlgorithmNameResamplingLanczos AlgorithmNameResampling = "LANCZOS"
	// AVERAGE
	AlgorithmNameResamplingAverage AlgorithmNameResampling = "AVERAGE"
	// RMS
	AlgorithmNameResamplingRms AlgorithmNameResampling = "RMS"
	// MODE
	AlgorithmNameResamplingMode AlgorithmNameResampling = "MODE"
	// MAX
	AlgorithmNameResamplingMax AlgorithmNameResampling = "MAX"
	// MIN
	AlgorithmNameResamplingMin AlgorithmNameResampling = "MIN"
	// MED
	AlgorithmNameResamplingMed AlgorithmNameResampling = "MED"
	// Q1
	AlgorithmNameResamplingQ1 AlgorithmNameResampling = "Q1"
	// Q3
	AlgorithmNameResamplingQ3 AlgorithmNameResampling = "Q3"
	// SUM
	AlgorithmNameResamplingSum AlgorithmNameResampling = "SUM"
)

// Values returns all known values for AlgorithmNameResampling. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AlgorithmNameResampling) Values() []AlgorithmNameResampling {
	return []AlgorithmNameResampling{
		"NEAR",
		"BILINEAR",
		"CUBIC",
		"CUBICSPLINE",
		"LANCZOS",
		"AVERAGE",
		"RMS",
		"MODE",
		"MAX",
		"MIN",
		"MED",
		"Q1",
		"Q3",
		"SUM",
	}
}

type ComparisonOperator string

// Enum values for ComparisonOperator
const (
	// EQUALS
	ComparisonOperatorEquals ComparisonOperator = "EQUALS"
	// NOT_EQUALS
	ComparisonOperatorNotEquals ComparisonOperator = "NOT_EQUALS"
	// STARTS_WITH
	ComparisonOperatorStartsWith ComparisonOperator = "STARTS_WITH"
)

// Values returns all known values for ComparisonOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ComparisonOperator) Values() []ComparisonOperator {
	return []ComparisonOperator{
		"EQUALS",
		"NOT_EQUALS",
		"STARTS_WITH",
	}
}

type DataCollectionType string

// Enum values for DataCollectionType
const (
	// PUBLIC
	DataCollectionTypePublic DataCollectionType = "PUBLIC"
	// PREMIUM
	DataCollectionTypePremium DataCollectionType = "PREMIUM"
	// USER
	DataCollectionTypeUser DataCollectionType = "USER"
)

// Values returns all known values for DataCollectionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataCollectionType) Values() []DataCollectionType {
	return []DataCollectionType{
		"PUBLIC",
		"PREMIUM",
		"USER",
	}
}

type EarthObservationJobErrorType string

// Enum values for EarthObservationJobErrorType
const (
	// CLIENT_ERROR
	EarthObservationJobErrorTypeClientError EarthObservationJobErrorType = "CLIENT_ERROR"
	// SERVER_ERROR
	EarthObservationJobErrorTypeServerError EarthObservationJobErrorType = "SERVER_ERROR"
)

// Values returns all known values for EarthObservationJobErrorType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (EarthObservationJobErrorType) Values() []EarthObservationJobErrorType {
	return []EarthObservationJobErrorType{
		"CLIENT_ERROR",
		"SERVER_ERROR",
	}
}

type EarthObservationJobExportStatus string

// Enum values for EarthObservationJobExportStatus
const (
	// IN_PROGRESS
	EarthObservationJobExportStatusInProgress EarthObservationJobExportStatus = "IN_PROGRESS"
	// SUCCEEDED
	EarthObservationJobExportStatusSucceeded EarthObservationJobExportStatus = "SUCCEEDED"
	// FAILED
	EarthObservationJobExportStatusFailed EarthObservationJobExportStatus = "FAILED"
)

// Values returns all known values for EarthObservationJobExportStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (EarthObservationJobExportStatus) Values() []EarthObservationJobExportStatus {
	return []EarthObservationJobExportStatus{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

type EarthObservationJobStatus string

// Enum values for EarthObservationJobStatus
const (
	// INITIALIZING
	EarthObservationJobStatusInitializing EarthObservationJobStatus = "INITIALIZING"
	// IN_PROGRESS
	EarthObservationJobStatusInProgress EarthObservationJobStatus = "IN_PROGRESS"
	// STOPPING
	EarthObservationJobStatusStopping EarthObservationJobStatus = "STOPPING"
	// COMPLETED
	EarthObservationJobStatusCompleted EarthObservationJobStatus = "COMPLETED"
	// STOPPED
	EarthObservationJobStatusStopped EarthObservationJobStatus = "STOPPED"
	// FAILED
	EarthObservationJobStatusFailed EarthObservationJobStatus = "FAILED"
	// DELETING
	EarthObservationJobStatusDeleting EarthObservationJobStatus = "DELETING"
	// DELETED
	EarthObservationJobStatusDeleted EarthObservationJobStatus = "DELETED"
)

// Values returns all known values for EarthObservationJobStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (EarthObservationJobStatus) Values() []EarthObservationJobStatus {
	return []EarthObservationJobStatus{
		"INITIALIZING",
		"IN_PROGRESS",
		"STOPPING",
		"COMPLETED",
		"STOPPED",
		"FAILED",
		"DELETING",
		"DELETED",
	}
}

type ExportErrorType string

// Enum values for ExportErrorType
const (
	// CLIENT_ERROR
	ExportErrorTypeClientError ExportErrorType = "CLIENT_ERROR"
	// SERVER_ERROR
	ExportErrorTypeServerError ExportErrorType = "SERVER_ERROR"
)

// Values returns all known values for ExportErrorType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExportErrorType) Values() []ExportErrorType {
	return []ExportErrorType{
		"CLIENT_ERROR",
		"SERVER_ERROR",
	}
}

type GroupBy string

// Enum values for GroupBy
const (
	// ALL
	GroupByAll GroupBy = "ALL"
	// YEARLY
	GroupByYearly GroupBy = "YEARLY"
)

// Values returns all known values for GroupBy. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (GroupBy) Values() []GroupBy {
	return []GroupBy{
		"ALL",
		"YEARLY",
	}
}

type LogicalOperator string

// Enum values for LogicalOperator
const (
	// AND
	LogicalOperatorAnd LogicalOperator = "AND"
)

// Values returns all known values for LogicalOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LogicalOperator) Values() []LogicalOperator {
	return []LogicalOperator{
		"AND",
	}
}

type MetadataProvider string

// Enum values for MetadataProvider
const (
	// PLANET_ORDER
	MetadataProviderPlanetOrder MetadataProvider = "PLANET_ORDER"
)

// Values returns all known values for MetadataProvider. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MetadataProvider) Values() []MetadataProvider {
	return []MetadataProvider{
		"PLANET_ORDER",
	}
}

type OutputType string

// Enum values for OutputType
const (
	// INT32
	OutputTypeInt32 OutputType = "INT32"
	// FLOAT32
	OutputTypeFloat32 OutputType = "FLOAT32"
	// INT16
	OutputTypeInt16 OutputType = "INT16"
	// FLOAT64
	OutputTypeFloat64 OutputType = "FLOAT64"
	// UINT16
	OutputTypeUint16 OutputType = "UINT16"
)

// Values returns all known values for OutputType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (OutputType) Values() []OutputType {
	return []OutputType{
		"INT32",
		"FLOAT32",
		"INT16",
		"FLOAT64",
		"UINT16",
	}
}

type PredefinedResolution string

// Enum values for PredefinedResolution
const (
	// HIGHEST
	PredefinedResolutionHighest PredefinedResolution = "HIGHEST"
	// LOWEST
	PredefinedResolutionLowest PredefinedResolution = "LOWEST"
	// AVERAGE
	PredefinedResolutionAverage PredefinedResolution = "AVERAGE"
)

// Values returns all known values for PredefinedResolution. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PredefinedResolution) Values() []PredefinedResolution {
	return []PredefinedResolution{
		"HIGHEST",
		"LOWEST",
		"AVERAGE",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	// ASCENDING
	SortOrderAscending SortOrder = "ASCENDING"
	// DESCENDING
	SortOrderDescending SortOrder = "DESCENDING"
)

// Values returns all known values for SortOrder. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"ASCENDING",
		"DESCENDING",
	}
}

type TargetOptions string

// Enum values for TargetOptions
const (
	// INPUT
	TargetOptionsInput TargetOptions = "INPUT"
	// OUTPUT
	TargetOptionsOutput TargetOptions = "OUTPUT"
)

// Values returns all known values for TargetOptions. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TargetOptions) Values() []TargetOptions {
	return []TargetOptions{
		"INPUT",
		"OUTPUT",
	}
}

type TemporalStatistics string

// Enum values for TemporalStatistics
const (
	// MEAN
	TemporalStatisticsMean TemporalStatistics = "MEAN"
	// MEDIAN
	TemporalStatisticsMedian TemporalStatistics = "MEDIAN"
	// STANDARD_DEVIATION
	TemporalStatisticsStandardDeviation TemporalStatistics = "STANDARD_DEVIATION"
)

// Values returns all known values for TemporalStatistics. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TemporalStatistics) Values() []TemporalStatistics {
	return []TemporalStatistics{
		"MEAN",
		"MEDIAN",
		"STANDARD_DEVIATION",
	}
}

type Unit string

// Enum values for Unit
const (
	// METERS
	UnitMeters Unit = "METERS"
)

// Values returns all known values for Unit. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Unit) Values() []Unit {
	return []Unit{
		"METERS",
	}
}

type VectorEnrichmentJobDocumentType string

// Enum values for VectorEnrichmentJobDocumentType
const (
	VectorEnrichmentJobDocumentTypeCsv VectorEnrichmentJobDocumentType = "CSV"
)

// Values returns all known values for VectorEnrichmentJobDocumentType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (VectorEnrichmentJobDocumentType) Values() []VectorEnrichmentJobDocumentType {
	return []VectorEnrichmentJobDocumentType{
		"CSV",
	}
}

type VectorEnrichmentJobErrorType string

// Enum values for VectorEnrichmentJobErrorType
const (
	// CLIENT_ERROR
	VectorEnrichmentJobErrorTypeClientError VectorEnrichmentJobErrorType = "CLIENT_ERROR"
	// SERVER_ERROR
	VectorEnrichmentJobErrorTypeServerError VectorEnrichmentJobErrorType = "SERVER_ERROR"
)

// Values returns all known values for VectorEnrichmentJobErrorType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (VectorEnrichmentJobErrorType) Values() []VectorEnrichmentJobErrorType {
	return []VectorEnrichmentJobErrorType{
		"CLIENT_ERROR",
		"SERVER_ERROR",
	}
}

type VectorEnrichmentJobExportErrorType string

// Enum values for VectorEnrichmentJobExportErrorType
const (
	VectorEnrichmentJobExportErrorTypeClientError VectorEnrichmentJobExportErrorType = "CLIENT_ERROR"
	VectorEnrichmentJobExportErrorTypeServerError VectorEnrichmentJobExportErrorType = "SERVER_ERROR"
)

// Values returns all known values for VectorEnrichmentJobExportErrorType. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (VectorEnrichmentJobExportErrorType) Values() []VectorEnrichmentJobExportErrorType {
	return []VectorEnrichmentJobExportErrorType{
		"CLIENT_ERROR",
		"SERVER_ERROR",
	}
}

type VectorEnrichmentJobExportStatus string

// Enum values for VectorEnrichmentJobExportStatus
const (
	VectorEnrichmentJobExportStatusInProgress VectorEnrichmentJobExportStatus = "IN_PROGRESS"
	VectorEnrichmentJobExportStatusSucceeded  VectorEnrichmentJobExportStatus = "SUCCEEDED"
	VectorEnrichmentJobExportStatusFailed     VectorEnrichmentJobExportStatus = "FAILED"
)

// Values returns all known values for VectorEnrichmentJobExportStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (VectorEnrichmentJobExportStatus) Values() []VectorEnrichmentJobExportStatus {
	return []VectorEnrichmentJobExportStatus{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

type VectorEnrichmentJobStatus string

// Enum values for VectorEnrichmentJobStatus
const (
	VectorEnrichmentJobStatusInitializing VectorEnrichmentJobStatus = "INITIALIZING"
	VectorEnrichmentJobStatusInProgress   VectorEnrichmentJobStatus = "IN_PROGRESS"
	VectorEnrichmentJobStatusStopping     VectorEnrichmentJobStatus = "STOPPING"
	VectorEnrichmentJobStatusStopped      VectorEnrichmentJobStatus = "STOPPED"
	VectorEnrichmentJobStatusCompleted    VectorEnrichmentJobStatus = "COMPLETED"
	VectorEnrichmentJobStatusFailed       VectorEnrichmentJobStatus = "FAILED"
	VectorEnrichmentJobStatusDeleting     VectorEnrichmentJobStatus = "DELETING"
	VectorEnrichmentJobStatusDeleted      VectorEnrichmentJobStatus = "DELETED"
)

// Values returns all known values for VectorEnrichmentJobStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (VectorEnrichmentJobStatus) Values() []VectorEnrichmentJobStatus {
	return []VectorEnrichmentJobStatus{
		"INITIALIZING",
		"IN_PROGRESS",
		"STOPPING",
		"STOPPED",
		"COMPLETED",
		"FAILED",
		"DELETING",
		"DELETED",
	}
}

type VectorEnrichmentJobType string

// Enum values for VectorEnrichmentJobType
const (
	VectorEnrichmentJobTypeReverseGeocoding VectorEnrichmentJobType = "REVERSE_GEOCODING"
	VectorEnrichmentJobTypeMapMatching      VectorEnrichmentJobType = "MAP_MATCHING"
)

// Values returns all known values for VectorEnrichmentJobType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VectorEnrichmentJobType) Values() []VectorEnrichmentJobType {
	return []VectorEnrichmentJobType{
		"REVERSE_GEOCODING",
		"MAP_MATCHING",
	}
}

type ZonalStatistics string

// Enum values for ZonalStatistics
const (
	// MEAN
	ZonalStatisticsMean ZonalStatistics = "MEAN"
	// MEDIAN
	ZonalStatisticsMedian ZonalStatistics = "MEDIAN"
	// STANDARD_DEVIATION
	ZonalStatisticsStandardDeviation ZonalStatistics = "STANDARD_DEVIATION"
	// MAX
	ZonalStatisticsMax ZonalStatistics = "MAX"
	// MIN
	ZonalStatisticsMin ZonalStatistics = "MIN"
	// SUM
	ZonalStatisticsSum ZonalStatistics = "SUM"
)

// Values returns all known values for ZonalStatistics. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ZonalStatistics) Values() []ZonalStatistics {
	return []ZonalStatistics{
		"MEAN",
		"MEDIAN",
		"STANDARD_DEVIATION",
		"MAX",
		"MIN",
		"SUM",
	}
}
