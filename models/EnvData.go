/*
 * Copyright 2019 CJWW Development
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package models

import "time"
import "fmt"

type EnvData struct {
	CollectedAt  time.Time `firestore:"collectedAt"`
	Humidity     float64   `firestore:"humidity"`
	TemperatureC float64   `firestore:"temperature-c"`
	TemperatureF float64   `firestore:"temperature-f"`
}

func (model EnvData) Display() {
	collectedAt := fmt.Sprintf("Collected at: %s", model.CollectedAt)
	hum := fmt.Sprintf("Humidity: %d", model.Humidity)
	tempc := fmt.Sprintf("Temp c: %d", model.TemperatureC)
	tempf := fmt.Sprintf("Temp f: %d", model.TemperatureF)
	fmt.Println(collectedAt)
	fmt.Println(hum)
	fmt.Println(tempc)
	fmt.Println(tempf)
}