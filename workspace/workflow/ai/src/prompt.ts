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

import {
  Activity,
  ActivityJobExecutor,
  execute,
  toArrayBuffer,
  uploadSupplementary,
} from '@bosca/workflow-activities-api'
import { Job } from 'bullmq'
import { ContentService, IdRequest, SupplementaryIdRequest, WorkflowActivityModel, WorkflowJob } from '@bosca/protobufs'
import {
  BaseLanguageModel,
  Ollama,
  ChatOpenAI,
  ChatGoogleGenerativeAI,
  ChatPromptTemplate,
  JsonOutputParser,
  getModel,
} from '@bosca/ai'

import { logger, useServiceAccountClient } from '@bosca/common'

export class PromptActivity extends Activity {
  get id(): string {
    return 'ai.prompt'
  }

  newJobExecutor(job: Job, workflowJob: WorkflowJob): ActivityJobExecutor<any> {
    return new Executor(this, job, workflowJob)
  }
}

class Executor extends ActivityJobExecutor<PromptActivity> {
  async execute() {
    const service = useServiceAccountClient(ContentService)
    const source = await service.getSource(new IdRequest({ id: 'workflow' }))
    const key = this.definition.supplementaryId
      ? this.definition.activity!.inputs!['supplementaryId'] + this.definition.supplementaryId
      : this.definition.activity!.inputs!['supplementaryId']
    const downloadUrl = await service.getMetadataSupplementaryDownloadUrl(
      new SupplementaryIdRequest({
        id: this.definition.metadataId,
        key: key,
      }),
    )
    let payload = (await execute(downloadUrl)).toString()

    const m = this.definition.models[0]

    if (!m) throw new Error('missing model')

    const prompt = this.definition.prompts[0]

    if (!prompt) throw new Error('missing prompt')

    const model = getModel(m)
    const promptTemplate = ChatPromptTemplate.fromMessages([
      ['system', prompt.prompt!.systemPrompt],
      ['user', prompt.prompt!.userPrompt],
    ])
    let chain = promptTemplate.pipe(model)
    if (prompt.prompt?.outputType == 'application/json') {
      chain = chain.pipe(new JsonOutputParser())
    }
    try {
      if (prompt.prompt?.inputType == 'application/json') {
        // ensure proper json
        const json = JSON.parse(payload)
        payload = JSON.stringify(json)
      }
      const response: any = await chain.invoke({ input: payload })
      logger.debug({ response: response }, 'got response from chain')

      const responseText = prompt.prompt?.outputType == 'application/json' ? JSON.stringify(response) : response.toString()
      await uploadSupplementary(
        this.definition.metadataId!,
        'Prompt Response',
        prompt.prompt?.outputType || 'text/plain',
        this.definition.supplementaryId
          ? this.definition.activity!.outputs['supplementaryId'] + this.definition.supplementaryId
          : this.definition.activity!.outputs['supplementaryId'],
        source.id,
        undefined,
        undefined,
        toArrayBuffer(responseText),
      )
    } catch (e) {
      logger.error({ error: e }, 'error in chain')
      throw e
    }
  }
}
