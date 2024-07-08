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
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/security/identity"
	"context"
)

func (ds *DataStore) SetMetadataWorkflowState(ctx context.Context, metadata *content.Metadata, toStateId string, status string, success bool, complete bool) error {
	subjectId, err := identity.GetSubjectId(ctx)
	if err != nil {
		return err
	}

	txn, err := ds.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = txn.ExecContext(ctx, "insert into metadata_workflow_transition_history (metadata_id, from_state_id, to_state_id, subject, status, success, complete) values ($1::uuid, $2, $3, $4, $5, $6, $7)", metadata.Id, metadata.WorkflowStateId, toStateId, subjectId, status, success, complete)
	if err != nil {
		txn.Rollback()
		return err
	}
	if !success {
		_, err = txn.ExecContext(ctx, "update metadata set workflow_state_pending_id = null where id = $1::uuid", metadata.Id)
		if err != nil {
			txn.Rollback()
			return err
		}
	} else {
		if complete {
			_, err = txn.ExecContext(ctx, "update metadata set workflow_state_id = $1, workflow_state_pending_id = null where id = $2::uuid", toStateId, metadata.Id)
			if err != nil {
				txn.Rollback()
				return err
			}
		} else {
			_, err = txn.ExecContext(ctx, "update metadata set workflow_state_pending_id = $1 where id = $2::uuid", toStateId, metadata.Id)
			if err != nil {
				txn.Rollback()
				return err
			}
		}
	}
	return txn.Commit()
}

