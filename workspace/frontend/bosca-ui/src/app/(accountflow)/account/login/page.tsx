import Image from 'next/image'
import BoscaSvg from "../../../../../public/bosca.svg"
import {AuthForm} from "@/components/authForm"

interface PageProps {
  searchParams: any
}

export default async function Page({searchParams}: PageProps) {
  return (
    <div className="flex min-h-full flex-1 flex-col justify-center py-12 sm:px-6 lg:px-8">
      <div className="sm:mx-auto sm:w-full sm:max-w-md items-center flex flex-col">
        <Image src={BoscaSvg} alt="Bosca" className="w-32 h-32" height="32" width="60"/>
        <h2 className="mt-6 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">
          Sign in to your account
        </h2>
      </div>
      <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-[480px]">
        <div className="bg-white px-6 py-12 shadow sm:rounded-lg sm:px-12">
          <AuthForm flowType="login" flowId={searchParams.flow} returnTo={searchParams.returnTo} login={true}/>
        </div>
        <p className="mt-10 text-center text-sm text-gray-500">
          Not a member?{' '}
          <a href={searchParams.returnTo ? "/account/registration?returnTo=" + searchParams.returnTo : "/account/registration"} className="font-semibold leading-6 text-green-600 hover:text-green-500">
            Sign up
          </a>
        </p>
      </div>
    </div>
  )
}
