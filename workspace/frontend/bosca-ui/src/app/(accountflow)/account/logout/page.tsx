import {getAuthorizationToken, getHost} from "@/util/security";
import {redirect} from "next/navigation";

export default async function Page() {
  const session = getAuthorizationToken()
  try {
    const data = await fetch(getHost() + '/self-service/logout/browser', {
      headers: {
        'Cookie': 'ory_kratos_session=' + session
      }
    })
    const json = await data.json()
    redirect(json.logout_url)
  } catch (e) {
    console.error(e)
    redirect('/account/login')
  }
  return (<div></div>)
}
