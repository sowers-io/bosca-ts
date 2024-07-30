import {Metadata} from "next"

import {getAuthorizationToken} from "@/util/security";
import {redirect} from "next/navigation";

export const metadata: Metadata = {
  title: "Bosca",
  description: "Bosca: The AI Content Management System",
}

export default function Page() {
  if (!getAuthorizationToken()) {
    return redirect('/account/login')
  }
  return redirect('/dashboard')
}