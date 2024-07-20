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
import { execute, toArrayBuffer } from '../../util/http'
import { Ollama } from '@langchain/community/llms/ollama'

import { ChatOpenAI } from '@langchain/openai'
import { ChatPromptTemplate } from '@langchain/core/prompts'
import { JsonOutputParser } from '@langchain/core/output_parsers'
import { WorkflowActivityModel } from '../../generated/protobuf/bosca/workflow/activities_pb'
import { BaseLanguageModel } from '@langchain/core/dist/language_models/base'
import { uploadSupplementary } from '../../util/uploader'

export class PromptActivity extends Activity {
  get id(): string {
    return 'ai.prompt'
  }

  private getModel(m: WorkflowActivityModel): BaseLanguageModel {
    let temperature: any = m.model!.configuration!.temperature
    if (temperature) {
      temperature = parseFloat(temperature)
    } else {
      temperature = 0
    }
    switch (m.model!.type) {
      case 'ollama-llm':
        return new Ollama({
          model: m.model!.name,
          temperature: temperature as number,
          format: 'json',
        })
      case 'openai-llm':
        return new ChatOpenAI({
          model: m.model!.name,
          temperature: temperature as number,
          apiKey: process.env.OPENAI_KEY,
          streaming: false,
        })
      default:
        throw new Error('unsupported model type: ' + m.model!.type)
    }
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
    const downloadResponse = await execute(downloadUrl)
    const payload = await downloadResponse.json()
    const m = activity.models[0]

    if (!m) throw new Error('missing model')

    const prompt = activity.prompts[0]

    if (!prompt) throw new Error('missing prompt')

    const model = this.getModel(m)
    const promptTemplate = ChatPromptTemplate.fromMessages([
      ['system', prompt.prompt!.systemPrompt],
      ['user', prompt.prompt!.userPrompt],
    ])
    const chain = promptTemplate.pipe(model).pipe(new JsonOutputParser())
    const response = await chain.invoke({ input: JSON.stringify(payload) })
    const output = JSON.stringify(response)
    await uploadSupplementary(
      activity.metadataId,
      'Prompt Response',
      'application/json',
      activity.activity!.outputs['supplementaryId'],
      source.id,
      undefined,
      toArrayBuffer(output)
    )
  }
}
