import Image from 'next/image'
import BoscaSvg from "../../../../public/bosca.svg"
import {getHost} from "@/util/security";

interface PageProps {
  searchParams: any
}

export default async function Page({searchParams}: PageProps) {
  let message = 'Sorry, there was an unexpected error.'

  if (searchParams.id) {
    message = (await (await fetch(getHost() + '/self-service/errors?id=' + searchParams.id)).json()).error.reason
  }

  return (
    <div className="flex min-h-full flex-1 flex-col justify-center py-12 sm:px-6 lg:px-8">
      <div className="sm:mx-auto sm:w-full sm:max-w-md items-center flex flex-col">
        <Image src={BoscaSvg} alt="Bosca" className="w-32 h-32" height="32" width="60"/>
        <h2 className="mt-6 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">
          Error
        </h2>
      </div>
      <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-[480px]">
        <div className="bg-white px-6 py-12 shadow sm:rounded-lg sm:px-12">
          <h3>{message}</h3>

          <div className="mt-10">
            <a href="/" className="rounded-md px-2 py-1.5 text-sm font-medium text-green-800 hover:bg-green-100 focus:outline-none focus:ring-2 focus:ring-green-600 focus:ring-offset-2 focus:ring-offset-green-50">
              Continue
            </a>
          </div>
        </div>
      </div>
    </div>
  )
}
