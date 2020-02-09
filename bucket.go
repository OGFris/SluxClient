//          SluxClient  Copyright (C) 2020  C. Ilyes
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

package SluxClient

import "net/url"

func (s *SluxClient) CreateBucket(name string) (err error) {
	_, err = s.Do("/bucket", url.Values{
		"operation": {"create"},
		"bucket":    {name},
	})
	if err != nil {
		return
	}

	return
}

func (s *SluxClient) DeleteBucket(name string) (err error) {
	_, err = s.Do("/bucket", url.Values{
		"operation": {"delete"},
		"bucket":    {name},
	})
	if err != nil {
		return
	}

	return
}
