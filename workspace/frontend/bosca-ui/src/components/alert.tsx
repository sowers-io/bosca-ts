import {CheckCircleIcon, InformationCircleIcon, XCircleIcon} from "@heroicons/react/24/solid";
import {ReactNode} from "react";

interface AlertProps {
  title?: ReactNode
  body?: ReactNode
}

export function ErrorAlert({title, body}: AlertProps) {
  return (
    <div className="rounded-md bg-red-50 p-4">
      <div className="flex">
        <div className="flex-shrink-0">
          <XCircleIcon className="h-5 w-5 text-red-400" aria-hidden="true"/>
        </div>
        <div className="ml-3">
          <h3 className="text-sm font-medium text-red-800">
            {title}
          </h3>
          <div className="mt-2 text-sm text-red-700">
            {body}
          </div>
        </div>
      </div>
    </div>
  )
}

export function InfoAlert({title, body}: AlertProps) {
  return (
    <div className="rounded-md bg-blue-50 p-4">
      <div className="flex">
        <div className="flex-shrink-0">
          <InformationCircleIcon className="h-5 w-5 text-blue-400" aria-hidden="true"/>
        </div>
        <div className="ml-3">
          <h3 className="text-sm font-medium text-blue-800">
            {title}
          </h3>
          <div className="mt-2 text-sm text-blue-700">
            {body}
          </div>
        </div>
      </div>
    </div>
  )
}

export function SuccessAlert({title, body}: AlertProps) {
  return (
    <div className="rounded-md bg-green-50 p-4">
      <div className="flex">
        <div className="flex-shrink-0">
          <CheckCircleIcon className="h-5 w-5 text-green-400" aria-hidden="true"/>
        </div>
        <div className="ml-3">
          <h3 className="text-sm font-medium text-green-800">
            {title}
          </h3>
          <div className="mt-2 text-sm text-green-700">
            {body}
          </div>
        </div>
      </div>
    </div>
  )
}