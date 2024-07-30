"use client"

import {Dashboard} from "@uppy/react";
import Uppy from "@uppy/core";
import Tus from "@uppy/tus";

import '@uppy/core/dist/style.min.css';
import '@uppy/dashboard/dist/style.min.css';
import '@/app/uppy.css';
import {useState} from "react";

interface UploadContentProps {
  collection?: string | null | undefined
}

export default function UploadContent({collection}: UploadContentProps) {
  const [uppy] = useState(() => {
    const uppy = new Uppy()
    uppy.use(Tus, {
      endpoint: 'http://localhost:8099/uploads',
      withCredentials: true
    }).on('file-added', (file) => {
      if (collection) {
        file.meta.collection = collection
      }
    }).on('complete', (result) => {
      window.dispatchEvent(new Event('files-changed'))
    })
    return uppy
  });

  return (
    <div>
      <Dashboard uppy={uppy} proudlyDisplayPoweredByUppy={false} width='100%' height="300px" style={{marginTop: '20px'}}/>
    </div>
  )
}
