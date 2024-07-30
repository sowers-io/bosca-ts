import {headers} from "next/headers";
import {redirect} from "next/navigation";
import {Messages} from "@/components/messages";
import {getHost} from "@/util/security";

interface AuthFormProps {
  flowType: string
  flowId: string | null
  returnTo: string | null
  login?: boolean
}

async function getData(flowType: string, flowId: string) {
  const hdrs = new Headers()
  headers().forEach((value, key) => {
    if (key !== 'cookie') return
    hdrs.append(key, value)
  })
  hdrs.set('host', 'accounts.bosca.io')
  hdrs.set('accept', 'application/json')
  hdrs.set('referer', 'https://login.bosca.io/' + flowType)
  const flow = await fetch(getHost() + '/self-service/' + flowType + '/flows?id=' + flowId, {
    headers: hdrs
  })
  return await flow.json()
}

const fieldOrder = ["traits.firstName", "firstName", "traits.lastName", "lastName", "traits.email", "identifier", "password"]

export async function AuthForm({flowType, flowId, returnTo, login}: AuthFormProps) {
  if (!flowId) {
    if (returnTo) {
      const r = encodeURI(returnTo)
      return redirect(getHost() + '/self-service/' + flowType + '/browser?return_to=' + r)
    }
    return redirect(getHost() + '/self-service/' + flowType + '/browser')
  }

  const data = await getData(flowType, flowId)

  console.log(data)

  if (!data.ui) {
    return redirect('/error')
  }

  const fields = []

  if (data.ui.messages && data.ui.messages.length > 0) {
    for (const message of data.ui.messages) {
      console.log(message)
      if (message.id === 4000010) {
        return redirect('/account/verification')
      }
    }
    fields.push(<Messages messages={data.ui.messages}/>)
  }

  data.ui.nodes.sort((a: any, b: any) => {
    const aIndex = fieldOrder.indexOf(a.attributes.name);
    const bIndex = fieldOrder.indexOf(b.attributes.name);
    if (aIndex === -1) return 1;
    if (bIndex === -1) return -1;
    return aIndex - bIndex;
  });

  for (const node of data.ui.nodes) {
    if (node.messages && node.messages.length > 0) {
      fields.push(<Messages messages={node.messages}/>)
    }
    if (node.attributes.type === "hidden") {
      fields.push(
        <input key={node.attributes.name} type="hidden" name={node.attributes.name} defaultValue={node.attributes.value}/>
      )
    } else if (node.attributes.type === "submit") {
      if (login) {
        fields.push(
          <div key="__remember__me__" className="flex items-center justify-between">
            <div className="flex items-center">
            </div>
            <div className="text-sm leading-6">
              <a href={returnTo ? "/account/recovery?returnTo=" + returnTo : "/account/recovery"} className="font-semibold text-green-600 hover:text-green-500">
                Forgot password?
              </a>
            </div>
          </div>
        )
      }
      fields.push(
        <div key={"key-" + node.meta.label.id.toString()}>
          <button
            type={node.attributes.type}
            id={node.meta.label.id.toString()}
            name={node.attributes.name}
            value={node.attributes.value}
            className="flex w-full justify-center rounded-md bg-green-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-green-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-green-600"
          >
            {node.meta.label.text}
          </button>
        </div>
      )
    } else {
      if (node.type === 'text') {
        fields.push(
          <div key={"key-" + node.meta.label.id.toString()} className="mt-2">
            {node.meta.label.text}
          </div>
        )
      } else if (node.type === 'a') {
        fields.push(
          <div key={"key-" + node.meta.label.id.toString()} className="mt-2">
            <a
              id={node.meta.label.id.toString()}
              href={node.attributes.href}
              className="rounded-md px-2 py-1.5 text-sm font-medium text-green-800 hover:bg-green-100 focus:outline-none focus:ring-2 focus:ring-green-600 focus:ring-offset-2 focus:ring-offset-green-50"
            >
              {node.meta.label.text}
            </a>
          </div>
        )
      } else {
        fields.push(
          <div key={"key-" + node.meta.label.id.toString()}>
            <label htmlFor={node.meta.label.id.toString()} className="block text-sm font-medium leading-6 text-gray-900">
              {node.meta.label.text}
            </label>
            <div className="mt-2">
              <input
                id={node.meta.label.id.toString()}
                name={node.attributes.name}
                type={node.attributes.type}
                required={node.attributes.required}
                defaultValue={node.attributes.value}
                className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-green-600 sm:text-sm sm:leading-6"
              />
            </div>
          </div>
        )
      }
    }
  }
  return (
    <form className="space-y-6" action={data.ui.action} method={data.ui.method} encType="application/x-www-form-urlencoded">
      {fields}
    </form>
  )
}