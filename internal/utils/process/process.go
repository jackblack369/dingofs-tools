/*
 *  Copyright (c) 2022 NetEase Inc.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

/*
 * Project: DingoCli
 * Created Date: 2022-06-21
 * Author: chengyi (Cyber-SiKu)
 */

package process

import (
	"log"
	"os"
)

type Cache struct {
	// bufs []*bytes.Buffer
	// mtx  *sync.RWMutex
}

func (c *Cache) Write(p []byte) (n int, err error) {
	// c.mtx.Lock()
	// defer c.mtx.Unlock()
	// c.bufs = append(c.bufs, bytes.NewBuffer(p))
	return len(p), nil
}

var C *Cache

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile | log.Lmicroseconds)
	C = &Cache{
		// bufs: make([]*bytes.Buffer, 0),
		// mtx:  &sync.RWMutex{},
	}
	log.SetOutput(C)
}

func SetShow(show bool) {
	if show {
		log.SetOutput(os.Stdout)
	} else {
		C = &Cache{}
		log.SetOutput(C)
	}
}
