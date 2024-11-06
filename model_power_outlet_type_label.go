/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.4 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// PowerOutletTypeLabel the model 'PowerOutletTypeLabel'
type PowerOutletTypeLabel string

// List of PowerOutlet_type_label
const (
	POWEROUTLETTYPELABEL_C5 PowerOutletTypeLabel = "C5"
	POWEROUTLETTYPELABEL_C7 PowerOutletTypeLabel = "C7"
	POWEROUTLETTYPELABEL_C13 PowerOutletTypeLabel = "C13"
	POWEROUTLETTYPELABEL_C15 PowerOutletTypeLabel = "C15"
	POWEROUTLETTYPELABEL_C19 PowerOutletTypeLabel = "C19"
	POWEROUTLETTYPELABEL_C21 PowerOutletTypeLabel = "C21"
	POWEROUTLETTYPELABEL_PNE_4_H PowerOutletTypeLabel = "P+N+E 4H"
	POWEROUTLETTYPELABEL_PNE_6_H PowerOutletTypeLabel = "P+N+E 6H"
	POWEROUTLETTYPELABEL_PNE_9_H PowerOutletTypeLabel = "P+N+E 9H"
	POWEROUTLETTYPELABEL__2_PE_4_H PowerOutletTypeLabel = "2P+E 4H"
	POWEROUTLETTYPELABEL__2_PE_6_H PowerOutletTypeLabel = "2P+E 6H"
	POWEROUTLETTYPELABEL__2_PE_9_H PowerOutletTypeLabel = "2P+E 9H"
	POWEROUTLETTYPELABEL__3_PE_4_H PowerOutletTypeLabel = "3P+E 4H"
	POWEROUTLETTYPELABEL__3_PE_6_H PowerOutletTypeLabel = "3P+E 6H"
	POWEROUTLETTYPELABEL__3_PE_9_H PowerOutletTypeLabel = "3P+E 9H"
	POWEROUTLETTYPELABEL__3_PNE_4_H PowerOutletTypeLabel = "3P+N+E 4H"
	POWEROUTLETTYPELABEL__3_PNE_6_H PowerOutletTypeLabel = "3P+N+E 6H"
	POWEROUTLETTYPELABEL__3_PNE_9_H PowerOutletTypeLabel = "3P+N+E 9H"
	POWEROUTLETTYPELABEL_IEC_60906_1 PowerOutletTypeLabel = "IEC 60906-1"
	POWEROUTLETTYPELABEL__2_PT_10_A__NBR_14136 PowerOutletTypeLabel = "2P+T 10A (NBR 14136)"
	POWEROUTLETTYPELABEL__2_PT_20_A__NBR_14136 PowerOutletTypeLabel = "2P+T 20A (NBR 14136)"
	POWEROUTLETTYPELABEL_NEMA_1_15_R PowerOutletTypeLabel = "NEMA 1-15R"
	POWEROUTLETTYPELABEL_NEMA_5_15_R PowerOutletTypeLabel = "NEMA 5-15R"
	POWEROUTLETTYPELABEL_NEMA_5_20_R PowerOutletTypeLabel = "NEMA 5-20R"
	POWEROUTLETTYPELABEL_NEMA_5_30_R PowerOutletTypeLabel = "NEMA 5-30R"
	POWEROUTLETTYPELABEL_NEMA_5_50_R PowerOutletTypeLabel = "NEMA 5-50R"
	POWEROUTLETTYPELABEL_NEMA_6_15_R PowerOutletTypeLabel = "NEMA 6-15R"
	POWEROUTLETTYPELABEL_NEMA_6_20_R PowerOutletTypeLabel = "NEMA 6-20R"
	POWEROUTLETTYPELABEL_NEMA_6_30_R PowerOutletTypeLabel = "NEMA 6-30R"
	POWEROUTLETTYPELABEL_NEMA_6_50_R PowerOutletTypeLabel = "NEMA 6-50R"
	POWEROUTLETTYPELABEL_NEMA_10_30_R PowerOutletTypeLabel = "NEMA 10-30R"
	POWEROUTLETTYPELABEL_NEMA_10_50_R PowerOutletTypeLabel = "NEMA 10-50R"
	POWEROUTLETTYPELABEL_NEMA_14_20_R PowerOutletTypeLabel = "NEMA 14-20R"
	POWEROUTLETTYPELABEL_NEMA_14_30_R PowerOutletTypeLabel = "NEMA 14-30R"
	POWEROUTLETTYPELABEL_NEMA_14_50_R PowerOutletTypeLabel = "NEMA 14-50R"
	POWEROUTLETTYPELABEL_NEMA_14_60_R PowerOutletTypeLabel = "NEMA 14-60R"
	POWEROUTLETTYPELABEL_NEMA_15_15_R PowerOutletTypeLabel = "NEMA 15-15R"
	POWEROUTLETTYPELABEL_NEMA_15_20_R PowerOutletTypeLabel = "NEMA 15-20R"
	POWEROUTLETTYPELABEL_NEMA_15_30_R PowerOutletTypeLabel = "NEMA 15-30R"
	POWEROUTLETTYPELABEL_NEMA_15_50_R PowerOutletTypeLabel = "NEMA 15-50R"
	POWEROUTLETTYPELABEL_NEMA_15_60_R PowerOutletTypeLabel = "NEMA 15-60R"
	POWEROUTLETTYPELABEL_NEMA_L1_15_R PowerOutletTypeLabel = "NEMA L1-15R"
	POWEROUTLETTYPELABEL_NEMA_L5_15_R PowerOutletTypeLabel = "NEMA L5-15R"
	POWEROUTLETTYPELABEL_NEMA_L5_20_R PowerOutletTypeLabel = "NEMA L5-20R"
	POWEROUTLETTYPELABEL_NEMA_L5_30_R PowerOutletTypeLabel = "NEMA L5-30R"
	POWEROUTLETTYPELABEL_NEMA_L5_50_R PowerOutletTypeLabel = "NEMA L5-50R"
	POWEROUTLETTYPELABEL_NEMA_L6_15_R PowerOutletTypeLabel = "NEMA L6-15R"
	POWEROUTLETTYPELABEL_NEMA_L6_20_R PowerOutletTypeLabel = "NEMA L6-20R"
	POWEROUTLETTYPELABEL_NEMA_L6_30_R PowerOutletTypeLabel = "NEMA L6-30R"
	POWEROUTLETTYPELABEL_NEMA_L6_50_R PowerOutletTypeLabel = "NEMA L6-50R"
	POWEROUTLETTYPELABEL_NEMA_L10_30_R PowerOutletTypeLabel = "NEMA L10-30R"
	POWEROUTLETTYPELABEL_NEMA_L14_20_R PowerOutletTypeLabel = "NEMA L14-20R"
	POWEROUTLETTYPELABEL_NEMA_L14_30_R PowerOutletTypeLabel = "NEMA L14-30R"
	POWEROUTLETTYPELABEL_NEMA_L14_50_R PowerOutletTypeLabel = "NEMA L14-50R"
	POWEROUTLETTYPELABEL_NEMA_L14_60_R PowerOutletTypeLabel = "NEMA L14-60R"
	POWEROUTLETTYPELABEL_NEMA_L15_20_R PowerOutletTypeLabel = "NEMA L15-20R"
	POWEROUTLETTYPELABEL_NEMA_L15_30_R PowerOutletTypeLabel = "NEMA L15-30R"
	POWEROUTLETTYPELABEL_NEMA_L15_50_R PowerOutletTypeLabel = "NEMA L15-50R"
	POWEROUTLETTYPELABEL_NEMA_L15_60_R PowerOutletTypeLabel = "NEMA L15-60R"
	POWEROUTLETTYPELABEL_NEMA_L21_20_R PowerOutletTypeLabel = "NEMA L21-20R"
	POWEROUTLETTYPELABEL_NEMA_L21_30_R PowerOutletTypeLabel = "NEMA L21-30R"
	POWEROUTLETTYPELABEL_NEMA_L22_20_R PowerOutletTypeLabel = "NEMA L22-20R"
	POWEROUTLETTYPELABEL_NEMA_L22_30_R PowerOutletTypeLabel = "NEMA L22-30R"
	POWEROUTLETTYPELABEL_CS6360_C PowerOutletTypeLabel = "CS6360C"
	POWEROUTLETTYPELABEL_CS6364_C PowerOutletTypeLabel = "CS6364C"
	POWEROUTLETTYPELABEL_CS8164_C PowerOutletTypeLabel = "CS8164C"
	POWEROUTLETTYPELABEL_CS8264_C PowerOutletTypeLabel = "CS8264C"
	POWEROUTLETTYPELABEL_CS8364_C PowerOutletTypeLabel = "CS8364C"
	POWEROUTLETTYPELABEL_CS8464_C PowerOutletTypeLabel = "CS8464C"
	POWEROUTLETTYPELABEL_ITA_TYPE_E__CEE_7_5 PowerOutletTypeLabel = "ITA Type E (CEE 7/5)"
	POWEROUTLETTYPELABEL_ITA_TYPE_F__CEE_7_3 PowerOutletTypeLabel = "ITA Type F (CEE 7/3)"
	POWEROUTLETTYPELABEL_ITA_TYPE_G__BS_1363 PowerOutletTypeLabel = "ITA Type G (BS 1363)"
	POWEROUTLETTYPELABEL_ITA_TYPE_H PowerOutletTypeLabel = "ITA Type H"
	POWEROUTLETTYPELABEL_ITA_TYPE_I PowerOutletTypeLabel = "ITA Type I"
	POWEROUTLETTYPELABEL_ITA_TYPE_J PowerOutletTypeLabel = "ITA Type J"
	POWEROUTLETTYPELABEL_ITA_TYPE_K PowerOutletTypeLabel = "ITA Type K"
	POWEROUTLETTYPELABEL_ITA_TYPE_L__CEI_23_50 PowerOutletTypeLabel = "ITA Type L (CEI 23-50)"
	POWEROUTLETTYPELABEL_ITA_TYPE_M__BS_546 PowerOutletTypeLabel = "ITA Type M (BS 546)"
	POWEROUTLETTYPELABEL_ITA_TYPE_N PowerOutletTypeLabel = "ITA Type N"
	POWEROUTLETTYPELABEL_ITA_TYPE_O PowerOutletTypeLabel = "ITA Type O"
	POWEROUTLETTYPELABEL_ITA_MULTISTANDARD PowerOutletTypeLabel = "ITA Multistandard"
	POWEROUTLETTYPELABEL_USB_TYPE_A PowerOutletTypeLabel = "USB Type A"
	POWEROUTLETTYPELABEL_USB_MICRO_B PowerOutletTypeLabel = "USB Micro B"
	POWEROUTLETTYPELABEL_USB_TYPE_C PowerOutletTypeLabel = "USB Type C"
	POWEROUTLETTYPELABEL_MOLEX_MICRO_FIT_1X2 PowerOutletTypeLabel = "Molex Micro-Fit 1x2"
	POWEROUTLETTYPELABEL_MOLEX_MICRO_FIT_2X2 PowerOutletTypeLabel = "Molex Micro-Fit 2x2"
	POWEROUTLETTYPELABEL_MOLEX_MICRO_FIT_2X4 PowerOutletTypeLabel = "Molex Micro-Fit 2x4"
	POWEROUTLETTYPELABEL_DC_TERMINAL PowerOutletTypeLabel = "DC Terminal"
	POWEROUTLETTYPELABEL_EATON_C39 PowerOutletTypeLabel = "Eaton C39"
	POWEROUTLETTYPELABEL_HDOT_CX PowerOutletTypeLabel = "HDOT Cx"
	POWEROUTLETTYPELABEL_SAF_D_GRID PowerOutletTypeLabel = "Saf-D-Grid"
	POWEROUTLETTYPELABEL_NEUTRIK_POWER_CON__20_A PowerOutletTypeLabel = "Neutrik powerCON (20A)"
	POWEROUTLETTYPELABEL_NEUTRIK_POWER_CON__32_A PowerOutletTypeLabel = "Neutrik powerCON (32A)"
	POWEROUTLETTYPELABEL_NEUTRIK_POWER_CON_TRUE1 PowerOutletTypeLabel = "Neutrik powerCON TRUE1"
	POWEROUTLETTYPELABEL_NEUTRIK_POWER_CON_TRUE1_TOP PowerOutletTypeLabel = "Neutrik powerCON TRUE1 TOP"
	POWEROUTLETTYPELABEL_UBIQUITI_SMART_POWER PowerOutletTypeLabel = "Ubiquiti SmartPower"
	POWEROUTLETTYPELABEL_HARDWIRED PowerOutletTypeLabel = "Hardwired"
	POWEROUTLETTYPELABEL_OTHER PowerOutletTypeLabel = "Other"
)

