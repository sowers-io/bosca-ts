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

import { Resolvers } from '../generated/resolvers'
import { RequestContext } from '../context'
import { Configuration, FrontendApi } from '@ory/kratos-client-fetch'

export const resolvers: Resolvers<RequestContext> = {
  Mutation: {
    login: async (_, args) => {
      const configuration = new Configuration({
        basePath: process.env.KRATOS_BASE_PATH,
      })
      const client = new FrontendApi(configuration)
      const loginFlow = await client.createNativeLoginFlow({})
      const updatedFlow = await client.updateLoginFlow({
        flow: loginFlow.id,
        updateLoginFlowBody: {
          method: 'password',
          identifier: args.username,
          password: args.password,
          password_identifier: args.username,
        },
      })
      return updatedFlow.session_token || null
    },
  },
}
