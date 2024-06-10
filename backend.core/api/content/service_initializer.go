/*
 * Copyright 2024 Sowers, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package content

import (
	"context"
	"log/slog"
	"os"
)

func initializeService(dataStore *DataStore) {
	ctx := context.Background()
	if added, err := dataStore.AddRootCollection(ctx); added {
		if err != nil {
			slog.Error("error initializing root collection: %v", slog.Any("error", err))
			os.Exit(1)
		}
	} else if err != nil {
		slog.Error("failed to initialize root collection permission", slog.Any("error", err))
		os.Exit(1)
	} else {
		slog.Info("root collection already initialized")
	}
}
