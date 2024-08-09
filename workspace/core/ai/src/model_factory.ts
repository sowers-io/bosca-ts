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

import { BaseLanguageModel, ChatGoogleGenerativeAI, ChatOpenAI, Ollama } from './index'
import { WorkflowActivityModel } from '@bosca/protobufs'

export function getModel(m: WorkflowActivityModel): BaseLanguageModel {
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
    case 'google-llm':
      return new ChatGoogleGenerativeAI({
        model: m.model!.name,
        temperature: temperature as number,
        apiKey: process.env.GOOGLE_API_KEY,
        streaming: false,
      })
    default:
      throw new Error('unsupported model type: ' + m.model!.type)
  }
}