// All allowed values of PowerOutletTypeLabel enum
var AllowedPowerOutletTypeLabelEnumValues = []PowerOutletTypeLabel{
	"C5",
	"C7",
	"C13",
	"C15",
	"C19",
	"C21",
	"P+N+E 4H",
	"P+N+E 6H",
	"P+N+E 9H",
	"2P+E 4H",
	"2P+E 6H",
	"2P+E 9H",
	"3P+E 4H",
	"3P+E 6H",
	"3P+E 9H",
	"3P+N+E 4H",
	"3P+N+E 6H",
	"3P+N+E 9H",
	"IEC 60906-1",
	"2P+T 10A (NBR 14136)",
	"2P+T 20A (NBR 14136)",
	"NEMA 1-15R",
	"NEMA 5-15R",
	"NEMA 5-20R",
	"NEMA 5-30R",
	"NEMA 5-50R",
	"NEMA 6-15R",
	"NEMA 6-20R",
	"NEMA 6-30R",
	"NEMA 6-50R",
	"NEMA 10-30R",
	"NEMA 10-50R",
	"NEMA 14-20R",
	"NEMA 14-30R",
	"NEMA 14-50R",
	"NEMA 14-60R",
	"NEMA 15-15R",
	"NEMA 15-20R",
	"NEMA 15-30R",
	"NEMA 15-50R",
	"NEMA 15-60R",
	"NEMA L1-15R",
	"NEMA L5-15R",
	"NEMA L5-20R",
	"NEMA L5-30R",
	"NEMA L5-50R",
	"NEMA L6-15R",
	"NEMA L6-20R",
	"NEMA L6-30R",
	"NEMA L6-50R",
	"NEMA L10-30R",
	"NEMA L14-20R",
	"NEMA L14-30R",
	"NEMA L14-50R",
	"NEMA L14-60R",
	"NEMA L15-20R",
	"NEMA L15-30R",
	"NEMA L15-50R",
	"NEMA L15-60R",
	"NEMA L21-20R",
	"NEMA L21-30R",
	"NEMA L22-20R",
	"NEMA L22-30R",
	"CS6360C",
	"CS6364C",
	"CS8164C",
	"CS8264C",
	"CS8364C",
	"CS8464C",
	"ITA Type E (CEE 7/5)",
	"ITA Type F (CEE 7/3)",
	"ITA Type G (BS 1363)",
	"ITA Type H",
	"ITA Type I",
	"ITA Type J",
	"ITA Type K",
	"ITA Type L (CEI 23-50)",
	"ITA Type M (BS 546)",
	"ITA Type N",
	"ITA Type O",
	"ITA Multistandard",
	"USB Type A",
	"USB Micro B",
	"USB Type C",
	"Molex Micro-Fit 1x2",
	"Molex Micro-Fit 2x2",
	"Molex Micro-Fit 2x4",
	"DC Terminal",
	"Eaton C39",
	"HDOT Cx",
	"Saf-D-Grid",
	"Neutrik powerCON (20A)",
	"Neutrik powerCON (32A)",
	"Neutrik powerCON TRUE1",
	"Neutrik powerCON TRUE1 TOP",
	"Ubiquiti SmartPower",
	"Hardwired",
	"Other",
}

