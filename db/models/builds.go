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

package models

import serverpb "github.com/maxlandon/wiregost/proto/v1/gen/go/server"

// SELECT ----------------------------------------------------------------------

// GhostBuilds - Given a GhostBuild model, this function checks all fields and builds a
// query that may match one or more GhostBuilds, and returns them.
// Some fields will automatically return only one GhostBuild, such as ID
func GhostBuilds(GhostBuild serverpb.GhostBuild) (GhostBuilds []serverpb.GhostBuild) {
	return
}

// CREATE ----------------------------------------------------------------------

// AddGhostBuild - Add a GhostBuild to the database, if no existing GhostBuild matches it.
func AddGhostBuild(GhostBuild serverpb.GhostBuild) (added serverpb.GhostBuild) {
	return
}

// DELETE ----------------------------------------------------------------------

// DeleteGhostBuilds - Given a GhostBuild model, this function checks all fields and builds a
// query that may match one or more GhostBuilds, and deletes the ones found.
// For instance, if the GhostBuild has an ID, this one only will be deleted, but if it has
// several addresses, all matches will be deleted at once.
func DeleteGhostBuilds(GhostBuild serverpb.GhostBuild) (deleted []serverpb.GhostBuild) {
	return
}

// DeleteGhostBuildByID - Delete a GhostBuild given its ID
func DeleteGhostBuildByID(id uint32) (deleted bool) {
	return
}
