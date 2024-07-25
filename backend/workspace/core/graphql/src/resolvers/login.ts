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
