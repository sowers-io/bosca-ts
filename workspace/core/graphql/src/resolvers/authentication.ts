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
import { Configuration, FrontendApi } from '@ory/kratos-client-fetch'
import { GraphQLRequestContext, logger, getAuthenticationToken } from '@bosca/common'

export const resolvers: Resolvers<GraphQLRequestContext> = {
  Mutation: {
    login: async (_, args) => {
      try {
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
      } catch (e) {
        logger.error({ error: e }, 'failed to login')
        throw e
      }
    },
    signup: async (_, args) => {
      try {
        const configuration = new Configuration({
          basePath: process.env.KRATOS_BASE_PATH,
        })
        const client = new FrontendApi(configuration)
        const loginFlow = await client.createNativeRegistrationFlow({})
        const updatedFlow = await client.updateRegistrationFlow({
          flow: loginFlow.id,
          updateRegistrationFlowBody: {
            method: 'password',
            traits: {
              email: args.email,
              firstName: args.firstName,
              lastName: args.lastName,
            },
            password: args.password,
          },
        })
        return updatedFlow.session_token || null
      } catch (e) {
        logger.error({ error: e }, 'failed to signup')
        throw e
      }
    },
    setPassword: async (_, args, context) => {
      try {
        const token = getAuthenticationToken(context)
        if (!token) return false 
        const configuration = new Configuration({
          basePath: process.env.KRATOS_BASE_PATH,
        })
        const client = new FrontendApi(configuration)
        const flow = await client.createNativeSettingsFlow({
          xSessionToken: token,
        })
        const updatedFlow = await client.updateSettingsFlow({
          flow: flow.id,
          xSessionToken: token,
          updateSettingsFlowBody: {
            method: 'password',
            password: args.password,
          },
        })
        return updatedFlow.state === 'success'
      } catch (e) {
        logger.error({ error: e }, 'failed to set password')
        throw e
      }
    },
  },
}
