// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package util

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/omec-project/amf/logger"
	"github.com/omec-project/nas/nasMessage"
	"github.com/omec-project/openapi/models"
)

func SnssaiHexToModels(hexString string) (*models.Snssai, error) {
	sst, err := strconv.ParseInt(hexString[:2], 16, 32)
	if err != nil {
		return nil, err
	}
	sNssai := models.Snssai{
		Sst: int32(sst),
		Sd:  hexString[2:],
	}
	return &sNssai, nil
}

func SnssaiModelsToHex(snssai models.Snssai) string {
	sst := fmt.Sprintf("%02x", snssai.Sst)
	return sst + snssai.Sd
}

func SeperateAmfId(amfid string) (regionId, setId, ptrId string, err error) {
	if len(amfid) != 6 {
		err = fmt.Errorf("len of amfId[%s] != 6", amfid)
		return
	}
	// regionId: 16bits, setId: 10bits, ptrId: 6bits
	regionId = amfid[:2]
	byteArray, err1 := hex.DecodeString(amfid[2:])
	if err1 != nil {
		err = err1
		return
	}
	byteSetId := []byte{byteArray[0] >> 6, byteArray[0]<<2 | byteArray[1]>>6}
	setId = hex.EncodeToString(byteSetId)[1:]
	bytePtrId := []byte{byteArray[1] & 0x3f}
	ptrId = hex.EncodeToString(bytePtrId)
	return
}

func PlmnIdStringToModels(plmnId string) (plmnID models.PlmnId) {
	plmnID.Mcc = plmnId[:3]
	plmnID.Mnc = plmnId[3:]
	return
}

func TACConfigToModels(intString string) (hexString string) {
	tmp, err := strconv.ParseUint(intString, 10, 32)
	if err != nil {
		logger.UtilLog.Errorf("ParseUint error: %+v", err)
		return
	}
	hexString = fmt.Sprintf("%06x", tmp)
	return
}

func AnTypeToNas(anType models.AccessType) uint8 {
	switch anType {
	case models.AccessType__3_GPP_ACCESS:
		return nasMessage.AccessType3GPP
	case models.AccessType_NON_3_GPP_ACCESS:
		return nasMessage.AccessTypeNon3GPP
	}

	return nasMessage.AccessTypeBoth
}