func (v *PowerOutletTypeLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PowerOutletTypeLabel(value)
	for _, existing := range AllowedPowerOutletTypeLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PowerOutletTypeLabel", value)
}

// NewPowerOutletTypeLabelFromValue returns a pointer to a valid PowerOutletTypeLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPowerOutletTypeLabelFromValue(v string) (*PowerOutletTypeLabel, error) {
	ev := PowerOutletTypeLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PowerOutletTypeLabel: valid values are %v", v, AllowedPowerOutletTypeLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PowerOutletTypeLabel) IsValid() bool {
	for _, existing := range AllowedPowerOutletTypeLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PowerOutlet_type_label value
func (v PowerOutletTypeLabel) Ptr() *PowerOutletTypeLabel {
	return &v
}

type NullablePowerOutletTypeLabel struct {
	value *PowerOutletTypeLabel
	isSet bool
}

func (v NullablePowerOutletTypeLabel) Get() *PowerOutletTypeLabel {
	return v.value
}

func (v *NullablePowerOutletTypeLabel) Set(val *PowerOutletTypeLabel) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerOutletTypeLabel) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerOutletTypeLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerOutletTypeLabel(val *PowerOutletTypeLabel) *NullablePowerOutletTypeLabel {
	return &NullablePowerOutletTypeLabel{value: val, isSet: true}
}

func (v NullablePowerOutletTypeLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerOutletTypeLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

