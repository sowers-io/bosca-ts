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

import { Activity } from '../activity'
import { WorkflowActivityJob } from '../../generated/protobuf/bosca/workflow/execution_context_pb'
import { useServiceClient } from '../../util/util'
import { ContentService } from '../../generated/protobuf/bosca/content/service_connect'
import { IdRequest, SupplementaryIdRequest } from '../../generated/protobuf/bosca/requests_pb'
import { execute } from '../../util/http'
import { Ollama } from "@langchain/community/llms/ollama";
import { uploadSupplementary } from '../../util/uploader'
import { HumanMessage, SystemMessage } from '@langchain/core/messages'

export class PromptActivity extends Activity {
  get id(): string {
    return 'ai.prompt'
  }

  async execute(activity: WorkflowActivityJob) {
    const service = useServiceClient(ContentService)
    const source = await service.getSource(new IdRequest({ id: 'workflow' }))
    const downloadUrl = await service.getMetadataSupplementaryDownloadUrl(
      new SupplementaryIdRequest({
        id: activity.metadataId,
        key: activity.activity!.inputs!['supplementaryId'],
      })
    )
    const response = await execute(downloadUrl)
    const markdown = await response.text()

    for (const model of activity.models) {
      const ollama = new Ollama({
        model: model.model!.name,
      })
      for (const prompt of activity.prompts) {
        const chatRequest = []
        if (prompt.prompt!.systemPrompt && prompt.prompt!.systemPrompt.length > 0) {
          chatRequest.push(new SystemMessage(prompt.prompt!.systemPrompt))
        }
        chatRequest.push(new HumanMessage(prompt.prompt!.userPrompt.replace('{table}', markdown)))
        console.log(prompt.prompt!.userPrompt.replace('{table}', markdown))
        const response = await ollama.invoke(chatRequest)
        console.log(response)
        // await uploadSupplementary(
        //   activity.metadataId,
        //   'AI Prompt Result',
        //   'text/plain',
        //   activity.activity!.outputs!['supplementaryId'],
        //   source.id,
        //   undefined,
        //   response.message.content
        // )
      }
    }
  }
}
