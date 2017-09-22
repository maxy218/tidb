// Copyright 2017 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	log "github.com/Sirupsen/logrus"
)

type Status int32

const (
	Ongoing Status = iota
	Succeeded
	Failed
	Error
)

type response struct {
	data    string
	status  Status
	errCode uint16
}

type authHandler interface {
	handleStart(mechanism *string, data []byte, initial_response []byte) *response
	handleContinue(data []byte) *response
}

func (xa *XAuth) createAuthHandler(method string) authHandler {
	switch method {
	case "MYSQL41":
		return &saslMysql41Auth{
			m_state: S_starting,
			xauth:   xa,
		}
	//@TODO support in next pr
	case "PLAIN":
		//return &saslPlainAuth{}
		return nil
	default:
		log.Errorf("unknown x-protocol auth handler type [%s].", method)
		return nil
	}
}
