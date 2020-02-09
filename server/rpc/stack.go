// Wiregost - Golang Exploitation Framework
// Copyright © 2020 Para
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package rpc

import (
	"strings"
	"time"

	"github.com/gogo/protobuf/proto"

	clientpb "github.com/maxlandon/wiregost/protobuf/client"
	"github.com/maxlandon/wiregost/server/module"
)

func rpcStackUse(data []byte, timeout time.Duration, resp RPCResponse) {
	stackReq := &clientpb.StackReq{}
	err := proto.Unmarshal(data, stackReq)
	if err != nil {
		resp(data, err)
	}

	// Find module
	path := strings.Join(stackReq.Path, "/")

	// Load it on the workspace stack
	wsID := uint(stackReq.WorkspaceID)
	stack := (*module.Stacks)[wsID]
	err = stack.LoadModule(path)

	// If errors, send status
	if err != nil {
		stackErr := &clientpb.Stack{
			Path: stackReq.Path,
			Err:  err.Error(),
		}

		data, err = proto.Marshal(stackErr)
		resp(data, err)
		return
	}

	// If no errors, send module
	mod := (*stack.Loaded)[path]
	module := []*clientpb.Module{mod.ToProtobuf()}
	stackUse := &clientpb.Stack{
		Path:    stackReq.Path,
		Modules: module,
		Err:     "",
	}

	data, err = proto.Marshal(stackUse)
	resp(data, err)
}

func rpcStackPop(data []byte, timeout time.Duration, resp RPCResponse) {
	stackReq := &clientpb.StackReq{}
	err := proto.Unmarshal(data, stackReq)
	if err != nil {
		resp(data, err)
	}

	// data, err := proto.Marshal()
	resp(data, err)
}

func rpcStackList(data []byte, timeout time.Duration, resp RPCResponse) {
	stackReq := &clientpb.StackReq{}
	err := proto.Unmarshal(data, stackReq)
	if err != nil {
		resp(data, err)
	}

	// data, err := proto.Marshal()
	resp(data, err)
}
