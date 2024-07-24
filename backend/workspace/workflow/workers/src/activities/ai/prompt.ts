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

import { Activity, ActivityJobExecutor } from '../activity'
import { useServiceClient } from '../../util/util'
import { execute, toArrayBuffer } from '../../util/http'
import { Ollama } from '@langchain/community/llms/ollama'

import { ChatOpenAI } from '@langchain/openai'
import { ChatPromptTemplate } from '@langchain/core/prompts'
import { JsonOutputParser } from '@langchain/core/output_parsers'
import { uploadSupplementary } from '../../util/uploader'
import { Job } from 'bullmq/dist/esm/classes/job'
import { ContentService, IdRequest, SupplementaryIdRequest, WorkflowActivityModel, WorkflowJob } from '@bosca/protobufs'
import { BaseLanguageModel } from '@langchain/core/language_models/base'

export class PromptActivity extends Activity {
  get id(): string {
    return 'ai.prompt'
  }

  newJobExecutor(job: Job, workflowJob: WorkflowJob): ActivityJobExecutor<any> {
    return new Executor(this, job, workflowJob)
  }
}

class Executor extends ActivityJobExecutor<PromptActivity> {
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

  async execute() {
    // const service = useServiceClient(ContentService)
    // const source = await service.getSource(new IdRequest({ id: 'workflow' }))
    // const downloadUrl = await service.getMetadataSupplementaryDownloadUrl(
    //   new SupplementaryIdRequest({
    //     id: this.definition.metadataId,
    //     key: this.definition.activity!.inputs!['supplementaryId'],
    //   })
    // )
    // const downloadResponse = await execute(downloadUrl)
    // const payload = await downloadResponse.json()
    // const m = this.definition.models[0]
    //
    // if (!m) throw new Error('missing model')
    //
    // const prompt = this.definition.prompts[0]
    //
    // if (!prompt) throw new Error('missing prompt')
    //
    // const model = this.getModel(m)
    // const promptTemplate = ChatPromptTemplate.fromMessages([
    //   ['system', prompt.prompt!.systemPrompt],
    //   ['user', prompt.prompt!.userPrompt],
    // ])
    // const chain = promptTemplate.pipe(model).pipe(new JsonOutputParser())
    // const response = await chain.invoke({ input: JSON.stringify(payload) })
    // const output = JSON.stringify(response)
    // await uploadSupplementary(
    //   this.definition.metadataId,
    //   'Prompt Response',
    //   'application/json',
    //   this.definition.activity!.outputs['supplementaryId'],
    //   source.id,
    //   undefined,
    //   toArrayBuffer(output)
    // )
  }
}
