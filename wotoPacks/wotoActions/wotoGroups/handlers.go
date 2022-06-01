package wotoGroups

import (
	we "wp-server/wotoPacks/core/wotoErrors"
	wv "wp-server/wotoPacks/core/wotoValues"
	"wp-server/wotoPacks/database/groupsDatabase"
	"wp-server/wotoPacks/interfaces"
	wa "wp-server/wotoPacks/wotoActions"
)

func HandleGroupsAction(req interfaces.ReqBase) error {
	batchValues := req.GetBatchValues()
	var err error
	var handler wv.ReqHandler

	for _, currentBatch := range batchValues {
		handler = _batchHandlers[currentBatch]
		if handler == nil {
			return wa.ErrInvalidBatch
		}

		err = handler(req)
		if err != nil {
			return err
		}
	}

	return req.LetExit()
}

func batchGetGroupInfoById(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginGetGroupInfoById)
	}

	var entryData = new(GetGroupInfoData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		return err
	}

	meta := req.GetMe().GetMeta()
	if meta != nil && !meta.GetBoolNoErr("can_get_group_info") {
		return we.SendPermissionDenied(req, OriginGetGroupInfoById)
	}

	if entryData.GroupId.IsInvalid() {
		return we.SendInvalidGroupId(req, OriginGetGroupInfoById)
	}

	info := groupsDatabase.GetGroupInfo(entryData.GroupId)
	if info == nil {
		return we.SendGroupNotFound(req, OriginGetGroupInfoById)
	}

	return req.SendResult(toGetGroupInfoResult(info))
	// return we.SendMethodNotImplemented(req, OriginGetGroupInfo)
}

func batchGetGroupCallInfo(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginGetGroupCallInfo)
	}

	// var entryData = new(SomethingData)
	// err := req.ParseJsonData(entryData)
	// if err != nil {
	// return err
	// }

	// doer := req.GetMe()
	// if doer != nil && !doer.CanCreateAccount() {
	// return we.SendAlreadyAuthorized(req, OriginRegisterUser)
	// }

	// return req.SendResult(nil)
	return we.SendMethodNotImplemented(req, "")
}

func batchCreateGroupCall(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginGetGroupCallInfo)
	}

	// var entryData = new(SomethingData)
	// err := req.ParseJsonData(entryData)
	// if err != nil {
	// return err
	// }

	// doer := req.GetMe()
	// if doer != nil && !doer.CanCreateAccount() {
	// return we.SendAlreadyAuthorized(req, OriginRegisterUser)
	// }

	// return req.SendResult(nil)
	return we.SendMethodNotImplemented(req, "")
}

func batchDeleteGroupCall(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginGetGroupCallInfo)
	}

	// var entryData = new(SomethingData)
	// err := req.ParseJsonData(entryData)
	// if err != nil {
	// return err
	// }

	// doer := req.GetMe()
	// if doer != nil && !doer.CanCreateAccount() {
	// return we.SendAlreadyAuthorized(req, OriginRegisterUser)
	// }

	// return req.SendResult(nil)
	return we.SendMethodNotImplemented(req, "")
}

func batchGetGroupCallQueue(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginGetGroupCallInfo)
	}

	// var entryData = new(SomethingData)
	// err := req.ParseJsonData(entryData)
	// if err != nil {
	// return err
	// }

	// doer := req.GetMe()
	// if doer != nil && !doer.CanCreateAccount() {
	// return we.SendAlreadyAuthorized(req, OriginRegisterUser)
	// }

	// return req.SendResult(nil)
	return we.SendMethodNotImplemented(req, "")
}

func batchGetGroupMediaQueue(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginGetGroupCallInfo)
	}

	// var entryData = new(SomethingData)
	// err := req.ParseJsonData(entryData)
	// if err != nil {
	// return err
	// }

	// doer := req.GetMe()
	// if doer != nil && !doer.CanCreateAccount() {
	// return we.SendAlreadyAuthorized(req, OriginRegisterUser)
	// }

	// return req.SendResult(nil)
	return we.SendMethodNotImplemented(req, "")
}

func batchGetGroupCallHistory(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginGetGroupCallInfo)
	}

	// var entryData = new(SomethingData)
	// err := req.ParseJsonData(entryData)
	// if err != nil {
	// return err
	// }

	// doer := req.GetMe()
	// if doer != nil && !doer.CanCreateAccount() {
	// return we.SendAlreadyAuthorized(req, OriginRegisterUser)
	// }

	// return req.SendResult(nil)
	return we.SendMethodNotImplemented(req, "")
}

func batchAddToQueue(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginGetGroupCallInfo)
	}

	// var entryData = new(SomethingData)
	// err := req.ParseJsonData(entryData)
	// if err != nil {
	// return err
	// }

	// doer := req.GetMe()
	// if doer != nil && !doer.CanCreateAccount() {
	// return we.SendAlreadyAuthorized(req, OriginRegisterUser)
	// }

	// return req.SendResult(nil)
	return we.SendMethodNotImplemented(req, "")
}

func batchAddMediaToQueue(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginGetGroupCallInfo)
	}

	// var entryData = new(SomethingData)
	// err := req.ParseJsonData(entryData)
	// if err != nil {
	// return err
	// }

	// doer := req.GetMe()
	// if doer != nil && !doer.CanCreateAccount() {
	// return we.SendAlreadyAuthorized(req, OriginRegisterUser)
	// }

	// return req.SendResult(nil)
	return we.SendMethodNotImplemented(req, "")
}
