// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nas

import (
	"errors"
	"fmt"

	"github.com/omec-project/amf/context"
	"github.com/omec-project/amf/gmm"
	"github.com/omec-project/amf/logger"
	"github.com/omec-project/nas"
	"github.com/omec-project/openapi/models"
	"github.com/omec-project/util/fsm"
)

func Dispatch(ue *context.AmfUe, accessType models.AccessType, procedureCode int64, msg *nas.Message) error {
	if msg.GmmMessage == nil {
		return errors.New("Gmm Message is nil")
	}

	if msg.GsmMessage != nil {
		return errors.New("GSM Message should include in GMM Message")
	}

	if ue.State[accessType] == nil {
		return fmt.Errorf("UE State is empty (accessType=%q). Can't send GSM Message", accessType)
	}

	logger.ContextLog.Info("*** NAS dispatch state: ", ue.State[accessType])
	logger.ContextLog.Info("*** NAS dispatch gmm.ArgAccessType: ", accessType)
	logger.ContextLog.Info("*** NAS dispatch gmm.ArgProcedureCode: ", procedureCode)

	return gmm.GmmFSM.SendEvent(ue.State[accessType], gmm.GmmMessageEvent, fsm.ArgsType{
		gmm.ArgAmfUe:         ue,
		gmm.ArgAccessType:    accessType,
		gmm.ArgNASMessage:    msg.GmmMessage,
		gmm.ArgProcedureCode: procedureCode,
	})
}
